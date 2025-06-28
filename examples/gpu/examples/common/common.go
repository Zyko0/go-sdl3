package common

import (
	"errors"

	"github.com/Zyko0/go-sdl3/sdl"
)

type Context struct {
	ExampleName  string
	BasePath     string
	Window       *sdl.Window
	Device       *sdl.GPUDevice
	LeftPressed  bool
	RightPressed bool
	DownPressed  bool
	UpPressed    bool
	DeltaTime    float32
}

func (context *Context) Init(windowFlags sdl.WindowFlags) error {
	var err error
	context.Device, err = sdl.CreateGPUDevice(
		sdl.GPU_SHADERFORMAT_SPIRV|sdl.GPU_SHADERFORMAT_DXIL|sdl.GPU_SHADERFORMAT_MSL,
		true,
		"",
	)
	if err != nil {
		return errors.New("GPUCreateDevice failed: " + err.Error())
	}

	context.Window, err = sdl.CreateWindow(context.ExampleName, 640, 480, windowFlags)
	if err != nil {
		return errors.New("CreateWindow failed: " + err.Error())
	}

	err = context.Device.ClaimWindow(context.Window)
	if err != nil {
		return errors.New("GPUClaimWindow failed: " + err.Error())
	}

	return nil
}

func (context *Context) Quit() {
	context.Device.ReleaseWindow(context.Window)
	context.Window.Destroy()
	context.Device.Destroy()
}

type Example struct {
	Name   string
	Init   func(context *Context) error
	Update func(context *Context) error
	Draw   func(context *Context) error
	Quit   func(context *Context)
}
