package sdl

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

var EndLoop = errors.New("graceful termination")

// This file contains extra glue code for types and enums that couldn't be
// generated automatically.
// It includes union types, function callbacks, #defines and more.

// Functions

// Types

type Pointer = internal.Pointer

// SDL_Time - SDL times are signed, 64-bit integers representing nanoseconds since the Unix epoch (Jan 1, 1970).
// (https://wiki.libsdl.org/SDL3/SDL_Time)
type Time int64

// SDL_GLProfile - Possible values to be set for the [SDL_GL_CONTEXT_PROFILE_MASK](SDL_GL_CONTEXT_PROFILE_MASK) attribute.
// (https://wiki.libsdl.org/SDL3/SDL_GLProfile)
type GLProfile uint32

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

// This is required because the original struct contains a string and since
// it comes from C, we need to handle the pointer manually.
type textEditingEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Text      *byte
	Start     int32
	Length    int32
}

func (e *Event) TextEditingEvent() *TextEditingEvent {
	impl := (*textEditingEvent)(unsafe.Pointer(e))
	if impl == nil {
		return nil
	}
	return &TextEditingEvent{
		Type:      impl.Type,
		Reserved:  impl.Reserved,
		Timestamp: impl.Timestamp,
		WindowID:  impl.WindowID,
		Text:      internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Text))),
		Start:     impl.Start,
		Length:    impl.Length,
	}
}

// The original structure contains a candidates pointer, we want to
// turn it into a slice.
type textEditingCandidatesEvent struct {
	Type              EventType
	Reserved          uint32
	Timestamp         uint64
	WindowID          WindowID
	Candidates        *string
	NumCandidates     int32
	SelectedCandidate int32
	Horizontal        bool
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

// SDL_TextEditingCandidatesEvent - Keyboard IME candidates event structure (event.edit_candidates.*)
// (https://wiki.libsdl.org/SDL3/SDL_TextEditingCandidatesEvent)
type TextEditingCandidatesEvent struct {
	Type              EventType
	Reserved          uint32
	Timestamp         uint64
	WindowID          WindowID
	Candidates        []string
	NumCandidates     int32
	SelectedCandidate int32
	Horizontal        bool
}

func (e *Event) TextEditingCandidatesEvent() *TextEditingCandidatesEvent {
	impl := (*textEditingCandidatesEvent)(unsafe.Pointer(e))
	if impl == nil {
		return nil
	}
	return &TextEditingCandidatesEvent{
		Type:              impl.Type,
		Reserved:          impl.Reserved,
		Timestamp:         impl.Timestamp,
		WindowID:          impl.WindowID,
		Candidates:        internal.ClonePtrSlice[string](uintptr(unsafe.Pointer(impl.Candidates)), int(impl.NumCandidates)),
		SelectedCandidate: impl.SelectedCandidate,
		NumCandidates:     impl.NumCandidates,
		Horizontal:        impl.Horizontal,
	}
}

// This is required because the original struct contains a string and since
// it comes from C, we need to handle the pointer manually.
type textInputEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Text      *byte
}

func (e *Event) TextInputEvent() *TextInputEvent {
	impl := (*textInputEvent)(unsafe.Pointer(e))
	if impl == nil {
		return nil
	}
	return &TextInputEvent{
		Type:      impl.Type,
		Reserved:  impl.Reserved,
		Timestamp: impl.Timestamp,
		WindowID:  impl.WindowID,
		Text:      internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Text))),
	}
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

// This is required because the original struct contains a string and since
// it comes from C, we need to handle the pointer manually.
type dropEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	X         float32
	Y         float32
	Source    *byte
	Data      *byte
}

func (e *Event) DropEvent() *DropEvent {
	impl := (*dropEvent)(unsafe.Pointer(e))
	if impl == nil {
		return nil
	}
	return &DropEvent{
		Type:      impl.Type,
		Reserved:  impl.Reserved,
		Timestamp: impl.Timestamp,
		WindowID:  impl.WindowID,
		X:         impl.X,
		Y:         impl.Y,
		Source:    internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Source))),
		Data:      internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Data))),
	}
}

type clipboardEvent struct {
	Type         EventType
	Reserved     uint32
	Timestamp    uint64
	Owner        bool
	NumMimeTypes int32
	MimeTypes    *string
}

