//go:build windows

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl.dll
	sdlBlob    []byte
	sdlLibName = "sdl.dll"
)
