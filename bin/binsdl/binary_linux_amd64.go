//go:build linux && amd64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_amd64.so.gz
	sdlBlob    []byte
	sdlLibName = "libSDL3.so.0"
)
