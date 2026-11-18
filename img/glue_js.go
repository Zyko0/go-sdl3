//go:build js

package img

import (
	"syscall/js"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// cStringArrayToJSHeap copies the NULL-terminated array built by
// cStringArray, passed as *string, to the heap. The returned function frees
// the copy.
func cStringArrayToJSHeap(xpm *string) (js.Value, func()) {
	ptrs := (**byte)(unsafe.Pointer(xpm))
	n := 0
	for *(**byte)(unsafe.Add(unsafe.Pointer(ptrs), uintptr(n)*unsafe.Sizeof(ptrs))) != nil {
		n++
	}
	return internal.ClonePtrArrayToJSHeap(ptrs, n+1)
}

func (a *Animation) Frames() []*sdl.Surface {
	return internal.GetObjectSliceFromJSPtr[sdl.Surface](js.ValueOf(int32(a.frames)), int(a.count))
}

func (a *Animation) Delays() []int32 {
	return internal.GetNumericSliceFromJSPtr[int32](js.ValueOf(int32(a.delays)), int(a.count))
}
