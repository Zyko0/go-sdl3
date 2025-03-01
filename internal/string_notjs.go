//go:build !js

package internal

import (
	"unsafe"

	puregogen "github.com/Zyko0/purego-gen"
)

// ClonePtrString returns a newly allocated string from a uintptr
func ClonePtrString(ptr uintptr) string {
	return "" + puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}

// PtrToString returns a string pointing to the provided ptr char data
func PtrToString(ptr uintptr) string {
	return puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}
