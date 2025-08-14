//go:build darwin && amd64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_amd64.dylib.gz
	sdlBlob    []byte
	sdlLibName = "libSDL3.dylib"
)
