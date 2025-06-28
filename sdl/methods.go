/*

This file has been generated initially with cmd/methodgen
to have a starting point. Functions need to be made Go-like,
return errors, strings and slices cloned/freed if necessary based
on the documentation.

*/

package sdl

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// ThreadPriority

// SDL_SetCurrentThreadPriority - Set the priority for the current thread.
// (https://wiki.libsdl.org/SDL3/SDL_SetCurrentThreadPriority)
func (priority ThreadPriority) SetCurrent() bool {
	panic("not implemented")
	return iSetCurrentThreadPriority(priority)
}

// IOStreamInterface

// SDL_OpenIO - Create a custom SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_OpenIO)
func (iface *IOStreamInterface) OpenIO(userdata *byte) *IOStream {
	panic("not implemented")
	//return iOpenIO(iface, userdata)
}

// ClipboardDataCallback

// SDL_SetClipboardData - Offer clipboard data to the OS.
// (https://wiki.libsdl.org/SDL3/SDL_SetClipboardData)
func (callback ClipboardDataCallback) SetClipboardData(cleanup ClipboardCleanupCallback, userdata *byte, mime_types *string, num_mime_types uintptr) bool {
	panic("not implemented")
	//return iSetClipboardData(callback, cleanup, userdata, mime_types, num_mime_types)
}

// WindowID

