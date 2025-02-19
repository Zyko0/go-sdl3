// https://examples.libsdl.org/SDL3/renderer/14-viewport/

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

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/14-viewport", WindowWidth, WindowHeight, 0)
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

		var viewport sdl.Rect
		dstRect := sdl.FRect{
			X: 0,
			Y: 0,
			W: float32(texture.W),
			H: float32(texture.H),
		}

		/* Setting a viewport has the effect of limiting the area that rendering
		   can happen, and making coordinate (0, 0) live somewhere else in the
		   window. It does _not_ scale rendering to fit the viewport. */

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.Clear()                    /* start with a blank canvas. */

		/* Draw once with the whole window as the viewport. */
		viewport.X = 0
		viewport.Y = 0
		viewport.W = WindowWidth / 2
		viewport.H = WindowHeight / 2
		renderer.SetViewport(nil) /* NULL means "use the whole window" */
		renderer.RenderTexture(texture, nil, &dstRect)

		/* top right quarter of the window. */
		viewport.X = WindowWidth / 2
		viewport.Y = WindowHeight / 2
		viewport.W = WindowWidth / 2
		viewport.H = WindowHeight / 2
		renderer.SetViewport(&viewport)
		renderer.RenderTexture(texture, nil, &dstRect)

		/* bottom 20% of the window. Note it clips the width! */
		viewport.X = 0
		viewport.Y = WindowHeight - (WindowHeight / 5)
		viewport.W = WindowWidth / 5
		viewport.H = WindowHeight / 5
		renderer.SetViewport(&viewport)
		renderer.RenderTexture(texture, nil, &dstRect)

		/* what happens if you try to draw above the viewport? It should clip! */
		viewport.X = 100
		viewport.Y = 200
		viewport.W = WindowWidth
		viewport.H = WindowHeight
		renderer.SetViewport(&viewport)
		dstRect.Y = -50
		renderer.RenderTexture(texture, nil, &dstRect)

		renderer.Present() /* put it all on the screen! */
	}
}
