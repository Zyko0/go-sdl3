//go:build windows && arm64

package binmix

import (
	_ "embed"
)

var (
	// TODO: not available for downloads apparently
	// //go:embed assets/img_arm64.dll

	mixBlob    []byte
	mixLibName = "SDL3_mixer.dll"
)

func init() {
	panic("unsupported mixer binary for windows arm64")
}
