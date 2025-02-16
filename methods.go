/*

This file has been generated initially with cmd/methodgen
to have a starting point. Functions need to be made Go-like,
return errors, strings and slices cloned/freed if necessary based
on the documentation.

*/

package sdl

import (
	"image/color"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// ThreadPriority

func (priority ThreadPriority) SetCurrent() bool {
	panic("not implemented")
	return iSetCurrentThreadPriority(priority)
}

// IOStreamInterface

func (iface *IOStreamInterface) OpenIO(userdata *byte) *IOStream {
	panic("not implemented")
	//return iOpenIO(iface, userdata)
}

// ClipboardDataCallback

func (callback ClipboardDataCallback) SetClipboardData(cleanup ClipboardCleanupCallback, userdata *byte, mime_types *string, num_mime_types uintptr) bool {
	panic("not implemented")
	//return iSetClipboardData(callback, cleanup, userdata, mime_types, num_mime_types)
}

// WindowID

func (id WindowID) WindowFromID() *Window {
	panic("not implemented")
	return iGetWindowFromID(id)
}

// TouchID

func (touchID TouchID) TouchDeviceName() string {
	panic("not implemented")
	return iGetTouchDeviceName(touchID)
}

func (touchID TouchID) TouchDeviceType() TouchDeviceType {
	panic("not implemented")
	return iGetTouchDeviceType(touchID)
}

func (touchID TouchID) TouchFingers(count *int32) **Finger {
	panic("not implemented")
	//return iGetTouchFingers(touchID, count)
}

// GPUShaderFormat

func (format_flags GPUShaderFormat) GPUSupportsShaderFormats(name string) bool {
	panic("not implemented")
	return iGPUSupportsShaderFormats(format_flags, name)
}

func (format_flags GPUShaderFormat) CreateGPUDevice(debug_mode bool, name string) *GPUDevice {
	panic("not implemented")
	return iCreateGPUDevice(format_flags, debug_mode, name)
}

// Storage

func (storage *Storage) Close() bool {
	panic("not implemented")
	return iCloseStorage(storage)
}

func (storage *Storage) Ready() bool {
	panic("not implemented")
	return iStorageReady(storage)
}

func (storage *Storage) FileSize(path string, length *uint64) bool {
	panic("not implemented")
	return iGetStorageFileSize(storage, path, length)
}

func (storage *Storage) ReadFile(path string, destination *byte, length uint64) bool {
	panic("not implemented")
	//return iReadStorageFile(storage, path, destination, length)
}

func (storage *Storage) WriteFile(path string, source *byte, length uint64) bool {
	panic("not implemented")
	//return iWriteStorageFile(storage, path, source, length)
}

func (storage *Storage) CreateDirectory(path string) bool {
	panic("not implemented")
	return iCreateStorageDirectory(storage, path)
}

func (storage *Storage) EnumerateDirectory(path string, callback EnumerateDirectoryCallback, userdata *byte) bool {
	panic("not implemented")
	//return iEnumerateStorageDirectory(storage, path, callback, userdata)
}

func (storage *Storage) RemovePath(path string) bool {
	panic("not implemented")
	return iRemoveStoragePath(storage, path)
}

func (storage *Storage) RenamePath(oldpath string, newpath string) bool {
	panic("not implemented")
	return iRenameStoragePath(storage, oldpath, newpath)
}

func (storage *Storage) CopyFile(oldpath string, newpath string) bool {
	panic("not implemented")
	return iCopyStorageFile(storage, oldpath, newpath)
}

func (storage *Storage) PathInfo(path string, info *PathInfo) bool {
	panic("not implemented")
	return iGetStoragePathInfo(storage, path, info)
}

func (storage *Storage) SpaceRemaining() uint64 {
	panic("not implemented")
	return iGetStorageSpaceRemaining(storage)
}

func (storage *Storage) GlobDirectory(path string, pattern string, flags GlobFlags, count *int32) *string {
	panic("not implemented")
	//return iGlobStorageDirectory(storage, path, pattern, flags, count)
}

// AudioDeviceID

func (devid AudioDeviceID) AudioDeviceName() string {
	panic("not implemented")
	return iGetAudioDeviceName(devid)
}

func (devid AudioDeviceID) AudioDeviceFormat(spec *AudioSpec, sample_frames *int32) bool {
	panic("not implemented")
	return iGetAudioDeviceFormat(devid, spec, sample_frames)
}

func (devid AudioDeviceID) AudioDeviceChannelMap(count *int32) *int {
	panic("not implemented")
	//return iGetAudioDeviceChannelMap(devid, count)
}

func (devid AudioDeviceID) OpenAudioDevice(spec *AudioSpec) AudioDeviceID {
	panic("not implemented")
	return iOpenAudioDevice(devid, spec)
}

func (devid AudioDeviceID) IsAudioDevicePhysical() bool {
	panic("not implemented")
	return iIsAudioDevicePhysical(devid)
}

func (devid AudioDeviceID) IsAudioDevicePlayback() bool {
	panic("not implemented")
	return iIsAudioDevicePlayback(devid)
}

func (devid AudioDeviceID) PauseAudioDevice() bool {
	panic("not implemented")
	return iPauseAudioDevice(devid)
}

func (devid AudioDeviceID) ResumeAudioDevice() bool {
	panic("not implemented")
	return iResumeAudioDevice(devid)
}

func (devid AudioDeviceID) AudioDevicePaused() bool {
	panic("not implemented")
	return iAudioDevicePaused(devid)
}

func (devid AudioDeviceID) AudioDeviceGain() float32 {
	panic("not implemented")
	return iGetAudioDeviceGain(devid)
}

func (devid AudioDeviceID) SetAudioDeviceGain(gain float32) bool {
	panic("not implemented")
	return iSetAudioDeviceGain(devid, gain)
}

func (devid AudioDeviceID) CloseAudioDevice() {
	panic("not implemented")
	iCloseAudioDevice(devid)
}

func (devid AudioDeviceID) BindAudioStreams(streams **AudioStream, num_streams int32) bool {
	panic("not implemented")
	return iBindAudioStreams(devid, streams, num_streams)
}

func (devid AudioDeviceID) BindAudioStream(stream *AudioStream) bool {
	panic("not implemented")
	return iBindAudioStream(devid, stream)
}

func (devid AudioDeviceID) OpenAudioDeviceStream(spec *AudioSpec, callback AudioStreamCallback, userdata *byte) *AudioStream {
	panic("not implemented")
	//return iOpenAudioDeviceStream(devid, spec, callback, userdata)
}

func (devid AudioDeviceID) SetAudioPostmixCallback(callback AudioPostmixCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetAudioPostmixCallback(devid, callback, userdata)
}

// Camera

func (camera *Camera) PermissionState() int32 {
	panic("not implemented")
	return iGetCameraPermissionState(camera)
}

func (camera *Camera) ID() CameraID {
	panic("not implemented")
	return iGetCameraID(camera)
}

func (camera *Camera) Properties() PropertiesID {
	panic("not implemented")
	return iGetCameraProperties(camera)
}

func (camera *Camera) Format(spec *CameraSpec) bool {
	panic("not implemented")
	return iGetCameraFormat(camera, spec)
}

func (camera *Camera) AcquireFrame(timestampNS *uint64) *Surface {
	panic("not implemented")
	return iAcquireCameraFrame(camera, timestampNS)
}

func (camera *Camera) ReleaseFrame(frame *Surface) {
	panic("not implemented")
	iReleaseCameraFrame(camera, frame)
}

func (camera *Camera) Close() {
	panic("not implemented")
	iCloseCamera(camera)
}

// GamepadButton

func (button GamepadButton) GamepadStringForButton() string {
	panic("not implemented")
	return iGetGamepadStringForButton(button)
}

// GPUTextureFormat

func (format GPUTextureFormat) TexelBlockSize() uint32 {
	panic("not implemented")
	return iGPUTextureFormatTexelBlockSize(format)
}

func (format GPUTextureFormat) CalculateSize(width uint32, height uint32, depth_or_layer_count uint32) uint32 {
	panic("not implemented")
	return iCalculateGPUTextureFormatSize(format, width, height, depth_or_layer_count)
}

// SpinLock

func (lock *SpinLock) TryLockSpinlock() bool {
	panic("not implemented")
	return iTryLockSpinlock(lock)
}

func (lock *SpinLock) LockSpinlock() {
	panic("not implemented")
	iLockSpinlock(lock)
}

func (lock *SpinLock) UnlockSpinlock() {
	panic("not implemented")
	iUnlockSpinlock(lock)
}

// Thread

func (thread *Thread) Name() string {
	panic("not implemented")
	return iGetThreadName(thread)
}

func (thread *Thread) ID() ThreadID {
	panic("not implemented")
	return iGetThreadID(thread)
}

func (thread *Thread) Wait(status *int32) {
	panic("not implemented")
	iWaitThread(thread, status)
}

func (thread *Thread) State() ThreadState {
	panic("not implemented")
	return iGetThreadState(thread)
}

func (thread *Thread) Detach() {
	panic("not implemented")
	iDetachThread(thread)
}

// Sensor

func (sensor *Sensor) Properties() PropertiesID {
	panic("not implemented")
	return iGetSensorProperties(sensor)
}

func (sensor *Sensor) Name() string {
	panic("not implemented")
	return iGetSensorName(sensor)
}

func (sensor *Sensor) Type() SensorType {
	panic("not implemented")
	return iGetSensorType(sensor)
}

func (sensor *Sensor) NonPortableType() int32 {
	panic("not implemented")
	return iGetSensorNonPortableType(sensor)
}

func (sensor *Sensor) ID() SensorID {
	panic("not implemented")
	return iGetSensorID(sensor)
}

func (sensor *Sensor) Data(data *float32, num_values int32) bool {
	panic("not implemented")
	return iGetSensorData(sensor, data, num_values)
}

func (sensor *Sensor) Close() {
	panic("not implemented")
	iCloseSensor(sensor)
}

// GamepadAxis

func (axis GamepadAxis) GamepadStringForAxis() string {
	panic("not implemented")
	return iGetGamepadStringForAxis(axis)
}

// Cursor

func (cursor *Cursor) Destroy() {
	panic("not implemented")
	iDestroyCursor(cursor)
}

// X11EventHook

func (callback X11EventHook) Set(userdata *byte) {
	panic("not implemented")
	//iSetX11EventHook(callback, userdata)
}

// TLSID

func (id *TLSID) TLS() *byte {
	panic("not implemented")
	//return iGetTLS(id)
}

func (id *TLSID) SetTLS(value *byte, destructor TLSDestructorCallback) bool {
	panic("not implemented")
	//return iSetTLS(id, value, destructor)
}

// Rect

func (A *Rect) HasIntersection(B *Rect) bool {
	panic("not implemented")
	return iHasRectIntersection(A, B)
}

func (A *Rect) Intersection(B *Rect, result *Rect) bool {
	panic("not implemented")
	return iGetRectIntersection(A, B, result)
}

func (A *Rect) Union(B *Rect, result *Rect) bool {
	panic("not implemented")
	return iGetRectUnion(A, B, result)
}

func (rect *Rect) AndLineIntersection(X1 *int32, Y1 *int32, X2 *int32, Y2 *int32) bool {
	panic("not implemented")
	return iGetRectAndLineIntersection(rect, X1, Y1, X2, Y2)
}

// JoystickID

func (instance_id JoystickID) JoystickNameForID() string {
	panic("not implemented")
	return iGetJoystickNameForID(instance_id)
}

func (instance_id JoystickID) JoystickPathForID() string {
	panic("not implemented")
	return iGetJoystickPathForID(instance_id)
}

func (instance_id JoystickID) JoystickPlayerIndexForID() int32 {
	panic("not implemented")
	return iGetJoystickPlayerIndexForID(instance_id)
}

func (instance_id JoystickID) JoystickGUIDForID() GUID {
	panic("not implemented")
	return iGetJoystickGUIDForID(instance_id)
}

func (instance_id JoystickID) JoystickVendorForID() uint16 {
	panic("not implemented")
	return iGetJoystickVendorForID(instance_id)
}

func (instance_id JoystickID) JoystickProductForID() uint16 {
	panic("not implemented")
	return iGetJoystickProductForID(instance_id)
}

func (instance_id JoystickID) JoystickProductVersionForID() uint16 {
	panic("not implemented")
	return iGetJoystickProductVersionForID(instance_id)
}

func (instance_id JoystickID) JoystickTypeForID() JoystickType {
	panic("not implemented")
	return iGetJoystickTypeForID(instance_id)
}

func (instance_id JoystickID) OpenJoystick() *Joystick {
	panic("not implemented")
	return iOpenJoystick(instance_id)
}

func (instance_id JoystickID) JoystickFromID() *Joystick {
	panic("not implemented")
	return iGetJoystickFromID(instance_id)
}

func (instance_id JoystickID) DetachVirtualJoystick() bool {
	panic("not implemented")
	return iDetachVirtualJoystick(instance_id)
}

func (instance_id JoystickID) IsJoystickVirtual() bool {
	panic("not implemented")
	return iIsJoystickVirtual(instance_id)
}

func (instance_id JoystickID) SetGamepadMapping(mapping string) bool {
	panic("not implemented")
	return iSetGamepadMapping(instance_id, mapping)
}

func (instance_id JoystickID) IsGamepad() bool {
	panic("not implemented")
	return iIsGamepad(instance_id)
}

func (instance_id JoystickID) GamepadNameForID() string {
	panic("not implemented")
	return iGetGamepadNameForID(instance_id)
}

func (instance_id JoystickID) GamepadPathForID() string {
	panic("not implemented")
	return iGetGamepadPathForID(instance_id)
}

func (instance_id JoystickID) GamepadPlayerIndexForID() int32 {
	panic("not implemented")
	return iGetGamepadPlayerIndexForID(instance_id)
}

func (instance_id JoystickID) GamepadGUIDForID() GUID {
	panic("not implemented")
	return iGetGamepadGUIDForID(instance_id)
}

func (instance_id JoystickID) GamepadVendorForID() uint16 {
	panic("not implemented")
	return iGetGamepadVendorForID(instance_id)
}

func (instance_id JoystickID) GamepadProductForID() uint16 {
	panic("not implemented")
	return iGetGamepadProductForID(instance_id)
}

func (instance_id JoystickID) GamepadProductVersionForID() uint16 {
	panic("not implemented")
	return iGetGamepadProductVersionForID(instance_id)
}

func (instance_id JoystickID) GamepadTypeForID() GamepadType {
	panic("not implemented")
	return iGetGamepadTypeForID(instance_id)
}

func (instance_id JoystickID) RealGamepadTypeForID() GamepadType {
	panic("not implemented")
	return iGetRealGamepadTypeForID(instance_id)
}

func (instance_id JoystickID) GamepadMappingForID() string {
	panic("not implemented")
	//return iGetGamepadMappingForID(instance_id)
}

func (instance_id JoystickID) OpenGamepad() *Gamepad {
	panic("not implemented")
	return iOpenGamepad(instance_id)
}

func (instance_id JoystickID) GamepadFromID() *Gamepad {
	panic("not implemented")
	return iGetGamepadFromID(instance_id)
}

// AtomicInt

func (a *AtomicInt) CompareAndSwap(oldval int32, newval int32) bool {
	panic("not implemented")
	return iCompareAndSwapAtomicInt(a, oldval, newval)
}

func (a *AtomicInt) Set(v int32) int32 {
	panic("not implemented")
	return iSetAtomicInt(a, v)
}

func (a *AtomicInt) Get() int32 {
	panic("not implemented")
	return iGetAtomicInt(a)
}

func (a *AtomicInt) Add(v int32) int32 {
	panic("not implemented")
	return iAddAtomicInt(a, v)
}

// GPUComputePass

func (compute_pass *GPUComputePass) BindGPUComputePipeline(compute_pipeline *GPUComputePipeline) {
	panic("not implemented")
	iBindGPUComputePipeline(compute_pass, compute_pipeline)
}

func (compute_pass *GPUComputePass) BindGPUComputeSamplers(first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
	panic("not implemented")
	iBindGPUComputeSamplers(compute_pass, first_slot, texture_sampler_bindings, num_bindings)
}

func (compute_pass *GPUComputePass) BindGPUComputeStorageTextures(first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
	panic("not implemented")
	iBindGPUComputeStorageTextures(compute_pass, first_slot, storage_textures, num_bindings)
}

func (compute_pass *GPUComputePass) BindGPUComputeStorageBuffers(first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
	panic("not implemented")
	iBindGPUComputeStorageBuffers(compute_pass, first_slot, storage_buffers, num_bindings)
}

func (compute_pass *GPUComputePass) DispatchGPUCompute(groupcount_x uint32, groupcount_y uint32, groupcount_z uint32) {
	panic("not implemented")
	iDispatchGPUCompute(compute_pass, groupcount_x, groupcount_y, groupcount_z)
}

func (compute_pass *GPUComputePass) DispatchGPUComputeIndirect(buffer *GPUBuffer, offset uint32) {
	panic("not implemented")
	iDispatchGPUComputeIndirect(compute_pass, buffer, offset)
}

func (compute_pass *GPUComputePass) End() {
	panic("not implemented")
	iEndGPUComputePass(compute_pass)
}

// Texture

func (texture *Texture) Properties() PropertiesID {
	panic("not implemented")
	return iGetTextureProperties(texture)
}

func (texture *Texture) RendererFrom() *Renderer {
	panic("not implemented")
	return iGetRendererFromTexture(texture)
}

func (texture *Texture) Size(w *float32, h *float32) bool {
	panic("not implemented")
	return iGetTextureSize(texture, w, h)
}

func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) bool {
	panic("not implemented")
	return iSetTextureColorMod(texture, r, g, b)
}

