//go:build darwin && amd64

package binmidi

import (
	_ "embed"
)

var (
	//go:embed assets/midi_amd64.dylib.gz
	midiBlob    []byte
	midiLibName = "libSDL_native_midi.dylib"
)
