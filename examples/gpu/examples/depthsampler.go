package main

import (
	"errors"
	"math"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/mathgl/mgl32"
)

type DepthSampler struct {
	scenePipeline     *sdl.GPUGraphicsPipeline
	sceneVertexBuffer *sdl.GPUBuffer
	sceneIndexBuffer  *sdl.GPUBuffer
	sceneColorTexture *sdl.GPUTexture
	sceneDepthTexture *sdl.GPUTexture

	effectPipeline     *sdl.GPUGraphicsPipeline
	effectVertexBuffer *sdl.GPUBuffer
	effectIndexBuffer  *sdl.GPUBuffer
	effectSampler      *sdl.GPUSampler

	time                    float64
	sceneWidth, sceneHeight uint32
}

var DepthSamplerExample = &DepthSampler{}

func (e *DepthSampler) String() string {
	return "DepthSampler"
}

func (e *DepthSampler) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create the shaders and pipelines
	{
		sceneVertexShader, err := common.LoadShader(
			context.Device, "PositionColorTransform.vert", 0, 1, 0, 0,
		)
		if err != nil {
			return errors.New("failed to create scene vertex shader: " + err.Error())
		}

		sceneFragmentShader, err := common.LoadShader(
			context.Device, "SolidColorDepth.frag", 0, 1, 0, 0,
		)
		if err != nil {
			return errors.New("failed to create scene fragment shader: " + err.Error())
		}

		effectVertexShader, err := common.LoadShader(
			context.Device, "TexturedQuad.vert", 0, 0, 0, 0,
		)
		if err != nil {
			return errors.New("failed to create effect vertex shader: " + err.Error())
		}

		effectFragmentShader, err := common.LoadShader(
			context.Device, "DepthOutline.frag", 2, 1, 0, 0,
		)
		if err != nil {
			return errors.New("failed to create effect fragment shader: " + err.Error())
		}

		sceneColorTargetDescriptions := []sdl.GPUColorTargetDescription{
			sdl.GPUColorTargetDescription{
				Format: sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
			},
		}

		sceneVertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
			sdl.GPUVertexBufferDescription{
				Slot:             0,
				InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
				InstanceStepRate: 0,
				Pitch:            uint32(unsafe.Sizeof(common.PositionColorVertex{})),
			},
		}

		sceneVertexAttributes := []sdl.GPUVertexAttribute{
			sdl.GPUVertexAttribute{
				BufferSlot: 0,
				Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT3,
				Location:   0,
				Offset:     0,
			},
			sdl.GPUVertexAttribute{
				BufferSlot: 0,
				Format:     sdl.GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM,
				Location:   1,
				Offset:     uint32(unsafe.Sizeof(float32(0)) * 3),
			},
		}

		scenePipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
			TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
				ColorTargetDescriptions: sceneColorTargetDescriptions,
				HasDepthStencilTarget:   true,
				DepthStencilFormat:      sdl.GPU_TEXTUREFORMAT_D16_UNORM,
			},
			DepthStencilState: sdl.GPUDepthStencilState{
				EnableDepthTest:   true,
				EnableDepthWrite:  true,
				EnableStencilTest: false,
				CompareOp:         sdl.GPU_COMPAREOP_LESS,
				WriteMask:         0xFF,
			},
			RasterizerState: sdl.GPURasterizerState{
				CullMode:  sdl.GPU_CULLMODE_NONE,
				FillMode:  sdl.GPU_FILLMODE_FILL,
				FrontFace: sdl.GPU_FRONTFACE_COUNTER_CLOCKWISE,
			},
			VertexInputState: sdl.GPUVertexInputState{
				VertexBufferDescriptions: sceneVertexBufferDescriptions,
				VertexAttributes:         sceneVertexAttributes,
			},
			PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
			VertexShader:   sceneVertexShader,
			FragmentShader: sceneFragmentShader,
		}

		e.scenePipeline, err = context.Device.CreateGraphicsPipeline(&scenePipelineCreateInfo)
		if err != nil {
			return errors.New("failed to create scene pipeline: " + err.Error())
		}

		effectColorTargetDescriptions := []sdl.GPUColorTargetDescription{
			sdl.GPUColorTargetDescription{
				Format: context.Device.SwapchainTextureFormat(context.Window),
				BlendState: sdl.GPUColorTargetBlendState{
					EnableBlend:         true,
					SrcColorBlendfactor: sdl.GPU_BLENDFACTOR_ONE,
					DstColorBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
					ColorBlendOp:        sdl.GPU_BLENDOP_ADD,
					SrcAlphaBlendfactor: sdl.GPU_BLENDFACTOR_ONE,
					DstAlphaBlendfactor: sdl.GPU_BLENDFACTOR_ONE_MINUS_SRC_ALPHA,
					AlphaBlendOp:        sdl.GPU_BLENDOP_ADD,
				},
			},
		}

		effectVertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
			sdl.GPUVertexBufferDescription{
				Slot:             0,
				InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
				InstanceStepRate: 0,
				Pitch:            uint32(unsafe.Sizeof(common.PositionTextureVertex{})),
			},
		}

		effectVertexAttributes := []sdl.GPUVertexAttribute{
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

		effectPipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
			TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
				ColorTargetDescriptions: effectColorTargetDescriptions,
			},
			VertexInputState: sdl.GPUVertexInputState{
				VertexBufferDescriptions: effectVertexBufferDescriptions,
				VertexAttributes:         effectVertexAttributes,
			},
			PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
			VertexShader:   effectVertexShader,
			FragmentShader: effectFragmentShader,
		}

		e.effectPipeline, err = context.Device.CreateGraphicsPipeline(&effectPipelineCreateInfo)
		if err != nil {
			return errors.New("failed to create effect pipeline: " + err.Error())
		}

		context.Device.ReleaseShader(effectVertexShader)
		context.Device.ReleaseShader(effectFragmentShader)

		context.Device.ReleaseShader(sceneVertexShader)
		context.Device.ReleaseShader(sceneFragmentShader)
	}

	// create the scene textures
	{
		w, h, err := context.Window.SizeInPixels()
		if err != nil {
			return errors.New("failed to get window size in pixels: " + err.Error())
		}
		e.sceneWidth = uint32(w) / 4
		e.sceneHeight = uint32(h) / 4

		e.sceneColorTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
			Type:              sdl.GPU_TEXTURETYPE_2D,
			Width:             e.sceneWidth,
			Height:            e.sceneHeight,
			LayerCountOrDepth: 1,
			NumLevels:         1,
			SampleCount:       sdl.GPU_SAMPLECOUNT_1,
			Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
			Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER | sdl.GPU_TEXTUREUSAGE_COLOR_TARGET,
		})
		if err != nil {
			return errors.New("failed to create scene color texture: " + err.Error())
		}

		e.sceneDepthTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
			Type:              sdl.GPU_TEXTURETYPE_2D,
			Width:             e.sceneWidth,
			Height:            e.sceneHeight,
			LayerCountOrDepth: 1,
			NumLevels:         1,
			SampleCount:       sdl.GPU_SAMPLECOUNT_1,
			Format:            sdl.GPU_TEXTUREFORMAT_D16_UNORM,
			Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER | sdl.GPU_TEXTUREUSAGE_DEPTH_STENCIL_TARGET,
		})
		if err != nil {
			return errors.New("failed to create scene depth texture: " + err.Error())
		}
	}

	// create outline effect sampler
	e.effectSampler, err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_NEAREST,
		MagFilter:    sdl.GPU_FILTER_NEAREST,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_NEAREST,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}

	// create and upload scene vertex and index buffers
	{
		e.sceneVertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
			Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
			Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 24),
		})
		if err != nil {
			return errors.New("failed to create scene vertex buffer: " + err.Error())
		}

		e.sceneIndexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
			Usage: sdl.GPU_BUFFERUSAGE_INDEX,
			Size:  uint32(unsafe.Sizeof(uint16(0)) * 36),
		})
		if err != nil {
			return errors.New("failed to create scene index buffer: " + err.Error())
		}

		bufferTransferBuffer, err := context.Device.CreateTransferBuffer(
			&sdl.GPUTransferBufferCreateInfo{
				Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
				Size: uint32(
					unsafe.Sizeof(common.PositionColorVertex{})*24 +
						unsafe.Sizeof(uint16(0))*36,
				),
			},
		)
		if err != nil {
			return errors.New("failed to create buffer transfer buffer: " + err.Error())
		}

		bufferTransferDataPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
		if err != nil {
			return errors.New("failed to map buffer transfer buffer: " + err.Error())
		}

		vertexData := unsafe.Slice(
			(*common.PositionColorVertex)(unsafe.Pointer(bufferTransferDataPtr)), 24,
		)

		vertexData[0] = common.NewPosColorVert(-10, -10, -10, 255, 0, 0, 255)
		vertexData[1] = common.NewPosColorVert(10, -10, -10, 255, 0, 0, 255)
		vertexData[2] = common.NewPosColorVert(10, 10, -10, 255, 0, 0, 255)
		vertexData[3] = common.NewPosColorVert(-10, 10, -10, 255, 0, 0, 255)

		vertexData[4] = common.NewPosColorVert(-10, -10, 10, 255, 255, 0, 255)
		vertexData[5] = common.NewPosColorVert(10, -10, 10, 255, 255, 0, 255)
		vertexData[6] = common.NewPosColorVert(10, 10, 10, 255, 255, 0, 255)
		vertexData[7] = common.NewPosColorVert(-10, 10, 10, 255, 255, 0, 255)

		vertexData[8] = common.NewPosColorVert(-10, -10, -10, 255, 0, 255, 255)
		vertexData[9] = common.NewPosColorVert(-10, 10, -10, 255, 0, 255, 255)
		vertexData[10] = common.NewPosColorVert(-10, 10, 10, 255, 0, 255, 255)
		vertexData[11] = common.NewPosColorVert(-10, -10, 10, 255, 0, 255, 255)

		vertexData[12] = common.NewPosColorVert(10, -10, -10, 0, 255, 0, 255)
		vertexData[13] = common.NewPosColorVert(10, 10, -10, 0, 255, 0, 255)
		vertexData[14] = common.NewPosColorVert(10, 10, 10, 0, 255, 0, 255)
		vertexData[15] = common.NewPosColorVert(10, -10, 10, 0, 255, 0, 255)

		vertexData[16] = common.NewPosColorVert(-10, -10, -10, 0, 255, 255, 255)
		vertexData[17] = common.NewPosColorVert(-10, -10, 10, 0, 255, 255, 255)
		vertexData[18] = common.NewPosColorVert(10, -10, 10, 0, 255, 255, 255)
		vertexData[19] = common.NewPosColorVert(10, -10, -10, 0, 255, 255, 255)

		vertexData[20] = common.NewPosColorVert(-10, 10, -10, 0, 0, 255, 255)
		vertexData[21] = common.NewPosColorVert(-10, 10, 10, 0, 0, 255, 255)
		vertexData[22] = common.NewPosColorVert(10, 10, 10, 0, 0, 255, 255)
		vertexData[23] = common.NewPosColorVert(10, 10, -10, 0, 0, 255, 255)

		indexData := unsafe.Slice(
			(*uint16)(unsafe.Pointer(
				bufferTransferDataPtr+unsafe.Sizeof(common.PositionColorVertex{})*24,
			)), 36,
		)
		indices := []uint16{
			0, 1, 2, 0, 2, 3,
			4, 5, 6, 4, 6, 7,
			8, 9, 10, 8, 10, 11,
			12, 13, 14, 12, 14, 15,
			16, 17, 18, 16, 18, 19,
			20, 21, 22, 20, 22, 23,
		}
		copy(indexData, indices)

		context.Device.UnmapTransferBuffer(bufferTransferBuffer)

		// upload the transfer data to the gpu buffers
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
				Buffer: e.sceneVertexBuffer,
				Offset: 0,
				Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 24),
			},
			false,
		)

		copyPass.UploadToGPUBuffer(
			&sdl.GPUTransferBufferLocation{
				TransferBuffer: bufferTransferBuffer,
				Offset:         uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 24),
			},
			&sdl.GPUBufferRegion{
				Buffer: e.sceneIndexBuffer,
				Offset: 0,
				Size:   uint32(unsafe.Sizeof(uint16(0)) * 36),
			},
			false,
		)

		copyPass.End()
		uploadCmdBuf.Submit()
		context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	}

	// create and upload outline effect vertex and index buffers
	{
		e.effectVertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
			Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
			Size:  uint32(unsafe.Sizeof(common.PositionTextureVertex{}) * 4),
		})
		if err != nil {
			return errors.New("failed to create effect vertex buffer: " + err.Error())
		}

		e.effectIndexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
			Usage: sdl.GPU_BUFFERUSAGE_INDEX,
			Size:  uint32(unsafe.Sizeof(uint16(0)) * 6),
		})
		if err != nil {
			return errors.New("failed to create effect index buffer: " + err.Error())
		}

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

		bufferTransferDataPtr, err := context.Device.MapTransferBuffer(bufferTransferBuffer, false)
		if err != nil {
			return errors.New("failed to map buffer transfer buffer: " + err.Error())
		}

		vertexData := unsafe.Slice(
			(*common.PositionTextureVertex)(unsafe.Pointer(bufferTransferDataPtr)), 4,
		)

		vertexData[0] = common.NewPosTexVert(-1, 1, 0, 0, 0)
		vertexData[1] = common.NewPosTexVert(1, 1, 0, 1, 0)
		vertexData[2] = common.NewPosTexVert(1, -1, 0, 1, 1)
		vertexData[3] = common.NewPosTexVert(-1, -1, 0, 0, 1)

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
				Buffer: e.effectVertexBuffer,
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
				Buffer: e.effectIndexBuffer,
				Offset: 0,
				Size:   uint32(unsafe.Sizeof(uint16(0)) * 6),
			},
			false,
		)

		copyPass.End()
		uploadCmdBuf.Submit()
		context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	}

	e.time = 0
	return nil
}

