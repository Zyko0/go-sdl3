//go:build windows && amd64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf_amd64.dll
	ttfBlob    []byte
	ttfLibName = "SDL3_ttf.dll"
)
