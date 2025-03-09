//go:build !js

package img

import (
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/internal"
)

func (a *Animation) Frames() []*sdl.Surface {
	return internal.PtrToSlice[*sdl.Surface](uintptr(a.frames), int(a.count))
}

func (a *Animation) Delays() []int32 {
	return internal.PtrToSlice[int32](uintptr(a.delays), int(a.count))
}
