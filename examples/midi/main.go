// Plays a C major scale through the platform's native MIDI device.
//
// Unlike the other examples this one cannot run in CI, or unattended at all:
// headless Linux runners expose no ALSA sequencer and there is no device to
// play to.
package main

import (
	"encoding/binary"

	"github.com/Zyko0/go-sdl3/bin/binmidi"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/midi"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Ticks per quarter note. The track carries no tempo event, so players fall
// back to 120 BPM and a quarter note lasts half a second.
const division = 96

// song encodes notes as a format 0 MIDI file: a single track holding each note
// for a quarter note and releasing it before the next one begins.
func song(notes []byte) []byte {
	const velocity = 0x64

	var track []byte
	for _, note := range notes {
		// Both delta times fit in 7 bits, so each is a single byte rather than
		// a variable-length quantity.
		track = append(track,
			0x00, 0x90, note, velocity, // note on, no delay
			division, 0x80, note, 0x00, // note off, a quarter note later
		)
	}
	track = append(track, 0x00, 0xFF, 0x2F, 0x00) // end of track

	buf := []byte("MThd")
	buf = binary.BigEndian.AppendUint32(buf, 6) // bytes of header that follow
	buf = binary.BigEndian.AppendUint16(buf, 0) // format 0
	buf = binary.BigEndian.AppendUint16(buf, 1) // track count
	buf = binary.BigEndian.AppendUint16(buf, division)
	buf = append(buf, "MTrk"...)
	buf = binary.BigEndian.AppendUint32(buf, uint32(len(track)))

	return append(buf, track...)
}

func main() {
	defer binsdl.Load().Unload()  // sdl.LoadLibrary(sdl.Path())
	defer binmidi.Load().Unload() // midi.LoadLibrary(midi.Path())

	// midi drives the OS sequencer itself, so no SDL subsystem is needed.
	if err := midi.Init(); err != nil {
		panic(err)
	}
	defer midi.Quit()

	stream, err := sdl.IOFromConstMem(song([]byte{60, 62, 64, 65, 67, 69, 71, 72}))
	if err != nil {
		panic(err)
	}

	s, err := midi.LoadSong_IO(stream, true) // true: closes the stream for us
	if err != nil {
		panic(err)
	}
	defer s.Destroy()

	s.Start(0) // 0 plays the scale once
	// The ALSA backend drops SetVolume unless a song is already playing.
	midi.SetVolume(1)

	for midi.Active() {
		sdl.Delay(100)
	}
}
