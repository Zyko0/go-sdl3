//go:build !js

package ttf

import (
	"runtime"

	puregogen "github.com/Zyko0/purego-gen"
)

// Path returns the library installation path based on the operating
// system
func Path() string {
	switch runtime.GOOS {
	case "windows":
		return "SDL3_ttf.dll"
	case "linux", "freebsd":
		return "libSDL3_ttf.so.0"
	case "darwin":
		return "libSDL3_ttf.dylib"
	default:
		return ""
	}
}

// LoadLibrary loads SDL_ttf library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	runtime.LockOSThread()

	_hnd_ttf, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()
	initialize_ex()

	return nil
}

// CloseLibrary releases resources associated with the library.
func CloseLibrary() error {
	return puregogen.CloseLibrary(_hnd_ttf)
}
