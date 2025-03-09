//go:build darwin && amd64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf.dylib
	ttfBlob    []byte
	ttfLibName = "libSDL3_ttf.dylib"
)
