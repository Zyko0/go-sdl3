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
func BytePtrPtrToStrSlice(byteptrptr **byte, count uint32, noempty bool) []string {
	pointer := *(*uintptr)(unsafe.Pointer(byteptrptr))
	strslice := make([]string, int(count))

	for i := range count {
		for noempty && len(ClonePtrString(pointer)) == 0 {
			// Some string arrays are returning with multiple null values, so single null termination does not get all values
			// Since we have the count and have stated noempty == true, we can be sure of a number of strings and skip nulls when we expect strings.
			pointer += uintptr(1) // Null terminated str
		}
		strslice[i] = ClonePtrString(pointer)
		pointer += uintptr(len(strslice[i]) + 1) // Null terminated str
	}

	return strslice
}
