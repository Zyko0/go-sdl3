package sdl

import puregogen "github.com/Zyko0/purego-gen"

// LoadLibrary loads SDL library and initializes all functions.
func LoadLibrary(path string) error {
	var err error

	_hnd_sdl, err = puregogen.OpenLibrary(path)
	if err != nil {
		return err
	}

	initialize()

	return nil
}

// CloseLibrary releases resources associated with the library.
func CloseLibrary() error {
	return puregogen.CloseLibrary(_hnd_sdl)
}