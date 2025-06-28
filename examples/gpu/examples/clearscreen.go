package main

import (
	"errors"

	"github.com/Zyko0/go-sdl3/sdl"
)

func Init(context *Context) error {
	return context.Init(sdl.WINDOW_RESIZABLE)
}

func Update(context *Context) error {
	return nil
}

func Draw(context *Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("AcquireCommandBuffer failed: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.AcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("AcquireGPUSwapchainTexture failed: " + err.Error())
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

func Quit(context *Context) {
	context.Quit()
}

var ClearScreenExample = Example{
	Name:   "ClearScreen",
	Init:   Init,
	Update: Update,
	Draw:   Draw,
	Quit:   Quit,
}
