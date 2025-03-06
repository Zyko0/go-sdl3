// https://examples.libsdl.org/SDL3/renderer/07-streaming-textures/

package main

import (
	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480

	TextureSize = 150
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/07-streaming-textures", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	texture, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STREAMING, TextureSize, TextureSize)
	if err != nil {
		panic(err)
	}

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
		var surface *sdl.Surface

		/* we'll have some color move around over a few seconds. */
		var direction float32
		if now%2000 >= 1000 {
			direction = 1
		} else {
			direction = -1
		}
		scale := (float32(int(now%1000)-500) / 500) * direction

		/* To update a streaming texture, you need to lock it first. This gets you access to the pixels.
		   Note that this is considered a _write-only_ operation: the buffer you get from locking
		   might not actually have the existing contents of the texture, and you have to write to every
		   locked pixel! */

		/* You can use SDL_LockTexture() to get an array of raw pixels, but we're going to use
		   SDL_LockTextureToSurface() here, because it wraps that array in a temporary SDL_Surface,
		   letting us use the surface drawing functions instead of lighting up individual pixels. */
		err := texture.LockToSurface(nil, &surface)
		if err == nil {
			var r sdl.Rect

			details, _ := surface.Format.Details()
			surface.FillRect(nil, sdl.MapRGB(details, nil, 0, 0, 0)) /* make the whole surface black */
			r.W = TextureSize
			r.H = TextureSize / 10
			r.X = 0
			r.Y = int32(float32(TextureSize-r.H) * (scale + 1) / 2)
			surface.FillRect(&r, sdl.MapRGB(details, nil, 0, 255, 0)) /* make a strip of the surface green */
			texture.Unlock()                                          /* upload the changes (and frees the temporary surface)! */
		}

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(66, 66, 66, 255) /* grey, full alpha */
		renderer.Clear()                       /* start with a blank canvas. */

		/* Just draw the static texture a few times. You can think of it like a
		   stamp, there isn't a limit to the number of times you can draw with it. */

		/* Center this one. It'll draw the latest version of the texture we drew while it was locked. */
		dstRect.X = float32(WindowWidth-TextureSize) / 2
		dstRect.Y = float32(WindowHeight-TextureSize) / 2
		dstRect.W, dstRect.H = TextureSize, TextureSize
		renderer.RenderTexture(texture, nil, &dstRect)

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
