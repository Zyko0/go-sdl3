package basicvertexbuffer

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
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

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPU_FILLMODE_FILL
	pipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// create vertex buffer

	vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	// to get data into the vertex buffer, we have to use a transfer buffer

	transferBuffer, err := context.Device.CreateTransferBuffer(&sdl.GPUTransferBufferCreateInfo{
		Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
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
		unsafe.Sizeof(common.PositionColorVertex{})*3,
	)

	transferData[0] = common.PositionColorVertex{X: -1, Y: -1, Z: 0, R: 255, G: 0, B: 0, A: 255}
	transferData[1] = common.PositionColorVertex{X: 1, Y: -1, Z: 0, R: 0, G: 255, B: 0, A: 255}
	transferData[2] = common.PositionColorVertex{X: 0, Y: 1, Z: 0, R: 0, G: 0, B: 255, A: 255}

	context.Device.UnmapTransferBuffer(transferBuffer)

	// upload the transfer data to the vertex buffer

	updateCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := updateCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         0,
		},
		&sdl.GPUBufferRegion{
			Buffer: vertexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 3),
		},
		false,
	)

	copyPass.End()
	updateCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer)

	return nil
}

func update(context *common.Context) error {
	return nil
}

func draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.AcquireGPUSwapchainTexture(context.Window)
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

		renderPass.BindGraphicsPipeline(pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: vertexBuffer, Offset: 0},
		})
		renderPass.DrawPrimitives(3, 1, 0, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(pipeline)
	context.Device.ReleaseBuffer(vertexBuffer)

	context.Quit()
}

var Example = common.Example{
	Name:   "BasicVertexBuffer",
	Init:   _init,
	Update: update,
	Draw:   draw,
	Quit:   quit,
}
