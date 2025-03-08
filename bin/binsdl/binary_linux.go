//go:build linux

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl.so
	sdlBlob    []byte
	sdlLibName = "libSDL3.so.0"
)
