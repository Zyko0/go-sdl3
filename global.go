package sdl

import "unsafe"

// Init

func Init(flags InitFlags) error {
	if !iInit(flags) {
		return lastError()
	}

	return nil
}

func InitSubSystem(flags InitFlags) error {
	if !iInitSubSystem(flags) {
		return lastError()
	}

	return nil
}

func QuitSubSystem(flags InitFlags) {
	iQuitSubSystem(flags)
}

func WasInit(flags InitFlags) InitFlags {
	return iWasInit(flags)
}

func Quit() {
	iQuit()
}

// TODO: IsMainThread
// TODO: RunOnMainThread
// TODO: SetAppMetadata
// TODO: SetAppMetadataProperty
// TODO: GetAppMetadataProperty

// Hints

func SetHintWithPriority(name, value string, priority HintPriority) error {
	if !iSetHintWithPriority(name, value, priority) {
		return lastError()
	}

	return nil
}

func SetHint(name, value string) error {
	if !iSetHint(name, value) {
		return lastError()
	}

	return nil
}

func ResetHint(name string) error {
	if !iResetHint(name) {
		return lastError()
	}

	return nil
}

func ResetHints() {
	iResetHints()
}

func GetHint(name string) string {
	return iGetHint(name)
}

func GetHintBoolean(name string, defaultValue bool) bool {
	return iGetHintBoolean(name, defaultValue)
}

// TODO: AddHintCallback
// TODO: RemoveHintCallback

// Error

// TODO: is there a need?

// Properties

func GetGlobalProperties() (PropertiesID, error) {
	properties := iGetGlobalProperties()
	if properties == 0 {
		return 0, lastError()
	}

	return properties, nil
}

func CreateProperties() (PropertiesID, error) {
	properties := iCreateProperties()
	if properties == 0 {
		return 0, lastError()
	}

	return properties, nil
}

// Log

// TODO: is there a need?

// Events

func PumpEvents() {
	iPumpEvents()
}

// TODO: PeepEvents

func HasEvent(typ EventType) bool {
	return iHasEvent(uint32(typ))
}

func HasEvents(minType, maxType EventType) bool {
	return iHasEvents(uint32(minType), uint32(maxType))
}

func FlushEvent(typ EventType) {
	iFlushEvent(uint32(typ))
}

func FlushEvents(minType, maxType EventType) {
	iFlushEvents(uint32(minType), uint32(maxType))
}

func PollEvent(event *Event) bool {
	return iPollEvent(event)
}

func WaitEvent(event *Event) error {
	if !iWaitEvent(event) {
		return lastError()
	}

	return nil
}

func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	return iWaitEventTimeout(event, timeoutMS)
}

func PushEvent(event *Event) error {
	if !iPushEvent(event) {
		return lastError()
	}

	return nil
}

// TODO: SetEventFilter
// TODO: GetEventFilter
// TODO: AddEventWatch
// TODO: RemoveEventWatch
// TODO: FilterEvents

func SetEventEnabled(typ EventType, enabled bool) {
	iSetEventEnabled(uint32(typ), enabled)
}

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

func IOFromConstMem(mem []byte) (*IOStream, error) {
	stream := iIOFromConstMem(
		uintptr(unsafe.Pointer(unsafe.SliceData(mem))),
		uintptr(len(mem)),
	)
	if stream == nil {
		return nil, lastError()
	}

	return stream, nil
}

// TODO:

// AsyncIO

// TODO:

// Main

// TODO: is this needed?
