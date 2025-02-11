package sdl

// Code generated by cmd/ffi2go. DO NOT EDIT.

type AsyncIO struct{}

type AsyncIOOutcome struct {
	Asyncio          *AsyncIO
	Type             AsyncIOTaskType
	Result           AsyncIOResult
	Buffer           uintptr
	Offset           uint64
	BytesRequested   uint64
	BytesTransferred uint64
	Userdata         uintptr
}

type AsyncIOQueue struct{}

type AtomicInt struct {
	Value int
}

type AtomicU32 struct {
	Value uint32
}

type Thread struct{}

type Mutex struct{}

type RWLock struct{}

type Semaphore struct{}

type Condition struct{}

type InitState struct {
	Status   AtomicInt
	Thread   ThreadID
	Reserved uintptr
}

type IOStreamInterface struct {
	Version uint32
	Size    uintptr
	Seek    uintptr
	Read    uintptr
	Write   uintptr
	Flush   uintptr
	Close   uintptr
}

type IOStream struct{}

type AudioSpec struct {
	Format   AudioFormat
	Channels int
	Freq     int
}

type AudioStream struct{}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

type FColor struct {
	R float32
	G float32
	B float32
	A float32
}

type Palette struct {
	Ncolors  int
	Colors   *Color
	Version  uint32
	Refcount int
}

type PixelFormatDetails struct {
	Format        PixelFormat
	BitsPerPixel  uint8
	BytesPerPixel uint8
	Padding       [2]uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rbits         uint8
	Gbits         uint8
	Bbits         uint8
	Abits         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
}

type Point struct {
	X int
	Y int
}

type FPoint struct {
	X float32
	Y float32
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

type FRect struct {
	X float32
	Y float32
	W float32
	H float32
}

type Surface struct {
	Flags    SurfaceFlags
	Format   PixelFormat
	W        int
	H        int
	Pitch    int
	Pixels   uintptr
	Refcount int
	Reserved uintptr
}

type Camera struct{}

type CameraSpec struct {
	Format               PixelFormat
	Colorspace           Colorspace
	Width                int
	Height               int
	FramerateNumerator   int
	FramerateDenominator int
}

type DisplayModeData struct{}

type DisplayMode struct {
	DisplayID              DisplayID
	Format                 PixelFormat
	W                      int
	H                      int
	PixelDensity           float32
	RefreshRate            float32
	RefreshRateNumerator   int
	RefreshRateDenominator int
	Internal               *DisplayModeData
}

type Window struct{}

type GLContextState struct{}

type DialogFileFilter struct {
	Name    string
	Pattern string
}

type Sensor struct{}

type Joystick struct{}

type VirtualJoystickTouchpadDesc struct {
	Nfingers uint16
	Padding  [3]uint16
}

type VirtualJoystickSensorDesc struct {
	Type SensorType
	Rate float32
}

type VirtualJoystickDesc struct {
	Version           uint32
	Type              uint16
	Padding           uint16
	VendorId          uint16
	ProductId         uint16
	Naxes             uint16
	Nbuttons          uint16
	Nballs            uint16
	Nhats             uint16
	Ntouchpads        uint16
	Nsensors          uint16
	Padding2          [2]uint16
	ButtonMask        uint32
	AxisMask          uint32
	Name              string
	Touchpads         *VirtualJoystickTouchpadDesc
	Sensors           *VirtualJoystickSensorDesc
	Userdata          uintptr
	Update            uintptr
	SetPlayerIndex    uintptr
	Rumble            uintptr
	RumbleTriggers    uintptr
	SetLED            uintptr
	SendEffect        uintptr
	SetSensorsEnabled uintptr
	Cleanup           uintptr
}

type Gamepad struct{}

type Cursor struct{}

type Finger struct {
	Id       FingerID
	X        float32
	Y        float32
	Pressure float32
}

type CommonEvent struct {
	Type      uint32
	Reserved  uint32
	Timestamp uint64
}

type DisplayEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	DisplayID DisplayID
	Data1     int32
	Data2     int32
}

type WindowEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Data1     int32
	Data2     int32
}

type KeyboardDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     KeyboardID
}

type KeyboardEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     KeyboardID
	Scancode  Scancode
	Key       Keycode
	Mod       Keymod
	Raw       uint16
	Down      bool
	Repeat    bool
}

type TextEditingEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Text      string
	Start     int32
	Length    int32
}

type TextEditingCandidatesEvent struct {
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

type TextInputEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Text      string
}

type MouseDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     MouseID
}

type MouseMotionEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     MouseID
	State     MouseButtonFlags
	X         float32
	Y         float32
	Xrel      float32
	Yrel      float32
}

type MouseButtonEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     MouseID
	Button    uint8
	Down      bool
	Clicks    uint8
	Padding   uint8
	X         float32
	Y         float32
}

type MouseWheelEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     MouseID
	X         float32
	Y         float32
	Direction MouseWheelDirection
	MouseX    float32
	MouseY    float32
}

type JoyAxisEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type JoyBallEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Ball      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Xrel      int16
	Yrel      int16
}

type JoyHatEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Hat       uint8
	Value     uint8
	Padding1  uint8
	Padding2  uint8
}

type JoyButtonEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Button    uint8
	Down      bool
	Padding1  uint8
	Padding2  uint8
}

type JoyDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
}

type JoyBatteryEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	State     PowerState
	Percent   int
}

type GamepadAxisEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Axis      uint8
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
	Value     int16
	Padding4  uint16
}

type GamepadButtonEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Button    uint8
	Down      bool
	Padding1  uint8
	Padding2  uint8
}

type GamepadDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
}

type GamepadTouchpadEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     JoystickID
	Touchpad  int32
	Finger    int32
	X         float32
	Y         float32
	Pressure  float32
}

type GamepadSensorEvent struct {
	Type            EventType
	Reserved        uint32
	Timestamp       uint64
	Which           JoystickID
	Sensor          int32
	Data            [3]float32
	SensorTimestamp uint64
}

type AudioDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     AudioDeviceID
	Recording bool
	Padding1  uint8
	Padding2  uint8
	Padding3  uint8
}

type CameraDeviceEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	Which     CameraID
}

type RenderEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
}

type TouchFingerEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	TouchID   TouchID
	FingerID  FingerID
	X         float32
	Y         float32
	Dx        float32
	Dy        float32
	Pressure  float32
	WindowID  WindowID
}

type PenProximityEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     PenID
}

type PenMotionEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     PenID
	PenState  PenInputFlags
	X         float32
	Y         float32
}

type PenTouchEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     PenID
	PenState  PenInputFlags
	X         float32
	Y         float32
	Eraser    bool
	Down      bool
}

type PenButtonEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     PenID
	PenState  PenInputFlags
	X         float32
	Y         float32
	Button    uint8
	Down      bool
}

type PenAxisEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Which     PenID
	PenState  PenInputFlags
	X         float32
	Y         float32
	Axis      PenAxis
	Value     float32
}

type DropEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	X         float32
	Y         float32
	Source    string
	Data      string
}

type ClipboardEvent struct {
	Type         EventType
	Reserved     uint32
	Timestamp    uint64
	Owner        bool
	NumMimeTypes int32
	MimeTypes    *string
}

type SensorEvent struct {
	Type            EventType
	Reserved        uint32
	Timestamp       uint64
	Which           SensorID
	Data            [6]float32
	SensorTimestamp uint64
}

type QuitEvent struct {
	Type      EventType
	Reserved  uint32
	Timestamp uint64
}

type UserEvent struct {
	Type      uint32
	Reserved  uint32
	Timestamp uint64
	WindowID  WindowID
	Code      int32
	Data1     uintptr
	Data2     uintptr
}

type PathInfo struct {
	Type       PathType
	Size       uint64
	CreateTime Time
	ModifyTime Time
	AccessTime Time
}

type GPUDevice struct{}

type GPUBuffer struct{}

type GPUTransferBuffer struct{}

type GPUTexture struct{}

type GPUSampler struct{}

type GPUShader struct{}

type GPUComputePipeline struct{}

type GPUGraphicsPipeline struct{}

type GPUCommandBuffer struct{}

type GPURenderPass struct{}

type GPUComputePass struct{}

type GPUCopyPass struct{}

type GPUFence struct{}

type GPUViewport struct {
	X        float32
	Y        float32
	W        float32
	H        float32
	MinDepth float32
	MaxDepth float32
}

type GPUTextureTransferInfo struct {
	TransferBuffer *GPUTransferBuffer
	Offset         uint32
	PixelsPerRow   uint32
	RowsPerLayer   uint32
}

type GPUTransferBufferLocation struct {
	TransferBuffer *GPUTransferBuffer
	Offset         uint32
}

type GPUTextureLocation struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	X        uint32
	Y        uint32
	Z        uint32
}