func (texture *Texture) SetColorModFloat(r float32, g float32, b float32) error {
	if !iSetTextureColorModFloat(texture, r, g, b) {
		return internal.LastErr()
	}

	return nil
}

func (texture *Texture) ColorMod(r *uint8, g *uint8, b *uint8) bool {
	panic("not implemented")
	return iGetTextureColorMod(texture, r, g, b)
}

func (texture *Texture) ColorModFloat(r *float32, g *float32, b *float32) bool {
	panic("not implemented")
	return iGetTextureColorModFloat(texture, r, g, b)
}

func (texture *Texture) SetAlphaMod(alpha uint8) bool {
	panic("not implemented")
	return iSetTextureAlphaMod(texture, alpha)
}

func (texture *Texture) SetAlphaModFloat(alpha float32) bool {
	panic("not implemented")
	return iSetTextureAlphaModFloat(texture, alpha)
}

func (texture *Texture) AlphaMod(alpha *uint8) bool {
	panic("not implemented")
	return iGetTextureAlphaMod(texture, alpha)
}

func (texture *Texture) AlphaModFloat(alpha *float32) bool {
	panic("not implemented")
	return iGetTextureAlphaModFloat(texture, alpha)
}

func (texture *Texture) SetBlendMode(blendMode BlendMode) bool {
	panic("not implemented")
	return iSetTextureBlendMode(texture, blendMode)
}

func (texture *Texture) BlendMode(blendMode *BlendMode) bool {
	panic("not implemented")
	return iGetTextureBlendMode(texture, blendMode)
}

func (texture *Texture) SetScaleMode(scaleMode ScaleMode) bool {
	panic("not implemented")
	return iSetTextureScaleMode(texture, scaleMode)
}

func (texture *Texture) ScaleMode(scaleMode *ScaleMode) bool {
	panic("not implemented")
	return iGetTextureScaleMode(texture, scaleMode)
}

func (texture *Texture) Update(rect *Rect, pixels *byte, pitch int32) bool {
	panic("not implemented")
	//return iUpdateTexture(texture, rect, pixels, pitch)
}

func (texture *Texture) UpdateYUV(rect *Rect, Yplane *uint8, Ypitch int32, Uplane *uint8, Upitch int32, Vplane *uint8, Vpitch int32) bool {
	panic("not implemented")
	return iUpdateYUVTexture(texture, rect, Yplane, Ypitch, Uplane, Upitch, Vplane, Vpitch)
}

func (texture *Texture) UpdateNV(rect *Rect, Yplane *uint8, Ypitch int32, UVplane *uint8, UVpitch int32) bool {
	panic("not implemented")
	return iUpdateNVTexture(texture, rect, Yplane, Ypitch, UVplane, UVpitch)
}

func (texture *Texture) Lock(rect *Rect, pixels **byte, pitch *int32) bool {
	panic("not implemented")
	//return iLockTexture(texture, rect, pixels, pitch)
}

func (texture *Texture) LockToSurface(rect *Rect, surface **Surface) error {
	if !iLockTextureToSurface(texture, rect, surface) {
		return internal.LastErr()
	}

	return nil
}

func (texture *Texture) Unlock() {
	iUnlockTexture(texture)
}

func (texture *Texture) Destroy() {
	panic("not implemented")
	iDestroyTexture(texture)
}

// RWLock

func (rwlock *RWLock) LockForReading() {
	panic("not implemented")
	iLockRWLockForReading(rwlock)
}

func (rwlock *RWLock) LockForWriting() {
	panic("not implemented")
	iLockRWLockForWriting(rwlock)
}

func (rwlock *RWLock) TryLockForReading() bool {
	panic("not implemented")
	return iTryLockRWLockForReading(rwlock)
}

func (rwlock *RWLock) TryLockForWriting() bool {
	panic("not implemented")
	return iTryLockRWLockForWriting(rwlock)
}

func (rwlock *RWLock) Unlock() {
	panic("not implemented")
	iUnlockRWLock(rwlock)
}

func (rwlock *RWLock) Destroy() {
	panic("not implemented")
	iDestroyRWLock(rwlock)
}

// AudioFormat

func (format AudioFormat) Name() string {
	panic("not implemented")
	return iGetAudioFormatName(format)
}

func (format AudioFormat) SilenceValueForFormat() int32 {
	panic("not implemented")
	return iGetSilenceValueForFormat(format)
}

// Point

func (points *Point) RectEnclosings(count int32, clip *Rect, result *Rect) bool {
	panic("not implemented")
	return iGetRectEnclosingPoints(points, count, clip, result)
}

// GPURenderPass

func (render_pass *GPURenderPass) BindGPUGraphicsPipeline(graphics_pipeline *GPUGraphicsPipeline) {
	panic("not implemented")
	iBindGPUGraphicsPipeline(render_pass, graphics_pipeline)
}

func (render_pass *GPURenderPass) SetGPUViewport(viewport *GPUViewport) {
	panic("not implemented")
	iSetGPUViewport(render_pass, viewport)
}

func (render_pass *GPURenderPass) SetGPUScissor(scissor *Rect) {
	panic("not implemented")
	iSetGPUScissor(render_pass, scissor)
}

func (render_pass *GPURenderPass) SetGPUStencilReference(reference uint8) {
	panic("not implemented")
	iSetGPUStencilReference(render_pass, reference)
}

