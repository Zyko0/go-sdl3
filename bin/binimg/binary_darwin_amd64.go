//go:build darwin && amd64

package binimg

import (
	_ "embed"
)

var (
	//go:embed assets/img_amd64.dylib.gz
	imgBlob    []byte
	imgLibName = "libSDL3_image.dylib"
)
