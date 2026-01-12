package sdl

import (
	"runtime"
	"slices"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// Init

// SDL_Init - Initialize the SDL library.
// (https://wiki.libsdl.org/SDL3/SDL_Init)
func Init(flags InitFlags) error {
	if !iInit(flags) {
		return internal.LastErr()
	}

	return nil
}

// SDL_InitSubSystem - Compatibility function to initialize the SDL library.
// (https://wiki.libsdl.org/SDL3/SDL_InitSubSystem)
func InitSubSystem(flags InitFlags) error {
	if !iInitSubSystem(flags) {
		return internal.LastErr()
	}

	return nil
}

// SDL_QuitSubSystem - Shut down specific SDL subsystems.
// (https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem)
func QuitSubSystem(flags InitFlags) {
	iQuitSubSystem(flags)
}

// SDL_WasInit - Get a mask of the specified subsystems which are currently initialized.
// (https://wiki.libsdl.org/SDL3/SDL_WasInit)
func WasInit(flags InitFlags) InitFlags {
	return iWasInit(flags)
}

// SDL_Quit - Clean up all initialized subsystems.
// (https://wiki.libsdl.org/SDL3/SDL_Quit)
func Quit() {
	iQuit()
}

// SDL_IsMainThread - Return whether this is the main thread.
// (https://wiki.libsdl.org/SDL3/SDL_IsMainThread)
func IsMainThread() bool {
	return iIsMainThread()
}

// SDL_SetAppMetadata - Specify basic metadata about your app.
// (https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata)
func SetAppMetadata(appName, appVersion, appIdentifier string) error {
	if !iSetAppMetadata(appName, appVersion, appIdentifier) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetAppMetadataProperty - Specify metadata about your app through a set of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty)
func SetAppMetadataProperty(name, value string) error {
	if !iSetAppMetadataProperty(name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetAppMetadataProperty - Get metadata about your app.
// (https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty)
func GetAppMetadataProperty(name string) string {
	return iGetAppMetadataProperty(name)
}

// Hints

// SDL_SetHintWithPriority - Set a hint with a specific priority.
// (https://wiki.libsdl.org/SDL3/SDL_SetHintWithPriority)
func SetHintWithPriority(name, value string, priority HintPriority) error {
	if !iSetHintWithPriority(name, value, priority) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetHint - Set a hint with normal priority.
// (https://wiki.libsdl.org/SDL3/SDL_SetHint)
func SetHint(name, value string) error {
	if !iSetHint(name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ResetHint - Reset a hint to the default value.
// (https://wiki.libsdl.org/SDL3/SDL_ResetHint)
func ResetHint(name string) error {
	if !iResetHint(name) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ResetHints - Reset all hints to the default values.
// (https://wiki.libsdl.org/SDL3/SDL_ResetHints)
func ResetHints() {
	iResetHints()
}

// SDL_GetHint - Get the value of a hint.
// (https://wiki.libsdl.org/SDL3/SDL_GetHint)
func GetHint(name string) string {
	return iGetHint(name)
}

// SDL_GetHintBoolean - Get the boolean value of a hint variable.
// (https://wiki.libsdl.org/SDL3/SDL_GetHintBoolean)
func GetHintBoolean(name string, defaultValue bool) bool {
	return iGetHintBoolean(name, defaultValue)
}

// TODO: AddHintCallback
// TODO: RemoveHintCallback

// Error

// SDL_OutOfMemory - Set an error indicating that memory allocation failed.
// (https://wiki.libsdl.org/SDL3/SDL_OutOfMemory)
func OutOfMemory() bool {
	return iOutOfMemory()
}

// Properties

// SDL_GetGlobalProperties - Get the global SDL properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties)
func GetGlobalProperties() (PropertiesID, error) {
	properties := iGetGlobalProperties()
	if properties == 0 {
		return 0, internal.LastErr()
	}

	return properties, nil
}

// SDL_CreateProperties - Create a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateProperties)
func CreateProperties() (PropertiesID, error) {
	properties := iCreateProperties()
	if properties == 0 {
		return 0, internal.LastErr()
	}

	return properties, nil
}

// SDL_CopyProperties - Copy a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_CopyProperties)
func CopyProperties(src, dst PropertiesID) error {
	if !iCopyProperties(src, dst) {
		return internal.LastErr()
	}

	return nil
}

// Log

// TODO: is there a need?

// Events

// SDL_PumpEvents - Pump the event loop, gathering events from the input devices.
// (https://wiki.libsdl.org/SDL3/SDL_PumpEvents)
func PumpEvents() {
	iPumpEvents()
}

// TODO: PeepEvents()

// NumEvents returns the number of messages from the event queue, it is similar to
// calling SDL_PeepEvents(nil, 0, SDL_PEEKEVENT, SDL_EVENT_FIRST, SDL_EVENT_LAST).
// You may have to call PumpEvents() before calling this function.
// https://wiki.libsdl.org/SDL3/SDL_PeepEvents
func NumEvents() (int32, error) {
	count := iPeepEvents(nil, 0, PEEKEVENT, uint32(EVENT_FIRST), uint32(EVENT_LAST))
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// AppendEvents adds the messages from the event queue to the provided slice and returns it.
// It is useful to avoid re-allocation, assuming events already has the required capacity, but
// you may also pass nil as an argument.
// A good pattern would be to zero the same slice each update: events = events[:0] before
// calling PumpEvents() and then AppendEvents().
// This function calls SDL_PeepEvents with the action GETEVENT, which means that it will
// remove the messages from the event queue, just like PollEvent.
// This helper exists to minimize calls to the library, so that there is
// only a fixed cost of purego/Cgo syscall overhead as opposed to calling PollEvent in a loop.
func AppendEvents(events []Event) ([]Event, error) {
	count, err := NumEvents()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return events, nil
	}
	initialLength := len(events)

	available := cap(events) - len(events)
	if available < int(count) {
		events = slices.Grow(events, int(count)-available)
	}
	events = events[len(events) : len(events)+int(count)]

	num := iPeepEvents(unsafe.SliceData(events[initialLength:]), count, GETEVENT, uint32(EVENT_FIRST), uint32(EVENT_LAST))
	if num < 0 {
		return nil, internal.LastErr()
	}

	return events, nil
}

// SDL_HasEvent - Check for the existence of a certain event type in the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_HasEvent)
func HasEvent(typ EventType) bool {
	return iHasEvent(uint32(typ))
}

// SDL_HasEvents - Check for the existence of certain event types in the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_HasEvents)
func HasEvents(minType, maxType EventType) bool {
	return iHasEvents(uint32(minType), uint32(maxType))
}

// SDL_FlushEvent - Clear events of a specific type from the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_FlushEvent)
func FlushEvent(typ EventType) {
	iFlushEvent(uint32(typ))
}

// SDL_FlushEvents - Clear events of a range of types from the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_FlushEvents)
func FlushEvents(minType, maxType EventType) {
	iFlushEvents(uint32(minType), uint32(maxType))
}

// SDL_PollEvent - Poll for currently pending events.
// (https://wiki.libsdl.org/SDL3/SDL_PollEvent)
func PollEvent(event *Event) bool {
	return iPollEvent(event)
}

// SDL_WaitEvent - Wait indefinitely for the next available event.
// (https://wiki.libsdl.org/SDL3/SDL_WaitEvent)
func WaitEvent(event *Event) error {
	if !iWaitEvent(event) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WaitEventTimeout - Wait until the specified timeout (in milliseconds) for the next available event.
// (https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout)
func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	return iWaitEventTimeout(event, timeoutMS)
}

// SDL_PushEvent - Add an event to the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_PushEvent)
func PushEvent(event *Event) error {
	if !iPushEvent(event) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetEventFilter - Set up a filter to process all events before they are added to the internal event queue.
// (https://wiki.libsdl.org/SDL3/SDL_SetEventFilter)
func SetEventFilter(filter EventFilter) {
	iSetEventFilter(filter, 0)
}

// SDL_GetEventFilter - Query the current event filter.
// (https://wiki.libsdl.org/SDL3/SDL_GetEventFilter)
func GetEventFilter() EventFilter {
	var filter EventFilter
	var userData uintptr

	if !iGetEventFilter(&filter, &userData) {
		return 0
	}

	return filter
}

// SDL_AddEventWatch - Add a callback to be triggered when an event is added to the event queue.
// (https://wiki.libsdl.org/SDL3/SDL_AddEventWatch)
func AddEventWatch(filter EventFilter) error {
	if !iAddEventWatch(filter, 0) {
		internal.LastErr()
	}

	return nil
}

// SDL_RemoveEventWatch - Remove an event watch callback added with [SDL_AddEventWatch](SDL_AddEventWatch)().
// (https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch)
func RemoveEventWatch(filter EventFilter) {
	iRemoveEventWatch(filter, 0)
}

// SDL_FilterEvents - Run a specific filter function on the current event queue, removing any events for which the filter returns false.
// (https://wiki.libsdl.org/SDL3/SDL_FilterEvents)
func FilterEvents(filter EventFilter) {
	iFilterEvents(filter, 0)
}

// SDL_SetEventEnabled - Set the state of processing events by type.
// (https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled)
func SetEventEnabled(typ EventType, enabled bool) {
	iSetEventEnabled(uint32(typ), enabled)
}

// SDL_EventEnabled - Query the state of processing events by type.
// (https://wiki.libsdl.org/SDL3/SDL_EventEnabled)
func EventEnabled(typ EventType) bool {
	return iEventEnabled(uint32(typ))
}

// Timer

// TODO:

// Shared object

// TODO:

// Thread

// TODO:

// Mutex

// TODO:

// Atomic

// TODO:

// IOStream

// SDL_IOFromFile - Use this function to create a new [SDL_IOStream](SDL_IOStream) structure for reading from and/or writing to a named file.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromFile)
func IOFromFile(file, mode string) (*IOStream, error) {
	stream := iIOFromFile(file, mode)
	if stream == nil {
		return nil, internal.LastErr()
	}

	runtime.KeepAlive(file)
	runtime.KeepAlive(mode)

	return stream, nil
}

// SDL_IOFromConstMem - Use this function to prepare a read-only memory buffer for use with SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem)
// Note: This function is unsafe as it is required for `mem` not to be garbage collected while the IOStream is in use.
// Please use IOFromBytes or IOFromDynamicMem, unless you can guarantee that `mem` will be kept alive.
func IOFromConstMem(mem []byte) (*IOStream, error) {
	stream := iIOFromConstMem(
		uintptr(unsafe.Pointer(unsafe.SliceData(mem))),
		uintptr(len(mem)),
	)
	if stream == nil {
		return nil, internal.LastErr()
	}

	runtime.KeepAlive(mem)

	return stream, nil
}

// SDL_IOFromDynamicMem - Use this function to create an SDL_IOStream that is backed by dynamically allocated memory.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromDynamicMem)
func IOFromDynamicMem() *IOStream {
	return iIOFromDynamicMem()
}

// TODO:

// AsyncIO

// TODO:

// Main

// TODO: is this needed?

// Render

// SDL_GetNumRenderDrivers - Get the number of 2D rendering drivers available for the current display.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers)
func GetNumRenderDrivers() int {
	return int(iGetNumRenderDrivers())
}

// SDL_GetRenderDriver - Use this function to get the name of a built in 2D rendering driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver)
func GetRenderDriver(index int) string {
	return iGetRenderDriver(int32(index))
}

// SDL_CreateWindowAndRenderer - Create a window and default renderer.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer)
func CreateWindowAndRenderer(title string, width, height int, flags WindowFlags) (*Window, *Renderer, error) {
	var window *Window
	var renderer *Renderer

	if !iCreateWindowAndRenderer(title, int32(width), int32(height), flags, &window, &renderer) {
		return nil, nil, internal.LastErr()
	}

	return window, renderer, nil
}

// SDL_CreateRendererWithProperties - Create a 2D rendering context for a window, with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties)
func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	renderer := iCreateRendererWithProperties(props)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// Pixels

// SDL_GetPixelFormatForMasks - Convert a bpp value and RGBA masks to an enumerated pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks)
func GetPixelFormatForMasks(bpp int, rmask, gmask, bmask, amask uint32) PixelFormat {
	return iGetPixelFormatForMasks(int32(bpp), rmask, gmask, bmask, amask)
}

// SDL_CreatePalette - Create a palette structure with the specified number of color entries.
// (https://wiki.libsdl.org/SDL3/SDL_CreatePalette)
func CreatePalette(numColors int) (*Palette, error) {
	palette := iCreatePalette(int32(numColors))
	if palette == nil {
		return nil, internal.LastErr()
	}

	return palette, nil
}

// SDL_MapRGB - Map an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGB)
func MapRGB(format *PixelFormatDetails, palette *Palette, r, g, b byte) uint32 {
	return iMapRGB(format, palette, r, g, b)
}

// SDL_MapRGBA - Map an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGBA)
func MapRGBA(format *PixelFormatDetails, palette *Palette, r, g, b, a byte) uint32 {
	return iMapRGBA(format, palette, r, g, b, a)
}

// SDL_GetRGB - Get RGB values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGB)
func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b uint8) {
	iGetRGB(pixel, format, palette, &r, &g, &b)

	return r, g, b
}

// SDL_GetRGBA - Get RGBA values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGBA)
func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b, a uint8) {
	iGetRGBA(pixel, format, palette, &r, &g, &b, &a)

	return r, g, b, a
}

// Surface

// SDL_CreateSurface - Allocate a new surface with a specific pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurface)
func CreateSurface(width, height int, format PixelFormat) (*Surface, error) {
	surface := iCreateSurface(int32(width), int32(height), format)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// SDL_CreateSurfaceFrom - Allocate a new surface with a specific pixel format and existing pixel data.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom)
func CreateSurfaceFrom(width, height int, format PixelFormat, pixels []byte, pitch int) (*Surface, error) {
	surface := iCreateSurfaceFrom(int32(width), int32(height), format, uintptr(unsafe.Pointer(unsafe.SliceData(pixels))), int32(pitch))
	if surface == nil {
		return nil, internal.LastErr()
	}
	runtime.KeepAlive(pixels)

	return surface, nil
}

// SDL_LoadBMP_IO - Load a BMP image from a seekable SDL data stream.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO)
func LoadBMP_IO(src *IOStream, closeIO bool) (*Surface, error) {
	surface := iLoadBMP_IO(src, closeIO)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// SDL_LoadBMP - Load a BMP image from a file.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP)
func LoadBMP(file string) (*Surface, error) {
	surface := iLoadBMP(file)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TODO: ConvertPixels => void* data
// TODO: ConvertPixelsAndColorspace => ^
// TODO: PremultiplyAlpha => ^

// Blend mode

// SDL_ComposeCustomBlendMode - Compose a custom blend mode for renderers.
// (https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode)
func ComposeCustomBlendMode(srcFactor, dstFactor BlendFactor, colorOp BlendOperation, srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOp BlendOperation) BlendMode {
	return iComposeCustomBlendMode(srcFactor, dstFactor, colorOp, srcAlphaFactor, dstAlphaFactor, alphaOp)
}

// GPU

// SDL_GPUSupportsShaderFormats - Checks for GPU runtime support.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats)
func GPUSupportShaderFormats(formatFlags GPUShaderFormat, name string) bool {
	var namePtr *byte
	if name != "" {
		namePtr = internal.StringToPtr(name)
	}
	return iGPUSupportsShaderFormats(formatFlags, namePtr)
}

// SDL_CreateGPUDevice - Creates a GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice)
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) (*GPUDevice, error) {
	var namePtr *byte
	if name != "" {
		namePtr = internal.StringToPtr(name)
	}
	device := iCreateGPUDevice(formatFlags, debugMode, namePtr)
	if device == nil {
		return nil, internal.LastErr()
	}

	return device, nil
}

// SDL_CreateGPUDeviceWithProperties - Creates a GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUDeviceWithProperties)
func CreateGPUDeviceWithProperties(props PropertiesID) (*GPUDevice, error) {
	device := iCreateGPUDeviceWithProperties(props)
	if device == nil {
		return nil, internal.LastErr()
	}

	return device, nil
}

// SDL_GetNumGPUDrivers - Get the number of GPU drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGPUDrivers)
func NumGPUDrivers() int {
	return int(iGetNumGPUDrivers())
}

// SDL_GetGPUDriver - Get the name of a built in GPU driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUDriver)
func GetGPUDriver(index int32) string {
	return iGetGPUDriver(index)
}

// Video

// SDL_GetNumVideoDrivers - Get the number of video drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers)
func GetNumVideoDrivers() int {
	return int(iGetNumVideoDrivers())
}

// SDL_GetVideoDriver - Get the name of a built in video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return iGetVideoDriver(int32(index))
}

// SDL_GetCurrentVideoDriver - Get the name of the currently initialized video driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver)
func GetCurrentVideoDriver() string {
	return iGetCurrentVideoDriver()
}

// SDL_GetSystemTheme - Get the current system theme.
// (https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme)
func GetSystemTheme() SystemTheme {
	return iGetSystemTheme()
}

// SDL_GetDisplays - Get a list of currently connected displays.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplays)
func GetDisplays() ([]DisplayID, error) {
	var count int32

	ptr := iGetDisplays(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[DisplayID](ptr, int(count)), nil
}

// SDL_GetPrimaryDisplay - Return the primary display.
// (https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay)
func GetPrimaryDisplay() DisplayID {
	return iGetPrimaryDisplay()
}

// SDL_GetDisplayForPoint - Get the display containing a point.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint)
func GetDisplayForPoint(point *Point) DisplayID {
	return iGetDisplayForPoint(point)
}

// SDL_GetDisplayForRect - Get the display primarily containing a rect.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect)
func GetDisplayForRect(rect *Rect) DisplayID {
	return iGetDisplayForRect(rect)
}

// SDL_GetDisplayForWindow - Get the display associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow)
func GetDisplayForWindow(window *Window) DisplayID {
	return iGetDisplayForWindow(window)
}

// SDL_GetWindows - Get a list of valid windows.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindows)
func GetWindows() ([]*Window, error) {
	var count int32

	ptr := iGetWindows(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Window](ptr, int(count)), nil
}

// SDL_CreateWindow - Create a window with the specified dimensions and flags.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindow)
func CreateWindow(title string, width, height int, flags WindowFlags) (*Window, error) {
	window := iCreateWindow(title, int32(width), int32(height), flags)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_CreateWindowWithProperties - Create a window with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties)
func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	window := iCreateWindowWithProperties(props)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_GetGrabbedWindow - Get the window that currently has an input grab enabled.
// (https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow)
func GetGrabbedWindow() *Window {
	return iGetGrabbedWindow()
}

// SDL_ScreenSaverEnabled - Check whether the screensaver is currently enabled.
// (https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled)
func ScreenSaverEnabled() bool {
	return iScreenSaverEnabled()
}

// SDL_EnableScreenSaver - Allow the screen to be blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver)
func EnableScreenSaver() error {
	if !iEnableScreenSaver() {
		return internal.LastErr()
	}

	return nil
}

// SDL_DisableScreenSaver - Prevent the screen from being blanked by a screen saver.
// (https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver)
func DisableScreenSaver() error {
	if !iDisableScreenSaver() {
		return internal.LastErr()
	}

	return nil
}

// SDL_GL_LoadLibrary - Dynamically load an OpenGL library.
// (https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary)
func GL_LoadLibrary(path string) error {
	if !iGL_LoadLibrary(path) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GL_GetProcAddress - Get an OpenGL function by name.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress)
func GL_GetProcAddress(proc string) FunctionPointer {
	return iGL_GetProcAddress(proc)
}

// SDL_EGL_GetProcAddress - Get an EGL library function by name.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress)
func EGL_GetProcAddress(proc string) FunctionPointer {
	return iEGL_GetProcAddress(proc)
}

// SDL_GL_ExtensionSupported - Check if an OpenGL extension is supported for the current context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported)
func GL_ExtensionSupported(extension string) bool {
	return iGL_ExtensionSupported(extension)
}

// SDL_GL_UnloadLibrary - Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
// (https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary)
func GL_UnloadLibrary() {
	iGL_UnloadLibrary()
}

// SDL_GL_CreateContext - Create an OpenGL context for an OpenGL window, and make it current.
// (https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext)
func GL_CreateContext(window *Window) (GLContext, error) {
	ctx := iGL_CreateContext(window)
	if ctx == nil {
		return nil, internal.LastErr()
	}

	return ctx, nil
}

// SDL_GL_MakeCurrent - Set up an OpenGL context for rendering into an OpenGL window.
// (https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent)
func GL_MakeCurrent(window *Window, ctx GLContext) error {
	if !iGL_MakeCurrent(window, ctx) {
		return internal.LastErr()
	}

	return nil
}

// SDL_EGL_GetWindowSurface - Get the EGL surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface)
func EGL_GetWindowSurface(window *Window) EGLSurface {
	return iEGL_GetWindowSurface(window)
}

// SDL_GL_SwapWindow - Update a window with OpenGL rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow)
func GL_SwapWindow(window *Window) error {
	if !iGL_SwapWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GL_DestroyContext - Delete an OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext)
func GL_DestroyContext(ctx GLContext) {
	iGL_DestroyContext(ctx)
}

// SDL_GL_ResetAttributes - Reset all previously set OpenGL context attributes to their default values.
// (https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes)
func GL_ResetAttributes() {
	iGL_ResetAttributes()
}

// SDL_GL_SetAttribute - Set an OpenGL window attribute before window creation.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute)
func GL_SetAttribute(attr GLAttr, value int32) error {
	if !iGL_SetAttribute(attr, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GL_GetAttribute - Get the actual value for an attribute from the current context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute)
func GL_GetAttribute(attr GLAttr) (int32, error) {
	var value int32

	if !iGL_GetAttribute(attr, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_GL_GetCurrentWindow - Get the currently active OpenGL window.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow)
func GL_GetCurrentWindow() (*Window, error) {
	window := iGL_GetCurrentWindow()
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_GL_GetCurrentContext - Get the currently active OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext)
func GL_GetCurrentContext() (GLContext, error) {
	ctx := iGL_GetCurrentContext()
	if ctx == nil {
		return nil, internal.LastErr()
	}

	return ctx, nil
}

// SDL_EGL_GetCurrentDisplay - Get the currently active EGL display.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay)
func EGL_GetCurrentDisplay() (EGLDisplay, error) {
	display := iEGL_GetCurrentDisplay()
	if display == 0 {
		return 0, internal.LastErr()
	}

	return display, nil
}

// SDL_EGL_GetCurrentConfig - Get the currently active EGL config.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig)
func EGL_GetCurrentConfig() (EGLConfig, error) {
	config := iEGL_GetCurrentConfig()
	if config == 0 {
		return 0, internal.LastErr()
	}

	return config, nil
}

// SDL_EGL_SetAttributeCallbacks - Sets the callbacks for defining custom EGLAttrib arrays for EGL initialization.
// (https://wiki.libsdl.org/SDL3/SDL_EGL_SetAttributeCallbacks)
func EGL_SetAttributeCallbacks(
	platformAttribCallback EGLAttribArrayCallback,
	surfaceAttribCallback EGLIntArrayCallback,
	contextAttribCallback EGLIntArrayCallback,
) {
	iEGL_SetAttributeCallbacks(platformAttribCallback, surfaceAttribCallback, contextAttribCallback, 0)
}

// SDL_GL_SetSwapInterval - Set the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval)
func GL_SetSwapInterval(interval int32) error {
	if !iGL_SetSwapInterval(interval) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GL_GetSwapInterval - Get the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval)
func GL_GetSwapInterval() (int32, error) {
	var interval int32
	if !iGL_GetSwapInterval(&interval) {
		return 0, internal.LastErr()
	}

	return interval, nil
}

// Audio

// SDL_GetNumAudioDrivers - Use this function to get the number of built-in audio drivers.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumAudioDrivers)
func GetNumAudioDrivers() int {
	return int(iGetNumAudioDrivers())
}

// SDL_GetAudioDriver - Use this function to get the name of a built in audio driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDriver)
func GetAudioDriver(index int) string {
	return iGetAudioDriver(int32(index))
}

// SDL_GetCurrentAudioDriver - Get the name of the current audio driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentAudioDriver)
func GetCurrentAudioDriver() string {
	return iGetCurrentAudioDriver()
}

// SDL_GetAudioPlaybackDevices - Get a list of currently-connected audio playback devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioPlaybackDevices)
func GetAudioPlaybackDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioPlaybackDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

// SDL_GetAudioRecordingDevices - Get a list of currently-connected audio recording devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioRecordingDevices)
func GetAudioRecordingDevices() ([]AudioDeviceID, error) {
	var count int32

	ptr := iGetAudioRecordingDevices(&count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[AudioDeviceID](ptr, int(count)), nil
}

// SDL_UnbindAudioStreams - Unbind a list of audio streams from their audio devices.
// (https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStreams)
func UnbindAudioStreams(streams []*AudioStream) {
	iUnbindAudioStreams(unsafe.SliceData(streams), int32(len(streams)))
}

// SDL_CreateAudioStream - Create a new audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_CreateAudioStream)
func CreateAudioStream(srcSpec *AudioSpec, dstSpec *AudioSpec) (*AudioStream, error) {
	stream := iCreateAudioStream(srcSpec, dstSpec)
	if stream == nil {
		return nil, internal.LastErr()
	}

	return stream, nil
}

// SDL_LoadWAV_IO - Load the audio data of a WAVE file into memory.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO)
func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV_IO(src, closeIO, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// SDL_LoadWAV - Loads a WAV from a file path.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV)
func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	var count uint32
	var ptr *byte

	if !iLoadWAV(path, spec, &ptr, &count) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// SDL_MixAudio - Mix audio data in a specified format.
// (https://wiki.libsdl.org/SDL3/SDL_MixAudio)
func MixAudio(src []byte, format AudioFormat, volume float32) ([]byte, error) {
	dst := make([]byte, len(src))
	if !iMixAudio(unsafe.SliceData(dst), unsafe.SliceData(src), format, uint32(len(src)), volume) {
		return nil, internal.LastErr()
	}

	return dst, nil
}

// SDL_ConvertAudioSamples - Convert some audio data of one format to another format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertAudioSamples)
func ConvertAudioSamples(srcSpec *AudioSpec, srcData []byte, dstSpec *AudioSpec) ([]byte, error) {
	var ptr *byte
	var count int32

	if !iConvertAudioSamples(
		srcSpec, unsafe.SliceData(srcData), int32(len(srcData)),
		dstSpec, &ptr, &count,
	) {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[byte](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// Time

// TODO:

// SDL_GetTicks - Get the number of milliseconds since SDL library initialization.
// (https://wiki.libsdl.org/SDL3/SDL_GetTicks)
func Ticks() uint64 {
	return iGetTicks()
}

// SDL_GetTicksNS - Get the number of nanoseconds since SDL library initialization.
// (https://wiki.libsdl.org/SDL3/SDL_GetTicksNS)
func TicksNS() uint64 {
	return iGetTicksNS()
}

// SDL_GetPerformanceCounter - Get the current value of the high resolution counter.
// (https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter)
func GetPerformanceCounter() uint64 {
	return iGetPerformanceCounter()
}

// SDL_GetPerformanceFrequency - Get the count per second of the high resolution counter.
// (https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency)
func GetPerformanceFrequency() uint64 {
	return iGetPerformanceFrequency()
}

// SDL_Delay - Wait a specified number of milliseconds before returning.
// (https://wiki.libsdl.org/SDL3/SDL_Delay)
func Delay(ms uint32) {
	iDelay(ms)
}

// SDL_DelayNS - Wait a specified number of nanoseconds before returning.
// (https://wiki.libsdl.org/SDL3/SDL_DelayNS)
func DelayNS(ns uint64) {
	iDelayNS(ns)
}

// SDL_DelayPrecise - Wait a specified number of nanoseconds before returning.
// (https://wiki.libsdl.org/SDL3/SDL_DelayPrecise)
func DelayPrecise(ns uint64) {
	iDelayPrecise(ns)
}

// Filesystem

// TODO:

// Storage

// TODO:

// Dialog

type dialogFileFilter struct {
	Name    *byte
	Pattern *byte
}

// SDL_ShowOpenFileDialog - Displays a dialog that lets the user select a file on their filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog)
func ShowOpenFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string, allowMany bool) {
	nullableFilters := make([]dialogFileFilter, len(filters))
	for i, filter := range filters {
		nullableFilters[i] = dialogFileFilter{
			Name:    internal.StringToNullablePtr(filter.Name),
			Pattern: internal.StringToNullablePtr(filter.Pattern),
		}
	}
	iShowOpenFileDialog(callback, 0, window, unsafe.SliceData(nullableFilters), int32(len(nullableFilters)), internal.StringToNullablePtr(defaultLocation), allowMany)
	runtime.KeepAlive(nullableFilters)
	runtime.KeepAlive(defaultLocation)
}

// SDL_ShowSaveFileDialog - Displays a dialog that lets the user choose a new or existing file on their filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog)
func ShowSaveFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string) {
	nullableFilters := make([]dialogFileFilter, len(filters))
	for i, filter := range filters {
		nullableFilters[i] = dialogFileFilter{
			Name:    internal.StringToNullablePtr(filter.Name),
			Pattern: internal.StringToNullablePtr(filter.Pattern),
		}
	}
	iShowSaveFileDialog(callback, 0, window, unsafe.SliceData(nullableFilters), int32(len(nullableFilters)), internal.StringToNullablePtr(defaultLocation))
	runtime.KeepAlive(nullableFilters)
	runtime.KeepAlive(defaultLocation)
}

// SDL_ShowOpenFolderDialog - Displays a dialog that lets the user select a folder on their filesystem.
// (https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog)
func ShowOpenFolderDialog(callback DialogFileCallback, window *Window, defaultLocation string, allowMany bool) {
	iShowOpenFolderDialog(callback, 0, window, internal.StringToNullablePtr(defaultLocation), allowMany)
	runtime.KeepAlive(defaultLocation)
}

// SDL_ShowFileDialogWithProperties - Create and launch a file dialog with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties)
func ShowFileDialogWithProperties(callback DialogFileCallback, typ FileDialogType, props PropertiesID) {
	iShowFileDialogWithProperties(typ, callback, 0, props)
}

// SDL_EnumerateDirectory - Enumerate a directory through a callback function.
// (https://wiki.libsdl.org/SDL3/SDL_EnumerateDirectory)
func EnumerateDirectory(path string, callback EnumerateDirectoryCallback) error {
	if !iEnumerateDirectory(path, callback, 0) {
		return internal.LastErr()
	}

	return nil
}

// Tray

// SDL_CreateTray - Create an icon to be placed in the operating system's tray, or equivalent.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTray)
func CreateTray(icon *Surface, tooltip string) *Tray {
	return iCreateTray(icon, internal.StringToNullablePtr(tooltip))
}

// SDL_CreateTraySubmenu - Create a submenu for a system tray entry.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu)
func CreateTraySubmenu(entry *TrayEntry) *TrayMenu {
	return iCreateTraySubmenu(entry)
}

// SDL_RemoveTrayEntry - Removes a tray entry.
// (https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry)
func RemoveTrayEntry(entry *TrayEntry) {
	iRemoveTrayEntry(entry)
}

// SDL_UpdateTrays - Update the trays.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateTrays)
func UpdateTrays() {
	iUpdateTrays()
}

// Message

// SDL_ShowMessageBox - Create a modal message box.
// (https://wiki.libsdl.org/SDL3/SDL_ShowMessageBox)
func ShowMessageBox(data *MessageBoxData) (int32, error) {
	var buttonID int32

	if !iShowMessageBox(data.as(), &buttonID) {
		return 0, internal.LastErr()
	}
	runtime.KeepAlive(data)

	return buttonID, nil
}

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

// SDL_CreateProcessWithProperties - Create a new process with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateProcessWithProperties)
func CreateProcessWithProperties(props PropertiesID) *Process {
	return iCreateProcessWithProperties(props)
}

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

	s := internal.PtrToSlice[*locale](ptr, int(count))
	locales := make([]*Locale, len(s))
	for i, loc := range s {
		locales[i] = &Locale{
			Language: internal.ClonePtrString(uintptr(unsafe.Pointer(loc.Language))),
			Country:  internal.ClonePtrString(uintptr(unsafe.Pointer(loc.Country))),
		}
	}

	return locales, nil
}

// System

// TODO: platform specific stuff

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

// SDL_GetKeyboardFocus - Query the window which currently has keyboard focus.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus)
func GetKeyboardFocus() *Window {
	return iGetKeyboardFocus()
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

// SDL_GetModState - Get the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetModState)
func GetModState() Keymod {
	return iGetModState()
}

// SDL_SetModState - Set the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_SetModState)
func SetModState(state Keymod) {
	iSetModState(state)
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

// SDL_SetRelativeMouseTransform - Set a user-defined function by which to transform relative mouse inputs.
// (https://wiki.libsdl.org/SDL3/SDL_SetRelativeMouseTransform)
func SetRelativeMouseTransform(callback MouseMotionTransformCallback) error {
	if !iSetRelativeMouseTransform(callback, 0) {
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

// SDL_CreateAnimatedCursor - Create an animated color cursor.
// (https://wiki.libsdl.org/SDL3/SDL_CreateAnimatedCursor)
func CreateAnimatedCursor(frames []CursorFrameInfo, hotX, hotY int32) (*Cursor, error) {
	cursor := iCreateAnimatedCursor(unsafe.SliceData(frames), int32(len(frames)), hotX, hotY)
	if cursor == nil {
		return nil, internal.LastErr()
	}
	runtime.KeepAlive(frames)

	return cursor, nil
}

// SDL_CreateSystemCursor - Create a system cursor.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor)
func CreateSystemCursor(id SystemCursor) (*Cursor, error) {
	cursor := iCreateSystemCursor(id)
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

// SDL_GetGamepadTypeFromString - Convert a string into [SDL_GamepadType](SDL_GamepadType) enum.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString)
func GetGamepadTypeFromString(str string) GamepadType {
	return iGetGamepadTypeFromString(str)
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
func AttachVirtualJoystick(desc *VirtualJoystickDesc) (JoystickID, error) {
	id := iAttachVirtualJoystick(desc.as())
	if id == 0 {
		return 0, internal.LastErr()
	}

	runtime.KeepAlive(desc)

	return id, nil
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

// SDL_GetHaptics - Get a list of currently connected haptic devices.
// (https://wiki.libsdl.org/SDL3/SDL_GetHaptics)
func GetHaptics() []HapticID {
	var count int32

	ptr := iGetHaptics(&count)
	if ptr == 0 {
		return nil
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[HapticID](ptr, int(count))
}

// SDL_IsMouseHaptic - Query whether or not the current mouse has haptic capabilities.
// (https://wiki.libsdl.org/SDL3/SDL_IsMouseHaptic)
func IsMouseHaptic() bool {
	return iIsMouseHaptic()
}

// SDL_OpenHapticFromMouse - Try to open a haptic device from the current mouse.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromMouse)
func OpenHapticFromMouse() (*Haptic, error) {
	haptic := iOpenHapticFromMouse()
	if haptic == nil {
		return nil, internal.LastErr()
	}

	return haptic, nil
}

// Rect

// SDL_GetRectEnclosingPoints - Calculate a minimal rectangle enclosing a set of points.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectEnclosingPoints)
func GetRectEnclosingPoints(points []Point, clip *Rect) (*Rect, bool) {
	var result Rect

	ret := iGetRectEnclosingPoints(unsafe.SliceData(points), int32(len(points)), clip, &result)

	return &result, ret
}

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

// SDL_SetClipboardData - Offer clipboard data to the OS.
// (https://wiki.libsdl.org/SDL3/SDL_SetClipboardData)
func SetClipboardData(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, mime_types []string) error {
	mimeTypes := make([]*byte, len(mime_types))
	for i, mime := range mime_types {
		mimeTypes[i] = internal.StringToNullablePtr(mime)
	}
	if !iSetClipboardData(callback, cleanup, 0, unsafe.SliceData(mimeTypes), uintptr(len(mimeTypes))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetVersion - Get the version of SDL that is linked against your program.
// (https://wiki.libsdl.org/SDL3/SDL_GetVersion)
func GetVersion() Version {
	return Version(iGetVersion())
}

// SDL_Vulkan_LoadLibrary - Dynamically load the Vulkan loader library.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_LoadLibrary)
func Vulkan_LoadLibrary(path string) error {
	success := iVulkan_LoadLibrary(path)
	if success {
		return nil
	} else {
		return internal.LastErr()
	}
}

// SDL_Vulkan_GetVkGetInstanceProcAddr - Get the address of the `vkGetInstanceProcAddr` function.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetVkGetInstanceProcAddr)
func Vulkan_GetVkGetInstanceProcAddr() (FunctionPointer, error) {
	fp := iVulkan_GetVkGetInstanceProcAddr()
	if fp != 0 {
		return fp, nil
	} else {
		return 0, internal.LastErr()
	}
}

// SDL_Vulkan_UnloadLibrary - Unload the Vulkan library previously loaded by [SDL_Vulkan_LoadLibrary](SDL_Vulkan_LoadLibrary)().
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_UnloadLibrary)
func Vulkan_UnloadLibrary() {
	iVulkan_UnloadLibrary()
}

// SDL_Vulkan_GetInstanceExtensions - Get the Vulkan instance extensions needed for vkCreateInstance.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetInstanceExtensions)
func Vulkan_GetInstanceExtensions() ([]string, error) {
	var count uint32

	ptr := iVulkan_GetInstanceExtensions(&count)
	if ptr == nil {
		return nil, internal.LastErr()
	}
	// Dont free pointer, its owned by sdl
	return internal.BytePtrPtrToStrSlice(ptr, count, true), nil
}

// SDL_Vulkan_CreateSurface - Create a Vulkan rendering surface for a window.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_CreateSurface)
func Vulkan_CreateSurface(window *Window, instance VkInstance, allocator VkAllocationCallbacksPtr) (VkSurfaceKHR, error) {
	var surface VkSurfaceKHR
	success := iVulkan_CreateSurface(window, instance, (*VkAllocationCallbacks)(unsafe.Pointer(uintptr(allocator))), &surface)
	if !success {
		return 0, internal.LastErr()
	}

	return surface, nil
}

// SDL_Vulkan_DestroySurface - Destroy the Vulkan rendering surface of a window.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_DestroySurface)
func Vulkan_DestroySurface(instance VkInstance, surface VkSurfaceKHR, allocator VkAllocationCallbacksPtr) {
	iVulkan_DestroySurface(instance, surface, (*VkAllocationCallbacks)(unsafe.Pointer(uintptr(allocator))))
}

// SDL_Vulkan_GetPresentationSupport - Query support for presentation via a given physical device and queue family.
// (https://wiki.libsdl.org/SDL3/SDL_Vulkan_GetPresentationSupport)
func Vulkan_GetPresentationSupport(instance VkInstance, physicalDevice VkPhysicalDevice, queueFamilyIndex uint32) (bool, error) {
	iClearError()
	support := iVulkan_GetPresentationSupport(instance, physicalDevice, queueFamilyIndex)
	if !support {
		return false, internal.LastErr()
	}

	return support, nil
}
