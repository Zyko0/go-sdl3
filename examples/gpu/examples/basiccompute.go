package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type BasicCompute struct {
	drawPipeline *sdl.GPUGraphicsPipeline
	texture      *sdl.GPUTexture
	sampler      *sdl.GPUSampler
	vertexBuffer *sdl.GPUBuffer
}

var BasicComputeExample = &BasicCompute{}

func (e *BasicCompute) String() string {
	return "BasicCompute"
}

func (e *BasicCompute) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	vertexShader, err := common.LoadShader(
		context.Device, "TexturedQuad.vert", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "TexturedQuad.frag", 1, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	fillTexturePipeline, err := common.CreateComputePipelineFromShader(
		context.Device, "FillTexture.comp", sdl.GPUComputePipelineCreateInfo{
			NumReadwriteStorageTextures: 1,
			ThreadcountX:                8,
			ThreadcountY:                8,
			ThreadcountZ:                1,
		},
	)
	if err != nil {
		return errors.New("failed to create compute pipeline: " + err.Error())
	}

	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		{
			Format: context.Device.SwapchainTextureFormat(context.Window),
		},
	}

	vertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
		{
			Slot:             0,
			InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
			InstanceStepRate: 0,
			Pitch:            uint32(unsafe.Sizeof(common.PositionTextureVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT3,
			Location:   0,
			Offset:     0,
		},
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT2,
			Location:   1,
			Offset:     uint32(unsafe.Sizeof(float32(0)) * 3),
		},
	}

	drawPipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			ColorTargetDescriptions: colorTargetDescriptions,
		},
		VertexInputState: sdl.GPUVertexInputState{
			VertexBufferDescriptions: vertexBufferDescriptions,
			VertexAttributes:         vertexAttributes,
		},
		PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
	}

	e.drawPipeline, err = context.Device.CreateGraphicsPipeline(&drawPipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create draw pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	w, h, err := context.Window.Size()
	if err != nil {
		return errors.New("failed to get window size: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(w),
		Height:            uint32(h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE | sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	e.sampler, err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 6),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	transferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size: uint32(
				unsafe.Sizeof(common.PositionTextureVertex{}) * 6,
			),
		},
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	transferDataPtr, err := context.Device.MapTransferBuffer(transferBuffer, false)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionTextureVertex)(unsafe.Pointer(transferDataPtr)), 6,
	)

	vertexData[0] = common.NewPosTexVert(-1, -1, 0, 0, 0)
	vertexData[1] = common.NewPosTexVert(1, -1, 0, 1, 0)
	vertexData[2] = common.NewPosTexVert(1, 1, 0, 1, 1)
	vertexData[3] = common.NewPosTexVert(-1, -1, 0, 0, 0)
	vertexData[4] = common.NewPosTexVert(1, 1, 0, 1, 1)
	vertexData[5] = common.NewPosTexVert(-1, 1, 0, 0, 1)

	context.Device.UnmapTransferBuffer(transferBuffer)

	cmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: e.vertexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 6),
		},
		false,
	)

	copyPass.End()

	computePass := cmdBuf.BeginComputePass(
		[]sdl.GPUStorageTextureReadWriteBinding{
			{
				Texture: e.texture,
			},
		},
		[]sdl.GPUStorageBufferReadWriteBinding{},
	)

	computePass.BindGPUComputePipeline(fillTexturePipeline)
	computePass.Dispatch(uint32(w)/8, uint32(h)/8, 1)
	computePass.End()

	cmdBuf.Submit()

	context.Device.ReleaseComputePipeline(fillTexturePipeline)
	context.Device.ReleaseTransferBuffer(transferBuffer)

	return nil
}

func (e *BasicCompute) Update(context *common.Context) error {
	return nil
}

func (e *BasicCompute) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		renderPass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{{
				Texture:    swapchainTexture.Texture,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				StoreOp:    sdl.GPU_STOREOP_STORE,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
				Cycle:      false,
			}}, nil,
		)

		renderPass.BindGraphicsPipeline(e.drawPipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{Buffer: e.vertexBuffer, Offset: 0},
		})
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			{Texture: e.texture, Sampler: e.sampler},
		})
		renderPass.DrawPrimitives(6, 1, 0, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *BasicCompute) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.drawPipeline)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseSampler(e.sampler)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Quit()
}
