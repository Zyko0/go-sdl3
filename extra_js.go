//go:build js

package sdl

import (
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
)

func (s *Surface) Pixels() []byte {
	return internal.GetByteSliceFromJSPtr(js.ValueOf(s.pixels), int(s.H*s.Pitch))
}
