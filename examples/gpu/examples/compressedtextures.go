package main

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/sdl/sdlgpu"
)

const (
	BC_IMAGE_COUNT   = 12
	ASTC_IMAGE_COUNT = 14
)

type CompressedTextures struct {
	srcTextures [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]*sdl.GPUTexture
	dstTextures [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]*sdl.GPUTexture

	textureFormats [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]sdl.GPUTextureFormat
	textureNames   [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]string

	currentTextureIndex int
}

var CompressedTexturesExample = &CompressedTextures{
	textureFormats: [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]sdl.GPUTextureFormat{
		// BCn formats
		sdl.GPU_TEXTUREFORMAT_BC1_RGBA_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC2_RGBA_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC3_RGBA_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC4_R_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC5_RG_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC6H_RGB_FLOAT,
		sdl.GPU_TEXTUREFORMAT_BC6H_RGB_UFLOAT,
		sdl.GPU_TEXTUREFORMAT_BC7_RGBA_UNORM,
		sdl.GPU_TEXTUREFORMAT_BC1_RGBA_UNORM_SRGB,
		sdl.GPU_TEXTUREFORMAT_BC2_RGBA_UNORM_SRGB,
		sdl.GPU_TEXTUREFORMAT_BC3_RGBA_UNORM_SRGB,
		sdl.GPU_TEXTUREFORMAT_BC7_RGBA_UNORM_SRGB,

		// ASTC formats
		sdl.GPU_TEXTUREFORMAT_ASTC_4x4_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_5x4_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_5x5_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_6x5_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_6x6_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_8x5_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_8x6_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_8x8_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_10x5_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_10x6_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_10x8_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_10x10_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_12x10_UNORM,
		sdl.GPU_TEXTUREFORMAT_ASTC_12x12_UNORM,
	},
	textureNames: [BC_IMAGE_COUNT + ASTC_IMAGE_COUNT]string{
		"bcn/BC1.dds",
		"bcn/BC2.dds",
		"bcn/BC3.dds",
		"bcn/BC4.dds",
		"bcn/BC5.dds",
		"bcn/BC6H_S.dds",
		"bcn/BC6H_U.dds",
		"bcn/BC7.dds",
		"bcn/BC1_SRGB.dds",
		"bcn/BC2_SRGB.dds",
		"bcn/BC3_SRGB.dds",
		"bcn/BC7_SRGB.dds",

		"astc/4x4.astc",
		"astc/5x4.astc",
		"astc/5x5.astc",
		"astc/6x5.astc",
		"astc/6x6.astc",
		"astc/8x5.astc",
		"astc/8x6.astc",
		"astc/8x8.astc",
		"astc/10x5.astc",
		"astc/10x6.astc",
		"astc/10x8.astc",
		"astc/10x10.astc",
		"astc/12x10.astc",
		"astc/12x12.astc",
	},
}

func (e *CompressedTextures) String() string {
	return "CompressedTextures"
}

