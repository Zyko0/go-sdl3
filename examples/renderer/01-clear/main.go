// https://examples.libsdl.org/SDL3/renderer/02-primitives/

package main

import (
	"math"
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binary"
)

func main() {
	defer binary.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)

	runtime.LockOSThread()

	defer sdl.Quit()
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/01-clear", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	running := true
	for running {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				running = false
			}
		}

		// Rendering

		now := float64(sdl.GetTicks()) / 1000 /* convert from milliseconds to seconds. */
		/* choose the color for the frame we will draw. The sine wave trick makes it fade between colors smoothly. */
		red := 0.5 + 0.5*float32(math.Sin(now))
		green := 0.5 + 0.5*float32(math.Sin(now+math.Pi*2/3))
		blue := 0.5 + 0.5*float32(math.Sin(now+math.Pi*4/3))
		renderer.SetRenderDrawColorFloat(red, green, blue, 1) /* new color, full alpha. */

		/* clear the window to the draw color. */
		renderer.RenderClear()

		/* put the newly-cleared rendering on the screen. */
		renderer.RenderPresent()
	}
}
