//go:build !js

package sdl

import "github.com/Zyko0/go-sdl3/internal"

func (s *Surface) Pixels() []byte {
	return internal.PtrToSlice[byte](uintptr(s.pixels), int(s.H*s.Pitch))
}
