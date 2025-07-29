package mixer

import internal "github.com/Zyko0/go-sdl3/internal"

// MIX_Version - Get the version of SDL_mixer that is linked against your program.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Version)
func Version() int32 {
	return iVersion()
}

// MIX_Init - Initialize the SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Init)
func Init() error {
	if !iInit() {
		return internal.LastErr()
	}

	return nil
}

// MIX_Quit - Deinitialize the SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Quit)
func Quit() {
	iQuit()
}
