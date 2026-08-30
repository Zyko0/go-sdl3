package testing

import (
	"go/types"
	"runtime"
	"sort"
	"strings"
	"testing"
)

// structTypes returns every package-level struct type, exported or not, keyed
// by name.
func structTypes(t *testing.T, importPath string) map[string]*types.Struct {
	t.Helper()

	pkg := loadPackage(t, importPath)
	out := map[string]*types.Struct{}
	for _, name := range pkg.Scope().Names() {
		obj, ok := pkg.Scope().Lookup(name).(*types.TypeName)
		if !ok {
			continue
		}
		if s, ok := obj.Type().Underlying().(*types.Struct); ok {
			out[name] = s
		}
	}
	return out
}

// cLayout reports whether a type can share memory with C. Strings, slices and
// maps cannot, which is how the public Go-native structs are told apart from
// the private mirrors that are actually cast to and from C. Pointers stop the
// recursion: the pointee never contributes to the layout.
func cLayout(typ types.Type) bool {
	switch u := typ.Underlying().(type) {
	case *types.Basic:
		return u.Kind() != types.String
	case *types.Pointer:
		return true
	case *types.Array:
		return cLayout(u.Elem())
	case *types.Struct:
		for f := range u.Fields() {
			if !cLayout(f.Type()) {
				return false
			}
		}
		return true
	default:
		return false
	}
}

// unexport spells a name the way the private C-layout twins are spelled: a
// leading acronym is lowered whole, so GPUGraphicsPipelineCreateInfo pairs with
// gpuGraphicsPipelineCreateInfo, not gPUGraphics...
func unexport(s string) string {
	var upper int
	for upper < len(s) && s[upper] >= 'A' && s[upper] <= 'Z' {
		upper++
	}
	switch {
	case upper == 0:
		return s
	case upper < len(s):
		// The last capital of the run starts the next word.
		upper = max(upper-1, 1)
	}
	return strings.ToLower(s[:upper]) + s[upper:]
}

// Test_StructLayout compares every generated struct against the field offsets
// c2ffi recorded for the C type it mirrors. A mismatch is silent memory
// corruption at runtime, so it is worth catching without loading a library,
// which also means it runs on any GOOS/GOARCH.
func Test_StructLayout(t *testing.T) {
	sizes := types.SizesFor("gc", runtime.GOARCH)
	if sizes == nil {
		t.Skip("no gc sizes for", runtime.GOARCH)
	}

	for _, library := range libraries {
		t.Run(library, func(t *testing.T) {
			cfg := readJSON[ffiConfig](t, assetPath("config", library))
			entries := readJSON[[]ffiEntry](t, assetPath("ffi", library))
			goStructs := structTypes(t, modulePath+"/"+cfg.OutDir)

			var checked int
			var untested []string
			for _, e := range entries {
				if e.Tag != "struct" || len(e.Fields) == 0 {
					continue
				}
				if !cfg.includes(e.Location) {
					continue
				}

				// The generators strip the library prefix. A C-layout twin of a
				// Go-native struct keeps the same name with a lowercase initial,
				// so both spellings are candidates.
				base := strings.TrimPrefix(e.Name, cfg.Prefix)
				var compared bool
				for _, name := range []string{base, unexport(base)} {
					s, ok := goStructs[name]
					// Opaque handles are emitted as empty structs, and a few
					// types are hand-written with a subset of the C fields.
					// Neither can be compared positionally.
					if !ok || s.NumFields() != len(e.Fields) || !cLayout(s) {
						continue
					}
					compared = true
					checkStruct(t, sizes, name, s, e)
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

func checkStruct(t *testing.T, sizes types.Sizes, name string, s *types.Struct, e ffiEntry) {
	t.Helper()

	if got, want := sizes.Sizeof(s), int64(e.BitSize/8); got != want {
		t.Errorf("%s (Go %s): size %d, C says %d", e.Name, name, got, want)
		return
	}

	fields := make([]*types.Var, s.NumFields())
	for i := range fields {
		fields[i] = s.Field(i)
	}
	offsets := sizes.Offsetsof(fields)
	for i, f := range e.Fields {
		if got, want := offsets[i], int64(f.BitOffset/8); got != want {
			t.Errorf("%s.%s (Go %s.%s): offset %d, C says %d",
				e.Name, f.Name, name, fields[i].Name(), got, want)
		}
	}
}
