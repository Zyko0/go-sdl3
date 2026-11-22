package internal

import "reflect"

// Go on wasm is a 64-bit target while emscripten is wasm32: pointers,
// uintptr, int and strings take 8 or 16 bytes in Go but 4 bytes in C, so
// structs holding them are re-encoded field by field instead of copied.

type cLayout struct {
	size  uintptr
	align uintptr
	// same is true when the Go memory layout is byte-identical to C.
	same   bool
	fields []cField
}

type cField struct {
	goOff uintptr
	cOff  uintptr
	typ   reflect.Type
}

var layouts = map[reflect.Type]*cLayout{}

func layoutOf(t reflect.Type) *cLayout {
	if l, ok := layouts[t]; ok {
		return l
	}
	l := &cLayout{}
	switch t.Kind() {
	case reflect.Bool, reflect.Int8, reflect.Uint8:
		l.size, l.same = 1, true
	case reflect.Int16, reflect.Uint16:
		l.size, l.same = 2, true
	case reflect.Int32, reflect.Uint32, reflect.Float32:
		l.size, l.same = 4, true
	case reflect.Int64, reflect.Uint64, reflect.Float64:
		l.size, l.same = 8, true
	case reflect.Int, reflect.Uint, reflect.Uintptr, reflect.UnsafePointer,
		reflect.Pointer, reflect.String:
		l.size = 4
	case reflect.Array:
		e := layoutOf(t.Elem())
		l.size = e.size * uintptr(t.Len())
		l.align = e.align
		l.same = e.same && e.size == t.Elem().Size()
	case reflect.Struct:
		l.same = true
		for i := range t.NumField() {
			f := t.Field(i)
			fl := layoutOf(f.Type)
			off := alignUp(l.size, fl.align)
			l.fields = append(l.fields, cField{goOff: f.Offset, cOff: off, typ: f.Type})
			l.same = l.same && fl.same && off == f.Offset
			l.size = off + fl.size
			l.align = max(l.align, fl.align)
		}
		l.size = alignUp(l.size, max(l.align, 1))
		l.same = l.same && l.size == t.Size()
	default:
		panic("internal: " + t.String() + " has no C representation")
	}
	if l.align == 0 {
		l.align = max(l.size, 1)
	}
	layouts[t] = l
	return l
}

func alignUp(n, align uintptr) uintptr {
	return (n + align - 1) &^ (align - 1)
}

// CLayout returns the size and field offsets of the struct t in the emscripten
// (wasm32) C ABI. Unlike the rest of the layout code it needs no js runtime, so
// it is also built natively, where it can be checked against the C headers.
func CLayout(t reflect.Type) (size uintptr, offsets []uintptr) {
	l := layoutOf(t)
	for _, f := range l.fields {
		offsets = append(offsets, f.cOff)
	}
	return l.size, offsets
}
