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

func NewDialogFileCallback(fn func(fileList []string, filter int32)) DialogFileCallback {
	return DialogFileCallback(purego.NewCallback(func(_ uintptr, fileList uintptr, filter int32) uintptr {
		files := make([]string, 0)
		for ptr := fileList; ptr != 0; ptr += unsafe.Sizeof(uintptr(0)) {
			strPtr := *(*uintptr)(unsafe.Pointer(ptr))
			if strPtr == 0 {
				break
			}
			str := internal.PtrToString(strPtr)
			files = append(files, str)
		}
		fn(files, filter)
		return 0
	}))
}

func NewEnumerateDirectoryCallback(fn func(dirname, fname string)) EnumerateDirectoryCallback {
	return EnumerateDirectoryCallback(purego.NewCallback(func(_, dirname, fname uintptr) uintptr {
		fn(internal.PtrToString(dirname), internal.PtrToString(fname))
		return 0
	}))
}

func NewEventFilter(fn func(event *Event) bool) EventFilter {
	return EventFilter(purego.NewCallback(func(_ uintptr, event *Event) uintptr {
		if fn(event) {
			return 1
		}
		return 0
	}))
}
