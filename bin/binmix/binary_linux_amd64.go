//go:build linux && amd64

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix_amd64.so.gz
	mixBlob    []byte
	mixLibName = "libSDL3_mixer.so.0"
)
