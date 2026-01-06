package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type CopyConsistency struct {
	pipeline *sdl.GPUGraphicsPipeline

	vertexBuffer      *sdl.GPUBuffer
	leftVertexBuffer  *sdl.GPUBuffer
	rightVertexBuffer *sdl.GPUBuffer
	indexBuffer       *sdl.GPUBuffer

	texture      *sdl.GPUTexture
	leftTexture  *sdl.GPUTexture
	rightTexture *sdl.GPUTexture

	sampler *sdl.GPUSampler
}

var CopyConsistencyExample = &CopyConsistency{}

func (e *CopyConsistency) String() string {
	return "CopyConsistency"
}

func (e *CopyConsistency) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create the shaders

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

	// create the pipeline

	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		{
			Format: context.Device.SwapchainTextureFormat(context.Window),
			BlendState: sdl.GPUColorTargetBlendState{
				EnableBlend:         true,
				AlphaBlendOp:        sdl.GPU_BLENDOP_ADD,
				ColorBlendOp:        sdl.GPU_BLENDOP_ADD,
				SrcColorBlendfactor: sdl.GPU_BLENDFACTOR_SRC_ALPHA,
				SrcAlphaBlendfactor: sdl.GPU_BLENDFACTOR_SRC_ALPHA,
				DstColorBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
				DstAlphaBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
			},
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

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
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

	e.pipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// create the textures

	textureCreateInfo := &sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             16,
		Height:            16,
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	}

	e.leftTexture, err = context.Device.CreateTexture(textureCreateInfo)
	if err != nil {
		return errors.New("failed to create left texture: " + err.Error())
	}

	e.rightTexture, err = context.Device.CreateTexture(textureCreateInfo)
	if err != nil {
		return errors.New("failed to create right texture: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(textureCreateInfo)
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	// load the texture data

	leftImage, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	rightImage, err := common.LoadBMP("ravioli_inverted.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	imageWidth := leftImage.W
	imageHeight := leftImage.H

	// create the sampler

	e.sampler, err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_NEAREST,
		MagFilter:    sdl.GPU_FILTER_NEAREST,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_NEAREST,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}

	// create the buffers

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	})
	if err != nil {
		return errors.New("failed to create vertex buffer: " + err.Error())
	}

	e.leftVertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	})
	if err != nil {
		return errors.New("failed to create left vertex buffer: " + err.Error())
	}

	e.rightVertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	})
	if err != nil {
		return errors.New("failed to create right vertex buffer: " + err.Error())
	}

	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(unsafe.Sizeof(uint16(0)) * 6),
	})
	if err != nil {
		return errors.New("failed to create index buffer: " + err.Error())
	}

	// set up buffer data

	bufferTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size: uint32(
				unsafe.Sizeof(common.PositionTextureVertex{})*8 +
					unsafe.Sizeof(uint16(0))*6,
			),
		},
	)
	if err != nil {
		return errors.New("failed to create buffer transfer buffer: " + err.Error())
	}

	bufferTransferPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map buffer transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionTextureVertex)(unsafe.Pointer(bufferTransferPtr)), 8,
	)

	vertexData[0] = common.NewPosTexVert(-1.0, 1.0, 0, 0, 0)
	vertexData[1] = common.NewPosTexVert(0.0, 1.0, 0, 1, 0)
	vertexData[2] = common.NewPosTexVert(0.0, -1.0, 0, 1, 1)
	vertexData[3] = common.NewPosTexVert(-1.0, -1.0, 0, 0, 1)

	vertexData[4] = common.NewPosTexVert(0.0, 1.0, 0, 0, 0)
	vertexData[5] = common.NewPosTexVert(1.0, 1.0, 0, 1, 0)
	vertexData[6] = common.NewPosTexVert(1.0, -1.0, 0, 1, 1)
	vertexData[7] = common.NewPosTexVert(0.0, -1.0, 0, 0, 1)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			bufferTransferPtr+unsafe.Sizeof(common.PositionTextureVertex{})*8,
		)), 6,
	)

	indexData[0] = 0
	indexData[1] = 1
	indexData[2] = 2
	indexData[3] = 0
	indexData[4] = 2
	indexData[5] = 3

	context.Device.UnmapTransferBuffer(bufferTransferBuffer)

	// set up texture data

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(imageWidth * imageHeight * 8),
		},
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureTransferPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	leftTextureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr)),
		imageWidth*imageHeight*4,
	)
	copy(leftTextureData, leftImage.Data)

	rightTextureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr+uintptr(imageWidth*imageHeight*4))),
		imageWidth*imageHeight*4,
	)
	copy(rightTextureData, rightImage.Data)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	// upload the transfer data to the gpu resources

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         0,
	}, &sdl.GPUBufferRegion{
		Buffer: e.leftVertexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	}, false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	}, &sdl.GPUBufferRegion{
		Buffer: e.rightVertexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	}, false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 8),
	}, &sdl.GPUBufferRegion{
		Buffer: e.indexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(uint16(0)) * 6),
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         0,
	}, &sdl.GPUTextureRegion{
		Texture: e.leftTexture,
		W:       uint32(imageWidth),
		H:       uint32(imageHeight),
		D:       1,
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         uint32(imageWidth * imageHeight * 4),
	}, &sdl.GPUTextureRegion{
		Texture: e.rightTexture,
		W:       uint32(imageWidth),
		H:       uint32(imageHeight),
		D:       1,
	}, false)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	return nil
}

