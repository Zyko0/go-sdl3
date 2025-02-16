// https://examples.libsdl.org/SDL3/renderer/11-color-mods/

package main

import (
	"math"
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

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/11-color-mods", WindowWidth, WindowHeight, 0)
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

		var dstRect sdl.FRect
		now := float64(sdl.GetTicks()) / 1000 /* convert from milliseconds to seconds. */
		/* choose the modulation values for the center texture. The sine wave trick makes it fade between colors smoothly. */
		red := float32(0.5 + 0.5*math.Sin(now))
		green := float32(0.5 + 0.5*math.Sin(now+math.Pi*2/3))
		blue := float32(0.5 + 0.5*math.Sin(now+math.Pi*4/3))

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetRenderDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.RenderClear()                    /* start with a blank canvas. */

		/* Just draw the static texture a few times. You can think of it like a
		   stamp, there isn't a limit to the number of times you can draw with it. */

		/* Color modulation multiplies each pixel's red, green, and blue intensities by the mod values,
		   so multiplying by 1.0f will leave a color intensity alone, 0.0f will shut off that color
		   completely, etc. */

		/* top left; let's make this one blue! */
		dstRect.X = 0
		dstRect.Y = 0
		dstRect.W = float32(texture.W)
		dstRect.H = float32(texture.H)
		texture.SetColorModFloat(red, green, blue)
		renderer.RenderTexture(texture, nil, &dstRect)

		/* center this one, and have it cycle through red/green/blue modulations. */
		dstRect.X = float32(WindowWidth-texture.W) / 2
		dstRect.Y = float32(WindowHeight-texture.H) / 2
		dstRect.W = float32(texture.W)
		dstRect.H = float32(texture.H)
		texture.SetColorModFloat(red, green, blue)
		renderer.RenderTexture(texture, nil, &dstRect)

		/* bottom right; let's make this one red! */
		dstRect.X = float32(WindowWidth - texture.W)
		dstRect.Y = float32(WindowHeight - texture.H)
		dstRect.W = float32(texture.W)
		dstRect.H = float32(texture.H)
		texture.SetColorModFloat(1, 0, 0) /* kill all green and blue. */
		renderer.RenderTexture(texture, nil, &dstRect)

		renderer.RenderPresent() /* put it all on the screen! */
	}
}
