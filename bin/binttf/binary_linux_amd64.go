//go:build linux && amd64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf_amd64.so.gz
	ttfBlob    []byte
	ttfLibName = "libSDL3_ttf.so.0"
)
