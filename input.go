package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// Keyboard

// SDL_HasKeyboard - Return whether a keyboard is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasKeyboard)
func HasKeyboard() bool {
	return iHasKeyboard()
}

// SDL_GetKeyboards - Get a list of currently connected keyboards.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboards)
func GetKeyboards() ([]KeyboardID, error) {
	var count int32

	ptr := iGetKeyboards(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[KeyboardID](ptr, int(count)), nil
}

// SDL_GetKeyboardState - Get a snapshot of the current state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState)
func GetKeyboardState() []bool {
	var count int32

	ptr := iGetKeyboardState(&count)

	return internal.PtrToSlice[bool](uintptr(unsafe.Pointer(ptr)), int(count))
}

// SDL_ResetKeyboard - Clear the state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard)
func ResetKeyboard() {
	iResetKeyboard()
}

// SDL_GetScancodeFromName - Get a scancode from a human-readable name.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName)
func GetScancodeFromName(name string) Scancode {
	return iGetScancodeFromName(name)
}

// SDL_GetKeyFromName - Get a key code from a human-readable name.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName)
func GetKeyFromName(name string) Keycode {
	return iGetKeyFromName(name)
}

// SDL_HasScreenKeyboardSupport - Check whether the platform has screen keyboard support.
// (https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport)
func HasScreenKeyboardSupport() bool {
	return iHasScreenKeyboardSupport()
}

// Mouse

// SDL_HasMouse - Return whether a mouse is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasMouse)
func HasMouse() bool {
	return iHasMouse()
}

// SDL_GetMice - Get a list of currently connected mice.
// (https://wiki.libsdl.org/SDL3/SDL_GetMice)
func GetMice() ([]MouseID, error) {
	var count int32

	ptr := iGetMice(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[MouseID](ptr, int(count)), nil
}

// SDL_GetMouseFocus - Get the window which currently has mouse focus.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus)
func GetMouseFocus() *Window {
	return iGetMouseFocus()
}

// SDL_GetMouseState - Query SDL's cache for the synchronous mouse button state and the window-relative SDL-cursor position.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseState)
func GetMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetMouseState(&x, &y)

	return flags, x, y
}

// SDL_GetGlobalMouseState - Query the platform for the asynchronous mouse button state and the desktop-relative platform-cursor position.
// (https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState)
func GetGlobalMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetGlobalMouseState(&x, &y)

	return flags, x, y
}

// SDL_GetRelativeMouseState - Query SDL's cache for the synchronous mouse button state and accumulated mouse delta since last call.
// (https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState)
func GetRelativeMouseState() (MouseButtonFlags, float32, float32) {
	var x, y float32

	flags := iGetRelativeMouseState(&x, &y)

	return flags, x, y
}

