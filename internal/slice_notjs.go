//go:build !js

package internal

import (
	"unsafe"
)

// ClonePtrSlice returns a newly allocated slice from a uintptr
func ClonePtrSlice[T any](ptr uintptr, count int) []T {
	var s []T

	if count > 0 {
		s = append(s, unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)...)
	}

	return s
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

// PtrToSlice returns a slice pointing to the provided uintptr data
func PtrToSlice[T any](ptr uintptr, count int) []T {
	return unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)
}

// ClonePtrStringSlice returns the count strings of the C string array ptr.
func ClonePtrStringSlice(ptr uintptr, count int) []string {
	if count <= 0 {
		return nil
	}

	strs := make([]string, count)
	for i, p := range PtrToSlice[uintptr](ptr, count) {
		strs[i] = ClonePtrString(p)
	}

	return strs
}
