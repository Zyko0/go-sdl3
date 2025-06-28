package common

import (
	"errors"
	"fmt"

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
		return errors.New("failed to create gpu device: " + err.Error())
	}

	fmt.Println("Driver: " + context.Device.Driver())

	context.Window, err = sdl.CreateWindow(context.ExampleName, 640, 480, windowFlags)
	if err != nil {
		return errors.New("failed to create window: " + err.Error())
	}

	err = context.Device.ClaimWindow(context.Window)
	if err != nil {
		return errors.New("failed to claim window: " + err.Error())
	}

	return nil
}

func (context *Context) Quit() {
	context.Device.ReleaseWindow(context.Window)
	context.Window.Destroy()
	context.Device.Destroy()
}