func (e *CopyConsistency) Update(context *common.Context) error {
	return nil
}

func (e *CopyConsistency) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		colorTargetInfos := []sdl.GPUColorTargetInfo{{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}}

		// copy left-side resources
		copyPass := cmdbuf.BeginCopyPass()
		copyPass.CopyGPUBufferToBuffer(
			&sdl.GPUBufferLocation{
				Buffer: e.leftVertexBuffer,
			}, &sdl.GPUBufferLocation{
				Buffer: e.vertexBuffer,
			},
			uint32(unsafe.Sizeof(common.PositionTextureVertex{})*4),
			false,
		)
		copyPass.CopyGPUTextureToTexture(
			&sdl.GPUTextureLocation{Texture: e.leftTexture},
			&sdl.GPUTextureLocation{Texture: e.texture},
			16, 16, 1, false,
		)
		copyPass.End()

		// draw the left side
		renderPass := cmdbuf.BeginRenderPass(colorTargetInfos, nil)
		renderPass.BindGraphicsPipeline(e.pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{Buffer: e.vertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.indexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			{
				Texture: e.texture, Sampler: e.sampler,
			},
		})
		renderPass.DrawIndexedPrimitives(6, 1, 0, 0, 0)
		renderPass.End()

		// copy right-side resources
		copyPass = cmdbuf.BeginCopyPass()
		copyPass.CopyGPUBufferToBuffer(
			&sdl.GPUBufferLocation{
				Buffer: e.rightVertexBuffer,
			}, &sdl.GPUBufferLocation{
				Buffer: e.vertexBuffer,
			},
			uint32(unsafe.Sizeof(common.PositionTextureVertex{})*4),
			false,
		)
		copyPass.CopyGPUTextureToTexture(
			&sdl.GPUTextureLocation{Texture: e.rightTexture},
			&sdl.GPUTextureLocation{Texture: e.texture},
			16, 16, 1, false,
		)
		copyPass.End()

		// draw the right side
		colorTargetInfos[0].LoadOp = sdl.GPU_LOADOP_LOAD
		renderPass = cmdbuf.BeginRenderPass(colorTargetInfos, nil)
		renderPass.BindGraphicsPipeline(e.pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{Buffer: e.vertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.indexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			{
				Texture: e.texture, Sampler: e.sampler,
			},
		})
		renderPass.DrawIndexedPrimitives(6, 1, 0, 0, 0)
		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *CopyConsistency) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)

	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.leftVertexBuffer)
	context.Device.ReleaseBuffer(e.rightVertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)

	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseTexture(e.leftTexture)
	context.Device.ReleaseTexture(e.rightTexture)

	context.Device.ReleaseSampler(e.sampler)

	context.Quit()
}
