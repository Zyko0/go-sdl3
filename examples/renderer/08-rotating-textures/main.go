// https://examples.libsdl.org/SDL3/renderer/08-rotating-textures/

package main

import (
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
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/08-rotating-textures", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

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

		var center sdl.FPoint
		var dstRect sdl.FRect
		now := sdl.Ticks()

		/* we'll have a texture rotate around over 2 seconds (2000 milliseconds). 360 degrees in a circle! */
		rotation := float32(now%2000) / 2000 * 360

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

		renderer.RenderTextureRotated(texture, nil, &dstRect, float64(rotation), &center, sdl.FLIP_NONE)

		renderer.Present() /* put it all on the screen! */
	}
}
