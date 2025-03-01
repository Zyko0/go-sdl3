//go:build !js

package sdl

import (
	"errors"
	"runtime"

	"github.com/Zyko0/go-sdl3/internal"
	puregogen "github.com/Zyko0/purego-gen"
)

// LoadLibrary loads SDL library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	runtime.LockOSThread()

	_hnd_sdl, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()

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
