package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// Video

func GetNumVideoDrivers() int {
	return int(iGetNumVideoDrivers())
}

func GetVideoDriver(index int) string {
	return iGetVideoDriver(int32(index))
}

func GetCurrentVideoDriver() string {
	return iGetCurrentVideoDriver()
}

func GetSystemTheme() SystemTheme {
	return iGetSystemTheme()
}

func GetDisplays() ([]DisplayID, error) {
	var count int32

	ptr := iGetDisplays(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[DisplayID](ptr, int(count)), nil
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
	var count int32

	ptr := iGetWindows(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Window](ptr, int(count)), nil
}

func CreateWindow(title string, width, height int, flags WindowFlags) (*Window, error) {
	window := iCreateWindow(title, int32(width), int32(height), flags)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	window := iCreateWindowWithProperties(props)
	if window == nil {
		return nil, internal.LastErr()
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
		return internal.LastErr()
	}

	return nil
}

func DisableScreenSaver() error {
	if !iDisableScreenSaver() {
		return internal.LastErr()
	}

	return nil
}

// TODO: GL_ functions??

// Audio

func GetNumAudioDrivers() int {
	return int(iGetNumAudioDrivers())
}

func GetAudioDriver(index int) string {
	return iGetAudioDriver(int32(index))
}

func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioPlaybackDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioRecordingDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

func UnbindAudioStreams(streams []*AudioStream) {
	iUnbindAudioStreams(unsafe.SliceData(streams), int32(len(streams)))
}

func CreateAudioStream(srcSpec *AudioSpec) (*AudioStream, *AudioSpec, error) {
	dstSpec := &AudioSpec{}
	stream := iCreateAudioStream(srcSpec, dstSpec)
	if stream == nil {
		return nil, nil, internal.LastErr()
	}

	return stream, dstSpec, nil
}

func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV_IO(src, closeIO, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV(path, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

func MixAudio(src []byte, format AudioFormat, volume float32) ([]byte, error) {
	dst := make([]byte, len(src))
	if !iMixAudio(unsafe.SliceData(dst), unsafe.SliceData(src), format, uint32(len(src)), volume) {
		return nil, internal.LastErr()
	}

	return dst, nil
}

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
