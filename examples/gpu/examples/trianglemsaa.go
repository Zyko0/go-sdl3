package main

import (
	"errors"
	"fmt"

	"github.com/Zyko0/go-sdl3/examples/gpu/examples/common"
	"github.com/Zyko0/go-sdl3/sdl"
)

type TriangleMSAA struct {
	pipelines          [4]*sdl.GPUGraphicsPipeline
	msaaRenderTextures [4]*sdl.GPUTexture
	resolveTexture     *sdl.GPUTexture
	sampleCounts       int

	currentSampleCount int
}

var TriangleMSAAExample = &TriangleMSAA{}

func (e *TriangleMSAA) String() string {
	return "TriangleMSAA"
}

func (e *TriangleMSAA) Init(context *common.Context) error {
	err := context.Init(0)
	if err != nil {
		return err
	}

	// create the shaders

	vertexShader, err := common.LoadShader(
		context.Device, "RawTriangle.vert", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create vertex shader: " + err.Error())
	}

	fragmentShader, err := common.LoadShader(
		context.Device, "SolidColor.frag", 0, 0, 0, 0,
	)
	if err != nil {
		return errors.New("failed to create fragment shader: " + err.Error())
	}

	// create the pipelines

	rtFormat := context.Device.SwapchainTextureFormat(context.Window)
	colorTargetDescriptions := []sdl.GPUColorTargetDescription{
		sdl.GPUColorTargetDescription{
			Format: rtFormat,
		},
	}

	pipelineCreateInfo := sdl.GPUGraphicsPipelineCreateInfo{
		TargetInfo: sdl.GPUGraphicsPipelineTargetInfo{
			ColorTargetDescriptions: colorTargetDescriptions,
		},
		PrimitiveType:  sdl.GPU_PRIMITIVETYPE_TRIANGLELIST,
		VertexShader:   vertexShader,
		FragmentShader: fragmentShader,
	}

	e.sampleCounts = 0
	for i := range e.pipelines {
		sampleCount := sdl.GPUSampleCount(i)
		if !context.Device.TextureSupportsSampleCount(rtFormat, sampleCount) {
			fmt.Printf("sample count %d not supported\n", (1 << sampleCount))
			continue
		}

		pipelineCreateInfo.MultisampleState.SampleCount = sampleCount
		e.pipelines[e.sampleCounts], err = context.Device.CreateGraphicsPipeline(&pipelineCreateInfo)
		if err != nil {
			return errors.New("failed to create pipeline: " + err.Error())
		}

		// create the render target textures
		textureCreateInfo := sdl.GPUTextureCreateInfo{
			Type:              sdl.GPU_TEXTURETYPE_2D,
			Width:             640,
			Height:            480,
			LayerCountOrDepth: 1,
			NumLevels:         1,
			Format:            rtFormat,
			Usage:             sdl.GPU_TEXTUREUSAGE_COLOR_TARGET,
			SampleCount:       sampleCount,
		}
		if sampleCount == sdl.GPU_SAMPLECOUNT_1 {
			textureCreateInfo.Usage |= sdl.GPU_TEXTUREUSAGE_SAMPLER
		}
		e.msaaRenderTextures[e.sampleCounts], err = context.Device.CreateTexture(&textureCreateInfo)
		if err != nil {
			fmt.Println("failed to create msaa render target texture")
			context.Device.ReleaseGraphicsPipeline(e.pipelines[e.sampleCounts])
			continue
		}

		e.sampleCounts++
	}

	// create the resolve texture

	e.resolveTexture, err = context.Device.CreateTexture(&sdl.GPUTextureCreateInfo{
		Type:              sdl.GPU_TEXTURETYPE_2D,
		Width:             640,
		Height:            480,
		LayerCountOrDepth: 1,
		NumLevels:         1,
		Format:            rtFormat,
		Usage:             sdl.GPU_TEXTUREUSAGE_COLOR_TARGET | sdl.GPU_TEXTUREUSAGE_SAMPLER,
	})
	if err != nil {
		return errors.New("failed to create resolve texture: " + err.Error())
	}

	// clean up shader resources
	context.Device.ReleaseShader(vertexShader)
	context.Device.ReleaseShader(fragmentShader)

	// print the instructions
	fmt.Println("Press Left/Right to cycle between sample counts")
	fmt.Printf("Current sample count: %d\n", (1 << e.currentSampleCount))

	return nil
}

func (e *TriangleMSAA) Update(context *common.Context) error {
	changed := false

	if context.LeftPressed {
		e.currentSampleCount -= 1
		if e.currentSampleCount < 0 {
			e.currentSampleCount = e.sampleCounts - 1
		}
		changed = true
	}

	if context.RightPressed {
		e.currentSampleCount = (e.currentSampleCount + 1) % e.sampleCounts
		changed = true
	}

	if changed {
		fmt.Printf("Current sample count: %d\n", (1 << e.currentSampleCount))
	}

	return nil
}

func (e *TriangleMSAA) Draw(context *common.Context) error {
	cmdbuf, err := context.Device.AcquireCommandBuffer()
	if err != nil {
		return errors.New("failed to acquire command buffer: " + err.Error())
	}

	swapchainTexture, err := cmdbuf.WaitAndAcquireGPUSwapchainTexture(context.Window)
	if err != nil {
		return errors.New("failed to acquire swapchain texture: " + err.Error())
	}

	if swapchainTexture != nil {
		colorTargetInfo := sdl.GPUColorTargetInfo{
			Texture:    e.msaaRenderTextures[e.currentSampleCount],
			ClearColor: sdl.FColor{R: 1, G: 1, B: 1, A: 1},
			LoadOp:     sdl.GPU_LOADOP_CLEAR,
		}

		if sdl.GPUSampleCount(e.currentSampleCount) == sdl.GPU_SAMPLECOUNT_1 {
			colorTargetInfo.StoreOp = sdl.GPU_STOREOP_STORE
		} else {
			colorTargetInfo.StoreOp = sdl.GPU_STOREOP_RESOLVE
			colorTargetInfo.ResolveTexture = e.resolveTexture
		}

		renderPass := cmdbuf.BeginRenderPass([]sdl.GPUColorTargetInfo{colorTargetInfo}, nil)
		renderPass.BindGraphicsPipeline(e.pipelines[e.currentSampleCount])
		renderPass.DrawPrimitives(3, 1, 0, 0)
		renderPass.End()

		blitSourceTexture := colorTargetInfo.ResolveTexture
		if blitSourceTexture == nil {
			blitSourceTexture = colorTargetInfo.Texture
		}

		cmdbuf.BlitGPUTexture(&sdl.GPUBlitInfo{
			Source: sdl.GPUBlitRegion{
				Texture: blitSourceTexture,
				X:       160,
				W:       320,
				H:       240,
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

func (e *TriangleMSAA) Quit(context *common.Context) {
	for i := range e.sampleCounts {
		context.Device.ReleaseGraphicsPipeline(e.pipelines[i])
		context.Device.ReleaseTexture(e.msaaRenderTextures[i])
	}
	context.Device.ReleaseTexture(e.resolveTexture)

	e.currentSampleCount = 0

	context.Quit()
}
