package sdl

import (
	"errors"
	"unsafe"

	puregogen "github.com/Zyko0/purego-gen"
)

// lastError returns the last sdl error given by SDL_GetError
func lastError() error {
	if msg := iGetError(); msg != "" {
		return errors.New(msg)
	}
	return nil
}

// sdlFree frees an sdl-allocated pointer via SDL_free
func sdlFree(ptr uintptr) {
	ifree(ptr)
}

// clonePtrSlice returns a newly allocated slice from a uintptr
func clonePtrSlice[T any](ptr uintptr, count int) []T {
	var s []T

	if count > 0 {
		s = append(s, unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)...)
	}

	return s
}

// ptrToSlice returns a slice pointing to the provided uintptr data
func ptrToSlice[T any](ptr uintptr, count int) []T {
	return unsafe.Slice(*(**T)(unsafe.Pointer(&ptr)), count)
}

// clonePtrString returns a newly allocated string from a uintptr
func clonePtrString(ptr uintptr) string {
	return "" + puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}

// ptrToString returns a string pointing to the provided ptr char data
func ptrToString[T any](ptr uintptr) string {
	return puregogen.BytePtrToString(*(**byte)(unsafe.Pointer(&ptr)))
}
