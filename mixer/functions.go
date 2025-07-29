package mixer

import internal "github.com/Zyko0/go-sdl3/internal"

func Version() int32 {
	return iVersion()
}

func Init() error {
	if !iInit() {
		return internal.LastErr()
	}

	return nil
}

func Quit() {
	iQuit()
}
