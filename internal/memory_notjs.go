//go:build !js

package internal

import "unsafe"

type Pointer = unsafe.Pointer

// SliceData returns a pointer to the first element of s that C can
// dereference.
func SliceData[T any](s []T) *T {
	return unsafe.SliceData(s)
}

// CStructAt returns the struct stored at p in C layout, such as an event
// variant inside the raw bytes of an SDL_Event.
func CStructAt[T any](p unsafe.Pointer) *T {
	return (*T)(p)
}