func (e *DepthSampler) Update(context *common.Context) error {
	e.time += float64(context.DeltaTime)
	return nil
}

func (e *DepthSampler) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		// render the 3D scene (color and depth pass)
		var nearPlane float32 = 20
		var farPlane float32 = 60

		proj := mgl32.Perspective(
			75*math.Pi/180,
			float32(e.sceneWidth)/float32(e.sceneHeight),
			nearPlane, farPlane,
		)
		view := mgl32.LookAtV(
			mgl32.Vec3{
				float32(math.Cos(e.time) * 30),
				30,
				float32(math.Sin(e.time) * 30),
			},
			mgl32.Vec3{0, 0, 0},
			mgl32.Vec3{0, 1, 0},
		)

		viewproj := proj.Mul4(view)

		colorTargetInfo := sdl.GPUColorTargetInfo{
			Texture:    e.sceneColorTexture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 0},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}

		depthStencilTargetInfo := sdl.GPUDepthStencilTargetInfo{
			Texture:        e.sceneDepthTexture,
			Cycle:          true,
			ClearDepth:     1,
			ClearStencil:   0,
			LoadOp:         sdl.GPU_LOADOP_CLEAR,
			StoreOp:        sdl.GPU_STOREOP_STORE,
			StencilLoadOp:  sdl.GPU_LOADOP_CLEAR,
			StencilStoreOp: sdl.GPU_STOREOP_STORE,
		}

		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&viewproj)), unsafe.Sizeof(viewproj),
		))

		fragUniformData := []float32{nearPlane, farPlane}
		cmdbuf.PushFragmentUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&fragUniformData[0])), unsafe.Sizeof(fragUniformData),
		))

		renderPass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{colorTargetInfo}, &depthStencilTargetInfo,
		)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: e.sceneVertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.sceneIndexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindGraphicsPipeline(e.scenePipeline)
		renderPass.DrawIndexedPrimitives(
			36, 1, 0, 0, 0,
		)
		renderPass.End()

		// render the outline effect that samples from the color/depth textures
		swapchainTargetInfo := sdl.GPUColorTargetInfo{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0.2, G: 0.5, B: 0.4, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}

		renderPass = cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{swapchainTargetInfo}, nil,
		)
		renderPass.BindGraphicsPipeline(e.effectPipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: e.effectVertexBuffer, Offset: 0},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.effectIndexBuffer, Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{Texture: e.sceneColorTexture, Sampler: e.effectSampler},
			sdl.GPUTextureSamplerBinding{Texture: e.sceneDepthTexture, Sampler: e.effectSampler},
		})
		renderPass.DrawIndexedPrimitives(
			6, 1, 0, 0, 0,
		)
		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *DepthSampler) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.scenePipeline)
	context.Device.ReleaseTexture(e.sceneColorTexture)
	context.Device.ReleaseTexture(e.sceneDepthTexture)
	context.Device.ReleaseBuffer(e.sceneVertexBuffer)
	context.Device.ReleaseBuffer(e.sceneIndexBuffer)

	context.Device.ReleaseGraphicsPipeline(e.effectPipeline)
	context.Device.ReleaseBuffer(e.effectVertexBuffer)
	context.Device.ReleaseBuffer(e.effectIndexBuffer)
	context.Device.ReleaseSampler(e.effectSampler)

	context.Quit()
}
