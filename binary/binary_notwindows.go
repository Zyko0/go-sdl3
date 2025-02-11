//go:build !windows

package binary

import (
	_ "embed"
	"log"
	"runtime"
)

// TODO: handle static binaries for other platforms

var (
	// //go:embed assets/sdl.so
	sdlBlob []byte
)

func init() {
	log.Fatal("binary: " + runtime.GOOS + " not supported yet for static binary")
}
