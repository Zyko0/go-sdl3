//go:build windows && amd64

package binmix

import (
	_ "embed"
)

var (
	// //go:embed assets/mix_amd64.dll
	mixBlob    []byte
	mixLibName = "SDL3_mixer.dll"
)

func init() {
	panic("unsupported mixer binary for windows amd64")
}
