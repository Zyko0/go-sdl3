package basictriangle

import (
	"errors"
	"fmt"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

var (
	fillPipeline  *sdl.GPUGraphicsPipeline
	linePipeline  *sdl.GPUGraphicsPipeline
	smallViewport = sdl.GPUViewport{
		X: 160, Y: 120, W: 320, H: 240, MinDepth: 0.1, MaxDepth: 1.0,
	}
	scissorRect = sdl.Rect{X: 320, Y: 240, W: 320, H: 240}

	useWireframeMode bool
	useSmallViewport bool
	useScissorRect   bool
)

func _init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create shaders

	vertexShader, err := common.LoadShader(
		context.Device, "RawTriangle.vert", 0, 0, 0, 0,
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

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			NumColorTargets:         uint32(len(colorTargetDescriptions)),
			ColorTargetDescriptions: &colorTargetDescriptions[0],
		},
		PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
	}

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPU_FILLMODE_FILL
	fillPipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create fill pipeline: " + err.Error())
	}

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPU_FILLMODE_LINE
	linePipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create line pipeline: " + err.Error())
	}

	// clean up shaders

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// print instructions

	fmt.Println("Press Left to toggle wireframe mode")
	fmt.Println("Press Down to toggle small viewport")
	fmt.Println("Press Right to toggle scissor rect")

	return nil
}

func update(context *common.Context) error {
	if context.LeftPressed {
		useWireframeMode = !useWireframeMode
	}
	if context.DownPressed {
		useSmallViewport = !useSmallViewport
	}
	if context.RightPressed {
		useScissorRect = !useScissorRect
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

		if useWireframeMode {
			renderPass.BindGraphicsPipeline(linePipeline)
		} else {
			renderPass.BindGraphicsPipeline(fillPipeline)
		}

		if useSmallViewport {
			renderPass.SetGPUViewport(&smallViewport)
		}

		if useScissorRect {
			renderPass.SetScissor(&scissorRect)
		}

		renderPass.DrawPrimitives(3, 1, 0, 0)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(fillPipeline)
	context.Device.ReleaseGraphicsPipeline(linePipeline)

	useWireframeMode = false
	useSmallViewport = false
	useScissorRect = false

	context.Quit()
}

var Example = common.Example{
	Name:   "BasicTriangle",
	Init:   _init,
	Update: update,
	Draw:   draw,
	Quit:   quit,
}
