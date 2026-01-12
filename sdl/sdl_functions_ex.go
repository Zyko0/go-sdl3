package sdl

var (
	//puregogen:library path:windows=sdl.dll path:unix=sdl.so alias=sdl
	//puregogen:function symbol=SDL_ShowMessageBox
	iShowMessageBox func(data *messageBoxData, buttonid *int32) bool

	//puregogen:function symbol=SDL_CreateGPUShader
	iCreateGPUShader func(device *GPUDevice, createinfo *gpuShaderCreateInfo) *GPUShader

	//puregogen:function symbol=SDL_CreateGPUComputePipeline
	iCreateGPUComputePipeline func(device *GPUDevice, createinfo *gpuComputePipelineCreateInfo) *GPUComputePipeline

	//puregogen:function symbol=SDL_CreateGPUGraphicsPipeline
	iCreateGPUGraphicsPipeline func(device *GPUDevice, createinfo *gpuGraphicsPipelineCreateInfo) *GPUGraphicsPipeline

	//puregogen:function symbol=SDL_SetClipboardData
	iSetClipboardData func(callback ClipboardDataCallback, cleanup ClipboardCleanupCallback, userdata uintptr, mime_types **byte, num_mime_types uintptr) bool

	//puregogen:function symbol=SDL_ShowOpenFileDialog
	iShowOpenFileDialog func(callback DialogFileCallback, userdata uintptr, window *Window, filters *dialogFileFilter, num_filters int32, default_location *byte, allow_many bool)

	//puregogen:function symbol=SDL_ShowSaveFileDialog
	iShowSaveFileDialog func(callback DialogFileCallback, userdata uintptr, window *Window, filters *dialogFileFilter, num_filters int32, default_location *byte)

	//puregogen:function symbol=SDL_AttachVirtualJoystick
	iAttachVirtualJoystick func(desc *virtualJoystickDesc) JoystickID

	//puregogen:function symbol=SDL_CreateGPURenderState
	iCreateGPURenderState func(renderer *Renderer, info *gpuRenderStateCreateInfo) *GPURenderState
)
