package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type GenerateMipmaps struct {
	mipmapTexture *sdl.GPUTexture
}

var GenerateMipmapsExample = &GenerateMipmaps{}

func (e *GenerateMipmaps) String() string {
	return "GenerateMipmaps"
}

func (e *GenerateMipmaps) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	e.mipmapTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER | sdl.GPU_TEXTUREUSAGE_COLOR_TARGET,
		Width:             32,
		Height:            32,
		LayerCountOrDepth: 1,
		NumLevels:         3,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	byteCount := uint32(32 * 32 * 4)
	textureTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  byteCount,
		},
	)
	if err != nil {
		return errors.New("failed to create texture transfer buffer: " + err.Error())
	}

	textureTransferPtr, err := context.Device.MapTransferBuffer(textureTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map texture transfer buffer: " + err.Error())
	}

	textureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(textureTransferPtr)),
		byteCount,
	)

	image, err := common.LoadBMP("cube0.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	copy(textureData, image.Data)

	context.Device.UnmapTransferBuffer(textureTransferBuffer)

	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdbuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: textureTransferBuffer,
	}, &sdl.GPUTextureRegion{
		Texture: e.mipmapTexture,
		W:       32,
		H:       32,
		D:       1,
	}, false)

	copyPass.End()
	cmdbuf.GenerateMipmapsForGPUTexture(e.mipmapTexture)

	cmdbuf.Submit()

	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	return nil
}

func (e *GenerateMipmaps) Update(context *common.Context) error {
	return nil
}

func (e *GenerateMipmaps) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		// blit the smallest mip level
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:  e.mipmapTexture,
				W:        8,
				H:        8,
				MipLevel: 2,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width,
				H:       swapchainTexture.Height,
			},
			LoadOp: sdl.GPU_LOADOP_DONT_CARE,
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *GenerateMipmaps) Quit(context *common.Context) {
	context.Device.ReleaseTexture(e.mipmapTexture)
	context.Quit()
}
