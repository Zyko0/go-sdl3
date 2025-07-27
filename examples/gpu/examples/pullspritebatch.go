package main

import (
	"errors"
	"math"
	"math/rand"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/mathgl/mgl32"
)

type PullSpriteBatch struct {
	renderPipeline           *sdl.GPUGraphicsPipeline
	sampler                  *sdl.GPUSampler
	texture                  *sdl.GPUTexture
	spriteDataTransferBuffer *sdl.GPUTransferBuffer
	spriteDataBuffer         *sdl.GPUBuffer

	SPRITE_COUNT uint32
}

type PullSpriteInstance struct {
	X, Y, Z                float32
	Rotation               float32
	W, H, _, _             float32
	TexU, TexV, TexW, TexH float32
	R, G, B, A             float32
}

var PullSpriteBatchExample = &PullSpriteBatch{
	SPRITE_COUNT: 8192,
}

func (e *PullSpriteBatch) String() string {
	return "PullSpriteBatch"
}

func (e *PullSpriteBatch) Init(context *common.Context) error {
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
		context.Device, "PullSpriteBatch.vert", 0, 1, 1, 0,
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

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			NumColorTargets:         uint32(len(colorTargetDescriptions)),
			ColorTargetDescriptions: &colorTargetDescriptions[0],
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

	// load the image data

	imageData, imageWidth, imageHeight, err := common.LoadBMP("ravioli_atlas.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(imageWidth * imageHeight * 4),
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
		imageWidth*imageHeight*4,
	)

	copy(textureData, imageData)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	// create the gpu resources

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imageWidth),
		Height:            uint32(imageHeight),
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

	e.spriteDataTransferBuffer, err = context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
		},
	)
	if err != nil {
		return errors.New("failed to create sprite compute transfer buffer: " + err.Error())
	}

	e.spriteDataBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ,
		Size:  e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
	})
	if err != nil {
		return errors.New("failed to create sprite compute buffer: " + err.Error())
	}

	// transfer the up-front data

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
		W:       uint32(imageWidth),
		H:       uint32(imageHeight),
		D:       1,
	}, false)

	copyPass.End()
	uploadCmdBuf.Submit()

	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	return nil
}

func (e *PullSpriteBatch) Update(context *common.Context) error {
	return nil
}

func (e *PullSpriteBatch) Draw(context *common.Context) error {
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
			e.spriteDataTransferBuffer, true,
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

		context.Device.UnmapTransferBuffer(e.spriteDataTransferBuffer)

		// upload instance data
		copyPass := cmdBuf.BeginCopyPass()
		copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
			TransferBuffer: e.spriteDataTransferBuffer,
			Offset:         0,
		}, &sdl.GPUBufferRegion{
			Buffer: e.spriteDataBuffer,
			Offset: 0,
			Size:   e.SPRITE_COUNT * uint32(unsafe.Sizeof(ComputeSpriteInstance{})),
		}, true)
		copyPass.End()

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
		renderPass.BindVertexStorageBuffers([]*sdl.GPUBuffer{
			e.spriteDataBuffer,
		})
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{Texture: e.texture, Sampler: e.sampler},
		})
		cmdBuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&cameraMatrix)), unsafe.Sizeof(cameraMatrix),
		))
		renderPass.DrawPrimitives(
			e.SPRITE_COUNT*6, 1, 0, 0,
		)

		renderPass.End()
	}

	cmdBuf.Submit()

	return nil
}

func (e *PullSpriteBatch) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.renderPipeline)
	context.Device.ReleaseSampler(e.sampler)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseTransferBuffer(e.spriteDataTransferBuffer)
	context.Device.ReleaseBuffer(e.spriteDataBuffer)

	context.Quit()
}
