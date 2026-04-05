//go:build linux && arm64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_arm64.so.gz
	sdlBlob    []byte
	sdlLibName = "libSDL3.so.0"
)
