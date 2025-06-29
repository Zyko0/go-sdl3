package instancedindexed

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
	indexBuffer  *sdl.GPUBuffer

	useVertexOffset = false
	useIndexOffset  = false
	useIndexBuffer  = true
)

func _init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create shaders

	vertexShader, err := common.LoadShader(
		context.Device, "PositionColorInstanced.vert", 0, 0, 0, 0,
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
		Size:  uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 9),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(unsafe.Sizeof(uint16(0)) * 6),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	// to get data into the vertex buffer, we have to use a transfer buffer

	transferBuffer, err := context.Device.CreateTransferBuffer(&sdl.GPUTransferBufferCreateInfo{
		Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
		Size: uint32(
			unsafe.Sizeof(common.PositionColorVertex{})*9 +
				unsafe.Sizeof(uint16(0))*6,
		),
	})
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	transferDataPtr, err := context.Device.MapTransferBuffer(transferBuffer, false)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionColorVertex)(unsafe.Pointer(transferDataPtr)),
		unsafe.Sizeof(common.PositionColorVertex{})*9,
	)

	vertexData[0] = common.NewPosColorVert(-1, -1, 0, 255, 0, 0, 255)
	vertexData[1] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[2] = common.NewPosColorVert(0, 1, 0, 0, 0, 255, 255)

	vertexData[3] = common.NewPosColorVert(-1, -1, 0, 255, 165, 0, 255)
	vertexData[4] = common.NewPosColorVert(1, -1, 0, 0, 128, 0, 255)
	vertexData[5] = common.NewPosColorVert(0, 1, 0, 0, 255, 255, 255)

	vertexData[6] = common.NewPosColorVert(-1, -1, 0, 255, 255, 255, 255)
	vertexData[7] = common.NewPosColorVert(1, -1, 0, 255, 255, 255, 255)
	vertexData[8] = common.NewPosColorVert(0, 1, 0, 255, 255, 255, 255)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			transferDataPtr+unsafe.Sizeof(common.PositionColorVertex{})*9,
		)),
		unsafe.Sizeof(uint16(0))*6,
	)

	for i := range 6 {
		indexData[i] = uint16(i)
	}

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
			Buffer: vertexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 9),
		},
		false,
	)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         uint32(unsafe.Sizeof(common.PositionColorVertex{}) * 9),
		},
		&sdl.GPUBufferRegion{
			Buffer: indexBuffer,
			Offset: 0,
			Size:   uint32(unsafe.Sizeof(uint16(0)) * 6),
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer)

	return nil
}

func update(context *common.Context) error {
	if context.LeftPressed {
		useVertexOffset = !useVertexOffset
		fmt.Printf("Using vertex offset: %v\n", useVertexOffset)
	}

	if context.RightPressed {
		useIndexOffset = !useIndexOffset
		fmt.Printf("Using index offset: %v\n", useIndexOffset)
	}

	if context.UpPressed {
		useIndexBuffer = !useIndexBuffer
		fmt.Printf("Using index buffer: %v\n", useIndexBuffer)
	}

	return nil
}

func draw(context *common.Context) error {
	var vertexOffset uint32 = 0
	if useVertexOffset {
		vertexOffset = 3
	}

	var indexOffset uint32 = 0
	if useIndexOffset {
		indexOffset = 3
	}

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

		renderPass.BindGraphicsPipeline(pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			sdl.GPUBufferBinding{Buffer: vertexBuffer, Offset: 0},
		})

		if useIndexBuffer {
			renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
				Buffer: indexBuffer, Offset: 0,
			}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
			renderPass.DrawIndexedPrimitives(
				3, 16, indexOffset, int32(vertexOffset), 0,
			)
		} else {
			renderPass.DrawPrimitives(3, 16, vertexOffset, 0)
		}

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(pipeline)
	context.Device.ReleaseBuffer(vertexBuffer)
	context.Device.ReleaseBuffer(indexBuffer)

	useVertexOffset = false
	useIndexOffset = false
	useIndexBuffer = true

	context.Quit()
}

var Example = common.Example{
	Name:   "InstancedIndexed",
	Init:   _init,
	Update: update,
	Draw:   draw,
	Quit:   quit,
}
