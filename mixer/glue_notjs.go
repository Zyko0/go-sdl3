//go:build !js

package mixer

import internal "github.com/Zyko0/go-sdl3/internal"

func (c *Chunk) Buffer() []byte {
	return internal.PtrToSlice[byte](uintptr(c.abuf), int(c.alen))
}
