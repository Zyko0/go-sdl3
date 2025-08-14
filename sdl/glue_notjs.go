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
	return CleanupPropertyCallback(purego.NewCallback(func(_, value uintptr) uintptr {
		fn(value)
		return 0
	}))
}

func NewEnumeratePropertiesCallback(fn func(props PropertiesID, name string)) EnumeratePropertiesCallback {
	return EnumeratePropertiesCallback(purego.NewCallback(func(_ uintptr, props PropertiesID, name uintptr) uintptr {
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
	return AudioStreamCallback(purego.NewCallback(func(_ uintptr, stream *AudioStream, additionalAmount, totalAmount int32) uintptr {
		fn(stream, additionalAmount, totalAmount)
		return 0
	}))
}

func NewAudioPostmixCallback(fn func(spec *AudioSpec, buffer []float32)) AudioPostmixCallback {
	return AudioPostmixCallback(purego.NewCallback(func(_ uintptr, spec *AudioSpec, buffer *float32, bufLen int32) uintptr {
		fn(spec, unsafe.Slice(buffer, bufLen/4))
		runtime.KeepAlive(buffer)
		return 0
	}))
}

func NewClipboardDataCallback(fn func(mimeType string) []byte) ClipboardDataCallback {
	return ClipboardDataCallback(purego.NewCallback(func(_ uintptr, mimeType uintptr, size *uintptr) uintptr {
		data := fn(internal.PtrToString(mimeType))
		if size != nil {
			*size = uintptr(len(data))
		}
		return uintptr(unsafe.Pointer(unsafe.SliceData(data)))
	}))
}

func NewClipboardCleanupCallback(fn func()) ClipboardCleanupCallback {
	return ClipboardCleanupCallback(purego.NewCallback(func(_ uintptr) uintptr {
		fn()
		return 0
	}))
}
