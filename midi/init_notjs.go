//go:build !js

package midi

import (
	"runtime"

	puregogen "github.com/Zyko0/purego-gen"
)

// Path returns the library installation path based on the operating
// system
func Path() string {
	// SDL_native_midi sets no SOVERSION, so the Linux name carries no
	// trailing .0 unlike every other library in this repository.
	switch runtime.GOOS {
	case "windows":
		return "SDL_native_midi.dll"
	case "linux", "freebsd":
		return "libSDL_native_midi.so"
	case "darwin":
		return "libSDL_native_midi.dylib"
	default:
		return ""
	}
}

// LoadLibrary loads SDL_native_midi library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	runtime.LockOSThread()

	_hnd_midi, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()

	return nil
}

// CloseLibrary releases resources associated with the library.
func CloseLibrary() error {
	return puregogen.CloseLibrary(_hnd_midi)
}
