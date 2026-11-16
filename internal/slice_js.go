//go:build js

package internal

import (
	"reflect"
	"unsafe"
)

// heapSlice decodes count C-layout T values starting at addr.
func heapSlice[T any](addr, count int) []T {
	t := reflect.TypeFor[T]()
	l := layoutOf(t)
	buf := make([]byte, int(l.size)*count)
	readHeap(addr, buf)
	s := make([]T, count)
	for i := range s {
		decodeValue(unsafe.Pointer(&s[i]), buf[uintptr(i)*l.size:], t)
	}

	return s
}

// ClonePtrSlice returns a newly allocated slice from a uintptr
func ClonePtrSlice[T any](ptr uintptr, count int) []T {
	if count <= 0 {
		return nil
	}
	if addr, ok := lookupHeap(ptr); ok {
		return heapSlice[T](addr, count)
	}

	return append([]T(nil), unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)...)
}

// ClonePtrSliceValues copies count pointers from ptr and the values they
// point to, for C arrays whose elements share the array's allocation.
func ClonePtrSliceValues[T any](ptr uintptr, count int) []*T {
	if count <= 0 {
		return nil
	}

	values := make([]T, count)
	ptrs := make([]*T, count)
	for i, p := range PtrToSlice[*T](ptr, count) {
		values[i] = *p
		ptrs[i] = &values[i]
	}

	return ptrs
}

// PtrToSlice returns a slice pointing to the provided uintptr data. Heap
// memory cannot be aliased from Go, so for it the slice is a copy.
func PtrToSlice[T any](ptr uintptr, count int) []T {
	if addr, ok := lookupHeap(ptr); ok {
		return heapSlice[T](addr, count)
	}

	return unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)
}

// BytePtrPtrToStrSlice returns the count strings of the C string array
// byteptrptr.
func BytePtrPtrToStrSlice(byteptrptr **byte, count uint32, noempty bool) []string {
	addr, ok := lookupHeap(uintptr(unsafe.Pointer(byteptrptr)))
	if !ok {
		panic("internal: string array was not returned by SDL")
	}

	return heapSlice[string](addr, int(count))
}

// ClonePtrStringSlice returns the count strings of the C string array ptr.
func ClonePtrStringSlice(ptr uintptr, count int) []string {
	return ClonePtrSlice[string](ptr, count)
}
