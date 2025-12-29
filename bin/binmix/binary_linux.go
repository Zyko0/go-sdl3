//go:build linux

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix.so.gz
	mixBlob    []byte
	mixLibName = "libSDL3_mixer.so.0"
)
