package sdl

import (
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

// TODO: IsMainThread
// TODO: RunOnMainThread
// TODO: SetAppMetadata
// TODO: SetAppMetadataProperty
// TODO: GetAppMetadataProperty

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

// TODO: is there a need?

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

// Log

// TODO: is there a need?

// Events

// SDL_PumpEvents - Pump the event loop, gathering events from the input devices.
// (https://wiki.libsdl.org/SDL3/SDL_PumpEvents)
func PumpEvents() {
	iPumpEvents()
}

// TODO: PeepEvents

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

// TODO: SetEventFilter
// TODO: GetEventFilter
// TODO: AddEventWatch
// TODO: RemoveEventWatch
// TODO: FilterEvents

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

// SDL_IOFromConstMem - Use this function to prepare a read-only memory buffer for use with SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_IOFromConstMem)
func IOFromConstMem(mem []byte) (*IOStream, error) {
	stream := iIOFromConstMem(
		uintptr(unsafe.Pointer(unsafe.SliceData(mem))),
		uintptr(len(mem)),
	)
	if stream == nil {
		return nil, internal.LastErr()
	}

	return stream, nil
}

// TODO:

// AsyncIO

// TODO:

// Main

// TODO: is this needed?
