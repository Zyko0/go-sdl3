package sdl

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// This file contains extra glue code for types and enums that couldn't be
// generated automatically.
// It includes union types, function callbacks, #defines and more.

// Functions

// Types

type Time int64

// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_guid.h#L61
type GUID *[16]uint8

// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_gamepad.h#L262
// TODO: Union type
type GamepadBinding struct {
	InputType  int32
	InputData  [12]byte
	OutputType GamepadBindingType
	OutputData [12]byte
}

// TODO: union type
// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_events.h#L986
type Event struct {
	Type EventType
	data [124]byte
	//EventData [48]byte // 52 is size of SDL_TextEditingEvent minus Type (uint32)
	//_         [76]byte // Padding (128 required in total)
}

func (e *Event) CommonEvent() *CommonEvent {
	return (*CommonEvent)(unsafe.Pointer(e))
}

func (e *Event) DisplayEvent() *DisplayEvent {
	return (*DisplayEvent)(unsafe.Pointer(e))
}

func (e *Event) WindowEvent() *WindowEvent {
	return (*WindowEvent)(unsafe.Pointer(e))
}

func (e *Event) KeyboardDeviceEvent() *KeyboardDeviceEvent {
	return (*KeyboardDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) KeyboardEvent() *KeyboardEvent {
	return (*KeyboardEvent)(unsafe.Pointer(e))
}

func (e *Event) TextEditingEvent() *TextEditingEvent {
	return (*TextEditingEvent)(unsafe.Pointer(e))
}

func (e *Event) TextEditingCandidatesEvent() *TextEditingCandidatesEvent {
	return (*TextEditingCandidatesEvent)(unsafe.Pointer(e))
}

func (e *Event) TextInputEvent() *TextInputEvent {
	return (*TextInputEvent)(unsafe.Pointer(e))
}

func (e *Event) MouseMotionEvent() *MouseMotionEvent {
	return (*MouseMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) MouseButtonEvent() *MouseButtonEvent {
	return (*MouseButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) MouseWheelEvent() *MouseWheelEvent {
	return (*MouseWheelEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyAxisEvent() *JoyAxisEvent {
	return (*JoyAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyBallEvent() *JoyBallEvent {
	return (*JoyBallEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyHatEvent() *JoyHatEvent {
	return (*JoyHatEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyButtonEvent() *JoyButtonEvent {
	return (*JoyButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyDeviceEvent() *JoyDeviceEvent {
	return (*JoyDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) JoyBatteryEvent() *JoyBatteryEvent {
	return (*JoyBatteryEvent)(unsafe.Pointer(e))
}

func (e *Event) GamepadAxisEvent() *GamepadAxisEvent {
	return (*GamepadAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) GamepadButtonEvent() *GamepadButtonEvent {
	return (*GamepadButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) GamepadDeviceEvent() *GamepadDeviceEvent {
	return (*GamepadDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) GamepadTouchpadEvent() *GamepadTouchpadEvent {
	return (*GamepadTouchpadEvent)(unsafe.Pointer(e))
}

func (e *Event) GamepadSensorEvent() *GamepadSensorEvent {
	return (*GamepadSensorEvent)(unsafe.Pointer(e))
}

func (e *Event) AudioDeviceEvent() *AudioDeviceEvent {
	return (*AudioDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) CameraDeviceEvent() *CameraDeviceEvent {
	return (*CameraDeviceEvent)(unsafe.Pointer(e))
}

func (e *Event) RenderEvent() *RenderEvent {
	return (*RenderEvent)(unsafe.Pointer(e))
}

func (e *Event) TouchFingerEvent() *TouchFingerEvent {
	return (*TouchFingerEvent)(unsafe.Pointer(e))
}

func (e *Event) PenProximityEvent() *PenProximityEvent {
	return (*PenProximityEvent)(unsafe.Pointer(e))
}

func (e *Event) PenMotionEvent() *PenMotionEvent {
	return (*PenMotionEvent)(unsafe.Pointer(e))
}

func (e *Event) PenTouchEvent() *PenTouchEvent {
	return (*PenTouchEvent)(unsafe.Pointer(e))
}

func (e *Event) PenButtonEvent() *PenButtonEvent {
	return (*PenButtonEvent)(unsafe.Pointer(e))
}

func (e *Event) PenAxisEvent() *PenAxisEvent {
	return (*PenAxisEvent)(unsafe.Pointer(e))
}

func (e *Event) DropEvent() *DropEvent {
	return (*DropEvent)(unsafe.Pointer(e))
}

func (e *Event) ClipboardEvent() *ClipboardEvent {
	return (*ClipboardEvent)(unsafe.Pointer(e))
}

func (e *Event) SensorEvent() *SensorEvent {
	return (*SensorEvent)(unsafe.Pointer(e))
}

func (e *Event) QuitEvent() *QuitEvent {
	return (*QuitEvent)(unsafe.Pointer(e))
}

func (e *Event) UserEvent() *UserEvent {
	return (*UserEvent)(unsafe.Pointer(e))
}

// TODO: union type
type HapticEffect struct {
	Type       uint16
	HapticData [66]byte // 68 is full size of SDL_HapticCondition
}

type HitTest uintptr // TODO: supposed to be a callback but can't find the signature

type va_list uintptr // TODO: not done yet

type Surface struct {
	Flags    SurfaceFlags
	Format   PixelFormat
	W        int32
	H        int32
	Pitch    int32
	pixels   *byte
	Refcount int32
	Reserved uintptr
}

func (s *Surface) Pixels() []byte {
	return internal.PtrToSlice[byte](
		uintptr(unsafe.Pointer(s.pixels)),
		int(s.H*s.Pitch),
	)
}

// Custom types

type SwapchainTexture struct {
	Texture *GPUTexture
	Width   uint32
	Height  uint32
}

// Callback types

type (
	CleanupPropertyCallback     func()
	EnumeratePropertiesCallback func()
	TLSDestructorCallback       func()
	AudioStreamCallback         func()
	AudioPostmixCallback        func()
	ClipboardDataCallback       func()
	ClipboardCleanupCallback    func()
	FunctionPointer             func()
	EGLAttribArrayCallback      func()
	EGLIntArrayCallback         func()
	DialogFileCallback          func()
	EnumerateDirectoryCallback  func()
	HintCallback                func()
	MainThreadCallback          func()
	LogOutputFunction           func()
	X11EventHook                func()
	TimerCallback               func()
	NSTimerCallback             func()
	main_func                   func()
	AppInit_func                func()
	AppIterate_func             func()
	AppEvent_func               func()
	AppQuit_func                func()
	EventFilter                 func()
)

// Constants (#define)

const (
	// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_init.h#L80-L87
	INIT_AUDIO    InitFlags = 0x00000010
	INIT_VIDEO    InitFlags = 0x00000020
	INIT_JOYSTICK InitFlags = 0x00000200
	INIT_HAPTIC   InitFlags = 0x00001000
	INIT_GAMEPAD  InitFlags = 0x00002000
	INIT_EVENTS   InitFlags = 0x00004000
	INIT_SENSOR   InitFlags = 0x00008000
	INIT_CAMERA   InitFlags = 0x00010000
)

const (
	// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_video.h#L189-L213
	WINDOW_FULLSCREEN          WindowFlags = 0x0000000000000001
	WINDOW_OPENGL              WindowFlags = 0x0000000000000002
	WINDOW_OCCLUDED            WindowFlags = 0x0000000000000004
	WINDOW_HIDDEN              WindowFlags = 0x0000000000000008
	WINDOW_BORDERLESS          WindowFlags = 0x0000000000000010
	WINDOW_RESIZABLE           WindowFlags = 0x0000000000000020
	WINDOW_MINIMIZED           WindowFlags = 0x0000000000000040
	WINDOW_MAXIMIZED           WindowFlags = 0x0000000000000080
	WINDOW_MOUSE_GRABBED       WindowFlags = 0x0000000000000100
	WINDOW_INPUT_FOCUS         WindowFlags = 0x0000000000000200
	WINDOW_MOUSE_FOCUS         WindowFlags = 0x0000000000000400
	WINDOW_EXTERNAL            WindowFlags = 0x0000000000000800
	WINDOW_MODAL               WindowFlags = 0x0000000000001000
	WINDOW_HIGH_PIXEL_DENSITY  WindowFlags = 0x0000000000002000
	WINDOW_MOUSE_CAPTURE       WindowFlags = 0x0000000000004000
	WINDOW_MOUSE_RELATIVE_MODE WindowFlags = 0x0000000000008000
	WINDOW_ALWAYS_ON_TOP       WindowFlags = 0x0000000000010000
	WINDOW_UTILITY             WindowFlags = 0x0000000000020000
	WINDOW_TOOLTIP             WindowFlags = 0x0000000000040000
	WINDOW_POPUP_MENU          WindowFlags = 0x0000000000080000
	WINDOW_KEYBOARD_GRABBED    WindowFlags = 0x0000000000100000
	WINDOW_VULKAN              WindowFlags = 0x0000000010000000
	WINDOW_METAL               WindowFlags = 0x0000000020000000
	WINDOW_TRANSPARENT         WindowFlags = 0x0000000040000000
	WINDOW_NOT_FOCUSABLE       WindowFlags = 0x0000000080000000
)

const (
	HAT_CENTERED  uint8 = 0x00
	HAT_UP        uint8 = 0x01
	HAT_RIGHT     uint8 = 0x02
	HAT_DOWN      uint8 = 0x04
	HAT_LEFT      uint8 = 0x08
	HAT_RIGHTUP   uint8 = HAT_RIGHT | HAT_UP
	HAT_RIGHTDOWN uint8 = HAT_RIGHT | HAT_DOWN
	HAT_LEFTUP    uint8 = HAT_LEFT | HAT_UP
	HAT_LEFTDOWN  uint8 = HAT_LEFT | HAT_DOWN
)

const DEBUG_TEXT_FONT_CHARACTER_SIZE = 8
