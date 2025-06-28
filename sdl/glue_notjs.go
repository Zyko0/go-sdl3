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

func NewCleanupPropertyCallback(fn func(value uintptr)) CleanupPropertyCallback {
	return CleanupPropertyCallback(purego.NewCallback(func(userData, value uintptr) uintptr {
		fn(value)
		return 0
	}))
}

func NewEnumeratePropertiesCallback(fn func(props PropertiesID, name string)) EnumeratePropertiesCallback {
	return EnumeratePropertiesCallback(purego.NewCallback(func(userData uintptr, props PropertiesID, name uintptr) uintptr {
		fn(props, internal.PtrToString(name))
		return 0
	}))
}

func NewTLSDestructorCallback(fn func(value uintptr)) TLSDestructorCallback {
	return TLSDestructorCallback(purego.NewCallback(func(value uintptr) uintptr {
		fn(value)
		return 0
	}))
}

func NewAudioStreamCallback(fn func(stream *AudioStream, additionalAmount, totalAmount int32)) AudioStreamCallback {
	return AudioStreamCallback(purego.NewCallback(func(userData uintptr, stream *AudioStream, additionalAmount, totalAmount int32) uintptr {
		fn(stream, additionalAmount, totalAmount)
		return 0
	}))
}

func NewAudioPostmixCallback(fn func(spec *AudioSpec, buffer []float32)) AudioPostmixCallback {
	return AudioPostmixCallback(purego.NewCallback(func(userData uintptr, spec *AudioSpec, buffer *float32, bufLen int32) uintptr {
		fn(spec, unsafe.Slice(buffer, bufLen/4))
		runtime.KeepAlive(buffer)
		return 0
	}))
}
