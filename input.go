package sdl

import (
	"unsafe"
)

// Keyboard

func HasKeyboard() bool {
	return iHasKeyboard()
}

func GetKeyboards() ([]KeyboardID, error) {
	var count int

	ptr := iGetKeyboards(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[KeyboardID](ptr, count), nil
}

func GetKeyboardState() []bool {
	var count int

	ptr := iGetKeyboardState(&count)

	return ptrToSlice[bool](uintptr(unsafe.Pointer(ptr)), count)
}

func ResetKeyboard() {
	iResetKeyboard()
}

func GetScancodeFromName(name string) Scancode {
	return iGetScancodeFromName(name)
}

func GetKeyFromName(name string) Keycode {
	return iGetKeyFromName(name)
}

func HasScreenKeyboardSupport() bool {
	return iHasScreenKeyboardSupport()
}

// Mouse

func HasMouse() bool {
	return iHasMouse()
}

func GetMice() ([]MouseID, error) {
	var count int

	ptr := iGetMice(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)
	
	return clonePtrSlice[MouseID](ptr, count), nil
}

func GetMouseFocus() *Window {
	return iGetMouseFocus()
}

func GetMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetMouseState(&x, &y)

	return flags, x, y
}

func GetGlobalMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetGlobalMouseState(&x, &y)

	return flags, x, y
}

func GetRelativeMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetRelativeMouseState(&x, &y)

	return flags, x, y
}

func WarpMouseGlobal(x, y float32) error {
	if !iWarpMouseGlobal(x, y) {
		return lastError()
	}

	return nil
}

func CaptureMouse(enabled bool) error {
	if !iCaptureMouse(enabled) {
		return lastError()
	}

	return nil
}

func CreateCursor(data, mask []byte, width, height, hotX, hotY int) (*Cursor, error) {
	cursor := iCreateCursor(
		unsafe.SliceData(data),
		unsafe.SliceData(mask),
		width, height, hotX, hotY,
	)
	if cursor == nil {
		return nil, lastError()
	}

	return cursor, nil
}

func SetCursor(cursor *Cursor) error {
	if !iSetCursor(cursor) {
		return lastError()
	}

	return nil
}

func GetCursor() *Cursor {
	return iGetCursor()
}

func GetDefaultCursor() (*Cursor, error) {
	cursor := iGetDefaultCursor()
	if cursor == nil {
		return nil, lastError()
	}

	return cursor, nil
}

func ShowCursor() error {
	if !iShowCursor() {
		return lastError()
	}

	return nil
}

func HideCursor() error {
	if !iHideCursor() {
		return lastError()
	}

	return nil
}

func CursorVisible() bool {
	return iCursorVisible()
}

// Touch

func GetTouchDevices() ([]TouchID, error) {
	var count int

	ptr := iGetTouchDevices(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[TouchID](ptr, count), nil
}

// Gamepad

func AddGamepadMapping(mapping string) error {
	if iAddGamepadMapping(mapping) == -1 {
		return lastError()
	}

	return nil
}

func AddGamepadMappingsFromFile(file string) error {
	if iAddGamepadMappingsFromFile(file) == -1 {
		return lastError()
	}

	return nil
}

func ReloadGamepadMappings() error {
	if !iReloadGamepadMappings() {
		return lastError()
	}

	return nil
}

func GetGamepadMappings() ([]string, error) {
	var count int

	ptr := iGetGamepadMappings(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)
	
	return clonePtrSlice[string](ptr, count), nil
}

func HasGamepad() bool {
	return iHasGamepad()
}

func GetGamepads() ([]JoystickID, error) {
	var count int

	ptr := iGetGamepads(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)
	
	return clonePtrSlice[JoystickID](ptr, count), nil
}

func GetGamepadFromPlayerIndex(playerIndex int) *Gamepad {
	return iGetGamepadFromPlayerIndex(playerIndex)
}

func SetGamepadEventsEnabled(enabled bool) {
	iSetGamepadEventsEnabled(enabled)
}

func GamepadEventsEnabled() bool {
	return iGamepadEventsEnabled()
}

func UpdateGamepads() {
	iUpdateGamepads()
}

func GetGamepadAxisFromString(str string) GamepadAxis {
	return iGetGamepadAxisFromString(str)
}

func GetGamepadButtonFromString(str string) GamepadButton {
	return iGetGamepadButtonFromString(str)
}

// Joystick

func LockJoysticks() {
	iLockJoysticks()
}

func UnlockJoysticks() {
	iUnlockJoysticks()
}

func HasJoystick() bool {
	return iHasJoystick()
}

func GetJoysticks() ([]JoystickID, error) {
	var count int

	ptr := iGetJoysticks(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)
	
	return clonePtrSlice[JoystickID](ptr, count), nil
}

func GetJoystickFromPlayerIndex(playerIndex int) *Joystick {
	return iGetJoystickFromPlayerIndex(playerIndex)
}

func AttachVirtualJoystick(desc *VirtualJoystickDesc) JoystickID {
	return iAttachVirtualJoystick(desc)
}

func SetJoystickEventsEnabled(enabled bool) {
	iSetJoystickEventsEnabled(enabled)
}

func JoystickEventsEnabled() bool {
	return iJoystickEventsEnabled()
}

func UpdateJoysticks() {
	iUpdateJoysticks()
}

// Haptic

// TODO:

// Camera

func GetNumCameraDrivers() int {
	return iGetNumCameraDrivers()
}

func GetCameraDriver(index int) string {
	return iGetCameraDriver(index)
}

func GetCurrentCameraDriver() string {
	return iGetCurrentCameraDriver()
}

func GetCameras() ([]CameraID, error) {
	var count int

	ptr := iGetCameras(&count)
	if ptr == 0 {
		return nil, lastError()
	}
	defer sdlFree(ptr)

	return clonePtrSlice[CameraID](ptr, count), nil
}

// Clipboard

func SetClipboardText(text string) error {
	if !iSetClipboardText(text) {
		return lastError()
	}

	return nil
}

func GetClipboardText() (string, error) {
	ptr := iGetClipboardText()
	if ptr == 0 {
		return "", lastError()
	}
	defer sdlFree(ptr)

	return clonePtrString(ptr), nil
}