// SDL_GetWindowFromID - Get a window from a stored ID.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID)
func (id WindowID) Window() (*Window, error) {
	window := iGetWindowFromID(id)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// TouchID

// SDL_GetTouchDeviceName - Get the touch device name as reported from the driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName)
func (touchID TouchID) DeviceName() (string, error) {
	name := iGetTouchDeviceName(touchID)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetTouchDeviceType - Get the type of the given touch device.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType)
func (touchID TouchID) DeviceType() TouchDeviceType {
	return iGetTouchDeviceType(touchID)
}

// SDL_GetTouchFingers - Get a list of active fingers for a given touch device.
// (https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers)
func (touchID TouchID) Fingers() ([]*Finger, error) {
	var count int32

	ptr := iGetTouchFingers(touchID, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Finger](ptr, int(count)), nil
}

// Storage

// SDL_CloseStorage - Closes and frees a storage container.
// (https://wiki.libsdl.org/SDL3/SDL_CloseStorage)
func (storage *Storage) Close() error {
	if !iCloseStorage(storage) {
		return internal.LastErr()
	}

	return nil
}

// SDL_StorageReady - Checks if the storage container is ready to use.
// (https://wiki.libsdl.org/SDL3/SDL_StorageReady)
func (storage *Storage) Ready() bool {
	return iStorageReady(storage)
}

// SDL_GetStorageFileSize - Query the size of a file within a storage container.
// (https://wiki.libsdl.org/SDL3/SDL_GetStorageFileSize)
func (storage *Storage) FileSize(path string) (uint64, error) {
	var length uint64

	if !iGetStorageFileSize(storage, path, &length) {
		return 0, internal.LastErr()
	}

	return length, nil
}

// SDL_ReadStorageFile - Synchronously read a file from a storage container into a client-provided buffer.
// (https://wiki.libsdl.org/SDL3/SDL_ReadStorageFile)
func (storage *Storage) ReadFile(path string, length uint64) ([]byte, error) {
	data := make([]byte, length)

	if !iReadStorageFile(storage, path, uintptr(unsafe.Pointer(unsafe.SliceData(data))), length) {
		return nil, internal.LastErr()
	}

	return data, nil
}

// SDL_WriteStorageFile - Synchronously write a file from client memory into a storage container.
// (https://wiki.libsdl.org/SDL3/SDL_WriteStorageFile)
func (storage *Storage) WriteFile(path string, source []byte) error {
	if !iWriteStorageFile(storage, path, uintptr(unsafe.Pointer(unsafe.SliceData(source))), uint64(len(source))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CreateStorageDirectory - Create a directory in a writable storage container.
// (https://wiki.libsdl.org/SDL3/SDL_CreateStorageDirectory)
func (storage *Storage) CreateDirectory(path string) error {
	if !iCreateStorageDirectory(storage, path) {
		return internal.LastErr()
	}

	return nil
}

// SDL_EnumerateStorageDirectory - Enumerate a directory in a storage container through a callback function.
// (https://wiki.libsdl.org/SDL3/SDL_EnumerateStorageDirectory)
func (storage *Storage) EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata *byte) bool {
	panic("not implemented")
	//return iEnumerateStorageDirectory(storage, path, callback, userdata)
}

// SDL_RemoveStoragePath - Remove a file or an empty directory in a writable storage container.
// (https://wiki.libsdl.org/SDL3/SDL_RemoveStoragePath)
func (storage *Storage) RemovePath(path string) error {
	if !iRemoveStoragePath(storage, path) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenameStoragePath - Rename a file or directory in a writable storage container.
// (https://wiki.libsdl.org/SDL3/SDL_RenameStoragePath)
func (storage *Storage) RenamePath(oldpath, newpath string) error {
	if !iRenameStoragePath(storage, oldpath, newpath) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CopyStorageFile - Copy a file in a writable storage container.
// (https://wiki.libsdl.org/SDL3/SDL_CopyStorageFile)
func (storage *Storage) CopyFile(oldpath, newpath string) error {
	if !iCopyStorageFile(storage, oldpath, newpath) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetStoragePathInfo - Get information about a filesystem path in a storage container.
// (https://wiki.libsdl.org/SDL3/SDL_GetStoragePathInfo)
func (storage *Storage) PathInfo(path string) (*PathInfo, error) {
	var info PathInfo

	if !iGetStoragePathInfo(storage, path, &info) {
		return nil, internal.LastErr()
	}

	return &info, nil
}

// SDL_GetStorageSpaceRemaining - Queries the remaining space in a storage container.
// (https://wiki.libsdl.org/SDL3/SDL_GetStorageSpaceRemaining)
func (storage *Storage) SpaceRemaining() uint64 {
	return iGetStorageSpaceRemaining(storage)
}

// SDL_GlobStorageDirectory - Enumerate a directory tree, filtered by pattern, and return a list.
// (https://wiki.libsdl.org/SDL3/SDL_GlobStorageDirectory)
func (storage *Storage) GlobDirectory(path string, pattern string, flags GlobFlags) ([]string, error) {
	var count int32
	ptr := iGlobStorageDirectory(storage, path, pattern, flags, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[string](ptr, int(count)), nil
}

// AudioDeviceID

// SDL_GetAudioDeviceName - Get the human-readable name of a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceName)
func (devid AudioDeviceID) Name() (string, error) {
	name := iGetAudioDeviceName(devid)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetAudioDeviceFormat - Get the current audio format of a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceFormat)
func (devid AudioDeviceID) Format() (*AudioSpec, int32, error) {
	var spec AudioSpec
	var sampleFrames int32

	if !iGetAudioDeviceFormat(devid, &spec, &sampleFrames) {
		return nil, 0, internal.LastErr()
	}

	return &spec, sampleFrames, nil
}

// SDL_GetAudioDeviceChannelMap - Get the current channel map of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceChannelMap)
func (devid AudioDeviceID) ChannelMap() []int {
	var count int32

	ptr := iGetAudioDeviceChannelMap(devid, &count)
	if ptr == 0 {
		return nil
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[int](ptr, int(count))
}

// SDL_OpenAudioDevice - Open a specific audio device.
// (https://wiki.libsdl.org/SDL3/SDL_OpenAudioDevice)
func (devid AudioDeviceID) OpenAudioDevice(spec *AudioSpec) (AudioDeviceID, error) {
	ret := iOpenAudioDevice(devid, spec)
	if ret == 0 {
		return 0, internal.LastErr()
	}

	return ret, nil
}

// SDL_IsAudioDevicePhysical - Determine if an audio device is physical (instead of logical).
// (https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePhysical)
func (devid AudioDeviceID) IsPhysical() bool {
	return iIsAudioDevicePhysical(devid)
}

// SDL_IsAudioDevicePlayback - Determine if an audio device is a playback device (instead of recording).
// (https://wiki.libsdl.org/SDL3/SDL_IsAudioDevicePlayback)
func (devid AudioDeviceID) IsPlayback() bool {
	return iIsAudioDevicePlayback(devid)
}

// SDL_PauseAudioDevice - Use this function to pause audio playback on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_PauseAudioDevice)
func (devid AudioDeviceID) Pause() error {
	if !iPauseAudioDevice(devid) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ResumeAudioDevice - Use this function to unpause audio playback on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeAudioDevice)
func (devid AudioDeviceID) Resume() error {
	if !iResumeAudioDevice(devid) {
		return internal.LastErr()
	}

	return nil
}

// SDL_AudioDevicePaused - Use this function to query if an audio device is paused.
// (https://wiki.libsdl.org/SDL3/SDL_AudioDevicePaused)
func (devid AudioDeviceID) Paused() bool {
	return iAudioDevicePaused(devid)
}

// SDL_GetAudioDeviceGain - Get the gain of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioDeviceGain)
func (devid AudioDeviceID) Gain() (float32, error) {
	gain := iGetAudioDeviceGain(devid)
	if gain == -1 {
		return -1, internal.LastErr()
	}

	return gain, nil
}

// SDL_SetAudioDeviceGain - Change the gain of an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioDeviceGain)
func (devid AudioDeviceID) SetGain(gain float32) error {
	if !iSetAudioDeviceGain(devid, gain) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CloseAudioDevice - Close a previously-opened audio device.
// (https://wiki.libsdl.org/SDL3/SDL_CloseAudioDevice)
func (devid AudioDeviceID) Close() {
	iCloseAudioDevice(devid)
}

// SDL_BindAudioStreams - Bind a list of audio streams to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_BindAudioStreams)
func (devid AudioDeviceID) BindAudioStreams(streams []*AudioStream) error {
	if !iBindAudioStreams(devid, unsafe.SliceData(streams), int32(len(streams))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BindAudioStream - Bind a single audio stream to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_BindAudioStream)
func (devid AudioDeviceID) BindAudioStream(stream *AudioStream) error {
	if !iBindAudioStream(devid, stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_OpenAudioDeviceStream - Convenience function for straightforward audio init for the common case.
// (https://wiki.libsdl.org/SDL3/SDL_OpenAudioDeviceStream)
func (devid AudioDeviceID) OpenAudioDeviceStream(spec *AudioSpec, callback AudioStreamCallback, userdata *byte) *AudioStream {
	panic("not implemented")
	//return iOpenAudioDeviceStream(devid, spec, callback, userdata)
}

// SDL_SetAudioPostmixCallback - Set a callback that fires when data is about to be fed to an audio device.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioPostmixCallback)
func (devid AudioDeviceID) SetAudioPostmixCallback(callback AudioPostmixCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetAudioPostmixCallback(devid, callback, userdata)
}

// Camera

// SDL_GetCameraPermissionState - Query if camera access has been approved by the user.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraPermissionState)
func (camera *Camera) PermissionState() int32 {
	return iGetCameraPermissionState(camera)
}

// SDL_GetCameraID - Get the instance ID of an opened camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraID)
func (camera *Camera) ID() (CameraID, error) {
	id := iGetCameraID(camera)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetCameraProperties - Get the properties associated with an opened camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraProperties)
func (camera *Camera) Properties() (PropertiesID, error) {
	props := iGetCameraProperties(camera)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetCameraFormat - Get the spec that a camera is using when generating images.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraFormat)
func (camera *Camera) Format() (*CameraSpec, error) {
	var spec CameraSpec

	if !iGetCameraFormat(camera, &spec) {
		return nil, internal.LastErr()
	}

	return &spec, nil
}

// SDL_AcquireCameraFrame - Acquire a frame.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireCameraFrame)
func (camera *Camera) AcquireFrame() (*Surface, uint64, error) {
	var timestampNS uint64

	surface := iAcquireCameraFrame(camera, &timestampNS)
	if timestampNS == 0 {
		return nil, 0, internal.LastErr()
	}

	return surface, timestampNS, nil
}

// SDL_ReleaseCameraFrame - Release a frame of video acquired from a camera.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseCameraFrame)
func (camera *Camera) ReleaseFrame(frame *Surface) {
	iReleaseCameraFrame(camera, frame)
}

// SDL_CloseCamera - Use this function to shut down camera processing and close the camera device.
// (https://wiki.libsdl.org/SDL3/SDL_CloseCamera)
func (camera *Camera) Close() {
	iCloseCamera(camera)
}

// GamepadButton

// SDL_GetGamepadStringForButton - Convert from an SDL_GamepadButton enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton)
func (button GamepadButton) GamepadStringForButton() string {
	return iGetGamepadStringForButton(button)
}

// GPUTextureFormat

// SDL_GPUTextureFormatTexelBlockSize - Obtains the texel block size for a texture format.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureFormatTexelBlockSize)
func (format GPUTextureFormat) TexelBlockSize() uint32 {
	return iGPUTextureFormatTexelBlockSize(format)
}

// SDL_CalculateGPUTextureFormatSize - Calculate the size in bytes of a texture format with dimensions.
// (https://wiki.libsdl.org/SDL3/SDL_CalculateGPUTextureFormatSize)
func (format GPUTextureFormat) CalculateSize(width, height, depthOrLayerCount uint32) uint32 {
	return iCalculateGPUTextureFormatSize(format, width, height, depthOrLayerCount)
}

// SpinLock

// SDL_TryLockSpinlock - Try to lock a spin lock by setting it to a non-zero value.
// (https://wiki.libsdl.org/SDL3/SDL_TryLockSpinlock)
func (lock *SpinLock) TryLock() bool {
	return iTryLockSpinlock(lock)
}

// SDL_LockSpinlock - Lock a spin lock by setting it to a non-zero value.
// (https://wiki.libsdl.org/SDL3/SDL_LockSpinlock)
func (lock *SpinLock) Lock() {
	iLockSpinlock(lock)
}

// SDL_UnlockSpinlock - Unlock a spin lock by setting it to 0.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockSpinlock)
func (lock *SpinLock) Unlock() {
	iUnlockSpinlock(lock)
}

// Thread

// SDL_GetThreadName - Get the thread name as it was specified in SDL_CreateThread().
// (https://wiki.libsdl.org/SDL3/SDL_GetThreadName)
func (thread *Thread) Name() string {
	panic("not implemented")
	return iGetThreadName(thread)
}

// SDL_GetThreadID - Get the thread identifier for the specified thread.
// (https://wiki.libsdl.org/SDL3/SDL_GetThreadID)
func (thread *Thread) ID() ThreadID {
	panic("not implemented")
	return iGetThreadID(thread)
}

// SDL_WaitThread - Wait for a thread to finish.
// (https://wiki.libsdl.org/SDL3/SDL_WaitThread)
func (thread *Thread) Wait(status *int32) {
	panic("not implemented")
	iWaitThread(thread, status)
}

// SDL_GetThreadState - Get the current state of a thread.
// (https://wiki.libsdl.org/SDL3/SDL_GetThreadState)
func (thread *Thread) State() ThreadState {
	panic("not implemented")
	return iGetThreadState(thread)
}

// SDL_DetachThread - Let a thread clean up on exit without intervention.
// (https://wiki.libsdl.org/SDL3/SDL_DetachThread)
func (thread *Thread) Detach() {
	panic("not implemented")
	iDetachThread(thread)
}

// Sensor

// SDL_GetSensorProperties - Get the properties associated with a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorProperties)
func (sensor *Sensor) Properties() (PropertiesID, error) {
	props := iGetSensorProperties(sensor)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetSensorName - Get the implementation dependent name of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorName)
func (sensor *Sensor) Name() (string, error) {
	name := iGetSensorName(sensor)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetSensorType - Get the type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorType)
func (sensor *Sensor) Type() SensorType {
	return iGetSensorType(sensor)
}

// SDL_GetSensorNonPortableType - Get the platform dependent type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableType)
func (sensor *Sensor) NonPortableType() int32 {
	return iGetSensorNonPortableType(sensor)
}

// SDL_GetSensorID - Get the instance ID of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorID)
func (sensor *Sensor) ID() (SensorID, error) {
	id := iGetSensorID(sensor)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetSensorData - Get the current state of an opened sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorData)
func (sensor *Sensor) Data(data []float32, numValues int32) error {
	if !iGetSensorData(sensor, unsafe.SliceData(data), numValues) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CloseSensor - Close a sensor previously opened with SDL_OpenSensor().
// (https://wiki.libsdl.org/SDL3/SDL_CloseSensor)
func (sensor *Sensor) Close() {
	iCloseSensor(sensor)
}

// GamepadAxis

// SDL_GetGamepadStringForAxis - Convert from an SDL_GamepadAxis enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis)
func (axis GamepadAxis) GamepadStringForAxis() string {
	return iGetGamepadStringForAxis(axis)
}

// Cursor

// SDL_DestroyCursor - Free a previously-created cursor.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyCursor)
func (cursor *Cursor) Destroy() {
	iDestroyCursor(cursor)
}

// X11EventHook

// SDL_SetX11EventHook - Set a callback for every X11 event.
// (https://wiki.libsdl.org/SDL3/SDL_SetX11EventHook)
func (callback X11EventHook) Set(userdata *byte) {
	panic("not implemented")
	//iSetX11EventHook(callback, userdata)
}

// TLSID

// SDL_GetTLS - Get the current thread's value associated with a thread local storage ID.
// (https://wiki.libsdl.org/SDL3/SDL_GetTLS)
func (id *TLSID) TLS() *byte {
	panic("not implemented")
	//return iGetTLS(id)
}

// SDL_SetTLS - Set the current thread's value associated with a thread local storage ID.
// (https://wiki.libsdl.org/SDL3/SDL_SetTLS)
func (id *TLSID) SetTLS(value *byte, destructor TLSDestructorCallback) bool {
	panic("not implemented")
	//return iSetTLS(id, value, destructor)
}

// Rect

// SDL_HasRectIntersection - Determine whether two rectangles intersect.
// (https://wiki.libsdl.org/SDL3/SDL_HasRectIntersection)
func (a *Rect) HasIntersection(b *Rect) bool {
	return iHasRectIntersection(a, b)
}

// SDL_GetRectIntersection - Calculate the intersection of two rectangles.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectIntersection)
func (a *Rect) Intersection(b *Rect) *Rect {
	var result Rect

	if !iGetRectIntersection(a, b, &result) {
		return nil
	}

	return &result
}

// SDL_GetRectUnion - Calculate the union of two rectangles.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectUnion)
func (a *Rect) Union(b *Rect) (*Rect, error) {
	var result Rect

	if !iGetRectUnion(a, b, &result) {
		return nil, internal.LastErr()
	}

	return &result, nil
}

// JoystickID

// SDL_GetJoystickNameForID - Get the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID)
func (id JoystickID) JoystickNameForID() (string, error) {
	name := iGetJoystickNameForID(id)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetJoystickPathForID - Get the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID)
func (id JoystickID) JoystickPathForID() (string, error) {
	path := iGetJoystickPathForID(id)
	if path == "" {
		return "", internal.LastErr()
	}

	return path, nil
}

// SDL_GetJoystickPlayerIndexForID - Get the player index of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID)
func (id JoystickID) JoystickPlayerIndexForID() int32 {
	return iGetJoystickPlayerIndexForID(id)
}

// SDL_GetJoystickGUIDForID - Get the implementation-dependent GUID of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID)
func (id JoystickID) JoystickGUIDForID() GUID {
	panic("not implemented")
	return iGetJoystickGUIDForID(id)
}

// SDL_GetJoystickVendorForID - Get the USB vendor ID of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID)
func (id JoystickID) JoystickVendorForID() uint16 {
	return iGetJoystickVendorForID(id)
}

// SDL_GetJoystickProductForID - Get the USB product ID of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID)
func (id JoystickID) JoystickProductForID() uint16 {
	return iGetJoystickProductForID(id)
}

// SDL_GetJoystickProductVersionForID - Get the product version of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID)
func (id JoystickID) JoystickProductVersionForID() uint16 {
	return iGetJoystickProductVersionForID(id)
}

// SDL_GetJoystickTypeForID - Get the type of a joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID)
func (id JoystickID) JoystickTypeForID() JoystickType {
	return iGetJoystickTypeForID(id)
}

// SDL_OpenJoystick - Open a joystick for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenJoystick)
func (id JoystickID) OpenJoystick() (*Joystick, error) {
	joystick := iOpenJoystick(id)
	if joystick == nil {
		return nil, internal.LastErr()
	}

	return joystick, nil
}

// SDL_GetJoystickFromID - Get the SDL_Joystick associated with an instance ID, if it has been opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID)
func (id JoystickID) Joystick() (*Joystick, error) {
	joystick := iGetJoystickFromID(id)
	if joystick == nil {
		return nil, internal.LastErr()
	}

	return joystick, nil
}

// SDL_DetachVirtualJoystick - Detach a virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick)
func (id JoystickID) DetachVirtualJoystick() error {
	if !iDetachVirtualJoystick(id) {
		return internal.LastErr()
	}

	return nil
}

// SDL_IsJoystickVirtual - Query whether or not a joystick is virtual.
// (https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual)
func (id JoystickID) IsJoystickVirtual() bool {
	return iIsJoystickVirtual(id)
}

// SDL_SetGamepadMapping - Set the current mapping of a joystick or gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping)
func (instance_id JoystickID) SetGamepadMapping(mapping string) bool {
	panic("not implemented")
	return iSetGamepadMapping(instance_id, mapping)
}

// SDL_IsGamepad - Check if the given joystick is supported by the gamepad interface.
// (https://wiki.libsdl.org/SDL3/SDL_IsGamepad)
func (id JoystickID) IsGamepad() bool {
	return iIsGamepad(id)
}

// SDL_GetGamepadNameForID - Get the implementation dependent name of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID)
func (id JoystickID) GamepadName() (string, error) {
	name := iGetGamepadNameForID(id)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetGamepadPathForID - Get the implementation dependent path of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID)
func (id JoystickID) GamepadPath() (string, error) {
	path := iGetGamepadPathForID(id)
	if path == "" {
		return "", internal.LastErr()
	}

	return path, nil
}

// SDL_GetGamepadPlayerIndexForID - Get the player index of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID)
func (id JoystickID) GamepadPlayerIndex() int32 {
	return iGetGamepadPlayerIndexForID(id)
}

// SDL_GetGamepadGUIDForID - Get the implementation-dependent GUID of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID)
func (instance_id JoystickID) GamepadGUIDForID() GUID {
	panic("not implemented")
	return iGetGamepadGUIDForID(instance_id)
}

// SDL_GetGamepadVendorForID - Get the USB vendor ID of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID)
func (id JoystickID) GamepadVendor() uint16 {
	return iGetGamepadVendorForID(id)
}

// SDL_GetGamepadProductForID - Get the USB product ID of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID)
func (id JoystickID) GamepadProduct() uint16 {
	return iGetGamepadProductForID(id)
}

// SDL_GetGamepadProductVersionForID - Get the product version of a gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID)
func (id JoystickID) GamepadProductVersion() uint16 {
	return iGetGamepadProductVersionForID(id)
}

// SDL_GetGamepadTypeForID - Get the type of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID)
func (id JoystickID) GamepadType() GamepadType {
	return iGetGamepadTypeForID(id)
}

// SDL_GetRealGamepadTypeForID - Get the type of a gamepad, ignoring any mapping override.
// (https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID)
func (id JoystickID) RealGamepadType() GamepadType {
	return iGetRealGamepadTypeForID(id)
}

// SDL_GetGamepadMappingForID - Get the mapping of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID)
func (id JoystickID) GamepadMapping() string {
	ptr := iGetGamepadMappingForID(id)
	if ptr == 0 {
		return ""
	}
	defer internal.Free(ptr)

	return internal.ClonePtrString(ptr)
}

// SDL_OpenGamepad - Open a gamepad for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenGamepad)
func (id JoystickID) OpenGamepad() (*Gamepad, error) {
	gamepad := iOpenGamepad(id)
	if gamepad == nil {
		return nil, internal.LastErr()
	}

	return gamepad, nil
}

// SDL_GetGamepadFromID - Get the SDL_Gamepad associated with a joystick instance ID, if it has been opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID)
func (id JoystickID) Gamepad() (*Gamepad, error) {
	gamepad := iGetGamepadFromID(id)
	if gamepad == nil {
		return nil, internal.LastErr()
	}

	return gamepad, nil
}

// AtomicInt

// SDL_CompareAndSwapAtomicInt - Set an atomic variable to a new value if it is currently an old value.
// (https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicInt)
func (a *AtomicInt) CompareAndSwap(oldval, newval int32) bool {
	return iCompareAndSwapAtomicInt(a, oldval, newval)
}

// SDL_SetAtomicInt - Set an atomic variable to a value.
// (https://wiki.libsdl.org/SDL3/SDL_SetAtomicInt)
func (a *AtomicInt) Set(v int32) int32 {
	return iSetAtomicInt(a, v)
}

// SDL_GetAtomicInt - Get the value of an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_GetAtomicInt)
func (a *AtomicInt) Get() int32 {
	return iGetAtomicInt(a)
}

// SDL_AddAtomicInt - Add to an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_AddAtomicInt)
func (a *AtomicInt) Add(v int32) int32 {
	return iAddAtomicInt(a, v)
}

// GPUComputePass

// SDL_BindGPUComputePipeline - Binds a compute pipeline on a command buffer for use in compute dispatch.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputePipeline)
func (cp *GPUComputePass) BindGPUComputePipeline(pipeline *GPUComputePipeline) {
	iBindGPUComputePipeline(cp, pipeline)
}

// SDL_BindGPUComputeSamplers - Binds texture-sampler pairs for use on the compute shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeSamplers)
func (cp *GPUComputePass) BindSamplers(bindings []GPUTextureSamplerBinding) {
	iBindGPUComputeSamplers(cp, 0, unsafe.SliceData(bindings), uint32(len(bindings)))
}

// SDL_BindGPUComputeStorageTextures - Binds storage textures as readonly for use on the compute pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageTextures)
func (cp *GPUComputePass) BindStorageTextures(textures []*GPUTexture) {
	iBindGPUComputeStorageTextures(cp, 0, unsafe.SliceData(textures), uint32(len(textures)))
}

// SDL_BindGPUComputeStorageBuffers - Binds storage buffers as readonly for use on the compute pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUComputeStorageBuffers)
func (cp *GPUComputePass) BindStorageBuffers(buffers []*GPUBuffer) {
	iBindGPUComputeStorageBuffers(cp, 0, unsafe.SliceData(buffers), uint32(len(buffers)))
}

// SDL_DispatchGPUCompute - Dispatches compute work.
// (https://wiki.libsdl.org/SDL3/SDL_DispatchGPUCompute)
func (cp *GPUComputePass) Dispatch(groupcountX, groupcountY, groupcountZ uint32) {
	iDispatchGPUCompute(cp, groupcountX, groupcountY, groupcountZ)
}

// SDL_DispatchGPUComputeIndirect - Dispatches compute work with parameters set from a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DispatchGPUComputeIndirect)
func (cp *GPUComputePass) DispatchIndirect(buffer *GPUBuffer, offset uint32) {
	iDispatchGPUComputeIndirect(cp, buffer, offset)
}

// SDL_EndGPUComputePass - Ends the current compute pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPUComputePass)
func (cp *GPUComputePass) End() {
	iEndGPUComputePass(cp)
}

// Texture

// SDL_GetTextureProperties - Get the properties associated with a texture.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties)
func (texture *Texture) Properties() (PropertiesID, error) {
	props := iGetTextureProperties(texture)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetRendererFromTexture - Get the renderer that created an SDL_Texture.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture)
func (texture *Texture) Renderer() (*Renderer, error) {
	renderer := iGetRendererFromTexture(texture)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// SDL_GetTextureSize - Get the size of a texture, as floating point values.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureSize)
func (texture *Texture) Size() (float32, float32, error) {
	var w, h float32
	if !iGetTextureSize(texture, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_SetTextureColorMod - Set an additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod)
func (texture *Texture) SetColorMod(r, g, b uint8) error {
	if !iSetTextureColorMod(texture, r, g, b) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetTextureColorModFloat - Set an additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat)
func (texture *Texture) SetColorModFloat(r float32, g float32, b float32) error {
	if !iSetTextureColorModFloat(texture, r, g, b) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetTextureColorMod - Get the additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod)
func (texture *Texture) ColorMod() (uint8, uint8, uint8, error) {
	var r, g, b uint8
	if !iGetTextureColorMod(texture, &r, &g, &b) {
		return 0, 0, 0, internal.LastErr()
	}

	return r, g, b, nil
}

// SDL_GetTextureColorModFloat - Get the additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat)
func (texture *Texture) ColorModFloat() (float32, float32, float32, error) {
	var r, g, b float32
	if !iGetTextureColorModFloat(texture, &r, &g, &b) {
		return 0, 0, 0, internal.LastErr()
	}

	return r, g, b, nil
}

// SDL_SetTextureAlphaMod - Set an additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod)
func (texture *Texture) SetAlphaMod(alpha uint8) error {
	if !iSetTextureAlphaMod(texture, alpha) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetTextureAlphaModFloat - Set an additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat)
func (texture *Texture) SetAlphaModFloat(alpha float32) error {
	if !iSetTextureAlphaModFloat(texture, alpha) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetTextureAlphaMod - Get the additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod)
func (texture *Texture) AlphaMod() (uint8, error) {
	var alpha uint8
	if !iGetTextureAlphaMod(texture, &alpha) {
		return 0, internal.LastErr()
	}

	return alpha, nil
}

// SDL_GetTextureAlphaModFloat - Get the additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat)
func (texture *Texture) AlphaModFloat() (float32, error) {
	var alpha float32
	if !iGetTextureAlphaModFloat(texture, &alpha) {
		return 0, internal.LastErr()
	}

	return alpha, nil
}

// SDL_SetTextureBlendMode - Set the blend mode for a texture, used by SDL_RenderTexture().
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode)
func (texture *Texture) SetBlendMode(blendMode BlendMode) error {
	if !iSetTextureBlendMode(texture, blendMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetTextureBlendMode - Get the blend mode used for texture copy operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode)
func (texture *Texture) BlendMode() (BlendMode, error) {
	var mode BlendMode
	if !iGetTextureBlendMode(texture, &mode) {
		return mode, internal.LastErr()
	}

	return mode, nil
}

// SDL_SetTextureScaleMode - Set the scale mode used for texture scale operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode)
func (texture *Texture) SetScaleMode(scaleMode ScaleMode) error {
	if !iSetTextureScaleMode(texture, scaleMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetTextureScaleMode - Get the scale mode used for texture scale operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode)
func (texture *Texture) ScaleMode() (ScaleMode, error) {
	var mode ScaleMode
	if !iGetTextureScaleMode(texture, &mode) {
		return mode, internal.LastErr()
	}

	return mode, nil
}

// SDL_UpdateTexture - Update the given texture rectangle with new pixel data.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateTexture)
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int32) error {
	if !iUpdateTexture(texture, rect, uintptr(unsafe.Pointer(unsafe.SliceData(pixels))), pitch) {
		return internal.LastErr()
	}

	runtime.KeepAlive(pixels)

	return nil
}

// SDL_UpdateYUVTexture - Update a rectangle within a planar YV12 or IYUV texture with new pixel data.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture)
func (texture *Texture) UpdateYUV(rect *Rect, Yplane *uint8, Ypitch int32, Uplane *uint8, Upitch int32, Vplane *uint8, Vpitch int32) bool {
	panic("not implemented")
	return iUpdateYUVTexture(texture, rect, Yplane, Ypitch, Uplane, Upitch, Vplane, Vpitch)
}

// SDL_UpdateNVTexture - Update a rectangle within a planar NV12 or NV21 texture with new pixels.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture)
func (texture *Texture) UpdateNV(rect *Rect, Yplane *uint8, Ypitch int32, UVplane *uint8, UVpitch int32) bool {
	panic("not implemented")
	return iUpdateNVTexture(texture, rect, Yplane, Ypitch, UVplane, UVpitch)
}

// SDL_LockTexture - Lock a portion of the texture for **write-only** pixel access.
// Returns pixels data, pitch, error.
// (https://wiki.libsdl.org/SDL3/SDL_LockTexture)
func (texture *Texture) Lock(rect *Rect) ([]byte, int32, error) {
	var ptr uintptr
	var pitch int32

	if !iLockTexture(texture, rect, &ptr, &pitch) {
		return nil, 0, internal.LastErr()
	}

	return unsafe.Slice(*(**byte)(unsafe.Pointer(&ptr)), texture.H*pitch), pitch, nil
}

// SDL_LockTextureToSurface - Lock a portion of the texture for **write-only** pixel access, and expose it as a SDL surface.
// (https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface)
func (texture *Texture) LockToSurface(rect *Rect, surface **Surface) error {
	if !iLockTextureToSurface(texture, rect, surface) {
		return internal.LastErr()
	}

	return nil
}

// SDL_UnlockTexture - Unlock a texture, uploading the changes to video memory, if needed.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockTexture)
func (texture *Texture) Unlock() {
	iUnlockTexture(texture)
}

// SDL_DestroyTexture - Destroy the specified texture.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyTexture)
func (texture *Texture) Destroy() {
	if texture == nil {
		return
	}
	iDestroyTexture(texture)
}

// RWLock

// SDL_LockRWLockForReading - Lock the read/write lock for _read only_ operations.
// (https://wiki.libsdl.org/SDL3/SDL_LockRWLockForReading)
func (rwlock *RWLock) LockForReading() {
	iLockRWLockForReading(rwlock)
}

// SDL_LockRWLockForWriting - Lock the read/write lock for _write_ operations.
// (https://wiki.libsdl.org/SDL3/SDL_LockRWLockForWriting)
func (rwlock *RWLock) LockForWriting() {
	iLockRWLockForWriting(rwlock)
}

// SDL_TryLockRWLockForReading - Try to lock a read/write lock _for reading_ without blocking.
// (https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForReading)
func (rwlock *RWLock) TryLockForReading() bool {
	return iTryLockRWLockForReading(rwlock)
}

// SDL_TryLockRWLockForWriting - Try to lock a read/write lock _for writing_ without blocking.
// (https://wiki.libsdl.org/SDL3/SDL_TryLockRWLockForWriting)
func (rwlock *RWLock) TryLockForWriting() bool {
	return iTryLockRWLockForWriting(rwlock)
}

// SDL_UnlockRWLock - Unlock the read/write lock.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockRWLock)
func (rwlock *RWLock) Unlock() {
	iUnlockRWLock(rwlock)
}

// SDL_DestroyRWLock - Destroy a read/write lock created with SDL_CreateRWLock().
// (https://wiki.libsdl.org/SDL3/SDL_DestroyRWLock)
func (rwlock *RWLock) Destroy() {
	iDestroyRWLock(rwlock)
}

// AudioFormat

// SDL_GetAudioFormatName - Get the human readable name of an audio format.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioFormatName)
func (format AudioFormat) Name() string {
	return iGetAudioFormatName(format)
}

// SDL_GetSilenceValueForFormat - Get the appropriate memset value for silencing an audio format.
// (https://wiki.libsdl.org/SDL3/SDL_GetSilenceValueForFormat)
func (format AudioFormat) SilenceValueForFormat() int32 {
	return iGetSilenceValueForFormat(format)
}

// GPURenderPass

// SDL_BindGPUGraphicsPipeline - Binds a graphics pipeline on a render pass to be used in rendering.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUGraphicsPipeline)
func (rp *GPURenderPass) BindGraphicsPipeline(pipeline *GPUGraphicsPipeline) {
	iBindGPUGraphicsPipeline(rp, pipeline)
}

// SDL_SetGPUViewport - Sets the current viewport state on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUViewport)
func (rp *GPURenderPass) SetGPUViewport(viewport *GPUViewport) {
	iSetGPUViewport(rp, viewport)
}

// SDL_SetGPUScissor - Sets the current scissor state on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUScissor)
func (rp *GPURenderPass) SetScissor(scissor *Rect) {
	iSetGPUScissor(rp, scissor)
}

// SDL_SetGPUStencilReference - Sets the current stencil reference value on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUStencilReference)
func (rp *GPURenderPass) SetStencilReference(reference uint8) {
	iSetGPUStencilReference(rp, reference)
}

// SDL_BindGPUVertexBuffers - Binds vertex buffers on a command buffer for use with subsequent draw calls.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexBuffers)
func (rp *GPURenderPass) BindVertexBuffers(bindings []GPUBufferBinding) {
	iBindGPUVertexBuffers(rp, 0, unsafe.SliceData(bindings), uint32(len(bindings)))
}

// SDL_BindGPUIndexBuffer - Binds an index buffer on a command buffer for use with subsequent draw calls.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUIndexBuffer)
func (rp *GPURenderPass) BindIndexBuffer(binding *GPUBufferBinding, indexElementSize GPUIndexElementSize) {
	iBindGPUIndexBuffer(rp, binding, indexElementSize)
}

// SDL_BindGPUVertexSamplers - Binds texture-sampler pairs for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexSamplers)
func (rp *GPURenderPass) BindVertexSamplers(bindings []GPUTextureSamplerBinding) {
	iBindGPUVertexSamplers(rp, 0, unsafe.SliceData(bindings), uint32(len(bindings)))
}

// SDL_BindGPUVertexStorageTextures - Binds storage textures for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageTextures)
func (rp *GPURenderPass) BindVertexStorageTextures(textures []*GPUTexture) {
	iBindGPUVertexStorageTextures(rp, 0, unsafe.SliceData(textures), uint32(len(textures)))
}

// SDL_BindGPUVertexStorageBuffers - Binds storage buffers for use on the vertex shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUVertexStorageBuffers)
func (rp *GPURenderPass) BindVertexStorageBuffers(buffers []*GPUBuffer) {
	iBindGPUVertexStorageBuffers(rp, 0, unsafe.SliceData(buffers), uint32(len(buffers)))
}

// SDL_BindGPUFragmentSamplers - Binds texture-sampler pairs for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentSamplers)
func (rp *GPURenderPass) BindFragmentSamplers(bindings []GPUTextureSamplerBinding) {
	iBindGPUFragmentSamplers(rp, 0, unsafe.SliceData(bindings), uint32(len(bindings)))
}

// SDL_BindGPUFragmentStorageTextures - Binds storage textures for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageTextures)
func (rp *GPURenderPass) BindFragmentStorageTextures(textures []*GPUTexture) {
	iBindGPUFragmentStorageTextures(rp, 0, unsafe.SliceData(textures), uint32(len(textures)))
}

// SDL_BindGPUFragmentStorageBuffers - Binds storage buffers for use on the fragment shader.
// (https://wiki.libsdl.org/SDL3/SDL_BindGPUFragmentStorageBuffers)
func (rp *GPURenderPass) BindFragmentStorageBuffers(buffers []*GPUBuffer) {
	iBindGPUFragmentStorageBuffers(rp, 0, unsafe.SliceData(buffers), uint32(len(buffers)))
}

// SDL_DrawGPUIndexedPrimitives - Draws data using bound graphics state with an index buffer and instancing enabled.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitives)
func (rp *GPURenderPass) DrawIndexedPrimitives(numIndices, numInstances, firstIndex uint32, vertexOffset int32, firstInstance uint32) {
	iDrawGPUIndexedPrimitives(rp, numIndices, numInstances, firstIndex, vertexOffset, firstInstance)
}

// SDL_DrawGPUPrimitives - Draws data using bound graphics state.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitives)
func (rp *GPURenderPass) DrawPrimitives(numVertices, numInstances, firstVertex, firstInstance uint32) {
	iDrawGPUPrimitives(rp, numVertices, numInstances, firstVertex, firstInstance)
}

// SDL_DrawGPUPrimitivesIndirect - Draws data using bound graphics state and with draw parameters set from a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUPrimitivesIndirect)
func (rp *GPURenderPass) DrawPrimitivesIndirect(buffer *GPUBuffer, offset, drawCount uint32) {
	iDrawGPUPrimitivesIndirect(rp, buffer, offset, drawCount)
}

// SDL_DrawGPUIndexedPrimitivesIndirect - Draws data using bound graphics state with an index buffer enabled and with draw parameters set from a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_DrawGPUIndexedPrimitivesIndirect)
func (rp *GPURenderPass) DrawIndexedPrimitivesIndirect(buffer *GPUBuffer, offset, drawCount uint32) {
	iDrawGPUIndexedPrimitivesIndirect(rp, buffer, offset, drawCount)
}

// SDL_EndGPURenderPass - Ends the given render pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPURenderPass)
func (rp *GPURenderPass) End() {
	iEndGPURenderPass(rp)
}

// AsyncIOQueue

// SDL_DestroyAsyncIOQueue - Destroy a previously-created async I/O task queue.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyAsyncIOQueue)
func (queue *AsyncIOQueue) Destroy() {
	iDestroyAsyncIOQueue(queue)
}

// SDL_GetAsyncIOResult - Query an async I/O task queue for completed tasks.
// (https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOResult)
func (queue *AsyncIOQueue) Result() (*AsyncIOOutcome, bool) {
	var outcome AsyncIOOutcome

	ret := iGetAsyncIOResult(queue, &outcome)

	return &outcome, ret
}

// SDL_WaitAsyncIOResult - Block until an async I/O task queue has a completed task.
// (https://wiki.libsdl.org/SDL3/SDL_WaitAsyncIOResult)
func (queue *AsyncIOQueue) WaitAsyncIOResult(timeoutMS int32) (*AsyncIOOutcome, bool) {
	var outcome AsyncIOOutcome

	ret := iWaitAsyncIOResult(queue, &outcome, timeoutMS)

	return &outcome, ret
}

// SDL_SignalAsyncIOQueue - Wake up any threads that are blocking in SDL_WaitAsyncIOResult().
// (https://wiki.libsdl.org/SDL3/SDL_SignalAsyncIOQueue)
func (queue *AsyncIOQueue) Signal() {
	iSignalAsyncIOQueue(queue)
}

// PixelFormatDetails

// SDL_MapRGB - Map an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGB)
func (format *PixelFormatDetails) MapRGB(palette *Palette, r, g, b uint8) uint32 {
	return iMapRGB(format, palette, r, g, b)
}

// SDL_MapRGBA - Map an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGBA)
func (format *PixelFormatDetails) MapRGBA(palette *Palette, r, g, b, a uint8) uint32 {
	return iMapRGBA(format, palette, r, g, b, a)
}

// Surface

// SDL_DestroySurface - Free a surface.
// (https://wiki.libsdl.org/SDL3/SDL_DestroySurface)
func (surface *Surface) Destroy() {
	if surface == nil {
		return
	}
	iDestroySurface(surface)
}

// SDL_GetSurfaceProperties - Get the properties associated with a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties)
func (surface *Surface) Properties() (PropertiesID, error) {
	props := iGetSurfaceProperties(surface)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_SetSurfaceColorspace - Set the colorspace used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace)
func (surface *Surface) SetColorspace(colorspace Colorspace) error {
	if !iSetSurfaceColorspace(surface, colorspace) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetSurfaceColorspace - Get the colorspace used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace)
func (surface *Surface) Colorspace() Colorspace {
	return iGetSurfaceColorspace(surface)
}

// SDL_CreateSurfacePalette - Create a palette and associate it with a surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette)
func (surface *Surface) CreatePalette() (*Palette, error) {
	palette := iCreateSurfacePalette(surface)
	if palette == nil {
		return nil, internal.LastErr()
	}

	return palette, nil
}

// SDL_SetSurfacePalette - Set the palette used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette)
func (surface *Surface) SetPalette(palette *Palette) error {
	if !iSetSurfacePalette(surface, palette) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetSurfacePalette - Get the palette used by a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette)
func (surface *Surface) Palette() *Palette {
	return iGetSurfacePalette(surface)
}

// SDL_AddSurfaceAlternateImage - Add an alternate version of a surface.
// (https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage)
func (surface *Surface) AddAlternateImage(image *Surface) error {
	if !iAddSurfaceAlternateImage(surface, image) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SurfaceHasAlternateImages - Return whether a surface has alternate versions available.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages)
func (surface *Surface) HasAlternateImages() bool {
	return iSurfaceHasAlternateImages(surface)
}

// SDL_GetSurfaceImages - Get an array including all versions of a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages)
func (surface *Surface) Images() ([]*Surface, error) {
	var count int32

	ptr := iGetSurfaceImages(surface, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*Surface](ptr, int(count)), nil
}

// SDL_RemoveSurfaceAlternateImages - Remove all alternate versions of a surface.
// (https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages)
func (surface *Surface) RemoveAlternateImages() {
	iRemoveSurfaceAlternateImages(surface)
}

// SDL_LockSurface - Set up a surface for directly accessing the pixels.
// (https://wiki.libsdl.org/SDL3/SDL_LockSurface)
func (surface *Surface) Lock() error {
	if !iLockSurface(surface) {
		return internal.LastErr()
	}

	return nil
}

// SDL_UnlockSurface - Release a surface after directly accessing the pixels.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockSurface)
func (surface *Surface) Unlock() {
	iUnlockSurface(surface)
}

// SDL_SaveBMP_IO - Save a surface to a seekable SDL data stream in BMP format.
// (https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO)
func (surface *Surface) SaveBMP_IO(dst *IOStream, closeio bool) error {
	if !iSaveBMP_IO(surface, dst, closeio) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SaveBMP - Save a surface to a file.
// (https://wiki.libsdl.org/SDL3/SDL_SaveBMP)
func (surface *Surface) SaveBMP(file string) error {
	if !iSaveBMP(surface, file) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetSurfaceRLE - Set the RLE acceleration hint for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE)
func (surface *Surface) SetRLE(enabled bool) error {
	if !iSetSurfaceRLE(surface, enabled) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SurfaceHasRLE - Returns whether the surface is RLE enabled.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE)
func (surface *Surface) HasRLE() bool {
	return iSurfaceHasRLE(surface)
}

// SDL_SetSurfaceColorKey - Set the color key (transparent pixel) in a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey)
func (surface *Surface) SetColorKey(enabled bool, key uint32) error {
	if !iSetSurfaceColorKey(surface, enabled, key) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SurfaceHasColorKey - Returns whether the surface has a color key.
// (https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey)
func (surface *Surface) HasColorKey() bool {
	return iSurfaceHasColorKey(surface)
}

// SDL_GetSurfaceColorKey - Get the color key (transparent pixel) for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey)
func (surface *Surface) ColorKey() (uint32, error) {
	var key uint32
	if !iGetSurfaceColorKey(surface, &key) {
		return 0, internal.LastErr()
	}

	return key, nil
}

// SDL_SetSurfaceColorMod - Set an additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod)
func (surface *Surface) SetColorMod(r, g, b uint8) error {
	if !iSetSurfaceColorMod(surface, r, g, b) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetSurfaceColorMod - Get the additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod)
func (surface *Surface) ColorMod() (uint8, uint8, uint8, error) {
	var r, g, b uint8

	if !iGetSurfaceColorMod(surface, &r, &g, &b) {
		return 0, 0, 0, internal.LastErr()
	}

	return r, g, b, nil
}

// SDL_SetSurfaceAlphaMod - Set an additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod)
func (surface *Surface) SetAlphaMod(alpha uint8) error {
	if !iSetSurfaceAlphaMod(surface, alpha) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetSurfaceAlphaMod - Get the additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod)
func (surface *Surface) AlphaMod() (uint8, error) {
	var alpha uint8

	if !iGetSurfaceAlphaMod(surface, &alpha) {
		return 0, internal.LastErr()
	}

	return alpha, nil
}

// SDL_SetSurfaceBlendMode - Set the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode)
func (surface *Surface) SetBlendMode(blendMode BlendMode) error {
	if !iSetSurfaceBlendMode(surface, blendMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetSurfaceBlendMode - Get the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode)
func (surface *Surface) BlendMode() (BlendMode, error) {
	var mode BlendMode

	if !iGetSurfaceBlendMode(surface, &mode) {
		return mode, internal.LastErr()
	}

	return mode, nil
}

// SDL_SetSurfaceClipRect - Set the clipping rectangle for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect)
func (surface *Surface) SetClipRect(rect *Rect) bool {
	return iSetSurfaceClipRect(surface, rect)
}

// SDL_GetSurfaceClipRect - Get the clipping rectangle for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect)
func (surface *Surface) ClipRect() (*Rect, error) {
	var rect Rect

	if !iGetSurfaceClipRect(surface, &rect) {
		return nil, internal.LastErr()
	}

	return &rect, nil
}

// SDL_FlipSurface - Flip a surface vertically or horizontally.
// (https://wiki.libsdl.org/SDL3/SDL_FlipSurface)
func (surface *Surface) Flip(flip FlipMode) error {
	if !iFlipSurface(surface, flip) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DuplicateSurface - Creates a new surface identical to the existing surface.
// (https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface)
func (surface *Surface) Duplicate() (*Surface, error) {
	dup := iDuplicateSurface(surface)
	if dup == nil {
		return nil, internal.LastErr()
	}

	return dup, nil
}

// SDL_ScaleSurface - Creates a new surface identical to the existing surface, scaled to the desired size.
// (https://wiki.libsdl.org/SDL3/SDL_ScaleSurface)
func (surface *Surface) Scale(width, height int32, scaleMode ScaleMode) (*Surface, error) {
	scaled := iScaleSurface(surface, width, height, scaleMode)
	if scaled == nil {
		return nil, internal.LastErr()
	}

	return scaled, nil
}

// SDL_ConvertSurface - Copy an existing surface to a new surface of the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertSurface)
func (surface *Surface) Convert(format PixelFormat) (*Surface, error) {
	s := iConvertSurface(surface, format)
	if s == nil {
		return nil, internal.LastErr()
	}

	return s, nil
}

// SDL_ConvertSurfaceAndColorspace - Copy an existing surface to a new surface of the specified format and colorspace.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace)
func (surface *Surface) ConvertWithColorspace(format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) (*Surface, error) {
	converted := iConvertSurfaceAndColorspace(surface, format, palette, colorspace, props)
	if converted == nil {
		return nil, internal.LastErr()
	}

	return converted, nil
}

// SDL_PremultiplySurfaceAlpha - Premultiply the alpha in a surface.
// (https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha)
func (surface *Surface) PremultiplyAlpha(linear bool) error {
	if !iPremultiplySurfaceAlpha(surface, linear) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ClearSurface - Clear a surface with a specific color, with floating point precision.
// (https://wiki.libsdl.org/SDL3/SDL_ClearSurface)
func (surface *Surface) Clear(r, g, b, a float32) error {
	if !iClearSurface(surface, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// SDL_FillSurfaceRect - Perform a fast fill of a rectangle with a specific color.
// (https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect)
func (dst *Surface) FillRect(rect *Rect, color uint32) error {
	if !iFillSurfaceRect(dst, rect, color) {
		return internal.LastErr()
	}

	return nil
}

// SDL_FillSurfaceRects - Perform a fast fill of a set of rectangles with a specific color.
// (https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects)
func (dst *Surface) FillRects(rects []Rect, color uint32) error {
	if !iFillSurfaceRects(dst, unsafe.SliceData(rects), int32(len(rects)), color) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurface - Performs a fast blit from the source surface to the destination surface with clipping.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurface)
func (src *Surface) Blit(srcrect *Rect, dst *Surface, dstrect *Rect) error {
	if !iBlitSurface(src, srcrect, dst, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurfaceUnchecked - Perform low-level surface blitting only.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked)
func (src *Surface) BlitUnchecked(srcrect *Rect, dst *Surface, dstrect *Rect) error {
	if !iBlitSurfaceUnchecked(src, srcrect, dst, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurfaceScaled - Perform a scaled blit to a destination surface, which may be of a different format.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled)
func (src *Surface) BlitScaled(srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) error {
	if !iBlitSurfaceScaled(src, srcrect, dst, dstrect, scaleMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurfaceUncheckedScaled - Perform low-level surface scaled blitting only.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled)
func (src *Surface) BlitUncheckedScaled(srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) error {
	if !iBlitSurfaceUncheckedScaled(src, srcrect, dst, dstrect, scaleMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurfaceTiled - Perform a tiled blit to a destination surface, which may be of a different format.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled)
func (src *Surface) BlitTiled(srcrect *Rect, dst *Surface, dstrect *Rect) error {
	if !iBlitSurfaceTiled(src, srcrect, dst, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurfaceTiledWithScale - Perform a scaled and tiled blit to a destination surface, which may be of a different format.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale)
func (src *Surface) BlitTiledWithScale(srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) error {
	if !iBlitSurfaceTiledWithScale(src, srcrect, scale, scaleMode, dst, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_BlitSurface9Grid - Perform a scaled blit using the 9-grid algorithm to a destination surface, which may be of a different format.
// (https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid)
func (src *Surface) Blit9Grid(srcrect *Rect, leftWidth, rightWidth, topHeight, bottomHeight int32, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) error {
	if !iBlitSurface9Grid(src, srcrect, leftWidth, rightWidth, topHeight, bottomHeight, scale, scaleMode, dst, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_MapSurfaceRGB - Map an RGB triple to an opaque pixel value for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB)
func (surface *Surface) MapRGB(r, g, b uint8) uint32 {
	return iMapSurfaceRGB(surface, r, g, b)
}

// SDL_MapSurfaceRGBA - Map an RGBA quadruple to a pixel value for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA)
func (surface *Surface) MapRGBA(r, g, b, a uint8) uint32 {
	return iMapSurfaceRGBA(surface, r, g, b, a)
}

// SDL_ReadSurfacePixel - Retrieves a single pixel from a surface.
// (https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel)
func (surface *Surface) ReadPixel(x, y int32) (uint8, uint8, uint8, uint8, error) {
	var r, g, b, a uint8

	if !iReadSurfacePixel(surface, x, y, &r, &g, &b, &a) {
		return 0, 0, 0, 0, internal.LastErr()
	}

	return r, g, b, a, nil
}

// SDL_ReadSurfacePixelFloat - Retrieves a single pixel from a surface.
// (https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat)
func (surface *Surface) ReadPixelFloat(x, y int32) (float32, float32, float32, float32, error) {
	var r, g, b, a float32

	if !iReadSurfacePixelFloat(surface, x, y, &r, &g, &b, &a) {
		return 0, 0, 0, 0, internal.LastErr()
	}

	return r, g, b, a, nil
}

// SDL_WriteSurfacePixel - Writes a single pixel to a surface.
// (https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel)
func (surface *Surface) WritePixel(x, y int32, r, g, b, a uint8) error {
	if !iWriteSurfacePixel(surface, x, y, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteSurfacePixelFloat - Writes a single pixel to a surface.
// (https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat)
func (surface *Surface) WritePixelFloat(x, y int32, r, g, b, a float32) error {
	if !iWriteSurfacePixelFloat(surface, x, y, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CreateColorCursor - Create a color cursor.
// (https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor)
func (surface *Surface) CreateColorCursor(hotX, hotY int32) (*Cursor, error) {
	cursor := iCreateColorCursor(surface, hotX, hotY)
	if cursor == nil {
		return nil, internal.LastErr()
	}

	return cursor, nil
}

// SDL_CreateSoftwareRenderer - Create a 2D software rendering context for a surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer)
func (surface *Surface) CreateSoftwareRenderer() (*Renderer, error) {
	renderer := iCreateSoftwareRenderer(surface)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// Event

// SDL_GetWindowFromEvent - Get window associated with an event.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent)
func (event *Event) Window() *Window {
	return iGetWindowFromEvent(event)
}

// GPUDevice

// SDL_DestroyGPUDevice - Destroys a GPU context previously returned by SDL_CreateGPUDevice.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyGPUDevice)
func (device *GPUDevice) Destroy() {
	iDestroyGPUDevice(device)
}

// SDL_GetGPUDeviceDriver - Returns the name of the backend used to create this GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUDeviceDriver)
func (device *GPUDevice) Driver() string {
	return iGetGPUDeviceDriver(device)
}

// SDL_GetGPUShaderFormats - Returns the supported shader formats for this GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUShaderFormats)
func (device *GPUDevice) ShaderFormats() GPUShaderFormat {
	return iGetGPUShaderFormats(device)
}

// SDL_CreateGPUComputePipeline - Creates a pipeline object to be used in a compute workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUComputePipeline)
func (device *GPUDevice) CreateComputePipeline(createinfo *GPUComputePipelineCreateInfo) (*GPUComputePipeline, error) {
	pipeline := iCreateGPUComputePipeline(device, &gpuComputePipelineCreateInfo{
		CodeSize:                    uintptr(createinfo.CodeSize),
		Code:                        unsafe.SliceData(createinfo.Code),
		Entrypoint:                  internal.StringToNullablePtr(createinfo.Entrypoint),
		Format:                      createinfo.Format,
		NumSamplers:                 createinfo.NumSamplers,
		NumReadonlyStorageTextures:  createinfo.NumReadonlyStorageTextures,
		NumReadonlyStorageBuffers:   createinfo.NumReadonlyStorageBuffers,
		NumReadwriteStorageTextures: createinfo.NumReadwriteStorageTextures,
		NumReadwriteStorageBuffers:  createinfo.NumReadwriteStorageBuffers,
		NumUniformBuffers:           createinfo.NumUniformBuffers,
		ThreadcountX:                createinfo.ThreadcountX,
		ThreadcountY:                createinfo.ThreadcountY,
		ThreadcountZ:                createinfo.ThreadcountZ,
		Props:                       createinfo.Props,
	})
	if pipeline == nil {
		return nil, internal.LastErr()
	}

	return pipeline, nil
}

// SDL_CreateGPUGraphicsPipeline - Creates a pipeline object to be used in a graphics workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUGraphicsPipeline)
func (device *GPUDevice) CreateGraphicsPipeline(createinfo *GPUGraphicsPipelineCreateInfo) (*GPUGraphicsPipeline, error) {
	pipeline := iCreateGPUGraphicsPipeline(device, createinfo)
	if pipeline == nil {
		return nil, internal.LastErr()
	}

	return pipeline, nil
}

// SDL_CreateGPUSampler - Creates a sampler object to be used when binding textures in a graphics workflow.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUSampler)
func (device *GPUDevice) CreateSampler(createinfo *GPUSamplerCreateInfo) *GPUSampler {
	return iCreateGPUSampler(device, createinfo)
}

// SDL_CreateGPUShader - Creates a shader to be used when creating a graphics pipeline.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUShader)
func (device *GPUDevice) CreateGPUShader(createinfo *GPUShaderCreateInfo) (*GPUShader, error) {
	shader := iCreateGPUShader(device, &gpuShaderCreateInfo{
		CodeSize:           uintptr(createinfo.CodeSize),
		Code:               unsafe.SliceData(createinfo.Code),
		Entrypoint:         internal.StringToNullablePtr(createinfo.Entrypoint),
		Format:             createinfo.Format,
		Stage:              createinfo.Stage,
		NumSamplers:        createinfo.NumSamplers,
		NumStorageTextures: createinfo.NumStorageTextures,
		NumStorageBuffers:  createinfo.NumStorageBuffers,
		NumUniformBuffers:  createinfo.NumUniformBuffers,
		Props:              createinfo.Props,
	})
	if shader == nil {
		return nil, internal.LastErr()
	}

	return shader, nil
}

// SDL_CreateGPUTexture - Creates a texture object to be used in graphics or compute workflows.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUTexture)
func (device *GPUDevice) CreateTexture(createinfo *GPUTextureCreateInfo) (*GPUTexture, error) {
	texture := iCreateGPUTexture(device, createinfo)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// SDL_CreateGPUBuffer - Creates a buffer object to be used in graphics or compute workflows.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUBuffer)
func (device *GPUDevice) CreateBuffer(createinfo *GPUBufferCreateInfo) (*GPUBuffer, error) {
	buffer := iCreateGPUBuffer(device, createinfo)
	if buffer == nil {
		return nil, internal.LastErr()
	}

	return buffer, nil
}

// SDL_CreateGPUTransferBuffer - Creates a transfer buffer to be used when uploading to or downloading from graphics resources.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUTransferBuffer)
func (device *GPUDevice) CreateTransferBuffer(createinfo *GPUTransferBufferCreateInfo) (*GPUTransferBuffer, error) {
	buffer := iCreateGPUTransferBuffer(device, createinfo)
	if buffer == nil {
		return nil, internal.LastErr()
	}

	return buffer, nil
}

// SDL_SetGPUBufferName - Sets an arbitrary string constant to label a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUBufferName)
func (device *GPUDevice) SetBufferName(buffer *GPUBuffer, text string) {
	iSetGPUBufferName(device, buffer, text)
}

// SDL_SetGPUTextureName - Sets an arbitrary string constant to label a texture.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUTextureName)
func (device *GPUDevice) SetTextureName(texture *GPUTexture, text string) {
	iSetGPUTextureName(device, texture, text)
}

// SDL_ReleaseGPUTexture - Frees the given texture as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTexture)
func (device *GPUDevice) ReleaseTexture(texture *GPUTexture) {
	iReleaseGPUTexture(device, texture)
}

// SDL_ReleaseGPUSampler - Frees the given sampler as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUSampler)
func (device *GPUDevice) ReleaseSampler(sampler *GPUSampler) {
	iReleaseGPUSampler(device, sampler)
}

// SDL_ReleaseGPUBuffer - Frees the given buffer as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUBuffer)
func (device *GPUDevice) ReleaseBuffer(buffer *GPUBuffer) {
	iReleaseGPUBuffer(device, buffer)
}

// SDL_ReleaseGPUTransferBuffer - Frees the given transfer buffer as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUTransferBuffer)
func (device *GPUDevice) ReleaseTransferBuffer(buffer *GPUTransferBuffer) {
	iReleaseGPUTransferBuffer(device, buffer)
}

// SDL_ReleaseGPUComputePipeline - Frees the given compute pipeline as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUComputePipeline)
func (device *GPUDevice) ReleaseComputePipeline(pipeline *GPUComputePipeline) {
	iReleaseGPUComputePipeline(device, pipeline)
}

// SDL_ReleaseGPUShader - Frees the given shader as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUShader)
func (device *GPUDevice) ReleaseShader(shader *GPUShader) {
	iReleaseGPUShader(device, shader)
}

// SDL_ReleaseGPUGraphicsPipeline - Frees the given graphics pipeline as soon as it is safe to do so.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUGraphicsPipeline)
func (device *GPUDevice) ReleaseGraphicsPipeline(pipeline *GPUGraphicsPipeline) {
	iReleaseGPUGraphicsPipeline(device, pipeline)
}

// SDL_AcquireGPUCommandBuffer - Acquire a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireGPUCommandBuffer)
func (device *GPUDevice) AcquireCommandBuffer() (*GPUCommandBuffer, error) {
	buffer := iAcquireGPUCommandBuffer(device)
	if buffer == nil {
		return nil, internal.LastErr()
	}

	return buffer, nil
}

// SDL_MapGPUTransferBuffer - Maps a transfer buffer into application address space.
// (https://wiki.libsdl.org/SDL3/SDL_MapGPUTransferBuffer)
func (device *GPUDevice) MapTransferBuffer(buffer *GPUTransferBuffer, cycle bool) (uintptr, error) {
	ptr := iMapGPUTransferBuffer(device, buffer, cycle)
	if ptr == 0 {
		return 0, internal.LastErr()
	}

	return ptr, nil
}

// SDL_UnmapGPUTransferBuffer - Unmaps a previously mapped transfer buffer.
// (https://wiki.libsdl.org/SDL3/SDL_UnmapGPUTransferBuffer)
func (device *GPUDevice) UnmapTransferBuffer(buffer *GPUTransferBuffer) {
	iUnmapGPUTransferBuffer(device, buffer)
}

// SDL_WindowSupportsGPUSwapchainComposition - Determines whether a swapchain composition is supported by the window.
// (https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUSwapchainComposition)
func (device *GPUDevice) WindowSupportsSwapchainComposition(window *Window, swapchainComposition GPUSwapchainComposition) bool {
	return iWindowSupportsGPUSwapchainComposition(device, window, swapchainComposition)
}

// SDL_WindowSupportsGPUPresentMode - Determines whether a presentation mode is supported by the window.
// (https://wiki.libsdl.org/SDL3/SDL_WindowSupportsGPUPresentMode)
func (device *GPUDevice) WindowSupportsPresentMode(window *Window, presentMode GPUPresentMode) bool {
	return iWindowSupportsGPUPresentMode(device, window, presentMode)
}

// SDL_ClaimWindowForGPUDevice - Claims a window, creating a swapchain structure for it.
// (https://wiki.libsdl.org/SDL3/SDL_ClaimWindowForGPUDevice)
func (device *GPUDevice) ClaimWindow(window *Window) error {
	if !iClaimWindowForGPUDevice(device, window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ReleaseWindowFromGPUDevice - Unclaims a window, destroying its swapchain structure.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseWindowFromGPUDevice)
func (device *GPUDevice) ReleaseWindow(window *Window) {
	iReleaseWindowFromGPUDevice(device, window)
}

// SDL_SetGPUSwapchainParameters - Changes the swapchain parameters for the given claimed window.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUSwapchainParameters)
func (device *GPUDevice) SetSwapchainParameters(window *Window, swapchainComposition GPUSwapchainComposition, presentMode GPUPresentMode) error {
	if !iSetGPUSwapchainParameters(device, window, swapchainComposition, presentMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetGPUAllowedFramesInFlight - Configures the maximum allowed number of frames in flight.
// (https://wiki.libsdl.org/SDL3/SDL_SetGPUAllowedFramesInFlight)
func (device *GPUDevice) SetAllowedFramesInFlight(allowedFramesInFlight uint32) error {
	if !iSetGPUAllowedFramesInFlight(device, allowedFramesInFlight) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetGPUSwapchainTextureFormat - Obtains the texture format of the swapchain for the given window.
// (https://wiki.libsdl.org/SDL3/SDL_GetGPUSwapchainTextureFormat)
func (device *GPUDevice) SwapchainTextureFormat(window *Window) GPUTextureFormat {
	return iGetGPUSwapchainTextureFormat(device, window)
}

// SDL_WaitForGPUSwapchain - Blocks the thread until a swapchain texture is available to be acquired.
// (https://wiki.libsdl.org/SDL3/SDL_WaitForGPUSwapchain)
func (device *GPUDevice) WaitForSwapchain(window *Window) error {
	if !iWaitForGPUSwapchain(device, window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WaitForGPUIdle - Blocks the thread until the GPU is completely idle.
// (https://wiki.libsdl.org/SDL3/SDL_WaitForGPUIdle)
func (device *GPUDevice) WaitForIdle() error {
	if !iWaitForGPUIdle(device) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WaitForGPUFences - Blocks the thread until the given fences are signaled.
// (https://wiki.libsdl.org/SDL3/SDL_WaitForGPUFences)
func (device *GPUDevice) WaitForFences(waitAll bool, fences []*GPUFence) error {
	if !iWaitForGPUFences(device, waitAll, unsafe.SliceData(fences), uint32(len(fences))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_QueryGPUFence - Checks the status of a fence.
// (https://wiki.libsdl.org/SDL3/SDL_QueryGPUFence)
func (device *GPUDevice) QueryFence(fence *GPUFence) bool {
	return iQueryGPUFence(device, fence)
}

// SDL_ReleaseGPUFence - Releases a fence obtained from SDL_SubmitGPUCommandBufferAndAcquireFence.
// (https://wiki.libsdl.org/SDL3/SDL_ReleaseGPUFence)
func (device *GPUDevice) ReleaseFence(fence *GPUFence) {
	iReleaseGPUFence(device, fence)
}

// SDL_GPUTextureSupportsFormat - Determines whether a texture format is supported for a given type and usage.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsFormat)
func (device *GPUDevice) TextureSupportsFormat(format GPUTextureFormat, typ GPUTextureType, usage GPUTextureUsageFlags) bool {
	return iGPUTextureSupportsFormat(device, format, typ, usage)
}

// SDL_GPUTextureSupportsSampleCount - Determines if a sample count for a texture format is supported.
// (https://wiki.libsdl.org/SDL3/SDL_GPUTextureSupportsSampleCount)
func (device *GPUDevice) TextureSupportsSampleCount(format GPUTextureFormat, sampleCount GPUSampleCount) bool {
	return iGPUTextureSupportsSampleCount(device, format, sampleCount)
}

// Haptic

// SDL_GetHapticID - Get the instance ID of an opened haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticID)
func (haptic *Haptic) ID() (HapticID, error) {
	id := iGetHapticID(haptic)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetHapticName - Get the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticName)
func (haptic *Haptic) Name() (string, error) {
	name := iGetHapticName(haptic)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_CloseHaptic - Close a haptic device previously opened with SDL_OpenHaptic().
// (https://wiki.libsdl.org/SDL3/SDL_CloseHaptic)
func (haptic *Haptic) Close() {
	iCloseHaptic(haptic)
}

// SDL_GetMaxHapticEffects - Get the number of effects a haptic device can store.
// (https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffects)
func (haptic *Haptic) MaxEffects() (int32, error) {
	count := iGetMaxHapticEffects(haptic)
	if count < 0 {
		return count, internal.LastErr()
	}

	return count, nil
}

// SDL_GetMaxHapticEffectsPlaying - Get the number of effects a haptic device can play at the same time.
// (https://wiki.libsdl.org/SDL3/SDL_GetMaxHapticEffectsPlaying)
func (haptic *Haptic) MaxEffectsPlaying() (int32, error) {
	count := iGetMaxHapticEffectsPlaying(haptic)
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// SDL_GetHapticFeatures - Get the haptic device's supported features in bitwise manner.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticFeatures)
func (haptic *Haptic) Features() (uint32, error) {
	mask := iGetHapticFeatures(haptic)
	if mask == 0 {
		return 0, internal.LastErr()
	}

	return mask, nil
}

// SDL_GetNumHapticAxes - Get the number of haptic axes the device has.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumHapticAxes)
func (haptic *Haptic) NumAxes() (int32, error) {
	count := iGetNumHapticAxes(haptic)
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// SDL_HapticEffectSupported - Check to see if an effect is supported by a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_HapticEffectSupported)
func (haptic *Haptic) EffectSupported(effect *HapticEffect) bool {
	return iHapticEffectSupported(haptic, effect)
}

// SDL_CreateHapticEffect - Create a new haptic effect on a specified device.
// (https://wiki.libsdl.org/SDL3/SDL_CreateHapticEffect)
func (haptic *Haptic) CreateEffect(effect *HapticEffect) (int32, error) {
	id := iCreateHapticEffect(haptic, effect)
	if id == -1 {
		return -1, internal.LastErr()
	}

	return id, nil
}

// SDL_UpdateHapticEffect - Update the properties of an effect.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateHapticEffect)
func (haptic *Haptic) UpdateEffect(effect int32, data *HapticEffect) error {
	if !iUpdateHapticEffect(haptic, effect, data) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RunHapticEffect - Run the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_RunHapticEffect)
func (haptic *Haptic) RunEffect(effect int32, iterations uint32) error {
	if !iRunHapticEffect(haptic, effect, iterations) {
		return internal.LastErr()
	}

	return nil
}

// SDL_StopHapticEffect - Stop the haptic effect on its associated haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticEffect)
func (haptic *Haptic) StopEffect(effect int32) error {
	if !iStopHapticEffect(haptic, effect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DestroyHapticEffect - Destroy a haptic effect on the device.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyHapticEffect)
func (haptic *Haptic) DestroyEffect(effect int32) {
	iDestroyHapticEffect(haptic, effect)
}

// SDL_GetHapticEffectStatus - Get the status of the current effect on the specified haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticEffectStatus)
func (haptic *Haptic) EffectStatus(effect int32) bool {
	return iGetHapticEffectStatus(haptic, effect)
}

// SDL_SetHapticGain - Set the global gain of the specified haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_SetHapticGain)
func (haptic *Haptic) SetGain(gain int32) error {
	if !iSetHapticGain(haptic, gain) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetHapticAutocenter - Set the global autocenter of the device.
// (https://wiki.libsdl.org/SDL3/SDL_SetHapticAutocenter)
func (haptic *Haptic) SetAutocenter(autocenter int32) error {
	if !iSetHapticAutocenter(haptic, autocenter) {
		return internal.LastErr()
	}

	return nil
}

// SDL_PauseHaptic - Pause a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_PauseHaptic)
func (haptic *Haptic) Pause() error {
	if !iPauseHaptic(haptic) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ResumeHaptic - Resume a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeHaptic)
func (haptic *Haptic) Resume() error {
	if !iResumeHaptic(haptic) {
		return internal.LastErr()
	}

	return nil
}

// SDL_StopHapticEffects - Stop all the currently playing effects on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticEffects)
func (haptic *Haptic) StopEffects() error {
	if !iStopHapticEffects(haptic) {
		return internal.LastErr()
	}

	return nil
}

// SDL_HapticRumbleSupported - Check whether rumble is supported on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_HapticRumbleSupported)
func (haptic *Haptic) RumbleSupported() bool {
	return iHapticRumbleSupported(haptic)
}

// SDL_InitHapticRumble - Initialize a haptic device for simple rumble playback.
// (https://wiki.libsdl.org/SDL3/SDL_InitHapticRumble)
func (haptic *Haptic) InitRumble() error {
	if !iInitHapticRumble(haptic) {
		return internal.LastErr()
	}

	return nil
}

// SDL_PlayHapticRumble - Run a simple rumble effect on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_PlayHapticRumble)
func (haptic *Haptic) PlayRumble(strength float32, length uint32) error {
	if !iPlayHapticRumble(haptic, strength, length) {
		return internal.LastErr()
	}

	return nil
}

// SDL_StopHapticRumble - Stop the simple rumble on a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_StopHapticRumble)
func (haptic *Haptic) StopRumble() error {
	if !iStopHapticRumble(haptic) {
		return internal.LastErr()
	}

	return nil
}

// FileDialogType

// SDL_ShowFileDialogWithProperties - Create and launch a file dialog with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties)
func (typ FileDialogType) ShowFileDialogWithProperties(callback DialogFileCallback, userdata *byte, props PropertiesID) {
	panic("not implemented")
	//iShowFileDialogWithProperties(typ, callback, userdata, props)
}

// Gamepad

// SDL_GetGamepadMapping - Get the current mapping of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping)
func (gamepad *Gamepad) Mapping() (string, error) {
	ptr := iGetGamepadMapping(gamepad)
	if ptr == 0 {
		return "", internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrString(ptr), nil
}

// SDL_GetGamepadProperties - Get the properties associated with an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties)
func (gamepad *Gamepad) Properties() (PropertiesID, error) {
	props := iGetGamepadProperties(gamepad)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetGamepadID - Get the instance ID of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadID)
func (gamepad *Gamepad) ID() (JoystickID, error) {
	id := iGetGamepadID(gamepad)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetGamepadName - Get the implementation-dependent name for an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadName)
func (gamepad *Gamepad) Name() string {
	return iGetGamepadName(gamepad)
}

// SDL_GetGamepadPath - Get the implementation-dependent path for an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath)
func (gamepad *Gamepad) Path() string {
	return iGetGamepadPath(gamepad)
}

// SDL_GetGamepadType - Get the type of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadType)
func (gamepad *Gamepad) Type() GamepadType {
	return iGetGamepadType(gamepad)
}

// SDL_GetRealGamepadType - Get the type of an opened gamepad, ignoring any mapping override.
// (https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType)
func (gamepad *Gamepad) RealType() GamepadType {
	return iGetRealGamepadType(gamepad)
}

// SDL_GetGamepadPlayerIndex - Get the player index of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex)
func (gamepad *Gamepad) PlayerIndex() int32 {
	return iGetGamepadPlayerIndex(gamepad)
}

// SDL_SetGamepadPlayerIndex - Set the player index of an opened gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex)
func (gamepad *Gamepad) SetPlayerIndex(playerIndex int32) error {
	if !iSetGamepadPlayerIndex(gamepad, playerIndex) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetGamepadVendor - Get the USB vendor ID of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor)
func (gamepad *Gamepad) Vendor() uint16 {
	return iGetGamepadVendor(gamepad)
}

// SDL_GetGamepadProduct - Get the USB product ID of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct)
func (gamepad *Gamepad) Product() uint16 {
	return iGetGamepadProduct(gamepad)
}

// SDL_GetGamepadProductVersion - Get the product version of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion)
func (gamepad *Gamepad) ProductVersion() uint16 {
	return iGetGamepadProductVersion(gamepad)
}

// SDL_GetGamepadFirmwareVersion - Get the firmware version of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion)
func (gamepad *Gamepad) FirmwareVersion() uint16 {
	return iGetGamepadFirmwareVersion(gamepad)
}

// SDL_GetGamepadSerial - Get the serial number of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial)
func (gamepad *Gamepad) Serial() string {
	return iGetGamepadSerial(gamepad)
}

// SDL_GetGamepadSteamHandle - Get the Steam Input handle of an opened gamepad, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle)
func (gamepad *Gamepad) SteamHandle() uint64 {
	return iGetGamepadSteamHandle(gamepad)
}

// SDL_GetGamepadConnectionState - Get the connection state of a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState)
func (gamepad *Gamepad) ConnectionState() (JoystickConnectionState, error) {
	state := iGetGamepadConnectionState(gamepad)
	if state == JOYSTICK_CONNECTION_INVALID {
		return JOYSTICK_CONNECTION_INVALID, internal.LastErr()
	}

	return state, nil
}

// SDL_GetGamepadPowerInfo - Get the battery state of a gamepad.
// Returns power state, percent.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo)
func (gamepad *Gamepad) PowerInfo() (PowerState, int32) {
	var percent int32

	state := iGetGamepadPowerInfo(gamepad, &percent)

	return state, percent
}

// SDL_GamepadConnected - Check if a gamepad has been opened and is currently connected.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadConnected)
func (gamepad *Gamepad) Connected() bool {
	return iGamepadConnected(gamepad)
}

// SDL_GetGamepadJoystick - Get the underlying joystick from a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick)
func (gamepad *Gamepad) Joystick() (*Joystick, error) {
	joystick := iGetGamepadJoystick(gamepad)
	if joystick == nil {
		return nil, internal.LastErr()
	}

	return joystick, nil
}

// SDL_GetGamepadBindings - Get the SDL joystick layer bindings for a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadBindings)
func (gamepad *Gamepad) Bindings() ([]*GamepadBinding, error) {
	var count int32

	ptr := iGetGamepadBindings(gamepad, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*GamepadBinding](ptr, int(count)), nil
}

// SDL_GamepadHasAxis - Query whether a gamepad has a given axis.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis)
func (gamepad *Gamepad) HasAxis(axis GamepadAxis) bool {
	return iGamepadHasAxis(gamepad, axis)
}

// SDL_GetGamepadAxis - Get the current state of an axis control on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis)
func (gamepad *Gamepad) Axis(axis GamepadAxis) int16 {
	return iGetGamepadAxis(gamepad, axis)
}

// SDL_GamepadHasButton - Query whether a gamepad has a given button.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton)
func (gamepad *Gamepad) HasButton(button GamepadButton) bool {
	return iGamepadHasButton(gamepad, button)
}

// SDL_GetGamepadButton - Get the current state of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton)
func (gamepad *Gamepad) Button(button GamepadButton) bool {
	return iGetGamepadButton(gamepad, button)
}

// SDL_GetGamepadButtonLabel - Get the label of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel)
func (gamepad *Gamepad) ButtonLabel(button GamepadButton) GamepadButtonLabel {
	return iGetGamepadButtonLabel(gamepad, button)
}

// SDL_GetNumGamepadTouchpads - Get the number of touchpads on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads)
func (gamepad *Gamepad) NumTouchpads() int32 {
	return iGetNumGamepadTouchpads(gamepad)
}

// SDL_GetNumGamepadTouchpadFingers - Get the number of supported simultaneous fingers on a touchpad on a game gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers)
func (gamepad *Gamepad) NumTouchpadFingers(touchpad int32) int32 {
	return iGetNumGamepadTouchpadFingers(gamepad, touchpad)
}

// SDL_GetGamepadTouchpadFinger - Get the current state of a finger on a touchpad on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger)
func (gamepad *Gamepad) TouchpadFinger(touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) bool {
	panic("not implemented")
	return iGetGamepadTouchpadFinger(gamepad, touchpad, finger, down, x, y, pressure)
}

// SDL_GamepadHasSensor - Return whether a gamepad has a particular sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor)
func (gamepad *Gamepad) HasSensor(typ SensorType) bool {
	return iGamepadHasSensor(gamepad, typ)
}

// SDL_SetGamepadSensorEnabled - Set whether data reporting for a gamepad sensor is enabled.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled)
func (gamepad *Gamepad) SetSensorEnabled(typ SensorType, enabled bool) error {
	if !iSetGamepadSensorEnabled(gamepad, typ, enabled) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GamepadSensorEnabled - Query whether sensor data reporting is enabled for a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled)
func (gamepad *Gamepad) SensorEnabled(typ SensorType) bool {
	return iGamepadSensorEnabled(gamepad, typ)
}

// SDL_GetGamepadSensorDataRate - Get the data rate (number of events per second) of a gamepad sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate)
func (gamepad *Gamepad) SensorDataRate(typ SensorType) float32 {
	return iGetGamepadSensorDataRate(gamepad, typ)
}

// SDL_GetGamepadSensorData - Get the current state of a gamepad sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData)
func (gamepad *Gamepad) SensorData(typ SensorType, data *float32, num_values int32) bool {
	panic("not implemented")
	return iGetGamepadSensorData(gamepad, typ, data, num_values)
}

// SDL_RumbleGamepad - Start a rumble effect on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad)
func (gamepad *Gamepad) Rumble(lowFrequency, highFrequency uint16, durationMS uint32) error {
	if !iRumbleGamepad(gamepad, lowFrequency, highFrequency, durationMS) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RumbleGamepadTriggers - Start a rumble effect in the gamepad's triggers.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers)
func (gamepad *Gamepad) RumbleTriggers(left, right uint16, durationMS uint32) error {
	if !iRumbleGamepadTriggers(gamepad, left, right, durationMS) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetGamepadLED - Update a gamepad's LED color.
// (https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED)
func (gamepad *Gamepad) SetLED(red, green, blue uint8) error {
	if !iSetGamepadLED(gamepad, red, green, blue) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SendGamepadEffect - Send a gamepad specific effect packet.
// (https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect)
func (gamepad *Gamepad) SendEffect(data []byte) error {
	if !iSendGamepadEffect(gamepad, uintptr(unsafe.Pointer(unsafe.SliceData(data))), int32(len(data))) {
		return internal.LastErr()
	}
	runtime.KeepAlive(data)

	return nil
}

// SDL_CloseGamepad - Close a gamepad previously opened with SDL_OpenGamepad().
// (https://wiki.libsdl.org/SDL3/SDL_CloseGamepad)
func (gamepad *Gamepad) Close() {
	iCloseGamepad(gamepad)
}

// SDL_GetGamepadAppleSFSymbolsNameForButton - Return the sfSymbolsName for a given button on a gamepad on Apple platforms.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton)
func (gamepad *Gamepad) AppleSFSymbolsNameForButton(button GamepadButton) string {
	panic("not implemented")
	return iGetGamepadAppleSFSymbolsNameForButton(gamepad, button)
}

// SDL_GetGamepadAppleSFSymbolsNameForAxis - Return the sfSymbolsName for a given axis on a gamepad on Apple platforms.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis)
func (gamepad *Gamepad) AppleSFSymbolsNameForAxis(axis GamepadAxis) string {
	panic("not implemented")
	return iGetGamepadAppleSFSymbolsNameForAxis(gamepad, axis)
}

// Keymod

// SDL_SetModState - Set the current key modifier state for the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_SetModState)
func (modstate Keymod) SetModState() {
	panic("not implemented")
	iSetModState(modstate)
}

// GPUCopyPass

// SDL_UploadToGPUTexture - Uploads data from a transfer buffer to a texture.
// (https://wiki.libsdl.org/SDL3/SDL_UploadToGPUTexture)
func (copy_pass *GPUCopyPass) UploadToGPUTexture(source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	iUploadToGPUTexture(copy_pass, source, destination, cycle)
}

// SDL_UploadToGPUBuffer - Uploads data from a transfer buffer to a buffer.
// (https://wiki.libsdl.org/SDL3/SDL_UploadToGPUBuffer)
func (copy_pass *GPUCopyPass) UploadToGPUBuffer(source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	iUploadToGPUBuffer(copy_pass, source, destination, cycle)
}

// SDL_CopyGPUTextureToTexture - Performs a texture-to-texture copy.
// (https://wiki.libsdl.org/SDL3/SDL_CopyGPUTextureToTexture)
func (copy_pass *GPUCopyPass) CopyGPUTextureToTexture(source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool) {
	panic("not implemented")
	iCopyGPUTextureToTexture(copy_pass, source, destination, w, h, d, cycle)
}

// SDL_CopyGPUBufferToBuffer - Performs a buffer-to-buffer copy.
// (https://wiki.libsdl.org/SDL3/SDL_CopyGPUBufferToBuffer)
func (copy_pass *GPUCopyPass) CopyGPUBufferToBuffer(source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {
	panic("not implemented")
	iCopyGPUBufferToBuffer(copy_pass, source, destination, size, cycle)
}

// SDL_DownloadFromGPUTexture - Copies data from a texture to a transfer buffer on the GPU timeline.
// (https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUTexture)
func (copy_pass *GPUCopyPass) DownloadFromGPUTexture(source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	panic("not implemented")
	iDownloadFromGPUTexture(copy_pass, source, destination)
}

// SDL_DownloadFromGPUBuffer - Copies data from a buffer to a transfer buffer on the GPU timeline.
// (https://wiki.libsdl.org/SDL3/SDL_DownloadFromGPUBuffer)
func (copy_pass *GPUCopyPass) DownloadFromGPUBuffer(source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	panic("not implemented")
	iDownloadFromGPUBuffer(copy_pass, source, destination)
}

// SDL_EndGPUCopyPass - Ends the current copy pass.
// (https://wiki.libsdl.org/SDL3/SDL_EndGPUCopyPass)
func (copy_pass *GPUCopyPass) End() {
	iEndGPUCopyPass(copy_pass)
}

// HapticID

// SDL_GetHapticNameForID - Get the implementation dependent name of a haptic device.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticNameForID)
func (id HapticID) HapticName() (string, error) {
	name := iGetHapticNameForID(id)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_OpenHaptic - Open a haptic device for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHaptic)
func (id HapticID) OpenHaptic() (*Haptic, error) {
	device := iOpenHaptic(id)
	if device == nil {
		return nil, internal.LastErr()
	}

	return device, nil
}

// SDL_GetHapticFromID - Get the SDL_Haptic associated with an instance ID, if it has been opened.
// (https://wiki.libsdl.org/SDL3/SDL_GetHapticFromID)
func (id HapticID) Haptic() (*Haptic, error) {
	haptic := iGetHapticFromID(id)
	if haptic == nil {
		return nil, internal.LastErr()
	}

	return haptic, nil
}

// Renderer

// SDL_GetRenderWindow - Get the window associated with a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow)
func (renderer *Renderer) Window() (*Window, error) {
	window := iGetRenderWindow(renderer)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_GetRendererName - Get the name of a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererName)
func (renderer *Renderer) Name() (string, error) {
	name := iGetRendererName(renderer)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetRendererProperties - Get the properties associated with a renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties)
func (renderer *Renderer) Properties() PropertiesID {
	return iGetRendererProperties(renderer)
}

// SDL_GetRenderOutputSize - Get the output size in pixels of a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize)
func (renderer *Renderer) RenderOutputSize() (int32, int32, error) {
	var w, h int32
	if !iGetRenderOutputSize(renderer, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_GetCurrentRenderOutputSize - Get the current output size in pixels of a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize)
func (renderer *Renderer) CurrentOutputSize() (int32, int32, error) {
	var w, h int32
	if !iGetCurrentRenderOutputSize(renderer, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_CreateTexture - Create a texture for a rendering context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTexture)
func (renderer *Renderer) CreateTexture(format PixelFormat, access TextureAccess, w, h int) (*Texture, error) {
	texture := iCreateTexture(renderer, format, access, int32(w), int32(h))
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// SDL_CreateTextureFromSurface - Create a texture from an existing surface.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface)
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	texture := iCreateTextureFromSurface(renderer, surface)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// SDL_CreateTextureWithProperties - Create a texture for a rendering context with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties)
func (renderer *Renderer) CreateTextureWithProperties(props PropertiesID) (*Texture, error) {
	texture := iCreateTextureWithProperties(renderer, props)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// SDL_SetRenderTarget - Set a texture as the current rendering target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget)
func (renderer *Renderer) SetRenderTarget(texture *Texture) error {
	if !iSetRenderTarget(renderer, texture) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderTarget - Get the current render target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget)
func (renderer *Renderer) RenderTarget() *Texture {
	return iGetRenderTarget(renderer)
}

// SDL_SetRenderLogicalPresentation - Set a device independent resolution and presentation mode for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation)
func (renderer *Renderer) SetLogicalPresentation(w int32, h int32, mode RendererLogicalPresentation) error {
	if !iSetRenderLogicalPresentation(renderer, w, h, mode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderLogicalPresentation - Get device independent resolution and presentation mode for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation)
func (renderer *Renderer) LogicalPresentation() (int32, int32, RendererLogicalPresentation, error) {
	var w, h int32
	var mode RendererLogicalPresentation

	if !iGetRenderLogicalPresentation(renderer, &w, &h, &mode) {
		return 0, 0, 0, internal.LastErr()
	}
	return w, h, mode, nil
}

// SDL_GetRenderLogicalPresentationRect - Get the final presentation rectangle for rendering.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect)
func (renderer *Renderer) LogicalPresentationRect() (*FRect, error) {
	var rect FRect

	if !iGetRenderLogicalPresentationRect(renderer, &rect) {
		return nil, internal.LastErr()
	}

	return &rect, nil
}

// SDL_RenderCoordinatesFromWindow - Get a point in render coordinates when given a point in window coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow)
func (renderer *Renderer) RenderCoordinatesFromWindow(window_x float32, window_y float32, x *float32, y *float32) bool {
	panic("not implemented")
	return iRenderCoordinatesFromWindow(renderer, window_x, window_y, x, y)
}

// SDL_RenderCoordinatesToWindow - Get a point in window coordinates when given a point in render coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow)
func (renderer *Renderer) RenderCoordinatesToWindow(x float32, y float32, window_x *float32, window_y *float32) bool {
	panic("not implemented")
	return iRenderCoordinatesToWindow(renderer, x, y, window_x, window_y)
}

// SDL_ConvertEventToRenderCoordinates - Convert the coordinates in an event to render coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates)
func (renderer *Renderer) ConvertEventToRenderCoordinates(event *Event) bool {
	panic("not implemented")
	return iConvertEventToRenderCoordinates(renderer, event)
}

// SDL_SetRenderViewport - Set the drawing area for rendering on the current target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport)
func (renderer *Renderer) SetViewport(rect *Rect) error {
	if !iSetRenderViewport(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderViewport - Get the drawing area for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport)
func (renderer *Renderer) Viewport() (Rect, error) {
	var r Rect

	if !iGetRenderViewport(renderer, &r) {
		return r, internal.LastErr()
	}

	return r, nil
}

// SDL_RenderViewportSet - Return whether an explicit rectangle was set as the viewport.
// (https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet)
func (renderer *Renderer) ViewportSet() bool {
	return iRenderViewportSet(renderer)
}

// SDL_GetRenderSafeArea - Get the safe area for rendering within the current viewport.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea)
func (renderer *Renderer) SafeArea() (Rect, error) {
	var r Rect

	if !iGetRenderSafeArea(renderer, &r) {
		return r, internal.LastErr()
	}

	return r, nil
}

// SDL_SetRenderClipRect - Set the clip rectangle for rendering on the specified target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect)
func (renderer *Renderer) SetClipRect(rect *Rect) error {
	if !iSetRenderClipRect(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderClipRect - Get the clip rectangle for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect)
func (renderer *Renderer) ClipRect() (Rect, error) {
	var r Rect

	if !iGetRenderClipRect(renderer, &r) {
		return r, internal.LastErr()
	}

	return r, nil
}

// SDL_RenderClipEnabled - Get whether clipping is enabled on the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled)
func (renderer *Renderer) ClipEnabled() error {
	if !iRenderClipEnabled(renderer) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetRenderScale - Set the drawing scale for rendering on the current target.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderScale)
func (renderer *Renderer) SetScale(scaleX, scaleY float32) error {
	if !iSetRenderScale(renderer, scaleX, scaleY) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderScale - Get the drawing scale for the current target.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderScale)
func (renderer *Renderer) Scale() (float32, float32, error) {
	var scaleX, scaleY float32

	if !iGetRenderScale(renderer, &scaleX, &scaleY) {
		return 0, 0, internal.LastErr()
	}

	return scaleX, scaleY, nil
}

// SDL_SetRenderDrawColor - Set the color used for drawing operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor)
func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) error {
	if !iSetRenderDrawColor(renderer, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetRenderDrawColorFloat - Set the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat)
func (renderer *Renderer) SetDrawColorFloat(r, g, b, a float32) error {
	if !iSetRenderDrawColorFloat(renderer, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderDrawColor - Get the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor)
func (renderer *Renderer) DrawColor() (Color, error) {
	var clr Color

	if !iGetRenderDrawColor(renderer, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

// SDL_GetRenderDrawColorFloat - Get the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat)
func (renderer *Renderer) DrawColorFloat() (FColor, error) {
	var clr FColor

	if !iGetRenderDrawColorFloat(renderer, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

// SDL_SetRenderColorScale - Set the color scale used for render operations.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale)
func (renderer *Renderer) SetColorScale(scale float32) error {
	if !iSetRenderColorScale(renderer, scale) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderColorScale - Get the color scale used for render operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale)
func (renderer *Renderer) ColorScale() (float32, error) {
	var scale float32

	if !iGetRenderColorScale(renderer, &scale) {
		return 0, internal.LastErr()
	}

	return scale, nil
}

// SDL_SetRenderDrawBlendMode - Set the blend mode used for drawing operations (Fill and Line).
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode)
func (renderer *Renderer) SetDrawBlendMode(blendMode BlendMode) error {
	if !iSetRenderDrawBlendMode(renderer, blendMode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderDrawBlendMode - Get the blend mode used for drawing operations.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode)
func (renderer *Renderer) DrawBlendMode() (BlendMode, error) {
	var mode BlendMode

	if !iGetRenderDrawBlendMode(renderer, &mode) {
		return mode, internal.LastErr()
	}

	return mode, nil
}

// SDL_RenderClear - Clear the current rendering target with the drawing color.
// (https://wiki.libsdl.org/SDL3/SDL_RenderClear)
func (renderer *Renderer) Clear() error {
	if !iRenderClear(renderer) {
		return internal.LastErr()
	}
	return nil
}

// SDL_RenderPoint - Draw a point on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPoint)
func (renderer *Renderer) RenderPoint(x, y float32) error {
	if !iRenderPoint(renderer, x, y) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderPoints - Draw multiple points on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPoints)
func (renderer *Renderer) RenderPoints(points []FPoint) error {
	if !iRenderPoints(renderer, unsafe.SliceData(points), int32(len(points))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderLine - Draw a line on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderLine)
func (renderer *Renderer) RenderLine(x1 float32, y1 float32, x2 float32, y2 float32) error {
	if !iRenderLine(renderer, x1, y1, x2, y2) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderLines - Draw a series of connected lines on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderLines)
func (renderer *Renderer) RenderLines(points []FPoint) error {
	if !iRenderLines(renderer, unsafe.SliceData(points), int32(len(points))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderRect - Draw a rectangle on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderRect)
func (renderer *Renderer) RenderRect(rect *FRect) error {
	if !iRenderRect(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderRects - Draw some number of rectangles on the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderRects)
func (renderer *Renderer) RenderRects(rects []FRect) error {
	if !iRenderRects(renderer, unsafe.SliceData(rects), int32(len(rects))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderFillRect - Fill a rectangle on the current rendering target with the drawing color at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderFillRect)
func (renderer *Renderer) RenderFillRect(rect *FRect) error {
	if !iRenderFillRect(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderFillRects - Fill some number of rectangles on the current rendering target with the drawing color at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderFillRects)
func (renderer *Renderer) RenderFillRects(rects []FRect) error {
	if !iRenderFillRects(renderer, unsafe.SliceData(rects), int32(len(rects))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderTexture - Copy a portion of the texture to the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTexture)
func (renderer *Renderer) RenderTexture(texture *Texture, srcrect *FRect, dstrect *FRect) error {
	if !iRenderTexture(renderer, texture, srcrect, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderTextureRotated - Copy a portion of the source texture to the current rendering target, with rotation and flipping, at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated)
func (renderer *Renderer) RenderTextureRotated(texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) error {
	if !iRenderTextureRotated(renderer, texture, srcrect, dstrect, angle, center, flip) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderTextureAffine - Copy a portion of the source texture to the current rendering target, with affine transform, at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine)
func (renderer *Renderer) RenderTextureAffine(texture *Texture, srcrect *FRect, origin, right, down *FPoint) error {
	if !iRenderTextureAffine(renderer, texture, srcrect, origin, right, down) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderTextureTiled - Tile a portion of the texture to the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled)
func (renderer *Renderer) RenderTextureTiled(texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) error {
	if !iRenderTextureTiled(renderer, texture, srcrect, scale, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderTexture9Grid - Perform a scaled copy using the 9-grid algorithm to the current rendering target at subpixel precision.
// (https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid)
func (renderer *Renderer) RenderTexture9Grid(texture *Texture, srcrect *FRect, leftWidth, rightWidth, topHeight, bottomHeight, scale float32, dstrect *FRect) error {
	if !iRenderTexture9Grid(renderer, texture, srcrect, leftWidth, rightWidth, topHeight, bottomHeight, scale, dstrect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderGeometry - Render a list of triangles, optionally using a texture and indices into the vertex array Color and alpha modulation is done per vertex (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
// (https://wiki.libsdl.org/SDL3/SDL_RenderGeometry)
func (renderer *Renderer) RenderGeometry(texture *Texture, vertices []Vertex, indices []int32) error {
	if !iRenderGeometry(renderer, texture, unsafe.SliceData(vertices), int32(len(vertices)), unsafe.SliceData(indices), int32(len(indices))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderGeometryRaw - Render a list of triangles, optionally using a texture and indices into the vertex arrays Color and alpha modulation is done per vertex (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
// (https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw)
func (renderer *Renderer) RenderGeometryRaw(texture *Texture, xy []float32, xyStride int32, color []FColor, colorStride int32, uv []float32, uvStride int32, indices []uint16) error {
	if !iRenderGeometryRaw(
		renderer,
		texture,
		unsafe.SliceData(xy), xyStride,
		unsafe.SliceData(color), colorStride,
		unsafe.SliceData(uv), uvStride, int32(len(uv)/2),
		uintptr(unsafe.Pointer(unsafe.SliceData(indices))), int32(len(indices)), 2,
	) {
		return internal.LastErr()
	}
	runtime.KeepAlive(xy)
	runtime.KeepAlive(color)
	runtime.KeepAlive(uv)
	runtime.KeepAlive(indices)

	return nil
}

// SDL_RenderReadPixels - Read pixels from the current rendering target.
// (https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels)
func (renderer *Renderer) ReadPixels(rect *Rect) (*Surface, error) {
	surface := iRenderReadPixels(renderer, rect)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// SDL_RenderPresent - Update the screen with any rendering performed since the previous call.
// (https://wiki.libsdl.org/SDL3/SDL_RenderPresent)
func (renderer *Renderer) Present() error {
	if !iRenderPresent(renderer) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DestroyRenderer - Destroy the rendering context for a window and free all associated textures.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer)
func (renderer *Renderer) Destroy() {
	iDestroyRenderer(renderer)
}

// SDL_FlushRenderer - Force the rendering context to flush any pending commands and state.
// (https://wiki.libsdl.org/SDL3/SDL_FlushRenderer)
func (renderer *Renderer) Flush() error {
	if !iFlushRenderer(renderer) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderMetalLayer - Get the CAMetalLayer associated with the given Metal renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer)
func (renderer *Renderer) RenderMetalLayer() *byte {
	panic("not implemented")
	//return iGetRenderMetalLayer(renderer)
}

// SDL_GetRenderMetalCommandEncoder - Get the Metal command encoder for the current frame.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder)
func (renderer *Renderer) RenderMetalCommandEncoder() *byte {
	panic("not implemented")
	//return iGetRenderMetalCommandEncoder(renderer)
}

// SDL_AddVulkanRenderSemaphores - Add a set of synchronization semaphores for the current frame.
// (https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores)
func (renderer *Renderer) AddVulkanRenderSemaphores(wait_stage_mask uint32, wait_semaphore int64, signal_semaphore int64) bool {
	panic("not implemented")
	return iAddVulkanRenderSemaphores(renderer, wait_stage_mask, wait_semaphore, signal_semaphore)
}

// SDL_SetRenderVSync - Toggle VSync of the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync)
func (renderer *Renderer) SetVSync(vsync int32) error {
	if !iSetRenderVSync(renderer, vsync) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetRenderVSync - Get VSync of the given renderer.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync)
func (renderer *Renderer) VSync() (int32, error) {
	var vsync int32

	if !iGetRenderVSync(renderer, &vsync) {
		return 0, internal.LastErr()
	}

	return vsync, nil
}

// SDL_RenderDebugText - Draw debug text to an SDL_Renderer.
// (https://wiki.libsdl.org/SDL3/SDL_RenderDebugText)
func (renderer *Renderer) DebugText(x float32, y float32, str string) error {
	if !iRenderDebugText(renderer, x, y, str) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RenderDebugText - Draw debug text to an SDL_Renderer.
// (https://wiki.libsdl.org/SDL3/SDL_RenderDebugText)
func (renderer *Renderer) DebugTextFormat(x float32, y float32, format string, values ...any) bool {
	return iRenderDebugText(renderer, x, y, fmt.Sprintf(format, values...))
}

// DateFormat

// SDL_GetDateTimeLocalePreferences - Gets the current preferred date and time format for the system locale.
// (https://wiki.libsdl.org/SDL3/SDL_GetDateTimeLocalePreferences)
func (dateFormat *DateFormat) DateTimeLocalePreferences(timeFormat *TimeFormat) bool {
	panic("not implemented")
	return iGetDateTimeLocalePreferences(dateFormat, timeFormat)
}

// AsyncIO

// SDL_GetAsyncIOSize - Use this function to get the size of the data stream in an SDL_AsyncIO.
// (https://wiki.libsdl.org/SDL3/SDL_GetAsyncIOSize)
func (asyncio *AsyncIO) Size() int64 {
	panic("not implemented")
	return iGetAsyncIOSize(asyncio)
}

// SDL_ReadAsyncIO - Start an async read.
// (https://wiki.libsdl.org/SDL3/SDL_ReadAsyncIO)
func (asyncio *AsyncIO) Read(ptr *byte, offset uint64, size uint64, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iReadAsyncIO(asyncio, ptr, offset, size, queue, userdata)
}

// SDL_WriteAsyncIO - Start an async write.
// (https://wiki.libsdl.org/SDL3/SDL_WriteAsyncIO)
func (asyncio *AsyncIO) Write(ptr *byte, offset uint64, size uint64, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iWriteAsyncIO(asyncio, ptr, offset, size, queue, userdata)
}

// SDL_CloseAsyncIO - Close and free any allocated resources for an async I/O object.
// (https://wiki.libsdl.org/SDL3/SDL_CloseAsyncIO)
func (asyncio *AsyncIO) Close(flush bool, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iCloseAsyncIO(asyncio, flush, queue, userdata)
}

// InitState

// SDL_ShouldInit - Return whether initialization should be done.
// (https://wiki.libsdl.org/SDL3/SDL_ShouldInit)
func (state *InitState) ShouldInit() bool {
	panic("not implemented")
	return iShouldInit(state)
}

// SDL_ShouldQuit - Return whether cleanup should be done.
// (https://wiki.libsdl.org/SDL3/SDL_ShouldQuit)
func (state *InitState) ShouldQuit() bool {
	panic("not implemented")
	return iShouldQuit(state)
}

// SDL_SetInitialized - Finish an initialization state transition.
// (https://wiki.libsdl.org/SDL3/SDL_SetInitialized)
func (state *InitState) SetInitialized(initialized bool) {
	panic("not implemented")
	iSetInitialized(state, initialized)
}

// CameraID

// SDL_GetCameraSupportedFormats - Get the list of native formats/sizes a camera supports.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraSupportedFormats)
func (instance_id CameraID) CameraSupportedFormats(count *int32) **CameraSpec {
	panic("not implemented")
	//return iGetCameraSupportedFormats(instance_id, count)
}

// SDL_GetCameraName - Get the human-readable device name for a camera.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraName)
func (instance_id CameraID) CameraName() string {
	panic("not implemented")
	return iGetCameraName(instance_id)
}

// SDL_GetCameraPosition - Get the position of the camera in relation to the system.
// (https://wiki.libsdl.org/SDL3/SDL_GetCameraPosition)
func (instance_id CameraID) CameraPosition() CameraPosition {
	panic("not implemented")
	return iGetCameraPosition(instance_id)
}

// SDL_OpenCamera - Open a video recording device (a "camera").
// (https://wiki.libsdl.org/SDL3/SDL_OpenCamera)
func (instance_id CameraID) OpenCamera(spec *CameraSpec) *Camera {
	panic("not implemented")
	return iOpenCamera(instance_id, spec)
}

// DisplayID

// SDL_GetDisplayProperties - Get the properties associated with a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties)
func (displayID DisplayID) Properties() (PropertiesID, error) {
	props := iGetDisplayProperties(displayID)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetDisplayName - Get the name of a display in UTF-8 encoding.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayName)
func (displayID DisplayID) Name() (string, error) {
	name := iGetDisplayName(displayID)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetDisplayBounds - Get the desktop area represented by a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds)
func (displayID DisplayID) Bounds() (*Rect, error) {
	var r Rect

	if !iGetDisplayBounds(displayID, &r) {
		return nil, internal.LastErr()
	}

	return &r, nil
}

// SDL_GetDisplayUsableBounds - Get the usable desktop area represented by a display, in screen coordinates.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds)
func (displayID DisplayID) UsableBounds() (*Rect, error) {
	var r Rect

	if !iGetDisplayUsableBounds(displayID, &r) {
		return nil, internal.LastErr()
	}

	return &r, nil
}

// SDL_GetNaturalDisplayOrientation - Get the orientation of a display when it is unrotated.
// (https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation)
func (displayID DisplayID) NaturalDisplayOrientation() DisplayOrientation {
	return iGetNaturalDisplayOrientation(displayID)
}

// SDL_GetCurrentDisplayOrientation - Get the orientation of a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation)
func (displayID DisplayID) CurrentDisplayOrientation() DisplayOrientation {
	return iGetCurrentDisplayOrientation(displayID)
}

// SDL_GetDisplayContentScale - Get the content scale of a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale)
func (displayID DisplayID) ContentScale() (float32, error) {
	scale := iGetDisplayContentScale(displayID)
	if scale == 0 {
		return 0, internal.LastErr()
	}

	return scale, nil
}

// SDL_GetFullscreenDisplayModes - Get a list of fullscreen display modes available on a display.
// (https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes)
func (displayID DisplayID) FullscreenDisplayModes() ([]*DisplayMode, error) {
	var count int32

	ptr := iGetFullscreenDisplayModes(displayID, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*DisplayMode](ptr, int(count)), nil
}

// SDL_GetClosestFullscreenDisplayMode - Get the closest match to the requested display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode)
func (displayID DisplayID) ClosestFullscreenDisplayMode(w, h int32, refreshRate float32, includeHighDensityModes bool) (*DisplayMode, error) {
	var mode DisplayMode

	if !iGetClosestFullscreenDisplayMode(displayID, w, h, refreshRate, includeHighDensityModes, &mode) {
		return nil, internal.LastErr()
	}

	return &mode, nil
}

// SDL_GetDesktopDisplayMode - Get information about the desktop's display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode)
func (displayID DisplayID) DesktopDisplayMode() (*DisplayMode, error) {
	mode := iGetDesktopDisplayMode(displayID)
	if mode == nil {
		return nil, internal.LastErr()
	}

	return mode, nil
}

// SDL_GetCurrentDisplayMode - Get information about the current display mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode)
func (displayID DisplayID) CurrentDisplayMode() (*DisplayMode, error) {
	mode := iGetCurrentDisplayMode(displayID)
	if mode == nil {
		return nil, internal.LastErr()
	}

	return mode, nil
}

// KeyboardID

// SDL_GetKeyboardNameForID - Get the name of a keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID)
func (id KeyboardID) Name() (string, error) {
	name := iGetKeyboardNameForID(id)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// MouseID

// SDL_GetMouseNameForID - Get the name of a mouse.
// (https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID)
func (id MouseID) Name() (string, error) {
	name := iGetMouseNameForID(id)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// EventFilter

// SDL_GetEventFilter - Query the current event filter.
// (https://wiki.libsdl.org/SDL3/SDL_GetEventFilter)
func (filter *EventFilter) Get(userdata **byte) bool {
	panic("not implemented")
	//return iGetEventFilter(filter, userdata)
}

// AudioStream

// SDL_UnbindAudioStream - Unbind a single audio stream from its audio device.
// (https://wiki.libsdl.org/SDL3/SDL_UnbindAudioStream)
func (stream *AudioStream) Unbind() {
	iUnbindAudioStream(stream)
}

// SDL_GetAudioStreamDevice - Query an audio stream for its currently-bound device.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamDevice)
func (stream *AudioStream) Device() AudioDeviceID {
	return iGetAudioStreamDevice(stream)
}

// SDL_GetAudioStreamProperties - Get the properties associated with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamProperties)
func (stream *AudioStream) Properties() (PropertiesID, error) {
	props := iGetAudioStreamProperties(stream)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetAudioStreamFormat - Query the current format of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFormat)
func (stream *AudioStream) Format(src, dst *AudioSpec) error {
	if !iGetAudioStreamFormat(stream, src, dst) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetAudioStreamFormat - Change the input and output formats of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFormat)
func (stream *AudioStream) SetFormat(src, dst *AudioSpec) error {
	if !iSetAudioStreamFormat(stream, src, dst) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetAudioStreamFrequencyRatio - Get the frequency ratio of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamFrequencyRatio)
func (stream *AudioStream) FrequencyRatio() (float32, error) {
	ratio := iGetAudioStreamFrequencyRatio(stream)
	if ratio == 0 {
		return 0, internal.LastErr()
	}

	return ratio, nil
}

// SDL_SetAudioStreamFrequencyRatio - Change the frequency ratio of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamFrequencyRatio)
func (stream *AudioStream) SetFrequencyRatio(ratio float32) error {
	if !iSetAudioStreamFrequencyRatio(stream, ratio) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetAudioStreamGain - Get the gain of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamGain)
func (stream *AudioStream) Gain() (float32, error) {
	gain := iGetAudioStreamGain(stream)
	if gain == -1 {
		return -1, internal.LastErr()
	}

	return gain, nil
}

// SDL_SetAudioStreamGain - Change the gain of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGain)
func (stream *AudioStream) SetGain(gain float32) error {
	if !iSetAudioStreamGain(stream, gain) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetAudioStreamInputChannelMap - Get the current input channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamInputChannelMap)
func (stream *AudioStream) InputChannelMap() []int {
	var count int32

	ptr := iGetAudioStreamInputChannelMap(stream, &count)
	if ptr == 0 {
		return nil
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[int](ptr, int(count))
}

// SDL_GetAudioStreamOutputChannelMap - Get the current output channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamOutputChannelMap)
func (stream *AudioStream) OutputChannelMap() []int {
	var count int32

	ptr := iGetAudioStreamOutputChannelMap(stream, &count)
	if ptr == 0 {
		return nil
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[int](ptr, int(count))
}

// SDL_SetAudioStreamInputChannelMap - Set the current input channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamInputChannelMap)
func (stream *AudioStream) SetInputChannelMap(chmap []int32) error {
	if !iSetAudioStreamInputChannelMap(stream, unsafe.SliceData(chmap), int32(len(chmap))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetAudioStreamOutputChannelMap - Set the current output channel map of an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamOutputChannelMap)
func (stream *AudioStream) SetOutputChannelMap(chmap []int32) error {
	if !iSetAudioStreamOutputChannelMap(stream, unsafe.SliceData(chmap), int32(len(chmap))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_PutAudioStreamData - Add data to the stream.
// (https://wiki.libsdl.org/SDL3/SDL_PutAudioStreamData)
func (stream *AudioStream) PutData(buf []byte) error {
	if !iPutAudioStreamData(stream, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), int32(len(buf))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetAudioStreamData - Get converted/resampled data from the stream.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamData)
func (stream *AudioStream) Data(buf []byte) (int32, error) {
	count := iGetAudioStreamData(stream, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), int32(len(buf)))
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// SDL_GetAudioStreamAvailable - Get the number of converted/resampled bytes available.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamAvailable)
func (stream *AudioStream) Available() (int32, error) {
	count := iGetAudioStreamAvailable(stream)
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// SDL_GetAudioStreamQueued - Get the number of bytes currently queued.
// (https://wiki.libsdl.org/SDL3/SDL_GetAudioStreamQueued)
func (stream *AudioStream) Queued() (int32, error) {
	count := iGetAudioStreamQueued(stream)
	if count == -1 {
		return -1, internal.LastErr()
	}

	return count, nil
}

// SDL_FlushAudioStream - Tell the stream that you're done sending data, and anything being buffered should be converted/resampled and made available immediately.
// (https://wiki.libsdl.org/SDL3/SDL_FlushAudioStream)
func (stream *AudioStream) Flush() error {
	if !iFlushAudioStream(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ClearAudioStream - Clear any pending data in the stream.
// (https://wiki.libsdl.org/SDL3/SDL_ClearAudioStream)
func (stream *AudioStream) Clear() error {
	if !iClearAudioStream(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_PauseAudioStreamDevice - Use this function to pause audio playback on the audio device associated with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_PauseAudioStreamDevice)
func (stream *AudioStream) PauseDevice() error {
	if !iPauseAudioStreamDevice(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ResumeAudioStreamDevice - Use this function to unpause audio playback on the audio device associated with an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_ResumeAudioStreamDevice)
func (stream *AudioStream) ResumeDevice() error {
	if !iResumeAudioStreamDevice(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_AudioStreamDevicePaused - Use this function to query if an audio device associated with a stream is paused.
// (https://wiki.libsdl.org/SDL3/SDL_AudioStreamDevicePaused)
func (stream *AudioStream) DevicePaused() bool {
	return iAudioStreamDevicePaused(stream)
}

// SDL_LockAudioStream - Lock an audio stream for serialized access.
// (https://wiki.libsdl.org/SDL3/SDL_LockAudioStream)
func (stream *AudioStream) Lock() error {
	if !iLockAudioStream(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_UnlockAudioStream - Unlock an audio stream for serialized access.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockAudioStream)
func (stream *AudioStream) Unlock() error {
	if !iUnlockAudioStream(stream) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetAudioStreamGetCallback - Set a callback that runs when data is requested from an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamGetCallback)
func (stream *AudioStream) SetGetCallback(callback AudioStreamCallback, userdata uintptr) bool {
	return iSetAudioStreamGetCallback(stream, callback, userdata)
}

// SDL_SetAudioStreamPutCallback - Set a callback that runs when data is added to an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_SetAudioStreamPutCallback)
func (stream *AudioStream) SetPutCallback(callback AudioStreamCallback, userdata uintptr) bool {
	return iSetAudioStreamPutCallback(stream, callback, userdata)
}

// SDL_DestroyAudioStream - Free an audio stream.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyAudioStream)
func (stream *AudioStream) Destroy() {
	iDestroyAudioStream(stream)
}

// FRect

// SDL_HasRectIntersectionFloat - Determine whether two rectangles intersect with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_HasRectIntersectionFloat)
func (A *FRect) HasRectIntersectionFloat(B *FRect) bool {
	panic("not implemented")
	return iHasRectIntersectionFloat(A, B)
}

// SDL_GetRectIntersectionFloat - Calculate the intersection of two rectangles with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectIntersectionFloat)
func (A *FRect) RectIntersectionFloat(B *FRect, result *FRect) bool {
	panic("not implemented")
	return iGetRectIntersectionFloat(A, B, result)
}

// SDL_GetRectUnionFloat - Calculate the union of two rectangles with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectUnionFloat)
func (A *FRect) RectUnionFloat(B *FRect, result *FRect) bool {
	panic("not implemented")
	return iGetRectUnionFloat(A, B, result)
}

// SDL_GetRectAndLineIntersectionFloat - Calculate the intersection of a rectangle and line segment with float precision.
// (https://wiki.libsdl.org/SDL3/SDL_GetRectAndLineIntersectionFloat)
func (rect *FRect) RectAndLineIntersectionFloat(X1 *float32, Y1 *float32, X2 *float32, Y2 *float32) bool {
	panic("not implemented")
	return iGetRectAndLineIntersectionFloat(rect, X1, Y1, X2, Y2)
}

// LogOutputFunction

// SDL_SetLogOutputFunction - Replace the default log output function with one of your own.
// (https://wiki.libsdl.org/SDL3/SDL_SetLogOutputFunction)
func (callback LogOutputFunction) Set(userdata *byte) {
	panic("not implemented")
	//iSetLogOutputFunction(callback, userdata)
}

// Condition

// SDL_DestroyCondition - Destroy a condition variable.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyCondition)
func (cond *Condition) Destroy() {
	panic("not implemented")
	iDestroyCondition(cond)
}

// SDL_SignalCondition - Restart one of the threads that are waiting on the condition variable.
// (https://wiki.libsdl.org/SDL3/SDL_SignalCondition)
func (cond *Condition) Signal() {
	panic("not implemented")
	iSignalCondition(cond)
}

// SDL_BroadcastCondition - Restart all threads that are waiting on the condition variable.
// (https://wiki.libsdl.org/SDL3/SDL_BroadcastCondition)
func (cond *Condition) Broadcast() {
	panic("not implemented")
	iBroadcastCondition(cond)
}

// SDL_WaitCondition - Wait until a condition variable is signaled.
// (https://wiki.libsdl.org/SDL3/SDL_WaitCondition)
func (cond *Condition) Wait(mutex *Mutex) {
	panic("not implemented")
	iWaitCondition(cond, mutex)
}

// SDL_WaitConditionTimeout - Wait until a condition variable is signaled or a certain time has passed.
// (https://wiki.libsdl.org/SDL3/SDL_WaitConditionTimeout)
func (cond *Condition) WaitTimeout(mutex *Mutex, timeoutMS int32) bool {
	panic("not implemented")
	return iWaitConditionTimeout(cond, mutex, timeoutMS)
}

// Window

// SDL_GetWindowPixelDensity - Get the pixel density of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity)
func (window *Window) PixelDensity() (float32, error) {
	density := iGetWindowPixelDensity(window)
	if density == 0 {
		return 0, internal.LastErr()
	}

	return density, nil
}

// SDL_GetWindowDisplayScale - Get the content display scale relative to a window's pixel size.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale)
func (window *Window) DisplayScale() (float32, error) {
	scale := iGetWindowDisplayScale(window)
	if scale == 0 {
		return 0, internal.LastErr()
	}

	return scale, nil
}

// SDL_SetWindowFullscreenMode - Set the display mode to use when a window is visible and fullscreen.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreenMode)
func (window *Window) SetFullscreenMode(mode *DisplayMode) error {
	if !iSetWindowFullscreenMode(window, mode) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowFullscreenMode - Query the display mode to use when a window is visible at fullscreen.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode)
func (window *Window) FullscreenMode() *DisplayMode {
	return iGetWindowFullscreenMode(window)
}

// SDL_GetWindowICCProfile - Get the raw ICC profile data for the screen the window is currently on.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile)
func (window *Window) ICCProfile() ([]byte, error) {
	var size uintptr

	ptr := iGetWindowICCProfile(window, &size)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[byte](ptr, int(size)), nil
}

// SDL_GetWindowPixelFormat - Get the pixel format associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat)
func (window *Window) PixelFormat() (PixelFormat, error) {
	format := iGetWindowPixelFormat(window)
	if format == PIXELFORMAT_UNKNOWN {
		return PIXELFORMAT_UNKNOWN, internal.LastErr()
	}

	return format, nil
}

// SDL_CreatePopupWindow - Create a child popup window of the specified parent window.
// (https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow)
func (parent *Window) CreatePopup(offsetX, offsetY, w, h int32, flags WindowFlags) (*Window, error) {
	window := iCreatePopupWindow(parent, offsetX, offsetY, w, h, flags)
	if window == nil {
		return nil, internal.LastErr()
	}

	return window, nil
}

// SDL_GetWindowID - Get the numeric ID of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowID)
func (window *Window) ID() (WindowID, error) {
	id := iGetWindowID(window)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetWindowParent - Get parent of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowParent)
func (window *Window) Parent() *Window {
	return iGetWindowParent(window)
}

// SDL_GetWindowProperties - Get the properties associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties)
func (window *Window) Properties() (PropertiesID, error) {
	props := iGetWindowProperties(window)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetWindowFlags - Get the window flags.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags)
func (window *Window) Flags() WindowFlags {
	return iGetWindowFlags(window)
}

// SDL_SetWindowTitle - Set the title of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle)
func (window *Window) SetTitle(title string) error {
	if !iSetWindowTitle(window, title) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowTitle - Get the title of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle)
func (window *Window) Title() string {
	return iGetWindowTitle(window)
}

// SDL_SetWindowIcon - Set the icon for a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon)
func (window *Window) SetIcon(icon *Surface) error {
	if !iSetWindowIcon(window, icon) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowPosition - Request that the window's position be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition)
func (window *Window) SetPosition(x, y int32) error {
	if !iSetWindowPosition(window, x, y) {
		return internal.LastErr()
	}

	return internal.LastErr()
}

// SDL_GetWindowPosition - Get the position of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition)
func (window *Window) Position() (int32, int32, error) {
	var x, y int32

	if !iGetWindowPosition(window, &x, &y) {
		return 0, 0, internal.LastErr()
	}

	return x, y, nil
}

// SDL_SetWindowSize - Request that the size of a window's client area be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowSize)
func (window *Window) SetSize(w, h int32) error {
	if !iSetWindowSize(window, w, h) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowSize - Get the size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSize)
func (window *Window) Size() (int32, int32, error) {
	var w, h int32

	if !iGetWindowSize(window, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_GetWindowSafeArea - Get the safe area for this window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea)
func (window *Window) SafeArea() (Rect, error) {
	var r Rect

	if !iGetWindowSafeArea(window, &r) {
		return r, internal.LastErr()
	}

	return r, nil
}

// SDL_SetWindowAspectRatio - Request that the aspect ratio of a window's client area be set.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio)
func (window *Window) SetAspectRatio(minAspect, maxAspect float32) error {
	if !iSetWindowAspectRatio(window, minAspect, maxAspect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowAspectRatio - Get the size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio)
func (window *Window) AspectRatio() (float32, float32, error) {
	var minAspect, maxAspect float32

	if !iGetWindowAspectRatio(window, &minAspect, &maxAspect) {
		return 0, 0, internal.LastErr()
	}

	return minAspect, maxAspect, nil
}

// SDL_GetWindowBordersSize - Get the size of a window's borders (decorations) around the client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize)
func (window *Window) BordersSize() (BorderSize, error) {
	var size BorderSize

	if !iGetWindowBordersSize(window, &size.Top, &size.Left, &size.Bottom, &size.Right) {
		return size, internal.LastErr()
	}

	return size, nil
}

// SDL_GetWindowSizeInPixels - Get the size of a window's client area, in pixels.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels)
func (window *Window) SizeInPixels() (int32, int32, error) {
	var w, h int32

	if !iGetWindowSizeInPixels(window, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_SetWindowMinimumSize - Set the minimum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize)
func (window *Window) SetMinimumSize(minW, minH int32) error {
	if !iSetWindowMinimumSize(window, minW, minH) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowMinimumSize - Get the minimum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize)
func (window *Window) MinimumSize() (int32, int32, error) {
	var w, h int32

	if !iGetWindowMinimumSize(window, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_SetWindowMaximumSize - Set the maximum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize)
func (window *Window) SetMaximumSize(maxW, maxH int32) error {
	if !iSetWindowMaximumSize(window, maxW, maxH) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowMaximumSize - Get the maximum size of a window's client area.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize)
func (window *Window) MaximumSize() (int32, int32, error) {
	var w, h int32

	if !iGetWindowMaximumSize(window, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// SDL_SetWindowBordered - Set the border state of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered)
func (window *Window) SetBordered(bordered bool) error {
	if !iSetWindowBordered(window, bordered) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowResizable - Set the user-resizable state of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable)
func (window *Window) SetResizable(resizable bool) error {
	if !iSetWindowResizable(window, resizable) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowAlwaysOnTop - Set the window to always be above the others.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop)
func (window *Window) SetAlwaysOnTop(onTop bool) error {
	if !iSetWindowAlwaysOnTop(window, onTop) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ShowWindow - Show a window.
// (https://wiki.libsdl.org/SDL3/SDL_ShowWindow)
func (window *Window) Show() error {
	if !iShowWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_HideWindow - Hide a window.
// (https://wiki.libsdl.org/SDL3/SDL_HideWindow)
func (window *Window) Hide() error {
	if !iHideWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RaiseWindow - Request that a window be raised above other windows and gain the input focus.
// (https://wiki.libsdl.org/SDL3/SDL_RaiseWindow)
func (window *Window) Raise() error {
	if !iRaiseWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_MaximizeWindow - Request that the window be made as large as possible.
// (https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow)
func (window *Window) Maximize() error {
	if !iMaximizeWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_MinimizeWindow - Request that the window be minimized to an iconic representation.
// (https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow)
func (window *Window) Minimize() error {
	if !iMinimizeWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_RestoreWindow - Request that the size and position of a minimized or maximized window be restored.
// (https://wiki.libsdl.org/SDL3/SDL_RestoreWindow)
func (window *Window) Restore() error {
	if !iRestoreWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowFullscreen - Request that the window's fullscreen state be changed.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen)
func (window *Window) SetFullscreen(fullscreen bool) error {
	if !iSetWindowFullscreen(window, fullscreen) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SyncWindow - Block until any pending window state is finalized.
// (https://wiki.libsdl.org/SDL3/SDL_SyncWindow)
func (window *Window) Sync() error {
	if !iSyncWindow(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WindowHasSurface - Return whether the window has a surface associated with it.
// (https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface)
func (window *Window) HasSurface() bool {
	return iWindowHasSurface(window)
}

// SDL_GetWindowSurface - Get the SDL surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface)
func (window *Window) Surface() (*Surface, error) {
	surface := iGetWindowSurface(window)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// SDL_SetWindowSurfaceVSync - Toggle VSync for the window surface.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync)
func (window *Window) SetSurfaceVSync(vsync int32) error {
	if !iSetWindowSurfaceVSync(window, vsync) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowSurfaceVSync - Get VSync for the window surface.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync)
func (window *Window) SurfaceVSync() (int32, error) {
	var vsync int32

	if !iGetWindowSurfaceVSync(window, &vsync) {
		return 0, internal.LastErr()
	}

	return vsync, nil
}

// SDL_UpdateWindowSurface - Copy the window surface to the screen.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface)
func (window *Window) UpdateSurface() error {
	if !iUpdateWindowSurface(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_UpdateWindowSurfaceRects - Copy areas of the window surface to the screen.
// (https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects)
func (window *Window) UpdateSurfaceRects(rects []Rect) error {
	if !iUpdateWindowSurfaceRects(window, unsafe.SliceData(rects), int32(len(rects))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DestroyWindowSurface - Destroy the surface associated with the window.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface)
func (window *Window) DestroySurface() error {
	if !iDestroyWindowSurface(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowKeyboardGrab - Set a window's keyboard grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab)
func (window *Window) SetKeyboardGrab(grabbed bool) error {
	if !iSetWindowKeyboardGrab(window, grabbed) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowMouseGrab - Set a window's mouse grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab)
func (window *Window) SetMouseGrab(grabbed bool) error {
	if !iSetWindowMouseGrab(window, grabbed) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowKeyboardGrab - Get a window's keyboard grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab)
func (window *Window) KeyboardGrab() bool {
	return iGetWindowKeyboardGrab(window)
}

// SDL_GetWindowMouseGrab - Get a window's mouse grab mode.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab)
func (window *Window) MouseGrab() bool {
	return iGetWindowMouseGrab(window)
}

// SDL_SetWindowMouseRect - Confines the cursor to the specified area of a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect)
func (window *Window) SetMouseRect(rect *Rect) error {
	if !iSetWindowMouseRect(window, rect) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowMouseRect - Get the mouse confinement rectangle of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect)
func (window *Window) MouseRect() *Rect {
	return iGetWindowMouseRect(window)
}

// SDL_SetWindowOpacity - Set the opacity for a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity)
func (window *Window) SetOpacity(opacity float32) error {
	if !iSetWindowOpacity(window, opacity) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowOpacity - Get the opacity of a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity)
func (window *Window) Opacity() float32 {
	return iGetWindowOpacity(window)
}

// SDL_SetWindowParent - Set the window as a child of a parent window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowParent)
func (window *Window) SetParent(parent *Window) error {
	if !iSetWindowParent(window, parent) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowModal - Toggle the state of the window as modal.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowModal)
func (window *Window) SetModal(modal bool) error {
	if !iSetWindowModal(window, modal) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowFocusable - Set whether the window may have input focus.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable)
func (window *Window) SetFocusable(focusable bool) error {
	if !iSetWindowFocusable(window, focusable) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ShowWindowSystemMenu - Display the system-level window menu.
// (https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu)
func (window *Window) ShowSystemMenu(x, y int32) error {
	if !iShowWindowSystemMenu(window, x, y) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetWindowHitTest - Provide a callback that decides if a window region has special properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest)
func (window *Window) SetHitTest(callback HitTest, callback_data *byte) bool {
	panic("not implemented")
	//return iSetWindowHitTest(window, callback, callback_data)
}

// SDL_SetWindowShape - Set the shape of a transparent window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowShape)
func (window *Window) SetShape(shape *Surface) error {
	if !iSetWindowShape(window, shape) {
		return internal.LastErr()
	}

	return nil
}

// SDL_FlashWindow - Request a window to demand attention from the user.
// (https://wiki.libsdl.org/SDL3/SDL_FlashWindow)
func (window *Window) Flash(operation FlashOperation) error {
	if !iFlashWindow(window, operation) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DestroyWindow - Destroy a window.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyWindow)
func (window *Window) Destroy() {
	iDestroyWindow(window)
}

// SDL_StartTextInput - Start accepting Unicode text input events in a window.
// (https://wiki.libsdl.org/SDL3/SDL_StartTextInput)
func (window *Window) StartTextInput() error {
	if !iStartTextInput(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_StartTextInputWithProperties - Start accepting Unicode text input events in a window, with properties describing the input.
// (https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties)
func (window *Window) StartTextInputWithProperties(props PropertiesID) error {
	if !iStartTextInputWithProperties(window, props) {
		return internal.LastErr()
	}

	return nil
}

// SDL_TextInputActive - Check whether or not Unicode text input events are enabled for a window.
// (https://wiki.libsdl.org/SDL3/SDL_TextInputActive)
func (window *Window) TextInputActive() bool {
	return iTextInputActive(window)
}

// SDL_StopTextInput - Stop receiving any text input events in a window.
// (https://wiki.libsdl.org/SDL3/SDL_StopTextInput)
func (window *Window) StopTextInput() error {
	if !iStopTextInput(window) {
		return internal.LastErr()
	}

	return nil
}

// SDL_ClearComposition - Dismiss the composition window/IME without disabling the subsystem.
// (https://wiki.libsdl.org/SDL3/SDL_ClearComposition)
func (window *Window) ClearComposition() bool {
	panic("not implemented")
	return iClearComposition(window)
}

// SDL_SetTextInputArea - Set the area used to type Unicode text input.
// (https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea)
func (window *Window) SetTextInputArea(rect *Rect, cursor int32) error {
	if !iSetTextInputArea(window, rect, cursor) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetTextInputArea - Get the area used to type Unicode text input.
// (https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea)
func (window *Window) TextInputArea() (Rect, int32, error) {
	var r Rect
	var cursor int32

	if !iGetTextInputArea(window, &r, &cursor) {
		return r, 0, internal.LastErr()
	}

	return r, cursor, nil
}

// SDL_ScreenKeyboardShown - Check whether the screen keyboard is shown for given window.
// (https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown)
func (window *Window) ScreenKeyboardShown() bool {
	return iScreenKeyboardShown(window)
}

// SDL_WarpMouseInWindow - Move the mouse cursor to the given position within the window.
// (https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow)
func (window *Window) WarpMouseIn(x float32, y float32) {
	iWarpMouseInWindow(window, x, y)
}

// SDL_SetWindowRelativeMouseMode - Set relative mouse mode for a window.
// (https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode)
func (window *Window) SetRelativeMouseMode(enabled bool) error {
	if !iSetWindowRelativeMouseMode(window, enabled) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetWindowRelativeMouseMode - Query whether relative mouse mode is enabled for a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode)
func (window *Window) RelativeMouseMode() bool {
	return iGetWindowRelativeMouseMode(window)
}

// SDL_Metal_CreateView - Create a CAMetalLayer-backed NSView/UIView and attach it to the specified window.
// (https://wiki.libsdl.org/SDL3/SDL_Metal_CreateView)
func (window *Window) Metal_CreateView() MetalView {
	panic("not implemented")
	return iMetal_CreateView(window)
}

// SDL_CreateRenderer - Create a 2D rendering context for a window.
// (https://wiki.libsdl.org/SDL3/SDL_CreateRenderer)
func (window *Window) CreateRenderer(name string) (*Renderer, error) {
	renderer := iCreateRenderer(window, name)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// SDL_GetRenderer - Get the renderer associated with a window.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderer)
func (window *Window) Renderer() *Renderer {
	return iGetRenderer(window)
}

// Scancode

// SDL_GetKeyFromScancode - Get the key code corresponding to the given scancode according to the current keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode)
func (scancode Scancode) KeyFrom(modstate Keymod, keyEvent bool) Keycode {
	return iGetKeyFromScancode(scancode, modstate, keyEvent)
}

// SDL_SetScancodeName - Set a human-readable name for a scancode.
// (https://wiki.libsdl.org/SDL3/SDL_SetScancodeName)
func (scancode Scancode) SetName(name string) error {
	if !iSetScancodeName(scancode, name) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetScancodeName - Get a human-readable name for a scancode.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeName)
func (scancode Scancode) Name() string {
	return iGetScancodeName(scancode)
}

func (scancode Scancode) ToKeycode() Keycode {
	return Keycode(scancode | K_SCANCODE_MASK)
}

// SharedObject

// SDL_LoadFunction - Look up the address of the named function in a shared object.
// (https://wiki.libsdl.org/SDL3/SDL_LoadFunction)
func (handle *SharedObject) LoadFunction(name string) FunctionPointer {
	panic("not implemented")
	return iLoadFunction(handle, name)
}

// SDL_UnloadObject - Unload a shared object from memory.
// (https://wiki.libsdl.org/SDL3/SDL_UnloadObject)
func (handle *SharedObject) UnloadObject() {
	panic("not implemented")
	iUnloadObject(handle)
}

// TimerID

// SDL_RemoveTimer - Remove a timer created with SDL_AddTimer().
// (https://wiki.libsdl.org/SDL3/SDL_RemoveTimer)
func (id TimerID) RemoveTimer() bool {
	panic("not implemented")
	return iRemoveTimer(id)
}

// IOStream

// SDL_CloseIO - Close and free an allocated SDL_IOStream structure.
// (https://wiki.libsdl.org/SDL3/SDL_CloseIO)
func (context *IOStream) Close() error {
	if !iCloseIO(context) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetIOProperties - Get the properties associated with an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOProperties)
func (context *IOStream) Properties() (PropertiesID, error) {
	props := iGetIOProperties(context)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetIOStatus - Query the stream status of an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOStatus)
func (context *IOStream) Status() IOStatus {
	return iGetIOStatus(context)
}

// SDL_GetIOSize - Use this function to get the size of the data stream in an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_GetIOSize)
func (context *IOStream) Size() (int64, error) {
	size := iGetIOSize(context)
	if size < 0 {
		return -1, internal.LastErr()
	}

	return size, nil
}

// SDL_SeekIO - Seek within an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_SeekIO)
func (context *IOStream) Seek(offset int64, whence IOWhence) (int64, error) {
	seek := iSeekIO(context, offset, whence)
	if seek == -1 {
		return -1, internal.LastErr()
	}

	return seek, nil
}

// SDL_TellIO - Determine the current read/write offset in an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_TellIO)
func (context *IOStream) Tell() int64 {
	return iTellIO(context)
}

// SDL_ReadIO - Read from a data source.
// (https://wiki.libsdl.org/SDL3/SDL_ReadIO)
func (context *IOStream) Read(available []byte) (uint64, error) {
	count := iReadIO(context, uintptr(unsafe.Pointer(unsafe.SliceData(available))), uintptr(len(available)))
	if count == 0 {
		return 0, internal.LastErr()
	}

	return uint64(count), nil
}

// SDL_WriteIO - Write to an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_WriteIO)
func (context *IOStream) Write(data []byte) (uint64, error) {
	count := iWriteIO(context, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uintptr(len(data)))
	if count < uintptr(len(data)) {
		return uint64(count), internal.LastErr()
	}

	return uint64(count), nil
}

// SDL_IOprintf - Print to an SDL_IOStream data stream.
// (https://wiki.libsdl.org/SDL3/SDL_IOprintf)
func (context *IOStream) Printf(format string, values ...any) (uint64, error) {
	count := iIOprintf(context, fmt.Sprintf(format, values...))
	if count == 0 {
		return 0, internal.LastErr()
	}

	return uint64(count), nil
}

// SDL_FlushIO - Flush any buffered data in the stream.
// (https://wiki.libsdl.org/SDL3/SDL_FlushIO)
func (context *IOStream) Flush() error {
	if !iFlushIO(context) {
		return internal.LastErr()
	}

	return nil
}

// SDL_LoadFile_IO - Load all the data from an SDL data stream.
// (https://wiki.libsdl.org/SDL3/SDL_LoadFile_IO)
func (src *IOStream) LoadFile(closeio bool) ([]byte, error) {
	var size uintptr
	ptr := iLoadFile_IO(src, &size, closeio)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[byte](ptr, int(size)), nil
}

// SDL_SaveFile_IO - Save all the data into an SDL data stream.
// (https://wiki.libsdl.org/SDL3/SDL_SaveFile_IO)
func (src *IOStream) SaveFile(data []byte, closeio bool) error {
	if !iSaveFile_IO(src, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uintptr(len(data)), closeio) {
		return internal.LastErr()
	}
	runtime.KeepAlive(data)

	return nil
}

// SDL_ReadU8 - Use this function to read a byte from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU8)
func (src *IOStream) ReadU8() (uint8, error) {
	var value uint8

	if !iReadU8(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS8 - Use this function to read a signed byte from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS8)
func (src *IOStream) ReadS8() (int8, error) {
	var value int8

	if !iReadS8(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadU16LE - Use this function to read 16 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU16LE)
func (src *IOStream) ReadU16LE() (uint16, error) {
	var value uint16

	if !iReadU16LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS16LE - Use this function to read 16 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS16LE)
func (src *IOStream) ReadS16LE() (int16, error) {
	var value int16

	if !iReadS16LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadU16BE - Use this function to read 16 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU16BE)
func (src *IOStream) ReadU16BE() (uint16, error) {
	var value uint16

	if !iReadU16BE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS16BE - Use this function to read 16 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS16BE)
func (src *IOStream) ReadS16BE() (int16, error) {
	var value int16

	if !iReadS16BE(src, &value) {
		return 0, nil
	}

	return value, nil
}

// SDL_ReadU32LE - Use this function to read 32 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU32LE)
func (src *IOStream) ReadU32LE() (uint32, error) {
	var value uint32

	if !iReadU32LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS32LE - Use this function to read 32 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS32LE)
func (src *IOStream) ReadS32LE() (int32, error) {
	var value int32

	if !iReadS32LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadU32BE - Use this function to read 32 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU32BE)
func (src *IOStream) ReadU32BE() (uint32, error) {
	var value uint32

	if !iReadU32BE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS32BE - Use this function to read 32 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS32BE)
func (src *IOStream) ReadS32BE() (int32, error) {
	var value int32

	if !iReadS32BE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadU64LE - Use this function to read 64 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU64LE)
func (src *IOStream) ReadU64LE() (uint64, error) {
	var value uint64

	if !iReadU64LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS64LE - Use this function to read 64 bits of little-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS64LE)
func (src *IOStream) ReadS64LE() (int64, error) {
	var value int64

	if !iReadS64LE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadU64BE - Use this function to read 64 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadU64BE)
func (src *IOStream) ReadU64BE() (uint64, error) {
	var value uint64

	if !iReadU64BE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_ReadS64BE - Use this function to read 64 bits of big-endian data from an SDL_IOStream and return in native format.
// (https://wiki.libsdl.org/SDL3/SDL_ReadS64BE)
func (src *IOStream) ReadS64BE() (int64, error) {
	var value int64

	if !iReadS64BE(src, &value) {
		return 0, internal.LastErr()
	}

	return value, nil
}

// SDL_WriteU8 - Use this function to write a byte to an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU8)
func (dst *IOStream) WriteU8(value uint8) error {
	if !iWriteU8(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS8 - Use this function to write a signed byte to an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS8)
func (dst *IOStream) WriteS8(value int8) error {
	if !iWriteS8(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU16LE - Use this function to write 16 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU16LE)
func (dst *IOStream) WriteU16LE(value uint16) error {
	if !iWriteU16LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS16LE - Use this function to write 16 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS16LE)
func (dst *IOStream) WriteS16LE(value int16) error {
	if !iWriteS16LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU16BE - Use this function to write 16 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU16BE)
func (dst *IOStream) WriteU16BE(value uint16) error {
	if !iWriteU16BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS16BE - Use this function to write 16 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS16BE)
func (dst *IOStream) WriteS16BE(value int16) error {
	if !iWriteS16BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU32LE - Use this function to write 32 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU32LE)
func (dst *IOStream) WriteU32LE(value uint32) error {
	if !iWriteU32LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS32LE - Use this function to write 32 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS32LE)
func (dst *IOStream) WriteS32LE(value int32) error {
	if !iWriteS32LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU32BE - Use this function to write 32 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU32BE)
func (dst *IOStream) WriteU32BE(value uint32) error {
	if !iWriteU32BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS32BE - Use this function to write 32 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS32BE)
func (dst *IOStream) WriteS32BE(value int32) error {
	if !iWriteS32BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU64LE - Use this function to write 64 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU64LE)
func (dst *IOStream) WriteU64LE(value uint64) error {
	if !iWriteU64LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS64LE - Use this function to write 64 bits in native format to an SDL_IOStream as little-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS64LE)
func (dst *IOStream) WriteS64LE(value int64) error {
	if !iWriteS64LE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteU64BE - Use this function to write 64 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteU64BE)
func (dst *IOStream) WriteU64BE(value uint64) error {
	if !iWriteU64BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WriteS64BE - Use this function to write 64 bits in native format to an SDL_IOStream as big-endian data.
// (https://wiki.libsdl.org/SDL3/SDL_WriteS64BE)
func (dst *IOStream) WriteS64BE(value int64) error {
	if !iWriteS64BE(dst, value) {
		return internal.LastErr()
	}

	return nil
}

// Palette

// SDL_SetPaletteColors - Set a range of colors in a palette.
// (https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors)
func (palette *Palette) SetColors(colors []Color, firstcolor int32) error {
	if !iSetPaletteColors(palette, unsafe.SliceData(colors), firstcolor, int32(len(colors))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_DestroyPalette - Free a palette created with SDL_CreatePalette().
// (https://wiki.libsdl.org/SDL3/SDL_DestroyPalette)
func (palette *Palette) Destroy() {
	iDestroyPalette(palette)
}

// Joystick

// SDL_SetJoystickVirtualAxis - Set the state of an axis on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis)
func (joystick *Joystick) SetVirtualAxis(axis int32, value int16) error {
	if !iSetJoystickVirtualAxis(joystick, axis, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetJoystickVirtualBall - Generate ball motion on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall)
func (joystick *Joystick) SetVirtualBall(ball int32, xrel, yrel int16) error {
	if !iSetJoystickVirtualBall(joystick, ball, xrel, yrel) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetJoystickVirtualButton - Set the state of a button on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton)
func (joystick *Joystick) SetVirtualButton(button int32, down bool) error {
	if !iSetJoystickVirtualButton(joystick, button, down) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetJoystickVirtualHat - Set the state of a hat on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat)
func (joystick *Joystick) SetVirtualHat(hat int32, value uint8) error {
	if !iSetJoystickVirtualHat(joystick, hat, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetJoystickVirtualTouchpad - Set touchpad finger state on an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad)
func (joystick *Joystick) SetVirtualTouchpad(touchpad int32, finger int32, down bool, x, y, pressure float32) error {
	if !iSetJoystickVirtualTouchpad(joystick, touchpad, finger, down, x, y, pressure) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SendJoystickVirtualSensorData - Send a sensor update for an opened virtual joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData)
func (joystick *Joystick) SendVirtualSensorData(typ SensorType, sensorTimestamp uint64, data []float32) error {
	if !iSendJoystickVirtualSensorData(joystick, typ, sensorTimestamp, unsafe.SliceData(data), int32(len(data))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetJoystickProperties - Get the properties associated with a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties)
func (joystick *Joystick) Properties() (PropertiesID, error) {
	props := iGetJoystickProperties(joystick)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_GetJoystickName - Get the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickName)
func (joystick *Joystick) Name() (string, error) {
	name := iGetJoystickName(joystick)
	if name == "" {
		return "", internal.LastErr()
	}

	return name, nil
}

// SDL_GetJoystickPath - Get the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath)
func (joystick *Joystick) Path() (string, error) {
	path := iGetJoystickPath(joystick)
	if path == "" {
		return "", internal.LastErr()
	}

	return path, nil
}

// SDL_GetJoystickPlayerIndex - Get the player index of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex)
func (joystick *Joystick) PlayerIndex() int32 {
	return iGetJoystickPlayerIndex(joystick)
}

// SDL_SetJoystickPlayerIndex - Set the player index of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex)
func (joystick *Joystick) SetPlayerIndex(playerIndex int32) error {
	if !iSetJoystickPlayerIndex(joystick, playerIndex) {
		return internal.LastErr()
	}

	return nil
}

// SDL_GetJoystickGUID - Get the implementation-dependent GUID for the joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID)
func (joystick *Joystick) GUID() GUID {
	panic("not implemented")
	return iGetJoystickGUID(joystick)
}

// SDL_GetJoystickVendor - Get the USB vendor ID of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor)
func (joystick *Joystick) Vendor() uint16 {
	return iGetJoystickVendor(joystick)
}

// SDL_GetJoystickProduct - Get the USB product ID of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct)
func (joystick *Joystick) Product() uint16 {
	return iGetJoystickProduct(joystick)
}

// SDL_GetJoystickProductVersion - Get the product version of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion)
func (joystick *Joystick) ProductVersion() uint16 {
	return iGetJoystickProductVersion(joystick)
}

// SDL_GetJoystickFirmwareVersion - Get the firmware version of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion)
func (joystick *Joystick) FirmwareVersion() uint16 {
	return iGetJoystickFirmwareVersion(joystick)
}

// SDL_GetJoystickSerial - Get the serial number of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial)
func (joystick *Joystick) Serial() string {
	return iGetJoystickSerial(joystick)
}

// SDL_GetJoystickType - Get the type of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickType)
func (joystick *Joystick) Type() JoystickType {
	return iGetJoystickType(joystick)
}

// SDL_JoystickConnected - Get the status of a specified joystick.
// (https://wiki.libsdl.org/SDL3/SDL_JoystickConnected)
func (joystick *Joystick) Connected() bool {
	return iJoystickConnected(joystick)
}

// SDL_GetJoystickID - Get the instance ID of an opened joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickID)
func (joystick *Joystick) ID() (JoystickID, error) {
	id := iGetJoystickID(joystick)
	if id == 0 {
		return 0, internal.LastErr()
	}

	return id, nil
}

// SDL_GetNumJoystickAxes - Get the number of general axis controls on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes)
func (joystick *Joystick) NumAxes() (int32, error) {
	num := iGetNumJoystickAxes(joystick)
	if num == -1 {
		return -1, internal.LastErr()
	}

	return num, nil
}

// SDL_GetNumJoystickBalls - Get the number of trackballs on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls)
func (joystick *Joystick) NumBalls() (int32, error) {
	num := iGetNumJoystickBalls(joystick)
	if num == -1 {
		return -1, internal.LastErr()
	}

	return num, nil
}

// SDL_GetNumJoystickHats - Get the number of POV hats on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats)
func (joystick *Joystick) NumHats() (int32, error) {
	num := iGetNumJoystickHats(joystick)
	if num == -1 {
		return -1, internal.LastErr()
	}

	return num, nil
}

// SDL_GetNumJoystickButtons - Get the number of buttons on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons)
func (joystick *Joystick) NumButtons() (int32, error) {
	num := iGetNumJoystickButtons(joystick)
	if num == -1 {
		return -1, internal.LastErr()
	}

	return num, nil
}

// SDL_GetJoystickAxis - Get the current state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis)
func (joystick *Joystick) Axis(axis int32) (int16, error) {
	res := iGetJoystickAxis(joystick, axis)
	if res == 0 {
		return 0, internal.LastErr()
	}

	return res, nil
}

// SDL_GetJoystickAxisInitialState - Get the initial state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState)
func (joystick *Joystick) AxisInitialState(axis int32) (int16, bool) {
	var state int16

	has := iGetJoystickAxisInitialState(joystick, axis, &state)
	return state, has
}

// SDL_GetJoystickBall - Get the ball axis change since the last poll.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall)
func (joystick *Joystick) Ball(ball int32) (int32, int32, error) {
	var dx, dy int32

	if !iGetJoystickBall(joystick, ball, &dx, &dy) {
		return 0, 0, internal.LastErr()
	}

	return dx, dy, nil
}

// SDL_GetJoystickHat - Get the current state of a POV hat on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat)
func (joystick *Joystick) Hat(hat int32) uint8 {
	return iGetJoystickHat(joystick, hat)
}

// SDL_GetJoystickButton - Get the current state of a button on a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton)
func (joystick *Joystick) Button(button int32) bool {
	return iGetJoystickButton(joystick, button)
}

// SDL_RumbleJoystick - Start a rumble effect.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick)
func (joystick *Joystick) Rumble(lowFreqencyRumble, highFrequencyRumble uint16, durationMS uint32) bool {
	return iRumbleJoystick(joystick, lowFreqencyRumble, highFrequencyRumble, durationMS)
}

// SDL_RumbleJoystickTriggers - Start a rumble effect in the joystick's triggers.
// (https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers)
func (joystick *Joystick) RumbleTriggers(leftRumble, rightRumble uint16, durationMS uint32) error {
	if !iRumbleJoystickTriggers(joystick, leftRumble, rightRumble, durationMS) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetJoystickLED - Update a joystick's LED color.
// (https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED)
func (joystick *Joystick) SetLED(red, green, blue uint8) error {
	if !iSetJoystickLED(joystick, red, green, blue) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SendJoystickEffect - Send a joystick specific effect packet.
// (https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect)
func (joystick *Joystick) SendEffect(data []byte) error {
	if !iSendJoystickEffect(joystick, uintptr(unsafe.Pointer(unsafe.SliceData(data))), int32(len(data))) {
		return internal.LastErr()
	}

	return nil
}

// SDL_CloseJoystick - Close a joystick previously opened with SDL_OpenJoystick().
// (https://wiki.libsdl.org/SDL3/SDL_CloseJoystick)
func (joystick *Joystick) Close() {
	iCloseJoystick(joystick)
}

// SDL_GetJoystickConnectionState - Get the connection state of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState)
func (joystick *Joystick) ConnectionState() (JoystickConnectionState, error) {
	state := iGetJoystickConnectionState(joystick)
	if state == JOYSTICK_CONNECTION_INVALID {
		return JOYSTICK_CONNECTION_INVALID, internal.LastErr()
	}

	return state, nil
}

// SDL_GetJoystickPowerInfo - Get the battery state of a joystick.
// (https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo)
func (joystick *Joystick) PowerInfo(percent *int32) (*PowerInfo, error) {
	var info PowerInfo

	info.State = iGetJoystickPowerInfo(joystick, &info.Percent)
	if info.State == POWERSTATE_ERROR {
		return nil, internal.LastErr()
	}

	return &info, nil
}

// SDL_IsJoystickHaptic - Query if a joystick has haptic features.
// (https://wiki.libsdl.org/SDL3/SDL_IsJoystickHaptic)
func (joystick *Joystick) IsHaptic() bool {
	return iIsJoystickHaptic(joystick)
}

// SDL_OpenHapticFromJoystick - Open a haptic device for use from a joystick device.
// (https://wiki.libsdl.org/SDL3/SDL_OpenHapticFromJoystick)
func (joystick *Joystick) OpenHapticFrom() (*Haptic, error) {
	haptic := iOpenHapticFromJoystick(joystick)
	if haptic == nil {
		return nil, internal.LastErr()
	}

	return haptic, nil
}

// GamepadType

// SDL_GetGamepadStringForType - Convert from an SDL_GamepadType enum to a string.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType)
func (typ GamepadType) GamepadStringForType() string {
	return iGetGamepadStringForType(typ)
}

// SDL_GetGamepadButtonLabelForType - Get the label of a button on a gamepad.
// (https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType)
func (typ GamepadType) GamepadButtonLabelForType(button GamepadButton) GamepadButtonLabel {
	return iGetGamepadButtonLabelForType(typ, button)
}

// StorageInterface

// SDL_OpenStorage - Opens up a container using a client-provided storage interface.
// (https://wiki.libsdl.org/SDL3/SDL_OpenStorage)
func (iface *StorageInterface) OpenStorage(userdata *byte) *Storage {
	panic("not implemented")
	//return iOpenStorage(iface, userdata)
}

// Mutex

// SDL_LockMutex - Lock the mutex.
// (https://wiki.libsdl.org/SDL3/SDL_LockMutex)
func (mutex *Mutex) Lock() {
	iLockMutex(mutex)
}

// SDL_TryLockMutex - Try to lock a mutex without blocking.
// (https://wiki.libsdl.org/SDL3/SDL_TryLockMutex)
func (mutex *Mutex) TryLock() bool {
	return iTryLockMutex(mutex)
}

// SDL_UnlockMutex - Unlock the mutex.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockMutex)
func (mutex *Mutex) Unlock() {
	iUnlockMutex(mutex)
}

// SDL_DestroyMutex - Destroy a mutex created with SDL_CreateMutex().
// (https://wiki.libsdl.org/SDL3/SDL_DestroyMutex)
func (mutex *Mutex) Destroy() {
	iDestroyMutex(mutex)
}

// Time

// SDL_TimeToDateTime - Converts an SDL_Time in nanoseconds since the epoch to a calendar time in the SDL_DateTime format.
// (https://wiki.libsdl.org/SDL3/SDL_TimeToDateTime)
func (ticks Time) ToDateTime(dt *DateTime, localTime bool) bool {
	panic("not implemented")
	return iTimeToDateTime(ticks, dt, localTime)
}

// SDL_TimeToWindows - Converts an SDL time into a Windows FILETIME (100-nanosecond intervals since January 1, 1601).
// (https://wiki.libsdl.org/SDL3/SDL_TimeToWindows)
func (ticks Time) ToWindows(dwLowDateTime *uint32, dwHighDateTime *uint32) {
	panic("not implemented")
	iTimeToWindows(ticks, dwLowDateTime, dwHighDateTime)
}

// AtomicU32

// SDL_CompareAndSwapAtomicU32 - Set an atomic variable to a new value if it is currently an old value.
// (https://wiki.libsdl.org/SDL3/SDL_CompareAndSwapAtomicU32)
func (a *AtomicU32) CompareAndSwap(oldval uint32, newval uint32) bool {
	panic("not implemented")
	return iCompareAndSwapAtomicU32(a, oldval, newval)
}

// SDL_SetAtomicU32 - Set an atomic variable to a value.
// (https://wiki.libsdl.org/SDL3/SDL_SetAtomicU32)
func (a *AtomicU32) Set(v uint32) uint32 {
	panic("not implemented")
	return iSetAtomicU32(a, v)
}

// SDL_GetAtomicU32 - Get the value of an atomic variable.
// (https://wiki.libsdl.org/SDL3/SDL_GetAtomicU32)
func (a *AtomicU32) Get() uint32 {
	panic("not implemented")
	return iGetAtomicU32(a)
}

// PropertiesID

// SDL_LockProperties - Lock a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_LockProperties)
func (props PropertiesID) Lock() error {
	if !iLockProperties(props) {
		return internal.LastErr()
	}

	return nil
}

// SDL_UnlockProperties - Unlock a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_UnlockProperties)
func (props PropertiesID) Unlock() {
	iUnlockProperties(props)
}

// SDL_SetPointerPropertyWithCleanup - Set a pointer property in a group of properties with a cleanup function that is called when the property is deleted.
// (https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup)
func (props PropertiesID) SetPointerPropertyWithCleanup(name string, value *byte, cleanup CleanupPropertyCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetPointerPropertyWithCleanup(props, name, value, cleanup, userdata)
}

// SDL_SetPointerProperty - Set a pointer property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty)
func (props PropertiesID) SetPointerProperty(name string, value *byte) bool {
	panic("not implemented")
	//return iSetPointerProperty(props, name, value)
}

// SDL_SetStringProperty - Set a string property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetStringProperty)
func (props PropertiesID) SetStringProperty(name, value string) error {
	if !iSetStringProperty(props, name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetNumberProperty - Set an integer property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty)
func (props PropertiesID) SetNumberProperty(name string, value int64) error {
	if !iSetNumberProperty(props, name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetFloatProperty - Set a floating point property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty)
func (props PropertiesID) SetFloatProperty(name string, value float32) error {
	if !iSetFloatProperty(props, name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SetBooleanProperty - Set a boolean property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty)
func (props PropertiesID) SetBooleanProperty(name string, value bool) error {
	if !iSetBooleanProperty(props, name, value) {
		return internal.LastErr()
	}

	return nil
}

// SDL_HasProperty - Return whether a property exists in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_HasProperty)
func (props PropertiesID) HasProperty(name string) bool {
	return iHasProperty(props, name)
}

// SDL_GetPropertyType - Get the type of a property in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetPropertyType)
func (props PropertiesID) PropertyType(name string) PropertyType {
	return iGetPropertyType(props, name)
}

// SDL_GetPointerProperty - Get a pointer property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty)
func (props PropertiesID) PointerProperty(name string, default_value *byte) *byte {
	panic("not implemented")
	//return iGetPointerProperty(props, name, default_value)
}

// SDL_GetStringProperty - Get a string property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetStringProperty)
func (props PropertiesID) StringProperty(name, defaultValue string) string {
	return iGetStringProperty(props, name, defaultValue)
}

// SDL_GetNumberProperty - Get a number property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty)
func (props PropertiesID) NumberProperty(name string, defaultValue int64) int64 {
	return iGetNumberProperty(props, name, defaultValue)
}

// SDL_GetFloatProperty - Get a floating point property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty)
func (props PropertiesID) FloatProperty(name string, defaultValue float32) float32 {
	return iGetFloatProperty(props, name, defaultValue)
}

// SDL_GetBooleanProperty - Get a boolean property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty)
func (props PropertiesID) BooleanProperty(name string, defaultValue bool) bool {
	return iGetBooleanProperty(props, name, defaultValue)
}

// SDL_ClearProperty - Clear a property from a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_ClearProperty)
func (props PropertiesID) ClearProperty(name string) error {
	if !iClearProperty(props, name) {
		return internal.LastErr()
	}

	return nil
}

// SDL_EnumerateProperties - Enumerate the properties contained in a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties)
func (props PropertiesID) EnumerateProperties(callback EnumeratePropertiesCallback, userdata uintptr) bool {
	return iEnumerateProperties(props, callback, userdata)
}

// SDL_DestroyProperties - Destroy a group of properties.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyProperties)
func (props PropertiesID) Destroy() {
	iDestroyProperties(props)
}

// SDL_GPUSupportsProperties - Checks for GPU runtime support.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSupportsProperties)
func (props PropertiesID) GPUSupport() bool {
	return iGPUSupportsProperties(props)
}

// Semaphore

// SDL_DestroySemaphore - Destroy a semaphore.
// (https://wiki.libsdl.org/SDL3/SDL_DestroySemaphore)
func (sem *Semaphore) Destroy() {
	iDestroySemaphore(sem)
}

// SDL_WaitSemaphore - Wait until a semaphore has a positive value and then decrements it.
// (https://wiki.libsdl.org/SDL3/SDL_WaitSemaphore)
func (sem *Semaphore) Wait() {
	iWaitSemaphore(sem)
}

// SDL_TryWaitSemaphore - See if a semaphore has a positive value and decrement it if it does.
// (https://wiki.libsdl.org/SDL3/SDL_TryWaitSemaphore)
func (sem *Semaphore) TryWait() bool {
	return iTryWaitSemaphore(sem)
}

// SDL_WaitSemaphoreTimeout - Wait until a semaphore has a positive value and then decrements it.
// (https://wiki.libsdl.org/SDL3/SDL_WaitSemaphoreTimeout)
func (sem *Semaphore) WaitTimeout(timeoutMS int32) bool {
	return iWaitSemaphoreTimeout(sem, timeoutMS)
}

// SDL_SignalSemaphore - Atomically increment a semaphore's value and wake waiting threads.
// (https://wiki.libsdl.org/SDL3/SDL_SignalSemaphore)
func (sem *Semaphore) Signal() {
	iSignalSemaphore(sem)
}

// SDL_GetSemaphoreValue - Get the current value of a semaphore.
// (https://wiki.libsdl.org/SDL3/SDL_GetSemaphoreValue)
func (sem *Semaphore) Value() uint32 {
	return iGetSemaphoreValue(sem)
}

// PixelFormat

// SDL_GetPixelFormatName - Get the human readable name of a pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName)
func (format PixelFormat) Name() string {
	return iGetPixelFormatName(format)
}

// SDL_GetMasksForPixelFormat - Convert one of the enumerated pixel formats to a bpp value and RGBA masks.
// (https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat)
func (format PixelFormat) Masks(bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
	// TODO: make it return (color.RGBA, error) or something
	panic("not implemented")
	return iGetMasksForPixelFormat(format, bpp, Rmask, Gmask, Bmask, Amask)
}

// SDL_GetPixelFormatDetails - Create an SDL_PixelFormatDetails structure corresponding to a pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails)
func (format PixelFormat) Details() (*PixelFormatDetails, error) {
	details := iGetPixelFormatDetails(format)
	if details == nil {
		return nil, internal.LastErr()
	}

	return details, nil
}

// DateTime

// SDL_DateTimeToTime - Converts a calendar time to an SDL_Time in nanoseconds since the epoch.
// (https://wiki.libsdl.org/SDL3/SDL_DateTimeToTime)
func (dt *DateTime) ToTime() (Time, error) {
	var ticks Time

	if !iDateTimeToTime(dt, &ticks) {
		return 0, internal.LastErr()
	}

	return ticks, nil
}

// Keycode

// SDL_GetScancodeFromKey - Get the scancode corresponding to the given key code according to the current keyboard layout.
// (https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey)
func (key Keycode) ScancodeFromKey(modstate *Keymod) Scancode {
	return iGetScancodeFromKey(key, modstate)
}

// SDL_GetKeyName - Get a human-readable name for a key.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyName)
func (key Keycode) KeyName() string {
	return iGetKeyName(key)
}

// GPUCommandBuffer

// SDL_InsertGPUDebugLabel - Inserts an arbitrary string label into the command buffer callstream.
// (https://wiki.libsdl.org/SDL3/SDL_InsertGPUDebugLabel)
func (cb *GPUCommandBuffer) InsertDebugLabel(text string) {
	iInsertGPUDebugLabel(cb, text)
}

// SDL_PushGPUDebugGroup - Begins a debug group with an arbitrary name.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUDebugGroup)
func (cb *GPUCommandBuffer) PushDebugGroup(name string) {
	iPushGPUDebugGroup(cb, name)
}

// SDL_PopGPUDebugGroup - Ends the most-recently pushed debug group.
// (https://wiki.libsdl.org/SDL3/SDL_PopGPUDebugGroup)
func (cb *GPUCommandBuffer) PopDebugGroup() {
	iPopGPUDebugGroup(cb)
}

// SDL_PushGPUVertexUniformData - Pushes data to a vertex uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUVertexUniformData)
func (cb *GPUCommandBuffer) PushVertexUniformData(slotIndex uint32, data []byte) {
	iPushGPUVertexUniformData(cb, slotIndex, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(len(data)))
}

// SDL_PushGPUFragmentUniformData - Pushes data to a fragment uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUFragmentUniformData)
func (cb *GPUCommandBuffer) PushFragmentUniformData(slotIndex uint32, data []byte) {
	iPushGPUFragmentUniformData(cb, slotIndex, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(len(data)))
}

// SDL_PushGPUComputeUniformData - Pushes data to a uniform slot on the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_PushGPUComputeUniformData)
func (cb *GPUCommandBuffer) PushComputeUniformData(slotIndex uint32, data []byte) {
	iPushGPUComputeUniformData(cb, slotIndex, uintptr(unsafe.Pointer(unsafe.SliceData(data))), uint32(len(data)))
}

// SDL_BeginGPURenderPass - Begins a render pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPURenderPass)
func (cb *GPUCommandBuffer) BeginRenderPass(colorTargetInfos []GPUColorTargetInfo, depthStencilTargetInfo *GPUDepthStencilTargetInfo) *GPURenderPass {
	return iBeginGPURenderPass(cb, unsafe.SliceData(colorTargetInfos), uint32(len(colorTargetInfos)), depthStencilTargetInfo)
}

// SDL_BeginGPUComputePass - Begins a compute pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPUComputePass)
func (cb *GPUCommandBuffer) BeginComputePass(storageTextureBindings []GPUStorageTextureReadWriteBinding, storageBufferBindings []GPUStorageBufferReadWriteBinding) *GPUComputePass {
	return iBeginGPUComputePass(cb, unsafe.SliceData(storageTextureBindings), uint32(len(storageTextureBindings)), unsafe.SliceData(storageBufferBindings), uint32(len(storageBufferBindings)))
}

// SDL_BeginGPUCopyPass - Begins a copy pass on a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_BeginGPUCopyPass)
func (cb *GPUCommandBuffer) BeginCopyPass() *GPUCopyPass {
	return iBeginGPUCopyPass(cb)
}

// SDL_GenerateMipmapsForGPUTexture - Generates mipmaps for the given texture.
// (https://wiki.libsdl.org/SDL3/SDL_GenerateMipmapsForGPUTexture)
func (cb *GPUCommandBuffer) GenerateMipmapsForGPUTexture(texture *GPUTexture) {
	iGenerateMipmapsForGPUTexture(cb, texture)
}

// SDL_BlitGPUTexture - Blits from a source texture region to a destination texture region.
// (https://wiki.libsdl.org/SDL3/SDL_BlitGPUTexture)
func (cb *GPUCommandBuffer) BlitGPUTexture(info *GPUBlitInfo) {
	iBlitGPUTexture(cb, info)
}

// SDL_AcquireGPUSwapchainTexture - Acquire a texture to use in presentation.
// (https://wiki.libsdl.org/SDL3/SDL_AcquireGPUSwapchainTexture)
func (cb *GPUCommandBuffer) AcquireGPUSwapchainTexture(window *Window) (*SwapchainTexture, error) {
	var texture SwapchainTexture

	if !iAcquireGPUSwapchainTexture(cb, window, &texture.Texture, &texture.Width, &texture.Height) {
		return nil, internal.LastErr()
	}

	return &texture, nil
}

// SDL_WaitAndAcquireGPUSwapchainTexture - Blocks the thread until a swapchain texture is available to be acquired, and then acquires it.
// (https://wiki.libsdl.org/SDL3/SDL_WaitAndAcquireGPUSwapchainTexture)
func (cb *GPUCommandBuffer) WaitAndAcquireGPUSwapchainTexture(window *Window) (*SwapchainTexture, error) {
	var texture SwapchainTexture

	if !iWaitAndAcquireGPUSwapchainTexture(cb, window, &texture.Texture, &texture.Width, &texture.Height) {
		return nil, internal.LastErr()
	}

	return &texture, nil
}

// SDL_SubmitGPUCommandBuffer - Submits a command buffer so its commands can be processed on the GPU.
// (https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBuffer)
func (cb *GPUCommandBuffer) Submit() error {
	if !iSubmitGPUCommandBuffer(cb) {
		return internal.LastErr()
	}

	return nil
}

// SDL_SubmitGPUCommandBufferAndAcquireFence - Submits a command buffer so its commands can be processed on the GPU, and acquires a fence associated with the command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_SubmitGPUCommandBufferAndAcquireFence)
func (cb *GPUCommandBuffer) SubmitAndAcquireFence() (*GPUFence, error) {
	fence := iSubmitGPUCommandBufferAndAcquireFence(cb)
	if fence == nil {
		return nil, internal.LastErr()
	}

	return fence, nil
}

// SDL_CancelGPUCommandBuffer - Cancels a command buffer.
// (https://wiki.libsdl.org/SDL3/SDL_CancelGPUCommandBuffer)
func (cb *GPUCommandBuffer) Cancel() error {
	if !iCancelGPUCommandBuffer(cb) {
		return internal.LastErr()
	}

	return nil
}

// Process

// SDL_GetProcessProperties - Get the properties associated with a process.
// (https://wiki.libsdl.org/SDL3/SDL_GetProcessProperties)
func (process *Process) Properties() (PropertiesID, error) {
	props := iGetProcessProperties(process)
	if props == 0 {
		return 0, internal.LastErr()
	}

	return props, nil
}

// SDL_ReadProcess - Read all the output from a process.
// (https://wiki.libsdl.org/SDL3/SDL_ReadProcess)
func (process *Process) Read() (*ProcessData, error) {
	var exitCode int32
	var size uintptr

	ptr := iReadProcess(process, &size, &exitCode)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return &ProcessData{
		ExitCode: exitCode,
		Data:     internal.ClonePtrSlice[byte](ptr, int(size)),
	}, nil
}

// SDL_GetProcessInput - Get the SDL_IOStream associated with process standard input.
// (https://wiki.libsdl.org/SDL3/SDL_GetProcessInput)
func (process *Process) Input() (*IOStream, error) {
	input := iGetProcessInput(process)
	if input == nil {
		return nil, internal.LastErr()
	}

	return input, nil
}

// SDL_GetProcessOutput - Get the SDL_IOStream associated with process standard output.
// (https://wiki.libsdl.org/SDL3/SDL_GetProcessOutput)
func (process *Process) Output() (*IOStream, error) {
	output := iGetProcessOutput(process)
	if output == nil {
		return nil, internal.LastErr()
	}

	return output, nil
}

// SDL_KillProcess - Stop a process.
// (https://wiki.libsdl.org/SDL3/SDL_KillProcess)
func (process *Process) Kill(force bool) error {
	if !iKillProcess(process, force) {
		return internal.LastErr()
	}

	return nil
}

// SDL_WaitProcess - Wait for a process to finish.
// (https://wiki.libsdl.org/SDL3/SDL_WaitProcess)
func (process *Process) Wait(block bool) (int32, bool) {
	var exitCode int32

	ret := iWaitProcess(process, block, &exitCode)

	return exitCode, ret
}

// SDL_DestroyProcess - Destroy a previously created process object.
// (https://wiki.libsdl.org/SDL3/SDL_DestroyProcess)
func (process *Process) Destroy() {
	iDestroyProcess(process)
}

// SensorID

// SDL_GetSensorNameForID - Get the implementation dependent name of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNameForID)
func (id SensorID) SensorName() string {
	return iGetSensorNameForID(id)
}

// SDL_GetSensorTypeForID - Get the type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorTypeForID)
func (id SensorID) SensorType() SensorType {
	return iGetSensorTypeForID(id)
}

// SDL_GetSensorNonPortableTypeForID - Get the platform dependent type of a sensor.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorNonPortableTypeForID)
func (id SensorID) SensorNonPortableType() int32 {
	return iGetSensorNonPortableTypeForID(id)
}

// SDL_OpenSensor - Open a sensor for use.
// (https://wiki.libsdl.org/SDL3/SDL_OpenSensor)
func (id SensorID) OpenSensor() (*Sensor, error) {
	sensor := iOpenSensor(id)
	if sensor == nil {
		return nil, internal.LastErr()
	}

	return sensor, nil
}

// SDL_GetSensorFromID - Return the SDL_Sensor associated with an instance ID.
// (https://wiki.libsdl.org/SDL3/SDL_GetSensorFromID)
func (id SensorID) Sensor() (*Sensor, error) {
	sensor := iGetSensorFromID(id)
	if sensor == nil {
		return nil, internal.LastErr()
	}

	return sensor, nil
}
