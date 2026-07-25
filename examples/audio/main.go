// Verification harness for core SDL3 audio (not SDL3_mixer) on the
// WebAssembly (GOOS=js/GOARCH=wasm) path.
//
// It builds a WAV file in memory, decodes it with sdl.LoadWAV_IO (via an
// in-memory IOStream - no Emscripten filesystem needed), then queues the decoded
// PCM on each click/keypress. Browsers suspend audio until a user gesture, so
// you must click the window for sound to play.
//
// Run natively:  go run ./examples/audio
// Run in browser: go run ./cmd/wasmsdl serve ./examples/audio
package main

import (
	"bytes"
	"encoding/binary"
	"math"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/sdl"
)

const (
	sampleRate = 48000
	freq       = 440.0 // A4
	seconds    = 1.0
	amplitude  = 0.2
)

// makeWAV builds a mono 16-bit PCM WAV file (sine wave) entirely in memory, so
// sdl.LoadWAV_IO has something real to decode without any asset on disk.
func makeWAV() []byte {
	n := int(sampleRate * seconds)
	dataSize := n * 2 // 16-bit mono
	var b bytes.Buffer
	w := func(v any) { binary.Write(&b, binary.LittleEndian, v) }

	b.WriteString("RIFF")
	w(uint32(36 + dataSize))
	b.WriteString("WAVE")
	b.WriteString("fmt ")
	w(uint32(16))             // fmt chunk size
	w(uint16(1))              // PCM
	w(uint16(1))              // channels
	w(uint32(sampleRate))     // sample rate
	w(uint32(sampleRate * 2)) // byte rate = rate * blockAlign
	w(uint16(2))              // block align = channels * bytesPerSample
	w(uint16(16))             // bits per sample
	b.WriteString("data")
	w(uint32(dataSize))
	for i := 0; i < n; i++ {
		t := float64(i) / float64(sampleRate)
		s := int16(amplitude * math.Sin(2*math.Pi*freq*t) * 32767)
		w(s)
	}
	return b.Bytes()
}

func main() {
	defer binsdl.Load().Unload()
	defer sdl.Quit()

	if err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/audio", 480, 160, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	// Decode the in-memory WAV. spec is filled from the file's header.
	stream, err := sdl.IOFromConstMem(makeWAV())
	if err != nil {
		panic(err)
	}
	var spec sdl.AudioSpec
	pcm, err := sdl.LoadWAV_IO(stream, true, &spec)
	if err != nil {
		panic(err)
	}

	// Open a playback stream using the format reported by the WAV.
	audio := sdl.AUDIO_DEVICE_DEFAULT_PLAYBACK.OpenAudioDeviceStream(&spec, sdl.AudioStreamCallback(0))
	if audio == nil {
		panic("failed to open audio device stream")
	}
	defer audio.Destroy()
	if err := audio.ResumeDevice(); err != nil {
		panic(err)
	}

	sdl.RunLoop(func() error {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_QUIT:
				return sdl.EndLoop
			case sdl.EVENT_KEY_DOWN, sdl.EVENT_MOUSE_BUTTON_DOWN:
				// Queue the decoded PCM. (Browsers resume the audio context on
				// this same user gesture.)
				if err := audio.PutData(pcm); err != nil {
					return err
				}
			}
		}

		queued, _ := audio.Queued()

		renderer.SetDrawColor(20, 20, 30, 255)
		renderer.Clear()
		renderer.SetDrawColor(220, 220, 220, 255)
		renderer.DebugText(20, 40, "Click or press a key to play a decoded WAV tone.")
		if queued > 0 {
			renderer.DebugText(20, 70, "playing...")
		}
		renderer.Present()
		return nil
	})
}
