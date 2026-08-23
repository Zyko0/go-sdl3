//go:build linux && arm64

package binmidi

import (
	_ "embed"
)

var (
	//go:embed assets/midi_arm64.so.gz
	midiBlob    []byte
	midiLibName = "libSDL_native_midi.so"
)
