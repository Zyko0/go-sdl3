package examples

import (
	"errors"
	"math"
	"math/rand"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/mathgl/mgl32"
)

// this example is slightly inferior in terms of performance compared to PullSpriteBatch
// despite that, it is included as an example of a compute to graphics workflow.

type ComputeSpriteBatch struct {
	computePipeline             *sdl.GPUComputePipeline
	renderPipeline              *sdl.GPUGraphicsPipeline
	sampler                     *sdl.GPUSampler
	texture                     *sdl.GPUTexture
	spriteComputeTransferBuffer *sdl.GPUTransferBuffer
	spriteComputeBuffer         *sdl.GPUBuffer
	spriteVertexBuffer          *sdl.GPUBuffer
	spriteIndexBuffer           *sdl.GPUBuffer

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
		{
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
		{
			Slot:             0,
			InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
			InstanceStepRate: 0,
			Pitch:            uint32(unsafe.Sizeof(common.PositionTextureColorVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT4,
			Location:   0,
			Offset:     0,
		},
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT2,
			Location:   1,
			Offset:     16,
		},
		{
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

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(image.W * image.H * 4),
		},
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureTransferPtr, err := context.Device.MapTransferBuffer(
		textureTransferBuffer, false,
	)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	textureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr)),
		image.W*image.H*4,
	)

	copy(textureData, image.Data)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

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

	e.spriteComputeTransferBuffer, err = context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
		},
	)
	if err != nil {
		return errors.New("failed to create sprite compute transfer buffer: " + err.Error())
	}

	e.spriteComputeBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_COMPUTE_STORAGE_READ,
		Size:  e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
	})
	if err != nil {
		return errors.New("failed to create sprite compute buffer: " + err.Error())
	}

	e.spriteVertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_COMPUTE_STORAGE_WRITE | sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  e.SPRITE_COUNT * 4 * uint32(unsafe.Sizeof(common.PositionTextureColorVertex{})),
	})
	if err != nil {
		return errors.New("failed to create sprite vertex buffer: " + err.Error())
	}

	e.spriteIndexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  e.SPRITE_COUNT * 6 * uint32(unsafe.Sizeof(uint32(0))),
	})
	if err != nil {
		return errors.New("failed to create sprite vertex buffer: " + err.Error())
	}

	// transfer the up-front data

	indexBufferTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  e.SPRITE_COUNT * 6 * uint32(unsafe.Sizeof(uint32(0))),
		},
	)
	if err != nil {
		return errors.New("failed to create index buffer transfer buffer: " + err.Error())
	}

	indexTransferPtr, err := context.Device.MapTransferBuffer(
		indexBufferTransferBuffer, false,
	)
	if err != nil {
		return errors.New("failed to map index buffer transfer buffer: " + err.Error())
	}

	indexData := unsafe.Slice(
		(*uint32)(unsafe.Pointer(indexTransferPtr)),
		e.SPRITE_COUNT*6,
	)

	for i := range e.SPRITE_COUNT {
		indexData[i*6] = i * 4
		indexData[i*6+1] = i*4 + 1
		indexData[i*6+2] = i*4 + 2
		indexData[i*6+3] = i*4 + 3
		indexData[i*6+4] = i*4 + 2
		indexData[i*6+5] = i*4 + 1
	}

	context.Device.UnmapTransferBuffer(indexBufferTransferBuffer)

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
		Offset:         0, // zeroes out the rest
	}, &sdl.GPUTextureRegion{
		Texture: e.texture,
		W:       uint32(image.W),
		H:       uint32(image.H),
		D:       1,
	}, false)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: indexBufferTransferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: e.spriteIndexBuffer,
			Offset: 0,
			Size:   e.SPRITE_COUNT * 6 * uint32(unsafe.Sizeof(uint32(0))),
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(textureTransferBuffer)
	context.Device.ReleaseTransferBuffer(indexBufferTransferBuffer)

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
		dataPtr, err := context.Device.MapTransferBuffer(
			e.spriteComputeTransferBuffer, true,
		)
		if err != nil {
			return errors.New("failed to map sprite compute transfer buffer: " + err.Error())
		}

		data := unsafe.Slice(
			(*ComputeSpriteInstance)(unsafe.Pointer(dataPtr)),
			e.SPRITE_COUNT*uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
		)

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

		context.Device.UnmapTransferBuffer(e.spriteComputeTransferBuffer)

		// upload instance data
		copyPass := cmdBuf.BeginCopyPass()
		copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
			TransferBuffer: e.spriteComputeTransferBuffer,
			Offset:         0,
		}, &sdl.GPUBufferRegion{
			Buffer: e.spriteComputeBuffer,
			Offset: 0,
			Size:   e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
		}, true)
		copyPass.End()

		// set up compute pass to build vertex buffer
		computePass := cmdBuf.BeginComputePass(
			[]sdl.GPUStorageTextureReadWriteBinding{},
			[]sdl.GPUStorageBufferReadWriteBinding{
				{
					Buffer: e.spriteVertexBuffer,
					Cycle:  true,
				},
			},
		)

		computePass.BindGPUComputePipeline(e.computePipeline)
		computePass.BindStorageBuffers([]*sdl.GPUBuffer{
			e.spriteComputeBuffer,
		})
		computePass.Dispatch(e.SPRITE_COUNT/64, 1, 1)

		computePass.End()

		// render sprites
		renderPass := cmdBuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
			{
				Texture:    swapchainTexture.Texture,
				Cycle:      false,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				StoreOp:    sdl.GPU_STOREOP_STORE,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			},
		}, nil)

		renderPass.BindGraphicsPipeline(e.renderPipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{Buffer: e.spriteVertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.spriteIndexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_32BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			{Texture: e.texture, Sampler: e.sampler},
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
	context.Device.ReleaseTransferBuffer(e.spriteComputeTransferBuffer)
	context.Device.ReleaseBuffer(e.spriteComputeBuffer)
	context.Device.ReleaseBuffer(e.spriteVertexBuffer)
	context.Device.ReleaseBuffer(e.spriteIndexBuffer)

	context.Quit()
}
