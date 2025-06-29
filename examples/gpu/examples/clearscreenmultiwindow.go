package main

import (
	"errors"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type ClearScreenMultiWindow struct {
	secondWindow *sdl.Window
}

var ClearScreenMultiWindowExample = &ClearScreenMultiWindow{}

func (e *ClearScreenMultiWindow) String() string {
	return "ClearScreenMultiWindow"
}

func (e *ClearScreenMultiWindow) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	e.secondWindow, err = sdl.CreateWindow(
		"ClearScreenMultiWindow (2)", 640, 480, 0,
	)
	if err != nil {
		return errors.New("failed to create window: " + err.Error())
	}

	err = context.Device.ClaimWindow(e.secondWindow)
	if err != nil {
		return errors.New("failed to claim window: " + err.Error())
	}

	return nil
}

func (e *ClearScreenMultiWindow) Update(context *common.Context) error {
	return nil
}

func (e *ClearScreenMultiWindow) Draw(context *common.Context) error {
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

	swapchainTexture, err = cmdbuf.WaitAndAcquireGPUSwapchainTexture(e.secondWindow)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		colorTargetInfo := sdl.GPUColorTargetInfo{
			Texture:    swapchainTexture.Texture,
			ClearColor: sdl.FColor{R: 1, G: 0.5, B: 0.6, A: 1.0},
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

func (e *ClearScreenMultiWindow) Quit(context *common.Context) {
	context.Device.ReleaseWindow(e.secondWindow)
	e.secondWindow.Destroy()
	e.secondWindow = nil

	context.Quit()
}
