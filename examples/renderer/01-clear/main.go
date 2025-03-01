// https://examples.libsdl.org/SDL3/renderer/01-clear/

package main

import (
	"math"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)
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

	sdl.RunLoop(func () error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		// Rendering

		now := float64(sdl.Ticks()) / 1000 /* convert from milliseconds to seconds. */
		/* choose the color for the frame we will draw. The sine wave trick makes it fade between colors smoothly. */
		red := 0.5 + 0.5*float32(math.Sin(now))
		green := 0.5 + 0.5*float32(math.Sin(now+math.Pi*2/3))
		blue := 0.5 + 0.5*float32(math.Sin(now+math.Pi*4/3))
		renderer.SetDrawColorFloat(red, green, blue, 1) /* new color, full alpha. */

		/* clear the window to the draw color. */
		renderer.Clear()

		/* put the newly-cleared rendering on the screen. */
		renderer.Present()

		return nil
	})
}
