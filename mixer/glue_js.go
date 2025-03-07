//go:build js

package mixer

import (
	js "syscall/js"

	internal "github.com/Zyko0/go-sdl3/internal"
)

func (c *Chunk) Buffer() []byte {
	return internal.GetByteSliceFromJSPtr(js.ValueOf(c.abuf), int(c.alen))
}
