//go:build windows && arm64

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img_arm64.dll.gz
	imgBlob    []byte
	imgLibName = "SDL3_image.dll"
)
