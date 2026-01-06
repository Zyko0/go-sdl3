package examples

import (
	"errors"
	"fmt"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type Latency struct {
	lagTexture            *sdl.GPUTexture
	lagX                  uint32
	captureCursor         bool
	allowedFramesInFlight uint32
	fullscreen            bool
}

var LatencyExample = &Latency{
	lagX: 1,
}

func (e *Latency) String() string {
	return "Latency"
}

func (e *Latency) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	e.lagTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Format:            sdl.GPU_TEXTUREFORMAT_R8G8B8A8_UNORM,
		Usage:             sdl.GPU_TEXTUREUSAGE_SAMPLER,
		Width:             8,
		Height:            32,
		LayerCountOrDepth: 1,
		NumLevels:         1,
	})
	if err != nil {
		return errors.New("failed to create texture: " + err.Error())
	}

	byteCount := uint32(8 * 32 * 4)
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

	image, err := common.LoadBMP("latency.bmp")
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
		Texture: e.lagTexture,
		W:       8,
		H:       32,
		D:       1,
	}, false)

	copyPass.End()

	cmdbuf.Submit()

	context.Device.ReleaseTransferBuffer(textureTransferBuffer)

	fmt.Println("Press Left/Right to toggle capturing the mouse cursor.")
	fmt.Println("Press Down to change the number of allowed frames in flight.")
	fmt.Println("Press Up to toggle fullscreen mode.")
	fmt.Println("When the mouse cursor is captured the color directly above the cursor's point in the " +
		"result of the test.")
	fmt.Println("Negative lag can occur when the cursor is below the tear line when tearing is enabled " +
		"as the cursor is only moved during V-blank so it lags the framebuffer update.")
	fmt.Println("  Gray:  -1 frames lag")
	fmt.Println("  White:  0 frames lag")
	fmt.Println("  Green:  1 frames lag")
	fmt.Println("  Yellow: 2 frames lag")
	fmt.Println("  Red:    3 frames lag")
	fmt.Println("  Cyan:   4 frames lag")
	fmt.Println("  Purple: 5 frames lag")
	fmt.Println("  Blue:   6 frames lag")

	e.allowedFramesInFlight = 2
	e.fullscreen = false

	context.Device.SetAllowedFramesInFlight(e.allowedFramesInFlight)

	return nil
}

func (e *Latency) Update(context *common.Context) error {
	if context.LeftPressed || context.RightPressed {
		e.captureCursor = !e.captureCursor
	}

	if context.DownPressed {
		e.allowedFramesInFlight = common.ClampInt(
			(e.allowedFramesInFlight+1)%4, 1, 3,
		)
		context.Device.SetAllowedFramesInFlight(e.allowedFramesInFlight)
		fmt.Printf("Allowed frames in flight: %d\n", e.allowedFramesInFlight)
	}

	if context.UpPressed {
		e.fullscreen = !e.fullscreen
		context.Window.SetFullscreen(e.fullscreen)
	}

	return nil
}

func (e *Latency) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		// get the current mouse cursor position. we use SDL_GetGlobalMouseState() as it actively
		// queries the latest position of the mouse unlike SDL_GetMouseState() which uses a cached
		// value.
		_, cursorX, cursorY := sdl.GetGlobalMouseState()
		winX, winY, _ := context.Window.Position()
		cursorX -= float32(winX)
		cursorY -= float32(winY)

		if e.captureCursor {
			// move the cursor to a known position.
			cursorX = float32(e.lagX)
			context.Window.WarpMouseIn(cursorX, cursorY)
			if e.lagX >= swapchainTexture.Width-8 {
				e.lagX = 1
			} else {
				e.lagX++
			}
		}

		// draw a sprite directly under the cursor if permitted by the blitting engine.
		if cursorX >= 1 && cursorX <= float32(swapchainTexture.Width-8) &&
			cursorY >= 5 && cursorY <= float32(swapchainTexture.Height-27) {
			cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
				Source: sdl.GPUBlitRegion{
					Texture: e.lagTexture,
					W:       8,
					H:       32,
				},
				Destination: sdl.GPUBlitRegion{
					Texture: swapchainTexture.Texture,
					X:       uint32(cursorX) - 1,
					Y:       uint32(cursorY) - 5,
					W:       8,
					H:       32,
				},
				LoadOp:     sdl.GPU_LOADOP_CLEAR,
				ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
			})
		} else {
			renderPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{
				{
					Texture:    swapchainTexture.Texture,
					ClearColor: sdl.FColor{R: 0, G: 0, B: 0, A: 1},
					LoadOp:     sdl.GPU_LOADOP_CLEAR,
					StoreOp:    sdl.GPU_STOREOP_STORE,
				},
			}, nil)
			renderPass.End()
		}
	}

	cmdbuf.Submit()

	return nil
}

func (e *Latency) Quit(context *common.Context) {
	context.Device.ReleaseTexture(e.lagTexture)
	context.Quit()
}
