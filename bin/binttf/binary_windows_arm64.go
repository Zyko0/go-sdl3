//go:build windows && arm64

package binttf

import (
	_ "embed"
)

var (
	//go:embed assets/ttf_arm64.dll.gz
	ttfBlob    []byte
	ttfLibName = "SDL3_ttf.dll"
)
