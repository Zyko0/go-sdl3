package mixer

import (
	"runtime"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Group

// MIX_DestroyGroup - Destroy a mixing group.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyGroup)
func (group *Group) Destroy() {
	iDestroyGroup(group)
}

// MIX_GetGroupProperties - Get the properties associated with a group.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupProperties)
func (group *Group) Properties() (sdl.PropertiesID, error) {
	props := iGetGroupProperties(group)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// MIX_GetGroupMixer - Get the [MIX_Mixer](MIX_Mixer) that owns a [MIX_Group](MIX_Group).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetGroupMixer)
func (group *Group) Mixer() *Mixer {
	return iGetGroupMixer(group)
}

// MIX_SetGroupPostMixCallback - Set a callback that fires when a mixer group has completed mixing.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetGroupPostMixCallback)
func (group *Group) SetPostMixCallback(cb GroupMixCallback) bool {
	panic("not implemented")
	return iSetGroupPostMixCallback(group, cb, 0)
}

// AudioDecoder

// MIX_DestroyAudioDecoder - Destroy the specified audio decoder.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudioDecoder)
func (audiodecoder *AudioDecoder) Destroy() {
	iDestroyAudioDecoder(audiodecoder)
}

// MIX_GetAudioDecoderProperties - Get the properties associated with a [MIX_AudioDecoder](MIX_AudioDecoder).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderProperties)
func (audiodecoder *AudioDecoder) Properties() (sdl.PropertiesID, error) {
	props := iGetAudioDecoderProperties(audiodecoder)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// MIX_GetAudioDecoderFormat - Query the initial audio format of a [MIX_AudioDecoder](MIX_AudioDecoder).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDecoderFormat)
func (audiodecoder *AudioDecoder) Format(spec *sdl.AudioSpec) error {
	if !iGetAudioDecoderFormat(audiodecoder, spec) {
		return internal.LastErr()
	}

	return nil
}

// MIX_DecodeAudio - Decode more audio from a [MIX_AudioDecoder](MIX_AudioDecoder).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DecodeAudio)
func (audiodecoder *AudioDecoder) DecodeAudio(buffer []byte, spec *sdl.AudioSpec) (int32, error) {
	count := iDecodeAudio(audiodecoder, uintptr(unsafe.Pointer(unsafe.SliceData(buffer))), int32(len(buffer)), spec)
	if count == -1 {
		return -1, internal.LastErr()
	}
	runtime.KeepAlive(buffer)

	return count, nil
}

// Mixer

// MIX_DestroyMixer - Free a mixer.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyMixer)
func (mixer *Mixer) Destroy() {
	iDestroyMixer(mixer)
}

// MIX_GetMixerProperties - Get the properties associated with a mixer.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerProperties)
func (mixer *Mixer) Properties() sdl.PropertiesID {
	return iGetMixerProperties(mixer)
}

// MIX_GetMixerFormat - Get the audio format a mixer is generating.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetMixerFormat)
func (mixer *Mixer) Format(spec *sdl.AudioSpec) error {
	if !iGetMixerFormat(mixer, spec) {
		return internal.LastErr()
	}

	return nil
}

// MIX_LoadAudio_IO - Load audio for playback from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio_IO)
func (mixer *Mixer) LoadAudio_IO(stream *sdl.IOStream, predecode, closeio bool) (*Audio, error) {
	audio := iLoadAudio_IO(mixer, stream, predecode, closeio)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

// MIX_LoadAudio - Load audio for playback from a file.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadAudio)
func (mixer *Mixer) LoadAudio(path string, predecode bool) (*Audio, error) {
	audio := iLoadAudio(mixer, path, predecode)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

// MIX_LoadRawAudio_IO - Load raw PCM data from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio_IO)
func (mixer *Mixer) LoadRawAudio_IO(stream *sdl.IOStream, spec *sdl.AudioSpec, closeio bool) (*Audio, error) {
	audio := iLoadRawAudio_IO(mixer, stream, spec, closeio)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

// MIX_LoadRawAudio - Load raw PCM data from a memory buffer.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudio)
func (mixer *Mixer) LoadRawAudio(data []byte, spec *sdl.AudioSpec) (*Audio, error) {
	audio := iLoadRawAudio(mixer, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uintptr(len(data)), spec)
	if audio == nil {
		return nil, internal.LastErr()
	}
	runtime.KeepAlive(data)

	return audio, nil
}

// MIX_LoadRawAudioNoCopy - Load raw PCM data from a memory buffer without making a copy.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_LoadRawAudioNoCopy)
func (mixer *Mixer) LoadRawAudioNoCopy(data []byte, spec *sdl.AudioSpec) (*Audio, error) {
	audio := iLoadRawAudioNoCopy(mixer, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uintptr(len(data)), spec, false)
	if audio == nil {
		return nil, internal.LastErr()
	}
	runtime.KeepAlive(data)

	return audio, nil
}

// MIX_CreateSineWaveAudio - Create a [MIX_Audio](MIX_Audio) that generates a sinewave.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateSineWaveAudio)
func (mixer *Mixer) CreateSineWaveAudio(hz int32, amplitude float32) (*Audio, error) {
	audio := iCreateSineWaveAudio(mixer, hz, amplitude)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

// MIX_CreateTrack - Create a new track on a mixer.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateTrack)
func (mixer *Mixer) CreateTrack() (*Track, error) {
	track := iCreateTrack(mixer)
	if track == nil {
		return nil, internal.LastErr()
	}

	return track, nil
}

// MIX_PlayTag - Start (or restart) mixing all tracks with a specific tag for playback.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTag)
func (mixer *Mixer) PlayTag(tag string, options sdl.PropertiesID) error {
	if !iPlayTag(mixer, tag, options) {
		return internal.LastErr()
	}

	return nil
}

// MIX_PlayAudio - Play a [MIX_Audio](MIX_Audio) from start to finish without any management.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PlayAudio)
func (mixer *Mixer) PlayAudio(audio *Audio) error {
	if !iPlayAudio(mixer, audio) {
		return internal.LastErr()
	}

	return nil
}

// MIX_StopAllTracks - Halt all currently-playing tracks, possibly fading out over time.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_StopAllTracks)
func (mixer *Mixer) StopAllTracks(fadeOutMS int64) error {
	if !iStopAllTracks(mixer, fadeOutMS) {
		return internal.LastErr()
	}

	return nil
}

// MIX_StopTag - Halt all tracks with a specific tag, possibly fading out over time.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_StopTag)
func (mixer *Mixer) StopTag(tag string, fadeOutMS int64) error {
	if !iStopTag(mixer, tag, fadeOutMS) {
		return internal.LastErr()
	}

	return nil
}

// MIX_PauseAllTracks - Pause all currently-playing tracks.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PauseAllTracks)
func (mixer *Mixer) PauseAllTracks() error {
	if !iPauseAllTracks(mixer) {
		return internal.LastErr()
	}

	return nil
}

// MIX_PauseTag - Pause all tracks with a specific tag.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTag)
func (mixer *Mixer) PauseTag(tag string) error {
	if !iPauseTag(mixer, tag) {
		return internal.LastErr()
	}

	return nil
}

// MIX_ResumeAllTracks - Resume all currently-paused tracks.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeAllTracks)
func (mixer *Mixer) ResumeAllTracks() error {
	if !iResumeAllTracks(mixer) {
		return internal.LastErr()
	}

	return nil
}

// MIX_ResumeTag - Resume all tracks with a specific tag.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTag)
func (mixer *Mixer) ResumeTag(tag string) error {
	if !iResumeTag(mixer, tag) {
		return internal.LastErr()
	}

	return nil
}

// MIX_SetMasterGain - Set a mixer's master gain control.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetMasterGain)
func (mixer *Mixer) SetMasterGain(gain float32) error {
	if !iSetMasterGain(mixer, gain) {
		return internal.LastErr()
	}

	return nil
}

// MIX_GetMasterGain - Get a mixer's master gain control.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetMasterGain)
func (mixer *Mixer) MasterGain() float32 {
	return iGetMasterGain(mixer)
}

// MIX_SetTagGain - Set the gain control of all tracks with a specific tag.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTagGain)
func (mixer *Mixer) SetTagGain(tag string, gain float32) error {
	if !iSetTagGain(mixer, tag, gain) {
		return internal.LastErr()
	}

	return nil
}

// MIX_CreateGroup - Create a mixing group.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_CreateGroup)
func (mixer *Mixer) CreateGroup() (*Group, error) {
	group := iCreateGroup(mixer)
	if group == nil {
		return nil, internal.LastErr()
	}

	return group, nil
}

// MIX_SetPostMixCallback - Set a callback that fires when all mixing has completed.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetPostMixCallback)
func (mixer *Mixer) SetPostMixCallback(cb PostMixCallback) bool {
	panic("not implemented")
	return iSetPostMixCallback(mixer, cb, 0)
}

// MIX_Generate - Generate mixer output when not driving an audio device.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_Generate)
func (mixer *Mixer) Generate(buffer []byte) error {
	if !iGenerate(mixer, uintptr(unsafe.Pointer(unsafe.SliceData(buffer))), int32(len(buffer))) {
		return internal.LastErr()
	}
	runtime.KeepAlive(buffer)

	return nil
}

// Audio

// MIX_GetAudioProperties - Get the properties associated with a [MIX_Audio](MIX_Audio).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioProperties)
func (audio *Audio) Properties() (sdl.PropertiesID, error) {
	props := iGetAudioProperties(audio)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// MIX_GetAudioDuration - Get the length of a [MIX_Audio](MIX_Audio)'s playback in sample frames.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioDuration)
func (audio *Audio) Duration() int64 {
	return iGetAudioDuration(audio)
}

// MIX_GetAudioFormat - Query the initial audio format of a [MIX_Audio](MIX_Audio).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetAudioFormat)
func (audio *Audio) Format(spec *sdl.AudioSpec) error {
	if !iGetAudioFormat(audio, spec) {
		return internal.LastErr()
	}

	return nil
}

// MIX_DestroyAudio - Destroy the specified audio.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyAudio)
func (audio *Audio) Destroy() {
	iDestroyAudio(audio)
}

// MIX_AudioMSToFrames - Convert milliseconds to sample frames for a [MIX_Audio](MIX_Audio)'s format.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_AudioMSToFrames)
func (audio *Audio) MSToFrames(ms uint64) uint64 {
	return iAudioMSToFrames(audio, ms)
}

// MIX_AudioFramesToMS - Convert sample frames for a [MIX_Audio](MIX_Audio)'s format to milliseconds.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_AudioFramesToMS)
func (audio *Audio) FramesToMS(frames uint64) uint64 {
	return iAudioFramesToMS(audio, frames)
}

// Track

// MIX_DestroyTrack - Destroy the specified track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_DestroyTrack)
func (track *Track) Destroy() {
	iDestroyTrack(track)
}

// MIX_GetTrackProperties - Get the properties associated with a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackProperties)
func (track *Track) Properties() (sdl.PropertiesID, error) {
	props := iGetTrackProperties(track)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// MIX_GetTrackMixer - Get the [MIX_Mixer](MIX_Mixer) that owns a [MIX_Track](MIX_Track).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackMixer)
func (track *Track) Mixer() (*Mixer, error) {
	mix := iGetTrackMixer(track)
	if mix == nil {
		return nil, internal.LastErr()
	}

	return mix, nil
}

// MIX_SetTrackAudio - Set a [MIX_Track](MIX_Track)'s input to a [MIX_Audio](MIX_Audio).
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudio)
func (track *Track) SetAudio(audio *Audio) error {
	if !iSetTrackAudio(track, audio) {
		return internal.LastErr()
	}

	return nil
}

// MIX_SetTrackAudioStream - Set a [MIX_Track](MIX_Track)'s input to an SDL_AudioStream.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackAudioStream)
func (track *Track) SetAudioStream(stream *sdl.AudioStream) error {
	if !iSetTrackAudioStream(track, stream) {
		return internal.LastErr()
	}

	return nil
}

// MIX_SetTrackIOStream - Set a [MIX_Track](MIX_Track)'s input to an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackIOStream)
func (track *Track) SetIOStream(io *sdl.IOStream, closeio bool) error {
	if !iSetTrackIOStream(track, io, closeio) {
		return internal.LastErr()
	}

	return nil
}

// MIX_TagTrack - Assign an arbitrary tag to a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TagTrack)
func (track *Track) SetTag(tag string) error {
	if !iTagTrack(track, tag) {
		return internal.LastErr()
	}

	return nil
}

// MIX_UntagTrack - Remove an arbitrary tag from a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_UntagTrack)
func (track *Track) RemoveTag(tag string) {
	iUntagTrack(track, tag)
}

// MIX_SetTrackPlaybackPosition - Seek a playing track to a new position in its input.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackPlaybackPosition)
func (track *Track) SetPlaybackPosition(frames uint64) error {
	if !iSetTrackPlaybackPosition(track, frames) {
		return internal.LastErr()
	}

	return nil
}

// MIX_GetTrackPlaybackPosition - Get the current input position of a playing track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackPlaybackPosition)
func (track *Track) PlaybackPosition() (int64, error) {
	pos := iGetTrackPlaybackPosition(track)
	if pos == -1 {
		return -1, internal.LastErr()
	}

	return pos, nil
}

// MIX_TrackLooping - Query whether a given track is looping.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackLooping)
func (track *Track) Looping() bool {
	return iTrackLooping(track)
}

// MIX_GetTrackAudio - Query the [MIX_Audio](MIX_Audio) assigned to a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudio)
func (track *Track) Audio() *Audio {
	return iGetTrackAudio(track)
}

// MIX_GetTrackAudioStream - Query the SDL_AudioStream assigned to a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackAudioStream)
func (track *Track) AudioStream() *sdl.AudioStream {
	return iGetTrackAudioStream(track)
}

// MIX_GetTrackRemaining - Return the number of sample frames remaining to be mixed in a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackRemaining)
func (track *Track) Remaining() int64 {
	return iGetTrackRemaining(track)
}

// MIX_TrackMSToFrames - Convert milliseconds to sample frames for a track's current format.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackMSToFrames)
func (track *Track) MSToFrames(ms uint64) uint64 {
	return iTrackMSToFrames(track, ms)
}

// MIX_TrackFramesToMS - Convert sample frames for a track's current format to milliseconds.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackFramesToMS)
func (track *Track) FramesToMS(frames uint64) uint64 {
	return iTrackFramesToMS(track, frames)
}

// MIX_PlayTrack - Start (or restart) mixing a track for playback.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PlayTrack)
func (track *Track) Play(options sdl.PropertiesID) error {
	if !iPlayTrack(track, options) {
		return internal.LastErr()
	}

	return nil
}

// MIX_StopTrack - Halt a currently-playing track, possibly fading out over time.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_StopTrack)
func (track *Track) Stop(fadeOutFrames int64) error {
	if !iStopTrack(track, fadeOutFrames) {
		return internal.LastErr()
	}

	return nil
}

// MIX_PauseTrack - Pause a currently-playing track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_PauseTrack)
func (track *Track) Pause() error {
	if !iPauseTrack(track) {
		return internal.LastErr()
	}

	return nil
}

// MIX_ResumeTrack - Resume a currently-paused track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_ResumeTrack)
func (track *Track) Resume() error {
	if !iResumeTrack(track) {
		return internal.LastErr()
	}

	return nil
}

// MIX_TrackPlaying - Query if a track is currently playing.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPlaying)
func (track *Track) Playing() bool {
	return iTrackPlaying(track)
}

// MIX_TrackPaused - Query if a track is currently paused.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_TrackPaused)
func (track *Track) Paused() bool {
	return iTrackPaused(track)
}

// MIX_SetTrackGain - Set a track's gain control.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGain)
func (track *Track) SetGain(gain float32) error {
	if !iSetTrackGain(track, gain) {
		return internal.LastErr()
	}

	return nil
}

// MIX_GetTrackGain - Get a track's gain control.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackGain)
func (track *Track) Gain() float32 {
	return iGetTrackGain(track)
}

// MIX_SetTrackFrequencyRatio - Change the frequency ratio of a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackFrequencyRatio)
func (track *Track) SetFrequencyRatio(ratio float32) error {
	if !iSetTrackFrequencyRatio(track, ratio) {
		return internal.LastErr()
	}

	return nil
}

// MIX_GetTrackFrequencyRatio - Query the frequency ratio of a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrackFrequencyRatio)
func (track *Track) FrequencyRatio() float32 {
	return iGetTrackFrequencyRatio(track)
}

// MIX_SetTrackOutputChannelMap - Set the current output channel map of a track.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackOutputChannelMap)
func (track *Track) SetOutputChannelMap(channelMap []int32) error {
	if !iSetTrackOutputChannelMap(track, unsafe.SliceData(channelMap), int32(len(channelMap))) {
		return internal.LastErr()
	}
	runtime.KeepAlive(channelMap)

	return nil
}

// MIX_SetTrackStereo - Force a track to stereo output, with optionally left/right panning.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStereo)
func (track *Track) SetStereo(gains []StereoGains) error {
	if !iSetTrackStereo(track, unsafe.SliceData(gains)) {
		return internal.LastErr()
	}
	runtime.KeepAlive(gains)

	return nil
}

// MIX_SetTrack3DPosition - Set a track's position in 3D space.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrack3DPosition)
func (track *Track) Set3DPosition(position *Point3D) error {
	if !iSetTrack3DPosition(track, position) {
		return internal.LastErr()
	}

	return nil
}

// MIX_GetTrack3DPosition - Get a track's current position in 3D space.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_GetTrack3DPosition)
func (track *Track) Get3DPosition() (Point3D, error) {
	var position Point3D

	if !iGetTrack3DPosition(track, &position) {
		return Point3D{}, internal.LastErr()
	}

	return position, nil
}

// MIX_SetTrackGroup - Assign a track to a mixing group.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackGroup)
func (track *Track) SetGroup(group *Group) error {
	if !iSetTrackGroup(track, group) {
		return internal.LastErr()
	}

	return nil
}

// MIX_SetTrackStoppedCallback - Set a callback that fires when a [MIX_Track](MIX_Track) is stopped.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackStoppedCallback)
func (track *Track) SetStoppedCallback(cb TrackStoppedCallback) bool {
	panic("not implemented")
	return iSetTrackStoppedCallback(track, cb, 0)
}

// MIX_SetTrackRawCallback - Set a callback that fires when a [MIX_Track](MIX_Track) has initial decoded audio.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackRawCallback)
func (track *Track) SetRawCallback(cb TrackMixCallback) bool {
	panic("not implemented")
	return iSetTrackRawCallback(track, cb, 0)
}

// MIX_SetTrackCookedCallback - Set a callback that fires when the mixer has transformed a track's audio.
// (https://wiki.libsdl.org/SDL3_mixer/MIX_SetTrackCookedCallback)
func (track *Track) SetCookedCallback(cb TrackMixCallback) bool {
	panic("not implemented")
	return iSetTrackCookedCallback(track, cb, 0)
}
