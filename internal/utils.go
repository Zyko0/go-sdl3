package internal

import (
	"unsafe"

	puregogen "github.com/Zyko0/purego-gen"
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

// ClonePtrString returns a newly allocated string from a uintptr
func ClonePtrString(ptr uintptr) string {
	return "" + puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}

// PtrToString returns a string pointing to the provided ptr char data
func PtrToString[T any](ptr uintptr) string {
	return puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}
