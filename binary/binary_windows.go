//go:build windows

package binary

import (
	_ "embed"
)

var (
	//go:embed assets/sdl.dll
	sdlBlob []byte
)
