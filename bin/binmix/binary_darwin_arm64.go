//go:build darwin && arm64

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix_arm64.dylib
	mixBlob    []byte
	mixLibName = "libSDL3_mixer.dylib"
)
