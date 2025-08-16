package main

import (
	"errors"
	"math"
	"math/rand"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/sdl/sdlgpu"
	"github.com/go-gl/mathgl/mgl32"
)

// this example is slightly inferior in terms of performance compared to PullSpriteBatch
// despite that, it is included as an example of a compute to graphics workflow.

type ComputeSpriteBatch struct {
	computePipeline             *sdl.GPUComputePipeline
	renderPipeline              *sdl.GPUGraphicsPipeline
	sampler                     *sdl.GPUSampler
	texture                     *sdl.GPUTexture
	spriteComputeTransferBuffer *sdlgpu.TypedTransferBuffer[ComputeSpriteInstance]
	spriteComputeBuffer         *sdlgpu.TypedBuffer[ComputeSpriteInstance]
	spriteVertexBuffer          *sdlgpu.TypedBuffer[common.PositionTextureColorVertex]
	spriteIndexBuffer           *sdlgpu.TypedBuffer[uint32]

	SPRITE_COUNT uint32
}

type ComputeSpriteInstance struct {
	X, Y, Z                float32
	Rotation               float32
	W, H, _, _             float32
	TexU, TexV, TexW, TexH float32
	R, G, B, A             float32
}

var ComputeSpriteBatchExample = &ComputeSpriteBatch{
	SPRITE_COUNT: 8192,
}

func (e *ComputeSpriteBatch) String() string {
	return "ComputeSpriteBatch"
}

