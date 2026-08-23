package midi

import (
	"errors"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// SDL_native_midi only reports failures through SDL_SetError on its ALSA and
// dummy backends; the macOS and Windows ones never call it. Without these,
// internal.LastErr() would report a failure as success on those platforms.
var (
	ErrInit     = errors.New("midi: could not initialize the native midi subsystem")
	ErrLoadSong = errors.New("midi: could not load song")
)

// lastErr prefers the message SDL recorded, falling back to err when it holds
// none. On macOS and Windows the message can predate this call, since neither
// backend sets one of its own.
func lastErr(err error) error {
	if last := internal.LastErr(); last != nil {
		return last
	}

	return err
}

// NativeMidi_Init - Initialize the native MIDI subsystem.
func Init() error {
	if !iInit() {
		return lastErr(ErrInit)
	}

	return nil
}

// NativeMidi_Quit - Shut down the native MIDI subsystem.
func Quit() {
	iQuit()
}

// NativeMidi_LoadSong - Load a MIDI song from a file path.
func LoadSong(path string) (*Song, error) {
	song := iLoadSong(path)
	if song == nil {
		return nil, lastErr(ErrLoadSong)
	}

	return song, nil
}

// NativeMidi_LoadSong_IO - Load a MIDI song from an SDL_IOStream.
func LoadSong_IO(stream *sdl.IOStream, closeIO bool) (*Song, error) {
	song := iLoadSong_IO(stream, closeIO)
	if song == nil {
		return nil, lastErr(ErrLoadSong)
	}

	return song, nil
}

// NativeMidi_Pause - Pause playback of the current song.
func Pause() {
	iPause()
}

// NativeMidi_Resume - Resume playback of a paused song.
func Resume() {
	iResume()
}

// NativeMidi_Stop - Stop playback of the current song.
func Stop() {
	iStop()
}

// NativeMidi_Active - Query whether a song is currently playing.
func Active() bool {
	return iActive()
}

// NativeMidi_SetVolume - Set the volume applied to MIDI playback.
func SetVolume(volume float32) {
	iSetVolume(volume)
}
