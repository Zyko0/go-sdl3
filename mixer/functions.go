package mixer

import (
	"unsafe"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Mix_Version - This function gets the version of the dynamically linked SDL_mixer library.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Version)
func Version() int32 {
	return iVersion()
}

// Mix_Init - Initialize SDL_mixer.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Init)
func Init(flags InitFlags) InitFlags {
	return InitFlags(iInit(MIX_InitFlags(flags)))
}

// Mix_Quit - Deinitialize SDL_mixer.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Quit)
func Quit() {
	iQuit()
}

// Mix_OpenAudio - Open an audio device for playback.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_OpenAudio)
func OpenAudio(devid sdl.AudioDeviceID, spec *sdl.AudioSpec) error {
	if !iOpenAudio(devid, spec) {
		return internal.LastErr()
	}

	return nil
}

// Mix_PauseAudio - Suspend or resume the whole audio output.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PauseAudio)
func PauseAudio() {
	iPauseAudio(1)
}

// ResumeAudio calls Mix_PauseAudio with argument 0 to resume
func ResumeAudio() {
	iPauseAudio(0)
}

// Mix_QuerySpec - Find out what the actual audio device parameters are.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_QuerySpec)
func QuerySpec() *AudioParams {
	var params AudioParams

	if !iQuerySpec(&params.Frequency, &params.Format, &params.Channels) {
		return nil
	}

	return &params
}

// Mix_AllocateChannels - Dynamically change the number of channels managed by the mixer.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_AllocateChannels)
func AllocateChannels(numChans int32) int32 {
	return iAllocateChannels(numChans)
}

// Mix_LoadWAV_IO - Load a supported audio format into a chunk.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV_IO)
func LoadWAV_IO(src *sdl.IOStream, closeio bool) (*Chunk, error) {
	chunk := iLoadWAV_IO(src, closeio)
	if chunk == nil {
		return nil, internal.LastErr()
	}

	return chunk, nil
}

// Mix_LoadWAV - Load a supported audio format into a chunk.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_LoadWAV)
func LoadWAV(file string) (*Chunk, error) {
	chunk := iLoadWAV(file)
	if chunk == nil {
		return nil, internal.LastErr()
	}

	return chunk, nil
}

// Mix_LoadMUS_IO - Load a supported audio format into a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS_IO)
func LoadMUS_IO(src *sdl.IOStream, closeio bool) (*Music, error) {
	music := iLoadMUS_IO(src, closeio)
	if music == nil {
		return nil, internal.LastErr()
	}

	return music, nil
}

// Mix_LoadMUS - Load a supported audio format into a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUS)
func LoadMUS(file string) (*Music, error) {
	music := iLoadMUS(file)
	if music == nil {
		return nil, internal.LastErr()
	}

	return music, nil
}

// Mix_LoadMUSType_IO - Load an audio format into a music object, assuming a specific format.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_LoadMUSType_IO)
func LoadMUSType_IO(src *sdl.IOStream, typ MusicType, closeio bool) (*Music, error) {
	music := iLoadMUSType_IO(src, typ, closeio)
	if music == nil {
		return nil, internal.LastErr()
	}

	return music, nil
}

// Mix_QuickLoad_WAV - Load a WAV file from memory as quickly as possible.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_WAV)
func QuickLoadWAV(mem []byte) (*Chunk, error) {
	chunk := iQuickLoad_WAV(unsafe.SliceData(mem))
	if chunk == nil {
		return nil, internal.LastErr()
	}

	return chunk, nil
}

// Mix_QuickLoad_RAW - Load a raw audio data from memory as quickly as possible.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_QuickLoad_RAW)
func QuickLoadRAW(mem []byte) (*Chunk, error) {
	chunk := iQuickLoad_RAW(unsafe.SliceData(mem), uint32(len(mem)))
	if chunk == nil {
		return nil, internal.LastErr()
	}

	return chunk, nil
}

// Mix_GetNumChunkDecoders - Get a list of chunk decoders that this build of SDL_mixer provides.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumChunkDecoders)
func NumChunkDecoders() int32 {
	return iGetNumChunkDecoders()
}

// Mix_GetChunkDecoder - Get a chunk decoder's name.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetChunkDecoder)
func GetChunkDecoder(index int32) string {
	return iGetChunkDecoder(index)
}

// Mix_HasChunkDecoder - Check if a chunk decoder is available by name.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_HasChunkDecoder)
func HasChunkDecoder(name string) bool {
	return iHasChunkDecoder(name)
}

// Mix_GetNumMusicDecoders - Get a list of music decoders that this build of SDL_mixer provides.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumMusicDecoders)
func NumMusicDecoders() int32 {
	return iGetNumMusicDecoders()
}

// Mix_GetMusicDecoder - Get a music decoder's name.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicDecoder)
func GetMusicDecoder(index int32) string {
	return iGetMusicDecoder(index)
}

// Mix_HasMusicDecoder - Check if a music decoder is available by name.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_HasMusicDecoder)
func HasMusicDecoder(name string) bool {
	return iHasMusicDecoder(name)
}

// Mix_PlayChannel - Play an audio chunk on a specific channel.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PlayChannel)
func PlayChannel(channel int32, chunk *Chunk, loops int32) int32 {
	return iPlayChannel(channel, chunk, loops)
}

// Mix_PlayingMusic - Check the playing status of the music stream.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PlayingMusic)
func PlayingMusic() bool {
	return iPlayingMusic()
}

// Mix_PausedMusic - Query whether the music stream is paused.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PausedMusic)
func PausedMusic() bool {
	return iPausedMusic()
}

// Mix_PauseMusic - Pause the music stream.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PauseMusic)
func PauseMusic() {
	iPauseMusic()
}

// Mix_ResumeMusic - Resume the music stream.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_ResumeMusic)
func ResumeMusic() {
	iResumeMusic()
}

// Mix_HaltMusic - Halt playing of the music stream.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_HaltMusic)
func HaltMusic() {
	iHaltMusic()
}
