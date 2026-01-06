package examples

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

/* Special thanks to Matt Taylor for this overview of ToneMapping: https://64.github.io/ToneMapping/ */

// would also recommend checking out tonymcmapface and agx below
// https://github.com/godotengine/godot/pull/97095
// https://github.com/godotengine/godot/pull/87260
// https://github.com/godotengine/godot/blob/master/servers/rendering/renderer_rd/shaders/effects/tonemap.glsl
// you dont necessarily have to use compute shaders for post processing either
// _maybe im wrong cause it dispatches in groups_

type ToneMapping struct {
	hdrTexture      *sdl.GPUTexture
	toneMapTexture  *sdl.GPUTexture
	transferTexture *sdl.GPUTexture

	swapchainCompositions              [4]sdl.GPUSwapchainComposition
	swapchainCompositionNames          [4]string
	swapchainCompositionCount          int32
	swapchainCompositionSelectionIndex int32
	currentSwapchainComposition        sdl.GPUSwapchainComposition

	tonemapOperatorNames          [4]string
	tonemapOperatorCount          int32
	tonemapOperators              [4]*sdl.GPUComputePipeline
	tonemapOperatorSelectionIndex int32
	currentTonemapOperator        *sdl.GPUComputePipeline

	linearToSRGBPipeline   *sdl.GPUComputePipeline
	linearToST2084Pipeline *sdl.GPUComputePipeline

	w, h int
}

var ToneMappingExample = &ToneMapping{
	swapchainCompositions: [4]sdl.GPUSwapchainComposition{
		sdl.GPU_SWAPCHAINCOMPOSITION_SDR,
		sdl.GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR,
		sdl.GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR,
		sdl.GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084,
	},
	swapchainCompositionNames: [4]string{
		"SDL_GPU_SWAPCHAINCOMPOSITION_SDR",
		"SDL_GPU_SWAPCHAINCOMPOSITION_SDR_LINEAR",
		"SDL_GPU_SWAPCHAINCOMPOSITION_HDR_EXTENDED_LINEAR",
		"SDL_GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084",
	},
	swapchainCompositionCount: 4,

	tonemapOperatorNames: [4]string{
		"Reinhard",
		"ExtendedReinhardLuminance",
		"Hable",
		"ACES",
	},
	tonemapOperatorCount: 4,
}

func (e *ToneMapping) String() string {
	return "ToneMapping"
}

func (e *ToneMapping) ChangeSwapchainComposition(
	context *common.Context, selectionIndex uint32,
) {
	if context.Device.WindowSupportsSwapchainComposition(
		context.Window, e.swapchainCompositions[selectionIndex],
	) {
		e.currentSwapchainComposition = e.swapchainCompositions[selectionIndex]
		fmt.Println("Changing swapchain composition to " + e.swapchainCompositionNames[selectionIndex])
		context.Device.SetSwapchainParameters(
			context.Window, e.currentSwapchainComposition, sdl.GPU_PRESENTMODE_VSYNC,
		)
	} else {
		fmt.Println("Swapchain composition " + e.swapchainCompositionNames[selectionIndex] + " unsupported")
	}
}

func (e *ToneMapping) ChangeTonemapOperator(
	context *common.Context, selectionIndex uint32,
) {
	fmt.Println("Changing tonemap operator to " + e.tonemapOperatorNames[selectionIndex])
	e.currentTonemapOperator = e.tonemapOperators[selectionIndex]
}

func (e *ToneMapping) BuildPostProcessComputePipeline(
	device *sdl.GPUDevice, shaderFilename string,
) (*sdl.GPUComputePipeline, error) {
	return common.CreateComputePipelineFromShader(
		device, shaderFilename, sdl.GPUComputePipelineCreateInfo{
			NumReadonlyStorageTextures:  1,
			NumReadwriteStorageTextures: 1,
			ThreadcountX:                8,
			ThreadcountY:                8,
			ThreadcountZ:                1,
		},
	)
}