type GPUTextureRegion struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	X        uint32
	Y        uint32
	Z        uint32
	W        uint32
	H        uint32
	D        uint32
}

type GPUBlitRegion struct {
	Texture           *GPUTexture
	MipLevel          uint32
	LayerOrDepthPlane uint32
	X                 uint32
	Y                 uint32
	W                 uint32
	H                 uint32
}

type GPUBufferLocation struct {
	Buffer *GPUBuffer
	Offset uint32
}

type GPUBufferRegion struct {
	Buffer *GPUBuffer
	Offset uint32
	Size   uint32
}

type GPUSamplerCreateInfo struct {
	MinFilter        GPUFilter
	MagFilter        GPUFilter
	MipmapMode       GPUSamplerMipmapMode
	AddressModeU     GPUSamplerAddressMode
	AddressModeV     GPUSamplerAddressMode
	AddressModeW     GPUSamplerAddressMode
	MipLodBias       float32
	MaxAnisotropy    float32
	CompareOp        GPUCompareOp
	MinLod           float32
	MaxLod           float32
	EnableAnisotropy bool
	EnableCompare    bool
	Padding1         uint8
	Padding2         uint8
	Props            PropertiesID
}

type GPUVertexBufferDescription struct {
	Slot             uint32
	Pitch            uint32
	InputRate        GPUVertexInputRate
	InstanceStepRate uint32
}

type GPUVertexAttribute struct {
	Location   uint32
	BufferSlot uint32
	Format     GPUVertexElementFormat
	Offset     uint32
}

type GPUVertexInputState struct {
	VertexBufferDescriptions *GPUVertexBufferDescription
	NumVertexBuffers         uint32
	VertexAttributes         *GPUVertexAttribute
	NumVertexAttributes      uint32
}

type GPUStencilOpState struct {
	FailOp      GPUStencilOp
	PassOp      GPUStencilOp
	DepthFailOp GPUStencilOp
	CompareOp   GPUCompareOp
}

type GPUColorTargetBlendState struct {
	SrcColorBlendfactor  GPUBlendFactor
	DstColorBlendfactor  GPUBlendFactor
	ColorBlendOp         GPUBlendOp
	SrcAlphaBlendfactor  GPUBlendFactor
	DstAlphaBlendfactor  GPUBlendFactor
	AlphaBlendOp         GPUBlendOp
	ColorWriteMask       GPUColorComponentFlags
	EnableBlend          bool
	EnableColorWriteMask bool
	Padding1             uint8
	Padding2             uint8
}

type GPUShaderCreateInfo struct {
	CodeSize           uintptr
	Code               *uint8
	Entrypoint         string
	Format             GPUShaderFormat
	Stage              GPUShaderStage
	NumSamplers        uint32
	NumStorageTextures uint32
	NumStorageBuffers  uint32
	NumUniformBuffers  uint32
	Props              PropertiesID
}

type GPUTextureCreateInfo struct {
	Type              GPUTextureType
	Format            GPUTextureFormat
	Usage             GPUTextureUsageFlags
	Width             uint32
	Height            uint32
	LayerCountOrDepth uint32
	NumLevels         uint32
	SampleCount       GPUSampleCount
	Props             PropertiesID
}

type GPUBufferCreateInfo struct {
	Usage GPUBufferUsageFlags
	Size  uint32
	Props PropertiesID
}

type GPUTransferBufferCreateInfo struct {
	Usage GPUTransferBufferUsage
	Size  uint32
	Props PropertiesID
}

type GPURasterizerState struct {
	FillMode                GPUFillMode
	CullMode                GPUCullMode
	FrontFace               GPUFrontFace
	DepthBiasConstantFactor float32
	DepthBiasClamp          float32
	DepthBiasSlopeFactor    float32
	EnableDepthBias         bool
	EnableDepthClip         bool
	Padding1                uint8
	Padding2                uint8
}

type GPUMultisampleState struct {
	SampleCount GPUSampleCount
	SampleMask  uint32
	EnableMask  bool
	Padding1    uint8
	Padding2    uint8
	Padding3    uint8
}

type GPUDepthStencilState struct {
	CompareOp         GPUCompareOp
	BackStencilState  GPUStencilOpState
	FrontStencilState GPUStencilOpState
	CompareMask       uint8
	WriteMask         uint8
	EnableDepthTest   bool
	EnableDepthWrite  bool
	EnableStencilTest bool
	Padding1          uint8
	Padding2          uint8
	Padding3          uint8
}

