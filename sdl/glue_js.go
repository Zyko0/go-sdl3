//go:build js

package sdl

import (
	"syscall/js"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

func (s *Surface) Pixels() []byte {
	return internal.GetByteSliceFromJSPtr(js.ValueOf(int32(s.pixels)), int(s.H*s.Pitch))
}

// planeSize returns the bytes spanned by rows rows of rowBytes bytes, pitch
// bytes apart: the last row may end before the pitch.
func planeSize(rows, pitch, rowBytes int32) int {
	if rows <= 0 {
		return 0
	}
	return int((rows-1)*pitch + rowBytes)
}

// lockedTexture is the Go mirror of the heap pixels of a locked texture.
type lockedTexture struct {
	addr   int
	pixels []byte
}

// C cannot alias Go memory, so a locked texture exposes a Go copy of its
// pixels, written back to the heap by UnlockTexture.
var lockedTextures = map[*Texture]lockedTexture{}

func lockTextureMirror(texture *Texture, rect *Rect, addr int, pitch int32) uintptr {
	size := int(texture.H * pitch)
	if rect != nil {
		bpp := int32(texture.Format & 0xFF)
		// Packed FOURCC formats use 2 bytes per pixel; SDL refuses partial
		// locks of planar ones.
		if texture.Format != 0 && (texture.Format>>28)&0x0F != 1 {
			bpp = 2
		}
		size = planeSize(rect.H, pitch, rect.W*bpp)
	}
	// The Lock wrapper slices texture.H*pitch bytes whatever the rect.
	pixels := make([]byte, max(int(texture.H*pitch), size, 1))
	copy(pixels, internal.GetByteSliceFromJSPtr(js.ValueOf(addr), size))
	lockedTextures[texture] = lockedTexture{addr: addr, pixels: pixels[:size]}
	return uintptr(unsafe.Pointer(&pixels[0]))
}

func unlockTextureMirror(texture *Texture) {
	m, ok := lockedTextures[texture]
	if !ok {
		return
	}
	delete(lockedTextures, texture)
	internal.SetNumericSliceToJSPtr(js.ValueOf(m.addr), m.pixels)
}

// Callbacks

// Opaque pointers such as userdata and property values are never
// dereferenced by C, so they round-trip as raw 32-bit values.
func rawPointer(v js.Value) uintptr {
	return uintptr(uint32(v.Int()))
}

func NewCleanupPropertyCallback(fn func(value uintptr)) CleanupPropertyCallback {
	return CleanupPropertyCallback(internal.NewCallback("vpp", func(args []js.Value) any {
		fn(rawPointer(args[1]))
		return nil
	}))
}

func NewEnumeratePropertiesCallback(fn func(props PropertiesID, name string)) EnumeratePropertiesCallback {
	return EnumeratePropertiesCallback(internal.NewCallback("vpip", func(args []js.Value) any {
		fn(PropertiesID(args[1].Int()), internal.UTF8JSToString(args[2]))
		return nil
	}))
}

func NewTLSDestructorCallback(fn func(value uintptr)) TLSDestructorCallback {
	return TLSDestructorCallback(internal.NewCallback("vp", func(args []js.Value) any {
		fn(rawPointer(args[0]))
		return nil
	}))
}

func NewAudioStreamCallback(fn func(stream *AudioStream, additionalAmount, totalAmount int32)) AudioStreamCallback {
	return AudioStreamCallback(internal.NewCallback("vppii", func(args []js.Value) any {
		fn(internal.NewObject[AudioStream](args[1]), int32(args[2].Int()), int32(args[3].Int()))
		return nil
	}))
}

// The buffer is a copy written back after fn returns: Go cannot alias heap
// memory.
func NewAudioPostmixCallback(fn func(spec *AudioSpec, buffer []float32)) AudioPostmixCallback {
	return AudioPostmixCallback(internal.NewCallback("vpppi", func(args []js.Value) any {
		var spec AudioSpec
		internal.CopyJSToObject(&spec, args[1])
		buffer := internal.GetNumericSliceFromJSPtr[float32](args[2], args[3].Int()/4)
		fn(&spec, buffer)
		internal.SetNumericSliceToJSPtr(args[2], buffer)
		return nil
	}))
}

// SDL reads the returned data until the next request, so each call frees
// the copy made by the previous one.
func NewClipboardDataCallback(fn func(mimeType string) []byte) ClipboardDataCallback {
	var last js.Value
	return ClipboardDataCallback(internal.NewCallback("pppp", func(args []js.Value) any {
		if !last.IsUndefined() {
			js.Global().Call("_free", last)
			last = js.Value{}
		}
		data := fn(internal.UTF8JSToString(args[1]))
		if !args[2].IsNull() && args[2].Int() != 0 {
			internal.SetValue(args[2], js.ValueOf(len(data)), "*")
		}
		if len(data) == 0 {
			return 0
		}
		last = internal.CloneByteSliceToJSHeap(data)
		return last
	}))
}

func NewClipboardCleanupCallback(fn func()) ClipboardCleanupCallback {
	return ClipboardCleanupCallback(internal.NewCallback("vp", func(args []js.Value) any {
		fn()
		return nil
	}))
}

func NewDialogFileCallback(fn func(fileList []string, filter int32)) DialogFileCallback {
	return DialogFileCallback(internal.NewCallback("vppi", func(args []js.Value) any {
		files := make([]string, 0)
		if list := args[1].Int(); list != 0 {
			for i := 0; ; i++ {
				str := internal.GetValue(js.ValueOf(list+i*4), "*")
				if str.Int() == 0 {
					break
				}
				files = append(files, internal.UTF8JSToString(str))
			}
		}
		fn(files, int32(args[2].Int()))
		return nil
	}))
}

func NewEnumerateDirectoryCallback(fn func(dirname, fname string)) EnumerateDirectoryCallback {
	return EnumerateDirectoryCallback(internal.NewCallback("ippp", func(args []js.Value) any {
		fn(internal.UTF8JSToString(args[1]), internal.UTF8JSToString(args[2]))
		return 0
	}))
}

// The event is a copy written back after fn returns: Go cannot alias heap
// memory.
func NewEventFilter(fn func(event *Event) bool) EventFilter {
	return EventFilter(internal.NewCallback("ipp", func(args []js.Value) any {
		var event Event
		internal.CopyJSToObject(&event, args[1])
		ok := fn(&event)
		internal.SetNumericSliceToJSPtr(args[1], unsafe.Slice((*byte)(unsafe.Pointer(&event)), unsafe.Sizeof(event)))
		if ok {
			return 1
		}
		return 0
	}))
}

func NewHitTest(fn func(window *Window, area *Point) HitTestResult) HitTest {
	// Unlike other callbacks, userdata comes last.
	return HitTest(internal.NewCallback("ippp", func(args []js.Value) any {
		var area Point
		internal.CopyJSToObject(&area, args[1])
		return int32(fn(internal.NewObject[Window](args[0]), &area))
	}))
}

func NewMouseMotionTransformCallback(fn func(timestamp uint64, window *Window, mouseID MouseID, x, y *float32)) MouseMotionTransformCallback {
	return MouseMotionTransformCallback(internal.NewCallback("vpjpipp", func(args []js.Value) any {
		x := float32(internal.GetValue(args[4], "float").Float())
		y := float32(internal.GetValue(args[5], "float").Float())
		fn(internal.GetUint64(args[1]), internal.NewObject[Window](args[2]), MouseID(args[3].Int()), &x, &y)
		internal.SetValue(args[4], js.ValueOf(x), "float")
		internal.SetValue(args[5], js.ValueOf(y), "float")
		return nil
	}))
}

func NewTrayCallback(fn func(entry *TrayEntry)) TrayCallback {
	return TrayCallback(internal.NewCallback("vpp", func(args []js.Value) any {
		fn(internal.NewObject[TrayEntry](args[1]))
		return nil
	}))
}
