//go:build darwin

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl.dylib
	sdlBlob    []byte
	sdlLibName = "sdl.dylib"
)
