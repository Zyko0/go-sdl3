package mixer

import (
	"runtime"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Group

func (group *Group) Destroy() {
	iDestroyGroup(group)
}

func (group *Group) Properties() (sdl.PropertiesID, error) {
	props := iGetGroupProperties(group)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

func (group *Group) Mixer() *Mixer {
	return iGetGroupMixer(group)
}

func (group *Group) SetPostMixCallback(cb GroupMixCallback, userdata uintptr) bool {
	panic("not implemented")
	return iSetGroupPostMixCallback(group, cb, userdata)
}

// AudioDecoder

func (audiodecoder *AudioDecoder) Destroy() {
	iDestroyAudioDecoder(audiodecoder)
}

func (audiodecoder *AudioDecoder) Properties() (sdl.PropertiesID, error) {
	props := iGetAudioDecoderProperties(audiodecoder)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

func (audiodecoder *AudioDecoder) Format(spec *sdl.AudioSpec) error {
	if !iGetAudioDecoderFormat(audiodecoder, spec) {
		return internal.LastErr()
	}

	return nil
}

func (audiodecoder *AudioDecoder) DecodeAudio(buffer uintptr, buflen int32, spec *sdl.AudioSpec) (int32, error) {
	count := iDecodeAudio(audiodecoder, buffer, buflen, spec)
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// Mixer

func (mixer *Mixer) Destroy() {
	iDestroyMixer(mixer)
}

func (mixer *Mixer) Properties() sdl.PropertiesID {
	return iGetMixerProperties(mixer)
}

func (mixer *Mixer) Format(spec *sdl.AudioSpec) error {
	if !iGetMixerFormat(mixer, spec) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) LoadAudio_IO(io *sdl.IOStream, predecode bool, closeio bool) (*Audio, error) {
	audio := iLoadAudio_IO(mixer, io, predecode, closeio)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) LoadAudio(path string, predecode bool) (*Audio, error) {
	audio := iLoadAudio(mixer, path, predecode)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) LoadRawAudio_IO(io *sdl.IOStream, spec *sdl.AudioSpec, closeio bool) (*Audio, error) {
	audio := iLoadRawAudio_IO(mixer, io, spec, closeio)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) LoadRawAudio(data uintptr, datalen uintptr, spec *sdl.AudioSpec) (*Audio, error) {
	audio := iLoadRawAudio(mixer, data, datalen, spec)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) LoadRawAudioNoCopy(data uintptr, datalen uintptr, spec *sdl.AudioSpec, free_when_done bool) (*Audio, error) {
	audio := iLoadRawAudioNoCopy(mixer, data, datalen, spec, free_when_done)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) CreateSineWaveAudio(hz int32, amplitude float32) (*Audio, error) {
	audio := iCreateSineWaveAudio(mixer, hz, amplitude)
	if audio == nil {
		return nil, internal.LastErr()
	}

	return audio, nil
}

func (mixer *Mixer) CreateTrack() (*Track, error) {
	track := iCreateTrack(mixer)
	if track == nil {
		return nil, internal.LastErr()
	}

	return track, nil
}

func (mixer *Mixer) PlayTag(tag string, options sdl.PropertiesID) error {
	if !iPlayTag(mixer, tag, options) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) PlayAudio(audio *Audio) error {
	if !iPlayAudio(mixer, audio) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) StopAllTracks(fade_out_ms int64) error {
	if !iStopAllTracks(mixer, fade_out_ms) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) StopTag(tag string, fade_out_ms int64) error {
	if !iStopTag(mixer, tag, fade_out_ms) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) PauseAllTracks() error {
	if !iPauseAllTracks(mixer) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) PauseTag(tag string) error {
	if !iPauseTag(mixer, tag) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) ResumeAllTracks() error {
	if !iResumeAllTracks(mixer) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) ResumeTag(tag string) error {
	if !iResumeTag(mixer, tag) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) SetMasterGain(gain float32) error {
	if !iSetMasterGain(mixer, gain) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) MasterGain() float32 {
	return iGetMasterGain(mixer)
}

func (mixer *Mixer) SetTagGain(tag string, gain float32) error {
	if !iSetTagGain(mixer, tag, gain) {
		return internal.LastErr()
	}

	return nil
}

func (mixer *Mixer) CreateGroup() (*Group, error) {
	group := iCreateGroup(mixer)
	if group == nil {
		return nil, internal.LastErr()
	}

	return group, nil
}

func (mixer *Mixer) SetPostMixCallback(cb PostMixCallback, userdata uintptr) bool {
	panic("not implemented")
	return iSetPostMixCallback(mixer, cb, userdata)
}

func (mixer *Mixer) Generate(buffer []byte) error {
	if !iGenerate(mixer, uintptr(unsafe.Pointer(&buffer[0])), int32(len(buffer))) {
		runtime.KeepAlive(buffer)
		return internal.LastErr()
	}

	return nil
}

// Audio

func (audio *Audio) Properties() (sdl.PropertiesID, error) {
	props := iGetAudioProperties(audio)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

func (audio *Audio) Duration() int64 {
	return iGetAudioDuration(audio)
}

func (audio *Audio) Format(spec *sdl.AudioSpec) error {
	if !iGetAudioFormat(audio, spec) {
		return internal.LastErr()
	}

	return nil
}

func (audio *Audio) Destroy() {
	iDestroyAudio(audio)
}

func (audio *Audio) MSToFrames(ms uint64) uint64 {
	return iAudioMSToFrames(audio, ms)
}

func (audio *Audio) FramesToMS(frames uint64) uint64 {
	return iAudioFramesToMS(audio, frames)
}

// Track

func (track *Track) Destroy() {
	iDestroyTrack(track)
}

func (track *Track) Properties() (sdl.PropertiesID, error) {
	props := iGetTrackProperties(track)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

func (track *Track) Mixer() (*Mixer, error) {
	mix := iGetTrackMixer(track)
	if mix == nil {
		return nil, internal.LastErr()
	}

	return mix, nil
}

func (track *Track) SetAudio(audio *Audio) error {
	if !iSetTrackAudio(track, audio) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) SetAudioStream(stream *sdl.AudioStream) error {
	if !iSetTrackAudioStream(track, stream) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) SetIOStream(io *sdl.IOStream, closeio bool) error {
	if !iSetTrackIOStream(track, io, closeio) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) SetTag(tag string) error {
	if !iTagTrack(track, tag) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) RemoveTag(tag string) {
	iUntagTrack(track, tag)
}

func (track *Track) SetPlaybackPosition(frames uint64) error {
	if !iSetTrackPlaybackPosition(track, frames) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) PlaybackPosition() (int64, error) {
	pos := iGetTrackPlaybackPosition(track)
	if pos == -1 {
		return -1, internal.LastErr()
	}

	return pos, nil
}

func (track *Track) Looping() bool {
	return iTrackLooping(track)
}

func (track *Track) Audio() *Audio {
	return iGetTrackAudio(track)
}

func (track *Track) AudioStream() *sdl.AudioStream {
	return iGetTrackAudioStream(track)
}

func (track *Track) Remaining() int64 {
	return iGetTrackRemaining(track)
}

func (track *Track) MSToFrames(ms uint64) uint64 {
	return iTrackMSToFrames(track, ms)
}

func (track *Track) FramesToMS(frames uint64) uint64 {
	return iTrackFramesToMS(track, frames)
}

func (track *Track) Play(options sdl.PropertiesID) error {
	if !iPlayTrack(track, options) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Stop(fade_out_frames int64) error {
	if !iStopTrack(track, fade_out_frames) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Pause() error {
	if !iPauseTrack(track) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Resume() error {
	if !iResumeTrack(track) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Playing() bool {
	return iTrackPlaying(track)
}

func (track *Track) Paused() bool {
	return iTrackPaused(track)
}

func (track *Track) SetGain(gain float32) error {
	if !iSetTrackGain(track, gain) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Gain() float32 {
	return iGetTrackGain(track)
}

func (track *Track) SetFrequencyRatio(ratio float32) error {
	if !iSetTrackFrequencyRatio(track, ratio) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) FrequencyRatio() float32 {
	return iGetTrackFrequencyRatio(track)
}

func (track *Track) SetOutputChannelMap(channelMap []int32) error {
	if !iSetTrackOutputChannelMap(track, &channelMap[0], int32(len(channelMap))) {
		return internal.LastErr()
	}
	runtime.KeepAlive(channelMap)

	return nil
}

func (track *Track) SetStereo(gains []StereoGains) error {
	if !iSetTrackStereo(track, &gains[0]) {
		return internal.LastErr()
	}
	runtime.KeepAlive(gains)

	return nil
}

func (track *Track) Set3DPosition(position *Point3D) error {
	if !iSetTrack3DPosition(track, position) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) Get3DPosition() (Point3D, error) {
	var position Point3D

	if !iGetTrack3DPosition(track, &position) {
		return Point3D{}, internal.LastErr()
	}

	return position, nil
}

func (track *Track) SetGroup(group *Group) error {
	if !iSetTrackGroup(track, group) {
		return internal.LastErr()
	}

	return nil
}

func (track *Track) SetStoppedCallback(cb TrackStoppedCallback, userdata uintptr) bool {
	panic("not implemented")
	return iSetTrackStoppedCallback(track, cb, userdata)
}

func (track *Track) SetRawCallback(cb TrackMixCallback, userdata uintptr) bool {
	panic("not implemented")
	return iSetTrackRawCallback(track, cb, userdata)
}

func (track *Track) SetCookedCallback(cb TrackMixCallback, userdata uintptr) bool {
	panic("not implemented")
	return iSetTrackCookedCallback(track, cb, userdata)
}
