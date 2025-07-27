package main

import (
	"bytes"
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

// NOTE: This program is intended purely as a test suite for SDL GPU backends.
// It's not reflective of any best practices.

// If all combinations of src and dst texture formats look identical, the test passes.

type TextureTypeTest struct {
	srcTextures     [5]*sdl.GPUTexture
	dstTextures     [5]*sdl.GPUTexture
	srcTextureIndex uint32
	dstTextureIndex uint32

	textureTypeNames [5]string

	baseMipSlices   [5]uint32
	secondMipSlices [5]uint32
}

var TextureTypeTestExample = &TextureTypeTest{
	textureTypeNames: [5]string{
		"2D",
		"2DArray",
		"3D",
		"Cubemap",
		"CubemapArray",
	},
	baseMipSlices:   [5]uint32{0, 1, 1, 1, 7},
	secondMipSlices: [5]uint32{0, 1, 0, 1, 7},
}

func (e *TextureTypeTest) String() string {
	return "TextureTypeTest"
}

func (e *TextureTypeTest) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// load the images

	baseMips := [4]common.Image{}
	secondMips := [4]common.Image{}

	for i := range 4 {
		baseMips[i], err = common.LoadBMP(fmt.Sprintf("cube%d.bmp", i))
		if err != nil {
			return errors.New("failed to load base mip image: " + err.Error())
		}
		secondMips[i], err = common.LoadBMP(fmt.Sprintf("cube%dmip1.bmp", i))
		if err != nil {
			return errors.New("failed to load second mip image: " + err.Error())
		}
	}

	baseMipDataSize := uint32(baseMips[0].W * baseMips[0].H * 4)
	secondMipDataSize := uint32(secondMips[0].W * secondMips[0].H * 4)

	// create the textures

	createInfo := sdl.GPUTextureCreateInfo{
		Usage:     sdl.GPU_TEXTUREUSAGE_SAMPLER,
		Width:     64,
		Height:    64,
		Format:    sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		NumLevels: 2,
	}

	createTexture := func(i uint32, texType sdl.GPUTextureType, layerCount uint32) error {
		createInfo.Type = texType
		createInfo.LayerCountOrDepth = layerCount
		e.srcTextures[i], err = context.Device.CreateTexture(&createInfo)
		if err != nil {
			return errors.New("failed to create " + e.textureTypeNames[i] + " src texture: " + err.Error())
		}
		e.dstTextures[i], err = context.Device.CreateTexture(&createInfo)
		if err != nil {
			return errors.New("failed to create " + e.textureTypeNames[i] + " dst texture: " + err.Error())
		}
		return nil
	}

	createTexture(0, sdl.GPU_TEXTURETYPE_2D, 1)
	createTexture(1, sdl.GPU_TEXTURETYPE_2D_ARRAY, 2)
	createTexture(2, sdl.GPU_TEXTURETYPE_3D, 2)
	createTexture(3, sdl.GPU_TEXTURETYPE_CUBE, 6)
	createTexture(4, sdl.GPU_TEXTURETYPE_CUBE_ARRAY, 12)

	// create and populate the upload transfer buffer

	uploadTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_UPLOAD,
			Size:  uint32(baseMipDataSize * secondMipDataSize * 4),
		},
	)
	if err != nil {
		return errors.New("failed to create upload transfer buffer: " + err.Error())
	}

	transferPtr, err := context.Device.MapTransferBuffer(uploadTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map upload transfer buffer: " + err.Error())
	}

	for i := range 4 {
		offset := uintptr(i) * uintptr(baseMipDataSize+secondMipDataSize)
		baseMipData := unsafe.Slice(
			(*byte)(unsafe.Pointer(transferPtr+offset)),
			baseMipDataSize,
		)
		copy(baseMipData, baseMips[i].Data)
		secondMipData := unsafe.Slice(
			(*byte)(unsafe.Pointer(transferPtr+offset+uintptr(baseMipDataSize))),
			secondMipDataSize,
		)
		copy(secondMipData, secondMips[i].Data)
	}
	context.Device.UnmapTransferBuffer(uploadTransferBuffer)

	// create the download transfer buffer

	downloadTransferBuffer, err := context.Device.CreateTransferBuffer(
		&sdl.GPUTransferBufferCreateInfo{
			Usage: sdl.GPU_TRANSFERBUFFERUSAGE_DOWNLOAD,
			Size:  uint32(baseMipDataSize * secondMipDataSize * 5),
		},
	)
	if err != nil {
		return errors.New("failed to create download transfer buffer: " + err.Error())
	}

	// upload the texture data

	commandBuffer, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	copyPass := commandBuffer.BeginCopyPass()

	for i := range 4 {
		offset := uint32(i) * (baseMipDataSize + secondMipDataSize)

		baseMipInfo := sdl.GPUTextureTransferInfo{
			TransferBuffer: uploadTransferBuffer,
			Offset:         offset,
		}
		secondMipInfo := sdl.GPUTextureTransferInfo{
			TransferBuffer: uploadTransferBuffer,
			Offset:         offset + baseMipDataSize,
		}

		baseMipRegion := sdl.GPUTextureRegion{
			MipLevel: 0, X: uint32(i%2) * 32, Y: uint32(i/2) * 32,
			W: 32, H: 32, D: 1,
		}
		secondMipRegion := sdl.GPUTextureRegion{
			MipLevel: 1, X: uint32(i%2) * 16, Y: uint32(i/2) * 16,
			W: 16, H: 16, D: 1,
		}

		upload := func(textureIndex, layer, baseZ uint32) {
			baseMipRegion.Texture = e.srcTextures[textureIndex]
			baseMipRegion.Layer = layer
			baseMipRegion.Z = baseZ
			copyPass.UploadToGPUTexture(&baseMipInfo, &baseMipRegion, false)

			secondMipRegion.Texture = e.srcTextures[textureIndex]
			secondMipRegion.Layer = layer
			copyPass.UploadToGPUTexture(&secondMipInfo, &secondMipRegion, false)
		}

		upload(0, 0, 0) // 2D
		upload(1, 1, 0) // 2D array
		upload(2, 0, 1) // 3D
		upload(3, 1, 0) // cubemap
		upload(4, 7, 0) // cubemap array

		// on the final section, we'll download to sanity-check that the
		// values are what we'd expect

		if i == 3 {
			baseMipInfo := sdl.GPUTextureTransferInfo{
				TransferBuffer: downloadTransferBuffer,
				Offset:         0,
			}
			secondMipInfo := sdl.GPUTextureTransferInfo{
				TransferBuffer: downloadTransferBuffer,
				Offset:         baseMipDataSize,
			}

			download := func(textureIndex, layer, baseZ uint32) {
				baseMipRegion.Texture = e.srcTextures[textureIndex]
				baseMipRegion.Layer = layer
				baseMipRegion.Z = baseZ
				copyPass.DownloadFromGPUTexture(&baseMipRegion, &baseMipInfo)
				baseMipInfo.Offset += baseMipDataSize + secondMipDataSize

				secondMipRegion.Texture = e.srcTextures[textureIndex]
				secondMipRegion.Layer = layer
				copyPass.DownloadFromGPUTexture(&secondMipRegion, &secondMipInfo)
				secondMipInfo.Offset += baseMipDataSize + secondMipDataSize
			}

			download(0, 0, 0) // 2D
			download(1, 1, 0) // 2D array
			download(2, 0, 1) // 3D
			download(3, 1, 0) // cubemap
			download(4, 7, 0) // cubemap array
		}
	}

	copyPass.End()

	fence, err := commandBuffer.SubmitAndAcquireFence()
	if err != nil {
		return errors.New("failed to submit and acquire fence: " + err.Error())
	}

	// readback

	err = context.Device.WaitForFences(true, []*sdl.GPUFence{fence})
	if err != nil {
		return errors.New("failed to wait for fence: " + err.Error())
	}
	context.Device.ReleaseFence(fence)

	downloadPtr, err := context.Device.MapTransferBuffer(downloadTransferBuffer, false)
	if err != nil {
		return errors.New("failed to map download transfer buffer: " + err.Error())
	}

	for i := range 5 {
		offset := uintptr(i) * uintptr(baseMipDataSize+secondMipDataSize)

		data := unsafe.Slice(
			(*byte)(unsafe.Pointer(downloadPtr+offset)),
			baseMipDataSize,
		)
		if bytes.Equal(data, baseMips[3].Data) {
			fmt.Printf("SUCCESS: Download test for the %s base mip\n", e.textureTypeNames[i])
		} else {
			fmt.Printf("FAILURE: Download test for the %s base mip\n", e.textureTypeNames[i])
		}

		data = unsafe.Slice(
			(*byte)(unsafe.Pointer(downloadPtr+offset+uintptr(baseMipDataSize))),
			secondMipDataSize,
		)
		if bytes.Equal(data, secondMips[3].Data) {
			fmt.Printf("SUCCESS: Download test for the %s second mip\n", e.textureTypeNames[i])
		} else {
			fmt.Printf("FAILURE: Download test for the %s second mip\n", e.textureTypeNames[i])
		}
	}
	context.Device.UnmapTransferBuffer(downloadTransferBuffer)

	// clean up

	context.Device.ReleaseTransferBuffer(uploadTransferBuffer)
	context.Device.ReleaseTransferBuffer(downloadTransferBuffer)

	// set up the program

	e.srcTextureIndex = 0
	e.dstTextureIndex = 0

	fmt.Println("Press Left to cycle through source texture types.")
	fmt.Println("Press Right to cycle through destination texture types.")
	fmt.Println("(2D / 2D)")

	return nil
}

