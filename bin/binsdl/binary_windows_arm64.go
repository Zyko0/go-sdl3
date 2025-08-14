//go:build windows && arm64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_arm64.dll.gz
	sdlBlob    []byte
	sdlLibName = "SDL3.dll"
)
