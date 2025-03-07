// https://examples.libsdl.org/SDL3/renderer/15-cliprect/

package main

import (
	"math"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	assets "github.com/Zyko0/go-sdl3/examples/renderer/_assets"
)

const (
	WindowWidth   = 640
	WindowHeight  = 480
	ClipRectSize  = 250
	ClipRectSpeed = 200
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/15-cliprect", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	var clipRectPosition sdl.FPoint

	clipRectDirection := sdl.FPoint{
		X: 1,
		Y: 1,
	}
	lastTime := sdl.Ticks()

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

		cliprect := sdl.Rect{
			X: int32(math.Round(float64(clipRectPosition.X))),
			Y: int32(math.Round(float64(clipRectPosition.Y))),
			W: ClipRectSize,
			H: ClipRectSize,
		}
		now := sdl.Ticks()
		elapsed := float32(now-lastTime) / 1000 /* seconds since last iteration */
		distance := elapsed * ClipRectSpeed

		/* Set a new clipping rectangle position */
		clipRectPosition.X += distance * clipRectDirection.X
		if clipRectPosition.X < 0 {
			clipRectPosition.X = 0
			clipRectDirection.X = 1
		} else if clipRectPosition.X >= (WindowWidth - ClipRectSize) {
			clipRectPosition.X = (WindowWidth - ClipRectSize) - 1
			clipRectDirection.X = -1
		}

		clipRectPosition.Y += distance * clipRectDirection.Y
		if clipRectPosition.Y < 0 {
			clipRectPosition.Y = 0
			clipRectDirection.Y = 1
		} else if clipRectPosition.Y >= (WindowHeight - ClipRectSize) {
			clipRectPosition.Y = (WindowHeight - ClipRectSize) - 1
			clipRectDirection.Y = -1
		}
		renderer.SetClipRect(&cliprect)

		lastTime = now

		/* okay, now draw! */

		/* Note that SDL_RenderClear is _not_ affected by the clipping rectangle! */
		renderer.SetDrawColor(33, 33, 33, 255) /* grey, full alpha */
		renderer.Clear()                       /* start with a blank canvas. */

		/* stretch the texture across the entire window. Only the piece in the
		   clipping rectangle will actually render, though! */
		renderer.RenderTexture(texture, nil, nil)

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
