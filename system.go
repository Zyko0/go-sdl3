package sdl

// Time

// TODO:

func GetTicks() uint64 {
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

func ShowSimpleMessageBox(flags MessageBoxFlags, title, message string, window *Window) error {
	if !iShowSimpleMessageBox(flags, title, message, window) {
		return lastError()
	}

	return nil
}

// Power

type PowerInfo struct {
	Seconds int
	Percent int
	State   PowerState
}

func GetPowerInfo() (PowerInfo, error) {
	var info PowerInfo

	info.State = iGetPowerInfo(&info.Seconds, &info.Percent)
	if info.State == POWERSTATE_ERROR {
		return info, lastError()
	}

	return info, nil
}

// Sensor

func GetSensors() ([]SensorID, error) {
	var count int

	ptr := iGetSensors(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[SensorID](ptr, count), nil
}

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

func GetPreferredLocales() ([]*Locale, error) {
	var count int

	ptr := iGetPreferredLocales(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[*Locale](ptr, count), nil
}

// System

// TODO: platform specific stuff