//go:build js

package binmidi

import "github.com/Zyko0/go-sdl3/midi"

type library struct{}

func Load() library {
	return library{}
}

func (l library) Unload() {
	midi.CloseLibrary()
}
