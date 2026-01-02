//go:build js

package binmix

import "github.com/Zyko0/go-sdl3/mixer"

type library struct{}

func Load() library {
	return library{}
}

func (l library) Unload() {
	mixer.CloseLibrary()
}
