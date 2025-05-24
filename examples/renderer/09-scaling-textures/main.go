// https://examples.libsdl.org/SDL3/renderer/09-scaling-textures/

package main

import (
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	assets "github.com/Zyko0/go-sdl3/examples/renderer/_assets"
	"github.com/Zyko0/go-sdl3/sdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/09-scaling-textures", WindowWidth, WindowHeight, 0)
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

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		// Rendering

		var dstRect sdl.FRect
		now := sdl.Ticks()

		/* we'll have the texture grow and shrink over a few seconds. */
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}
		scale := (float32(int(now%1000)-500) / 500) * direction

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.Clear()                    /* start with a blank canvas. */

		/* center this one and make it grow and shrink. */
		dstRect.W = float32(texture.W) + float32(texture.W)*scale
		dstRect.H = float32(texture.H) + float32(texture.H)*scale
		dstRect.X = float32(WindowWidth-dstRect.W) / 2
		dstRect.Y = float32(WindowHeight-dstRect.H) / 2

		renderer.RenderTexture(texture, nil, &dstRect)

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
