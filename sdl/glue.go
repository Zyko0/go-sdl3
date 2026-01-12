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

// SDL_GamepadBinding - A mapping between one joystick input to a gamepad control.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadBinding)
// Union type
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
		reserved:  impl.Reserved,
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
	Type              EventType // SDL_EVENT_TEXT_EDITING_CANDIDATES
	Reserved          uint32
	Timestamp         uint64   // In nanoseconds, populated using SDL_GetTicksNS()
	WindowID          WindowID // The window with keyboard focus, if any
	Candidates        []string // The list of candidates, or NULL if there are no candidates available
	NumCandidates     int32    // The number of strings in `candidates`
	SelectedCandidate int32    // The index of the selected candidate, or -1 if no candidate is selected
	Horizontal        bool     // true if the list is horizontal, false if it's vertical
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
		reserved:  impl.Reserved,
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
		reserved:  impl.Reserved,
		Timestamp: impl.Timestamp,
		WindowID:  impl.WindowID,
		X:         impl.X,
		Y:         impl.Y,
		Source:    internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Source))),
		Data:      internal.ClonePtrString(uintptr(unsafe.Pointer(impl.Data))),
	}
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

type clipboardEvent struct {
	Type         EventType
	Reserved     uint32
	Timestamp    uint64
	Owner        bool
	NumMimeTypes int32
	MimeTypes    *string
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

// SDL_HapticEffect - The generic template for any haptic effect.
// (https://wiki.libsdl.org/SDL3/SDL_HapticEffect)
// Union type
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
	Flags    SurfaceFlags // The flags of the surface, read-only
	Format   PixelFormat  // The format of the surface, read-only
	W        int32        // The width of the surface, read-only.
	H        int32        // The height of the surface, read-only.
	Pitch    int32        // The distance in bytes between rows of pixels, read-only
	pixels   Pointer
	Refcount int32   // Application reference count, used when freeing surface
	reserved Pointer // Reserved for internal use
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
	Window      *Window // Parent window, can be NULL
	Title       string  // UTF-8 title
	Message     string  // UTF-8 message text
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme // SDL_MessageBoxColorScheme, can be NULL to use system settings
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
	Format             GPUShaderFormat // The format of the shader code.
	Stage              GPUShaderStage  // The stage the shader program corresponds to.
	NumSamplers        uint32          // The number of samplers defined in the shader.
	NumStorageTextures uint32          // The number of storage textures defined in the shader.
	NumStorageBuffers  uint32          // The number of storage buffers defined in the shader.
	NumUniformBuffers  uint32          // The number of uniform buffers defined in the shader.
	Props              PropertiesID    // A properties ID for extensions. Should be 0 if no extensions are needed.
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
	Format                      GPUShaderFormat // The format of the compute shader code.
	NumSamplers                 uint32          // The number of samplers defined in the shader.
	NumReadonlyStorageTextures  uint32          // The number of readonly storage textures defined in the shader.
	NumReadonlyStorageBuffers   uint32          // The number of readonly storage buffers defined in the shader.
	NumReadwriteStorageTextures uint32          // The number of read-write storage textures defined in the shader.
	NumReadwriteStorageBuffers  uint32          // The number of read-write storage buffers defined in the shader.
	NumUniformBuffers           uint32          // The number of uniform buffers defined in the shader.
	ThreadcountX                uint32          // The number of threads in the X dimension. This should match the value in the shader.
	ThreadcountY                uint32          // The number of threads in the Y dimension. This should match the value in the shader.
	ThreadcountZ                uint32          // The number of threads in the Z dimension. This should match the value in the shader.
	Props                       PropertiesID    // A properties ID for extensions. Should be 0 if no extensions are needed.
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

// SDL_GPUVertexInputState - A structure specifying the parameters of a graphics pipeline vertex input state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUVertexInputState)
type GPUVertexInputState struct {
	VertexBufferDescriptions []GPUVertexBufferDescription
	VertexAttributes         []GPUVertexAttribute
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

// SDL_GPUGraphicsPipelineCreateInfo - A structure specifying the parameters of a graphics pipeline state.
// (https://wiki.libsdl.org/SDL3/SDL_GPUGraphicsPipelineCreateInfo)
type GPUGraphicsPipelineCreateInfo struct {
	VertexShader      *GPUShader                    // The vertex shader used by the graphics pipeline.
	FragmentShader    *GPUShader                    // The fragment shader used by the graphics pipeline.
	VertexInputState  GPUVertexInputState           // The vertex layout of the graphics pipeline.
	PrimitiveType     GPUPrimitiveType              // The primitive topology of the graphics pipeline.
	RasterizerState   GPURasterizerState            // The rasterizer state of the graphics pipeline.
	MultisampleState  GPUMultisampleState           // The multisample state of the graphics pipeline.
	DepthStencilState GPUDepthStencilState          // The depth-stencil state of the graphics pipeline.
	TargetInfo        GPUGraphicsPipelineTargetInfo // Formats and blend modes for the render targets of the graphics pipeline.
	Props             PropertiesID                  // A properties ID for extensions. Should be 0 if no extensions are needed.
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

// SDL_Palette - A set of indexed colors representing a palette.
// (https://wiki.libsdl.org/SDL3/SDL_Palette)
type Palette struct {
	ncolors  int32
	colors   *Color
	version  uint32
	refcount int32
}

// SDL_VirtualJoystickDesc - The structure that describes a virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickDesc)
type VirtualJoystickDesc struct {
	Version           uint32                        // the version of this interface
	Type              JoystickType                  // `SDL_JoystickType`
	VendorId          uint16                        // the USB vendor ID of this joystick
	ProductId         uint16                        // the USB product ID of this joystick
	Naxes             uint16                        // the number of axes on this joystick
	Nbuttons          uint16                        // the number of buttons on this joystick
	Nballs            uint16                        // the number of balls on this joystick
	Nhats             uint16                        // the number of hats on this joystick
	ButtonMask        uint32                        // A mask of which buttons are valid for this controller e.g. (1 << SDL_GAMEPAD_BUTTON_SOUTH)
	AxisMask          uint32                        // A mask of which axes are valid for this controller e.g. (1 << SDL_GAMEPAD_AXIS_LEFTX)
	Name              string                        // the name of the joystick
	Touchpads         []VirtualJoystickTouchpadDesc // A pointer to an array of touchpad descriptions, required if `ntouchpads` is > 0
	Sensors           []VirtualJoystickSensorDesc   // A pointer to an array of sensor descriptions, required if `nsensors` is > 0
	Update            Pointer                       // Called when the joystick state should be updated
	SetPlayerIndex    Pointer                       // Called when the player index is set
	Rumble            Pointer                       // Implements SDL_RumbleJoystick()
	RumbleTriggers    Pointer                       // Implements SDL_RumbleJoystickTriggers()
	SetLED            Pointer                       // Implements SDL_SetJoystickLED()
	SendEffect        Pointer                       // Implements SDL_SendJoystickEffect()
	SetSensorsEnabled Pointer                       // Implements SDL_SetGamepadSensorEnabled()
}

type virtualJoystickDesc struct {
	Version           uint32                       // the version of this interface
	Type              uint16                       // `SDL_JoystickType`
	padding           uint16                       // unused
	VendorId          uint16                       // the USB vendor ID of this joystick
	ProductId         uint16                       // the USB product ID of this joystick
	Naxes             uint16                       // the number of axes on this joystick
	Nbuttons          uint16                       // the number of buttons on this joystick
	Nballs            uint16                       // the number of balls on this joystick
	Nhats             uint16                       // the number of hats on this joystick
	Ntouchpads        uint16                       // the number of touchpads on this joystick, requires `touchpads` to point at valid descriptions
	Nsensors          uint16                       // the number of sensors on this joystick, requires `sensors` to point at valid descriptions
	padding2          [2]uint16                    // unused
	ButtonMask        uint32                       // A mask of which buttons are valid for this controller e.g. (1 << SDL_GAMEPAD_BUTTON_SOUTH)
	AxisMask          uint32                       // A mask of which axes are valid for this controller e.g. (1 << SDL_GAMEPAD_AXIS_LEFTX)
	Name              *byte                        // the name of the joystick
	Touchpads         *VirtualJoystickTouchpadDesc // A pointer to an array of touchpad descriptions, required if `ntouchpads` is > 0
	Sensors           *VirtualJoystickSensorDesc   // A pointer to an array of sensor descriptions, required if `nsensors` is > 0
	Userdata          Pointer                      // User data pointer passed to callbacks
	Update            Pointer                      // Called when the joystick state should be updated
	SetPlayerIndex    Pointer                      // Called when the player index is set
	Rumble            Pointer                      // Implements SDL_RumbleJoystick()
	RumbleTriggers    Pointer                      // Implements SDL_RumbleJoystickTriggers()
	SetLED            Pointer                      // Implements SDL_SetJoystickLED()
	SendEffect        Pointer                      // Implements SDL_SendJoystickEffect()
	SetSensorsEnabled Pointer                      // Implements SDL_SetGamepadSensorEnabled()
	Cleanup           Pointer                      // Cleans up the userdata when the joystick is detached
}

func (desc *VirtualJoystickDesc) as() *virtualJoystickDesc {
	if desc == nil {
		return nil
	}

	var touchpads *VirtualJoystickTouchpadDesc
	var sensors *VirtualJoystickSensorDesc
	var nilPointer internal.Pointer
	if len(desc.Touchpads) > 0 {
		touchpads = unsafe.SliceData(desc.Touchpads)
	}
	if len(desc.Sensors) > 0 {
		sensors = unsafe.SliceData(desc.Sensors)
	}
	return &virtualJoystickDesc{
		Version:           desc.Version,
		Type:              uint16(desc.Type),
		padding:           0,
		VendorId:          desc.VendorId,
		ProductId:         desc.ProductId,
		Naxes:             desc.Naxes,
		Nbuttons:          desc.Nbuttons,
		Nballs:            desc.Nballs,
		Nhats:             desc.Nhats,
		Ntouchpads:        uint16(len(desc.Touchpads)),
		Nsensors:          uint16(len(desc.Sensors)),
		padding2:          [2]uint16{},
		ButtonMask:        desc.ButtonMask,
		AxisMask:          desc.AxisMask,
		Name:              internal.StringToNullablePtr(desc.Name),
		Touchpads:         touchpads,
		Sensors:           sensors,
		Userdata:          nilPointer,
		Update:            desc.Update,
		SetPlayerIndex:    desc.SetPlayerIndex,
		Rumble:            desc.Rumble,
		RumbleTriggers:    desc.RumbleTriggers,
		SetLED:            desc.SetLED,
		SendEffect:        desc.SendEffect,
		SetSensorsEnabled: desc.SetSensorsEnabled,
		Cleanup:           nilPointer,
	}
}

// SDL_GPURenderStateCreateInfo - A structure specifying the parameters of a GPU render state.
// (https://wiki.libsdl.org/SDL3/SDL_GPURenderStateCreateInfo)
type GPURenderStateCreateInfo struct {
	FragmentShader  *GPUShader                 // The fragment shader to use when this render state is active
	SamplerBindings []GPUTextureSamplerBinding // Additional fragment samplers to bind when this render state is active
	StorageTextures []*GPUTexture              // Storage textures to bind when this render state is active
	StorageBuffers  []*GPUBuffer               // Storage buffers to bind when this render state is active
	Props           PropertiesID               // A properties ID for extensions. Should be 0 if no extensions are needed.
}

type gpuRenderStateCreateInfo struct {
	FragmentShader     *GPUShader                // The fragment shader to use when this render state is active
	NumSamplerBindings int32                     // The number of additional fragment samplers to bind when this render state is active
	SamplerBindings    *GPUTextureSamplerBinding // Additional fragment samplers to bind when this render state is active
	NumStorageTextures int32                     // The number of storage textures to bind when this render state is active
	StorageTextures    **GPUTexture              // Storage textures to bind when this render state is active
	NumStorageBuffers  int32                     // The number of storage buffers to bind when this render state is active
	StorageBuffers     **GPUBuffer               // Storage buffers to bind when this render state is active
	Props              PropertiesID              // A properties ID for extensions. Should be 0 if no extensions are needed.
}

func (info *GPURenderStateCreateInfo) as() *gpuRenderStateCreateInfo {
	if info == nil {
		return nil
	}
	var samplerBindings *GPUTextureSamplerBinding
	var storageTextures **GPUTexture
	var storageBuffers **GPUBuffer
	if len(info.SamplerBindings) > 0 {
		samplerBindings = unsafe.SliceData(info.SamplerBindings)
	}
	if len(info.StorageTextures) > 0 {
		storageTextures = unsafe.SliceData(info.StorageTextures)
	}
	if len(info.StorageBuffers) > 0 {
		storageBuffers = unsafe.SliceData(info.StorageBuffers)
	}
	return &gpuRenderStateCreateInfo{
		FragmentShader:     info.FragmentShader,
		NumSamplerBindings: int32(len(info.SamplerBindings)),
		SamplerBindings:    samplerBindings,
		NumStorageTextures: int32(len(info.StorageTextures)),
		StorageTextures:    storageTextures,
		NumStorageBuffers:  int32(len(info.StorageBuffers)),
		StorageBuffers:     storageBuffers,
		Props:              info.Props,
	}
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

// SDL_AudioStreamDataCompleteCallback - A callback that fires for completed [SDL_PutAudioStreamDataNoCopy](SDL_PutAudioStreamDataNoCopy)() data.
// (https://wiki.libsdl.org/SDL3/SDL_AudioStreamDataCompleteCallback)
type AudioStreamDataCompleteCallback uintptr

// SDL_MouseMotionTransformCallback - A callback used to transform mouse motion delta from raw values.
// (https://wiki.libsdl.org/SDL3/SDL_MouseMotionTransformCallback)
type MouseMotionTransformCallback uintptr

// SDL_TrayCallback - A callback that is invoked when a tray entry is selected.
// (https://wiki.libsdl.org/SDL3/SDL_TrayCallback)
type TrayCallback uintptr

// Vulkan Types

// VkInstance is a pointer to a Vulkan_Instance
type VkInstance uintptr

// VkPhysicalDevice is a pointer to a Vulkan_PhysicalDevice
type VkPhysicalDevice uintptr

// VkSurfaceKHR is a pointer to a Vulkan_SurfaceKHR
type VkSurfaceKHR uintptr

// VkAllocationCallbacks is a placeholder type for Vulkan_AllocationCallbacks
type VkAllocationCallbacks struct{}

// VkAllocationCallbacksPtr is a pointer to for VkAllocationCallbacks
type VkAllocationCallbacksPtr uintptr