func (e *ToneMapping) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	image, err := common.LoadHDR("memorial.hdr")
	if err != nil {
		return errors.New("failed to load hdr image: " + err.Error())
	}
	e.w = image.W
	e.h = image.H

	err = context.Window.SetSize(int32(e.w), int32(e.h))
	if err != nil {
		return errors.New("failed to set window size: " + err.Error())
	}

	e.hdrTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R32G32B32A32_FLOAT,
		Width:             uint32(e.w),
		Height:            uint32(e.h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage: sdl.GPU_TEXTUREUSAGE_SAMPLER |
			sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ,
	})
	if err != nil {
		return errors.New("failed to create hdr texture: " + err.Error())
	}

	e.toneMapTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R16G16B16A16_FLOAT,
		Width:             uint32(e.w),
		Height:            uint32(e.h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage: sdl.GPU_TEXTUREUSAGE_SAMPLER |
			sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_READ |
			sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE,
	})
	if err != nil {
		return errors.New("failed to create hdr texture: " + err.Error())
	}

	e.transferTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(e.w),
		Height:            uint32(e.h),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage: sdl.GPU_TEXTUREUSAGE_SAMPLER |
			sdl.GPU_TEXTUREUSAGE_COMPUTE_STORAGE_WRITE,
	})
	if err != nil {
		return errors.New("failed to create hdr texture: " + err.Error())
	}

	imageDataTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(unsafe.Sizeof(float32(0))) * uint32(4*e.w*e.h),
		},
	)
	if err != nil {
		return errors.New("failed to create transfer buffer: " + err.Error())
	}

	imageTransferPtr, err := context.Device.MapTransferBuffer(
		imageDataTransferBuffer, false,
	)
	if err != nil {
		return errors.New("failed to map transfer buffer: " + err.Error())
	}

	imageTransfer := unsafe.Slice(
		(*float32)(unsafe.Pointer(imageTransferPtr)), 4*e.w*e.h,
	)

	copy(imageTransfer, image.Data)

	context.Device.UnmapTransferBuffer(imageDataTransferBuffer)

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(
		&sdl.GPUTextureTransferInfo{
			TransferBuffer: imageDataTransferBuffer,
			Offset:         0,
		},
		&sdl.GPUTextureRegion{
			Texture: e.hdrTexture,
			W:       uint32(e.w),
			H:       uint32(e.h),
			D:       1,
		},
		false,
	)

	copyPass.End()

	uploadCmdBuf.Submit()

	context.Device.ReleaseTransferBuffer(imageDataTransferBuffer)

	for i, name := range e.tonemapOperatorNames {
		e.tonemapOperators[i], err = e.BuildPostProcessComputePipeline(
			context.Device, "ToneMap"+name+".comp",
		)
		if err != nil {
			return errors.New("failed to build post process compute pipeline for: " + name)
		}
	}

	fmt.Println("yo")
	e.currentTonemapOperator = e.tonemapOperators[0]

	e.linearToSRGBPipeline, err = e.BuildPostProcessComputePipeline(
		context.Device, "LinearToSRGB.comp",
	)
	if err != nil {
		return errors.New("failed to build compute pipeline: " + err.Error())
	}

	e.linearToST2084Pipeline, err = e.BuildPostProcessComputePipeline(
		context.Device, "LinearToST2084.comp",
	)
	if err != nil {
		return errors.New("failed to build compute pipeline: " + err.Error())
	}

	fmt.Println("Press Left/Right to cycle swapchain composition")
	fmt.Println("Press Up/Down to cycle tonemap operators")

	return nil
}

func (e *ToneMapping) Update(context *common.Context) error {
	if context.LeftPressed {
		e.swapchainCompositionSelectionIndex -= 1
		if e.swapchainCompositionSelectionIndex < 0 {
			e.swapchainCompositionSelectionIndex = e.swapchainCompositionCount - 1
		}
		e.ChangeSwapchainComposition(context, uint32(e.swapchainCompositionSelectionIndex))
	} else if context.RightPressed {
		e.swapchainCompositionSelectionIndex = (e.swapchainCompositionSelectionIndex + 1) % e.swapchainCompositionCount

		e.ChangeSwapchainComposition(context, uint32(e.swapchainCompositionSelectionIndex))
	}

	if context.UpPressed {
		e.tonemapOperatorSelectionIndex -= 1
		if e.tonemapOperatorSelectionIndex < 0 {
			e.tonemapOperatorSelectionIndex = e.tonemapOperatorCount - 1
		}

		e.ChangeTonemapOperator(context, uint32(e.tonemapOperatorSelectionIndex))
	} else if context.DownPressed {
		e.tonemapOperatorSelectionIndex = (e.tonemapOperatorSelectionIndex + 1) % e.tonemapOperatorCount

		e.ChangeTonemapOperator(context, uint32(e.tonemapOperatorSelectionIndex))
	}

	return nil
}

func (e *ToneMapping) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		// tonemap

		computePass := cmdbuf.BeginComputePass(
			[]sdl.GPUStorageTextureReadWriteBinding{
				{
					Texture: e.toneMapTexture,
					Cycle:   true,
				},
			},
			[]sdl.GPUStorageBufferReadWriteBinding{},
		)

		computePass.BindGPUComputePipeline(e.currentTonemapOperator)
		computePass.BindStorageTextures([]*sdl.GPUTexture{e.hdrTexture})
		computePass.Dispatch(
			swapchainTexture.Width/8, swapchainTexture.Height/8, 1,
		)
		computePass.End()

		blitSourceTexture := e.toneMapTexture

		// transfer to target color space if necessary

		if e.currentSwapchainComposition == sdl.GPU_SWAPCHAINCOMPOSITION_SDR ||
			e.currentSwapchainComposition == sdl.GPU_SWAPCHAINCOMPOSITION_HDR10_ST2084 {

			computePass = cmdbuf.BeginComputePass(
				[]sdl.GPUStorageTextureReadWriteBinding{
					{
						Texture: e.transferTexture,
						Cycle:   true,
					},
				},
				[]sdl.GPUStorageBufferReadWriteBinding{},
			)

			if e.currentSwapchainComposition == sdl.GPU_SWAPCHAINCOMPOSITION_SDR {
				computePass.BindGPUComputePipeline(e.linearToSRGBPipeline)
			} else {
				computePass.BindGPUComputePipeline(e.linearToST2084Pipeline)
			}

			computePass.BindStorageTextures([]*sdl.GPUTexture{e.toneMapTexture})
			computePass.Dispatch(
				swapchainTexture.Width/8, swapchainTexture.Height/8, 1,
			)
			computePass.End()

			blitSourceTexture = e.transferTexture
		}

		// blit to swapchain

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: blitSourceTexture,
				W:       uint32(e.w),
				H:       uint32(e.h),
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

func (e *ToneMapping) Quit(context *common.Context) {
	for _, pipeline := range e.tonemapOperators {
		context.Device.ReleaseComputePipeline(pipeline)
	}

	context.Device.ReleaseComputePipeline(e.linearToSRGBPipeline)
	context.Device.ReleaseComputePipeline(e.linearToST2084Pipeline)

	context.Device.ReleaseTexture(e.hdrTexture)
	context.Device.ReleaseTexture(e.toneMapTexture)
	context.Device.ReleaseTexture(e.transferTexture)

	context.Quit()
}