func (e *CompressedTextures) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	var downloadTransferBuffer *sdlgpu.TypedTransferBuffer[byte]
	var firstTextureData []byte

	uploadCmdBuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := uploadCmdBuf.BeginCopyPass()

	// upload texture data
	for i := range e.srcTextures {
		if !context.Device.TextureSupportsFormat(
			e.textureFormats[i], sdl.GPU_TEXTURETYPE_2D, sdl.GPU_TEXTUREUSAGE_SAMPLER,
		) {
			continue
		}

		// load the image
		var image common.Image
		if i < BC_IMAGE_COUNT {
			image, err = common.LoadDDS(e.textureNames[i])
			if err != nil {
				return errors.New("failed to load dds texture: " + err.Error())
			}
		} else {
			image, err = common.LoadASTC(e.textureNames[i])
			if err != nil {
				return errors.New("failed to load astc texture: " + err.Error())
			}
		}

		// create the texture
		createInfo := sdl.GPUTextureCreateInfo{
			Format:            e.textureFormats[i],
			Width:             uint32(image.W),
			Height:            uint32(image.H),
			LayerCountOrDepth: 1,
			Type:              sdl.GPU_TEXTURETYPE_2D,
			Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
			NumLevels:         1,
		}
		e.srcTextures[i], err = context.Device.CreateTexture(&createInfo)
		if err != nil {
			return errors.New("failed to create src texture: " + err.Error())
		}
		e.dstTextures[i], err = context.Device.CreateTexture(&createInfo)
		if err != nil {
			return errors.New("failed to create dst texture: " + err.Error())
		}

		// set up texture transfer data
		textureTransferBuffer, err := sdlgpu.CreateTypedTransferBuffer[byte](
			context.Device, sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD, uint32(len(image.Data)), 0,
		)
		if err != nil {
			return errors.New("failed to create texture transfer buffer: " + err.Error())
		}

		textureData, err := textureTransferBuffer.Map(context.Device, false)
		if err != nil {
			return errors.New("failed to map texture transfer buffer: " + err.Error())
		}

		copy(textureData, image.Data)

		context.Device.UnmapTransferBuffer(textureTransferBuffer.Raw())

		// upload the texture data
		copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
			TransferBuffer: textureTransferBuffer.Raw(),
			Offset:         0, // zeros out the rest
		}, &sdl.GPUTextureRegion{
			Texture: e.srcTextures[i],
			W:       uint32(image.W),
			H:       uint32(image.H),
			D:       1,
		}, false)

		copyPass.CopyGPUTextureToTexture(&sdl.GPUTextureLocation{
			Texture: e.srcTextures[i],
		}, &sdl.GPUTextureLocation{
			Texture: e.dstTextures[i],
		}, 256, 256, 1, false)

		context.Device.ReleaseTransferBuffer(textureTransferBuffer.Raw())

		// testing if downloads work...
		if i == 0 {
			downloadTransferBuffer, err = sdlgpu.CreateTypedTransferBuffer[byte](
				context.Device, sdl.GPU_TRANSFERBUFFERUSAGE_DOWNLOAD, uint32(len(image.Data)), 0,
			)
			if err != nil {
				return errors.New("failed to create download transfer buffer: " + err.Error())
			}

			copyPass.DownloadFromGPUTexture(&sdl.GPUTextureRegion{
				Texture: e.srcTextures[i],
				W:       256,
				H:       256,
				D:       1,
			}, &sdl.GPUTextureTransferInfo{
				TransferBuffer: downloadTransferBuffer.Raw(),
			})

			firstTextureData = image.Data
		}
	}

	copyPass.End()

	fence, err := uploadCmdBuf.SubmitAndAcquireFence()
	if err != nil {
		return errors.New("failed to submit and acquire fence: " + err.Error())
	}

	// read the downloaded data
	err = context.Device.WaitForFences(true, []*sdl.GPUFence{fence})
	if err != nil {
		return errors.New("failed to wait for fence: " + err.Error())
	}
	context.Device.ReleaseFence(fence)

	downloadedData, err := downloadTransferBuffer.Map(context.Device, false)
	if err != nil {
		return errors.New("failed to map download transfer buffer: " + err.Error())
	}
	if bytes.Equal(downloadedData, firstTextureData) {
		fmt.Println("Success: Downloaded bytes match original texture bytes!")
	} else {
		fmt.Println("Failure: Downloaded bytes match original texture bytes!")
	}
	context.Device.UnmapTransferBuffer(downloadTransferBuffer.Raw())
	context.Device.ReleaseTransferBuffer(downloadTransferBuffer.Raw())

	// finally, print instructions
	fmt.Println("Press Left/Right to switch between textures")

	return nil
}

func (e *CompressedTextures) Update(context *common.Context) error {
	changed := false

	if context.LeftPressed {
		e.currentTextureIndex -= 1
		if e.currentTextureIndex < 0 {
			e.currentTextureIndex = len(e.srcTextures) - 1
		}
		changed = true
	}

	if context.RightPressed {
		e.currentTextureIndex = (e.currentTextureIndex + 1) % len(e.srcTextures)
		changed = true
	}

	if changed {
		if e.srcTextures[e.currentTextureIndex] == nil {
			fmt.Printf("Unsupported texture format: %s\n", e.textureNames[e.currentTextureIndex])
		} else {
			fmt.Printf("Setting texture to: %s\n", e.textureNames[e.currentTextureIndex])
		}
	}

	return nil
}

func (e *CompressedTextures) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		if e.srcTextures[e.currentTextureIndex] != nil {
			cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
				ClearColor: sdl.FColor{R: 1, G: 1, B: 1, A: 1},
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				Source: sdl.GPUBlitRegion{
					Texture: e.srcTextures[e.currentTextureIndex],
					W:       256,
					H:       256,
				},
				Destination: sdl.GPUBlitRegion{
					Texture: swapchainTexture.Texture,
					W:       256,
					H:       256,
				},
			})
			cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
				Source: sdl.GPUBlitRegion{
					Texture: e.dstTextures[e.currentTextureIndex],
					W:       256,
					H:       256,
				},
				Destination: sdl.GPUBlitRegion{
					Texture: swapchainTexture.Texture,
					X:       384,
					W:       256,
					H:       256,
				},
			})
		} else {
			renderPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
				sdl.GPUColorTargetInfo{
					Texture:    swapchainTexture.Texture,
					LoadOp:     sdl.GPU_LOADOP_CLEAR,
					ClearColor: sdl.FColor{R: 1, G: 1, B: 1, A: 1},
				},
			}, nil)
			renderPass.End()
		}
	}

	cmdbuf.Submit()

	return nil
}

func (e *CompressedTextures) Quit(context *common.Context) {
	for i := range e.srcTextures {
		context.Device.ReleaseTexture(e.srcTextures[i])
		context.Device.ReleaseTexture(e.dstTextures[i])
		e.srcTextures[i] = nil
		e.dstTextures[i] = nil
	}

	e.currentTextureIndex = 0

	context.Quit()
}
