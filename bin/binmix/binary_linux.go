//go:build linux

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix.so
	mixBlob    []byte
	mixLibName = "libSDL3_mixer.so.0"
)
