package mixer

import internal "github.com/Zyko0/go-sdl3/internal"

// Chunk

// Mix_FreeChunk - Free an audio chunk.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FreeChunk)
func (chunk *Chunk) Free() {
	iFreeChunk(chunk)
}

// Mix_VolumeChunk - Set the volume for a specific chunk.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_VolumeChunk)
func (chunk *Chunk) SetVolume(volume int32) int32 {
	return iVolumeChunk(chunk, volume)
}

// Music

// Mix_FreeMusic - Free a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FreeMusic)
func (music *Music) Free() {
	iFreeMusic(music)
}

// Mix_GetMusicType - Find out the format of a mixer music.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicType)
func (music *Music) Type() MusicType {
	return iGetMusicType(music)
}

// Mix_GetMusicTitle - Get the title for a music object, or its filename.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitle)
func (music *Music) Title() string {
	return iGetMusicTitle(music)
}

// Mix_GetMusicTitleTag - Get the title for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicTitleTag)
func (music *Music) TitleTag() string {
	return iGetMusicTitleTag(music)
}

// Mix_GetMusicArtistTag - Get the artist name for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicArtistTag)
func (music *Music) ArtistTag() string {
	return iGetMusicArtistTag(music)
}

// Mix_GetMusicAlbumTag - Get the album name for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicAlbumTag)
func (music *Music) AlbumTag() string {
	return iGetMusicAlbumTag(music)
}

// Mix_GetMusicCopyrightTag - Get the copyright text for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicCopyrightTag)
func (music *Music) CopyrightTag() string {
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
func (music *Music) FadeIn(loops, ms int32) error {
	if !iFadeInMusic(music, loops, ms) {
		return internal.LastErr()
	}

	return nil
}

// Mix_FadeInMusicPos - Play a new music object, fading in the audio, from a starting position.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_FadeInMusicPos)
func (music *Music) FadeInPos(loops, ms int32, position float64) error {
	if !iFadeInMusicPos(music, loops, ms, position) {
		return internal.LastErr()
	}

	return nil
}

// Mix_GetMusicVolume - Query the current volume value for a music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicVolume)
func (music *Music) Volume() int32 {
	return iGetMusicVolume(music)
}

// Mix_StartTrack - Start a track in music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_StartTrack)
func (music *Music) StartTrack(track int32) error {
	if !iStartTrack(music, track) {
		return internal.LastErr()
	}

	return nil
}

// Mix_GetNumTracks - Get number of tracks in music object.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetNumTracks)
func (music *Music) NumTracks() int32 {
	return iGetNumTracks(music)
}

// Mix_GetMusicPosition - Get the time current position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicPosition)
func (music *Music) Position() float64 {
	return iGetMusicPosition(music)
}

// Mix_MusicDuration - Get a music object's duration, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_MusicDuration)
func (music *Music) Duration() float64 {
	return iMusicDuration(music)
}

// Mix_GetMusicLoopStartTime - Get the loop start time position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopStartTime)
func (music *Music) LoopStartTime() float64 {
	return iGetMusicLoopStartTime(music)
}

// Mix_GetMusicLoopEndTime - Get the loop end time position of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopEndTime)
func (music *Music) LoopEndTime() float64 {
	return iGetMusicLoopEndTime(music)
}

// Mix_GetMusicLoopLengthTime - Get the loop time length of music stream, in seconds.
// (https://wiki.libsdl.org/SDL3_mixer/Mix_GetMusicLoopLengthTime)
func (music *Music) LoopLengthTime() float64 {
	return iGetMusicLoopLengthTime(music)
}
