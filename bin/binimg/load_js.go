//go:build js

package binimg

import "github.com/Zyko0/go-sdl3/img"

type library struct{}

func Load() library {
	return library{}
}

func (l library) Unload() {
	img.CloseLibrary()
}
