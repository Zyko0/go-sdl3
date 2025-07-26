package main

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type ComputeSampler struct {
	samplerNames [6]string

	pipeline     *sdl.GPUComputePipeline
	texture      *sdl.GPUTexture
	writeTexture *sdl.GPUTexture
	samplers     [6]*sdl.GPUSampler

	currentSamplerIndex int
}

var ComputeSamplerExample = &ComputeSampler{
	samplerNames: [6]string{
		"PointClamp",
		"PointWrap",
		"LinearClamp",
		"LinearWrap",
		"AnisotropicClamp",
		"AnisotropicWrap",
	},
}

func (e *ComputeSampler) String() string {
	return "ComputeSampler"
}

func (e *ComputeSampler) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// load the image

	imageData, imageWidth, imageHeight, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imageWidth),
		Height:            uint32(imageHeight),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	context.Device.SetTextureName(e.texture, "Ravioli Texture 🖼️")

	e.writeTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             640,
		Height:            480,
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage: sdl.GPU_TEXTUREUSAGE_SAMPLER |
			sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE,
	})

	e.pipeline, err = common.CreateComputePipelineFromShader(
		context.Device, "TexturedQuad.comp",
		sdl.GPUComputePipelineCreateInfo{
			NumSamplers:                 1,
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

	// PointClamp
	e.samplers[0], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_NEAREST,
		MagFilter:    sdl.GPU_FILTER_NEAREST,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_NEAREST,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}
	// PointWrap
	e.samplers[1], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_NEAREST,
		MagFilter:    sdl.GPU_FILTER_NEAREST,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_NEAREST,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}
	// LinearClamp
	e.samplers[2], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_LINEAR,
		MagFilter:    sdl.GPU_FILTER_LINEAR,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_LINEAR,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}
	// LinearWrap
	e.samplers[3], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:    sdl.GPU_FILTER_LINEAR,
		MagFilter:    sdl.GPU_FILTER_LINEAR,
		MipmapMode:   sdl.GPU_SAMPLERMIPMAPMODE_LINEAR,
		AddressModeU: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeV: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeW: sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}
	// AnisotropicClamp
	e.samplers[4], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:        sdl.GPU_FILTER_LINEAR,
		MagFilter:        sdl.GPU_FILTER_LINEAR,
		MipmapMode:       sdl.GPU_SAMPLERMIPMAPMODE_LINEAR,
		AddressModeU:     sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeV:     sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		AddressModeW:     sdl.GPU_SAMPLERADDRESSMODE_CLAMP_TO_EDGE,
		EnableAnisotropy: true,
		MaxAnisotropy:    4,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}
	// AnisotropicWrap
	e.samplers[5], err = context.Device.CreateSampler(&sdl.GPUSamplerCreateInfo{
		MinFilter:        sdl.GPU_FILTER_LINEAR,
		MagFilter:        sdl.GPU_FILTER_LINEAR,
		MipmapMode:       sdl.GPU_SAMPLERMIPMAPMODE_LINEAR,
		AddressModeU:     sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeV:     sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		AddressModeW:     sdl.GPU_SAMPLERADDRESSMODE_REPEAT,
		EnableAnisotropy: true,
		MaxAnisotropy:    4,
	})
	if err != nil {
		return errors.New("failed to create sampler: " + err.Error())
	}

	// set up texture data

	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(imageWidth * imageHeight * 4),
		},
	)

	textureTransferDataPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map texture transfer buffer: " + err.Error())
	}

	textureData := unsafe.Slice(
		(*uint8)(unsafe.Pointer(textureTransferDataPtr)),
		imageWidth*imageHeight*4,
	)

	copy(textureData, imageData)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	// upload the transfer data to the gpu resources

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(
		&sdl.GPUTextureTransferInfo{
			TransferBuffer: textureTransferBuffer,
			Offset:         0,
		},
		&sdl.GPUTextureRegion{
			Texture: e.texture,
			W:       uint32(imageWidth),
			H:       uint32(imageHeight),
			D:       1,
		},
		false,
	)

	copyPass.End()
	uploadCmdBuf.Submit()

	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	// finally, print instructions
	fmt.Println("Press Left/Right to switch between sampler states")
	fmt.Printf("Setting sampler state to: %s\n", e.samplerNames[0])

	return nil
}

func (e *ComputeSampler) Update(context *common.Context) error {
	if context.LeftPressed {
		e.currentSamplerIndex -= 1
		if e.currentSamplerIndex < 0 {
			e.currentSamplerIndex = len(e.samplers) - 1
		}
		fmt.Println("Setting sampler state to: " + e.samplerNames[e.currentSamplerIndex])
	}

	if context.RightPressed {
		e.currentSamplerIndex = (e.currentSamplerIndex + 1) % len(e.samplers)
		fmt.Println("Setting sampler state to: " + e.samplerNames[e.currentSamplerIndex])
	}

	return nil
}

func (e *ComputeSampler) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		computePass := cmdbuf.BeginComputePass([]sdl.GPUStorageTextureReadWriteBinding{
			sdl.GPUStorageTextureReadWriteBinding{
				Texture:  e.writeTexture,
				Layer:    0,
				MipLevel: 0,
				Cycle:    true,
			},
		}, []sdl.GPUStorageBufferReadWriteBinding{})

		computePass.BindGPUComputePipeline(e.pipeline)
		computePass.BindSamplers([]sdl.GPUTextureSamplerBinding{
			sdl.GPUTextureSamplerBinding{
				Texture: e.texture,
				Sampler: e.samplers[e.currentSamplerIndex],
			},
		})

		var texcoordMultiplier float32 = 0.25
		cmdbuf.PushComputeUniformData(0, unsafe.Slice(
			(*byte)(unsafe.Pointer(&texcoordMultiplier)),
			unsafe.Sizeof(float32(0)),
		))

		computePass.Dispatch(
			swapchainTexture.Width/8, swapchainTexture.Height/8, 1,
		)
		computePass.End()

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.writeTexture,
				W:       640,
				H:       480,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width,
				H:       swapchainTexture.Height,
			},
			LoadOp: sdl.GPU_LOADOP_DONT_CARE,
			Filter: sdl.GPU_FILTER_NEAREST,
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *ComputeSampler) Quit(context *common.Context) {
	context.Device.ReleaseComputePipeline(e.pipeline)
	context.Device.ReleaseTexture(e.texture)
	context.Device.ReleaseTexture(e.writeTexture)

	for _, sampler := range e.samplers {
		context.Device.ReleaseSampler(sampler)
	}

	e.currentSamplerIndex = 0

	context.Quit()
}
