//go:build linux && arm64

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img_arm64.so.gz
	imgBlob    []byte
	imgLibName = "libSDL3_image.so.0"
)
