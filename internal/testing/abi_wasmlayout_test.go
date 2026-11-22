package testing

import (
	"fmt"
	"go/types"
	"reflect"
	"sort"
	"strings"
	"testing"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// Test_WasmLayout compares the C layout internal.CLayout derives from every Go
// struct against the emscripten (wasm32) layout of the C struct it mirrors.
// Pointers, size_t and long shrink to 4 bytes there, so a Go field type that
// only matches C on 64-bit hosts shows up here and nowhere else.
func Test_WasmLayout(t *testing.T) {
	for _, library := range libraries {
		t.Run(library, func(t *testing.T) {
			cfg := readJSON[ffiConfig](t, assetPath("config", library))
			entries := readJSON[[]ffiEntry](t, assetPath("ffi", library))
			goStructs := structTypes(t, modulePath+"/"+cfg.OutDir)
			c := newWasm32(entries)

			var checked int
			var untested []string
			for _, e := range entries {
				if e.Tag != "struct" || len(e.Fields) == 0 || !cfg.includes(e.Location) {
					continue
				}

				base := strings.TrimPrefix(e.Name, cfg.Prefix)
				var compared bool
				for _, name := range []string{base, unexport(base)} {
					s, ok := goStructs[name]
					if !ok || s.NumFields() != len(e.Fields) {
						continue
					}
					rt, ok := reflectType(s)
					if !ok {
						continue
					}
					compared = true

					size, offsets := internal.CLayout(rt)
					wantSize, _, wantOffsets := c.layout(e)
					if int64(size) != wantSize {
						t.Errorf("%s (Go %s): wasm32 size %d, C says %d", e.Name, name, size, wantSize)
						continue
					}
					for i, f := range e.Fields {
						if int64(offsets[i]) != wantOffsets[i] {
							t.Errorf("%s.%s (Go %s.%s): wasm32 offset %d, C says %d",
								e.Name, f.Name, name, s.Field(i).Name(), offsets[i], wantOffsets[i])
						}
					}
				}
				if compared {
					checked++
				} else {
					untested = append(untested, e.Name)
				}
			}

			sort.Strings(untested)
			t.Log("checked:", checked, "untested:", len(untested), untested)
		})
	}
}

// wasm32 lays out C types as emscripten does. The dumps come from a 64-bit
// Linux host, so their recorded sizes are ignored and only type names are used.
type wasm32 struct {
	records  map[string]ffiEntry
	typedefs map[string]*ffiEntry
}

var wasm32Sizes = map[string]int64{
	":_Bool": 1, ":char": 1, ":signed-char": 1, ":unsigned-char": 1,
	":short": 2, ":unsigned-short": 2,
	":int": 4, ":unsigned-int": 4, ":long": 4, ":unsigned-long": 4, ":enum": 4,
	":float": 4, ":pointer": 4, ":function-pointer": 4,
	":long-long": 8, ":unsigned-long-long": 8, ":double": 8,
	// glibc spells these long on 64-bit hosts, emscripten long long.
	"int64_t": 8, "uint64_t": 8,
}

func newWasm32(entries []ffiEntry) wasm32 {
	w := wasm32{records: map[string]ffiEntry{}, typedefs: map[string]*ffiEntry{}}
	for _, e := range entries {
		switch e.Tag {
		case "struct", "union":
			w.records[e.Name] = e
		case "typedef":
			w.typedefs[e.Name] = e.Type
		}
	}
	return w
}

// layout returns the size and alignment of typ, and the offset of each field
// when typ is a struct or union.
func (w wasm32) layout(typ ffiEntry) (size, align int64, offsets []int64) {
	if n, ok := wasm32Sizes[typ.Tag]; ok {
		return n, n, nil
	}
	switch typ.Tag {
	case ":array":
		size, align, _ = w.layout(*typ.Type)
		return size * int64(typ.Size), align, nil
	case ":struct", ":union":
		return w.layout(w.records[typ.Name])
	case "struct", "union":
		for _, f := range typ.Fields {
			fsize, falign, _ := w.layout(*f.Type)
			var off int64
			if typ.Tag == "struct" {
				off = (size + falign - 1) / falign * falign
			}
			offsets = append(offsets, off)
			size = max(size, off+fsize)
			align = max(align, falign)
		}
		return (size + align - 1) / align * align, align, offsets
	}
	td, ok := w.typedefs[typ.Tag]
	if !ok {
		panic("no wasm32 layout for C type " + typ.Tag)
	}
	return w.layout(*td)
}

var reflectBasics = map[types.BasicKind]reflect.Type{
	types.Bool:          reflect.TypeFor[bool](),
	types.Int:           reflect.TypeFor[int](),
	types.Int8:          reflect.TypeFor[int8](),
	types.Int16:         reflect.TypeFor[int16](),
	types.Int32:         reflect.TypeFor[int32](),
	types.Int64:         reflect.TypeFor[int64](),
	types.Uint:          reflect.TypeFor[uint](),
	types.Uint8:         reflect.TypeFor[uint8](),
	types.Uint16:        reflect.TypeFor[uint16](),
	types.Uint32:        reflect.TypeFor[uint32](),
	types.Uint64:        reflect.TypeFor[uint64](),
	types.Uintptr:       reflect.TypeFor[uintptr](),
	types.Float32:       reflect.TypeFor[float32](),
	types.Float64:       reflect.TypeFor[float64](),
	types.String:        reflect.TypeFor[string](),
	types.UnsafePointer: reflect.TypeFor[unsafe.Pointer](),
}

// reflectType rebuilds typ with reflect, which internal.CLayout reads, keeping
// only what it looks at: kinds, lengths and nesting. It reports false for types
// the js build cannot pass to C.
func reflectType(typ types.Type) (reflect.Type, bool) {
	switch u := typ.Underlying().(type) {
	case *types.Basic:
		r, ok := reflectBasics[u.Kind()]
		return r, ok
	case *types.Pointer:
		return reflect.TypeFor[*byte](), true
	case *types.Array:
		e, ok := reflectType(u.Elem())
		if !ok {
			return nil, false
		}
		return reflect.ArrayOf(int(u.Len()), e), true
	case *types.Struct:
		fields := make([]reflect.StructField, u.NumFields())
		for i := range fields {
			ft, ok := reflectType(u.Field(i).Type())
			if !ok {
				return nil, false
			}
			// StructOf rejects unexported names, and names play no part in layout.
			fields[i] = reflect.StructField{Name: fmt.Sprint("F", i), Type: ft}
		}
		return reflect.StructOf(fields), true
	}
	return nil, false
}
