//go:build !js

package img

import (
	"runtime"

	puregogen "github.com/Zyko0/purego-gen"
)

// Path returns the library installation path based on the operating
// system
func Path() string {
	switch runtime.GOOS {
	case "windows":
		return "SDL3_image.dll"
	case "linux", "freebsd":
		return "libSDL3_image.so.0"
	case "darwin":
		return "libSDL3_image.dylib"
	default:
		return ""
	}
}

// LoadLibrary loads SDL_image library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	runtime.LockOSThread()

	_hnd_img, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()

	return nil
}

// CloseLibrary releases resources associated with the library.
func CloseLibrary() error {
	return puregogen.CloseLibrary(_hnd_img)
}
