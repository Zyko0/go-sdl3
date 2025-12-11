package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Resolution struct {
	X, Y int32
}

type WindowResize struct {
	pipeline *sdl.GPUGraphicsPipeline

	resolutions []Resolution

	resolutionIndex int
}

var WindowResizeExample = &WindowResize{
	resolutions: []Resolution{
		{X: 640, Y: 480},
		{X: 1280, Y: 720},
		{X: 1024, Y: 1024},
		{X: 1600, Y: 900},
		{X: 1920, Y: 1080},
		{X: 3200, Y: 1800},
		{X: 3840, Y: 2160},
	},
}

func (e *WindowResize) String() string {
	return "WindowResize"
}

func (e *WindowResize) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	e.resolutionIndex = 0

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

	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		{
			Format: context.Device.SwapchainTextureFormat(context.Window),
		},
	}

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			ColorTargetDescriptions: colorTargetDescriptions,
		},
		PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
		RasterizerState: sdl.GPURasterizerState{
			FillMode: sdl.GPU_FILLMODE_FILL,
		},
	}

	pipelineCreateInfo.RasterizerState.FillMode = sdl.GPU_FILLMODE_FILL
	e.pipeline, err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
	if err != nil {
		return errors.New("failed to create pipeline: " + err.Error())
	}

	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	fmt.Println("Press Left and Right to resize the window!")

	return nil
}

func (e *WindowResize) Update(context *common.Context) error {
	changeResolution := false

	if context.RightPressed {
		e.resolutionIndex = (e.resolutionIndex + 1) % len(e.resolutions)
		changeResolution = true
	}

	if context.LeftPressed {
		e.resolutionIndex -= 1
		if e.resolutionIndex < 0 {
			e.resolutionIndex = len(e.resolutions) - 1
		}
		changeResolution = true
	}

	if changeResolution {
		currentResolution := e.resolutions[e.resolutionIndex]
		fmt.Printf("Setting resolution to: %d, %d\n", currentResolution.X, currentResolution.Y)
		context.Window.SetSize(currentResolution.X, currentResolution.Y)
		context.Window.SetPosition(sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED)
		context.Window.Sync()
	}

	return nil
}

func (e *WindowResize) Draw(context *common.Context) error {
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
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *WindowResize) Quit(context *common.Context) {
	context.Device.ReleaseGraphicsPipeline(e.pipeline)
	context.Quit()
}
