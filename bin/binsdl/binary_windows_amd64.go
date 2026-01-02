//go:build windows && amd64

package binsdl

import (
	_ "embed"
)

var (
	//go:embed assets/sdl_amd64.dll.gz
	sdlBlob    []byte
	sdlLibName = "SDL3.dll"
)
