package mixer

import (
	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Mix_Version - This function gets the version of the dynamically linked SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Version)
func GetVersion() sdl.Version {
	return sdl.Version(iVersion())
}

// MIX_Init - Initialize the SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Init)
func Init() error {
	if !iInit() {
		return internal.LastErr()
	}

	return nil
}

// MIX_Quit - Deinitialize the SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Quit)
func Quit() {
	iQuit()
}

// MIX_GetNumAudioDecoders - Report the number of audio decoders available for use.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetNumAudioDecoders)
func NumAudioDecoders() int32 {
	return iGetNumAudioDecoders()
}

// MIX_GetAudioDecoder - Report the name of a specific audio decoders.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoder)
func GetAudioDecoder(index int32) string {
	return iGetAudioDecoder(index)
}

// MIX_CreateMixerDevice - Create a mixer that plays sound directly to an audio device.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixerDevice)
func CreateMixerDevice(devid sdl.AudioDeviceID, spec *sdl.AudioSpec) (*Mixer, error) {
	mix := iCreateMixerDevice(devid, spec)
	if mix == nil {
		return nil, internal.LastErr()
	}

	return mix, nil
}

// MIX_CreateMixer - Create a mixer that generates audio to a memory buffer.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateMixer)
func CreateMixer(spec *sdl.AudioSpec) (*Mixer, error) {
	mix := iCreateMixer(spec)
	if mix == nil {
		return nil, internal.LastErr()
	}

	return mix, nil
}

// MIX_LoadAudioWithProperties - Load audio for playback through a collection of properties.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudioWithProperties)
func LoadAudioWithProperties(props sdl.PropertiesID) (*Audio, error) {
	audio := iLoadAudioWithProperties(props)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

// MIX_MSToFrames - Convert milliseconds to sample frames at a specific sample rate.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_MSToFrames)
func MSToFrames(sampleRate int32, ms int64) int64 {
	return iMSToFrames(sampleRate, ms)
}

// MIX_FramesToMS - Convert sample frames, at a specific sample rate, to milliseconds.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_FramesToMS)
func FramesToMS(sampleRate int32, frames int64) int64 {
	return iFramesToMS(sampleRate, frames)
}

// MIX_CreateAudioDecoder - Create a [MIX_AudioDecoder](MIX_AudioDecoder) from a path on the filesystem.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder)
func CreateAudioDecoder(path string, props sdl.PropertiesID) (*AudioDecoder, error) {
	dec := iCreateAudioDecoder(path, props)
	if dec == nil {
		return nil, internal.LastErr()
	}

	return dec, nil
}

// MIX_CreateAudioDecoder_IO - Create a [MIX_AudioDecoder](MIX_AudioDecoder) from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateAudioDecoder_IO)
func CreateAudioDecoder_IO(stream *sdl.IOStream, closeIO bool, props sdl.PropertiesID) (*AudioDecoder, error) {
	dec := iCreateAudioDecoder_IO(stream, closeIO, props)
	if dec == nil {
		return nil, internal.LastErr()
	}

	return dec, nil
}
