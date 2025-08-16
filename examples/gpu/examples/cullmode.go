package main

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/sdl/sdlgpu"
)

type CullMode struct {
	modeNames [6]string

	pipelines   [6]*sdl.GPUGraphicsPipeline
	currentMode int

	vertexBufferCW  *sdlgpu.TypedBuffer[common.PositionColorVertex]
	vertexBufferCCW *sdlgpu.TypedBuffer[common.PositionColorVertex]
}

var CullModeExample = &CullMode{
	modeNames: [6]string{
		"CCW_CullNone",
		"CCW_CullFront",
		"CCW_CullBack",
		"CW_CullNone",
		"CW_CullFront",
		"CW_CullBack",
	},
}

func (e *CullMode) String() string {
	return "CullMode"
}

func (e *CullMode) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create shaders

	vertexShader, err := common.LoadShader(
		context.Device, "PositionColor.vert", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "SolidColor.frag", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	// create pipelines

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
			Pitch:            uint32(unsafe.Sizeof(common.PositionColorVertex{})),
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
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_UBYTE4_NORM,
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

	for i := range len(e.pipelines) {
		pipelineCreateInfo.RasterizerState.CullMode = sdl.GPUCullMode(i % 3)
		if i > 2 {
			pipelineCreateInfo.RasterizerState.FrontFace = sdl.GPU_FRONTFACE_CLOCKWISE
		} else {
			pipelineCreateInfo.RasterizerState.FrontFace = sdl.GPU_FRONTFACE_COUNTER_CLOCKWISE
		}

		e.pipelines[i], err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
		if err != nil {
			return errors.New("failed to create pipeline: " + err.Error())
		}
	}

	// clean up shaders

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// create vertex buffer. they're the same except for the vertex order

	e.vertexBufferCW, err = sdlgpu.CreateTypedBuffer[common.PositionColorVertex](
		context.Device, sdl.GPU_BUFFERUSAGE_VERTEX, 3, 0,
	)
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	e.vertexBufferCCW, err = sdlgpu.CreateTypedBuffer[common.PositionColorVertex](
		context.Device, sdl.GPU_BUFFERUSAGE_VERTEX, 3, 0,
	)
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	// setup the transfer buffer

	transferBuffer, err := sdlgpu.CreateTypedTransferBuffer[common.PositionColorVertex](
		context.Device,
		sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
		e.vertexBufferCW.Length()+e.vertexBufferCCW.Length(), 0,
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	vertexData, err := transferBuffer.Map(context.Device, false)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	vertexData[0] = common.NewPosColorVert(-1, -1, 0, 255, 0, 0, 255)
	vertexData[1] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[2] = common.NewPosColorVert(0, 1, 0, 0, 0, 255, 255)
	vertexData[3] = common.NewPosColorVert(0, 1, 0, 255, 0, 0, 255)
	vertexData[4] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[5] = common.NewPosColorVert(-1, -1, 0, 0, 0, 255, 255)

	context.Device.UnmapTransferBuffer(transferBuffer.Raw())

	// upload the transfer data to the vertex buffer

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		transferBuffer.Location(0), e.vertexBufferCW.Region(0, e.vertexBufferCW.Length()), false,
	)
	copyPass.UploadToGPUBuffer(
		transferBuffer.Location(3), e.vertexBufferCCW.Region(0, e.vertexBufferCCW.Length()), false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer.Raw())

	// print instructions

	fmt.Println("Press Left/Right to switch between modes")
	fmt.Println("Current Mode: " + e.modeNames[e.currentMode])

	return nil
}

func (e *CullMode) Update(context *common.Context) error {
	if context.LeftPressed {
		e.currentMode--
		if e.currentMode < 0 {
			e.currentMode = len(e.pipelines) - 1
		}
		fmt.Println("Current Mode: " + e.modeNames[e.currentMode])
	}

	if context.RightPressed {
		e.currentMode = (e.currentMode + 1) % len(e.pipelines)
		fmt.Println("Current Mode: " + e.modeNames[e.currentMode])
	}

	return nil
}

func (e *CullMode) Draw(context *common.Context) error {
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
		renderPass.BindGraphicsPipeline(e.pipelines[e.currentMode])
		renderPass.SetGPUViewport(&sdl.GPUViewport{X: 0, Y: 0, W: 320, H: 480})
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			*e.vertexBufferCW.Binding(0),
		})
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.SetGPUViewport(&sdl.GPUViewport{X: 320, Y: 0, W: 320, H: 480})
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			*e.vertexBufferCCW.Binding(0),
		})
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *CullMode) Quit(context *common.Context) {
	for _, pipeline := range e.pipelines {
		context.Device.ReleaseGraphicsPipeline(pipeline)
	}

	context.Device.ReleaseBuffer(e.vertexBufferCW.Raw())
	context.Device.ReleaseBuffer(e.vertexBufferCCW.Raw())

	e.currentMode = 0

	context.Quit()
}
