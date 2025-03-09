//go:build js

package mixer

import (
	js "syscall/js"
	"unsafe"

	"github.com/Zyko0/go-sdl3/sdl"
	internal "github.com/Zyko0/go-sdl3/internal"
)

func init() {
	iVersion = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_Version",
		)

		return int32(ret.Int())
	}

	iInit = func(flags MIX_InitFlags) MIX_InitFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_flags := int32(flags)
		ret := js.Global().Get("Module").Call(
			"_SDL_Init",
			_flags,
		)

		return MIX_InitFlags(ret.Int())
	}

	iQuit = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_Quit",
		)
	}

	/*iOpenAudio = func(devid *sdl.AudioDeviceID, spec *sdl.AudioSpec) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_devid, ok := internal.GetJSPointer(devid)
		if !ok {
			_devid = internal.StackAlloc(int(unsafe.Sizeof(*devid)))
		}
		_spec, ok := internal.GetJSPointer(spec)
		if !ok {
			_spec = internal.StackAlloc(int(unsafe.Sizeof(*spec)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenAudio",
			_devid,
			_spec,
		)

		return internal.GetBool(ret)
	}*/

	iPauseAudio = func(pause_on int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_pause_on := int32(pause_on)
		js.Global().Get("Module").Call(
			"_SDL_PauseAudio",
			_pause_on,
		)
	}

	iQuerySpec = func(frequency *int32, format *sdl.AudioFormat, channels *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_frequency, ok := internal.GetJSPointer(frequency)
		if !ok {
			_frequency = internal.StackAlloc(int(unsafe.Sizeof(*frequency)))
		}
		_format, ok := internal.GetJSPointer(format)
		if !ok {
			_format = internal.StackAlloc(int(unsafe.Sizeof(*format)))
		}
		_channels, ok := internal.GetJSPointer(channels)
		if !ok {
			_channels = internal.StackAlloc(int(unsafe.Sizeof(*channels)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_QuerySpec",
			_frequency,
			_format,
			_channels,
		)

		return internal.GetBool(ret)
	}

	iAllocateChannels = func(numchans int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_numchans := int32(numchans)
		ret := js.Global().Get("Module").Call(
			"_SDL_AllocateChannels",
			_numchans,
		)

		return int32(ret.Int())
	}

	iLoadWAV_IO = func(src *sdl.IOStream, closeio bool) *Chunk {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadWAV_IO",
			_src,
			_closeio,
		)

		_obj := internal.NewObject[Chunk](ret)
		return _obj
	}

	iLoadWAV = func(file string) *Chunk {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadWAV",
			_file,
		)

		_obj := internal.NewObject[Chunk](ret)
		return _obj
	}

	iLoadMUS = func(file string) *Music {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadMUS",
			_file,
		)

		_obj := internal.NewObject[Music](ret)
		return _obj
	}

	iLoadMUS_IO = func(src *sdl.IOStream, closeio bool) *Music {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadMUS_IO",
			_src,
			_closeio,
		)

		_obj := internal.NewObject[Music](ret)
		return _obj
	}

	iLoadMUSType_IO = func(src *sdl.IOStream, typ MusicType, closeio bool) *Music {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_typ := int32(typ)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_SDL_LoadMUSType_IO",
			_src,
			_typ,
			_closeio,
		)

		_obj := internal.NewObject[Music](ret)
		return _obj
	}

	iQuickLoad_WAV = func(mem *uint8) *Chunk {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mem, ok := internal.GetJSPointer(mem)
		if !ok {
			_mem = internal.StackAlloc(int(unsafe.Sizeof(*mem)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_QuickLoad_WAV",
			_mem,
		)

		_obj := internal.NewObject[Chunk](ret)
		return _obj
	}

	iQuickLoad_RAW = func(mem *uint8, len uint32) *Chunk {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mem, ok := internal.GetJSPointer(mem)
		if !ok {
			_mem = internal.StackAlloc(int(unsafe.Sizeof(*mem)))
		}
		_len := int32(len)
		ret := js.Global().Get("Module").Call(
			"_SDL_QuickLoad_RAW",
			_mem,
			_len,
		)

		_obj := internal.NewObject[Chunk](ret)
		return _obj
	}

	iFreeChunk = func(chunk *Chunk) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		js.Global().Get("Module").Call(
			"_SDL_FreeChunk",
			_chunk,
		)
	}

	iFreeMusic = func(music *Music) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		js.Global().Get("Module").Call(
			"_SDL_FreeMusic",
			_music,
		)
	}

	iGetNumChunkDecoders = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumChunkDecoders",
		)

		return int32(ret.Int())
	}

	iGetChunkDecoder = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetChunkDecoder",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iHasChunkDecoder = func(name string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasChunkDecoder",
			_name,
		)

		return internal.GetBool(ret)
	}

	iGetNumMusicDecoders = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumMusicDecoders",
		)

		return int32(ret.Int())
	}

	iGetMusicDecoder = func(index int32) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_index := int32(index)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicDecoder",
			_index,
		)

		return internal.UTF8JSToString(ret)
	}

	iHasMusicDecoder = func(name string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_name := internal.StringOnJSStack(name)
		ret := js.Global().Get("Module").Call(
			"_SDL_HasMusicDecoder",
			_name,
		)

		return internal.GetBool(ret)
	}

	iGetMusicType = func(music *Music) MusicType {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicType",
			_music,
		)

		return MusicType(ret.Int())
	}

	iGetMusicTitle = func(music *Music) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicTitle",
			_music,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMusicTitleTag = func(music *Music) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicTitleTag",
			_music,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMusicArtistTag = func(music *Music) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicArtistTag",
			_music,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMusicAlbumTag = func(music *Music) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicAlbumTag",
			_music,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetMusicCopyrightTag = func(music *Music) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicCopyrightTag",
			_music,
		)

		return internal.UTF8JSToString(ret)
	}

	/*iSetPostMix = func(mix_func MixCallback, arg uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mix_func := int32(mix_func)
		_arg := internal.NewBigInt(arg)
		js.Global().Get("Module").Call(
			"_SDL_SetPostMix",
			_mix_func,
			_arg,
		)
	}*/

	/*iHookMusic = func(mix_func MixCallback, arg uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_mix_func := int32(mix_func)
		_arg := internal.NewBigInt(arg)
		js.Global().Get("Module").Call(
			"_SDL_HookMusic",
			_mix_func,
			_arg,
		)
	}*/

	/*iHookMusicFinished = func(music_finished MusicFinishedCallback) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music_finished := int32(music_finished)
		js.Global().Get("Module").Call(
			"_SDL_HookMusicFinished",
			_music_finished,
		)
	}*/

	/*iGetMusicHookData = func() uintptr {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicHookData",
		)

		return uintptr(internal.GetInt64(ret))
	}*/

	/*iChannelFinished = func(channel_finished ChannelFinishedCallback) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel_finished := int32(channel_finished)
		js.Global().Get("Module").Call(
			"_SDL_ChannelFinished",
			_channel_finished,
		)
	}*/

	/*iRegisterEffect = func(chann int32, f EffectFunc_t, d EffectDone_t, arg uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_chann := int32(chann)
		_f := int32(f)
		_d := int32(d)
		_arg := internal.NewBigInt(arg)
		ret := js.Global().Get("Module").Call(
			"_SDL_RegisterEffect",
			_chann,
			_f,
			_d,
			_arg,
		)

		return internal.GetBool(ret)
	}*/

	/*iUnregisterEffect = func(channel int32, f EffectFunc_t) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_f := int32(f)
		ret := js.Global().Get("Module").Call(
			"_SDL_UnregisterEffect",
			_channel,
			_f,
		)

		return internal.GetBool(ret)
	}*/

	iUnregisterAllEffects = func(channel int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		ret := js.Global().Get("Module").Call(
			"_SDL_UnregisterAllEffects",
			_channel,
		)

		return internal.GetBool(ret)
	}

	iSetPanning = func(channel int32, left uint8, right uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_left := int32(left)
		_right := int32(right)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPanning",
			_channel,
			_left,
			_right,
		)

		return internal.GetBool(ret)
	}

	iSetPosition = func(channel int32, angle int16, distance uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_angle := int32(angle)
		_distance := int32(distance)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetPosition",
			_channel,
			_angle,
			_distance,
		)

		return internal.GetBool(ret)
	}

	iSetDistance = func(channel int32, distance uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_distance := int32(distance)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetDistance",
			_channel,
			_distance,
		)

		return internal.GetBool(ret)
	}

	iSetReverseStereo = func(channel int32, flip int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_flip := int32(flip)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetReverseStereo",
			_channel,
			_flip,
		)

		return internal.GetBool(ret)
	}

	iReserveChannels = func(num int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_num := int32(num)
		ret := js.Global().Get("Module").Call(
			"_SDL_ReserveChannels",
			_num,
		)

		return int32(ret.Int())
	}

	iGroupChannel = func(which int32, tag int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_which := int32(which)
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupChannel",
			_which,
			_tag,
		)

		return internal.GetBool(ret)
	}

	iGroupChannels = func(from int32, to int32, tag int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_from := int32(from)
		_to := int32(to)
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupChannels",
			_from,
			_to,
			_tag,
		)

		return internal.GetBool(ret)
	}

	iGroupAvailable = func(tag int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupAvailable",
			_tag,
		)

		return int32(ret.Int())
	}

	iGroupCount = func(tag int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupCount",
			_tag,
		)

		return int32(ret.Int())
	}

	iGroupOldest = func(tag int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupOldest",
			_tag,
		)

		return int32(ret.Int())
	}

	iGroupNewer = func(tag int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		ret := js.Global().Get("Module").Call(
			"_SDL_GroupNewer",
			_tag,
		)

		return int32(ret.Int())
	}

	iPlayChannel = func(channel int32, chunk *Chunk, loops int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		_loops := int32(loops)
		ret := js.Global().Get("Module").Call(
			"_SDL_PlayChannel",
			_channel,
			_chunk,
			_loops,
		)

		return int32(ret.Int())
	}

	iPlayChannelTimed = func(channel int32, chunk *Chunk, loops int32, ticks int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		_loops := int32(loops)
		_ticks := int32(ticks)
		ret := js.Global().Get("Module").Call(
			"_SDL_PlayChannelTimed",
			_channel,
			_chunk,
			_loops,
			_ticks,
		)

		return int32(ret.Int())
	}

	iPlayMusic = func(music *Music, loops int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		_loops := int32(loops)
		ret := js.Global().Get("Module").Call(
			"_SDL_PlayMusic",
			_music,
			_loops,
		)

		return internal.GetBool(ret)
	}

	iFadeInMusic = func(music *Music, loops int32, ms int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		_loops := int32(loops)
		_ms := int32(ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeInMusic",
			_music,
			_loops,
			_ms,
		)

		return internal.GetBool(ret)
	}

	iFadeInMusicPos = func(music *Music, loops int32, ms int32, position float64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		_loops := int32(loops)
		_ms := int32(ms)
		_position := int32(position)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeInMusicPos",
			_music,
			_loops,
			_ms,
			_position,
		)

		return internal.GetBool(ret)
	}

	iFadeInChannel = func(channel int32, chunk *Chunk, loops int32, ms int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		_loops := int32(loops)
		_ms := int32(ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeInChannel",
			_channel,
			_chunk,
			_loops,
			_ms,
		)

		return int32(ret.Int())
	}

	iFadeInChannelTimed = func(channel int32, chunk *Chunk, loops int32, ms int32, ticks int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		_loops := int32(loops)
		_ms := int32(ms)
		_ticks := int32(ticks)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeInChannelTimed",
			_channel,
			_chunk,
			_loops,
			_ms,
			_ticks,
		)

		return int32(ret.Int())
	}

	iVolume = func(channel int32, volume int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_volume := int32(volume)
		ret := js.Global().Get("Module").Call(
			"_SDL_Volume",
			_channel,
			_volume,
		)

		return int32(ret.Int())
	}

	iVolumeChunk = func(chunk *Chunk, volume int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_chunk, ok := internal.GetJSPointer(chunk)
		if !ok {
			_chunk = internal.StackAlloc(int(unsafe.Sizeof(*chunk)))
		}
		_volume := int32(volume)
		ret := js.Global().Get("Module").Call(
			"_SDL_VolumeChunk",
			_chunk,
			_volume,
		)

		return int32(ret.Int())
	}

	iVolumeMusic = func(volume int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_volume := int32(volume)
		ret := js.Global().Get("Module").Call(
			"_SDL_VolumeMusic",
			_volume,
		)

		return int32(ret.Int())
	}

	iGetMusicVolume = func(music *Music) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicVolume",
			_music,
		)

		return int32(ret.Int())
	}

	iMasterVolume = func(volume int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_volume := int32(volume)
		ret := js.Global().Get("Module").Call(
			"_SDL_MasterVolume",
			_volume,
		)

		return int32(ret.Int())
	}

	iHaltChannel = func(channel int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		js.Global().Get("Module").Call(
			"_SDL_HaltChannel",
			_channel,
		)
	}

	iHaltGroup = func(tag int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		js.Global().Get("Module").Call(
			"_SDL_HaltGroup",
			_tag,
		)
	}

	iHaltMusic = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_HaltMusic",
		)
	}

	iExpireChannel = func(channel int32, ticks int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		_ticks := int32(ticks)
		ret := js.Global().Get("Module").Call(
			"_SDL_ExpireChannel",
			_channel,
			_ticks,
		)

		return int32(ret.Int())
	}

	iFadeOutChannel = func(which int32, ms int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_which := int32(which)
		_ms := int32(ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeOutChannel",
			_which,
			_ms,
		)

		return int32(ret.Int())
	}

	iFadeOutGroup = func(tag int32, ms int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		_ms := int32(ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeOutGroup",
			_tag,
			_ms,
		)

		return int32(ret.Int())
	}

	iFadeOutMusic = func(ms int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ms := int32(ms)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadeOutMusic",
			_ms,
		)

		return internal.GetBool(ret)
	}

	iFadingMusic = func() Fading {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_FadingMusic",
		)

		return Fading(ret.Int())
	}

	iFadingChannel = func(which int32) Fading {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_which := int32(which)
		ret := js.Global().Get("Module").Call(
			"_SDL_FadingChannel",
			_which,
		)

		return Fading(ret.Int())
	}

	iPause = func(channel int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		js.Global().Get("Module").Call(
			"_SDL_Pause",
			_channel,
		)
	}

	iPauseGroup = func(tag int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		js.Global().Get("Module").Call(
			"_SDL_PauseGroup",
			_tag,
		)
	}

	iResume = func(channel int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		js.Global().Get("Module").Call(
			"_SDL_Resume",
			_channel,
		)
	}

	iResumeGroup = func(tag int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		js.Global().Get("Module").Call(
			"_SDL_ResumeGroup",
			_tag,
		)
	}

	iPaused = func(channel int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		ret := js.Global().Get("Module").Call(
			"_SDL_Paused",
			_channel,
		)

		return int32(ret.Int())
	}

	iPauseMusic = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_PauseMusic",
		)
	}

	iResumeMusic = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_ResumeMusic",
		)
	}

	iRewindMusic = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_RewindMusic",
		)
	}

	iPausedMusic = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_PausedMusic",
		)

		return internal.GetBool(ret)
	}

	iModMusicJumpToOrder = func(order int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_order := int32(order)
		ret := js.Global().Get("Module").Call(
			"_SDL_ModMusicJumpToOrder",
			_order,
		)

		return internal.GetBool(ret)
	}

	iStartTrack = func(music *Music, track int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		_track := int32(track)
		ret := js.Global().Get("Module").Call(
			"_SDL_StartTrack",
			_music,
			_track,
		)

		return internal.GetBool(ret)
	}

	iGetNumTracks = func(music *Music) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumTracks",
			_music,
		)

		return int32(ret.Int())
	}

	iSetMusicPosition = func(position float64) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_position := int32(position)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetMusicPosition",
			_position,
		)

		return internal.GetBool(ret)
	}

	iGetMusicPosition = func(music *Music) float64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicPosition",
			_music,
		)

		return float64(ret.Int())
	}

	iMusicDuration = func(music *Music) float64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MusicDuration",
			_music,
		)

		return float64(ret.Int())
	}

	iGetMusicLoopStartTime = func(music *Music) float64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicLoopStartTime",
			_music,
		)

		return float64(ret.Int())
	}

	iGetMusicLoopEndTime = func(music *Music) float64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicLoopEndTime",
			_music,
		)

		return float64(ret.Int())
	}

	iGetMusicLoopLengthTime = func(music *Music) float64 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_music, ok := internal.GetJSPointer(music)
		if !ok {
			_music = internal.StackAlloc(int(unsafe.Sizeof(*music)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetMusicLoopLengthTime",
			_music,
		)

		return float64(ret.Int())
	}

	iPlaying = func(channel int32) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		ret := js.Global().Get("Module").Call(
			"_SDL_Playing",
			_channel,
		)

		return int32(ret.Int())
	}

	iPlayingMusic = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_PlayingMusic",
		)

		return internal.GetBool(ret)
	}

	iSetSoundFonts = func(paths string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_paths := internal.StringOnJSStack(paths)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetSoundFonts",
			_paths,
		)

		return internal.GetBool(ret)
	}

	iGetSoundFonts = func() string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetSoundFonts",
		)

		return internal.UTF8JSToString(ret)
	}

	/*iEachSoundFont = func(function EachSoundFontCallback, data uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_function := int32(function)
		_data := internal.NewBigInt(data)
		ret := js.Global().Get("Module").Call(
			"_SDL_EachSoundFont",
			_function,
			_data,
		)

		return internal.GetBool(ret)
	}*/

	iSetTimidityCfg = func(path string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_path := internal.StringOnJSStack(path)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTimidityCfg",
			_path,
		)

		return internal.GetBool(ret)
	}

	iGetTimidityCfg = func() string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTimidityCfg",
		)

		return internal.UTF8JSToString(ret)
	}

	iGetChunk = func(channel int32) *Chunk {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_channel := int32(channel)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetChunk",
			_channel,
		)

		_obj := internal.NewObject[Chunk](ret)
		return _obj
	}

	iCloseAudio = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_CloseAudio",
		)
	}

}
