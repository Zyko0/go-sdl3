package mixer

import sdl "github.com/Zyko0/go-sdl3"

// Code generated by cmd/ffi2go. DO NOT EDIT.

var (
	// Mix_Version => This function gets the version of the dynamically linked SDL_mixer library.
	//
	//puregogen:library path:windows=mixer.dll path:unix=mixer.so alias=mixer
	//puregogen:function symbol=Mix_Version
	iVersion func() int32

	// Mix_Init => Initialize SDL_mixer.
	//
	//puregogen:function symbol=Mix_Init
	iInit func(flags MIX_InitFlags) MIX_InitFlags

	// Mix_Quit => Deinitialize SDL_mixer.
	//
	//puregogen:function symbol=Mix_Quit
	iQuit func()

	// Mix_OpenAudio => Open an audio device for playback.
	//
	//puregogen:function symbol=Mix_OpenAudio
	iOpenAudio func(devid sdl.AudioDeviceID, spec *sdl.AudioSpec) bool

	// Mix_PauseAudio => Suspend or resume the whole audio output.
	//
	//puregogen:function symbol=Mix_PauseAudio
	iPauseAudio func(pause_on int32)

	// Mix_QuerySpec => Find out what the actual audio device parameters are.
	//
	//puregogen:function symbol=Mix_QuerySpec
	iQuerySpec func(frequency *int32, format *sdl.AudioFormat, channels *int32) bool

	// Mix_AllocateChannels => Dynamically change the number of channels managed by the mixer.
	//
	//puregogen:function symbol=Mix_AllocateChannels
	iAllocateChannels func(numchans int32) int32

	// Mix_LoadWAV_IO => Load a supported audio format into a chunk.
	//
	//puregogen:function symbol=Mix_LoadWAV_IO
	iLoadWAV_IO func(src *sdl.IOStream, closeio bool) *Chunk

	// Mix_LoadWAV => Load a supported audio format into a chunk.
	//
	//puregogen:function symbol=Mix_LoadWAV
	iLoadWAV func(file string) *Chunk

	// Mix_LoadMUS => Load a supported audio format into a music object.
	//
	//puregogen:function symbol=Mix_LoadMUS
	iLoadMUS func(file string) *Music

	// Mix_LoadMUS_IO => Load a supported audio format into a music object.
	//
	//puregogen:function symbol=Mix_LoadMUS_IO
	iLoadMUS_IO func(src *sdl.IOStream, closeio bool) *Music

	// Mix_LoadMUSType_IO => Load an audio format into a music object, assuming a specific format.
	//
	//puregogen:function symbol=Mix_LoadMUSType_IO
	iLoadMUSType_IO func(src *sdl.IOStream, typ MusicType, closeio bool) *Music

	// Mix_QuickLoad_WAV => Load a WAV file from memory as quickly as possible.
	//
	//puregogen:function symbol=Mix_QuickLoad_WAV
	iQuickLoad_WAV func(mem *uint8) *Chunk

	// Mix_QuickLoad_RAW => Load a raw audio data from memory as quickly as possible.
	//
	//puregogen:function symbol=Mix_QuickLoad_RAW
	iQuickLoad_RAW func(mem *uint8, len uint32) *Chunk

	// Mix_FreeChunk => Free an audio chunk.
	//
	//puregogen:function symbol=Mix_FreeChunk
	iFreeChunk func(chunk *Chunk)

	// Mix_FreeMusic => Free a music object.
	//
	//puregogen:function symbol=Mix_FreeMusic
	iFreeMusic func(music *Music)

	// Mix_GetNumChunkDecoders => Get a list of chunk decoders that this build of SDL_mixer provides.
	//
	//puregogen:function symbol=Mix_GetNumChunkDecoders
	iGetNumChunkDecoders func() int32

	// Mix_GetChunkDecoder => Get a chunk decoder's name.
	//
	//puregogen:function symbol=Mix_GetChunkDecoder
	iGetChunkDecoder func(index int32) string

	// Mix_HasChunkDecoder => Check if a chunk decoder is available by name.
	//
	//puregogen:function symbol=Mix_HasChunkDecoder
	iHasChunkDecoder func(name string) bool

	// Mix_GetNumMusicDecoders => Get a list of music decoders that this build of SDL_mixer provides.
	//
	//puregogen:function symbol=Mix_GetNumMusicDecoders
	iGetNumMusicDecoders func() int32

	// Mix_GetMusicDecoder => Get a music decoder's name.
	//
	//puregogen:function symbol=Mix_GetMusicDecoder
	iGetMusicDecoder func(index int32) string

	// Mix_HasMusicDecoder => Check if a music decoder is available by name.
	//
	//puregogen:function symbol=Mix_HasMusicDecoder
	iHasMusicDecoder func(name string) bool

	// Mix_GetMusicType => Find out the format of a mixer music.
	//
	//puregogen:function symbol=Mix_GetMusicType
	iGetMusicType func(music *Music) MusicType

	// Mix_GetMusicTitle => Get the title for a music object, or its filename.
	//
	//puregogen:function symbol=Mix_GetMusicTitle
	iGetMusicTitle func(music *Music) string

	// Mix_GetMusicTitleTag => Get the title for a music object.
	//
	//puregogen:function symbol=Mix_GetMusicTitleTag
	iGetMusicTitleTag func(music *Music) string

	// Mix_GetMusicArtistTag => Get the artist name for a music object.
	//
	//puregogen:function symbol=Mix_GetMusicArtistTag
	iGetMusicArtistTag func(music *Music) string

	// Mix_GetMusicAlbumTag => Get the album name for a music object.
	//
	//puregogen:function symbol=Mix_GetMusicAlbumTag
	iGetMusicAlbumTag func(music *Music) string

	// Mix_GetMusicCopyrightTag => Get the copyright text for a music object.
	//
	//puregogen:function symbol=Mix_GetMusicCopyrightTag
	iGetMusicCopyrightTag func(music *Music) string

	// Mix_SetPostMix => Set a function that is called after all mixing is performed.
	//
	//puregogen:function symbol=Mix_SetPostMix
	iSetPostMix func(mix_func MixCallback, arg uintptr)

	// Mix_HookMusic => Add your own music player or additional mixer function.
	//
	//puregogen:function symbol=Mix_HookMusic
	iHookMusic func(mix_func MixCallback, arg uintptr)

	// Mix_HookMusicFinished => Set a callback that runs when a music object has stopped playing.
	//
	//puregogen:function symbol=Mix_HookMusicFinished
	iHookMusicFinished func(music_finished MusicFinishedCallback)

	// Mix_GetMusicHookData => Get a pointer to the user data for the current music hook.
	//
	//puregogen:function symbol=Mix_GetMusicHookData
	iGetMusicHookData func() uintptr

	// Mix_ChannelFinished => Set a callback that runs when a channel has finished playing.
	//
	//puregogen:function symbol=Mix_ChannelFinished
	iChannelFinished func(channel_finished ChannelFinishedCallback)

	// Mix_RegisterEffect => Register a special effect function.
	//
	//puregogen:function symbol=Mix_RegisterEffect
	iRegisterEffect func(chann int32, f EffectFunc_t, d EffectDone_t, arg uintptr) bool

	// Mix_UnregisterEffect => Explicitly unregister a special effect function.
	//
	//puregogen:function symbol=Mix_UnregisterEffect
	iUnregisterEffect func(channel int32, f EffectFunc_t) bool

	// Mix_UnregisterAllEffects => Explicitly unregister all special effect functions.
	//
	//puregogen:function symbol=Mix_UnregisterAllEffects
	iUnregisterAllEffects func(channel int32) bool

	// Mix_SetPanning => Set the panning of a channel.
	//
	//puregogen:function symbol=Mix_SetPanning
	iSetPanning func(channel int32, left uint8, right uint8) bool

	// Mix_SetPosition => Set the position of a channel.
	//
	//puregogen:function symbol=Mix_SetPosition
	iSetPosition func(channel int32, angle int16, distance uint8) bool

	// Mix_SetDistance => Set the "distance" of a channel.
	//
	//puregogen:function symbol=Mix_SetDistance
	iSetDistance func(channel int32, distance uint8) bool

	// Mix_SetReverseStereo => Cause a channel to reverse its stereo.
	//
	//puregogen:function symbol=Mix_SetReverseStereo
	iSetReverseStereo func(channel int32, flip int32) bool

	// Mix_ReserveChannels => Reserve the first channels for the application.
	//
	//puregogen:function symbol=Mix_ReserveChannels
	iReserveChannels func(num int32) int32

	// Mix_GroupChannel => Assign a tag to a channel.
	//
	//puregogen:function symbol=Mix_GroupChannel
	iGroupChannel func(which int32, tag int32) bool

	// Mix_GroupChannels => Assign several consecutive channels to the same tag.
	//
	//puregogen:function symbol=Mix_GroupChannels
	iGroupChannels func(from int32, to int32, tag int32) bool

	// Mix_GroupAvailable => Finds the first available channel in a group of channels.
	//
	//puregogen:function symbol=Mix_GroupAvailable
	iGroupAvailable func(tag int32) int32

	// Mix_GroupCount => Returns the number of channels in a group.
	//
	//puregogen:function symbol=Mix_GroupCount
	iGroupCount func(tag int32) int32

	// Mix_GroupOldest => Find the "oldest" sample playing in a group of channels.
	//
	//puregogen:function symbol=Mix_GroupOldest
	iGroupOldest func(tag int32) int32

	// Mix_GroupNewer => Find the "most recent" sample playing in a group of channels.
	//
	//puregogen:function symbol=Mix_GroupNewer
	iGroupNewer func(tag int32) int32

	// Mix_PlayChannel => Play an audio chunk on a specific channel.
	//
	//puregogen:function symbol=Mix_PlayChannel
	iPlayChannel func(channel int32, chunk *Chunk, loops int32) int32

	// Mix_PlayChannelTimed => Play an audio chunk on a specific channel for a maximum time.
	//
	//puregogen:function symbol=Mix_PlayChannelTimed
	iPlayChannelTimed func(channel int32, chunk *Chunk, loops int32, ticks int32) int32

	// Mix_PlayMusic => Play a new music object.
	//
	//puregogen:function symbol=Mix_PlayMusic
	iPlayMusic func(music *Music, loops int32) bool

	// Mix_FadeInMusic => Play a new music object, fading in the audio.
	//
	//puregogen:function symbol=Mix_FadeInMusic
	iFadeInMusic func(music *Music, loops int32, ms int32) bool

	// Mix_FadeInMusicPos => Play a new music object, fading in the audio, from a starting position.
	//
	//puregogen:function symbol=Mix_FadeInMusicPos
	iFadeInMusicPos func(music *Music, loops int32, ms int32, position float64) bool

	// Mix_FadeInChannel => Play an audio chunk on a specific channel, fading in the audio.
	//
	//puregogen:function symbol=Mix_FadeInChannel
	iFadeInChannel func(channel int32, chunk *Chunk, loops int32, ms int32) int32

	// Mix_FadeInChannelTimed => Play an audio chunk on a specific channel, fading in the audio, for a maximum time.
	//
	//puregogen:function symbol=Mix_FadeInChannelTimed
	iFadeInChannelTimed func(channel int32, chunk *Chunk, loops int32, ms int32, ticks int32) int32

	// Mix_Volume => Set the volume for a specific channel.
	//
	//puregogen:function symbol=Mix_Volume
	iVolume func(channel int32, volume int32) int32

	// Mix_VolumeChunk => Set the volume for a specific chunk.
	//
	//puregogen:function symbol=Mix_VolumeChunk
	iVolumeChunk func(chunk *Chunk, volume int32) int32

	// Mix_VolumeMusic => Set the volume for the music channel.
	//
	//puregogen:function symbol=Mix_VolumeMusic
	iVolumeMusic func(volume int32) int32

	// Mix_GetMusicVolume => Query the current volume value for a music object.
	//
	//puregogen:function symbol=Mix_GetMusicVolume
	iGetMusicVolume func(music *Music) int32

	// Mix_MasterVolume => Set the master volume for all channels.
	//
	//puregogen:function symbol=Mix_MasterVolume
	iMasterVolume func(volume int32) int32

	// Mix_HaltChannel => Halt playing of a particular channel.
	//
	//puregogen:function symbol=Mix_HaltChannel
	iHaltChannel func(channel int32)

	// Mix_HaltGroup => Halt playing of a group of channels by arbitrary tag.
	//
	//puregogen:function symbol=Mix_HaltGroup
	iHaltGroup func(tag int32)

	// Mix_HaltMusic => Halt playing of the music stream.
	//
	//puregogen:function symbol=Mix_HaltMusic
	iHaltMusic func()

	// Mix_ExpireChannel => Change the expiration delay for a particular channel.
	//
	//puregogen:function symbol=Mix_ExpireChannel
	iExpireChannel func(channel int32, ticks int32) int32

	// Mix_FadeOutChannel => Halt a channel after fading it out for a specified time.
	//
	//puregogen:function symbol=Mix_FadeOutChannel
	iFadeOutChannel func(which int32, ms int32) int32

	// Mix_FadeOutGroup => Halt a playing group of channels by arbitrary tag, after fading them out for a specified time.
	//
	//puregogen:function symbol=Mix_FadeOutGroup
	iFadeOutGroup func(tag int32, ms int32) int32

	// Mix_FadeOutMusic => Halt the music stream after fading it out for a specified time.
	//
	//puregogen:function symbol=Mix_FadeOutMusic
	iFadeOutMusic func(ms int32) bool

	// Mix_FadingMusic => Query the fading status of the music stream.
	//
	//puregogen:function symbol=Mix_FadingMusic
	iFadingMusic func() Fading

	// Mix_FadingChannel => Query the fading status of a channel.
	//
	//puregogen:function symbol=Mix_FadingChannel
	iFadingChannel func(which int32) Fading

	// Mix_Pause => Pause a particular channel.
	//
	//puregogen:function symbol=Mix_Pause
	iPause func(channel int32)

	// Mix_PauseGroup => Pause playing of a group of channels by arbitrary tag.
	//
	//puregogen:function symbol=Mix_PauseGroup
	iPauseGroup func(tag int32)

	// Mix_Resume => Resume a particular channel.
	//
	//puregogen:function symbol=Mix_Resume
	iResume func(channel int32)

	// Mix_ResumeGroup => Resume playing of a group of channels by arbitrary tag.
	//
	//puregogen:function symbol=Mix_ResumeGroup
	iResumeGroup func(tag int32)

	// Mix_Paused => Query whether a particular channel is paused.
	//
	//puregogen:function symbol=Mix_Paused
	iPaused func(channel int32) int32

	// Mix_PauseMusic => Pause the music stream.
	//
	//puregogen:function symbol=Mix_PauseMusic
	iPauseMusic func()

	// Mix_ResumeMusic => Resume the music stream.
	//
	//puregogen:function symbol=Mix_ResumeMusic
	iResumeMusic func()

	// Mix_RewindMusic => Rewind the music stream.
	//
	//puregogen:function symbol=Mix_RewindMusic
	iRewindMusic func()

	// Mix_PausedMusic => Query whether the music stream is paused.
	//
	//puregogen:function symbol=Mix_PausedMusic
	iPausedMusic func() bool

	// Mix_ModMusicJumpToOrder => Jump to a given order in mod music.
	//
	//puregogen:function symbol=Mix_ModMusicJumpToOrder
	iModMusicJumpToOrder func(order int32) bool

	// Mix_StartTrack => Start a track in music object.
	//
	//puregogen:function symbol=Mix_StartTrack
	iStartTrack func(music *Music, track int32) bool

	// Mix_GetNumTracks => Get number of tracks in music object.
	//
	//puregogen:function symbol=Mix_GetNumTracks
	iGetNumTracks func(music *Music) int32

	// Mix_SetMusicPosition => Set the current position in the music stream, in seconds.
	//
	//puregogen:function symbol=Mix_SetMusicPosition
	iSetMusicPosition func(position float64) bool

	// Mix_GetMusicPosition => Get the time current position of music stream, in seconds.
	//
	//puregogen:function symbol=Mix_GetMusicPosition
	iGetMusicPosition func(music *Music) float64

	// Mix_MusicDuration => Get a music object's duration, in seconds.
	//
	//puregogen:function symbol=Mix_MusicDuration
	iMusicDuration func(music *Music) float64

	// Mix_GetMusicLoopStartTime => Get the loop start time position of music stream, in seconds.
	//
	//puregogen:function symbol=Mix_GetMusicLoopStartTime
	iGetMusicLoopStartTime func(music *Music) float64

	// Mix_GetMusicLoopEndTime => Get the loop end time position of music stream, in seconds.
	//
	//puregogen:function symbol=Mix_GetMusicLoopEndTime
	iGetMusicLoopEndTime func(music *Music) float64

	// Mix_GetMusicLoopLengthTime => Get the loop time length of music stream, in seconds.
	//
	//puregogen:function symbol=Mix_GetMusicLoopLengthTime
	iGetMusicLoopLengthTime func(music *Music) float64

	// Mix_Playing => Check the playing status of a specific channel.
	//
	//puregogen:function symbol=Mix_Playing
	iPlaying func(channel int32) int32

	// Mix_PlayingMusic => Check the playing status of the music stream.
	//
	//puregogen:function symbol=Mix_PlayingMusic
	iPlayingMusic func() bool

	// Mix_SetSoundFonts => Set SoundFonts paths to use by supported MIDI backends.
	//
	//puregogen:function symbol=Mix_SetSoundFonts
	iSetSoundFonts func(paths string) bool

	// Mix_GetSoundFonts => Get SoundFonts paths to use by supported MIDI backends.
	//
	//puregogen:function symbol=Mix_GetSoundFonts
	iGetSoundFonts func() string

	// Mix_EachSoundFont => Iterate SoundFonts paths to use by supported MIDI backends.
	//
	//puregogen:function symbol=Mix_EachSoundFont
	iEachSoundFont func(function EachSoundFontCallback, data uintptr) bool

	// Mix_SetTimidityCfg => Set full path of the Timidity config file.
	//
	//puregogen:function symbol=Mix_SetTimidityCfg
	iSetTimidityCfg func(path string) bool

	// Mix_GetTimidityCfg => Get full path of a previously-specified Timidity config file.
	//
	//puregogen:function symbol=Mix_GetTimidityCfg
	iGetTimidityCfg func() string

	// Mix_GetChunk => Get the Mix_Chunk currently associated with a mixer channel.
	//
	//puregogen:function symbol=Mix_GetChunk
	iGetChunk func(channel int32) *Chunk

	// Mix_CloseAudio => Close the mixer, halting all playing audio.
	//
	//puregogen:function symbol=Mix_CloseAudio
	iCloseAudio func()
)
