//go:build windows && amd64

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img_amd64.dll
	imgBlob    []byte
	imgLibName = "SDL3_image.dll"
)
