package main

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"reflect"
	"slices"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/content"
	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
	"golang.org/x/image/bmp"
)

type CopyAndReadback struct {
	originalTexture *sdl.GPUTexture
	textureCopy     *sdl.GPUTexture
	textureSmall    *sdl.GPUTexture

	originalBuffer *sdl.GPUBuffer
	bufferCopy     *sdl.GPUBuffer

	textureWidth, textureHeight uint32
}

var CopyAndReadbackExample = &CopyAndReadback{}

func (e *CopyAndReadback) String() string {
	return "CopyAndReadback"
}

func (e *CopyAndReadback) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// load the image

	imgBytes, err := content.ReadFile("images/ravioli.bmp")
	if err != nil {
		return errors.New("failed to read file: " + err.Error())
	}

	img, err := bmp.Decode(bytes.NewReader(imgBytes))

	imgRGBA, ok := img.(*image.NRGBA)
	if !ok {
		return fmt.Errorf("failed to cast: %s", reflect.TypeOf(img))
	}

	e.textureWidth = uint32(imgRGBA.Rect.Size().X)
	e.textureHeight = uint32(imgRGBA.Rect.Size().Y)

	// create texture resources

	e.originalTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imgRGBA.Rect.Size().X),
		Height:            uint32(imgRGBA.Rect.Size().Y),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create original texture: " + err.Error())
	}

	e.textureCopy, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imgRGBA.Rect.Size().X),
		Height:            uint32(imgRGBA.Rect.Size().Y),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create texture copy: " + err.Error())
	}

	e.textureSmall, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Width:             uint32(imgRGBA.Rect.Size().X / 2),
		Height:            uint32(imgRGBA.Rect.Size().Y / 2),
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER | sdl.GPU_TEXTUREUSAGE_COLOR_TARGET,
	})
	if err != nil {
		return errors.New("failed to create texture small: " + err.Error())
	}

	bufferData := [8]uint32{2, 4, 8, 16, 32, 64, 128}
	bufferDataSize := uint32(unsafe.Sizeof(uint32(0)) * uintptr(len(bufferData)))

	e.originalBuffer, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ, // arbitrary
		Size:  bufferDataSize,
	})
	if err != nil {
		return errors.New("failed to create original buffer: " + err.Error())
	}

	e.bufferCopy, err = context.Device.CreateBuffer(&sdl.GPUBufferCreateInfo{
		Usage: sdl.GPU_BUFFERUSAGE_GRAPHICS_STORAGE_READ, // arbitrary
		Size:  bufferDataSize,
	})
	if err != nil {
		return errors.New("failed to create buffer copy: " + err.Error())
	}

	downloadTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_DOWNLOAD,
			Size:  e.textureWidth*e.textureHeight*4 + bufferDataSize,
		},
	)
	if err != nil {
		return errors.New("failed to create download transfer buffer: " + err.Error())
	}

	uploadTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  e.textureWidth*e.textureHeight*4 + bufferDataSize,
		},
	)
	if err != nil {
		return errors.New("failed to create upload transfer buffer: " + err.Error())
	}

	uploadTransferPtr, err := context.Device.MapTransferBuffer(uploadTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map upload transfer buffer: " + err.Error())
	}

	uploadTextureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(uploadTransferPtr)),
		e.textureWidth*e.textureHeight*4,
	)

	copy(uploadTextureData, imgRGBA.Pix)

	uploadBufferData := unsafe.Slice(
		(*uint32)(unsafe.Pointer(uploadTransferPtr+uintptr(e.textureWidth*e.textureHeight*4))),
		len(bufferData),
	)

	copy(uploadBufferData, bufferData[:])

	context.Device.UnmapTransferBuffer(uploadTransferBuffer)

	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := cmdbuf.BeginCopyPass()

	// upload original texture

	copyPass.UploadToGPUTexture(&sdl.GPUTextureTransferInfo{
		TransferBuffer: uploadTransferBuffer,
		Offset:         0,
	}, &sdl.GPUTextureRegion{
		Texture: e.originalTexture,
		W:       e.textureWidth,
		H:       e.textureHeight,
		D:       1,
	}, false)

	// copy original to copy

	copyPass.CopyGPUTextureToTexture(
		&sdl.GPUTextureLocation{
			Texture: e.originalTexture,
			X:       0, Y: 0, Z: 0,
		}, &sdl.GPUTextureLocation{
			Texture: e.textureCopy,
			X:       0, Y: 0, Z: 0,
		},
		e.textureWidth,
		e.textureHeight,
		1, false,
	)

	// upload original buffer

	copyPass.UploadToGPUBuffer(&sdl.GPUTransferBufferLocation{
		TransferBuffer: uploadTransferBuffer,
		Offset:         e.textureWidth * e.textureHeight * 4,
	}, &sdl.GPUBufferRegion{
		Buffer: e.originalBuffer,
		Offset: 0,
		Size:   bufferDataSize,
	}, false)

	// copy original to copy

	copyPass.CopyGPUBufferToBuffer(&sdl.GPUBufferLocation{
		Buffer: e.originalBuffer,
		Offset: 0,
	}, &sdl.GPUBufferLocation{
		Buffer: e.bufferCopy,
		Offset: 0,
	}, bufferDataSize, false)

	copyPass.End()

	// render the half sized version

	cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
		Source: sdl.GPUBlitRegion{
			Texture: e.originalTexture,
			W:       e.textureWidth,
			H:       e.textureHeight,
		},
		Destination: sdl.GPUBlitRegion{
			Texture: e.textureSmall,
			W:       e.textureWidth / 2,
			H:       e.textureHeight / 2,
		},
		LoadOp: sdl.GPU_LOADOP_DONT_CARE,
		Filter: sdl.GPU_FILTER_LINEAR,
	})

	// download the original bytes from the copy

	copyPass = cmdbuf.BeginCopyPass()

	copyPass.DownloadFromGPUTexture(&sdl.GPUTextureRegion{
		Texture: e.textureCopy,
		W:       e.textureWidth,
		H:       e.textureHeight,
		D:       1,
	}, &sdl.GPUTextureTransferInfo{
		TransferBuffer: downloadTransferBuffer,
		Offset:         0,
	})

	copyPass.DownloadFromGPUBuffer(&sdl.GPUBufferRegion{
		Buffer: e.bufferCopy,
		Offset: 0,
		Size:   bufferDataSize,
	}, &sdl.GPUTransferBufferLocation{
		TransferBuffer: downloadTransferBuffer,
		Offset:         e.textureWidth * e.textureHeight * 4,
	})

	copyPass.End()

	fence, err := cmdbuf.SubmitAndAcquireFence()
	if err != nil {
		return errors.New("failed to submit and acquire fence: " + err.Error())
	}

	err = context.Device.WaitForFences(true, []*sdl.GPUFence{fence})
	if err != nil {
		return errors.New("failed to wait for fence: " + err.Error())
	}

	context.Device.ReleaseFence(fence)

	// compare the original bytes to the copied bytes

	downloadTransferPtr, err := context.Device.MapTransferBuffer(
		downloadTransferBuffer, false,
	)
	if err != nil {
		return errors.New("failed to map download transfer: " + err.Error())
	}

	downloadTextureData := unsafe.Slice(
		(*byte)(unsafe.Pointer(downloadTransferPtr)),
		e.textureWidth*e.textureHeight*4,
	)

	if bytes.Equal(imgRGBA.Pix, downloadTextureData) {
		fmt.Println("SUCCESS! Original texture bytes and the downloaded bytes match!")
	} else {
		fmt.Println("FAILURE! Original texture bytes do not match downloaded bytes!")
	}

	downloadBufferData := unsafe.Slice(
		(*uint32)(unsafe.Pointer(downloadTransferPtr+uintptr(e.textureWidth*e.textureHeight*4))),
		len(bufferData),
	)

	if slices.Equal(bufferData[:], downloadBufferData) {
		fmt.Println("SUCCESS! Original buffer bytes and the downloaded bytes match!")
	} else {
		fmt.Println("FAILURE! Original buffer bytes do not match downloaded bytes!")
	}

	context.Device.UnmapTransferBuffer(downloadTransferBuffer)

	// cleanup
	context.Device.ReleaseTransferBuffer(downloadTransferBuffer)
	context.Device.ReleaseTransferBuffer(uploadTransferBuffer)

	return nil
}

