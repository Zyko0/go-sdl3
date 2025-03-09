//go:build js

package img

import (
	"syscall/js"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/internal"
)

func (a *Animation) Frames() []*sdl.Surface {
	return internal.GetObjectSliceFromJSPtr[sdl.Surface](js.ValueOf(a.frames), int(a.count))
}

func (a *Animation) Delays() []int32 {
	return internal.GetNumericSliceFromJSPtr[int32](js.ValueOf(a.delays), int(a.count))
}
