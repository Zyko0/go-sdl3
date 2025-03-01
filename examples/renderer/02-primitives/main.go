// https://examples.libsdl.org/SDL3/renderer/02-primitives/

package main

import (
	"math/rand/v2"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

var (
	points = [500]sdl.FPoint{}
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/02-primitives", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	/* set up some random points */
	for i := range len(points) {
		points[i].X = (rand.Float32() * 440) + 100
		points[i].Y = (rand.Float32() * 280) + 100
	}

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		// Rendering

		var rect sdl.FRect

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(33, 33, 33, 255) /* dark gray, full alpha */
		renderer.Clear()                       /* start with a blank canvas. */

		/* draw a filled rectangle in the middle of the canvas. */
		renderer.SetDrawColor(0, 0, 255, 255) /* blue, full alpha */
		rect.X, rect.Y = 100, 100
		rect.W = 440
		rect.H = 280
		renderer.RenderFillRect(&rect)

		/* draw some points across the canvas. */
		renderer.SetDrawColor(255, 0, 0, 255) /* red, full alpha */
		renderer.RenderPoints(points[:])

		/* draw a unfilled rectangle in-set a little bit. */
		renderer.SetDrawColor(0, 255, 0, 255) /* green, full alpha */
		rect.X += 30
		rect.Y += 30
		rect.W -= 60
		rect.H -= 60
		renderer.RenderRect(&rect)

		/* draw two lines in an X across the whole canvas. */
		renderer.SetDrawColor(255, 255, 0, 255) /* yellow, full alpha */
		renderer.RenderLine(0, 0, 640, 480)
		renderer.RenderLine(0, 480, 640, 0)

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
