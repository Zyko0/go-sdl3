//go:build darwin && amd64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf_amd64.dylib
	ttfBlob    []byte
	ttfLibName = "libSDL3_ttf.dylib"
)
