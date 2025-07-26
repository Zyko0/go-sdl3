package main

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Texture2DArray struct {
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
	indexBuffer  *sdl.GPUBuffer
	texture      *sdl.GPUTexture
	sampler      *sdl.GPUSampler
}

var Texture2DArrayExample = &Texture2DArray{}

func (e *Texture2DArray) String() string {
	return "Texture2DArray"
}

func (e *Texture2DArray) Init(context *common.Context) error {
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
		context.Device, "TexturedQuadArray.frag", 1, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	// create the pipeline

	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		sdl.GPUColorTargetDescription{
			Format: context.Device.SwapchainTextureFormat(context.Window),
		},
	}

	vertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
		sdl.GPUVertexBufferDescription{
			Slot:             0,
			InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
			InstanceStepRate: 0,
			Pitch:            uint32(unsafe.Sizeof(common.PositionTextureVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		sdl.GPUVertexAttribute{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT3,
			Location:   0,
			Offset:     0,
		},
		sdl.GPUVertexAttribute{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT2,
			Location:   1,
			Offset:     uint32(unsafe.Sizeof(float32(0)) * 3),
		},
	}

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			NumColorTargets:         uint32(len(colorTargetDescriptions)),
			ColorTargetDescriptions: &colorTargetDescriptions[0],
		},
		VertexInputState: sdl.GPUVertexInputState{
			NumVertexBuffers:         uint32(len(vertexBufferDescriptions)),
			VertexBufferDescriptions: &vertexBufferDescriptions[0],
			NumVertexAttributes:      uint32(len(vertexAttributes)),
			VertexAttributes:         &vertexAttributes[0],
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

	// load the images

	imageData1, imageWidth, imageHeight, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	imageData2, _, _, err := common.LoadBMP("ravioli_inverted.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	// create the gpu resources

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	})
	if err != nil {
		return errors.New("failed to create vertex buffer: " + err.Error())
	}

	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(unsafe.Sizeof(uint16(0)) * 6),
	})
	if err != nil {
		return errors.New("failed to create index buffer: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D_ARRAY,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imageWidth),
		Height:            uint32(imageHeight),
		LayerCountOrDepth: 2,
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

	// set up buffer data

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
		return errors.New("failed to create buffer transfer buffer: " + err.Error())
	}

	bufferTransferPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map buffer transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionTextureVertex)(unsafe.Pointer(bufferTransferPtr)), 4,
	)

	vertexData[0] = common.NewPosTexVert(-1, 1, 0, 0, 0)
	vertexData[1] = common.NewPosTexVert(1, 1, 0, 1, 0)
	vertexData[2] = common.NewPosTexVert(1, -1, 0, 1, 1)
	vertexData[3] = common.NewPosTexVert(-1, -1, 0, 0, 1)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			bufferTransferPtr+unsafe.Sizeof(common.PositionTextureVertex{})*4,
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

	imageSizeInBytes := uint32(imageWidth * imageHeight * 4)

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  imageSizeInBytes * 2,
		},
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureTransferPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	textureData1 := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr)),
		imageSizeInBytes,
	)
	copy(textureData1, imageData1)

	textureData2 := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr+uintptr(imageSizeInBytes))),
		imageSizeInBytes,
	)
	copy(textureData2, imageData2)

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
		Buffer: e.vertexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	}, false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
	}, &sdl.GPUBufferRegion{
		Buffer: e.indexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(uint16(0)) * 6),
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         0,
	}, &sdl.GPUTextureRegion{
		Texture: e.texture,
		W:       uint32(imageWidth),
		H:       uint32(imageHeight),
		D:       1,
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         imageSizeInBytes,
	}, &sdl.GPUTextureRegion{
		Texture: e.texture,
		Layer:   1,
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

func (e *Texture2DArray) Update(context *common.Context) error {
	return nil
}

func (e *Texture2DArray) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		colorTargetInfos := []sdl.GPUColorTargetInfo{sdl.GPUColorTargetInfo{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}}

		renderPass := cmdbuf.BeginRenderPass(colorTargetInfos, nil)

		renderPass.BindGraphicsPipeline(e.pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: e.vertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.indexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{
				Texture: e.texture, Sampler: e.sampler,
			},
		})
		renderPass.DrawIndexedPrimitives(6, 1, 0, 0, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *Texture2DArray) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseSampler(e.sampler)

	context.Quit()
}
