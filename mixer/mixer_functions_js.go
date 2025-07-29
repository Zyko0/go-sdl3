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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_MIX_Version",
		)

		return int32(ret.Int())
	}

	iInit = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_MIX_Init",
		)

		return internal.GetBool(ret)
	}

	iQuit = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_MIX_Quit",
		)
	}

	iGetNumAudioDecoders = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_MIX_GetNumAudioDecoders",
		)

		return int32(ret.Int())
	}

	iGetAudioDecoder = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoder",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iCreateMixerDevice = func(devid sdl.AudioDeviceID, spec *sdl.AudioSpec) *Mixer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid := int32(devid)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateMixerDevice",
			_devid,
			_spec,
		)

		_obj := internal.NewObject[Mixer](ret)
		return _obj
	}

	iCreateMixer = func(spec *sdl.AudioSpec) *Mixer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateMixer",
			_spec,
		)

		_obj := internal.NewObject[Mixer](ret)
		return _obj
	}

	iDestroyMixer = func(mixer *Mixer) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		js.Global().Get("Module").Call(
			"_MIX_DestroyMixer",
			_mixer,
		)
	}

	iGetMixerProperties = func(mixer *Mixer) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerProperties",
			_mixer,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetMixerFormat = func(mixer *Mixer, spec *sdl.AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMixerFormat",
			_mixer,
			_spec,
		)

		return internal.GetBool(ret)
	}

	iLoadAudio_IO = func(mixer *Mixer, io *sdl.IOStream, predecode bool, closeio bool) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_io, ok := internal.GetJSPointer(io)
		if !ok {
			_io = internal.StackAlloc(int(unsafe.Sizeof(*io)))
		}
		_predecode := internal.NewBoolean(predecode)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudio_IO",
			_mixer,
			_io,
			_predecode,
			_closeio,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iLoadAudio = func(mixer *Mixer, path string, predecode bool) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_path := internal.StringOnJSStack(path)
		_predecode := internal.NewBoolean(predecode)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudio",
			_mixer,
			_path,
			_predecode,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iLoadAudioWithProperties = func(props sdl.PropertiesID) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props := uint32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadAudioWithProperties",
			_props,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iLoadRawAudio_IO = func(mixer *Mixer, io *sdl.IOStream, spec *sdl.AudioSpec, closeio bool) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_io, ok := internal.GetJSPointer(io)
		if !ok {
			_io = internal.StackAlloc(int(unsafe.Sizeof(*io)))
		}
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudio_IO",
			_mixer,
			_io,
			_spec,
			_closeio,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iLoadRawAudio = func(mixer *Mixer, data uintptr, datalen uintptr, spec *sdl.AudioSpec) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_data := internal.NewBigInt(data)
		_datalen := internal.NewBigInt(datalen)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudio",
			_mixer,
			_data,
			_datalen,
			_spec,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iLoadRawAudioNoCopy = func(mixer *Mixer, data uintptr, datalen uintptr, spec *sdl.AudioSpec, free_when_done bool) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_data := internal.NewBigInt(data)
		_datalen := internal.NewBigInt(datalen)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		_free_when_done := internal.NewBoolean(free_when_done)
		ret := js.Global().Get("Module").Call(
			"_MIX_LoadRawAudioNoCopy",
			_mixer,
			_data,
			_datalen,
			_spec,
			_free_when_done,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iCreateSineWaveAudio = func(mixer *Mixer, hz int32, amplitude float32) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_hz := int32(hz)
		_amplitude := int32(amplitude)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateSineWaveAudio",
			_mixer,
			_hz,
			_amplitude,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iGetAudioProperties = func(audio *Audio) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioProperties",
			_audio,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetAudioDuration = func(audio *Audio) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDuration",
			_audio,
		)

		return int64(internal.GetInt64(ret))
	}

	iGetAudioFormat = func(audio *Audio, spec *sdl.AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioFormat",
			_audio,
			_spec,
		)

		return internal.GetBool(ret)
	}

	iDestroyAudio = func(audio *Audio) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		js.Global().Get("Module").Call(
			"_MIX_DestroyAudio",
			_audio,
		)
	}

	iCreateTrack = func(mixer *Mixer) *Track {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateTrack",
			_mixer,
		)

		_obj := internal.NewObject[Track](ret)
		return _obj
	}

	iDestroyTrack = func(track *Track) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		js.Global().Get("Module").Call(
			"_MIX_DestroyTrack",
			_track,
		)
	}

	iGetTrackProperties = func(track *Track) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackProperties",
			_track,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetTrackMixer = func(track *Track) *Mixer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackMixer",
			_track,
		)

		_obj := internal.NewObject[Mixer](ret)
		return _obj
	}

	iSetTrackAudio = func(track *Track, audio *Audio) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackAudio",
			_track,
			_audio,
		)

		return internal.GetBool(ret)
	}

	iSetTrackAudioStream = func(track *Track, stream *sdl.AudioStream) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_stream, ok := internal.GetJSPointer(stream)
		if !ok {
			_stream = internal.StackAlloc(int(unsafe.Sizeof(*stream)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackAudioStream",
			_track,
			_stream,
		)

		return internal.GetBool(ret)
	}

	iSetTrackIOStream = func(track *Track, io *sdl.IOStream, closeio bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_io, ok := internal.GetJSPointer(io)
		if !ok {
			_io = internal.StackAlloc(int(unsafe.Sizeof(*io)))
		}
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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_TagTrack",
			_track,
			_tag,
		)

		return internal.GetBool(ret)
	}

	iUntagTrack = func(track *Track, tag string) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_tag := internal.StringOnJSStack(tag)
		js.Global().Get("Module").Call(
			"_MIX_UntagTrack",
			_track,
			_tag,
		)
	}

	iSetTrackPlaybackPosition = func(track *Track, frames uint64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackPlaybackPosition",
			_track,
			_frames,
		)

		return internal.GetBool(ret)
	}

	iGetTrackPlaybackPosition = func(track *Track) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackPlaybackPosition",
			_track,
		)

		return int64(internal.GetInt64(ret))
	}

	iTrackLooping = func(track *Track) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackLooping",
			_track,
		)

		return internal.GetBool(ret)
	}

	iGetTrackAudio = func(track *Track) *Audio {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackAudio",
			_track,
		)

		_obj := internal.NewObject[Audio](ret)
		return _obj
	}

	iGetTrackAudioStream = func(track *Track) *sdl.AudioStream {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackAudioStream",
			_track,
		)

		_obj := internal.NewObject[sdl.AudioStream](ret)
		return _obj
	}

	iGetTrackRemaining = func(track *Track) int64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackRemaining",
			_track,
		)

		return int64(internal.GetInt64(ret))
	}

	iTrackMSToFrames = func(track *Track, ms uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackMSToFrames",
			_track,
			_ms,
		)

		return uint64(internal.GetInt64(ret))
	}

	iTrackFramesToMS = func(track *Track, frames uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackFramesToMS",
			_track,
			_frames,
		)

		return uint64(internal.GetInt64(ret))
	}

	iAudioMSToFrames = func(audio *Audio, ms uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_AudioMSToFrames",
			_audio,
			_ms,
		)

		return uint64(internal.GetInt64(ret))
	}

	iAudioFramesToMS = func(audio *Audio, frames uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_AudioFramesToMS",
			_audio,
			_frames,
		)

		return uint64(internal.GetInt64(ret))
	}

	iMSToFrames = func(sample_rate int32, ms uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sample_rate := int32(sample_rate)
		_ms := internal.NewBigInt(ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_MSToFrames",
			_sample_rate,
			_ms,
		)

		return uint64(internal.GetInt64(ret))
	}

	iFramesToMS = func(sample_rate int32, frames uint64) uint64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_sample_rate := int32(sample_rate)
		_frames := internal.NewBigInt(frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_FramesToMS",
			_sample_rate,
			_frames,
		)

		return uint64(internal.GetInt64(ret))
	}

	iPlayTrack = func(track *Track, options sdl.PropertiesID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_options := uint32(options)
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayTrack",
			_track,
			_options,
		)

		return internal.GetBool(ret)
	}

	iPlayTag = func(mixer *Mixer, tag string, options sdl.PropertiesID) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_tag := internal.StringOnJSStack(tag)
		_options := uint32(options)
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayTag",
			_mixer,
			_tag,
			_options,
		)

		return internal.GetBool(ret)
	}

	iPlayAudio = func(mixer *Mixer, audio *Audio) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_audio, ok := internal.GetJSPointer(audio)
		if !ok {
			_audio = internal.StackAlloc(int(unsafe.Sizeof(*audio)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_PlayAudio",
			_mixer,
			_audio,
		)

		return internal.GetBool(ret)
	}

	iStopTrack = func(track *Track, fade_out_frames int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_fade_out_frames := internal.NewBigInt(fade_out_frames)
		ret := js.Global().Get("Module").Call(
			"_MIX_StopTrack",
			_track,
			_fade_out_frames,
		)

		return internal.GetBool(ret)
	}

	iStopAllTracks = func(mixer *Mixer, fade_out_ms int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_fade_out_ms := internal.NewBigInt(fade_out_ms)
		ret := js.Global().Get("Module").Call(
			"_MIX_StopAllTracks",
			_mixer,
			_fade_out_ms,
		)

		return internal.GetBool(ret)
	}

	iStopTag = func(mixer *Mixer, tag string, fade_out_ms int64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseTrack",
			_track,
		)

		return internal.GetBool(ret)
	}

	iPauseAllTracks = func(mixer *Mixer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseAllTracks",
			_mixer,
		)

		return internal.GetBool(ret)
	}

	iPauseTag = func(mixer *Mixer, tag string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_PauseTag",
			_mixer,
			_tag,
		)

		return internal.GetBool(ret)
	}

	iResumeTrack = func(track *Track) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeTrack",
			_track,
		)

		return internal.GetBool(ret)
	}

	iResumeAllTracks = func(mixer *Mixer) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeAllTracks",
			_mixer,
		)

		return internal.GetBool(ret)
	}

	iResumeTag = func(mixer *Mixer, tag string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_tag := internal.StringOnJSStack(tag)
		ret := js.Global().Get("Module").Call(
			"_MIX_ResumeTag",
			_mixer,
			_tag,
		)

		return internal.GetBool(ret)
	}

	iTrackPlaying = func(track *Track) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackPlaying",
			_track,
		)

		return internal.GetBool(ret)
	}

	iTrackPaused = func(track *Track) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_TrackPaused",
			_track,
		)

		return internal.GetBool(ret)
	}

	iSetMasterGain = func(mixer *Mixer, gain float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetMasterGain",
			_mixer,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iGetMasterGain = func(mixer *Mixer) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetMasterGain",
			_mixer,
		)

		return float32(ret.Int())
	}

	iSetTrackGain = func(track *Track, gain float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackGain",
			_track,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iGetTrackGain = func(track *Track) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackGain",
			_track,
		)

		return float32(ret.Int())
	}

	iSetTagGain = func(mixer *Mixer, tag string, gain float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_tag := internal.StringOnJSStack(tag)
		_gain := int32(gain)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTagGain",
			_mixer,
			_tag,
			_gain,
		)

		return internal.GetBool(ret)
	}

	iSetTrackFrequencyRatio = func(track *Track, ratio float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_ratio := int32(ratio)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackFrequencyRatio",
			_track,
			_ratio,
		)

		return internal.GetBool(ret)
	}

	iGetTrackFrequencyRatio = func(track *Track) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrackFrequencyRatio",
			_track,
		)

		return float32(ret.Int())
	}

	iSetTrackOutputChannelMap = func(track *Track, chmap *int32, count int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_chmap, ok := internal.GetJSPointer(chmap)
		if !ok {
			_chmap = internal.StackAlloc(int(unsafe.Sizeof(*chmap)))
		}
		_count := int32(count)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackOutputChannelMap",
			_track,
			_chmap,
			_count,
		)

		return internal.GetBool(ret)
	}

	iSetTrackStereo = func(track *Track, gains *StereoGains) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_gains, ok := internal.GetJSPointer(gains)
		if !ok {
			_gains = internal.StackAlloc(int(unsafe.Sizeof(*gains)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackStereo",
			_track,
			_gains,
		)

		return internal.GetBool(ret)
	}

	iSetTrack3DPosition = func(track *Track, position *Point3D) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_position, ok := internal.GetJSPointer(position)
		if !ok {
			_position = internal.StackAlloc(int(unsafe.Sizeof(*position)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrack3DPosition",
			_track,
			_position,
		)

		return internal.GetBool(ret)
	}

	iGetTrack3DPosition = func(track *Track, position *Point3D) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_position, ok := internal.GetJSPointer(position)
		if !ok {
			_position = internal.StackAlloc(int(unsafe.Sizeof(*position)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetTrack3DPosition",
			_track,
			_position,
		)

		return internal.GetBool(ret)
	}

	iCreateGroup = func(mixer *Mixer) *Group {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateGroup",
			_mixer,
		)

		_obj := internal.NewObject[Group](ret)
		return _obj
	}

	iDestroyGroup = func(group *Group) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_group, ok := internal.GetJSPointer(group)
		if !ok {
			_group = internal.StackAlloc(int(unsafe.Sizeof(*group)))
		}
		js.Global().Get("Module").Call(
			"_MIX_DestroyGroup",
			_group,
		)
	}

	iGetGroupProperties = func(group *Group) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_group, ok := internal.GetJSPointer(group)
		if !ok {
			_group = internal.StackAlloc(int(unsafe.Sizeof(*group)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetGroupProperties",
			_group,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetGroupMixer = func(group *Group) *Mixer {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_group, ok := internal.GetJSPointer(group)
		if !ok {
			_group = internal.StackAlloc(int(unsafe.Sizeof(*group)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetGroupMixer",
			_group,
		)

		_obj := internal.NewObject[Mixer](ret)
		return _obj
	}

	iSetTrackGroup = func(track *Track, group *Group) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_group, ok := internal.GetJSPointer(group)
		if !ok {
			_group = internal.StackAlloc(int(unsafe.Sizeof(*group)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackGroup",
			_track,
			_group,
		)

		return internal.GetBool(ret)
	}

	iSetTrackStoppedCallback = func(track *Track, cb TrackStoppedCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_cb := int32(cb)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackStoppedCallback",
			_track,
			_cb,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iSetTrackRawCallback = func(track *Track, cb TrackMixCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_cb := int32(cb)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackRawCallback",
			_track,
			_cb,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iSetTrackCookedCallback = func(track *Track, cb TrackMixCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_track, ok := internal.GetJSPointer(track)
		if !ok {
			_track = internal.StackAlloc(int(unsafe.Sizeof(*track)))
		}
		_cb := int32(cb)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetTrackCookedCallback",
			_track,
			_cb,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iSetGroupPostMixCallback = func(group *Group, cb GroupMixCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_group, ok := internal.GetJSPointer(group)
		if !ok {
			_group = internal.StackAlloc(int(unsafe.Sizeof(*group)))
		}
		_cb := int32(cb)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetGroupPostMixCallback",
			_group,
			_cb,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iSetPostMixCallback = func(mixer *Mixer, cb PostMixCallback, userdata uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_cb := int32(cb)
		_userdata := internal.NewBigInt(userdata)
		ret := js.Global().Get("Module").Call(
			"_MIX_SetPostMixCallback",
			_mixer,
			_cb,
			_userdata,
		)

		return internal.GetBool(ret)
	}

	iGenerate = func(mixer *Mixer, buffer uintptr, buflen int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mixer, ok := internal.GetJSPointer(mixer)
		if !ok {
			_mixer = internal.StackAlloc(int(unsafe.Sizeof(*mixer)))
		}
		_buffer := internal.NewBigInt(buffer)
		_buflen := int32(buflen)
		ret := js.Global().Get("Module").Call(
			"_MIX_Generate",
			_mixer,
			_buffer,
			_buflen,
		)

		return internal.GetBool(ret)
	}

	iCreateAudioDecoder = func(path string, props sdl.PropertiesID) *AudioDecoder {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		_props := uint32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateAudioDecoder",
			_path,
			_props,
		)

		_obj := internal.NewObject[AudioDecoder](ret)
		return _obj
	}

	iCreateAudioDecoder_IO = func(io *sdl.IOStream, closeio bool, props sdl.PropertiesID) *AudioDecoder {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_io, ok := internal.GetJSPointer(io)
		if !ok {
			_io = internal.StackAlloc(int(unsafe.Sizeof(*io)))
		}
		_closeio := internal.NewBoolean(closeio)
		_props := uint32(props)
		ret := js.Global().Get("Module").Call(
			"_MIX_CreateAudioDecoder_IO",
			_io,
			_closeio,
			_props,
		)

		_obj := internal.NewObject[AudioDecoder](ret)
		return _obj
	}

	iDestroyAudioDecoder = func(audiodecoder *AudioDecoder) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, ok := internal.GetJSPointer(audiodecoder)
		if !ok {
			_audiodecoder = internal.StackAlloc(int(unsafe.Sizeof(*audiodecoder)))
		}
		js.Global().Get("Module").Call(
			"_MIX_DestroyAudioDecoder",
			_audiodecoder,
		)
	}

	iGetAudioDecoderProperties = func(audiodecoder *AudioDecoder) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, ok := internal.GetJSPointer(audiodecoder)
		if !ok {
			_audiodecoder = internal.StackAlloc(int(unsafe.Sizeof(*audiodecoder)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoderProperties",
			_audiodecoder,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetAudioDecoderFormat = func(audiodecoder *AudioDecoder, spec *sdl.AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, ok := internal.GetJSPointer(audiodecoder)
		if !ok {
			_audiodecoder = internal.StackAlloc(int(unsafe.Sizeof(*audiodecoder)))
		}
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_GetAudioDecoderFormat",
			_audiodecoder,
			_spec,
		)

		return internal.GetBool(ret)
	}

	iDecodeAudio = func(audiodecoder *AudioDecoder, buffer uintptr, buflen int32, spec *sdl.AudioSpec) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_audiodecoder, ok := internal.GetJSPointer(audiodecoder)
		if !ok {
			_audiodecoder = internal.StackAlloc(int(unsafe.Sizeof(*audiodecoder)))
		}
		_buffer := internal.NewBigInt(buffer)
		_buflen := int32(buflen)
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_MIX_DecodeAudio",
			_audiodecoder,
			_buffer,
			_buflen,
			_spec,
		)

		return int32(ret.Int())
	}

}
