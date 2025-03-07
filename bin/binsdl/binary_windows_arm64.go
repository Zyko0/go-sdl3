//go:build windows && arm64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_arm64.dll
	sdlBlob    []byte
	sdlLibName = "sdl.dll"
)
