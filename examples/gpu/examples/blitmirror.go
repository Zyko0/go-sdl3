package examples

import (
	"errors"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type BlitMirror struct {
	texture       *sdl.GPUTexture
	textureWidth  uint32
	textureHeight uint32
}

var BlitMirrorExample = &BlitMirror{}

func (e *BlitMirror) String() string {
	return "BlitMirror"
}

func (e *BlitMirror) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// load the image

	image, err := common.LoadBMP("ravioli.bmp")
	if err != nil {
		return errors.New("failed to load image: " + err.Error())
	}

	e.textureWidth = uint32(image.W)
	e.textureHeight = uint32(image.H)

	// create the texture resource

	e.texture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             e.textureWidth,
		Height:            e.textureHeight,
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	uploadTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  e.textureWidth * e.textureHeight * 4,
		},
	)
	if err != nil {
		return errors.New("failed to create upload transfer buffer: " + err.Error())
	}

	uploadTransferPtr, err := context.Device.MapTransferBuffer(uploadTransferBuffer, false)
	if err != nil {
		return errors.New("failed to create map upload transfer buffer: " + err.Error())
	}

	textureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(uploadTransferPtr)),
		image.W*image.H*4,
	)
	copy(textureData, image.Data)

	context.Device.UnmapTransferBuffer(uploadTransferBuffer)

	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdbuf.BeginCopyPass()

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: uploadTransferBuffer,
	}, &sdl.GPUTextureRegion{
		Texture: e.texture,
		W:       e.textureWidth,
		H:       e.textureHeight,
		D:       1,
	}, false)

	copyPass.End()
	cmdbuf.Submit()

	context.Device.ReleaseTransferBuffer(uploadTransferBuffer)

	return nil
}

func (e *BlitMirror) Update(context *common.Context) error {
	return nil
}

func (e *BlitMirror) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		clearPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
			{
				Texture:    swapchainTexture.Texture,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				StoreOp:    sdl.GPU_STOREOP_STORE,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
				Cycle:      false,
			},
		}, nil)
		clearPass.End()

		// normal
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.texture,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp: sdl.GPU_LOADOP_DONT_CARE,
			Filter: sdl.GPU_FILTER_NEAREST,
		})

		// flipped horizontally
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.texture,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       swapchainTexture.Width / 2,
				W:       swapchainTexture.Width / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp:   sdl.GPU_LOADOP_LOAD,
			FlipMode: sdl.FLIP_HORIZONTAL,
			Filter:   sdl.GPU_FILTER_NEAREST,
		})

		// flipped vertically
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.texture,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width / 2,
				Y:       swapchainTexture.Height / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp:   sdl.GPU_LOADOP_LOAD,
			FlipMode: sdl.FLIP_VERTICAL,
			Filter:   sdl.GPU_FILTER_NEAREST,
		})

		// flipped horizontally and vertically
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.texture,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       swapchainTexture.Width / 2,
				W:       swapchainTexture.Width / 2,
				Y:       swapchainTexture.Height / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp:   sdl.GPU_LOADOP_LOAD,
			FlipMode: sdl.FLIP_HORIZONTAL | sdl.FLIP_VERTICAL,
			Filter:   sdl.GPU_FILTER_NEAREST,
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *BlitMirror) Quit(context *common.Context) {
	context.Device.ReleaseTexture(e.texture)

	context.Quit()
}
