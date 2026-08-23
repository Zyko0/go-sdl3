//go:build windows && amd64

package binmidi

import (
	_ "embed"
)

var (
	//go:embed assets/midi_amd64.dll.gz
	midiBlob    []byte
	midiLibName = "SDL_native_midi.dll"
)