// SDL_WarpMouseGlobal - Move the mouse to the given position in global screen space.
// (https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal)
func WarpMouseGlobal(x, y float32) error {
	if !iWarpMouseGlobal(x, y) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CaptureMouse - Capture the mouse and to track input outside an SDL window.
// (https://wiki.libsdl.org/SDL3/SDL_CaptureMouse)
func CaptureMouse(enabled bool) error {
	if !iCaptureMouse(enabled) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CreateCursor - Create a cursor using the specified bitmap data and mask (in MSB format).
// (https://wiki.libsdl.org/SDL3/SDL_CreateCursor)
func CreateCursor(data, mask []byte, width, height, hotX, hotY int) (*Cursor, error) {
	cursor := iCreateCursor(
		unsafe.SliceData(data),
		unsafe.SliceData(mask),
		int32(width), int32(height), int32(hotX), int32(hotY),
	)
	if cursor == nil {
		return nil, internal.LastErr()
	}

	return cursor, nil
}

// SDL_SetCursor - Set the active cursor.
// (https://wiki.libsdl.org/SDL3/SDL_SetCursor)
func SetCursor(cursor *Cursor) error {
	if !iSetCursor(cursor) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetCursor - Get the active cursor.
// (https://wiki.libsdl.org/SDL3/SDL_GetCursor)
func GetCursor() *Cursor {
	return iGetCursor()
}

// SDL_GetDefaultCursor - Get the default cursor.
// (https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor)
func GetDefaultCursor() (*Cursor, error) {
	cursor := iGetDefaultCursor()
	if cursor == nil {
		return nil, internal.LastErr()
	}

	return cursor, nil
}

// SDL_ShowCursor - Show the cursor.
// (https://wiki.libsdl.org/SDL3/SDL_ShowCursor)
func ShowCursor() error {
	if !iShowCursor() {
		return internal.LastErr()
	}

	return nil
}

// SDL_HideCursor - Hide the cursor.
// (https://wiki.libsdl.org/SDL3/SDL_HideCursor)
func HideCursor() error {
	if !iHideCursor() {
		return internal.LastErr()
	}

	return nil
}

// SDL_CursorVisible - Return whether the cursor is currently being shown.
// (https://wiki.libsdl.org/SDL3/SDL_CursorVisible)
func CursorVisible() bool {
	return iCursorVisible()
}

// Touch

// SDL_GetTouchDevices - Get a list of registered touch devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices)
func GetTouchDevices() ([]TouchID, error) {
	var count int32

	ptr := iGetTouchDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[TouchID](ptr, int(count)), nil
}

// Gamepad

// SDL_AddGamepadMapping - Add support for gamepads that SDL is unaware of or change the binding of an existing gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping)
func AddGamepadMapping(mapping string) error {
	if iAddGamepadMapping(mapping) == -1 {
		return internal.LastErr()
	}

	return nil
}

// SDL_AddGamepadMappingsFromFile - Load a set of gamepad mappings from a file.
// (https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile)
func AddGamepadMappingsFromFile(file string) error {
	if iAddGamepadMappingsFromFile(file) == -1 {
		return internal.LastErr()
	}

	return nil
}

// SDL_ReloadGamepadMappings - Reinitialize the SDL mapping database to its initial state.
// (https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings)
func ReloadGamepadMappings() error {
	if !iReloadGamepadMappings() {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetGamepadMappings - Get the current gamepad mappings.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings)
func GetGamepadMappings() ([]string, error) {
	var count int32

	ptr := iGetGamepadMappings(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[string](ptr, int(count)), nil
}

// SDL_HasGamepad - Return whether a gamepad is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasGamepad)
func HasGamepad() bool {
	return iHasGamepad()
}

// SDL_GetGamepads - Get a list of currently connected gamepads.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepads)
func GetGamepads() ([]JoystickID, error) {
	var count int32

	ptr := iGetGamepads(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[JoystickID](ptr, int(count)), nil
}

// SDL_GetGamepadFromPlayerIndex - Get the SDL_Gamepad associated with a player index.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex)
func GetGamepadFromPlayerIndex(playerIndex int) *Gamepad {
	return iGetGamepadFromPlayerIndex(int32(playerIndex))
}

// SDL_SetGamepadEventsEnabled - Set the state of gamepad event processing.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled)
func SetGamepadEventsEnabled(enabled bool) {
	iSetGamepadEventsEnabled(enabled)
}

// SDL_GamepadEventsEnabled - Query the state of gamepad event processing.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled)
func GamepadEventsEnabled() bool {
	return iGamepadEventsEnabled()
}

// SDL_UpdateGamepads - Manually pump gamepad updates if not using the loop.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads)
func UpdateGamepads() {
	iUpdateGamepads()
}

// SDL_GetGamepadAxisFromString - Convert a string into SDL_GamepadAxis enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString)
func GetGamepadAxisFromString(str string) GamepadAxis {
	return iGetGamepadAxisFromString(str)
}

// SDL_GetGamepadButtonFromString - Convert a string into an SDL_GamepadButton enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString)
func GetGamepadButtonFromString(str string) GamepadButton {
	return iGetGamepadButtonFromString(str)
}

// Joystick

// SDL_LockJoysticks - Locking for atomic access to the joystick API.
// (https://wiki.libsdl.org/SDL3/SDL_LockJoysticks)
func LockJoysticks() {
	iLockJoysticks()
}

// SDL_UnlockJoysticks - Unlocking for atomic access to the joystick API.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks)
func UnlockJoysticks() {
	iUnlockJoysticks()
}

// SDL_HasJoystick - Return whether a joystick is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_HasJoystick)
func HasJoystick() bool {
	return iHasJoystick()
}

// SDL_GetJoysticks - Get a list of currently connected joysticks.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoysticks)
func GetJoysticks() ([]JoystickID, error) {
	var count int32

	ptr := iGetJoysticks(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[JoystickID](ptr, int(count)), nil
}

// SDL_GetJoystickFromPlayerIndex - Get the SDL_Joystick associated with a player index.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex)
func GetJoystickFromPlayerIndex(playerIndex int) *Joystick {
	return iGetJoystickFromPlayerIndex(int32(playerIndex))
}

// SDL_AttachVirtualJoystick - Attach a new virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick)
func AttachVirtualJoystick(desc *VirtualJoystickDesc) JoystickID {
	return iAttachVirtualJoystick(desc)
}

// SDL_SetJoystickEventsEnabled - Set the state of joystick event processing.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled)
func SetJoystickEventsEnabled(enabled bool) {
	iSetJoystickEventsEnabled(enabled)
}

// SDL_JoystickEventsEnabled - Query the state of joystick event processing.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled)
func JoystickEventsEnabled() bool {
	return iJoystickEventsEnabled()
}

// SDL_UpdateJoysticks - Update the current state of the open joysticks.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks)
func UpdateJoysticks() {
	iUpdateJoysticks()
}

// Haptic

// TODO:

// Camera

// SDL_GetNumCameraDrivers - Use this function to get the number of built-in camera drivers.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumCameraDrivers)
func GetNumCameraDrivers() int {
	return int(iGetNumCameraDrivers())
}

// SDL_GetCameraDriver - Use this function to get the name of a built in camera driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraDriver)
func GetCameraDriver(index int) string {
	return iGetCameraDriver(int32(index))
}

// SDL_GetCurrentCameraDriver - Get the name of the current camera driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentCameraDriver)
func GetCurrentCameraDriver() string {
	return iGetCurrentCameraDriver()
}

// SDL_GetCameras - Get a list of currently connected camera devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameras)
func GetCameras() ([]CameraID, error) {
	var count int32

	ptr := iGetCameras(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[CameraID](ptr, int(count)), nil
}

// Clipboard

// SDL_SetClipboardText - Put UTF-8 text into the clipboard.
// (https://wiki.libsdl.org/SDL3/SDL_SetClipboardText)
func SetClipboardText(text string) error {
	if !iSetClipboardText(text) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetClipboardText - Get UTF-8 text from the clipboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetClipboardText)
func GetClipboardText() (string, error) {
	ptr := iGetClipboardText()
	if ptr == 0 {
		return "", internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrString(ptr), nil
}
