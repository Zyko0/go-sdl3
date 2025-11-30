package main

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/sdl/sdlgpu"
)

type Blit2DArray struct {
	pipeline           *sdl.GPUGraphicsPipeline
	vertexBuffer       *sdl.GPUBuffer
	indexBuffer        *sdl.GPUBuffer
	sourceTexture      *sdl.GPUTexture
	destinationTexture *sdl.GPUTexture
	sampler            *sdl.GPUSampler
}

var Blit2DArrayExample = &Blit2DArray{}

func (e *Blit2DArray) String() string {
	return "Blit2DArray"
}

func (e *Blit2DArray) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	type Layout struct {
		sdlgpu.BaseLayout
		Vertices []common.PositionTextureVertex `gpu:"len:8"`
		Indices  []uint16                       `gpu:"len:6"`
	}
	layout := &Layout{}
	sdlgpu.RegisterLayout(layout)

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

	// load the images

	image1, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	image2, err := common.LoadBMP("ravioli_inverted.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	imageW := image1.W
	imageH := image1.H

	// create the gpu resources

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  layout.Size(0),
	})
	if err != nil {
		return errors.New("failed to create vertex buffer: " + err.Error())
	}

	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  layout.Size(1),
	})
	if err != nil {
		return errors.New("failed to create index buffer: " + err.Error())
	}

	e.sourceTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Type:              sdl.GPU_TEXTURETYPE_2D_ARRAY,
		Width:             uint32(imageW),
		Height:            uint32(imageH),
		LayerCountOrDepth: 2,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	e.destinationTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Type:              sdl.GPU_TEXTURETYPE_2D_ARRAY,
		Width:             uint32(imageW / 2),
		Height:            uint32(imageH / 2),
		LayerCountOrDepth: 2,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER | sdl.GPU_TEXTUREUSAGE_COLOR_TARGET,
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
			Size:  layout.TotalSize(),
		},
	)
	if err != nil {
		return errors.New("failed to create buffer transfer buffer: " + err.Error())
	}

	bufferTransferPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map buffer transfer buffer: " + err.Error())
	}

	layout.Map(bufferTransferPtr)

	layout.Vertices[0] = common.NewPosTexVert(-1, 1, 0, 0, 0)
	layout.Vertices[1] = common.NewPosTexVert(0, 1, 0, 1, 0)
	layout.Vertices[2] = common.NewPosTexVert(0, -1, 0, 1, 1)
	layout.Vertices[3] = common.NewPosTexVert(-1, -1, 0, 0, 1)
	layout.Vertices[4] = common.NewPosTexVert(0, 1, 0, 0, 0)
	layout.Vertices[5] = common.NewPosTexVert(1, 1, 0, 1, 0)
	layout.Vertices[6] = common.NewPosTexVert(1, -1, 0, 1, 1)
	layout.Vertices[7] = common.NewPosTexVert(0, -1, 0, 0, 1)

	layout.Indices[0] = 0
	layout.Indices[1] = 1
	layout.Indices[2] = 2
	layout.Indices[3] = 0
	layout.Indices[4] = 2
	layout.Indices[5] = 3

	context.Device.UnmapTransferBuffer(bufferTransferBuffer)

	// set up texture data

	imageSizeInBytes := uint32(imageW * imageH * 4)

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
	copy(textureData1, image1.Data)

	textureData2 := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr+uintptr(imageSizeInBytes))),
		imageSizeInBytes,
	)
	copy(textureData2, image2.Data)

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
		Size:   layout.Size(0),
	}, false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         layout.Offset(1, 0),
	}, &sdl.GPUBufferRegion{
		Buffer: e.indexBuffer,
		Offset: 0,
		Size:   layout.Size(1),
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         0,
	}, &sdl.GPUTextureRegion{
		Texture: e.sourceTexture,
		Layer:   0,
		W:       uint32(imageW),
		H:       uint32(imageH),
		D:       1,
	}, false)

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         imageSizeInBytes,
	}, &sdl.GPUTextureRegion{
		Texture: e.sourceTexture,
		Layer:   1,
		W:       uint32(imageW),
		H:       uint32(imageH),
		D:       1,
	}, false)

	copyPass.End()

	uploadCmdBuf.BlitGPUTexture(&sdl.GPUBlitInfo{
		Source: sdl.GPUBlitRegion{
			Texture: e.sourceTexture,
			W:       uint32(imageW),
			H:       uint32(imageH),
		},
		Destination: sdl.GPUBlitRegion{
			Texture: e.destinationTexture,
			W:       uint32(imageW / 2),
			H:       uint32(imageH / 2),
		},
		LoadOp: sdl.GPU_LOADOP_DONT_CARE,
		Filter: sdl.GPU_FILTER_LINEAR,
	})

	uploadCmdBuf.BlitGPUTexture(&sdl.GPUBlitInfo{
		Source: sdl.GPUBlitRegion{
			Texture:           e.sourceTexture,
			LayerOrDepthPlane: 1,
			W:                 uint32(imageW),
			H:                 uint32(imageH),
		},
		Destination: sdl.GPUBlitRegion{
			Texture:           e.destinationTexture,
			LayerOrDepthPlane: 1,
			W:                 uint32(imageW / 2),
			H:                 uint32(imageH / 2),
		},
		LoadOp: sdl.GPU_LOADOP_LOAD,
		Filter: sdl.GPU_FILTER_LINEAR,
	})

	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	return nil
}

func (e *Blit2DArray) Update(context *common.Context) error {
	return nil
}

func (e *Blit2DArray) Draw(context *common.Context) error {
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
				Texture: e.sourceTexture, Sampler: e.sampler,
			},
		})
		renderPass.DrawIndexedPrimitives(6, 1, 0, 0, 0)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{
				Texture: e.destinationTexture, Sampler: e.sampler,
			},
		})
		renderPass.DrawIndexedPrimitives(6, 1, 0, 4, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *Blit2DArray) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)
	context.Device.ReleaseTexture(e.sourceTexture)
	context.Device.ReleaseTexture(e.destinationTexture)
	context.Device.ReleaseSampler(e.sampler)

	context.Quit()
}