func (render_pass *GPURenderPass) BindGPUVertexBuffers(first_slot uint32, bindings *GPUBufferBinding, num_bindings uint32) {
	panic("not implemented")
	iBindGPUVertexBuffers(render_pass, first_slot, bindings, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUIndexBuffer(binding *GPUBufferBinding, index_element_size GPUIndexElementSize) {
	panic("not implemented")
	iBindGPUIndexBuffer(render_pass, binding, index_element_size)
}

func (render_pass *GPURenderPass) BindGPUVertexSamplers(first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
	panic("not implemented")
	iBindGPUVertexSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUVertexStorageTextures(first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
	panic("not implemented")
	iBindGPUVertexStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUVertexStorageBuffers(first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
	panic("not implemented")
	iBindGPUVertexStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUFragmentSamplers(first_slot uint32, texture_sampler_bindings *GPUTextureSamplerBinding, num_bindings uint32) {
	panic("not implemented")
	iBindGPUFragmentSamplers(render_pass, first_slot, texture_sampler_bindings, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUFragmentStorageTextures(first_slot uint32, storage_textures **GPUTexture, num_bindings uint32) {
	panic("not implemented")
	iBindGPUFragmentStorageTextures(render_pass, first_slot, storage_textures, num_bindings)
}

func (render_pass *GPURenderPass) BindGPUFragmentStorageBuffers(first_slot uint32, storage_buffers **GPUBuffer, num_bindings uint32) {
	panic("not implemented")
	iBindGPUFragmentStorageBuffers(render_pass, first_slot, storage_buffers, num_bindings)
}

func (render_pass *GPURenderPass) DrawGPUIndexedPrimitives(num_indices uint32, num_instances uint32, first_index uint32, vertex_offset int32, first_instance uint32) {
	panic("not implemented")
	iDrawGPUIndexedPrimitives(render_pass, num_indices, num_instances, first_index, vertex_offset, first_instance)
}

func (render_pass *GPURenderPass) DrawGPUPrimitives(num_vertices uint32, num_instances uint32, first_vertex uint32, first_instance uint32) {
	panic("not implemented")
	iDrawGPUPrimitives(render_pass, num_vertices, num_instances, first_vertex, first_instance)
}

func (render_pass *GPURenderPass) DrawGPUPrimitivesIndirect(buffer *GPUBuffer, offset uint32, draw_count uint32) {
	panic("not implemented")
	iDrawGPUPrimitivesIndirect(render_pass, buffer, offset, draw_count)
}

func (render_pass *GPURenderPass) DrawGPUIndexedPrimitivesIndirect(buffer *GPUBuffer, offset uint32, draw_count uint32) {
	panic("not implemented")
	iDrawGPUIndexedPrimitivesIndirect(render_pass, buffer, offset, draw_count)
}

func (render_pass *GPURenderPass) End() {
	panic("not implemented")
	iEndGPURenderPass(render_pass)
}

// AsyncIOQueue

func (queue *AsyncIOQueue) Destroy() {
	panic("not implemented")
	iDestroyAsyncIOQueue(queue)
}

func (queue *AsyncIOQueue) AsyncIOResult(outcome *AsyncIOOutcome) bool {
	panic("not implemented")
	return iGetAsyncIOResult(queue, outcome)
}

func (queue *AsyncIOQueue) WaitAsyncIOResult(outcome *AsyncIOOutcome, timeoutMS int32) bool {
	panic("not implemented")
	return iWaitAsyncIOResult(queue, outcome, timeoutMS)
}

func (queue *AsyncIOQueue) Signal() {
	panic("not implemented")
	iSignalAsyncIOQueue(queue)
}

// PixelFormatDetails

func (format *PixelFormatDetails) MapRGB(palette *Palette, r uint8, g uint8, b uint8) uint32 {
	panic("not implemented")
	return iMapRGB(format, palette, r, g, b)
}

func (format *PixelFormatDetails) MapRGBA(palette *Palette, r uint8, g uint8, b uint8, a uint8) uint32 {
	panic("not implemented")
	return iMapRGBA(format, palette, r, g, b, a)
}

// Surface

func (surface *Surface) Destroy() {
	iDestroySurface(surface)
}

func (surface *Surface) Properties() PropertiesID {
	panic("not implemented")
	return iGetSurfaceProperties(surface)
}

func (surface *Surface) SetColorspace(colorspace Colorspace) bool {
	panic("not implemented")
	return iSetSurfaceColorspace(surface, colorspace)
}

func (surface *Surface) Colorspace() Colorspace {
	panic("not implemented")
	return iGetSurfaceColorspace(surface)
}

func (surface *Surface) CreatePalette() *Palette {
	panic("not implemented")
	return iCreateSurfacePalette(surface)
}

func (surface *Surface) SetPalette(palette *Palette) bool {
	panic("not implemented")
	return iSetSurfacePalette(surface, palette)
}

func (surface *Surface) Palette() *Palette {
	panic("not implemented")
	return iGetSurfacePalette(surface)
}

func (surface *Surface) AddAlternateImage(image *Surface) bool {
	panic("not implemented")
	return iAddSurfaceAlternateImage(surface, image)
}

func (surface *Surface) HasAlternateImages() bool {
	panic("not implemented")
	return iSurfaceHasAlternateImages(surface)
}

func (surface *Surface) Images(count *int32) **Surface {
	panic("not implemented")
	//return iGetSurfaceImages(surface, count)
}

func (surface *Surface) RemoveAlternateImages() {
	panic("not implemented")
	iRemoveSurfaceAlternateImages(surface)
}

func (surface *Surface) Lock() bool {
	panic("not implemented")
	return iLockSurface(surface)
}

func (surface *Surface) Unlock() {
	panic("not implemented")
	iUnlockSurface(surface)
}

func (surface *Surface) SaveBMP_IO(dst *IOStream, closeio bool) bool {
	panic("not implemented")
	return iSaveBMP_IO(surface, dst, closeio)
}

func (surface *Surface) SaveBMP(file string) bool {
	panic("not implemented")
	return iSaveBMP(surface, file)
}

func (surface *Surface) SetRLE(enabled bool) bool {
	panic("not implemented")
	return iSetSurfaceRLE(surface, enabled)
}

func (surface *Surface) HasRLE() bool {
	panic("not implemented")
	return iSurfaceHasRLE(surface)
}

func (surface *Surface) SetColorKey(enabled bool, key uint32) bool {
	panic("not implemented")
	return iSetSurfaceColorKey(surface, enabled, key)
}

func (surface *Surface) HasColorKey() bool {
	panic("not implemented")
	return iSurfaceHasColorKey(surface)
}

func (surface *Surface) ColorKey(key *uint32) bool {
	panic("not implemented")
	return iGetSurfaceColorKey(surface, key)
}

func (surface *Surface) SetColorMod(r uint8, g uint8, b uint8) bool {
	panic("not implemented")
	return iSetSurfaceColorMod(surface, r, g, b)
}

func (surface *Surface) ColorMod(r *uint8, g *uint8, b *uint8) bool {
	panic("not implemented")
	return iGetSurfaceColorMod(surface, r, g, b)
}

func (surface *Surface) SetAlphaMod(alpha uint8) bool {
	panic("not implemented")
	return iSetSurfaceAlphaMod(surface, alpha)
}

func (surface *Surface) AlphaMod(alpha *uint8) bool {
	panic("not implemented")
	return iGetSurfaceAlphaMod(surface, alpha)
}

func (surface *Surface) SetBlendMode(blendMode BlendMode) bool {
	panic("not implemented")
	return iSetSurfaceBlendMode(surface, blendMode)
}

func (surface *Surface) BlendMode(blendMode *BlendMode) bool {
	panic("not implemented")
	return iGetSurfaceBlendMode(surface, blendMode)
}

func (surface *Surface) SetClipRect(rect *Rect) bool {
	panic("not implemented")
	return iSetSurfaceClipRect(surface, rect)
}

func (surface *Surface) ClipRect(rect *Rect) bool {
	panic("not implemented")
	return iGetSurfaceClipRect(surface, rect)
}

func (surface *Surface) Flip(flip FlipMode) bool {
	panic("not implemented")
	return iFlipSurface(surface, flip)
}

func (surface *Surface) Duplicate() *Surface {
	panic("not implemented")
	return iDuplicateSurface(surface)
}

func (surface *Surface) Scale(width int32, height int32, scaleMode ScaleMode) *Surface {
	panic("not implemented")
	return iScaleSurface(surface, width, height, scaleMode)
}

func (surface *Surface) Convert(format PixelFormat) *Surface {
	panic("not implemented")
	return iConvertSurface(surface, format)
}

func (surface *Surface) ConvertAndColorspace(format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) *Surface {
	panic("not implemented")
	return iConvertSurfaceAndColorspace(surface, format, palette, colorspace, props)
}

func (surface *Surface) PremultiplyAlpha(linear bool) bool {
	panic("not implemented")
	return iPremultiplySurfaceAlpha(surface, linear)
}

func (surface *Surface) Clear(r float32, g float32, b float32, a float32) bool {
	panic("not implemented")
	return iClearSurface(surface, r, g, b, a)
}

func (dst *Surface) FillRect(rect *Rect, color uint32) error {
	if !iFillSurfaceRect(dst, rect, color) {
		return internal.LastErr()
	}

	return nil
}

func (dst *Surface) FillRects(rects *Rect, count int32, color uint32) bool {
	panic("not implemented")
	return iFillSurfaceRects(dst, rects, count, color)
}

func (src *Surface) Blit(srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	panic("not implemented")
	return iBlitSurface(src, srcrect, dst, dstrect)
}

func (src *Surface) BlitUnchecked(srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	panic("not implemented")
	return iBlitSurfaceUnchecked(src, srcrect, dst, dstrect)
}

func (src *Surface) BlitScaled(srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
	panic("not implemented")
	return iBlitSurfaceScaled(src, srcrect, dst, dstrect, scaleMode)
}

func (src *Surface) BlitUncheckedScaled(srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) bool {
	panic("not implemented")
	return iBlitSurfaceUncheckedScaled(src, srcrect, dst, dstrect, scaleMode)
}

func (src *Surface) BlitTiled(srcrect *Rect, dst *Surface, dstrect *Rect) bool {
	panic("not implemented")
	return iBlitSurfaceTiled(src, srcrect, dst, dstrect)
}

func (src *Surface) BlitTiledWithScale(srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
	panic("not implemented")
	return iBlitSurfaceTiledWithScale(src, srcrect, scale, scaleMode, dst, dstrect)
}

func (src *Surface) Blit9Grid(srcrect *Rect, left_width int32, right_width int32, top_height int32, bottom_height int32, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) bool {
	panic("not implemented")
	return iBlitSurface9Grid(src, srcrect, left_width, right_width, top_height, bottom_height, scale, scaleMode, dst, dstrect)
}

func (surface *Surface) MapRGB(r uint8, g uint8, b uint8) uint32 {
	panic("not implemented")
	return iMapSurfaceRGB(surface, r, g, b)
}

func (surface *Surface) MapRGBA(r uint8, g uint8, b uint8, a uint8) uint32 {
	panic("not implemented")
	return iMapSurfaceRGBA(surface, r, g, b, a)
}

func (surface *Surface) ReadPixel(x int32, y int32, r *uint8, g *uint8, b *uint8, a *uint8) bool {
	panic("not implemented")
	return iReadSurfacePixel(surface, x, y, r, g, b, a)
}

func (surface *Surface) ReadPixelFloat(x int32, y int32, r *float32, g *float32, b *float32, a *float32) bool {
	panic("not implemented")
	return iReadSurfacePixelFloat(surface, x, y, r, g, b, a)
}

func (surface *Surface) WritePixel(x int32, y int32, r uint8, g uint8, b uint8, a uint8) bool {
	panic("not implemented")
	return iWriteSurfacePixel(surface, x, y, r, g, b, a)
}

func (surface *Surface) WritePixelFloat(x int32, y int32, r float32, g float32, b float32, a float32) bool {
	panic("not implemented")
	return iWriteSurfacePixelFloat(surface, x, y, r, g, b, a)
}

func (surface *Surface) CreateColorCursor(hot_x int32, hot_y int32) *Cursor {
	panic("not implemented")
	return iCreateColorCursor(surface, hot_x, hot_y)
}

func (surface *Surface) CreateSoftwareRenderer() *Renderer {
	panic("not implemented")
	return iCreateSoftwareRenderer(surface)
}

// Event

func (event *Event) Window() *Window {
	panic("not implemented")
	return iGetWindowFromEvent(event)
}

// GPUDevice

func (device *GPUDevice) Destroy() {
	panic("not implemented")
	iDestroyGPUDevice(device)
}

func (device *GPUDevice) Driver() string {
	panic("not implemented")
	return iGetGPUDeviceDriver(device)
}

func (device *GPUDevice) GPUShaderFormats() GPUShaderFormat {
	panic("not implemented")
	return iGetGPUShaderFormats(device)
}

func (device *GPUDevice) CreateGPUComputePipeline(createinfo *GPUComputePipelineCreateInfo) *GPUComputePipeline {
	panic("not implemented")
	return iCreateGPUComputePipeline(device, createinfo)
}

func (device *GPUDevice) CreateGPUGraphicsPipeline(createinfo *GPUGraphicsPipelineCreateInfo) *GPUGraphicsPipeline {
	panic("not implemented")
	return iCreateGPUGraphicsPipeline(device, createinfo)
}

func (device *GPUDevice) CreateGPUSampler(createinfo *GPUSamplerCreateInfo) *GPUSampler {
	panic("not implemented")
	return iCreateGPUSampler(device, createinfo)
}

func (device *GPUDevice) CreateGPUShader(createinfo *GPUShaderCreateInfo) *GPUShader {
	panic("not implemented")
	return iCreateGPUShader(device, createinfo)
}

func (device *GPUDevice) CreateGPUTexture(createinfo *GPUTextureCreateInfo) *GPUTexture {
	panic("not implemented")
	return iCreateGPUTexture(device, createinfo)
}

func (device *GPUDevice) CreateGPUBuffer(createinfo *GPUBufferCreateInfo) *GPUBuffer {
	panic("not implemented")
	return iCreateGPUBuffer(device, createinfo)
}

func (device *GPUDevice) CreateGPUTransferBuffer(createinfo *GPUTransferBufferCreateInfo) *GPUTransferBuffer {
	panic("not implemented")
	return iCreateGPUTransferBuffer(device, createinfo)
}

func (device *GPUDevice) SetGPUBufferName(buffer *GPUBuffer, text string) {
	panic("not implemented")
	iSetGPUBufferName(device, buffer, text)
}

func (device *GPUDevice) SetGPUTextureName(texture *GPUTexture, text string) {
	panic("not implemented")
	iSetGPUTextureName(device, texture, text)
}

func (device *GPUDevice) ReleaseGPUTexture(texture *GPUTexture) {
	panic("not implemented")
	iReleaseGPUTexture(device, texture)
}

func (device *GPUDevice) ReleaseGPUSampler(sampler *GPUSampler) {
	panic("not implemented")
	iReleaseGPUSampler(device, sampler)
}

func (device *GPUDevice) ReleaseGPUBuffer(buffer *GPUBuffer) {
	panic("not implemented")
	iReleaseGPUBuffer(device, buffer)
}

func (device *GPUDevice) ReleaseGPUTransferBuffer(transfer_buffer *GPUTransferBuffer) {
	panic("not implemented")
	iReleaseGPUTransferBuffer(device, transfer_buffer)
}

func (device *GPUDevice) ReleaseGPUComputePipeline(compute_pipeline *GPUComputePipeline) {
	panic("not implemented")
	iReleaseGPUComputePipeline(device, compute_pipeline)
}

func (device *GPUDevice) ReleaseGPUShader(shader *GPUShader) {
	panic("not implemented")
	iReleaseGPUShader(device, shader)
}

func (device *GPUDevice) ReleaseGPUGraphicsPipeline(graphics_pipeline *GPUGraphicsPipeline) {
	panic("not implemented")
	iReleaseGPUGraphicsPipeline(device, graphics_pipeline)
}

func (device *GPUDevice) AcquireGPUCommandBuffer() *GPUCommandBuffer {
	panic("not implemented")
	return iAcquireGPUCommandBuffer(device)
}

func (device *GPUDevice) MapGPUTransferBuffer(transfer_buffer *GPUTransferBuffer, cycle bool) *byte {
	panic("not implemented")
	//return iMapGPUTransferBuffer(device, transfer_buffer, cycle)
}

func (device *GPUDevice) UnmapGPUTransferBuffer(transfer_buffer *GPUTransferBuffer) {
	panic("not implemented")
	iUnmapGPUTransferBuffer(device, transfer_buffer)
}

func (device *GPUDevice) WindowSupportsGPUSwapchainComposition(window *Window, swapchain_composition GPUSwapchainComposition) bool {
	panic("not implemented")
	return iWindowSupportsGPUSwapchainComposition(device, window, swapchain_composition)
}

func (device *GPUDevice) WindowSupportsGPUPresentMode(window *Window, present_mode GPUPresentMode) bool {
	panic("not implemented")
	return iWindowSupportsGPUPresentMode(device, window, present_mode)
}

func (device *GPUDevice) ClaimWindowFor(window *Window) bool {
	panic("not implemented")
	return iClaimWindowForGPUDevice(device, window)
}

func (device *GPUDevice) ReleaseWindowFrom(window *Window) {
	panic("not implemented")
	iReleaseWindowFromGPUDevice(device, window)
}

func (device *GPUDevice) SetGPUSwapchainParameters(window *Window, swapchain_composition GPUSwapchainComposition, present_mode GPUPresentMode) bool {
	panic("not implemented")
	return iSetGPUSwapchainParameters(device, window, swapchain_composition, present_mode)
}

func (device *GPUDevice) SetGPUAllowedFramesInFlight(allowed_frames_in_flight uint32) bool {
	panic("not implemented")
	return iSetGPUAllowedFramesInFlight(device, allowed_frames_in_flight)
}

func (device *GPUDevice) GPUSwapchainTextureFormat(window *Window) GPUTextureFormat {
	panic("not implemented")
	return iGetGPUSwapchainTextureFormat(device, window)
}

func (device *GPUDevice) WaitForGPUSwapchain(window *Window) bool {
	panic("not implemented")
	return iWaitForGPUSwapchain(device, window)
}

func (device *GPUDevice) WaitForGPUIdle() bool {
	panic("not implemented")
	return iWaitForGPUIdle(device)
}

func (device *GPUDevice) WaitForGPUFences(wait_all bool, fences **GPUFence, num_fences uint32) bool {
	panic("not implemented")
	return iWaitForGPUFences(device, wait_all, fences, num_fences)
}

func (device *GPUDevice) QueryGPUFence(fence *GPUFence) bool {
	panic("not implemented")
	return iQueryGPUFence(device, fence)
}

func (device *GPUDevice) ReleaseGPUFence(fence *GPUFence) {
	panic("not implemented")
	iReleaseGPUFence(device, fence)
}

func (device *GPUDevice) GPUTextureSupportsFormat(format GPUTextureFormat, typ GPUTextureType, usage GPUTextureUsageFlags) bool {
	panic("not implemented")
	return iGPUTextureSupportsFormat(device, format, typ, usage)
}

func (device *GPUDevice) GPUTextureSupportsSampleCount(format GPUTextureFormat, sample_count GPUSampleCount) bool {
	panic("not implemented")
	return iGPUTextureSupportsSampleCount(device, format, sample_count)
}

// Haptic

func (haptic *Haptic) ID() HapticID {
	panic("not implemented")
	return iGetHapticID(haptic)
}

func (haptic *Haptic) Name() string {
	panic("not implemented")
	return iGetHapticName(haptic)
}

func (haptic *Haptic) Close() {
	panic("not implemented")
	iCloseHaptic(haptic)
}

func (haptic *Haptic) MaxEffects() int32 {
	panic("not implemented")
	return iGetMaxHapticEffects(haptic)
}

func (haptic *Haptic) MaxEffectsPlaying() int32 {
	panic("not implemented")
	return iGetMaxHapticEffectsPlaying(haptic)
}

func (haptic *Haptic) Features() uint32 {
	panic("not implemented")
	return iGetHapticFeatures(haptic)
}

func (haptic *Haptic) NumAxes() int32 {
	panic("not implemented")
	return iGetNumHapticAxes(haptic)
}

func (haptic *Haptic) EffectSupported(effect *HapticEffect) bool {
	panic("not implemented")
	return iHapticEffectSupported(haptic, effect)
}

func (haptic *Haptic) CreateEffect(effect *HapticEffect) int32 {
	panic("not implemented")
	return iCreateHapticEffect(haptic, effect)
}

func (haptic *Haptic) UpdateEffect(effect int32, data *HapticEffect) bool {
	panic("not implemented")
	return iUpdateHapticEffect(haptic, effect, data)
}

func (haptic *Haptic) RunEffect(effect int32, iterations uint32) bool {
	panic("not implemented")
	return iRunHapticEffect(haptic, effect, iterations)
}

func (haptic *Haptic) StopEffect(effect int32) bool {
	panic("not implemented")
	return iStopHapticEffect(haptic, effect)
}

func (haptic *Haptic) DestroyEffect(effect int32) {
	panic("not implemented")
	iDestroyHapticEffect(haptic, effect)
}

func (haptic *Haptic) EffectStatus(effect int32) bool {
	panic("not implemented")
	return iGetHapticEffectStatus(haptic, effect)
}

func (haptic *Haptic) SetGain(gain int32) bool {
	panic("not implemented")
	return iSetHapticGain(haptic, gain)
}

func (haptic *Haptic) SetAutocenter(autocenter int32) bool {
	panic("not implemented")
	return iSetHapticAutocenter(haptic, autocenter)
}

func (haptic *Haptic) Pause() bool {
	panic("not implemented")
	return iPauseHaptic(haptic)
}

func (haptic *Haptic) Resume() bool {
	panic("not implemented")
	return iResumeHaptic(haptic)
}

func (haptic *Haptic) StopEffects() bool {
	panic("not implemented")
	return iStopHapticEffects(haptic)
}

func (haptic *Haptic) RumbleSupported() bool {
	panic("not implemented")
	return iHapticRumbleSupported(haptic)
}

func (haptic *Haptic) InitRumble() bool {
	panic("not implemented")
	return iInitHapticRumble(haptic)
}

func (haptic *Haptic) PlayRumble(strength float32, length uint32) bool {
	panic("not implemented")
	return iPlayHapticRumble(haptic, strength, length)
}

func (haptic *Haptic) StopRumble() bool {
	panic("not implemented")
	return iStopHapticRumble(haptic)
}

// FileDialogType

func (typ FileDialogType) ShowFileDialogWithProperties(callback DialogFileCallback, userdata *byte, props PropertiesID) {
	panic("not implemented")
	//iShowFileDialogWithProperties(typ, callback, userdata, props)
}

// Gamepad

func (gamepad *Gamepad) Mapping() string {
	panic("not implemented")
	//return iGetGamepadMapping(gamepad)
}

func (gamepad *Gamepad) Properties() PropertiesID {
	panic("not implemented")
	return iGetGamepadProperties(gamepad)
}

func (gamepad *Gamepad) ID() JoystickID {
	panic("not implemented")
	return iGetGamepadID(gamepad)
}

func (gamepad *Gamepad) Name() string {
	panic("not implemented")
	return iGetGamepadName(gamepad)
}

func (gamepad *Gamepad) Path() string {
	panic("not implemented")
	return iGetGamepadPath(gamepad)
}

func (gamepad *Gamepad) Type() GamepadType {
	panic("not implemented")
	return iGetGamepadType(gamepad)
}

func (gamepad *Gamepad) RealType() GamepadType {
	panic("not implemented")
	return iGetRealGamepadType(gamepad)
}

func (gamepad *Gamepad) PlayerIndex() int32 {
	panic("not implemented")
	return iGetGamepadPlayerIndex(gamepad)
}

func (gamepad *Gamepad) SetPlayerIndex(player_index int32) bool {
	panic("not implemented")
	return iSetGamepadPlayerIndex(gamepad, player_index)
}

func (gamepad *Gamepad) Vendor() uint16 {
	panic("not implemented")
	return iGetGamepadVendor(gamepad)
}

func (gamepad *Gamepad) Product() uint16 {
	panic("not implemented")
	return iGetGamepadProduct(gamepad)
}

func (gamepad *Gamepad) ProductVersion() uint16 {
	panic("not implemented")
	return iGetGamepadProductVersion(gamepad)
}

func (gamepad *Gamepad) FirmwareVersion() uint16 {
	panic("not implemented")
	return iGetGamepadFirmwareVersion(gamepad)
}

func (gamepad *Gamepad) Serial() string {
	panic("not implemented")
	return iGetGamepadSerial(gamepad)
}

func (gamepad *Gamepad) SteamHandle() uint64 {
	panic("not implemented")
	return iGetGamepadSteamHandle(gamepad)
}

func (gamepad *Gamepad) ConnectionState() JoystickConnectionState {
	panic("not implemented")
	return iGetGamepadConnectionState(gamepad)
}

func (gamepad *Gamepad) PowerInfo(percent *int32) PowerState {
	panic("not implemented")
	return iGetGamepadPowerInfo(gamepad, percent)
}

func (gamepad *Gamepad) Connected() bool {
	panic("not implemented")
	return iGamepadConnected(gamepad)
}

func (gamepad *Gamepad) Joystick() *Joystick {
	panic("not implemented")
	return iGetGamepadJoystick(gamepad)
}

func (gamepad *Gamepad) Bindings() ([]*GamepadBinding, error) {
	var count int32

	ptr := iGetGamepadBindings(gamepad, &count)
	if ptr == 0 {
		return nil, internal.LastErr()
	}
	defer internal.Free(ptr)

	return internal.ClonePtrSlice[*GamepadBinding](ptr, int(count)), nil
}

func (gamepad *Gamepad) HasAxis(axis GamepadAxis) bool {
	panic("not implemented")
	return iGamepadHasAxis(gamepad, axis)
}

func (gamepad *Gamepad) Axis(axis GamepadAxis) int16 {
	panic("not implemented")
	return iGetGamepadAxis(gamepad, axis)
}

func (gamepad *Gamepad) HasButton(button GamepadButton) bool {
	panic("not implemented")
	return iGamepadHasButton(gamepad, button)
}

func (gamepad *Gamepad) Button(button GamepadButton) bool {
	panic("not implemented")
	return iGetGamepadButton(gamepad, button)
}

func (gamepad *Gamepad) ButtonLabel(button GamepadButton) GamepadButtonLabel {
	panic("not implemented")
	return iGetGamepadButtonLabel(gamepad, button)
}

func (gamepad *Gamepad) NumTouchpads() int32 {
	panic("not implemented")
	return iGetNumGamepadTouchpads(gamepad)
}

func (gamepad *Gamepad) NumTouchpadFingers(touchpad int32) int32 {
	panic("not implemented")
	return iGetNumGamepadTouchpadFingers(gamepad, touchpad)
}

func (gamepad *Gamepad) TouchpadFinger(touchpad int32, finger int32, down *bool, x *float32, y *float32, pressure *float32) bool {
	panic("not implemented")
	return iGetGamepadTouchpadFinger(gamepad, touchpad, finger, down, x, y, pressure)
}

func (gamepad *Gamepad) HasSensor(typ SensorType) bool {
	panic("not implemented")
	return iGamepadHasSensor(gamepad, typ)
}

func (gamepad *Gamepad) SetSensorEnabled(typ SensorType, enabled bool) bool {
	panic("not implemented")
	return iSetGamepadSensorEnabled(gamepad, typ, enabled)
}

func (gamepad *Gamepad) SensorEnabled(typ SensorType) bool {
	panic("not implemented")
	return iGamepadSensorEnabled(gamepad, typ)
}

func (gamepad *Gamepad) SensorDataRate(typ SensorType) float32 {
	panic("not implemented")
	return iGetGamepadSensorDataRate(gamepad, typ)
}

func (gamepad *Gamepad) SensorData(typ SensorType, data *float32, num_values int32) bool {
	panic("not implemented")
	return iGetGamepadSensorData(gamepad, typ, data, num_values)
}

func (gamepad *Gamepad) Rumble(low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
	panic("not implemented")
	return iRumbleGamepad(gamepad, low_frequency_rumble, high_frequency_rumble, duration_ms)
}

func (gamepad *Gamepad) RumbleTriggers(left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
	panic("not implemented")
	return iRumbleGamepadTriggers(gamepad, left_rumble, right_rumble, duration_ms)
}

func (gamepad *Gamepad) SetLED(red uint8, green uint8, blue uint8) bool {
	panic("not implemented")
	return iSetGamepadLED(gamepad, red, green, blue)
}

func (gamepad *Gamepad) SendEffect(data *byte, size int32) bool {
	panic("not implemented")
	//return iSendGamepadEffect(gamepad, data, size)
}

func (gamepad *Gamepad) Close() {
	panic("not implemented")
	iCloseGamepad(gamepad)
}

func (gamepad *Gamepad) AppleSFSymbolsNameForButton(button GamepadButton) string {
	panic("not implemented")
	return iGetGamepadAppleSFSymbolsNameForButton(gamepad, button)
}

func (gamepad *Gamepad) AppleSFSymbolsNameForAxis(axis GamepadAxis) string {
	panic("not implemented")
	return iGetGamepadAppleSFSymbolsNameForAxis(gamepad, axis)
}

// Keymod

func (modstate Keymod) SetModState() {
	panic("not implemented")
	iSetModState(modstate)
}

// GPUCopyPass

func (copy_pass *GPUCopyPass) UploadToGPUTexture(source *GPUTextureTransferInfo, destination *GPUTextureRegion, cycle bool) {
	panic("not implemented")
	iUploadToGPUTexture(copy_pass, source, destination, cycle)
}

func (copy_pass *GPUCopyPass) UploadToGPUBuffer(source *GPUTransferBufferLocation, destination *GPUBufferRegion, cycle bool) {
	panic("not implemented")
	iUploadToGPUBuffer(copy_pass, source, destination, cycle)
}

func (copy_pass *GPUCopyPass) CopyGPUTextureToTexture(source *GPUTextureLocation, destination *GPUTextureLocation, w uint32, h uint32, d uint32, cycle bool) {
	panic("not implemented")
	iCopyGPUTextureToTexture(copy_pass, source, destination, w, h, d, cycle)
}

func (copy_pass *GPUCopyPass) CopyGPUBufferToBuffer(source *GPUBufferLocation, destination *GPUBufferLocation, size uint32, cycle bool) {
	panic("not implemented")
	iCopyGPUBufferToBuffer(copy_pass, source, destination, size, cycle)
}

func (copy_pass *GPUCopyPass) DownloadFromGPUTexture(source *GPUTextureRegion, destination *GPUTextureTransferInfo) {
	panic("not implemented")
	iDownloadFromGPUTexture(copy_pass, source, destination)
}

func (copy_pass *GPUCopyPass) DownloadFromGPUBuffer(source *GPUBufferRegion, destination *GPUTransferBufferLocation) {
	panic("not implemented")
	iDownloadFromGPUBuffer(copy_pass, source, destination)
}

func (copy_pass *GPUCopyPass) End() {
	panic("not implemented")
	iEndGPUCopyPass(copy_pass)
}

// HapticID

func (instance_id HapticID) HapticNameForID() string {
	panic("not implemented")
	return iGetHapticNameForID(instance_id)
}

func (instance_id HapticID) OpenHaptic() *Haptic {
	panic("not implemented")
	return iOpenHaptic(instance_id)
}

func (instance_id HapticID) HapticFromID() *Haptic {
	panic("not implemented")
	return iGetHapticFromID(instance_id)
}

// Renderer

func (renderer *Renderer) RenderWindow() *Window {
	panic("not implemented")
	return iGetRenderWindow(renderer)
}

func (renderer *Renderer) Name() string {
	panic("not implemented")
	return iGetRendererName(renderer)
}

func (renderer *Renderer) Properties() PropertiesID {
	panic("not implemented")
	return iGetRendererProperties(renderer)
}

func (renderer *Renderer) RenderOutputSize(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetRenderOutputSize(renderer, w, h)
}

func (renderer *Renderer) CurrentRenderOutputSize(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetCurrentRenderOutputSize(renderer, w, h)
}

func (renderer *Renderer) CreateTexture(format PixelFormat, access TextureAccess, w, h int) (*Texture, error) {
	texture := iCreateTexture(renderer, format, access, int32(w), int32(h))
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	texture := iCreateTextureFromSurface(renderer, surface)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

func (renderer *Renderer) CreateTextureWithProperties(props PropertiesID) *Texture {
	panic("not implemented")
	return iCreateTextureWithProperties(renderer, props)
}

func (renderer *Renderer) SetRenderTarget(texture *Texture) bool {
	panic("not implemented")
	return iSetRenderTarget(renderer, texture)
}

func (renderer *Renderer) RenderTarget() *Texture {
	panic("not implemented")
	return iGetRenderTarget(renderer)
}

func (renderer *Renderer) SetRenderLogicalPresentation(w int32, h int32, mode RendererLogicalPresentation) bool {
	panic("not implemented")
	return iSetRenderLogicalPresentation(renderer, w, h, mode)
}

func (renderer *Renderer) RenderLogicalPresentation(w *int32, h *int32, mode *RendererLogicalPresentation) bool {
	panic("not implemented")
	return iGetRenderLogicalPresentation(renderer, w, h, mode)
}

func (renderer *Renderer) RenderLogicalPresentationRect(rect *FRect) bool {
	panic("not implemented")
	return iGetRenderLogicalPresentationRect(renderer, rect)
}

func (renderer *Renderer) RenderCoordinatesFromWindow(window_x float32, window_y float32, x *float32, y *float32) bool {
	panic("not implemented")
	return iRenderCoordinatesFromWindow(renderer, window_x, window_y, x, y)
}

func (renderer *Renderer) RenderCoordinatesToWindow(x float32, y float32, window_x *float32, window_y *float32) bool {
	panic("not implemented")
	return iRenderCoordinatesToWindow(renderer, x, y, window_x, window_y)
}

func (renderer *Renderer) ConvertEventToRenderCoordinates(event *Event) bool {
	panic("not implemented")
	return iConvertEventToRenderCoordinates(renderer, event)
}

func (renderer *Renderer) SetRenderViewport(rect *Rect) bool {
	panic("not implemented")
	return iSetRenderViewport(renderer, rect)
}

func (renderer *Renderer) RenderViewport(rect *Rect) bool {
	panic("not implemented")
	return iGetRenderViewport(renderer, rect)
}

func (renderer *Renderer) RenderViewportSet() bool {
	panic("not implemented")
	return iRenderViewportSet(renderer)
}

func (renderer *Renderer) RenderSafeArea(rect *Rect) bool {
	panic("not implemented")
	return iGetRenderSafeArea(renderer, rect)
}

func (renderer *Renderer) SetRenderClipRect(rect *Rect) bool {
	panic("not implemented")
	return iSetRenderClipRect(renderer, rect)
}

func (renderer *Renderer) RenderClipRect(rect *Rect) bool {
	panic("not implemented")
	return iGetRenderClipRect(renderer, rect)
}

func (renderer *Renderer) RenderClipEnabled() bool {
	panic("not implemented")
	return iRenderClipEnabled(renderer)
}

func (renderer *Renderer) SetRenderScale(scaleX float32, scaleY float32) bool {
	panic("not implemented")
	return iSetRenderScale(renderer, scaleX, scaleY)
}

func (renderer *Renderer) RenderScale(scaleX *float32, scaleY *float32) bool {
	panic("not implemented")
	return iGetRenderScale(renderer, scaleX, scaleY)
}

func (renderer *Renderer) SetRenderDrawColor(r, g, b, a uint8) error {
	if !iSetRenderDrawColor(renderer, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) SetRenderDrawColorFloat(r float32, g float32, b float32, a float32) error {
	if !iSetRenderDrawColorFloat(renderer, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderDrawColor(r *uint8, g *uint8, b *uint8, a *uint8) (color.RGBA, error) {
	var clr color.RGBA

	if !iGetRenderDrawColor(renderer, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

func (renderer *Renderer) RenderDrawColorFloat(r *float32, g *float32, b *float32, a *float32) bool {
	panic("not implemented")
	return iGetRenderDrawColorFloat(renderer, r, g, b, a)
}

func (renderer *Renderer) SetRenderColorScale(scale float32) bool {
	panic("not implemented")
	return iSetRenderColorScale(renderer, scale)
}

func (renderer *Renderer) RenderColorScale(scale *float32) bool {
	panic("not implemented")
	return iGetRenderColorScale(renderer, scale)
}

func (renderer *Renderer) SetRenderDrawBlendMode(blendMode BlendMode) bool {
	panic("not implemented")
	return iSetRenderDrawBlendMode(renderer, blendMode)
}

func (renderer *Renderer) RenderDrawBlendMode(blendMode *BlendMode) bool {
	panic("not implemented")
	return iGetRenderDrawBlendMode(renderer, blendMode)
}

func (renderer *Renderer) RenderClear() error {
	if !iRenderClear(renderer) {
		return internal.LastErr()
	}
	return nil
}

func (renderer *Renderer) RenderPoint(x float32, y float32) bool {
	panic("not implemented")
	return iRenderPoint(renderer, x, y)
}

func (renderer *Renderer) RenderPoints(points []FPoint) error {
	if !iRenderPoints(renderer, unsafe.SliceData(points), int32(len(points))) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderLine(x1 float32, y1 float32, x2 float32, y2 float32) error {
	if !iRenderLine(renderer, x1, y1, x2, y2) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderLines(points []FPoint) error {
	if !iRenderLines(renderer, unsafe.SliceData(points), int32(len(points))) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderRect(rect *FRect) error {
	if !iRenderRect(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderRects(rects []FRect) error {
	if !iRenderRects(renderer, unsafe.SliceData(rects), int32(len(rects))) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderFillRect(rect *FRect) error {
	if !iRenderFillRect(renderer, rect) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderFillRects(rects []FRect) error {
	if !iRenderFillRects(renderer, unsafe.SliceData(rects), int32(len(rects))) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderTexture(texture *Texture, srcrect *FRect, dstrect *FRect) error {
	if !iRenderTexture(renderer, texture, srcrect, dstrect) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderTextureRotated(texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) error {
	if !iRenderTextureRotated(renderer, texture, srcrect, dstrect, angle, center, flip) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderTextureAffine(texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) bool {
	panic("not implemented")
	return iRenderTextureAffine(renderer, texture, srcrect, origin, right, down)
}

func (renderer *Renderer) RenderTextureTiled(texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) bool {
	panic("not implemented")
	return iRenderTextureTiled(renderer, texture, srcrect, scale, dstrect)
}

func (renderer *Renderer) RenderTexture9Grid(texture *Texture, srcrect *FRect, left_width float32, right_width float32, top_height float32, bottom_height float32, scale float32, dstrect *FRect) bool {
	panic("not implemented")
	return iRenderTexture9Grid(renderer, texture, srcrect, left_width, right_width, top_height, bottom_height, scale, dstrect)
}

func (renderer *Renderer) RenderGeometry(texture *Texture, vertices []Vertex, indices []int32) error {
	if !iRenderGeometry(renderer, texture, unsafe.SliceData(vertices), int32(len(vertices)), unsafe.SliceData(indices), int32(len(indices))) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderGeometryRaw(texture *Texture, xy *float32, xy_stride int32, color *FColor, color_stride int32, uv *float32, uv_stride int32, num_vertices int32, indices *byte, num_indices int32, size_indices int32) bool {
	panic("not implemented")
	//return iRenderGeometryRaw(renderer, texture, xy, xy_stride, color, color_stride, uv, uv_stride, num_vertices, indices, num_indices, size_indices)
}

func (renderer *Renderer) RenderReadPixels(rect *Rect) *Surface {
	panic("not implemented")
	return iRenderReadPixels(renderer, rect)
}

func (renderer *Renderer) RenderPresent() error {
	if !iRenderPresent(renderer) {
		return internal.LastErr()
	}
	return nil
}

func (renderer *Renderer) Destroy() {
	iDestroyRenderer(renderer)
}

func (renderer *Renderer) Flush() bool {
	panic("not implemented")
	return iFlushRenderer(renderer)
}

func (renderer *Renderer) RenderMetalLayer() *byte {
	panic("not implemented")
	//return iGetRenderMetalLayer(renderer)
}

func (renderer *Renderer) RenderMetalCommandEncoder() *byte {
	panic("not implemented")
	//return iGetRenderMetalCommandEncoder(renderer)
}

func (renderer *Renderer) AddVulkanRenderSemaphores(wait_stage_mask uint32, wait_semaphore int64, signal_semaphore int64) bool {
	panic("not implemented")
	return iAddVulkanRenderSemaphores(renderer, wait_stage_mask, wait_semaphore, signal_semaphore)
}

func (renderer *Renderer) SetRenderVSync(vsync int32) bool {
	panic("not implemented")
	return iSetRenderVSync(renderer, vsync)
}

func (renderer *Renderer) RenderVSync(vsync *int32) bool {
	panic("not implemented")
	return iGetRenderVSync(renderer, vsync)
}

func (renderer *Renderer) RenderDebugText(x float32, y float32, str string) error {
	if !iRenderDebugText(renderer, x, y, str) {
		return internal.LastErr()
	}

	return nil
}

func (renderer *Renderer) RenderDebugTextFormat(x float32, y float32, fmt string) bool {
	return iRenderDebugTextFormat(renderer, x, y, fmt)
}

// DateFormat

func (dateFormat *DateFormat) DateTimeLocalePreferences(timeFormat *TimeFormat) bool {
	panic("not implemented")
	return iGetDateTimeLocalePreferences(dateFormat, timeFormat)
}

// Time

func (ticks *Time) Current() bool {
	panic("not implemented")
	return iGetCurrentTime(ticks)
}

// AsyncIO

func (asyncio *AsyncIO) Size() int64 {
	panic("not implemented")
	return iGetAsyncIOSize(asyncio)
}

func (asyncio *AsyncIO) Read(ptr *byte, offset uint64, size uint64, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iReadAsyncIO(asyncio, ptr, offset, size, queue, userdata)
}

func (asyncio *AsyncIO) Write(ptr *byte, offset uint64, size uint64, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iWriteAsyncIO(asyncio, ptr, offset, size, queue, userdata)
}

func (asyncio *AsyncIO) Close(flush bool, queue *AsyncIOQueue, userdata *byte) bool {
	panic("not implemented")
	//return iCloseAsyncIO(asyncio, flush, queue, userdata)
}

// InitState

func (state *InitState) ShouldInit() bool {
	panic("not implemented")
	return iShouldInit(state)
}

func (state *InitState) ShouldQuit() bool {
	panic("not implemented")
	return iShouldQuit(state)
}

func (state *InitState) SetInitialized(initialized bool) {
	panic("not implemented")
	iSetInitialized(state, initialized)
}

// CameraID

func (instance_id CameraID) CameraSupportedFormats(count *int32) **CameraSpec {
	panic("not implemented")
	//return iGetCameraSupportedFormats(instance_id, count)
}

func (instance_id CameraID) CameraName() string {
	panic("not implemented")
	return iGetCameraName(instance_id)
}

func (instance_id CameraID) CameraPosition() CameraPosition {
	panic("not implemented")
	return iGetCameraPosition(instance_id)
}

func (instance_id CameraID) OpenCamera(spec *CameraSpec) *Camera {
	panic("not implemented")
	return iOpenCamera(instance_id, spec)
}

// DisplayID

func (displayID DisplayID) DisplayProperties() PropertiesID {
	panic("not implemented")
	return iGetDisplayProperties(displayID)
}

func (displayID DisplayID) DisplayName() string {
	panic("not implemented")
	return iGetDisplayName(displayID)
}

func (displayID DisplayID) DisplayBounds(rect *Rect) bool {
	panic("not implemented")
	return iGetDisplayBounds(displayID, rect)
}

func (displayID DisplayID) DisplayUsableBounds(rect *Rect) bool {
	panic("not implemented")
	return iGetDisplayUsableBounds(displayID, rect)
}

func (displayID DisplayID) NaturalDisplayOrientation() DisplayOrientation {
	panic("not implemented")
	return iGetNaturalDisplayOrientation(displayID)
}

func (displayID DisplayID) CurrentDisplayOrientation() DisplayOrientation {
	panic("not implemented")
	return iGetCurrentDisplayOrientation(displayID)
}

func (displayID DisplayID) DisplayContentScale() float32 {
	panic("not implemented")
	return iGetDisplayContentScale(displayID)
}

func (displayID DisplayID) FullscreenDisplayModes(count *int32) **DisplayMode {
	panic("not implemented")
	//return iGetFullscreenDisplayModes(displayID, count)
}

func (displayID DisplayID) ClosestFullscreenDisplayMode(w int32, h int32, refresh_rate float32, include_high_density_modes bool, closest *DisplayMode) bool {
	panic("not implemented")
	return iGetClosestFullscreenDisplayMode(displayID, w, h, refresh_rate, include_high_density_modes, closest)
}

func (displayID DisplayID) DesktopDisplayMode() *DisplayMode {
	panic("not implemented")
	return iGetDesktopDisplayMode(displayID)
}

func (displayID DisplayID) CurrentDisplayMode() *DisplayMode {
	panic("not implemented")
	return iGetCurrentDisplayMode(displayID)
}

// KeyboardID

func (instance_id KeyboardID) KeyboardNameForID() string {
	panic("not implemented")
	return iGetKeyboardNameForID(instance_id)
}

// MouseID

func (instance_id MouseID) MouseNameForID() string {
	panic("not implemented")
	return iGetMouseNameForID(instance_id)
}

// EventFilter

func (filter *EventFilter) Get(userdata **byte) bool {
	panic("not implemented")
	//return iGetEventFilter(filter, userdata)
}

// AudioStream

func (stream *AudioStream) Unbind() {
	panic("not implemented")
	iUnbindAudioStream(stream)
}

func (stream *AudioStream) Device() AudioDeviceID {
	panic("not implemented")
	return iGetAudioStreamDevice(stream)
}

func (stream *AudioStream) Properties() PropertiesID {
	panic("not implemented")
	return iGetAudioStreamProperties(stream)
}

func (stream *AudioStream) Format(src_spec *AudioSpec, dst_spec *AudioSpec) bool {
	panic("not implemented")
	return iGetAudioStreamFormat(stream, src_spec, dst_spec)
}

func (stream *AudioStream) SetFormat(src_spec *AudioSpec, dst_spec *AudioSpec) bool {
	panic("not implemented")
	return iSetAudioStreamFormat(stream, src_spec, dst_spec)
}

func (stream *AudioStream) FrequencyRatio() float32 {
	panic("not implemented")
	return iGetAudioStreamFrequencyRatio(stream)
}

func (stream *AudioStream) SetFrequencyRatio(ratio float32) bool {
	panic("not implemented")
	return iSetAudioStreamFrequencyRatio(stream, ratio)
}

func (stream *AudioStream) Gain() float32 {
	panic("not implemented")
	return iGetAudioStreamGain(stream)
}

func (stream *AudioStream) SetGain(gain float32) bool {
	panic("not implemented")
	return iSetAudioStreamGain(stream, gain)
}

func (stream *AudioStream) InputChannelMap(count *int32) *int {
	panic("not implemented")
	//return iGetAudioStreamInputChannelMap(stream, count)
}

func (stream *AudioStream) OutputChannelMap(count *int32) *int {
	panic("not implemented")
	//return iGetAudioStreamOutputChannelMap(stream, count)
}

func (stream *AudioStream) SetInputChannelMap(chmap *int32, count int32) bool {
	panic("not implemented")
	return iSetAudioStreamInputChannelMap(stream, chmap, count)
}

func (stream *AudioStream) SetOutputChannelMap(chmap *int32, count int32) bool {
	panic("not implemented")
	return iSetAudioStreamOutputChannelMap(stream, chmap, count)
}

func (stream *AudioStream) PutData(buf *byte, len int32) bool {
	panic("not implemented")
	//return iPutAudioStreamData(stream, buf, len)
}

func (stream *AudioStream) Data(buf *byte, len int32) int32 {
	panic("not implemented")
	//return iGetAudioStreamData(stream, buf, len)
}

func (stream *AudioStream) Available() int32 {
	panic("not implemented")
	return iGetAudioStreamAvailable(stream)
}

func (stream *AudioStream) Queued() int32 {
	panic("not implemented")
	return iGetAudioStreamQueued(stream)
}

func (stream *AudioStream) Flush() bool {
	panic("not implemented")
	return iFlushAudioStream(stream)
}

func (stream *AudioStream) Clear() bool {
	panic("not implemented")
	return iClearAudioStream(stream)
}

func (stream *AudioStream) PauseDevice() bool {
	panic("not implemented")
	return iPauseAudioStreamDevice(stream)
}

func (stream *AudioStream) ResumeDevice() bool {
	panic("not implemented")
	return iResumeAudioStreamDevice(stream)
}

func (stream *AudioStream) DevicePaused() bool {
	panic("not implemented")
	return iAudioStreamDevicePaused(stream)
}

func (stream *AudioStream) Lock() bool {
	panic("not implemented")
	return iLockAudioStream(stream)
}

func (stream *AudioStream) Unlock() bool {
	panic("not implemented")
	return iUnlockAudioStream(stream)
}

func (stream *AudioStream) SetGetCallback(callback AudioStreamCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetAudioStreamGetCallback(stream, callback, userdata)
}

func (stream *AudioStream) SetPutCallback(callback AudioStreamCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetAudioStreamPutCallback(stream, callback, userdata)
}

func (stream *AudioStream) Destroy() {
	panic("not implemented")
	iDestroyAudioStream(stream)
}

// FRect

func (A *FRect) HasRectIntersectionFloat(B *FRect) bool {
	panic("not implemented")
	return iHasRectIntersectionFloat(A, B)
}

func (A *FRect) RectIntersectionFloat(B *FRect, result *FRect) bool {
	panic("not implemented")
	return iGetRectIntersectionFloat(A, B, result)
}

func (A *FRect) RectUnionFloat(B *FRect, result *FRect) bool {
	panic("not implemented")
	return iGetRectUnionFloat(A, B, result)
}

func (rect *FRect) RectAndLineIntersectionFloat(X1 *float32, Y1 *float32, X2 *float32, Y2 *float32) bool {
	panic("not implemented")
	return iGetRectAndLineIntersectionFloat(rect, X1, Y1, X2, Y2)
}

// LogOutputFunction

func (callback LogOutputFunction) Set(userdata *byte) {
	panic("not implemented")
	//iSetLogOutputFunction(callback, userdata)
}

// Condition

func (cond *Condition) Destroy() {
	panic("not implemented")
	iDestroyCondition(cond)
}

func (cond *Condition) Signal() {
	panic("not implemented")
	iSignalCondition(cond)
}

func (cond *Condition) Broadcast() {
	panic("not implemented")
	iBroadcastCondition(cond)
}

func (cond *Condition) Wait(mutex *Mutex) {
	panic("not implemented")
	iWaitCondition(cond, mutex)
}

func (cond *Condition) WaitTimeout(mutex *Mutex, timeoutMS int32) bool {
	panic("not implemented")
	return iWaitConditionTimeout(cond, mutex, timeoutMS)
}

// Window

func (window *Window) PixelDensity() float32 {
	panic("not implemented")
	return iGetWindowPixelDensity(window)
}

func (window *Window) DisplayScale() float32 {
	panic("not implemented")
	return iGetWindowDisplayScale(window)
}

func (window *Window) SetFullscreenMode(mode *DisplayMode) bool {
	panic("not implemented")
	return iSetWindowFullscreenMode(window, mode)
}

func (window *Window) FullscreenMode() *DisplayMode {
	panic("not implemented")
	return iGetWindowFullscreenMode(window)
}

func (window *Window) ICCProfile(size *uintptr) *byte {
	panic("not implemented")
	//return iGetWindowICCProfile(window, size)
}

func (window *Window) PixelFormat() PixelFormat {
	panic("not implemented")
	return iGetWindowPixelFormat(window)
}

func (parent *Window) CreatePopup(offset_x int32, offset_y int32, w int32, h int32, flags WindowFlags) *Window {
	panic("not implemented")
	return iCreatePopupWindow(parent, offset_x, offset_y, w, h, flags)
}

func (window *Window) ID() WindowID {
	panic("not implemented")
	return iGetWindowID(window)
}

func (window *Window) Parent() *Window {
	panic("not implemented")
	return iGetWindowParent(window)
}

func (window *Window) Properties() PropertiesID {
	panic("not implemented")
	return iGetWindowProperties(window)
}

func (window *Window) Flags() WindowFlags {
	panic("not implemented")
	return iGetWindowFlags(window)
}

func (window *Window) SetTitle(title string) bool {
	panic("not implemented")
	return iSetWindowTitle(window, title)
}

func (window *Window) Title() string {
	panic("not implemented")
	return iGetWindowTitle(window)
}

func (window *Window) SetIcon(icon *Surface) bool {
	panic("not implemented")
	return iSetWindowIcon(window, icon)
}

func (window *Window) SetPosition(x int32, y int32) bool {
	panic("not implemented")
	return iSetWindowPosition(window, x, y)
}

func (window *Window) Position(x *int32, y *int32) bool {
	panic("not implemented")
	return iGetWindowPosition(window, x, y)
}

func (window *Window) SetSize(w int32, h int32) bool {
	panic("not implemented")
	return iSetWindowSize(window, w, h)
}

func (window *Window) Size(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetWindowSize(window, w, h)
}

func (window *Window) SafeArea(rect *Rect) bool {
	panic("not implemented")
	return iGetWindowSafeArea(window, rect)
}

func (window *Window) SetAspectRatio(min_aspect float32, max_aspect float32) bool {
	panic("not implemented")
	return iSetWindowAspectRatio(window, min_aspect, max_aspect)
}

func (window *Window) AspectRatio(min_aspect *float32, max_aspect *float32) bool {
	panic("not implemented")
	return iGetWindowAspectRatio(window, min_aspect, max_aspect)
}

func (window *Window) BordersSize(top *int32, left *int32, bottom *int32, right *int32) bool {
	panic("not implemented")
	return iGetWindowBordersSize(window, top, left, bottom, right)
}

func (window *Window) SizeInPixels(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetWindowSizeInPixels(window, w, h)
}

func (window *Window) SetMinimumSize(min_w int32, min_h int32) bool {
	panic("not implemented")
	return iSetWindowMinimumSize(window, min_w, min_h)
}

func (window *Window) MinimumSize(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetWindowMinimumSize(window, w, h)
}

func (window *Window) SetMaximumSize(max_w int32, max_h int32) bool {
	panic("not implemented")
	return iSetWindowMaximumSize(window, max_w, max_h)
}

func (window *Window) MaximumSize(w *int32, h *int32) bool {
	panic("not implemented")
	return iGetWindowMaximumSize(window, w, h)
}

func (window *Window) SetBordered(bordered bool) bool {
	panic("not implemented")
	return iSetWindowBordered(window, bordered)
}

func (window *Window) SetResizable(resizable bool) bool {
	panic("not implemented")
	return iSetWindowResizable(window, resizable)
}

func (window *Window) SetAlwaysOnTop(on_top bool) bool {
	panic("not implemented")
	return iSetWindowAlwaysOnTop(window, on_top)
}

func (window *Window) Show() bool {
	panic("not implemented")
	return iShowWindow(window)
}

func (window *Window) Hide() bool {
	panic("not implemented")
	return iHideWindow(window)
}

func (window *Window) Raise() bool {
	panic("not implemented")
	return iRaiseWindow(window)
}

func (window *Window) Maximize() bool {
	panic("not implemented")
	return iMaximizeWindow(window)
}

func (window *Window) Minimize() bool {
	panic("not implemented")
	return iMinimizeWindow(window)
}

func (window *Window) Restore() bool {
	panic("not implemented")
	return iRestoreWindow(window)
}

func (window *Window) SetFullscreen(fullscreen bool) bool {
	panic("not implemented")
	return iSetWindowFullscreen(window, fullscreen)
}

func (window *Window) Sync() bool {
	panic("not implemented")
	return iSyncWindow(window)
}

func (window *Window) HasSurface() bool {
	panic("not implemented")
	return iWindowHasSurface(window)
}

func (window *Window) Surface() *Surface {
	panic("not implemented")
	return iGetWindowSurface(window)
}

func (window *Window) SetSurfaceVSync(vsync int32) bool {
	panic("not implemented")
	return iSetWindowSurfaceVSync(window, vsync)
}

func (window *Window) SurfaceVSync(vsync *int32) bool {
	panic("not implemented")
	return iGetWindowSurfaceVSync(window, vsync)
}

func (window *Window) UpdateSurface() bool {
	panic("not implemented")
	return iUpdateWindowSurface(window)
}

func (window *Window) UpdateSurfaceRects(rects *Rect, numrects int32) bool {
	panic("not implemented")
	return iUpdateWindowSurfaceRects(window, rects, numrects)
}

func (window *Window) DestroySurface() bool {
	panic("not implemented")
	return iDestroyWindowSurface(window)
}

func (window *Window) SetKeyboardGrab(grabbed bool) bool {
	panic("not implemented")
	return iSetWindowKeyboardGrab(window, grabbed)
}

func (window *Window) SetMouseGrab(grabbed bool) bool {
	panic("not implemented")
	return iSetWindowMouseGrab(window, grabbed)
}

func (window *Window) KeyboardGrab() bool {
	panic("not implemented")
	return iGetWindowKeyboardGrab(window)
}

func (window *Window) MouseGrab() bool {
	panic("not implemented")
	return iGetWindowMouseGrab(window)
}

func (window *Window) SetMouseRect(rect *Rect) bool {
	panic("not implemented")
	return iSetWindowMouseRect(window, rect)
}

func (window *Window) MouseRect() *Rect {
	panic("not implemented")
	return iGetWindowMouseRect(window)
}

func (window *Window) SetOpacity(opacity float32) bool {
	panic("not implemented")
	return iSetWindowOpacity(window, opacity)
}

func (window *Window) Opacity() float32 {
	panic("not implemented")
	return iGetWindowOpacity(window)
}

func (window *Window) SetParent(parent *Window) bool {
	panic("not implemented")
	return iSetWindowParent(window, parent)
}

func (window *Window) SetModal(modal bool) bool {
	panic("not implemented")
	return iSetWindowModal(window, modal)
}

func (window *Window) SetFocusable(focusable bool) bool {
	panic("not implemented")
	return iSetWindowFocusable(window, focusable)
}

func (window *Window) ShowSystemMenu(x int32, y int32) bool {
	panic("not implemented")
	return iShowWindowSystemMenu(window, x, y)
}

func (window *Window) SetHitTest(callback HitTest, callback_data *byte) bool {
	panic("not implemented")
	//return iSetWindowHitTest(window, callback, callback_data)
}

func (window *Window) SetShape(shape *Surface) bool {
	panic("not implemented")
	return iSetWindowShape(window, shape)
}

func (window *Window) Flash(operation FlashOperation) bool {
	panic("not implemented")
	return iFlashWindow(window, operation)
}

func (window *Window) Destroy() {
	iDestroyWindow(window)
}

func (window *Window) GL_CreateContext() GLContext {
	panic("not implemented")
	return iGL_CreateContext(window)
}

func (window *Window) GL_MakeCurrent(context GLContext) bool {
	panic("not implemented")
	return iGL_MakeCurrent(window, context)
}

func (window *Window) EGL_GetSurface() EGLSurface {
	panic("not implemented")
	return iEGL_GetWindowSurface(window)
}

func (window *Window) GL_Swap() bool {
	panic("not implemented")
	return iGL_SwapWindow(window)
}

func (window *Window) StartTextInput() bool {
	panic("not implemented")
	return iStartTextInput(window)
}

func (window *Window) StartTextInputWithProperties(props PropertiesID) bool {
	panic("not implemented")
	return iStartTextInputWithProperties(window, props)
}

func (window *Window) TextInputActive() bool {
	panic("not implemented")
	return iTextInputActive(window)
}

func (window *Window) StopTextInput() bool {
	panic("not implemented")
	return iStopTextInput(window)
}

func (window *Window) ClearComposition() bool {
	panic("not implemented")
	return iClearComposition(window)
}

func (window *Window) SetTextInputArea(rect *Rect, cursor int32) bool {
	panic("not implemented")
	return iSetTextInputArea(window, rect, cursor)
}

func (window *Window) TextInputArea(rect *Rect, cursor *int32) bool {
	panic("not implemented")
	return iGetTextInputArea(window, rect, cursor)
}

func (window *Window) ScreenKeyboardShown() bool {
	panic("not implemented")
	return iScreenKeyboardShown(window)
}

func (window *Window) WarpMouseIn(x float32, y float32) {
	panic("not implemented")
	iWarpMouseInWindow(window, x, y)
}

func (window *Window) SetRelativeMouseMode(enabled bool) bool {
	panic("not implemented")
	return iSetWindowRelativeMouseMode(window, enabled)
}

func (window *Window) RelativeMouseMode() bool {
	panic("not implemented")
	return iGetWindowRelativeMouseMode(window)
}

func (window *Window) Metal_CreateView() MetalView {
	panic("not implemented")
	return iMetal_CreateView(window)
}

func (window *Window) CreateRenderer(name string) (*Renderer, error) {
	renderer := iCreateRenderer(window, name)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

func (window *Window) Renderer() *Renderer {
	return iGetRenderer(window)
}

// GLAttr

func (attr GLAttr) GL_SetAttribute(value int32) bool {
	panic("not implemented")
	return iGL_SetAttribute(attr, value)
}

func (attr GLAttr) GL_GetAttribute(value *int32) bool {
	panic("not implemented")
	return iGL_GetAttribute(attr, value)
}

// Scancode

func (scancode Scancode) KeyFrom(modstate Keymod, key_event bool) Keycode {
	panic("not implemented")
	return iGetKeyFromScancode(scancode, modstate, key_event)
}

func (scancode Scancode) SetName(name string) bool {
	panic("not implemented")
	return iSetScancodeName(scancode, name)
}

func (scancode Scancode) Name() string {
	panic("not implemented")
	return iGetScancodeName(scancode)
}

// SharedObject

func (handle *SharedObject) LoadFunction(name string) FunctionPointer {
	panic("not implemented")
	return iLoadFunction(handle, name)
}

func (handle *SharedObject) UnloadObject() {
	panic("not implemented")
	iUnloadObject(handle)
}

// TimerID

func (id TimerID) RemoveTimer() bool {
	panic("not implemented")
	return iRemoveTimer(id)
}

// IOStream

func (context *IOStream) CloseIO() bool {
	panic("not implemented")
	return iCloseIO(context)
}

func (context *IOStream) IOProperties() PropertiesID {
	panic("not implemented")
	return iGetIOProperties(context)
}

func (context *IOStream) IOStatus() IOStatus {
	panic("not implemented")
	return iGetIOStatus(context)
}

func (context *IOStream) IOSize() int64 {
	panic("not implemented")
	return iGetIOSize(context)
}

func (context *IOStream) SeekIO(offset int64, whence IOWhence) int64 {
	panic("not implemented")
	return iSeekIO(context, offset, whence)
}

func (context *IOStream) TellIO() int64 {
	panic("not implemented")
	return iTellIO(context)
}

func (context *IOStream) ReadIO(ptr *byte, size uintptr) uintptr {
	panic("not implemented")
	//return iReadIO(context, ptr, size)
}

func (context *IOStream) WriteIO(ptr *byte, size uintptr) uintptr {
	panic("not implemented")
	//return iWriteIO(context, ptr, size)
}

func (context *IOStream) IOprintf(fmt string) uintptr {
	panic("not implemented")
	return iIOprintf(context, fmt)
}

func (context *IOStream) IOvprintf(fmt string, ap va_list) uintptr {
	panic("not implemented")
	return iIOvprintf(context, fmt, ap)
}

func (context *IOStream) FlushIO() bool {
	panic("not implemented")
	return iFlushIO(context)
}

func (src *IOStream) LoadFile_IO(datasize *uintptr, closeio bool) *byte {
	panic("not implemented")
	//return iLoadFile_IO(src, datasize, closeio)
}

func (src *IOStream) SaveFile_IO(data *byte, datasize uintptr, closeio bool) bool {
	panic("not implemented")
	//return iSaveFile_IO(src, data, datasize, closeio)
}

func (src *IOStream) ReadU8(value *uint8) bool {
	panic("not implemented")
	return iReadU8(src, value)
}

func (src *IOStream) ReadS8(value *int8) bool {
	panic("not implemented")
	return iReadS8(src, value)
}

func (src *IOStream) ReadU16LE(value *uint16) bool {
	panic("not implemented")
	return iReadU16LE(src, value)
}

func (src *IOStream) ReadS16LE(value *int16) bool {
	panic("not implemented")
	return iReadS16LE(src, value)
}

func (src *IOStream) ReadU16BE(value *uint16) bool {
	panic("not implemented")
	return iReadU16BE(src, value)
}

func (src *IOStream) ReadS16BE(value *int16) bool {
	panic("not implemented")
	return iReadS16BE(src, value)
}

func (src *IOStream) ReadU32LE(value *uint32) bool {
	panic("not implemented")
	return iReadU32LE(src, value)
}

func (src *IOStream) ReadS32LE(value *int32) bool {
	panic("not implemented")
	return iReadS32LE(src, value)
}

func (src *IOStream) ReadU32BE(value *uint32) bool {
	panic("not implemented")
	return iReadU32BE(src, value)
}

func (src *IOStream) ReadS32BE(value *int32) bool {
	panic("not implemented")
	return iReadS32BE(src, value)
}

func (src *IOStream) ReadU64LE(value *uint64) bool {
	panic("not implemented")
	return iReadU64LE(src, value)
}

func (src *IOStream) ReadS64LE(value *int64) bool {
	panic("not implemented")
	return iReadS64LE(src, value)
}

func (src *IOStream) ReadU64BE(value *uint64) bool {
	panic("not implemented")
	return iReadU64BE(src, value)
}

func (src *IOStream) ReadS64BE(value *int64) bool {
	panic("not implemented")
	return iReadS64BE(src, value)
}

func (dst *IOStream) WriteU8(value uint8) bool {
	panic("not implemented")
	return iWriteU8(dst, value)
}

func (dst *IOStream) WriteS8(value int8) bool {
	panic("not implemented")
	return iWriteS8(dst, value)
}

func (dst *IOStream) WriteU16LE(value uint16) bool {
	panic("not implemented")
	return iWriteU16LE(dst, value)
}

func (dst *IOStream) WriteS16LE(value int16) bool {
	panic("not implemented")
	return iWriteS16LE(dst, value)
}

func (dst *IOStream) WriteU16BE(value uint16) bool {
	panic("not implemented")
	return iWriteU16BE(dst, value)
}

func (dst *IOStream) WriteS16BE(value int16) bool {
	panic("not implemented")
	return iWriteS16BE(dst, value)
}

func (dst *IOStream) WriteU32LE(value uint32) bool {
	panic("not implemented")
	return iWriteU32LE(dst, value)
}

func (dst *IOStream) WriteS32LE(value int32) bool {
	panic("not implemented")
	return iWriteS32LE(dst, value)
}

func (dst *IOStream) WriteU32BE(value uint32) bool {
	panic("not implemented")
	return iWriteU32BE(dst, value)
}

func (dst *IOStream) WriteS32BE(value int32) bool {
	panic("not implemented")
	return iWriteS32BE(dst, value)
}

func (dst *IOStream) WriteU64LE(value uint64) bool {
	panic("not implemented")
	return iWriteU64LE(dst, value)
}

func (dst *IOStream) WriteS64LE(value int64) bool {
	panic("not implemented")
	return iWriteS64LE(dst, value)
}

func (dst *IOStream) WriteU64BE(value uint64) bool {
	panic("not implemented")
	return iWriteU64BE(dst, value)
}

func (dst *IOStream) WriteS64BE(value int64) bool {
	panic("not implemented")
	return iWriteS64BE(dst, value)
}

func (src *IOStream) LoadWAV_IO(closeio bool, spec *AudioSpec, audio_buf **uint8, audio_len *uint32) bool {
	panic("not implemented")
	return iLoadWAV_IO(src, closeio, spec, audio_buf, audio_len)
}

func (src *IOStream) LoadBMP_IO(closeio bool) *Surface {
	panic("not implemented")
	return iLoadBMP_IO(src, closeio)
}

func (src *IOStream) AddGamepadMappingsFromIO(closeio bool) int32 {
	panic("not implemented")
	return iAddGamepadMappingsFromIO(src, closeio)
}

// Palette

func (palette *Palette) SetColors(colors *Color, firstcolor int32, ncolors int32) bool {
	panic("not implemented")
	return iSetPaletteColors(palette, colors, firstcolor, ncolors)
}

func (palette *Palette) Destroy() {
	panic("not implemented")
	iDestroyPalette(palette)
}

// Joystick

func (joystick *Joystick) SetVirtualAxis(axis int32, value int16) bool {
	panic("not implemented")
	return iSetJoystickVirtualAxis(joystick, axis, value)
}

func (joystick *Joystick) SetVirtualBall(ball int32, xrel int16, yrel int16) bool {
	panic("not implemented")
	return iSetJoystickVirtualBall(joystick, ball, xrel, yrel)
}

func (joystick *Joystick) SetVirtualButton(button int32, down bool) bool {
	panic("not implemented")
	return iSetJoystickVirtualButton(joystick, button, down)
}

func (joystick *Joystick) SetVirtualHat(hat int32, value uint8) bool {
	panic("not implemented")
	return iSetJoystickVirtualHat(joystick, hat, value)
}

func (joystick *Joystick) SetVirtualTouchpad(touchpad int32, finger int32, down bool, x float32, y float32, pressure float32) bool {
	panic("not implemented")
	return iSetJoystickVirtualTouchpad(joystick, touchpad, finger, down, x, y, pressure)
}

func (joystick *Joystick) SendVirtualSensorData(typ SensorType, sensor_timestamp uint64, data *float32, num_values int32) bool {
	panic("not implemented")
	return iSendJoystickVirtualSensorData(joystick, typ, sensor_timestamp, data, num_values)
}

func (joystick *Joystick) Properties() PropertiesID {
	panic("not implemented")
	return iGetJoystickProperties(joystick)
}

func (joystick *Joystick) Name() string {
	panic("not implemented")
	return iGetJoystickName(joystick)
}

func (joystick *Joystick) Path() string {
	panic("not implemented")
	return iGetJoystickPath(joystick)
}

func (joystick *Joystick) PlayerIndex() int32 {
	panic("not implemented")
	return iGetJoystickPlayerIndex(joystick)
}

func (joystick *Joystick) SetPlayerIndex(player_index int32) bool {
	panic("not implemented")
	return iSetJoystickPlayerIndex(joystick, player_index)
}

func (joystick *Joystick) GUID() GUID {
	panic("not implemented")
	return iGetJoystickGUID(joystick)
}

func (joystick *Joystick) Vendor() uint16 {
	panic("not implemented")
	return iGetJoystickVendor(joystick)
}

func (joystick *Joystick) Product() uint16 {
	panic("not implemented")
	return iGetJoystickProduct(joystick)
}

func (joystick *Joystick) ProductVersion() uint16 {
	panic("not implemented")
	return iGetJoystickProductVersion(joystick)
}

func (joystick *Joystick) FirmwareVersion() uint16 {
	panic("not implemented")
	return iGetJoystickFirmwareVersion(joystick)
}

func (joystick *Joystick) Serial() string {
	panic("not implemented")
	return iGetJoystickSerial(joystick)
}

func (joystick *Joystick) Type() JoystickType {
	panic("not implemented")
	return iGetJoystickType(joystick)
}

func (joystick *Joystick) Connected() bool {
	panic("not implemented")
	return iJoystickConnected(joystick)
}

func (joystick *Joystick) ID() JoystickID {
	panic("not implemented")
	return iGetJoystickID(joystick)
}

func (joystick *Joystick) NumAxes() int32 {
	panic("not implemented")
	return iGetNumJoystickAxes(joystick)
}

func (joystick *Joystick) NumBalls() int32 {
	panic("not implemented")
	return iGetNumJoystickBalls(joystick)
}

func (joystick *Joystick) NumHats() int32 {
	panic("not implemented")
	return iGetNumJoystickHats(joystick)
}

func (joystick *Joystick) NumButtons() int32 {
	panic("not implemented")
	return iGetNumJoystickButtons(joystick)
}

func (joystick *Joystick) Axis(axis int32) int16 {
	panic("not implemented")
	return iGetJoystickAxis(joystick, axis)
}

func (joystick *Joystick) AxisInitialState(axis int32, state *int16) bool {
	panic("not implemented")
	return iGetJoystickAxisInitialState(joystick, axis, state)
}

func (joystick *Joystick) Ball(ball int32, dx *int32, dy *int32) bool {
	panic("not implemented")
	return iGetJoystickBall(joystick, ball, dx, dy)
}

func (joystick *Joystick) Hat(hat int32) uint8 {
	panic("not implemented")
	return iGetJoystickHat(joystick, hat)
}

func (joystick *Joystick) Button(button int32) bool {
	panic("not implemented")
	return iGetJoystickButton(joystick, button)
}

func (joystick *Joystick) Rumble(low_frequency_rumble uint16, high_frequency_rumble uint16, duration_ms uint32) bool {
	panic("not implemented")
	return iRumbleJoystick(joystick, low_frequency_rumble, high_frequency_rumble, duration_ms)
}

func (joystick *Joystick) RumbleTriggers(left_rumble uint16, right_rumble uint16, duration_ms uint32) bool {
	panic("not implemented")
	return iRumbleJoystickTriggers(joystick, left_rumble, right_rumble, duration_ms)
}

func (joystick *Joystick) SetLED(red uint8, green uint8, blue uint8) bool {
	panic("not implemented")
	return iSetJoystickLED(joystick, red, green, blue)
}

func (joystick *Joystick) SendEffect(data *byte, size int32) bool {
	panic("not implemented")
	//return iSendJoystickEffect(joystick, data, size)
}

func (joystick *Joystick) Close() {
	panic("not implemented")
	iCloseJoystick(joystick)
}

func (joystick *Joystick) ConnectionState() JoystickConnectionState {
	panic("not implemented")
	return iGetJoystickConnectionState(joystick)
}

func (joystick *Joystick) PowerInfo(percent *int32) PowerState {
	panic("not implemented")
	return iGetJoystickPowerInfo(joystick, percent)
}

func (joystick *Joystick) IsHaptic() bool {
	panic("not implemented")
	return iIsJoystickHaptic(joystick)
}

func (joystick *Joystick) OpenHapticFrom() *Haptic {
	panic("not implemented")
	return iOpenHapticFromJoystick(joystick)
}

// GamepadType

func (typ GamepadType) GamepadStringForType() string {
	panic("not implemented")
	return iGetGamepadStringForType(typ)
}

func (typ GamepadType) GamepadButtonLabelForType(button GamepadButton) GamepadButtonLabel {
	panic("not implemented")
	return iGetGamepadButtonLabelForType(typ, button)
}

// StorageInterface

func (iface *StorageInterface) OpenStorage(userdata *byte) *Storage {
	panic("not implemented")
	//return iOpenStorage(iface, userdata)
}

// Mutex

func (mutex *Mutex) Lock() {
	panic("not implemented")
	iLockMutex(mutex)
}

func (mutex *Mutex) TryLock() bool {
	panic("not implemented")
	return iTryLockMutex(mutex)
}

func (mutex *Mutex) Unlock() {
	panic("not implemented")
	iUnlockMutex(mutex)
}

func (mutex *Mutex) Destroy() {
	panic("not implemented")
	iDestroyMutex(mutex)
}

// MessageBoxData

func (messageboxdata *MessageBoxData) ShowMessageBox(buttonid *int32) bool {
	panic("not implemented")
	return iShowMessageBox(messageboxdata, buttonid)
}

// MessageBoxFlags

func (flags MessageBoxFlags) ShowSimpleMessageBox(title string, message string, window *Window) bool {
	panic("not implemented")
	return iShowSimpleMessageBox(flags, title, message, window)
}

// Time

func (ticks Time) ToDateTime(dt *DateTime, localTime bool) bool {
	panic("not implemented")
	return iTimeToDateTime(ticks, dt, localTime)
}

func (ticks Time) ToWindows(dwLowDateTime *uint32, dwHighDateTime *uint32) {
	panic("not implemented")
	iTimeToWindows(ticks, dwLowDateTime, dwHighDateTime)
}

// AtomicU32

func (a *AtomicU32) CompareAndSwap(oldval uint32, newval uint32) bool {
	panic("not implemented")
	return iCompareAndSwapAtomicU32(a, oldval, newval)
}

func (a *AtomicU32) Set(v uint32) uint32 {
	panic("not implemented")
	return iSetAtomicU32(a, v)
}

func (a *AtomicU32) Get() uint32 {
	panic("not implemented")
	return iGetAtomicU32(a)
}

// PropertiesID

func (src PropertiesID) CopyProperties(dst PropertiesID) bool {
	panic("not implemented")
	return iCopyProperties(src, dst)
}

func (props PropertiesID) LockProperties() bool {
	panic("not implemented")
	return iLockProperties(props)
}

func (props PropertiesID) UnlockProperties() {
	panic("not implemented")
	iUnlockProperties(props)
}

func (props PropertiesID) SetPointerPropertyWithCleanup(name string, value *byte, cleanup CleanupPropertyCallback, userdata *byte) bool {
	panic("not implemented")
	//return iSetPointerPropertyWithCleanup(props, name, value, cleanup, userdata)
}

func (props PropertiesID) SetPointerProperty(name string, value *byte) bool {
	panic("not implemented")
	//return iSetPointerProperty(props, name, value)
}

func (props PropertiesID) SetStringProperty(name string, value string) bool {
	panic("not implemented")
	return iSetStringProperty(props, name, value)
}

func (props PropertiesID) SetNumberProperty(name string, value int64) bool {
	panic("not implemented")
	return iSetNumberProperty(props, name, value)
}

func (props PropertiesID) SetFloatProperty(name string, value float32) bool {
	panic("not implemented")
	return iSetFloatProperty(props, name, value)
}

func (props PropertiesID) SetBooleanProperty(name string, value bool) bool {
	panic("not implemented")
	return iSetBooleanProperty(props, name, value)
}

func (props PropertiesID) HasProperty(name string) bool {
	panic("not implemented")
	return iHasProperty(props, name)
}

func (props PropertiesID) PropertyType(name string) PropertyType {
	panic("not implemented")
	return iGetPropertyType(props, name)
}

func (props PropertiesID) PointerProperty(name string, default_value *byte) *byte {
	panic("not implemented")
	//return iGetPointerProperty(props, name, default_value)
}

func (props PropertiesID) StringProperty(name string, default_value string) string {
	panic("not implemented")
	return iGetStringProperty(props, name, default_value)
}

func (props PropertiesID) NumberProperty(name string, default_value int64) int64 {
	panic("not implemented")
	return iGetNumberProperty(props, name, default_value)
}

func (props PropertiesID) FloatProperty(name string, default_value float32) float32 {
	panic("not implemented")
	return iGetFloatProperty(props, name, default_value)
}

func (props PropertiesID) BooleanProperty(name string, default_value bool) bool {
	panic("not implemented")
	return iGetBooleanProperty(props, name, default_value)
}

func (props PropertiesID) ClearProperty(name string) bool {
	panic("not implemented")
	return iClearProperty(props, name)
}

func (props PropertiesID) EnumerateProperties(callback EnumeratePropertiesCallback, userdata *byte) bool {
	panic("not implemented")
	//return iEnumerateProperties(props, callback, userdata)
}

func (props PropertiesID) DestroyProperties() {
	panic("not implemented")
	iDestroyProperties(props)
}

func (props PropertiesID) GPUSupportsProperties() bool {
	panic("not implemented")
	return iGPUSupportsProperties(props)
}

func (props PropertiesID) CreateGPUDeviceWithProperties() *GPUDevice {
	panic("not implemented")
	return iCreateGPUDeviceWithProperties(props)
}

func (props PropertiesID) CreateProcessWithProperties() *Process {
	panic("not implemented")
	return iCreateProcessWithProperties(props)
}

// Semaphore

func (sem *Semaphore) Destroy() {
	panic("not implemented")
	iDestroySemaphore(sem)
}

func (sem *Semaphore) Wait() {
	panic("not implemented")
	iWaitSemaphore(sem)
}

func (sem *Semaphore) TryWait() bool {
	panic("not implemented")
	return iTryWaitSemaphore(sem)
}

func (sem *Semaphore) WaitTimeout(timeoutMS int32) bool {
	panic("not implemented")
	return iWaitSemaphoreTimeout(sem, timeoutMS)
}

func (sem *Semaphore) Signal() {
	panic("not implemented")
	iSignalSemaphore(sem)
}

func (sem *Semaphore) Value() uint32 {
	panic("not implemented")
	return iGetSemaphoreValue(sem)
}

// PixelFormat

func (format PixelFormat) Name() string {
	panic("not implemented")
	return iGetPixelFormatName(format)
}

func (format PixelFormat) Masks(bpp *int32, Rmask *uint32, Gmask *uint32, Bmask *uint32, Amask *uint32) bool {
	// TODO: make it return (color.RGBA, error) or something
	panic("not implemented")
	return iGetMasksForPixelFormat(format, bpp, Rmask, Gmask, Bmask, Amask)
}

func (format PixelFormat) Details() (*PixelFormatDetails, error) {
	details := iGetPixelFormatDetails(format)
	if details == nil {
		return nil, internal.LastErr()
	}

	return details, nil
}

// EGLAttribArrayCallback

func (platformAttribCallback EGLAttribArrayCallback) EGL_SetAttributeCallbacks(surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback, userdata *byte) {
	panic("not implemented")
	//iEGL_SetAttributeCallbacks(platformAttribCallback, surfaceAttribCallback, contextAttribCallback, userdata)
}

// DialogFileCallback

func (callback DialogFileCallback) ShowOpenFileDialog(userdata *byte, window *Window, filters *DialogFileFilter, nfilters int32, default_location string, allow_many bool) {
	panic("not implemented")
	//iShowOpenFileDialog(callback, userdata, window, filters, nfilters, default_location, allow_many)
}

func (callback DialogFileCallback) ShowSaveFileDialog(userdata *byte, window *Window, filters *DialogFileFilter, nfilters int32, default_location string) {
	panic("not implemented")
	//iShowSaveFileDialog(callback, userdata, window, filters, nfilters, default_location)
}

func (callback DialogFileCallback) ShowOpenFolderDialog(userdata *byte, window *Window, default_location string, allow_many bool) {
	panic("not implemented")
	//iShowOpenFolderDialog(callback, userdata, window, default_location, allow_many)
}

// LogOutputFunction

func (callback *LogOutputFunction) Get(userdata **byte) {
	panic("not implemented")
	//iGetLogOutputFunction(callback, userdata)
}

// DateTime

func (dt *DateTime) ToTime(ticks *Time) bool {
	panic("not implemented")
	return iDateTimeToTime(dt, ticks)
}

// Keycode

func (key Keycode) ScancodeFromKey(modstate *Keymod) Scancode {
	panic("not implemented")
	return iGetScancodeFromKey(key, modstate)
}

func (key Keycode) KeyName() string {
	panic("not implemented")
	return iGetKeyName(key)
}

// EventFilter

func (filter EventFilter) Set(userdata *byte) {
	panic("not implemented")
	//iSetEventFilter(filter, userdata)
}

func (filter EventFilter) AddEventWatch(userdata *byte) bool {
	panic("not implemented")
	//return iAddEventWatch(filter, userdata)
}

func (filter EventFilter) RemoveEventWatch(userdata *byte) {
	panic("not implemented")
	//iRemoveEventWatch(filter, userdata)
}

func (filter EventFilter) FilterEvents(userdata *byte) {
	panic("not implemented")
	//iFilterEvents(filter, userdata)
}

// Folder

func (folder Folder) User() string {
	panic("not implemented")
	return iGetUserFolder(folder)
}

// GPUCommandBuffer

func (command_buffer *GPUCommandBuffer) InsertGPUDebugLabel(text string) {
	panic("not implemented")
	iInsertGPUDebugLabel(command_buffer, text)
}

func (command_buffer *GPUCommandBuffer) PushGPUDebugGroup(name string) {
	panic("not implemented")
	iPushGPUDebugGroup(command_buffer, name)
}

func (command_buffer *GPUCommandBuffer) PopGPUDebugGroup() {
	panic("not implemented")
	iPopGPUDebugGroup(command_buffer)
}

func (command_buffer *GPUCommandBuffer) PushGPUVertexUniformData(slot_index uint32, data *byte, length uint32) {
	panic("not implemented")
	//iPushGPUVertexUniformData(command_buffer, slot_index, data, length)
}

func (command_buffer *GPUCommandBuffer) PushGPUFragmentUniformData(slot_index uint32, data *byte, length uint32) {
	panic("not implemented")
	//iPushGPUFragmentUniformData(command_buffer, slot_index, data, length)
}

func (command_buffer *GPUCommandBuffer) PushGPUComputeUniformData(slot_index uint32, data *byte, length uint32) {
	panic("not implemented")
	//iPushGPUComputeUniformData(command_buffer, slot_index, data, length)
}

func (command_buffer *GPUCommandBuffer) BeginGPURenderPass(color_target_infos *GPUColorTargetInfo, num_color_targets uint32, depth_stencil_target_info *GPUDepthStencilTargetInfo) *GPURenderPass {
	panic("not implemented")
	//return iBeginGPURenderPass(command_buffer, color_target_infos, num_color_targets, depth_stencil_target_info)
}

func (command_buffer *GPUCommandBuffer) BeginGPUComputePass(storage_texture_bindings *GPUStorageTextureReadWriteBinding, num_storage_texture_bindings uint32, storage_buffer_bindings *GPUStorageBufferReadWriteBinding, num_storage_buffer_bindings uint32) *GPUComputePass {
	panic("not implemented")
	return iBeginGPUComputePass(command_buffer, storage_texture_bindings, num_storage_texture_bindings, storage_buffer_bindings, num_storage_buffer_bindings)
}

func (command_buffer *GPUCommandBuffer) BeginGPUCopyPass() *GPUCopyPass {
	panic("not implemented")
	return iBeginGPUCopyPass(command_buffer)
}

func (command_buffer *GPUCommandBuffer) GenerateMipmapsForGPUTexture(texture *GPUTexture) {
	panic("not implemented")
	iGenerateMipmapsForGPUTexture(command_buffer, texture)
}

func (command_buffer *GPUCommandBuffer) BlitGPUTexture(info *GPUBlitInfo) {
	panic("not implemented")
	iBlitGPUTexture(command_buffer, info)
}

func (command_buffer *GPUCommandBuffer) AcquireGPUSwapchainTexture(window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
	panic("not implemented")
	return iAcquireGPUSwapchainTexture(command_buffer, window, swapchain_texture, swapchain_texture_width, swapchain_texture_height)
}

func (command_buffer *GPUCommandBuffer) WaitAndAcquireGPUSwapchainTexture(window *Window, swapchain_texture **GPUTexture, swapchain_texture_width *uint32, swapchain_texture_height *uint32) bool {
	panic("not implemented")
	return iWaitAndAcquireGPUSwapchainTexture(command_buffer, window, swapchain_texture, swapchain_texture_width, swapchain_texture_height)
}

func (command_buffer *GPUCommandBuffer) Submit() bool {
	panic("not implemented")
	return iSubmitGPUCommandBuffer(command_buffer)
}

func (command_buffer *GPUCommandBuffer) SubmitAndAcquireFence() *GPUFence {
	panic("not implemented")
	return iSubmitGPUCommandBufferAndAcquireFence(command_buffer)
}

func (command_buffer *GPUCommandBuffer) Cancel() bool {
	panic("not implemented")
	return iCancelGPUCommandBuffer(command_buffer)
}

// Process

func (process *Process) Properties() PropertiesID {
	panic("not implemented")
	return iGetProcessProperties(process)
}

func (process *Process) Read(datasize *uintptr, exitcode *int32) *byte {
	panic("not implemented")
	//return iReadProcess(process, datasize, exitcode)
}

func (process *Process) Input() *IOStream {
	panic("not implemented")
	return iGetProcessInput(process)
}

func (process *Process) Output() *IOStream {
	panic("not implemented")
	return iGetProcessOutput(process)
}

func (process *Process) Kill(force bool) bool {
	panic("not implemented")
	return iKillProcess(process, force)
}

func (process *Process) Wait(block bool, exitcode *int32) bool {
	panic("not implemented")
	return iWaitProcess(process, block, exitcode)
}

func (process *Process) Destroy() {
	panic("not implemented")
	iDestroyProcess(process)
}

// FPoint

func (points *FPoint) RectEnclosingPointsFloat(count int32, clip *FRect, result *FRect) bool {
	panic("not implemented")
	return iGetRectEnclosingPointsFloat(points, count, clip, result)
}

// SensorID

func (instance_id SensorID) SensorNameForID() string {
	panic("not implemented")
	return iGetSensorNameForID(instance_id)
}

func (instance_id SensorID) SensorTypeForID() SensorType {
	panic("not implemented")
	return iGetSensorTypeForID(instance_id)
}

func (instance_id SensorID) SensorNonPortableTypeForID() int32 {
	panic("not implemented")
	return iGetSensorNonPortableTypeForID(instance_id)
}

func (instance_id SensorID) OpenSensor() *Sensor {
	panic("not implemented")
	return iOpenSensor(instance_id)
}

func (instance_id SensorID) SensorFromID() *Sensor {
	panic("not implemented")
	return iGetSensorFromID(instance_id)
}

// SystemCursor

func (id SystemCursor) Create() *Cursor {
	panic("not implemented")
	return iCreateSystemCursor(id)
}

// LogPriority

func (priority LogPriority) SetLogPriorities() {
	panic("not implemented")
	iSetLogPriorities(priority)
}

func (priority LogPriority) SetPrefix(prefix string) bool {
	panic("not implemented")
	return iSetLogPriorityPrefix(priority, prefix)
}
