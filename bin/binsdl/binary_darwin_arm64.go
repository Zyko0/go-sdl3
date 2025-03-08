//go:build darwin && arm64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_arm64.dylib
	sdlBlob    []byte
	sdlLibName = "libSDL3.dylib"
)
