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

// StringToNullablePtr returns a uintptr pointing to the provided string data
func StringToNullablePtr(s string) *byte {
	if len(s) == 0 {
		return nil
	}
	return puregogen.BytePtrFromString(s)
}

// StringToPtr returns a uintptr pointing to the provided string data
func StringToPtr(s string) *byte {
	return puregogen.BytePtrFromString(s)
}
