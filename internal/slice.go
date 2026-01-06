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

// PtrToSlice returns a slice pointing to the provided uintptr data
func PtrToSlice[T any](ptr uintptr, count int) []T {
	return unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)
}

// BytePtrPtrToStrSlice returns a slice of strings from the pointer and count of strings
func BytePtrPtrToStrSlice(byteptrptr **byte, count uint32) []string {
	pointer := *(*uintptr)(unsafe.Pointer(byteptrptr))
	strslice := make([]string, int(count))

	for i := range count {
		strslice[i] = ClonePtrString(pointer)
		pointer += uintptr(len(strslice[i])) + 1 // Null terminated str
	}

	return strslice
}
