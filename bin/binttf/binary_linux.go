//go:build linux

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf.so
	ttfBlob    []byte
	ttfLibName = "libSDL3_ttf.so.0"
)