func (e *CopyAndReadback) Update(context *common.Context) error {
	return nil
}

func (e *CopyAndReadback) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		clearpass := cmdbuf.BeginRenderPass(
			[]sdl.GPUColorTargetInfo{sdl.GPUColorTargetInfo{
				Texture:    swapchainTexture.Texture,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				StoreOp:    sdl.GPU_STOREOP_STORE,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
				Cycle:      false,
			}}, nil,
		)
		clearpass.End()

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.originalTexture,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       swapchainTexture.Width / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp: sdl.GPU_LOADOP_LOAD,
			Filter: sdl.GPU_FILTER_NEAREST,
		})

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.textureCopy,
				W:       e.textureWidth,
				H:       e.textureHeight,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       swapchainTexture.Width / 2,
				W:       swapchainTexture.Width / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp: sdl.GPU_LOADOP_LOAD,
			Filter: sdl.GPU_FILTER_NEAREST,
		})

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: e.textureSmall,
				W:       e.textureWidth / 2,
				H:       e.textureHeight / 2,
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       swapchainTexture.Width / 4,
				Y:       swapchainTexture.Height / 2,
				W:       swapchainTexture.Width / 2,
				H:       swapchainTexture.Height / 2,
			},
			LoadOp: sdl.GPU_LOADOP_LOAD,
			Filter: sdl.GPU_FILTER_NEAREST,
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *CopyAndReadback) Quit(context *common.Context) {
	context.Device.ReleaseTexture(e.originalTexture)
	context.Device.ReleaseTexture(e.textureCopy)
	context.Device.ReleaseTexture(e.textureSmall)

	context.Device.ReleaseBuffer(e.originalBuffer)
	context.Device.ReleaseBuffer(e.bufferCopy)

	context.Quit()
}
