//go:build windows && arm64

package binmidi

import (
	_ "embed"
)

var (
	//go:embed assets/midi_arm64.dll.gz
	midiBlob    []byte
	midiLibName = "SDL_native_midi.dll"
)
