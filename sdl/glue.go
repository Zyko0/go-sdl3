package sdl

import (
	"errors"
	"fmt"
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

// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_guid.h#L61
type GUID *[16]uint8

// https://github.com/libsdl-org/SDL/blob/release-3.2.2/include/SDL3/SDL_version.h
type Version int32

func (v Version) Major() int {
	return int(v) / 1000000
}

func (v Version) Minor() int {
	return (int(v) / 1000) % 1000
}

func (v Version) Patch() int {
	return int(v) % 1000
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major(), v.Minor(), v.Patch())
}

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
	if e == nil {
		return nil
	}
	evt := *(*CommonEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) DisplayEvent() *DisplayEvent {
	if e == nil {
		return nil
	}
	evt := *(*DisplayEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) WindowEvent() *WindowEvent {
	if e == nil {
		return nil
	}
	evt := *(*WindowEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) KeyboardDeviceEvent() *KeyboardDeviceEvent {
	if e == nil {
		return nil
	}
	evt := *(*KeyboardDeviceEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) KeyboardEvent() *KeyboardEvent {
	if e == nil {
		return nil
	}
	evt := *(*KeyboardEvent)(unsafe.Pointer(e))
	return &evt
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
	if e == nil {
		return nil
	}
	evt := *(*MouseMotionEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) MouseButtonEvent() *MouseButtonEvent {
	if e == nil {
		return nil
	}
	evt := *(*MouseButtonEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) MouseWheelEvent() *MouseWheelEvent {
	if e == nil {
		return nil
	}
	evt := *(*MouseWheelEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyAxisEvent() *JoyAxisEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyAxisEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyBallEvent() *JoyBallEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyBallEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyHatEvent() *JoyHatEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyHatEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyButtonEvent() *JoyButtonEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyButtonEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyDeviceEvent() *JoyDeviceEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyDeviceEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) JoyBatteryEvent() *JoyBatteryEvent {
	if e == nil {
		return nil
	}
	evt := *(*JoyBatteryEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) GamepadAxisEvent() *GamepadAxisEvent {
	if e == nil {
		return nil
	}
	evt := *(*GamepadAxisEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) GamepadButtonEvent() *GamepadButtonEvent {
	if e == nil {
		return nil
	}
	evt := *(*GamepadButtonEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) GamepadDeviceEvent() *GamepadDeviceEvent {
	if e == nil {
		return nil
	}
	evt := *(*GamepadDeviceEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) GamepadTouchpadEvent() *GamepadTouchpadEvent {
	if e == nil {
		return nil
	}
	evt := *(*GamepadTouchpadEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) GamepadSensorEvent() *GamepadSensorEvent {
	if e == nil {
		return nil
	}
	evt := *(*GamepadSensorEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) AudioDeviceEvent() *AudioDeviceEvent {
	if e == nil {
		return nil
	}
	evt := *(*AudioDeviceEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) CameraDeviceEvent() *CameraDeviceEvent {
	if e == nil {
		return nil
	}
	evt := *(*CameraDeviceEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) RenderEvent() *RenderEvent {
	if e == nil {
		return nil
	}
	evt := *(*RenderEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) TouchFingerEvent() *TouchFingerEvent {
	if e == nil {
		return nil
	}
	evt := *(*TouchFingerEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) PenProximityEvent() *PenProximityEvent {
	if e == nil {
		return nil
	}
	evt := *(*PenProximityEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) PenMotionEvent() *PenMotionEvent {
	if e == nil {
		return nil
	}
	evt := *(*PenMotionEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) PenTouchEvent() *PenTouchEvent {
	if e == nil {
		return nil
	}
	evt := *(*PenTouchEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) PenButtonEvent() *PenButtonEvent {
	if e == nil {
		return nil
	}
	evt := *(*PenButtonEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) PenAxisEvent() *PenAxisEvent {
	if e == nil {
		return nil
	}
	evt := *(*PenAxisEvent)(unsafe.Pointer(e))
	return &evt
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
	if e == nil {
		return nil
	}
	evt := *(*SensorEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) QuitEvent() *QuitEvent {
	if e == nil {
		return nil
	}
	evt := *(*QuitEvent)(unsafe.Pointer(e))
	return &evt
}

func (e *Event) UserEvent() *UserEvent {
	if e == nil {
		return nil
	}
	evt := *(*UserEvent)(unsafe.Pointer(e))
	return &evt
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

type messageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int32
	Text     *byte
}

func (mbd *MessageBoxButtonData) as() messageBoxButtonData {
	return messageBoxButtonData{
		Flags:    mbd.Flags,
		ButtonID: mbd.ButtonID,
		Text:     internal.StringToNullablePtr(mbd.Text),
	}
}

type messageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window
	Title       *byte
	Message     *byte
	Numbuttons  int32
	Buttons     *messageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

func (md *MessageBoxData) as() *messageBoxData {
	if md == nil {
		return nil
	}
	buttons := make([]messageBoxButtonData, len(md.Buttons))
	for i, button := range md.Buttons {
		buttons[i] = button.as()
	}
	return &messageBoxData{
		Flags:       md.Flags,
		Window:      md.Window,
		Title:       internal.StringToNullablePtr(md.Title),
		Message:     internal.StringToNullablePtr(md.Message),
		Numbuttons:  int32(len(buttons)),
		Buttons:     unsafe.SliceData(buttons),
		ColorScheme: md.ColorScheme,
	}
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

func (info *GPUShaderCreateInfo) as() *gpuShaderCreateInfo {
	if info == nil {
		return nil
	}
	return &gpuShaderCreateInfo{
		CodeSize:           uintptr(len(info.Code)),
		Code:               unsafe.SliceData(info.Code),
		Entrypoint:         internal.StringToNullablePtr(info.Entrypoint),
		Format:             info.Format,
		Stage:              info.Stage,
		NumSamplers:        info.NumSamplers,
		NumStorageTextures: info.NumStorageTextures,
		NumStorageBuffers:  info.NumStorageBuffers,
		NumUniformBuffers:  info.NumUniformBuffers,
		Props:              info.Props,
	}
}

// SDL_GPUShaderCreateInfo - A structure specifying code and metadata for creating a shader object.
// (https://wiki.libsdl.org/SDL3/SDL_GPUShaderCreateInfo)
type GPUShaderCreateInfo struct {
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

func (info *GPUComputePipelineCreateInfo) as() *gpuComputePipelineCreateInfo {
	if info == nil {
		return nil
	}
	return &gpuComputePipelineCreateInfo{
		CodeSize:                    uintptr(len(info.Code)),
		Code:                        unsafe.SliceData(info.Code),
		Entrypoint:                  internal.StringToNullablePtr(info.Entrypoint),
		Format:                      info.Format,
		NumSamplers:                 info.NumSamplers,
		NumReadonlyStorageTextures:  info.NumReadonlyStorageTextures,
		NumReadonlyStorageBuffers:   info.NumReadonlyStorageBuffers,
		NumReadwriteStorageTextures: info.NumReadwriteStorageTextures,
		NumReadwriteStorageBuffers:  info.NumReadwriteStorageBuffers,
		NumUniformBuffers:           info.NumUniformBuffers,
		ThreadcountX:                info.ThreadcountX,
		ThreadcountY:                info.ThreadcountY,
		ThreadcountZ:                info.ThreadcountZ,
		Props:                       info.Props,
	}
}

// SDL_GPUComputePipelineCreateInfo - A structure specifying the parameters of a compute pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUComputePipelineCreateInfo)
type GPUComputePipelineCreateInfo struct {
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

type gpuGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription
	NumColorTargets         uint32
	DepthStencilFormat      GPUTextureFormat
	HasDepthStencilTarget   bool
	Padding1                uint8
	Padding2                uint8
	Padding3                uint8
}

func (info *GPUGraphicsPipelineTargetInfo) as() gpuGraphicsPipelineTargetInfo {
	return gpuGraphicsPipelineTargetInfo{
		ColorTargetDescriptions: unsafe.SliceData(info.ColorTargetDescriptions),
		NumColorTargets:         uint32(len(info.ColorTargetDescriptions)),
		DepthStencilFormat:      info.DepthStencilFormat,
		HasDepthStencilTarget:   info.HasDepthStencilTarget,
		Padding1:                0,
		Padding2:                0,
		Padding3:                0,
	}
}

// SDL_GPUGraphicsPipelineTargetInfo - A structure specifying the descriptions of render targets used in a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineTargetInfo)
type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions []GPUColorTargetDescription
	DepthStencilFormat      GPUTextureFormat
	HasDepthStencilTarget   bool
}

type gpuVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription
	NumVertexBuffers         uint32
	VertexAttributes         *GPUVertexAttribute
	NumVertexAttributes      uint32
}

func (state *GPUVertexInputState) as() gpuVertexInputState {
	return gpuVertexInputState{
		VertexBufferDescriptions: unsafe.SliceData(state.VertexBufferDescriptions),
		NumVertexBuffers:         uint32(len(state.VertexBufferDescriptions)),
		VertexAttributes:         unsafe.SliceData(state.VertexAttributes),
		NumVertexAttributes:      uint32(len(state.VertexAttributes)),
	}
}

// SDL_GPUVertexInputState - A structure specifying the parameters of a graphics pipeline vertex input state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputState)
type GPUVertexInputState struct {
	VertexBufferDescriptions []GPUVertexBufferDescription
	VertexAttributes         []GPUVertexAttribute
}

type gpuGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader
	FragmentShader    *GPUShader
	VertexInputState  gpuVertexInputState
	PrimitiveType     GPUPrimitiveType
	RasterizerState   GPURasterizerState
	MultisampleState  GPUMultisampleState
	DepthStencilState GPUDepthStencilState
	TargetInfo        gpuGraphicsPipelineTargetInfo
	Props             PropertiesID
}

func (info *GPUGraphicsPipelineCreateInfo) as() *gpuGraphicsPipelineCreateInfo {
	if info == nil {
		return nil
	}
	return &gpuGraphicsPipelineCreateInfo{
		VertexShader:      info.VertexShader,
		FragmentShader:    info.FragmentShader,
		VertexInputState:  info.VertexInputState.as(),
		PrimitiveType:     info.PrimitiveType,
		RasterizerState:   info.RasterizerState,
		MultisampleState:  info.MultisampleState,
		DepthStencilState: info.DepthStencilState,
		TargetInfo:        info.TargetInfo.as(),
		Props:             info.Props,
	}
}

// SDL_GPUGraphicsPipelineCreateInfo - A structure specifying the parameters of a graphics pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineCreateInfo)
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader
	FragmentShader    *GPUShader
	VertexInputState  GPUVertexInputState
	PrimitiveType     GPUPrimitiveType
	RasterizerState   GPURasterizerState
	MultisampleState  GPUMultisampleState
	DepthStencilState GPUDepthStencilState
	TargetInfo        GPUGraphicsPipelineTargetInfo
	Props             PropertiesID
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

// SDL_CleanupPropertyCallback - A callback used to free resources when a property is deleted.
// (https://wiki.libsdl.org/SDL3/SDL_CleanupPropertyCallback)
type CleanupPropertyCallback uintptr

// SDL_EnumeratePropertiesCallback - A callback used to enumerate all the properties in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_EnumeratePropertiesCallback)
type EnumeratePropertiesCallback uintptr

// SDL_TLSDestructorCallback - The callback used to cleanup data passed to [SDL_SetTLS](SDL_SetTLS).
// (https://wiki.libsdl.org/SDL3/SDL_TLSDestructorCallback)
type TLSDestructorCallback uintptr

// SDL_AudioStreamCallback - A callback that fires when data passes through an [SDL_AudioStream](SDL_AudioStream).
// (https://wiki.libsdl.org/SDL3/SDL_AudioStreamCallback)
type AudioStreamCallback uintptr

// SDL_AudioPostmixCallback - A callback that fires when data is about to be fed to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_AudioPostmixCallback)
type AudioPostmixCallback uintptr

// SDL_ClipboardDataCallback - Callback function that will be called when data for the specified mime-type is requested by the OS.
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardDataCallback)
type ClipboardDataCallback uintptr

// SDL_ClipboardCleanupCallback - Callback function that will be called when the clipboard is cleared, or new data is set.
// (https://wiki.libsdl.org/SDL3/SDL_ClipboardCleanupCallback)
type ClipboardCleanupCallback uintptr

// SDL_FunctionPointer - A generic function pointer.
// (https://wiki.libsdl.org/SDL3/SDL_FunctionPointer)
type FunctionPointer uintptr

// SDL_EGLAttribArrayCallback - EGL platform attribute initialization callback.
// (https://wiki.libsdl.org/SDL3/SDL_EGLAttribArrayCallback)
type EGLAttribArrayCallback uintptr

// SDL_EGLIntArrayCallback - EGL surface/context attribute initialization callback types.
// (https://wiki.libsdl.org/SDL3/SDL_EGLIntArrayCallback)
type EGLIntArrayCallback uintptr

// SDL_DialogFileCallback - Callback used by file dialog functions.
// (https://wiki.libsdl.org/SDL3/SDL_DialogFileCallback)
type DialogFileCallback uintptr

// SDL_EnumerateDirectoryCallback - Callback for directory enumeration.
// (https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectoryCallback)
type EnumerateDirectoryCallback uintptr

// SDL_HintCallback - A callback used to send notifications of hint value changes.
// (https://wiki.libsdl.org/SDL3/SDL_HintCallback)
type HintCallback uintptr

// SDL_MainThreadCallback - Callback run on the main thread.
// (https://wiki.libsdl.org/SDL3/SDL_MainThreadCallback)
type MainThreadCallback uintptr

// SDL_LogOutputFunction - The prototype for the log output callback function.
// (https://wiki.libsdl.org/SDL3/SDL_LogOutputFunction)
type LogOutputFunction uintptr

// SDL_X11EventHook - A callback to be used with [SDL_SetX11EventHook](SDL_SetX11EventHook).
// (https://wiki.libsdl.org/SDL3/SDL_X11EventHook)
type X11EventHook uintptr

// SDL_TimerCallback - Function prototype for the millisecond timer callback function.
// (https://wiki.libsdl.org/SDL3/SDL_TimerCallback)
type TimerCallback uintptr

// SDL_NSTimerCallback - Function prototype for the nanosecond timer callback function.
// (https://wiki.libsdl.org/SDL3/SDL_NSTimerCallback)
type NSTimerCallback uintptr

// SDL_main_func - The prototype for the application's main() function
// (https://wiki.libsdl.org/SDL3/SDL_main_func)
type main_func uintptr

// SDL_AppInit_func - Function pointer typedef for [SDL_AppInit](SDL_AppInit).
// (https://wiki.libsdl.org/SDL3/SDL_AppInit_func)
type AppInit_func uintptr

// SDL_AppIterate_func - Function pointer typedef for [SDL_AppIterate](SDL_AppIterate).
// (https://wiki.libsdl.org/SDL3/SDL_AppIterate_func)
type AppIterate_func uintptr

// SDL_AppEvent_func - Function pointer typedef for [SDL_AppEvent](SDL_AppEvent).
// (https://wiki.libsdl.org/SDL3/SDL_AppEvent_func)
type AppEvent_func uintptr

// SDL_AppQuit_func - Function pointer typedef for [SDL_AppQuit](SDL_AppQuit).
// (https://wiki.libsdl.org/SDL3/SDL_AppQuit_func)
type AppQuit_func uintptr

// SDL_EventFilter - A function pointer used for callbacks that watch the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_EventFilter)
type EventFilter uintptr

type AudioStreamDataCompleteCallback uintptr

type MouseMotionTransformCallback uintptr
