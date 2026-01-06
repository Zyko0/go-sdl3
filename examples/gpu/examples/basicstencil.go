package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type BasicStencil struct {
	maskerPipeline      *sdl.GPUGraphicsPipeline
	maskeePipeline      *sdl.GPUGraphicsPipeline
	vertexBuffer        *sdl.GPUBuffer
	depthStencilTexture *sdl.GPUTexture
}

var BasicStencilExample = &BasicStencil{}

func (e *BasicStencil) String() string {
	return "BasicStencil"
}

func (e *BasicStencil) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	vertexShader, err := common.LoadShader(
		context.Device, "PositionColor.vert", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "SolidColor.frag", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	var depthStencilFormat sdl.GPUTextureFormat

	if context.Device.TextureSupportsFormat(
		sdl.GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT, sdl.GPU_TEXTURETYPE_2D,
		sdl.GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
	) {
		depthStencilFormat = sdl.GPU_TEXTUREFORMAT_D24_UNORM_S8_UINT
	} else if context.Device.TextureSupportsFormat(
		sdl.GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT, sdl.GPU_TEXTURETYPE_2D,
		sdl.GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
	) {
		depthStencilFormat = sdl.GPU_TEXTUREFORMAT_D32_FLOAT_S8_UINT
	} else {
		return errors.New("stencil formats not supported")
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
			Pitch:            uint32(unsafe.Sizeof(common.PositionColorVertex{})),
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
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM,
			Location:   1,
			Offset:     uint32(unsafe.Sizeof(float32(0)) * 3),
		},
	}

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			ColorTargetDescriptions: colorTargetDescriptions,
			HasDepthStencilTarget:   true,
			DepthStencilFormat:      depthStencilFormat,
		},
		DepthStencilState: sdl.GPUDepthStencilState{
			EnableStencilTest: true,
			FrontStencilState: sdl.GPUStencilOpState{
				CompareOp:   sdl.GPU_COMPAREOP_NEVER,
				FailOp:      sdl.GPU_STENCILOP_REPLACE,
				PassOp:      sdl.GPU_STENCILOP_KEEP,
				DepthFailOp: sdl.GPU_STENCILOP_KEEP,
			},
			BackStencilState: sdl.GPUStencilOpState{
				CompareOp:   sdl.GPU_COMPAREOP_NEVER,
				FailOp:      sdl.GPU_STENCILOP_REPLACE,
				PassOp:      sdl.GPU_STENCILOP_KEEP,
				DepthFailOp: sdl.GPU_STENCILOP_KEEP,
			},
			WriteMask: 0xFF,
		},
		RasterizerState: sdl.GPURasterizerState{
			CullMode:  sdl.GPU_CULLMODE_NONE,
			FillMode:  sdl.GPU_FILLMODE_FILL,
			FrontFace: sdl.GPU_FRONTFACE_COUNTER_CLOCKWISE,
		},
		VertexInputState: sdl.GPUVertexInputState{
			VertexBufferDescriptions: vertexBufferDescriptions,
			VertexAttributes:         vertexAttributes,
		},
		PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
	}

	e.maskerPipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create masker pipeline: " + err.Error())
	}

	pipelineCreateInfo.DepthStencilState = sdl.GPUDepthStencilState{
		EnableStencilTest: true,
		FrontStencilState: sdl.GPUStencilOpState{
			CompareOp:   sdl.GPU_COMPAREOP_EQUAL,
			FailOp:      sdl.GPU_STENCILOP_KEEP,
			PassOp:      sdl.GPU_STENCILOP_KEEP,
			DepthFailOp: sdl.GPU_STENCILOP_KEEP,
		},
		BackStencilState: sdl.GPUStencilOpState{
			CompareOp:   sdl.GPU_COMPAREOP_NEVER,
			FailOp:      sdl.GPU_STENCILOP_KEEP,
			PassOp:      sdl.GPU_STENCILOP_KEEP,
			DepthFailOp: sdl.GPU_STENCILOP_KEEP,
		},
		CompareMask: 0xFF,
		WriteMask:   0,
	}

	e.maskeePipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create maskee pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 6),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	w, h, err := context.Window.Size()
	if err != nil {
		return errors.New("failed to get window size: " + err.Error())
	}

	e.depthStencilTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Width:             uint32(w),
		Height:            uint32(h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		SampleCount:       sdl.GPU_SAMPLECOUNT_1,
		Format:            depthStencilFormat,
		Usage:             sdl.GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
	})
	if err != nil {
		return errors.New("failed to create depth stencil texture: " + err.Error())
	}

	transferBuffer, err := context.Device.CreateTransferBuffer(&sdl.GPUTransferBufferCreateInfo{
		Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 6),
	})
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	transferDataPtr, err := context.Device.MapTransferBuffer(transferBuffer, false)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionColorVertex)(unsafe.Pointer(transferDataPtr)), 6,
	)

	vertexData[0] = common.NewPosColorVert(-0.5, -0.5, 0, 255, 255, 0, 255)
	vertexData[1] = common.NewPosColorVert(0.5, -0.5, 0, 255, 255, 0, 255)
	vertexData[2] = common.NewPosColorVert(0, 0.5, 0, 255, 255, 0, 255)
	vertexData[3] = common.NewPosColorVert(-1, -1, 0, 255, 0, 0, 255)
	vertexData[4] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[5] = common.NewPosColorVert(0, 1, 0, 0, 0, 255, 255)

	context.Device.UnmapTransferBuffer(transferBuffer)

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: e.vertexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 6),
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer)

	return nil
}

func (e *BasicStencil) Update(context *common.Context) error {
	return nil
}

func (e *BasicStencil) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		colorTargetInfo := sdl.GPUColorTargetInfo{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}

		depthStencilTargetInfo := sdl.GPUDepthStencilTargetInfo{
			Texture:        e.depthStencilTexture,
			Cycle:          true,
			ClearDepth:     0,
			ClearStencil:   0,
			LoadOp:         sdl.GPU_LOADOP_CLEAR,
			StoreOp:        sdl.GPU_STOREOP_DONT_CARE,
			StencilLoadOp:  sdl.GPU_LOADOP_CLEAR,
			StencilStoreOp: sdl.GPU_STOREOP_DONT_CARE,
		}

		renderPass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{colorTargetInfo}, &depthStencilTargetInfo,
		)

		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{Buffer: e.vertexBuffer, Offset: 0},
		})

		renderPass.SetStencilReference(1)
		renderPass.BindGraphicsPipeline(e.maskerPipeline)
		renderPass.DrawPrimitives(3, 1, 0, 0)

		renderPass.SetStencilReference(0)
		renderPass.BindGraphicsPipeline(e.maskeePipeline)
		renderPass.DrawPrimitives(3, 1, 3, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *BasicStencil) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.maskerPipeline)
	context.Device.ReleaseGraphicsPipeline(e.maskeePipeline)

	context.Device.ReleaseTexture(e.depthStencilTexture)
	context.Device.ReleaseBuffer(e.vertexBuffer)

	context.Quit()
}
