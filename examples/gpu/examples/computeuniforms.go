package main

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type GradientUniforms struct {
	time float32
}

type ComputeUniforms struct {
	gradientPipeline      *sdl.GPUComputePipeline
	gradientRenderTexture *sdl.GPUTexture
	gradientUniformValues GradientUniforms
}

var ComputeUniformsExample = &ComputeUniforms{}

func (e *ComputeUniforms) String() string {
	return "ComputeUniforms"
}

func (e *ComputeUniforms) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	e.gradientPipeline, err = common.CreateComputePipelineFromShader(
		context.Device, "GradientTexture.comp", sdl.GPUComputePipelineCreateInfo{
			NumReadwriteStorageTextures: 1,
			NumUniformBuffers:           1,
			ThreadcountX:                8,
			ThreadcountY:                8,
			ThreadcountZ:                1,
		},
	)
	if err != nil {
		return errors.New("failed to create compute pipeline: " + err.Error())
	}

	w, h, err := context.Window.Size()
	if err != nil {
		return errors.New("failed to get window size: " + err.Error())
	}

	e.gradientRenderTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(w),
		Height:            uint32(h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE | sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	return nil
}

func (e *ComputeUniforms) Update(context *common.Context) error {
	e.gradientUniformValues.time += context.DeltaTime
	return nil
}

func (e *ComputeUniforms) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		computePass := cmdbuf.BeginComputePass(
			[]sdl.GPUStorageTextureReadWriteBinding{
				sdl.GPUStorageTextureReadWriteBinding{
					Texture: e.gradientRenderTexture,
					Cycle:   true,
				},
			},
			[]sdl.GPUStorageBufferReadWriteBinding{},
		)

		computePass.BindGPUComputePipeline(e.gradientPipeline)
		cmdbuf.PushComputeUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&e.gradientUniformValues)),
			unsafe.Sizeof(e.gradientUniformValues),
		))
		computePass.Dispatch(
			swapchainTexture.Width/8, swapchainTexture.Height/8, 1,
		)

		computePass.End()

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.gradientRenderTexture,
				W:       swapchainTexture.Width,
				H:       swapchainTexture.Height,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width,
				H:       swapchainTexture.Height,
			},
			LoadOp: sdl.GPU_LOADOP_DONT_CARE,
			Filter: sdl.GPU_FILTER_LINEAR,
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *ComputeUniforms) Quit(context *common.Context) {
	context.Device.ReleaseComputePipeline(e.gradientPipeline)
	context.Device.ReleaseTexture(e.gradientRenderTexture)
	context.Quit()
}
