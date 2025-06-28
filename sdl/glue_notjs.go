//go:build !js

package sdl

import (
	"runtime"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	purego "github.com/ebitengine/purego"
)

func (s *Surface) Pixels() []byte {
	return internal.PtrToSlice[byte](uintptr(s.pixels), int(s.H*s.Pitch))
}

// Callbacks

func NewCleanupPropertyCallback(fn func(userData, value uintptr)) CleanupPropertyCallback {
	return CleanupPropertyCallback(purego.NewCallback(fn))
}

func NewEnumeratePropertiesCallback(fn func(userData uintptr, name string)) EnumeratePropertiesCallback {
	return EnumeratePropertiesCallback(purego.NewCallback(fn))
}

func NewTLSDestructorCallback(fn func(value uintptr)) TLSDestructorCallback {
	return TLSDestructorCallback(purego.NewCallback(fn))
}

func NewAudioStreamCallback(fn func(userData uintptr, stream *AudioStream, additionalAmount, totalAmount int32)) AudioStreamCallback {
	return AudioStreamCallback(purego.NewCallback(fn))
}

func NewAudioPostmixCallback(fn func(userData uintptr, spec *AudioSpec, buffer []float32)) AudioPostmixCallback {
	return AudioPostmixCallback(purego.NewCallback(func(userData uintptr, spec *AudioSpec, buffer *float32, bufLen int32) {
		fn(userData, spec, unsafe.Slice(buffer, bufLen))
		runtime.KeepAlive(buffer)
	}))
}
