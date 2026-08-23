//go:build darwin && arm64

package binmidi

import (
	_ "embed"
)

var (
	//go:embed assets/midi_arm64.dylib.gz
	midiBlob    []byte
	midiLibName = "libSDL_native_midi.dylib"
)
