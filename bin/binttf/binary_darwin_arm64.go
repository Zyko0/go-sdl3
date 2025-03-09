//go:build darwin && arm64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf_arm64.dylib
	ttfBlob    []byte
	ttfLibName = "libSDL3_ttf.dylib"
)