func (e *ComputeSpriteBatch) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	presentMode := sdl.GPU_PRESENTMODE_VSYNC
	if context.Device.WindowSupportsPresentMode(
		context.Window, sdl.GPU_PRESENTMODE_IMMEDIATE,
	) {
		presentMode = sdl.GPU_PRESENTMODE_IMMEDIATE
	} else if context.Device.WindowSupportsPresentMode(
		context.Window, sdl.GPU_PRESENTMODE_MAILBOX,
	) {
		presentMode = sdl.GPU_PRESENTMODE_MAILBOX
	}

	context.Device.SetSwapchainParameters(
		context.Window, sdl.GPU_SWAPCHAINCOMPOSITION_SDR, presentMode,
	)

	// create the shaders and pipelines

	vertShader, err := common.LoadShader(
		context.Device, "TexturedQuadColorWithMatrix.vert", 0, 1, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create scene vertex shader: " + err.Error())
	}

	fragShader, err := common.LoadShader(
		context.Device, "TexturedQuadColor.frag", 1, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create scene fragment shader: " + err.Error())
	}

	// create the sprite render pipeline

	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		sdl.GPUColorTargetDescription{
			Format: context.Device.SwapchainTextureFormat(context.Window),
			BlendState: sdl.GPUColorTargetBlendState{
				EnableBlend:         true,
				ColorBlendOp:        sdl.GPU_BLENDOP_ADD,
				AlphaBlendOp:        sdl.GPU_BLENDOP_ADD,
				SrcColorBlendfactor: sdl.GPU_BLENDFACTOR_SRC_ALPHA,
				DstColorBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
				SrcAlphaBlendfactor: sdl.GPU_BLENDFACTOR_SRC_ALPHA,
				DstAlphaBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
			},
		},
	}

	vertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
		sdl.GPUVertexBufferDescription{
			Slot:             0,
			InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
			InstanceStepRate: 0,
			Pitch:            uint32(unsafe.Sizeof(common.PositionTextureColorVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		sdl.GPUVertexAttribute{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT4,
			Location:   0,
			Offset:     0,
		},
		sdl.GPUVertexAttribute{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT2,
			Location:   1,
			Offset:     16,
		},
		sdl.GPUVertexAttribute{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT4,
			Location:   2,
			Offset:     32,
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
		VertexShader:   vertShader,
		FragmentShader: fragShader,
	}

	e.renderPipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create render pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertShader)
	context.Device.ReleaseShader(fragShader)

	// create the sprite batch compute pipeline

	e.computePipeline, err = common.CreateComputePipelineFromShader(
		context.Device, "SpriteBatch.comp",
		sdl.GPUComputePipelineCreateInfo{
			NumReadonlyStorageBuffers:  1,
			NumReadwriteStorageBuffers: 1,
			ThreadcountX:               64,
			ThreadcountY:               1,
			ThreadcountZ:               1,
		},
	)
	if err != nil {
		return errors.New("failed to create compute pipeline: " + err.Error())
	}

	image, err := common.LoadBMP("ravioli_atlas.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	textureTransferBuffer, err := sdlgpu.CreateTypedTransferBuffer[byte](
		context.Device, sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD, uint32(len(image.Data)), 0,
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureData, err := textureTransferBuffer.Map(context.Device, false)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	copy(textureData, image.Data)

	context.Device.UnmapTransferBuffer(textureTransferBuffer.Raw())

	// create the gpu resources

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

	e.spriteComputeTransferBuffer, err = sdlgpu.CreateTypedTransferBuffer[ComputeSpriteInstance](
		context.Device, sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD, e.SPRITE_COUNT, 0,
	)
	if err != nil {
		return errors.New("failed to create sprite compute transfer buffer: " + err.Error())
	}

	e.spriteComputeBuffer, err = sdlgpu.CreateTypedBuffer[ComputeSpriteInstance](
		context.Device, sdl.GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ, e.SPRITE_COUNT, 0,
	)
	if err != nil {
		return errors.New("failed to create sprite compute buffer: " + err.Error())
	}

	e.spriteVertexBuffer, err = sdlgpu.CreateTypedBuffer[common.PositionTextureColorVertex](
		context.Device,
		sdl.GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE|sdl.GPU_BUFFERUSAGE_VERTEX,
		e.SPRITE_COUNT*4, 0,
	)
	if err != nil {
		return errors.New("failed to create sprite vertex buffer: " + err.Error())
	}

	e.spriteIndexBuffer, err = sdlgpu.CreateTypedBuffer[uint32](
		context.Device, sdl.GPU_BUFFERUSAGE_INDEX,
		e.SPRITE_COUNT*6, 0,
	)
	if err != nil {
		return errors.New("failed to create sprite vertex buffer: " + err.Error())
	}

	// transfer the up-front data

	indexBufferTransferBuffer, err := sdlgpu.CreateTypedTransferBuffer[uint32](
		context.Device, sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD, e.SPRITE_COUNT*6, 0,
	)
	if err != nil {
		return errors.New("failed to create index buffer transfer buffer: " + err.Error())
	}

	indexData, err := indexBufferTransferBuffer.Map(context.Device, false)
	if err != nil {
		return errors.New("failed to map index buffer transfer buffer: " + err.Error())
	}

	for i := range e.SPRITE_COUNT {
		indexData[i*6] = i * 4
		indexData[i*6+1] = i*4 + 1
		indexData[i*6+2] = i*4 + 2
		indexData[i*6+3] = i*4 + 3
		indexData[i*6+4] = i*4 + 2
		indexData[i*6+5] = i*4 + 1
	}

	context.Device.UnmapTransferBuffer(indexBufferTransferBuffer.Raw())

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer.Raw(),
		Offset:         0, // zeroes out the rest
	}, &sdl.GPUTextureRegion{
		Texture: e.texture,
		W:       uint32(image.W),
		H:       uint32(image.H),
		D:       1,
	}, false)

	copyPass.UploadToGPUBuffer(
		indexBufferTransferBuffer.Location(0), e.spriteIndexBuffer.Region(0, e.SPRITE_COUNT*6), false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(textureTransferBuffer.Raw())
	context.Device.ReleaseTransferBuffer(indexBufferTransferBuffer.Raw())

	return nil
}

func (e *ComputeSpriteBatch) Update(context *common.Context) error {
	return nil
}

func (e *ComputeSpriteBatch) Draw(context *common.Context) error {
	cameraMatrix := mgl32.Ortho(0, 640, 480, 0, 0, -1)

	cmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdBuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	uCoords := [4]float32{0.0, 0.5, 0.0, 0.5}
	vCoords := [4]float32{0.0, 0.0, 0.5, 0.5}

	if swapchainTexture != nil {
		// build sprite instance transfer
		data, err := e.spriteComputeTransferBuffer.Map(context.Device, true)
		if err != nil {
			return errors.New("failed to map sprite compute transfer buffer: " + err.Error())
		}

		for i := range e.SPRITE_COUNT {
			ravioli := rand.Intn(4)
			data[i].X = float32(rand.Intn(640))
			data[i].Y = float32(rand.Intn(480))
			data[i].Z = 0
			data[i].Rotation = rand.Float32() * math.Pi * 2
			data[i].W = 32
			data[i].H = 32
			data[i].TexU = uCoords[ravioli]
			data[i].TexV = vCoords[ravioli]
			data[i].TexW = 0.5
			data[i].TexH = 0.5
			data[i].R = 1.0
			data[i].G = 1.0
			data[i].B = 1.0
			data[i].A = 1.0
		}

		context.Device.UnmapTransferBuffer(e.spriteComputeTransferBuffer.Raw())

		// upload instance data
		copyPass := cmdBuf.BeginCopyPass()
		copyPass.UploadToGPUBuffer(
			e.spriteComputeTransferBuffer.Location(0),
			e.spriteComputeBuffer.Region(0, e.SPRITE_COUNT),
			true,
		)
		copyPass.End()

		// set up compute pass to build vertex buffer
		computePass := cmdBuf.BeginComputePass(
			[]sdl.GPUStorageTextureReadWriteBinding{},
			[]sdl.GPUStorageBufferReadWriteBinding{
				sdl.GPUStorageBufferReadWriteBinding{
					Buffer: e.spriteVertexBuffer.Raw(),
					Cycle:  true,
				},
			},
		)

		computePass.BindGPUComputePipeline(e.computePipeline)
		computePass.BindStorageBuffers([]*sdl.GPUBuffer{
			e.spriteComputeBuffer.Raw(),
		})
		computePass.Dispatch(e.SPRITE_COUNT/64, 1, 1)

		computePass.End()

		// render sprites
		renderPass := cmdBuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
			sdl.GPUColorTargetInfo{
				Texture:    swapchainTexture.Texture,
				Cycle:      false,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				StoreOp:    sdl.GPU_STOREOP_STORE,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			},
		}, nil)

		renderPass.BindGraphicsPipeline(e.renderPipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			*e.spriteVertexBuffer.Binding(0),
		})
		renderPass.BindIndexBuffer(e.spriteIndexBuffer.Binding(0), sdl.GPU_INDEXELEMENTSIZE_32BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{Texture: e.texture, Sampler: e.sampler},
		})
		cmdBuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&cameraMatrix)), unsafe.Sizeof(cameraMatrix),
		))
		renderPass.DrawIndexedPrimitives(
			e.SPRITE_COUNT*6, 1, 0, 0, 0,
		)

		renderPass.End()
	}

	cmdBuf.Submit()

	return nil
}

func (e *ComputeSpriteBatch) Quit(context *common.Context) {
	context.Device.ReleaseComputePipeline(e.computePipeline)
	context.Device.ReleaseGraphicsPipeline(e.renderPipeline)
	context.Device.ReleaseSampler(e.sampler)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseTransferBuffer(e.spriteComputeTransferBuffer.Raw())
	context.Device.ReleaseBuffer(e.spriteComputeBuffer.Raw())
	context.Device.ReleaseBuffer(e.spriteVertexBuffer.Raw())
	context.Device.ReleaseBuffer(e.spriteIndexBuffer.Raw())

	context.Quit()
}
