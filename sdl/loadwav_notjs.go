//go:build !js

package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// SDL_LoadWAV_IO - Load the audio data of a WAVE file into memory.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO)
func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV_IO(src, closeIO, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// SDL_LoadWAV - Loads a WAV from a file path.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV)
func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV(path, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}
