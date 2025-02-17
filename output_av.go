package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// Video

// SDL_GetNumVideoDrivers - Get the number of video drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers)
func GetNumVideoDrivers() int {
	return int(iGetNumVideoDrivers())
}

// SDL_GetVideoDriver - Get the name of a built in video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return iGetVideoDriver(int32(index))
}

// SDL_GetCurrentVideoDriver - Get the name of the currently initialized video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver)
func GetCurrentVideoDriver() string {
	return iGetCurrentVideoDriver()
}

// SDL_GetSystemTheme - Get the current system theme.
// (https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme)
func GetSystemTheme() SystemTheme {
	return iGetSystemTheme()
}

// SDL_GetDisplays - Get a list of currently connected displays.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplays)
func GetDisplays() ([]DisplayID, error) {
	var count int32

	ptr := iGetDisplays(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[DisplayID](ptr, int(count)), nil
}

// SDL_GetPrimaryDisplay - Return the primary display.
// (https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay)
func GetPrimaryDisplay() DisplayID {
	return iGetPrimaryDisplay()
}

// SDL_GetDisplayForPoint - Get the display containing a point.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint)
func GetDisplayForPoint(point *Point) DisplayID {
	return iGetDisplayForPoint(point)
}

// SDL_GetDisplayForRect - Get the display primarily containing a rect.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect)
func GetDisplayForRect(rect *Rect) DisplayID {
	return iGetDisplayForRect(rect)
}

// SDL_GetDisplayForWindow - Get the display associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow)
func GetDisplayForWindow(window *Window) DisplayID {
	return iGetDisplayForWindow(window)
}

// SDL_GetWindows - Get a list of valid windows.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindows)
func GetWindows() ([]*Window, error) {
	var count int32

	ptr := iGetWindows(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Window](ptr, int(count)), nil
}

// SDL_CreateWindow - Create a window with the specified dimensions and flags.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindow)
func CreateWindow(title string, width, height int, flags WindowFlags) (*Window, error) {
	window := iCreateWindow(title, int32(width), int32(height), flags)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_CreateWindowWithProperties - Create a window with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties)
func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	window := iCreateWindowWithProperties(props)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_GetGrabbedWindow - Get the window that currently has an input grab enabled.
// (https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow)
func GetGrabbedWindow() *Window {
	return iGetGrabbedWindow()
}

// SDL_ScreenSaverEnabled - Check whether the screensaver is currently enabled.
// (https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled)
func ScreenSaverEnabled() bool {
	return iScreenSaverEnabled()
}

// SDL_EnableScreenSaver - Allow the screen to be blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver)
func EnableScreenSaver() error {
	if !iEnableScreenSaver() {
		return internal.LastErr()
	}

	return nil
}

// SDL_DisableScreenSaver - Prevent the screen from being blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver)
func DisableScreenSaver() error {
	if !iDisableScreenSaver() {
		return internal.LastErr()
	}

	return nil
}

// TODO: GL_ functions??

// Audio

// SDL_GetNumAudioDrivers - Use this function to get the number of built-in audio drivers.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers)
func GetNumAudioDrivers() int {
	return int(iGetNumAudioDrivers())
}

// SDL_GetAudioDriver - Use this function to get the name of a built in audio driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver)
func GetAudioDriver(index int) string {
	return iGetAudioDriver(int32(index))
}

// SDL_GetAudioPlaybackDevices - Get a list of currently-connected audio playback devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices)
func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioPlaybackDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

// SDL_GetAudioRecordingDevices - Get a list of currently-connected audio recording devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices)
func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioRecordingDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

// SDL_UnbindAudioStreams - Unbind a list of audio streams from their audio devices.
// (https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams)
func UnbindAudioStreams(streams []*AudioStream) {
	iUnbindAudioStreams(unsafe.SliceData(streams), int32(len(streams)))
}

// SDL_CreateAudioStream - Create a new audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream)
func CreateAudioStream(srcSpec *AudioSpec) (*AudioStream, *AudioSpec, error) {
	dstSpec := &AudioSpec{}
	stream := iCreateAudioStream(srcSpec, dstSpec)
	if stream == nil {
		return nil, nil, internal.LastErr()
	}

	return stream, dstSpec, nil
}

// SDL_LoadWAV_IO - Load the audio data of a WAVE file into memory.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO)
func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV_IO(src, closeIO, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// SDL_LoadWAV - Loads a WAV from a file path.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV)
func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV(path, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// SDL_MixAudio - Mix audio data in a specified format.
// (https://wiki.libsdl.org/SDL3/SDL_MixAudio)
func MixAudio(src []byte, format AudioFormat, volume float32) ([]byte, error) {
	dst := make([]byte, len(src))
	if !iMixAudio(unsafe.SliceData(dst), unsafe.SliceData(src), format, uint32(len(src)), volume) {
		return nil, internal.LastErr()
	}

	return dst, nil
}

// SDL_ConvertAudioSamples - Convert some audio data of one format to another format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples)
func ConvertAudioSamples(srcSpec *AudioSpec, srcData []byte) (*AudioSpec, []byte, error) {
	var ptr *byte
	var count int32
	dstSpec := &AudioSpec{}

	if !iConvertAudioSamples(
		srcSpec, unsafe.SliceData(srcData), int32(len(srcData)),
		dstSpec, &ptr, &count,
	) {
		return nil, nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return dstSpec, internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}
