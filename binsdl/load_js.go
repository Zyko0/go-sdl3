//go:build js

package binsdl

import sdl "github.com/Zyko0/go-sdl3"

type library struct{}

func Load() library {
	sdl.LoadLibrary("")
	return library{}
}

func (l library) Unload() {
	sdl.CloseLibrary()
}
