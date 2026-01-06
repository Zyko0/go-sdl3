package examples

import (
	"errors"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Clear3DSlice struct {
	texture3D *sdl.GPUTexture
}

var Clear3DSliceExample = &Clear3DSlice{}

func (e *Clear3DSlice) String() string {
	return "Clear3DSlice"
}

func (e *Clear3DSlice) Init(context *common.Context) error {
	err := context.Init(sdl.WINDOW_RESIZABLE)
	if err != nil {
		return err
	}

	swapchainFormat := context.Device.SwapchainTextureFormat(context.Window)

	e.texture3D, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_3D,
		Format:            swapchainFormat,
		Width:             64,
		Height:            64,
		LayerCountOrDepth: 4,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_COLOR_TARGET | sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	return nil
}

func (e *Clear3DSlice) Update(context *common.Context) error {
	return nil
}

func (e *Clear3DSlice) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{{
				Texture:           e.texture3D,
				Cycle:             true,
				LoadOp:            sdl.GPU_LOADOP_CLEAR,
				StoreOp:           sdl.GPU_STOREOP_STORE,
				ClearColor:        sdl.FColor{R: 1, G: 0, B: 0, A: 1},
				LayerOrDepthPlane: 0,
			}}, nil,
		).End()

		cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{{
				Texture:           e.texture3D,
				Cycle:             false,
				LoadOp:            sdl.GPU_LOADOP_CLEAR,
				StoreOp:           sdl.GPU_STOREOP_STORE,
				ClearColor:        sdl.FColor{R: 0, G: 1, B: 0, A: 1},
				LayerOrDepthPlane: 1,
			}}, nil,
		).End()

		cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{{
				Texture:           e.texture3D,
				Cycle:             false,
				LoadOp:            sdl.GPU_LOADOP_CLEAR,
				StoreOp:           sdl.GPU_STOREOP_STORE,
				ClearColor:        sdl.FColor{R: 0, G: 0, B: 1, A: 1},
				LayerOrDepthPlane: 2,
			}}, nil,
		).End()

		cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{{
				Texture:           e.texture3D,
				Cycle:             false,
				LoadOp:            sdl.GPU_LOADOP_CLEAR,
				StoreOp:           sdl.GPU_STOREOP_STORE,
				ClearColor:        sdl.FColor{R: 1, G: 0, B: 1, A: 1},
				LayerOrDepthPlane: 3,
			}}, nil,
		).End()

		for i := range 4 {
			destX := uint32(i%2) * (swapchainTexture.Width / 2)
			destY := uint32(0)
			if i > 1 {
				destY = swapchainTexture.Height / 2
			}
			cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
				Source: sdl.GPUBlitRegion{
					Texture:           e.texture3D,
					LayerOrDepthPlane: uint32(i),
					W:                 64,
					H:                 64,
				},
				Destination: sdl.GPUBlitRegion{
					Texture: swapchainTexture.Texture,
					X:       destX,
					Y:       destY,
					W:       swapchainTexture.Width / 2,
					H:       swapchainTexture.Height / 2,
				},
				LoadOp: sdl.GPU_LOADOP_LOAD,
				Filter: sdl.GPU_FILTER_NEAREST,
			})
		}
	}

	cmdbuf.Submit()

	return nil
}

func (e *Clear3DSlice) Quit(context *common.Context) {
	context.Device.ReleaseTexture(e.texture3D)
	context.Quit()
}