func (e *TextureTypeTest) Update(context *common.Context) error {
	changed := false

	if context.LeftPressed {
		e.srcTextureIndex = (e.srcTextureIndex + 1) % 5
		changed = true
	}

	if context.RightPressed {
		e.dstTextureIndex = (e.dstTextureIndex + 1) % 5
		changed = true
	}

	if changed {
		fmt.Printf(
			"(%s / %s)\n",
			e.textureTypeNames[e.srcTextureIndex],
			e.textureTypeNames[e.dstTextureIndex],
		)
	}

	return nil
}

func (e *TextureTypeTest) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		// clear the screen
		renderPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
			sdl.GPUColorTargetInfo{
				Texture:    swapchainTexture.Texture,
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			},
		}, nil)
		renderPass.End()

		// copy the source to the destination

		copyPass := cmdbuf.BeginCopyPass()

		for i := range 4 {
			x0 := (uint32(i) % 2) * 32
			y0 := (uint32(i) / 2) * 32
			x1 := (uint32(i) % 2) * 16
			y1 := (uint32(i) / 2) * 16

			src := sdl.GPUTextureLocation{
				Texture: e.srcTextures[e.srcTextureIndex],
				Layer:   e.baseMipSlices[e.srcTextureIndex],
				X:       x0, Y: y0, Z: 0,
			}
			if e.srcTextureIndex == 2 {
				src.Layer = 0
				src.Z = e.baseMipSlices[e.srcTextureIndex]
			}

			dst := sdl.GPUTextureLocation{
				Texture: e.dstTextures[e.dstTextureIndex],
				Layer:   e.baseMipSlices[e.dstTextureIndex],
				X:       x0, Y: y0, Z: 0,
			}
			if e.dstTextureIndex == 2 {
				dst.Layer = 0
				dst.Z = e.baseMipSlices[e.dstTextureIndex]
			}
			copyPass.CopyGPUTextureToTexture(&src, &dst, 32, 32, 1, false)

			src.MipLevel = 1
			src.X = x1
			src.Y = y1
			src.Layer = e.secondMipSlices[e.srcTextureIndex]
			if e.srcTextureIndex == 2 {
				src.Layer = 0
				src.Z = e.secondMipSlices[e.srcTextureIndex]
			}

			dst.MipLevel = 1
			dst.X = x1
			dst.Y = y1
			dst.Layer = e.secondMipSlices[e.dstTextureIndex]
			if e.dstTextureIndex == 2 {
				dst.Layer = 0
				dst.Z = e.secondMipSlices[e.dstTextureIndex]
			}
			copyPass.CopyGPUTextureToTexture(&src, &dst, 16, 16, 1, false)
		}

		copyPass.End()

		// blit the source texture and its mip
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:           e.srcTextures[e.srcTextureIndex],
				W:                 64,
				H:                 64,
				LayerOrDepthPlane: e.baseMipSlices[e.srcTextureIndex],
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				W:       128,
				H:       128,
			},
		})
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:           e.srcTextures[e.srcTextureIndex],
				W:                 32,
				H:                 32,
				MipLevel:          1,
				LayerOrDepthPlane: e.secondMipSlices[e.srcTextureIndex],
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       128,
				W:       64,
				H:       64,
			},
		})

		// blit the destination texture and its mip
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:           e.dstTextures[e.dstTextureIndex],
				W:                 64,
				H:                 64,
				LayerOrDepthPlane: e.baseMipSlices[e.dstTextureIndex],
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       256,
				W:       128,
				H:       128,
			},
		})
		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture:           e.dstTextures[e.dstTextureIndex],
				W:                 32,
				H:                 32,
				MipLevel:          1,
				LayerOrDepthPlane: e.secondMipSlices[e.dstTextureIndex],
			},
			Destination: sdl.GPUBlitRegion{
				Texture: swapchainTexture.Texture,
				X:       384,
				W:       64,
				H:       64,
			},
		})
	}

	cmdbuf.Submit()

	return nil
}

func (e *TextureTypeTest) Quit(context *common.Context) {
	for i := range 5 {
		context.Device.ReleaseTexture(e.srcTextures[i])
		context.Device.ReleaseTexture(e.dstTextures[i])
	}

	context.Quit()
}
