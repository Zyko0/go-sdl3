package sdl

import "github.com/Zyko0/go-sdl3/internal"

// Time

// TODO:

// SDL_GetTicks - Get the number of milliseconds since SDL library initialization.
// (https://wiki.libsdl.org/SDL3/SDL_GetTicks)
func Ticks() uint64 {
	return iGetTicks()
}

// Filesystem

// TODO:

// Storage

// TODO:

// Dialog

// TODO: ShowOpenFileDialog
// TODO: ShowSaveFileDialog
// TODO: ShowOpenFolderDialog
// TODO: ShowFileDialogWithProperties

// Message

// TODO: func ShowMessageBox(data *MessageBoxData)
// "buttonid => the pointer to which user id of hit button should be copied."
// I don't understand https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox

// SDL_ShowSimpleMessageBox - Display a simple modal message box.
// (https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox)
func ShowSimpleMessageBox(flags MessageBoxFlags, title, message string, window *Window) error {
	if !iShowSimpleMessageBox(flags, title, message, window) {
		return internal.LastErr()
	}

	return nil
}

// Power

type PowerInfo struct {
	Seconds int32
	Percent int32
	State   PowerState
}

// SDL_GetPowerInfo - Get the current power supply details.
// (https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo)
func GetPowerInfo() (PowerInfo, error) {
	var info PowerInfo

	info.State = iGetPowerInfo(&info.Seconds, &info.Percent)
	if info.State == POWERSTATE_ERROR {
		return info, internal.LastErr()
	}

	return info, nil
}

// Sensor

// SDL_GetSensors - Get a list of currently connected sensors.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensors)
func GetSensors() ([]SensorID, error) {
	var count int32

	ptr := iGetSensors(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[SensorID](ptr, int(count)), nil
}

// SDL_UpdateSensors - Update the current state of the open sensors.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateSensors)
func UpdateSensors() {
	iUpdateSensors()
}

// Process

// TODO: is this needed?

// Bits

// TODO: is this needed?

// Endian

// TODO: is this needed?

// Assert

// TODO: is this needed?

// CPU Info

// TODO: only intrinsics, how does that help in Go?

// Locale

// SDL_GetPreferredLocales - Report the user's preferred locale.
// (https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales)
func GetPreferredLocales() ([]*Locale, error) {
	var count int32

	ptr := iGetPreferredLocales(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Locale](ptr, int(count)), nil
}

// System

// TODO: platform specific stuff
