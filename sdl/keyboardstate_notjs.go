//go:build !js

package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// SDL_GetKeyboardState - Get a snapshot of the current state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState)
func GetKeyboardState() []bool {
	var count int32

	ptr := iGetKeyboardState(&count)

	return internal.PtrToSlice[bool](uintptr(unsafe.Pointer(ptr)), int(count))
}
