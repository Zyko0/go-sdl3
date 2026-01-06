package examples

import (
	"errors"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type ClearScreen struct{}

var ClearScreenExample = &ClearScreen{}

func (e *ClearScreen) String() string {
	return "ClearScreen"
}

func (e *ClearScreen) Init(context *common.Context) error {
	return context.Init(sdl.WINDOW_RESIZABLE)
}

func (e *ClearScreen) Update(context *common.Context) error {
	return nil
}

func (e *ClearScreen) Draw(context *common.Context) error {
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
			ClearColor: sdl.FColor{R: 0.3, G: 0.4, B: 0.5, A: 1.0},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
			StoreOp:    sdl.GPU_STOREOP_STORE,
		}

		renderPass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{colorTargetInfo}, nil,
		)

		renderPass.End()
	}

	cmdbuf.Submit()

	return nil
}

func (e *ClearScreen) Quit(context *common.Context) {
	context.Quit()
}