// SDL_ClipboardEvent - An event triggered when the clipboard contents have changed (event.clipboard.*)
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardEvent)
type ClipboardEvent struct {
	Type         EventType
	Reserved     uint32
	Timestamp    uint64
	Owner        bool
	NumMimeTypes int32
	MimeTypes    []string
}

func (e *Event) ClipboardEvent() *ClipboardEvent {
	impl := (*clipboardEvent)(unsafe.Pointer(e))
	if impl == nil {
		return nil
	}
	return &ClipboardEvent{
		Type:         impl.Type,
		Reserved:     impl.Reserved,
		Timestamp:    impl.Timestamp,
		Owner:        impl.Owner,
		NumMimeTypes: impl.NumMimeTypes,
		MimeTypes:    internal.ClonePtrSlice[string](uintptr(unsafe.Pointer(impl.MimeTypes)), int(impl.NumMimeTypes)),
	}
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

// SDL_HitTest - Callback used for hit-testing.
// (https://wiki.libsdl.org/SDL3/SDL_HitTest)
type HitTest uintptr // TODO: supposed to be a callback but can't find the signature

type va_list uintptr // TODO: not done yet

// SDL_Surface - A collection of pixels used in software blitting.
// (https://wiki.libsdl.org/SDL3/SDL_Surface)
type Surface struct {
	Flags    SurfaceFlags
	Format   PixelFormat
	W        int32
	H        int32
	Pitch    int32
	pixels   Pointer
	Refcount int32
	Reserved Pointer
}

// SDL_MessageBoxData - MessageBox structure containing title, text, window, etc.
// (https://wiki.libsdl.org/SDL3/SDL_MessageBoxData)
type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window
	Title       string
	Message     string
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

type gpuShaderCreateInfo struct {
	CodeSize           uintptr
	Code               *uint8
	Entrypoint         *byte
	Format             GPUShaderFormat
	Stage              GPUShaderStage
	NumSamplers        uint32
	NumStorageTextures uint32
	NumStorageBuffers  uint32
	NumUniformBuffers  uint32
	Props              PropertiesID
}

// SDL_GPUShaderCreateInfo - A structure specifying code and metadata for creating a shader object.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShaderCreateInfo)
type GPUShaderCreateInfo struct {
	CodeSize           uint64
	Code               []byte
	Entrypoint         string
	Format             GPUShaderFormat
	Stage              GPUShaderStage
	NumSamplers        uint32
	NumStorageTextures uint32
	NumStorageBuffers  uint32
	NumUniformBuffers  uint32
	Props              PropertiesID
}

type gpuComputePipelineCreateInfo struct {
	CodeSize                    uintptr
	Code                        *byte
	Entrypoint                  *byte
	Format                      GPUShaderFormat
	NumSamplers                 uint32
	NumReadonlyStorageTextures  uint32
	NumReadonlyStorageBuffers   uint32
	NumReadwriteStorageTextures uint32
	NumReadwriteStorageBuffers  uint32
	NumUniformBuffers           uint32
	ThreadcountX                uint32
	ThreadcountY                uint32
	ThreadcountZ                uint32
	Props                       PropertiesID
}

// SDL_GPUComputePipelineCreateInfo - A structure specifying the parameters of a compute pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUComputePipelineCreateInfo)
type GPUComputePipelineCreateInfo struct {
	CodeSize                    uint64
	Code                        []byte
	Entrypoint                  string
	Format                      GPUShaderFormat
	NumSamplers                 uint32
	NumReadonlyStorageTextures  uint32
	NumReadonlyStorageBuffers   uint32
	NumReadwriteStorageTextures uint32
	NumReadwriteStorageBuffers  uint32
	NumUniformBuffers           uint32
	ThreadcountX                uint32
	ThreadcountY                uint32
	ThreadcountZ                uint32
	Props                       PropertiesID
}

// SDL_Palette - A set of indexed colors representing a palette.
// (https://wiki.libsdl.org/SDL3/SDL_Palette)
type Palette struct {
	ncolors  int32
	colors   *Color
	Version  uint32
	Refcount int32
}

// Custom types

type locale struct {
	Language *byte
	Country  *byte
}

type SwapchainTexture struct {
	Texture *GPUTexture
	Width   uint32
	Height  uint32
}

type BorderSize struct {
	Top    int32
	Left   int32
	Bottom int32
	Right  int32
}

type ProcessData struct {
	ExitCode int32
	Data     []byte
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
