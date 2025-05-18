package mixer

import internal "github.com/Zyko0/go-sdl3/internal"

// MusicFinishedCallback

// Mix_HookMusicFinished - Set a callback that runs when a music object has stopped playing.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusicFinished)
func (music_finished MusicFinishedCallback) HookMusicFinished() {
	panic("not implemented")
	iHookMusicFinished(music_finished)
}

// ChannelFinishedCallback

// Mix_ChannelFinished - Set a callback that runs when a channel has finished playing.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_ChannelFinished)
func (channel_finished ChannelFinishedCallback) ChannelFinished() {
	panic("not implemented")
	iChannelFinished(channel_finished)
}

// EachSoundFontCallback

// Mix_EachSoundFont - Iterate SoundFonts paths to use by supported MIDI backends.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_EachSoundFont)
func (function EachSoundFontCallback) EachSoundFont(data uintptr) bool {
	panic("not implemented")
	return iEachSoundFont(function, data)
}

// MIX_InitFlags

// Mix_Init - Initialize SDL_mixer.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_Init)
func (flags MIX_InitFlags) Init() MIX_InitFlags {
	panic("not implemented")
	return iInit(flags)
}

// Chunk

// Mix_FreeChunk - Free an audio chunk.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FreeChunk)
func (chunk *Chunk) Free() {
	iFreeChunk(chunk)
}

/*func (chunk *Chunk) Volume(volume int32) int32 { // TODO: same name
	panic("not implemented")
	return iVolumeChunk(chunk, volume)
}*/

// Music

// Mix_FreeMusic - Free a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FreeMusic)
func (music *Music) Free() {
	iFreeMusic(music)
}

// Mix_GetMusicType - Find out the format of a mixer music.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicType)
func (music *Music) Type() MusicType {
	panic("not implemented")
	return iGetMusicType(music)
}

// Mix_GetMusicTitle - Get the title for a music object, or its filename.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitle)
func (music *Music) Title() string {
	panic("not implemented")
	return iGetMusicTitle(music)
}

// Mix_GetMusicTitleTag - Get the title for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitleTag)
func (music *Music) TitleTag() string {
	panic("not implemented")
	return iGetMusicTitleTag(music)
}

// Mix_GetMusicArtistTag - Get the artist name for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicArtistTag)
func (music *Music) ArtistTag() string {
	panic("not implemented")
	return iGetMusicArtistTag(music)
}

// Mix_GetMusicAlbumTag - Get the album name for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicAlbumTag)
func (music *Music) AlbumTag() string {
	panic("not implemented")
	return iGetMusicAlbumTag(music)
}

// Mix_GetMusicCopyrightTag - Get the copyright text for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicCopyrightTag)
func (music *Music) CopyrightTag() string {
	panic("not implemented")
	return iGetMusicCopyrightTag(music)
}

// Mix_PlayMusic - Play a new music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_PlayMusic)
func (music *Music) Play(loops int32) error {
	if !iPlayMusic(music, loops) {
		return internal.LastErr()
	}

	return nil
}

// Mix_FadeInMusic - Play a new music object, fading in the audio.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusic)
func (music *Music) FadeIn(loops int32, ms int32) bool {
	panic("not implemented")
	return iFadeInMusic(music, loops, ms)
}

// Mix_FadeInMusicPos - Play a new music object, fading in the audio, from a starting position.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusicPos)
func (music *Music) FadeInPos(loops int32, ms int32, position float64) bool {
	panic("not implemented")
	return iFadeInMusicPos(music, loops, ms, position)
}

// Mix_GetMusicVolume - Query the current volume value for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicVolume)
func (music *Music) Volume() int32 {
	panic("not implemented")
	return iGetMusicVolume(music)
}

// Mix_StartTrack - Start a track in music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_StartTrack)
func (music *Music) StartTrack(track int32) bool {
	panic("not implemented")
	return iStartTrack(music, track)
}

// Mix_GetNumTracks - Get number of tracks in music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumTracks)
func (music *Music) NumTracks() int32 {
	panic("not implemented")
	return iGetNumTracks(music)
}

// Mix_GetMusicPosition - Get the time current position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicPosition)
func (music *Music) Position() float64 {
	panic("not implemented")
	return iGetMusicPosition(music)
}

// Mix_MusicDuration - Get a music object's duration, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_MusicDuration)
func (music *Music) Duration() float64 {
	panic("not implemented")
	return iMusicDuration(music)
}

// Mix_GetMusicLoopStartTime - Get the loop start time position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopStartTime)
func (music *Music) LoopStartTime() float64 {
	panic("not implemented")
	return iGetMusicLoopStartTime(music)
}

// Mix_GetMusicLoopEndTime - Get the loop end time position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopEndTime)
func (music *Music) LoopEndTime() float64 {
	panic("not implemented")
	return iGetMusicLoopEndTime(music)
}

// Mix_GetMusicLoopLengthTime - Get the loop time length of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopLengthTime)
func (music *Music) LoopLengthTime() float64 {
	panic("not implemented")
	return iGetMusicLoopLengthTime(music)
}

// MixCallback

// Mix_SetPostMix - Set a function that is called after all mixing is performed.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_SetPostMix)
func (mix_func MixCallback) SetPostMix(arg uintptr) {
	panic("not implemented")
	iSetPostMix(mix_func, arg)
}

// Mix_HookMusic - Add your own music player or additional mixer function.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_HookMusic)
func (mix_func MixCallback) HookMusic(arg uintptr) {
	panic("not implemented")
	iHookMusic(mix_func, arg)
}
