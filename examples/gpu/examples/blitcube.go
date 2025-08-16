package main

import (
	"errors"
	"fmt"
	"math"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/sdl/sdlgpu"
	"github.com/go-gl/mathgl/mgl32"
)

type BlitCube struct {
	pipeline           *sdl.GPUGraphicsPipeline
	vertexBuffer       *sdlgpu.TypedBuffer[common.PositionVertex]
	indexBuffer        *sdlgpu.TypedBuffer[uint16]
	sourceTexture      *sdl.GPUTexture
	destinationTexture *sdl.GPUTexture
	sampler            *sdl.GPUSampler

	camPos mgl32.Vec3
}

var BlitCubeExample = &BlitCube{
	camPos: mgl32.Vec3{0, 0, 4},
}

func (e *BlitCube) String() string {
	return "BlitCube"
}

func (e *BlitCube) Init(context *common.Context) error {
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
		sdl.GPUColorTargetDescription{
			Format: context.Device.SwapchainTextureFormat(context.Window),
		},
	}

	vertexBufferDescriptions := []sdl.GPUVertexBufferDescription{
		sdl.GPUVertexBufferDescription{
			Slot:             0,
			InputRate:        sdl.GPU_VERTEXINPUTRATE_VERTEX,
			InstanceStepRate: 0,
			Pitch:            uint32(unsafe.Sizeof(common.PositionVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		sdl.GPUVertexAttribute{
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

	e.vertexBuffer, err = sdlgpu.CreateTypedBuffer[common.PositionVertex](
		context.Device, sdl.GPU_BUFFERUSAGE_VERTEX, 24, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex buffer: " + err.Error())
	}

	e.indexBuffer, err = sdlgpu.CreateTypedBuffer[uint16](
		context.Device, sdl.GPU_BUFFERUSAGE_INDEX, 36, 0,
	)
	if err != nil {
		return errors.New("failed to create index buffer: " + err.Error())
	}

	e.sourceTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Type:              sdl.GPU_TEXTURETYPE_CUBE,
		Width:             32,
		Height:            32,
		LayerCountOrDepth: 6,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	e.destinationTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Type:              sdl.GPU_TEXTURETYPE_CUBE,
		Width:             32,
		Height:            32,
		LayerCountOrDepth: 6,
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

	// set up texture data

	bytesPerImage := 32 * 32 * 4

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(bytesPerImage * 6),
		},
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureTransferPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map texture transfer buffer: " + err.Error())
	}

	imageNames := []string{
		"cube0.bmp", "cube1.bmp", "cube2.bmp",
		"cube3.bmp", "cube4.bmp", "cube5.bmp",
	}

	for i := range imageNames {
		image, err := common.LoadBMP(imageNames[i])
		if err != nil {
			return errors.New("failed to load " + imageNames[i] + ": " + err.Error())
		}
		textureData := unsafe.Slice(
			(*byte)(unsafe.Pointer(textureTransferPtr+(uintptr(bytesPerImage*i)))),
			bytesPerImage,
		)
		copy(textureData, image.Data)
	}

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	// upload the transfer data to the gpu resources

	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdbuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         0,
	}, e.vertexBuffer.Region(0, 24), false)

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: bufferTransferBuffer,
		Offset:         uint32(unsafe.Sizeof(common.PositionVertex{}) * 24),
	}, e.indexBuffer.Region(0, 36), false)

	for i := range 6 {
		copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
			TransferBuffer: textureTransferBuffer,
			Offset:         uint32(bytesPerImage * i),
		}, &sdl.GPUTextureRegion{
			Texture: e.sourceTexture,
			Layer:   uint32(i),
			W:       32,
			H:       32,
			D:       1,
		}, false)
	}

	copyPass.End()

	// blit to destination texture.
	// this serves no real purpose other than demonstrating cube->cube blits are possible!

	for i := range 6 {
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:           e.sourceTexture,
				LayerOrDepthPlane: uint32(i),
				W:                 32,
				H:                 32,
			},
			Destination: sdl.GPUBlitRegion{
				Texture:           e.destinationTexture,
				LayerOrDepthPlane: uint32(i),
				W:                 32,
				H:                 32,
			},
			LoadOp: sdl.GPU_LOADOP_DONT_CARE,
			Filter: sdl.GPU_FILTER_LINEAR,
		})
	}

	context.Device.ReleaseTransferBuffer(bufferTransferBuffer)
	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	cmdbuf.Submit()

	// print the instructions
	fmt.Println("Press Left/Right to view the opposite direction!")

	return nil
}

func (e *BlitCube) Update(context *common.Context) error {
	if context.LeftPressed || context.RightPressed {
		e.camPos[2] *= -1
	}

	return nil
}

func (e *BlitCube) Draw(context *common.Context) error {
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

		colorTargetInfos := []sdl.GPUColorTargetInfo{sdl.GPUColorTargetInfo{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}}

		renderPass := cmdbuf.BeginRenderPass(colorTargetInfos, nil)

		renderPass.BindGraphicsPipeline(e.pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			*e.vertexBuffer.Binding(0),
		})
		renderPass.BindIndexBuffer(e.indexBuffer.Binding(0), sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.BindFragmentSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{
				Texture: e.destinationTexture, Sampler: e.sampler,
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

func (e *BlitCube) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer.Raw())
	context.Device.ReleaseBuffer(e.indexBuffer.Raw())
	context.Device.ReleaseTexture(e.sourceTexture)
	context.Device.ReleaseTexture(e.destinationTexture)
	context.Device.ReleaseSampler(e.sampler)

	e.camPos[2] = float32(math.Abs(float64(e.camPos[2])))

	context.Quit()
}
