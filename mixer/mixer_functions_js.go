//go:build js

package mixer

import (
	js "syscall/js"
	"unsafe"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

func initialize() {
	iVersion = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_MIX_Version",
		)
		return int32(ret.Int())
	}

	iInit = func() bool {
		ret := js.Global().Get("Module").Call(
			"_MIX_Init",
		)
		return internal.GetBool(ret)
	}

	iQuit = func() {
		js.Global().Get("Module").Call(
			"_MIX_Quit",
		)
	}

	iGetNumAudioDecoders = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_MIX_GetNumAudioDecoders",
		)
		return int32(ret.Int())
	}

	iGetAudioDecoder = func(index int32) string {
		_index := index
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoder",
			_index,
		)
		return internal.UTF8JSToString(ret)
	}

	iCreateMixerDevice = func(devid sdl.AudioDeviceID, spec *sdl.AudioSpec) *Mixer {
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateMixerDevice",
			_devid,
			_spec,
		)
		return internal.NewObject[Mixer](ret)
	}

	iCreateMixer = func(spec *sdl.AudioSpec) *Mixer {
		internal.StackSave()
		defer internal.StackRestore()
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateMixer",
			_spec,
		)
		return internal.NewObject[Mixer](ret)
	}

	iDestroyMixer = func(mixer *Mixer) {
		_mixer, _ := internal.GetJSPointer(mixer)
		js.Global().Get("Module").Call(
			"_MIX_DestroyMixer",
			_mixer,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(mixer)))
	}

	iGetMixerProperties = func(mixer *Mixer) sdl.PropertiesID {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerProperties",
			_mixer,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetMixerFormat = func(mixer *Mixer, spec *sdl.AudioSpec) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_spec := internal.StackAlloc(internal.CSizeof[sdl.AudioSpec]())
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerFormat",
			_mixer,
			_spec,
		)
		internal.CopyJSToObject(spec, _spec)
		return internal.GetBool(ret)
	}

	iLoadAudio_IO = func(mixer *Mixer, io *sdl.IOStream, predecode bool, closeio bool) *Audio {
		_mixer, _ := internal.GetJSPointer(mixer)
		_io, _ := internal.GetJSPointer(io)
		_predecode := internal.NewBoolean(predecode)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudio_IO",
			_mixer,
			_io,
			_predecode,
			_closeio,
		)
		return internal.NewObject[Audio](ret)
	}

	iLoadAudio = func(mixer *Mixer, path string, predecode bool) *Audio {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_path := internal.StringOnJSStack(path)
		_predecode := internal.NewBoolean(predecode)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudio",
			_mixer,
			_path,
			_predecode,
		)
		return internal.NewObject[Audio](ret)
	}

	iLoadAudioWithProperties = func(props sdl.PropertiesID) *Audio {
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudioWithProperties",
			_props,
		)
		return internal.NewObject[Audio](ret)
	}

	iLoadRawAudio_IO = func(mixer *Mixer, io *sdl.IOStream, spec *sdl.AudioSpec, closeio bool) *Audio {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_io, _ := internal.GetJSPointer(io)
		_spec := internal.CloneObjectToJSStack(spec)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudio_IO",
			_mixer,
			_io,
			_spec,
			_closeio,
		)
		return internal.NewObject[Audio](ret)
	}

	iLoadRawAudio = func(mixer *Mixer, data uintptr, datalen uintptr, spec *sdl.AudioSpec) *Audio {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_data, free := internal.ClonePtrArrayToJSHeap(*(**byte)(unsafe.Pointer(&data)), int(datalen))
		defer free()
		_datalen := int32(datalen)
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudio",
			_mixer,
			_data,
			_datalen,
			_spec,
		)
		return internal.NewObject[Audio](ret)
	}

	iLoadRawAudioNoCopy = func(mixer *Mixer, data uintptr, datalen uintptr, spec *sdl.AudioSpec, free_when_done bool) *Audio {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		// C cannot alias Go memory: the audio keeps a heap copy, which SDL
		// owns and frees with the audio.
		_data := internal.CloneByteSliceToJSHeap(unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), datalen))
		_datalen := int32(datalen)
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudioNoCopy",
			_mixer,
			_data,
			_datalen,
			_spec,
			internal.NewBoolean(true),
		)
		if ret.Int() == 0 {
			js.Global().Call("_free", _data)
		}
		return internal.NewObject[Audio](ret)
	}

	iCreateSineWaveAudio = func(mixer *Mixer, hz int32, amplitude float32, ms int64) *Audio {
		_mixer, _ := internal.GetJSPointer(mixer)
		_hz := hz
		_amplitude := amplitude
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateSineWaveAudio",
			_mixer,
			_hz,
			_amplitude,
			_ms,
		)
		return internal.NewObject[Audio](ret)
	}

	iGetAudioProperties = func(audio *Audio) sdl.PropertiesID {
		_audio, _ := internal.GetJSPointer(audio)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioProperties",
			_audio,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetAudioDuration = func(audio *Audio) int64 {
		_audio, _ := internal.GetJSPointer(audio)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDuration",
			_audio,
		)
		return internal.GetInt64(ret)
	}

	iGetAudioFormat = func(audio *Audio, spec *sdl.AudioSpec) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_audio, _ := internal.GetJSPointer(audio)
		_spec := internal.StackAlloc(internal.CSizeof[sdl.AudioSpec]())
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioFormat",
			_audio,
			_spec,
		)
		internal.CopyJSToObject(spec, _spec)
		return internal.GetBool(ret)
	}

	iDestroyAudio = func(audio *Audio) {
		_audio, _ := internal.GetJSPointer(audio)
		js.Global().Get("Module").Call(
			"_MIX_DestroyAudio",
			_audio,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(audio)))
	}

	iCreateTrack = func(mixer *Mixer) *Track {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateTrack",
			_mixer,
		)
		return internal.NewObject[Track](ret)
	}

	iDestroyTrack = func(track *Track) {
		_track, _ := internal.GetJSPointer(track)
		js.Global().Get("Module").Call(
			"_MIX_DestroyTrack",
			_track,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(track)))
	}

	iGetTrackProperties = func(track *Track) sdl.PropertiesID {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackProperties",
			_track,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetTrackMixer = func(track *Track) *Mixer {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackMixer",
			_track,
		)
		return internal.NewObject[Mixer](ret)
	}

	iSetTrackAudio = func(track *Track, audio *Audio) bool {
		_track, _ := internal.GetJSPointer(track)
		_audio, _ := internal.GetJSPointer(audio)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackAudio",
			_track,
			_audio,
		)
		return internal.GetBool(ret)
	}

	iSetTrackAudioStream = func(track *Track, stream *sdl.AudioStream) bool {
		_track, _ := internal.GetJSPointer(track)
		_stream, _ := internal.GetJSPointer(stream)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackAudioStream",
			_track,
			_stream,
		)
		return internal.GetBool(ret)
	}

	iSetTrackIOStream = func(track *Track, io *sdl.IOStream, closeio bool) bool {
		_track, _ := internal.GetJSPointer(track)
		_io, _ := internal.GetJSPointer(io)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackIOStream",
			_track,
			_io,
			_closeio,
		)
		return internal.GetBool(ret)
	}

	iTagTrack = func(track *Track, tag string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_TagTrack",
			_track,
			_tag,
		)
		return internal.GetBool(ret)
	}

	iUntagTrack = func(track *Track, tag string) {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_tag := internal.StringOnJSStack(tag)
		js.Global().Get("Module").Call(
			"_MIX_UntagTrack",
			_track,
			_tag,
		)
	}

	iSetTrackPlaybackPosition = func(track *Track, frames int64) bool {
		_track, _ := internal.GetJSPointer(track)
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackPlaybackPosition",
			_track,
			_frames,
		)
		return internal.GetBool(ret)
	}

	iGetTrackPlaybackPosition = func(track *Track) int64 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackPlaybackPosition",
			_track,
		)
		return internal.GetInt64(ret)
	}

	iGetTrackAudio = func(track *Track) *Audio {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackAudio",
			_track,
		)
		return internal.NewObject[Audio](ret)
	}

	iGetTrackAudioStream = func(track *Track) *sdl.AudioStream {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackAudioStream",
			_track,
		)
		return internal.NewObject[sdl.AudioStream](ret)
	}

	iGetTrackRemaining = func(track *Track) int64 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackRemaining",
			_track,
		)
		return internal.GetInt64(ret)
	}

	iTrackMSToFrames = func(track *Track, ms int64) int64 {
		_track, _ := internal.GetJSPointer(track)
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackMSToFrames",
			_track,
			_ms,
		)
		return internal.GetInt64(ret)
	}

	iTrackFramesToMS = func(track *Track, frames int64) int64 {
		_track, _ := internal.GetJSPointer(track)
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackFramesToMS",
			_track,
			_frames,
		)
		return internal.GetInt64(ret)
	}

	iAudioMSToFrames = func(audio *Audio, ms int64) int64 {
		_audio, _ := internal.GetJSPointer(audio)
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_AudioMSToFrames",
			_audio,
			_ms,
		)
		return internal.GetInt64(ret)
	}

	iAudioFramesToMS = func(audio *Audio, frames int64) int64 {
		_audio, _ := internal.GetJSPointer(audio)
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_AudioFramesToMS",
			_audio,
			_frames,
		)
		return internal.GetInt64(ret)
	}

	iMSToFrames = func(sample_rate int32, ms int64) int64 {
		_sample_rate := sample_rate
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_MSToFrames",
			_sample_rate,
			_ms,
		)
		return internal.GetInt64(ret)
	}

	iFramesToMS = func(sample_rate int32, frames int64) int64 {
		_sample_rate := sample_rate
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_FramesToMS",
			_sample_rate,
			_frames,
		)
		return internal.GetInt64(ret)
	}

	iPlayTrack = func(track *Track, options sdl.PropertiesID) bool {
		_track, _ := internal.GetJSPointer(track)
		_options := int32(options)
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayTrack",
			_track,
			_options,
		)
		return internal.GetBool(ret)
	}

	iPlayTag = func(mixer *Mixer, tag string, options sdl.PropertiesID) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		_options := int32(options)
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayTag",
			_mixer,
			_tag,
			_options,
		)
		return internal.GetBool(ret)
	}

	iPlayAudio = func(mixer *Mixer, audio *Audio) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		_audio, _ := internal.GetJSPointer(audio)
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayAudio",
			_mixer,
			_audio,
		)
		return internal.GetBool(ret)
	}

	iStopTrack = func(track *Track, fade_out_frames int64) bool {
		_track, _ := internal.GetJSPointer(track)
		_fade_out_frames := internal.NewBigInt(fade_out_frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_StopTrack",
			_track,
			_fade_out_frames,
		)
		return internal.GetBool(ret)
	}

	iStopAllTracks = func(mixer *Mixer, fade_out_ms int64) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		_fade_out_ms := internal.NewBigInt(fade_out_ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_StopAllTracks",
			_mixer,
			_fade_out_ms,
		)
		return internal.GetBool(ret)
	}

	iStopTag = func(mixer *Mixer, tag string, fade_out_ms int64) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		_fade_out_ms := internal.NewBigInt(fade_out_ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_StopTag",
			_mixer,
			_tag,
			_fade_out_ms,
		)
		return internal.GetBool(ret)
	}

	iPauseTrack = func(track *Track) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseTrack",
			_track,
		)
		return internal.GetBool(ret)
	}

	iPauseAllTracks = func(mixer *Mixer) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseAllTracks",
			_mixer,
		)
		return internal.GetBool(ret)
	}

	iPauseTag = func(mixer *Mixer, tag string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseTag",
			_mixer,
			_tag,
		)
		return internal.GetBool(ret)
	}

	iResumeTrack = func(track *Track) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeTrack",
			_track,
		)
		return internal.GetBool(ret)
	}

	iResumeAllTracks = func(mixer *Mixer) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeAllTracks",
			_mixer,
		)
		return internal.GetBool(ret)
	}

	iResumeTag = func(mixer *Mixer, tag string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeTag",
			_mixer,
			_tag,
		)
		return internal.GetBool(ret)
	}

	iTrackPlaying = func(track *Track) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackPlaying",
			_track,
		)
		return internal.GetBool(ret)
	}

	iTrackPaused = func(track *Track) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackPaused",
			_track,
		)
		return internal.GetBool(ret)
	}

	iSetTrackGain = func(track *Track, gain float32) bool {
		_track, _ := internal.GetJSPointer(track)
		_gain := gain
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackGain",
			_track,
			_gain,
		)
		return internal.GetBool(ret)
	}

	iGetTrackGain = func(track *Track) float32 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackGain",
			_track,
		)
		return float32(ret.Float())
	}

	iSetTagGain = func(mixer *Mixer, tag string, gain float32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		_gain := gain
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTagGain",
			_mixer,
			_tag,
			_gain,
		)
		return internal.GetBool(ret)
	}

	iSetTrackFrequencyRatio = func(track *Track, ratio float32) bool {
		_track, _ := internal.GetJSPointer(track)
		_ratio := ratio
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackFrequencyRatio",
			_track,
			_ratio,
		)
		return internal.GetBool(ret)
	}

	iGetTrackFrequencyRatio = func(track *Track) float32 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackFrequencyRatio",
			_track,
		)
		return float32(ret.Float())
	}

	iSetTrackOutputChannelMap = func(track *Track, chmap *int32, count int32) bool {
		_track, _ := internal.GetJSPointer(track)
		_chmap, free := internal.ClonePtrArrayToJSHeap(chmap, int(count))
		defer free()
		_count := count
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackOutputChannelMap",
			_track,
			_chmap,
			_count,
		)
		return internal.GetBool(ret)
	}

	iSetTrackStereo = func(track *Track, gains *StereoGains) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_gains := internal.CloneObjectToJSStack(gains)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackStereo",
			_track,
			_gains,
		)
		return internal.GetBool(ret)
	}

	iSetTrack3DPosition = func(track *Track, position *Point3D) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_position := internal.CloneObjectToJSStack(position)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrack3DPosition",
			_track,
			_position,
		)
		return internal.GetBool(ret)
	}

	iGetTrack3DPosition = func(track *Track, position *Point3D) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_position := internal.StackAlloc(internal.CSizeof[Point3D]())
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrack3DPosition",
			_track,
			_position,
		)
		internal.CopyJSToObject(position, _position)
		return internal.GetBool(ret)
	}

	iCreateGroup = func(mixer *Mixer) *Group {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateGroup",
			_mixer,
		)
		return internal.NewObject[Group](ret)
	}

	iDestroyGroup = func(group *Group) {
		_group, _ := internal.GetJSPointer(group)
		js.Global().Get("Module").Call(
			"_MIX_DestroyGroup",
			_group,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(group)))
	}

	iGetGroupProperties = func(group *Group) sdl.PropertiesID {
		_group, _ := internal.GetJSPointer(group)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetGroupProperties",
			_group,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetGroupMixer = func(group *Group) *Mixer {
		_group, _ := internal.GetJSPointer(group)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetGroupMixer",
			_group,
		)
		return internal.NewObject[Mixer](ret)
	}

	iSetTrackGroup = func(track *Track, group *Group) bool {
		_track, _ := internal.GetJSPointer(track)
		_group, _ := internal.GetJSPointer(group)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackGroup",
			_track,
			_group,
		)
		return internal.GetBool(ret)
	}

	iSetTrackStoppedCallback = func(track *Track, cb TrackStoppedCallback, userdata uintptr) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackStoppedCallback",
			_track,
			int32(cb),
			int32(userdata),
		)
		return internal.GetBool(ret)
	}

	iSetTrackRawCallback = func(track *Track, cb TrackMixCallback, userdata uintptr) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackRawCallback",
			_track,
			int32(cb),
			int32(userdata),
		)
		return internal.GetBool(ret)
	}

	iSetTrackCookedCallback = func(track *Track, cb TrackMixCallback, userdata uintptr) bool {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackCookedCallback",
			_track,
			int32(cb),
			int32(userdata),
		)
		return internal.GetBool(ret)
	}

	iSetGroupPostMixCallback = func(group *Group, cb GroupMixCallback, userdata uintptr) bool {
		_group, _ := internal.GetJSPointer(group)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetGroupPostMixCallback",
			_group,
			int32(cb),
			int32(userdata),
		)
		return internal.GetBool(ret)
	}

	iSetPostMixCallback = func(mixer *Mixer, cb PostMixCallback, userdata uintptr) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetPostMixCallback",
			_mixer,
			int32(cb),
			int32(userdata),
		)
		return internal.GetBool(ret)
	}

	iCreateAudioDecoder = func(path string, props sdl.PropertiesID) *AudioDecoder {
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateAudioDecoder",
			_path,
			_props,
		)
		return internal.NewObject[AudioDecoder](ret)
	}

	iCreateAudioDecoder_IO = func(io *sdl.IOStream, closeio bool, props sdl.PropertiesID) *AudioDecoder {
		_io, _ := internal.GetJSPointer(io)
		_closeio := internal.NewBoolean(closeio)
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateAudioDecoder_IO",
			_io,
			_closeio,
			_props,
		)
		return internal.NewObject[AudioDecoder](ret)
	}

	iDestroyAudioDecoder = func(audiodecoder *AudioDecoder) {
		_audiodecoder, _ := internal.GetJSPointer(audiodecoder)
		js.Global().Get("Module").Call(
			"_MIX_DestroyAudioDecoder",
			_audiodecoder,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(audiodecoder)))
	}

	iGetAudioDecoderProperties = func(audiodecoder *AudioDecoder) sdl.PropertiesID {
		_audiodecoder, _ := internal.GetJSPointer(audiodecoder)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoderProperties",
			_audiodecoder,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetAudioDecoderFormat = func(audiodecoder *AudioDecoder, spec *sdl.AudioSpec) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, _ := internal.GetJSPointer(audiodecoder)
		_spec := internal.StackAlloc(internal.CSizeof[sdl.AudioSpec]())
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoderFormat",
			_audiodecoder,
			_spec,
		)
		internal.CopyJSToObject(spec, _spec)
		return internal.GetBool(ret)
	}

	iDecodeAudio = func(audiodecoder *AudioDecoder, buffer uintptr, buflen int32, spec *sdl.AudioSpec) int32 {
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, _ := internal.GetJSPointer(audiodecoder)
		dst := unsafe.Slice(*(**byte)(unsafe.Pointer(&buffer)), buflen)
		_buffer, free := internal.ClonePtrArrayToJSHeap(unsafe.SliceData(dst), int(buflen))
		defer free()
		_spec := internal.CloneObjectToJSStack(spec)
		ret := js.Global().Get("Module").Call(
			"_MIX_DecodeAudio",
			_audiodecoder,
			_buffer,
			buflen,
			_spec,
		)
		if n := ret.Int(); n > 0 {
			copy(dst, internal.GetByteSliceFromJSPtr(_buffer, n))
		}
		return int32(ret.Int())
	}

	iLockMixer = func(mixer *Mixer) {
		_mixer, _ := internal.GetJSPointer(mixer)
		js.Global().Get("Module").Call(
			"_MIX_LockMixer",
			_mixer,
		)
	}

	iUnlockMixer = func(mixer *Mixer) {
		_mixer, _ := internal.GetJSPointer(mixer)
		js.Global().Get("Module").Call(
			"_MIX_UnlockMixer",
			_mixer,
		)
	}

	iLoadAudioNoCopy = func(mixer *Mixer, data uintptr, datalen uintptr, free_when_done bool) *Audio {
		_mixer, _ := internal.GetJSPointer(mixer)
		// C cannot alias Go memory: the audio keeps a heap copy, which SDL
		// owns and frees with the audio.
		_data := internal.CloneByteSliceToJSHeap(unsafe.Slice(*(**byte)(unsafe.Pointer(&data)), datalen))
		_datalen := int32(datalen)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudioNoCopy",
			_mixer,
			_data,
			_datalen,
			internal.NewBoolean(true),
		)
		if ret.Int() == 0 {
			js.Global().Call("_free", _data)
		}
		return internal.NewObject[Audio](ret)
	}

	iSetTrackRawIOStream = func(track *Track, io *sdl.IOStream, spec *sdl.AudioSpec, closeio bool) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_io, _ := internal.GetJSPointer(io)
		_spec := internal.CloneObjectToJSStack(spec)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackRawIOStream",
			_track,
			_io,
			_spec,
			_closeio,
		)
		return internal.GetBool(ret)
	}

	iGetTrackTags = func(track *Track, count *int32) uintptr {
		internal.StackSave()
		defer internal.StackRestore()
		_track, _ := internal.GetJSPointer(track)
		_count := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackTags",
			_track,
			_count,
		)
		*count = int32(internal.GetValue(_count, "i32").Int())
		return internal.NewUintptr(ret)
	}

	iGetTaggedTracks = func(mixer *Mixer, tag string, count *int32) uintptr {
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, _ := internal.GetJSPointer(mixer)
		_tag := internal.StringOnJSStack(tag)
		_count := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTaggedTracks",
			_mixer,
			_tag,
			_count,
		)
		*count = int32(internal.GetValue(_count, "i32").Int())
		return internal.NewUintptr(ret)
	}

	iGetTrackFadeFrames = func(track *Track) int64 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackFadeFrames",
			_track,
		)
		return internal.GetInt64(ret)
	}

	iGetTrackLoops = func(track *Track) int32 {
		_track, _ := internal.GetJSPointer(track)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackLoops",
			_track,
		)
		return int32(ret.Int())
	}

	iSetTrackLoops = func(track *Track, num_loops int32) bool {
		_track, _ := internal.GetJSPointer(track)
		_num_loops := num_loops
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackLoops",
			_track,
			_num_loops,
		)
		return internal.GetBool(ret)
	}

	iSetMixerGain = func(mixer *Mixer, gain float32) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		_gain := gain
		ret := js.Global().Get("Module").Call(
			"_MIX_SetMixerGain",
			_mixer,
			_gain,
		)
		return internal.GetBool(ret)
	}

	iGetMixerGain = func(mixer *Mixer) float32 {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerGain",
			_mixer,
		)
		return float32(ret.Float())
	}

	iSetMixerFrequencyRatio = func(mixer *Mixer, ratio float32) bool {
		_mixer, _ := internal.GetJSPointer(mixer)
		_ratio := ratio
		ret := js.Global().Get("Module").Call(
			"_MIX_SetMixerFrequencyRatio",
			_mixer,
			_ratio,
		)
		return internal.GetBool(ret)
	}

	iGetMixerFrequencyRatio = func(mixer *Mixer) float32 {
		_mixer, _ := internal.GetJSPointer(mixer)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerFrequencyRatio",
			_mixer,
		)
		return float32(ret.Float())
	}

	iGenerate = func(mixer *Mixer, buffer uintptr, buflen int32) int32 {
		_mixer, _ := internal.GetJSPointer(mixer)
		dst := unsafe.Slice(*(**byte)(unsafe.Pointer(&buffer)), buflen)
		_buffer, free := internal.ClonePtrArrayToJSHeap(unsafe.SliceData(dst), int(buflen))
		defer free()
		ret := js.Global().Get("Module").Call(
			"_MIX_Generate",
			_mixer,
			_buffer,
			buflen,
		)
		if n := ret.Int(); n > 0 {
			copy(dst, internal.GetByteSliceFromJSPtr(_buffer, n))
		}
		return int32(ret.Int())
	}
}
