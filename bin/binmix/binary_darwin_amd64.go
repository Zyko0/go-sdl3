//go:build darwin && amd64

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix_amd64.dylib.gz
	mixBlob    []byte
	mixLibName = "libSDL3_mixer.dylib"
)
