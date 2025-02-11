package sdl

import "unsafe"

// Video

func GetNumVideoDrivers() int {
	return iGetNumVideoDrivers()
}

func GetVideoDriver(index int) string {
	return iGetVideoDriver(index)
}

func GetCurrentVideoDriver() string {
	return iGetCurrentVideoDriver()
}

func GetSystemTheme() SystemTheme {
	return iGetSystemTheme()
}

func GetDisplays() ([]DisplayID, error) {
	var count int

	ptr := iGetDisplays(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[DisplayID](ptr, count), nil
}

func GetPrimaryDisplay() DisplayID {
	return iGetPrimaryDisplay()
}

func GetDisplayForPoint(point *Point) DisplayID {
	return iGetDisplayForPoint(point)
}

func GetDisplayForRect(rect *Rect) DisplayID {
	return iGetDisplayForRect(rect)
}

func GetDisplayForWindow(window *Window) DisplayID {
	return iGetDisplayForWindow(window)
}

func GetWindows() ([]*Window, error) {
	var count int

	ptr := iGetWindows(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[*Window](ptr, count), nil
}

func CreateWindow(title string, width, height int, flags WindowFlags) (*Window, error) {
	window := iCreateWindow(title, width, height, flags)
	if window == nil {
		return nil, lastError()
	}

	return window, nil
}

func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	window := iCreateWindowWithProperties(props)
	if window == nil {
		return nil, lastError()
	}

	return window, nil
}

func GetGrabbedWindow() *Window {
	return iGetGrabbedWindow()
}

func ScreenSaverEnabled() bool {
	return iScreenSaverEnabled()
}

func EnableScreenSaver() error {
	if !iEnableScreenSaver() {
		return lastError()
	}

	return nil
}

func DisableScreenSaver() error {
	if !iDisableScreenSaver() {
		return lastError()
	}

	return nil
}

// TODO: GL_ functions??

// Audio

func GetNumAudioDrivers() int {
	return iGetNumAudioDrivers()
}

func GetAudioDriver(index int) string {
	return iGetAudioDriver(index)
}

func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count int

	ptr := iGetAudioPlaybackDevices(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[AudioDeviceID](ptr, count), nil
}

func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count int

	ptr := iGetAudioRecordingDevices(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[AudioDeviceID](ptr, count), nil
}

func UnbindAudioStreams(streams []*AudioStream) {
	iUnbindAudioStreams(unsafe.SliceData(streams), len(streams))
}

func CreateAudioStream(srcSpec *AudioSpec) (*AudioStream, *AudioSpec, error) {
	dstSpec := &AudioSpec{}
	stream := iCreateAudioStream(srcSpec, dstSpec)
	if stream == nil {
		return nil, nil, lastError()
	}

	return stream, dstSpec, nil
}

func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV_IO(src, closeIO, spec, &ptr, &count) {
		return nil, lastError()
	}
	defer sdlFree(uintptr(unsafe.Pointer(ptr)))

	return clonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV(path, spec, &ptr, &count) {
		return nil, lastError()
	}
	defer sdlFree(uintptr(unsafe.Pointer(ptr)))

	return clonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

func MixAudio(src []byte, format AudioFormat, volume float32) ([]byte, error) {
	dst := make([]byte, len(src))
	if !iMixAudio(unsafe.SliceData(dst), unsafe.SliceData(src), format, uint32(len(src)), volume) {
		return nil, lastError()
	}

	return dst, nil
}

func ConvertAudioSamples(srcSpec *AudioSpec, srcData []byte) (*AudioSpec, []byte, error) {
	var ptr *byte
	var count int
	dstSpec := &AudioSpec{}

	if !iConvertAudioSamples(
		srcSpec, unsafe.SliceData(srcData), len(srcData),
		dstSpec, &ptr, &count,
	) {
		return nil, nil, lastError()
	}
	defer sdlFree(uintptr(unsafe.Pointer(ptr)))

	return dstSpec, clonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), count), nil
}
