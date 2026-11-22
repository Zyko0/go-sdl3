//go:build js

package internal

import (
	"reflect"
	"syscall/js"
	"unsafe"
)

// allocator returns the address of n bytes in the emscripten heap or stack.
type allocator func(n int) int

func stackAllocator(n int) int {
	return stackAlloc.Invoke(int32(n)).Int()
}

type heapAllocs []int

func (h *heapAllocs) alloc(n int) int {
	addr := js.Global().Get("Module").Call("_malloc", max(n, 1)).Int()
	*h = append(*h, addr)
	return addr
}

func (h *heapAllocs) free() {
	for _, addr := range *h {
		js.Global().Get("Module").Call("_free", addr)
	}
	*h = nil
}

func readHeap(addr int, buf []byte) {
	js.CopyBytesToGo(buf, heapU8().Call("subarray", addr, addr+len(buf)))
}

func writeHeap(addr int, buf []byte) {
	js.CopyBytesToJS(heapU8().Call("subarray", addr, addr+len(buf)), buf)
}

func putU32(buf []byte, v uint32) {
	buf[0], buf[1], buf[2], buf[3] = byte(v), byte(v>>8), byte(v>>16), byte(v>>24)
}

func getU32(buf []byte) uint32 {
	return uint32(buf[0]) | uint32(buf[1])<<8 | uint32(buf[2])<<16 | uint32(buf[3])<<24
}

// encodeValue writes the Go value at src into buf using the C layout of t.
// Strings and Go-owned pointees are copied with alloc.
func encodeValue(buf []byte, src unsafe.Pointer, t reflect.Type, alloc allocator) {
	l := layoutOf(t)
	if l.same {
		copy(buf[:l.size], unsafe.Slice((*byte)(src), l.size))
		return
	}
	switch t.Kind() {
	case reflect.Int:
		putU32(buf, uint32(*(*int)(src)))
	case reflect.Uint:
		putU32(buf, uint32(*(*uint)(src)))
	case reflect.Uintptr:
		putU32(buf, encodeAddress(*(*uintptr)(src)))
	case reflect.UnsafePointer:
		putU32(buf, encodeAddress(uintptr(*(*unsafe.Pointer)(src))))
	case reflect.Pointer:
		putU32(buf, encodePointer(*(*unsafe.Pointer)(src), t.Elem(), alloc))
	case reflect.String:
		putU32(buf, encodeString(*(*string)(src), alloc))
	case reflect.Array:
		e, goSize := layoutOf(t.Elem()), t.Elem().Size()
		for i := range uintptr(t.Len()) {
			encodeValue(buf[i*e.size:], unsafe.Add(src, i*goSize), t.Elem(), alloc)
		}
	case reflect.Struct:
		for _, f := range l.fields {
			encodeValue(buf[f.cOff:], unsafe.Add(src, f.goOff), f.typ, alloc)
		}
	}
}

// encodeAddress passes registered Go stand-ins as their heap address and
// any other value (sizes, opaque userdata) unchanged.
func encodeAddress(v uintptr) uint32 {
	if addr, ok := lookupHeap(v); ok {
		return uint32(addr)
	}
	return uint32(v)
}

func encodePointer(p unsafe.Pointer, elem reflect.Type, alloc allocator) uint32 {
	if p == nil {
		return 0
	}
	if addr, ok := lookupHeap(uintptr(p)); ok {
		return uint32(addr)
	}
	// Go-owned pointee: C gets a copy of one element. Arrays must be passed
	// through SliceData so that their length is known.
	l := layoutOf(elem)
	if l.size == 0 {
		panic("internal: " + elem.String() + " was not returned by SDL")
	}
	buf := make([]byte, l.size)
	encodeValue(buf, p, elem, alloc)
	addr := alloc(len(buf))
	writeHeap(addr, buf)
	return uint32(addr)
}

func encodeString(s string, alloc allocator) uint32 {
	if s == "" {
		return 0
	}
	addr := alloc(len(s) + 1)
	writeHeap(addr, append([]byte(s), 0))
	return uint32(addr)
}

// decodeValue reads buf, laid out as the C version of t, into the Go value
// at dst. C pointers become registered Go stand-ins.
func decodeValue(dst unsafe.Pointer, buf []byte, t reflect.Type) {
	l := layoutOf(t)
	if l.same {
		copy(unsafe.Slice((*byte)(dst), l.size), buf[:l.size])
		return
	}
	switch t.Kind() {
	case reflect.Int:
		*(*int)(dst) = int(int32(getU32(buf)))
	case reflect.Uint:
		*(*uint)(dst) = uint(getU32(buf))
	case reflect.Uintptr:
		*(*uintptr)(dst) = uintptr(getU32(buf))
	case reflect.UnsafePointer:
		*(*unsafe.Pointer)(dst) = newObject(reflect.TypeFor[byte](), int(getU32(buf)))
	case reflect.Pointer:
		*(*unsafe.Pointer)(dst) = newObject(t.Elem(), int(getU32(buf)))
	case reflect.String:
		*(*string)(dst) = heapString(int(getU32(buf)))
	case reflect.Array:
		e, goSize := layoutOf(t.Elem()), t.Elem().Size()
		for i := range uintptr(t.Len()) {
			decodeValue(unsafe.Add(dst, i*goSize), buf[i*e.size:], t.Elem())
		}
	case reflect.Struct:
		for _, f := range l.fields {
			decodeValue(unsafe.Add(dst, f.goOff), buf[f.cOff:], f.typ)
		}
	}
}

func decodeFromHeap(dst unsafe.Pointer, t reflect.Type, addr int) {
	l := layoutOf(t)
	if l.size == 0 {
		return
	}
	buf := make([]byte, l.size)
	readHeap(addr, buf)
	decodeValue(dst, buf, t)
}

func heapString(addr int) string {
	if addr == 0 {
		return ""
	}
	return utf8String.Invoke(addr).String()
}

// CSizeof returns the size of T in the emscripten (wasm32) C ABI.
func CSizeof[T any]() int {
	return int(layoutOf(reflect.TypeFor[T]()).size)
}

// CStructAt returns the struct stored at p in C layout, such as an event
// variant inside the raw bytes of an SDL_Event.
func CStructAt[T any](p unsafe.Pointer) *T {
	if p == nil {
		return nil
	}
	t := reflect.TypeFor[T]()
	v := new(T)
	decodeValue(unsafe.Pointer(v), unsafe.Slice((*byte)(p), layoutOf(t).size), t)
	return v
}
