package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type DrawIndirect struct {
	pipeline     *sdl.GPUGraphicsPipeline
	vertexBuffer *sdl.GPUBuffer
	indexBuffer  *sdl.GPUBuffer
	drawBuffer   *sdl.GPUBuffer
}

var DrawIndirectExample = &DrawIndirect{}

func (e *DrawIndirect) String() string {
	return "DrawIndirect"
}

func (e *DrawIndirect) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create the shaders

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
			Pitch:            uint32(unsafe.Sizeof(common.PositionColorVertex{})),
		},
	}

	vertexAttributes := []sdl.GPUVertexAttribute{
		{
			BufferSlot: 0,
			Format:     sdl.GPU_VERTEXELEMENTFORMAT_FLOAT3,
			Location:   0,
			Offset:     0,
		},
		{
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

	e.pipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// create the buffers

	const vertexBufferSize = unsafe.Sizeof(common.PositionColorVertex{}) * 10
	e.vertexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_VERTEX,
		Size:  uint32(vertexBufferSize),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	const indexBufferSize = unsafe.Sizeof(uint16(0)) * 6
	e.indexBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDEX,
		Size:  uint32(indexBufferSize),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	const drawBufferSize = (unsafe.Sizeof(sdl.GPUIndexedIndirectDrawCommand{}) * 1) +
		(unsafe.Sizeof(sdl.GPUIndirectDrawCommand{}) * 2)

	e.drawBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_INDIRECT,
		Size:  uint32(drawBufferSize),
	})
	if err != nil {
		return errors.New("failed to create buffer: " + err.Error())
	}

	// setup the buffer data

	transferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size: uint32(
				vertexBufferSize + indexBufferSize + drawBufferSize,
			),
		},
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	transferDataPtr, err := context.Device.MapTransferBuffer(transferBuffer, false)
	if err != nil {
		return errors.New("failed to map buffer transfer buffer: " + err.Error())
	}

	vertexData := unsafe.Slice(
		(*common.PositionColorVertex)(unsafe.Pointer(transferDataPtr)), 10,
	)

	vertexData[0] = common.NewPosColorVert(-1, -1, 0, 255, 0, 0, 255)
	vertexData[1] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[2] = common.NewPosColorVert(1, 1, 0, 0, 0, 255, 255)
	vertexData[3] = common.NewPosColorVert(-1, 1, 0, 255, 255, 255, 255)

	vertexData[4] = common.NewPosColorVert(1, -1, 0, 0, 255, 0, 255)
	vertexData[5] = common.NewPosColorVert(0, -1, 0, 0, 0, 255, 255)
	vertexData[6] = common.NewPosColorVert(0.5, 1, 0, 255, 0, 0, 255)
	vertexData[7] = common.NewPosColorVert(-1, -1, 0, 0, 255, 0, 255)
	vertexData[8] = common.NewPosColorVert(0, -1, 0, 0, 0, 255, 255)
	vertexData[9] = common.NewPosColorVert(-0.5, 1, 0, 255, 0, 0, 255)

	indexData := unsafe.Slice(
		(*uint16)(unsafe.Pointer(
			transferDataPtr+vertexBufferSize,
		)), 6,
	)

	indexData[0] = 0
	indexData[1] = 1
	indexData[2] = 2
	indexData[3] = 0
	indexData[4] = 2
	indexData[5] = 3

	indexDrawCommand := (*sdl.GPUIndexedIndirectDrawCommand)(
		unsafe.Pointer(transferDataPtr + vertexBufferSize + indexBufferSize),
	)
	indexDrawCommand.NumIndices = 6
	indexDrawCommand.NumInstances = 1

	drawCommands := unsafe.Slice(
		(*sdl.GPUIndirectDrawCommand)(unsafe.Pointer(
			transferDataPtr+vertexBufferSize+indexBufferSize+
				unsafe.Sizeof(sdl.GPUIndexedIndirectDrawCommand{}),
		)),
		unsafe.Sizeof(sdl.GPUIndirectDrawCommand{})*2,
	)

	drawCommands[0].NumVertices = 3
	drawCommands[0].NumInstances = 1
	drawCommands[0].FirstVertex = 4

	drawCommands[1].NumVertices = 3
	drawCommands[1].NumInstances = 1
	drawCommands[1].FirstVertex = 7

	context.Device.UnmapTransferBuffer(transferBuffer)

	// upload the transfer data to the gpu buffers

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
			Buffer: e.vertexBuffer,
			Offset: 0,
			Size:   uint32(vertexBufferSize),
		},
		false,
	)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         uint32(vertexBufferSize),
		},
		&sdl.GPUBufferRegion{
			Buffer: e.indexBuffer,
			Offset: 0,
			Size:   uint32(indexBufferSize),
		},
		false,
	)

	copyPass.UploadToGPUBuffer(
		&sdl.GPUTransferBufferLocation{
			TransferBuffer: transferBuffer,
			Offset:         uint32(vertexBufferSize + indexBufferSize),
		},
		&sdl.GPUBufferRegion{
			Buffer: e.drawBuffer,
			Offset: 0,
			Size:   uint32(drawBufferSize),
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()
	context.Device.ReleaseTransferBuffer(transferBuffer)

	return nil
}

func (e *DrawIndirect) Update(context *common.Context) error {
	return nil
}

func (e *DrawIndirect) Draw(context *common.Context) error {
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

		renderPass.BindGraphicsPipeline(e.pipeline)
		renderPass.BindVertexBuffers([]sdl.GPUBufferBinding{
			{
				Buffer: e.vertexBuffer,
				Offset: 0,
			},
		})
		renderPass.BindIndexBuffer(&sdl.GPUBufferBinding{
			Buffer: e.indexBuffer,
			Offset: 0,
		}, sdl.GPU_INDEXELEMENTSIZE_16BIT)
		renderPass.DrawIndexedPrimitivesIndirect(e.drawBuffer, 0, 1)
		renderPass.DrawPrimitivesIndirect(
			e.drawBuffer,
			uint32(unsafe.Sizeof(sdl.GPUIndexedIndirectDrawCommand{})),
			2,
		)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *DrawIndirect) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Device.ReleaseBuffer(e.vertexBuffer)
	context.Device.ReleaseBuffer(e.indexBuffer)
	context.Device.ReleaseBuffer(e.drawBuffer)

	context.Quit()
}