type GPUColorTargetDescription struct {
	Format     GPUTextureFormat
	BlendState GPUColorTargetBlendState
}

type GPUGraphicsPipelineTargetInfo struct {
	ColorTargetDescriptions *GPUColorTargetDescription
	NumColorTargets         uint32
	DepthStencilFormat      GPUTextureFormat
	HasDepthStencilTarget   bool
	Padding1                uint8
	Padding2                uint8
	Padding3                uint8
}

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

type GPUComputePipelineCreateInfo struct {
	CodeSize                    uintptr
	Code                        *uint8
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

type GPUColorTargetInfo struct {
	Texture             *GPUTexture
	MipLevel            uint32
	LayerOrDepthPlane   uint32
	ClearColor          FColor
	LoadOp              GPULoadOp
	StoreOp             GPUStoreOp
	ResolveTexture      *GPUTexture
	ResolveMipLevel     uint32
	ResolveLayer        uint32
	Cycle               bool
	CycleResolveTexture bool
	Padding1            uint8
	Padding2            uint8
}

type GPUDepthStencilTargetInfo struct {
	Texture        *GPUTexture
	ClearDepth     float32
	LoadOp         GPULoadOp
	StoreOp        GPUStoreOp
	StencilLoadOp  GPULoadOp
	StencilStoreOp GPUStoreOp
	Cycle          bool
	ClearStencil   uint8
	Padding1       uint8
	Padding2       uint8
}

type GPUBlitInfo struct {
	Source      GPUBlitRegion
	Destination GPUBlitRegion
	LoadOp      GPULoadOp
	ClearColor  FColor
	FlipMode    FlipMode
	Filter      GPUFilter
	Cycle       bool
	Padding1    uint8
	Padding2    uint8
	Padding3    uint8
}

type GPUBufferBinding struct {
	Buffer *GPUBuffer
	Offset uint32
}

type GPUTextureSamplerBinding struct {
	Texture *GPUTexture
	Sampler *GPUSampler
}

type GPUStorageBufferReadWriteBinding struct {
	Buffer   *GPUBuffer
	Cycle    bool
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
}

type GPUStorageTextureReadWriteBinding struct {
	Texture  *GPUTexture
	MipLevel uint32
	Layer    uint32
	Cycle    bool
	Padding1 uint8
	Padding2 uint8
	Padding3 uint8
}

type Haptic struct{}

type HapticDirection struct {
	Type uint8
	Dir  [3]int32
}

type HapticConstant struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Level        int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticPeriodic struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Period       uint16
	Magnitude    int16
	Offset       int16
	Phase        uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticCondition struct {
	Type       uint16
	Direction  HapticDirection
	Length     uint32
	Delay      uint16
	Button     uint16
	Interval   uint16
	RightSat   [3]uint16
	LeftSat    [3]uint16
	RightCoeff [3]int16
	LeftCoeff  [3]int16
	Deadband   [3]uint16
	Center     [3]int16
}

type HapticRamp struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Start        int16
	End          int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticLeftRight struct {
	Type           uint16
	Length         uint32
	LargeMagnitude uint16
	SmallMagnitude uint16
}

type HapticCustom struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Channels     uint8
	Period       uint16
	Samples      uint16
	Data         *uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type SharedObject struct{}

type Locale struct {
	Language string
	Country  string
}

type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int
	Text     string
}

type MessageBoxColor struct {
	R uint8
	G uint8
	B uint8
}

type MessageBoxColorScheme struct {
	Colors [5]MessageBoxColor
}

type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window
	Title       string
	Message     string
	Numbuttons  int
	Buttons     *MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

type Process struct{}

type Vertex struct {
	Position FPoint
	Color    FColor
	TexCoord FPoint
}

type Renderer struct{}

type Texture struct {
	Format   PixelFormat
	W        int
	H        int
	Refcount int
}

type StorageInterface struct {
	Version        uint32
	Close          uintptr
	Ready          uintptr
	Enumerate      uintptr
	Info           uintptr
	ReadFile       uintptr
	WriteFile      uintptr
	Mkdir          uintptr
	Remove         uintptr
	Rename         uintptr
	Copy           uintptr
	SpaceRemaining uintptr
}

type Storage struct{}

type DateTime struct {
	Year       int
	Month      int
	Day        int
	Hour       int
	Minute     int
	Second     int
	Nanosecond int
	DayOfWeek  int
	UtcOffset  int
}
