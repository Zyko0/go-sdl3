//go:build js

package binsdl

type library struct{}

func Load() library {
	return library{}
}

func (l library) Unload() {
}
