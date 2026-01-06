package examples

import (
	"errors"
	"fmt"
	"math"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/go-gl/mathgl/mgl32"
)

type Cubemap struct {
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
	indexBuffer  *sdl.GPUBuffer
	texture      *sdl.GPUTexture
	sampler      *sdl.GPUSampler

	clearColors []sdl.FColor

	camPos mgl32.Vec3
}

var CubemapExample = &Cubemap{
	clearColors: []sdl.FColor{
		{R: 1, G: 0, B: 0, A: 1},
		{R: 0, G: 1, B: 0, A: 1},
		{R: 0, G: 0, B: 1, A: 1},
		{R: 1, G: 1, B: 0, A: 1},
		{R: 1, G: 0, B: 1, A: 1},
		{R: 0, G: 1, B: 1, A: 1},
	},
	camPos: mgl32.Vec3{0, 0, 4},
}

func (e *Cubemap) String() string {
	return "Cubemap"
}

func (e *Cubemap) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create the shaders

	vertexShader, err := common.LoadShader(
		context.Device, "Skybox.vert", 0, 1, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "Skybox.frag", 1, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	// create the pipeline

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
			Pitch:            uint32(unsafe.Sizeof(common.PositionVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT3,
			Location:   0,
			Offset:     0,
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

	// create the gpu resources

	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionVertex{}) * 24),
	})
	if err != nil {
		return errors.New("failed to create vertex buffer: " + err.Error())
	}

	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(unsafe.Sizeof(uint16(0)) * 36),
	})
	if err != nil {
		return errors.New("failed to create index buffer: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_CUBE,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             64,
		Height:            64,
		LayerCountOrDepth: 6,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_COLOR_TARGET | sdl.GPU_TEXTUREUSAGE_SAMPLER,
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
				unsafe.Sizeof(common.PositionVertex{})*24 +
					unsafe.Sizeof(uint16(0))*36,
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
		(*common.PositionVertex)(unsafe.Pointer(bufferTransferPtr)), 24,
	)

	vertexData[0] = common.NewPosVert(-10, -10, -10)
	vertexData[1] = common.NewPosVert(10, -10, -10)
	vertexData[2] = common.NewPosVert(10, 10, -10)
	vertexData[3] = common.NewPosVert(-10, 10, -10)

	vertexData[4] = common.NewPosVert(-10, -10, 10)
	vertexData[5] = common.NewPosVert(10, -10, 10)
	vertexData[6] = common.NewPosVert(10, 10, 10)
	vertexData[7] = common.NewPosVert(-10, 10, 10)

	vertexData[8] = common.NewPosVert(-10, -10, -10)
	vertexData[9] = common.NewPosVert(-10, 10, -10)
	vertexData[10] = common.NewPosVert(-10, 10, 10)
	vertexData[11] = common.NewPosVert(-10, -10, 10)

	vertexData[12] = common.NewPosVert(10, -10, -10)
	vertexData[13] = common.NewPosVert(10, 10, -10)
	vertexData[14] = common.NewPosVert(10, 10, 10)
	vertexData[15] = common.NewPosVert(10, -10, 10)

	vertexData[16] = common.NewPosVert(-10, -10, -10)
	vertexData[17] = common.NewPosVert(-10, -10, 10)
	vertexData[18] = common.NewPosVert(10, -10, 10)
	vertexData[19] = common.NewPosVert(10, -10, -10)

	vertexData[20] = common.NewPosVert(-10, 10, -10)
	vertexData[21] = common.NewPosVert(-10, 10, 10)
	vertexData[22] = common.NewPosVert(10, 10, 10)
	vertexData[23] = common.NewPosVert(10, 10, -10)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			bufferTransferPtr+unsafe.Sizeof(common.PositionVertex{})*24,
		)), 36,
	)
	indices := []uint16{
		0, 1, 2, 0, 2, 3,
		6, 5, 4, 7, 6, 4,
		8, 9, 10, 8, 10, 11,
		14, 13, 12, 15, 14, 12,
		16, 17, 18, 16, 18, 19,
		22, 21, 20, 23, 22, 20,
	}
	copy(indexData, indices)

	context.Device.UnmapTransferBuffer(bufferTransferBuffer)

	// upload the transfer data to the gpu resources

	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdbuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         0,
	}, &sdl.GPUBufferRegion{
		Buffer: e.vertexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(common.PositionVertex{}) * 24),
	}, false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         uint32(unsafe.Sizeof(common.PositionVertex{}) * 24),
	}, &sdl.GPUBufferRegion{
		Buffer: e.indexBuffer,
		Offset: 0,
		Size:   uint32(unsafe.Sizeof(uint16(0)) * 36),
	}, false)

	copyPass.End()
	context.Device.ReleaseTransferBuffer(bufferTransferBuffer)

	// clear the faces of the cube texture
	for i := range 6 {
		renderPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
			{
				Texture:           e.texture,
				LayerOrDepthPlane: uint32(i),
				ClearColor:        e.clearColors[i],
				LoadOp:            sdl.GPU_LOADOP_CLEAR,
				StoreOp:           sdl.GPU_STOREOP_STORE,
			},
		}, nil)
		renderPass.End()
	}

	cmdbuf.Submit()

	// print the instructions
	fmt.Println("Press Left/Right to view the opposite direction!")

	return nil
}

func (e *Cubemap) Update(context *common.Context) error {
	if context.LeftPressed || context.RightPressed {
		e.camPos[2] *= -1
	}

	return nil
}

func (e *Cubemap) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		proj := mgl32.Perspective(
			75*math.Pi/180,
			float32(640)/480,
			0.01, 100,
		)
		view := mgl32.LookAtV(
			e.camPos, mgl32.Vec3{0, 0, 0}, mgl32.Vec3{0, 1, 0},
		)

		viewproj := proj.Mul4(view)

		colorTargetInfos := []sdl.GPUColorTargetInfo{{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}}

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
		cmdbuf.PushVertexUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&viewproj)),
			unsafe.Sizeof(viewproj),
		))
		renderPass.DrawIndexedPrimitives(36, 1, 0, 0, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *Cubemap) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseSampler(e.sampler)

	e.camPos[2] = float32(math.Abs(float64(e.camPos[2])))

	context.Quit()
}
