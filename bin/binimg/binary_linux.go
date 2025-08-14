//go:build linux

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img.so.gz
	imgBlob    []byte
	imgLibName = "libSDL3_image.so.0"
)
