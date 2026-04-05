//go:build windows && arm64

package binmix

import (
	_ "embed"
)

var (
	//go:embed assets/mix_amd64.dll.gz
	mixBlob    []byte
	mixLibName = "SDL3_mixer.dll"
)