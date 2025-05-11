//go:build !js

package sdl

import (
	"errors"
	"runtime"

	"github.com/Zyko0/go-sdl3/internal"
	puregogen "github.com/Zyko0/purego-gen"
)

func init() {
	runtime.LockOSThread()
}

// Path returns the library installation path based on the operating
// system
func Path() string {
	switch runtime.GOOS {
	case "windows":
		return "SDL3.dll"
	case "linux", "freebsd":
		return "libSDL3.so.0"
	case "darwin":
		return "libSDL3.dylib"
	default:
		return ""
	}
}

// LoadLibrary loads SDL library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	_hnd_sdl, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()
	initialize_ex()

	// Set free, error functions
	internal.SetSDLFreeFunc(ifree)
	internal.SetSDLLastErrFunc(func() error {
		if msg := iGetError(); msg != "" {
			return errors.New(msg)
		}
		return nil
	})

	return nil
}

// CloseLibrary releases resources associated with the library.
func CloseLibrary() error {
	return puregogen.CloseLibrary(_hnd_sdl)
}

func RunLoop(updateFunc func() error) error {
	for {
		if err := updateFunc(); err != nil {
			if errors.Is(err, EndLoop) {
				return nil
			}
			return err
		}
	}
}
