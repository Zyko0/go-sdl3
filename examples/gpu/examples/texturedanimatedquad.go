package main

import (
	"errors"
	"math"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/mathgl/mgl32"
)

type TexturedAnimatedQuad struct {
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
	indexBuffer  *sdl.GPUBuffer
	texture      *sdl.GPUTexture
	sampler      *sdl.GPUSampler

	t float32
}

type FragMultiplyUniform struct {
	r, g, b, a float32
}

var TexturedAnimatedQuadExample = &TexturedAnimatedQuad{}

func (e *TexturedAnimatedQuad) String() string {
	return "TexturedAnimatedQuad"
}

func (e *TexturedAnimatedQuad) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create shaders

	vertexShader, err := common.LoadShader(
		context.Device, "TexturedQuadWithMatrix.vert", 0, 1, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "TexturedQuadWithMultiplyColor.frag", 1, 1, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	// load the image

	image, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
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

	// create gpu resources buffer

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(unsafe.Sizeof(uint16(0)) * 6),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(image.W),
		Height:            uint32(image.H),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

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

	// setup buffer data

	bufferTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size: uint32(
				unsafe.Sizeof(common.PositionTextureVertex{})*4 +
					unsafe.Sizeof(uint16(0))*6,
			),
		},
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	bufferTransferDataPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map buffer transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionTextureVertex)(unsafe.Pointer(bufferTransferDataPtr)), 4,
	)

	vertexData[0] = common.NewPosTexVert(-0.5, -0.5, 0, 0, 0)
	vertexData[1] = common.NewPosTexVert(0.5, -0.5, 0, 1, 0)
	vertexData[2] = common.NewPosTexVert(0.5, 0.5, 0, 1, 1)
	vertexData[3] = common.NewPosTexVert(-0.5, 0.5, 0, 0, 1)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			bufferTransferDataPtr+unsafe.Sizeof(common.PositionTextureVertex{})*4,
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
			Size:  uint32(image.W * image.H * 4),
		},
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	textureTransferDataPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	textureData := unsafe.Slice(
		(*uint8)(unsafe.Pointer(textureTransferDataPtr)),
		image.W*image.H*4,
	)

	copy(textureData, image.Data)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	// upload the transfer data to the gpu resources

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: bufferTransferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: e.vertexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
		},
		false,
	)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: bufferTransferBuffer,
			Offset:         uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
		},
		&sdl.GPUBufferRegion{
			Buffer: e.indexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(uint16(0)) * 6),
		},
		false,
	)

	copyPass.UploadToGPUTexture(
		&sdl.GPUTextureTransferInfo{
			TransferBuffer: textureTransferBuffer,
			Offset:         0,
		},
		&sdl.GPUTextureRegion{
			Texture: e.texture,
			W:       uint32(image.W),
			H:       uint32(image.H),
			D:       1,
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	return nil
}

func (e *TexturedAnimatedQuad) Update(context *common.Context) error {
	e.t += context.DeltaTime
	return nil
}

func (e *TexturedAnimatedQuad) Draw(context *common.Context) error {
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

		renderPass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{colorTargetInfo}, nil,
		)

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

		// bottom left
		matrixUniform := mgl32.Translate3D(-0.5, -0.5, 0).Mul4(
			mgl32.HomogRotate3DZ(e.t),
		)
		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&matrixUniform)), unsafe.Sizeof(matrixUniform),
		))
		fragMultiplyUniform := FragMultiplyUniform{
			r: 1, g: float32(0.5 + math.Sin(float64(e.t))*0.5), b: 1, a: 1,
		}
		cmdbuf.PushFragmentUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&fragMultiplyUniform)), unsafe.Sizeof(fragMultiplyUniform),
		))
		renderPass.DrawIndexedPrimitives(
			6, 1, 0, 0, 0,
		)

		// bottom right
		matrixUniform = mgl32.Translate3D(0.5, -0.5, 0).Mul4(
			mgl32.HomogRotate3DZ((2 * math.Pi) - e.t),
		)
		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&matrixUniform)), unsafe.Sizeof(matrixUniform),
		))
		fragMultiplyUniform = FragMultiplyUniform{
			r: 1, g: float32(0.5 + math.Cos(float64(e.t))*0.5), b: 1, a: 1,
		}
		cmdbuf.PushFragmentUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&fragMultiplyUniform)), unsafe.Sizeof(fragMultiplyUniform),
		))
		renderPass.DrawIndexedPrimitives(
			6, 1, 0, 0, 0,
		)

		// top left
		matrixUniform = mgl32.Translate3D(-0.5, 0.5, 0).Mul4(
			mgl32.HomogRotate3DZ(e.t),
		)
		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&matrixUniform)), unsafe.Sizeof(matrixUniform),
		))
		fragMultiplyUniform = FragMultiplyUniform{
			r: 1, g: float32(0.5 + math.Sin(float64(e.t))*0.2), b: 1, a: 1,
		}
		cmdbuf.PushFragmentUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&fragMultiplyUniform)), unsafe.Sizeof(fragMultiplyUniform),
		))
		renderPass.DrawIndexedPrimitives(
			6, 1, 0, 0, 0,
		)

		// top right
		matrixUniform = mgl32.Translate3D(0.5, 0.5, 0).Mul4(
			mgl32.HomogRotate3DZ(e.t),
		)
		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&matrixUniform)), unsafe.Sizeof(matrixUniform),
		))
		fragMultiplyUniform = FragMultiplyUniform{
			r: 1, g: float32(0.5 + math.Cos(float64(e.t))*1), b: 1, a: 1,
		}
		cmdbuf.PushFragmentUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&fragMultiplyUniform)), unsafe.Sizeof(fragMultiplyUniform),
		))
		renderPass.DrawIndexedPrimitives(
			6, 1, 0, 0, 0,
		)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *TexturedAnimatedQuad) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseSampler(e.sampler)

	e.t = 0

	context.Quit()
}
