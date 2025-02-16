// https://examples.libsdl.org/SDL3/renderer/17-read-pixels/

package main

import (
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
	assets "github.com/Zyko0/go-sdl3/examples/renderer/_assets"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)

	runtime.LockOSThread()

	defer sdl.Quit()
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/17-read-pixels", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	var convertedTexture *sdl.Texture
	var convertedTextureWidth, convertedTextureHeight int32

	/* Textures are pixel data that we upload to the video hardware for fast drawing. Lots of 2D
	   engines refer to these as "sprites." We'll do a static texture (upload once, draw many
	   times) with data from a bitmap file. */
	bmpStream, err := sdl.IOFromConstMem(assets.SampleBMP)
	if err != nil {
		panic(err)
	}

	/* SDL_Surface is pixel data the CPU can access. SDL_Texture is pixel data the GPU can access.
	   Load a .bmp into a surface, move it to a texture from there. */
	surface, err := sdl.LoadBMP_IO(bmpStream, true)
	if err != nil {
		panic(err)
	}

	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	surface.Destroy()

	running := true
	for running {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				running = false
			}
		}

		// Rendering

		now := sdl.Ticks()
		var center sdl.FPoint
		var dstRect sdl.FRect

		/* we'll have a texture rotate around over 2 seconds (2000 milliseconds). 360 degrees in a circle! */
		rotation := float64(now%2000) / 2000 * 360

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.Clear()                    /* start with a blank canvas. */

		/* Center this one, and draw it with some rotation so it spins! */
		dstRect.X = float32(WindowWidth-texture.W) / 2
		dstRect.Y = float32(WindowHeight-texture.H) / 2
		dstRect.W = float32(texture.W)
		dstRect.H = float32(texture.H)
		/* rotate it around the center of the texture; you can rotate it from a different point, too! */
		center.X = float32(texture.W) / 2
		center.Y = float32(texture.H) / 2
		renderer.RenderTextureRotated(texture, nil, &dstRect, rotation, &center, sdl.FLIP_NONE)

		/* this next whole thing is _super_ expensive. Seriously, don't do this in real life. */

		/* Download the pixels of what has just been rendered. This has to wait for the GPU to finish rendering it and everything before it,
		   and then make an expensive copy from the GPU to system RAM! */
		surface, _ = renderer.ReadPixels(nil)

		/* This is also expensive, but easier: convert the pixels to a format we want. */
		if surface != nil && surface.Format != sdl.PIXELFORMAT_RGBA8888 && surface.Format != sdl.PIXELFORMAT_BGRA8888 {
			converted, _ := surface.Convert(sdl.PIXELFORMAT_RGBA8888)
			surface.Destroy()
			surface = converted
		}

		if surface != nil {
			/* Rebuild converted_texture if the dimensions have changed (window resized, etc). */
			if surface.W != convertedTextureWidth || surface.H != convertedTextureHeight {
				convertedTexture.Destroy()
				convertedTexture, err = renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, int(surface.W), int(surface.H))
				if err != nil {
					panic(err)
				}
				convertedTextureWidth = surface.W
				convertedTextureHeight = surface.H
			}

			/* Turn each pixel into either black or white. This is a lousy technique but it works here.
			   In real life, something like Floyd-Steinberg dithering might work
			   better: https://en.wikipedia.org/wiki/Floyd%E2%80%93Steinberg_dithering*/
			for y := range surface.H {
				pixels := surface.Pixels()[y*surface.Pitch:]
				for x := range surface.W {
					p := pixels[x*4 : x*4+4]
					average := (uint32(p[1]) + uint32(p[2]) + uint32(p[3])) / 3
					if average == 0 {
						/* make pure black pixels red. */
						p[0] = 0xFF
						p[1] = 0
						p[2] = 0
						p[3] = 0xFF
					} else {
						/* make everything else either black or white. */
						var v byte
						if average > 50 {
							v = 0xFF
						}
						p[1] = v
						p[2] = v
						p[3] = v
					}
				}
			}

			/* upload the processed pixels back into a texture. */
			convertedTexture.Update(nil, surface.Pixels(), surface.Pitch)
			surface.Destroy()

			/* draw the texture to the top-left of the screen. */
			dstRect.X, dstRect.Y = 0, 0
			dstRect.W = float32(WindowWidth) / 4
			dstRect.H = float32(WindowHeight) / 4
			renderer.RenderTexture(convertedTexture, nil, &dstRect)
		}

		renderer.Present() /* put it all on the screen! */
	}
}
