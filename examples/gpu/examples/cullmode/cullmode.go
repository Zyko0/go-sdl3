package cullmode

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	modeNames = []string{
		"CW_CullNone",
		"CW_CullFront",
		"CW_CullBack",
		"CCW_CullNone",
		"CCW_CullFront",
		"CCW_CullBack",
	}

	pipelines   [6]*sdl.GPUGraphicsPipeline
	currentMode int

	vertexBufferCW  *sdl.GPUBuffer
	vertexBufferCCW *sdl.GPUBuffer
)

func _init(context *common.Context) error {
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

	for i := range len(pipelines) {
		pipelineCreateInfo.RasterizerState.CullMode = sdl.GPUCullMode(i % 3)
		if i > 2 {
			pipelineCreateInfo.RasterizerState.FrontFace = sdl.GPU_FRONTFACE_CLOCKWISE
		} else {
			pipelineCreateInfo.RasterizerState.FrontFace = sdl.GPU_FRONTFACE_COUNTER_CLOCKWISE
		}

		pipelines[i], err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
		if err != nil {
			return errors.New("failed to create pipeline: " + err.Error())
		}
	}

	// clean up shaders

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// create vertex buffer. they're the same except for the vertex order

	vertexBufferCW, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	vertexBufferCCW, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	// setup the transfer buffer

	transferBuffer, err := context.Device.CreateTransferBuffer(&sdl.GPUTransferBufferCreateInfo{
		Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 6),
	})
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	transferDataPtr, err := context.Device.MapTransferBuffer(transferBuffer, false)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	transferData := unsafe.Slice(
		(*common.PositionColorVertex)(unsafe.Pointer(transferDataPtr)),
		unsafe.Sizeof(common.PositionColorVertex{})*6,
	)

	transferData[0] = common.NewPosColorVert(-1, -1, 0, 255, 0, 0, 255)
	transferData[1] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	transferData[2] = common.NewPosColorVert(0, 1, 0, 0, 0, 255, 255)
	transferData[3] = common.NewPosColorVert(0, 1, 0, 255, 0, 0, 255)
	transferData[4] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	transferData[5] = common.NewPosColorVert(-1, -1, 0, 0, 0, 255, 255)

	context.Device.UnmapTransferBuffer(transferBuffer)

	// upload the transfer data to the vertex buffer

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: vertexBufferCW,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
		},
		false,
	)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
		},
		&sdl.GPUBufferRegion{
			Buffer: vertexBufferCCW,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer)

	// print instructions

	fmt.Println("Press Left/Right to switch between modes")
	fmt.Println("Current Mode: " + modeNames[currentMode])

	return nil
}

func update(context *common.Context) error {
	if context.LeftPressed {
		currentMode--
		if currentMode < 0 {
			currentMode = len(pipelines) - 1
		}
		fmt.Println("Current Mode: " + modeNames[currentMode])
	}

	if context.RightPressed {
		currentMode = (currentMode + 1) % len(pipelines)
		fmt.Println("Current Mode: " + modeNames[currentMode])
	}

	return nil
}

func draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire gpu swapchain texture: " + err.Error())
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
		renderPass.BindGraphicsPipeline(pipelines[currentMode])
		renderPass.SetGPUViewport(&sdl.GPUViewport{X: 0, Y: 0, W: 320, H: 480})
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: vertexBufferCW, Offset: 0},
		})
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.SetGPUViewport(&sdl.GPUViewport{X: 320, Y: 0, W: 320, H: 480})
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: vertexBufferCCW, Offset: 0},
		})
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func quit(context *common.Context) {
	for _, pipeline := range pipelines {
		context.Device.ReleaseGraphicsPipeline(pipeline)
	}

	context.Device.ReleaseBuffer(vertexBufferCW)
	context.Device.ReleaseBuffer(vertexBufferCCW)

	currentMode = 0

	context.Quit()
}

var Example = common.Example{
	Name:   "CullMode",
	Init:   _init,
	Update: update,
	Draw:   draw,
	Quit:   quit,
}
