// https://examples.libsdl.org/SDL3/renderer/03-lines/

package main

import (
	"image/color"
	"math"
	"math/rand/v2"
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binary"
)

var (
	line_points = []sdl.FPoint{
		{100, 354}, {220, 230}, {140, 230}, {320, 100}, {500, 230},
		{420, 230}, {540, 354}, {400, 354}, {100, 354},
	}
)

func main() {
	defer binary.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)

	runtime.LockOSThread()

	defer sdl.Quit()
	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/03-lines", 640, 480, 0)
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

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetRenderDrawColor(color.RGBA{100, 100, 100, 255}) /* grey, full alpha */
		renderer.RenderClear()                                      /* start with a blank canvas. */

		/* You can draw lines, one at a time, like these brown ones... */
		renderer.SetRenderDrawColor(color.RGBA{127, 49, 32, 255})
		renderer.RenderLine(240, 450, 400, 450)
		renderer.RenderLine(240, 356, 400, 356)
		renderer.RenderLine(240, 356, 240, 450)
		renderer.RenderLine(400, 356, 400, 450)

		/* You can also draw a series of connected lines in a single batch... */
		renderer.SetRenderDrawColor(color.RGBA{0, 255, 0, 255})
		renderer.RenderLines(line_points)

		/* here's a bunch of lines drawn out from a center point in a circle. */
		/* we randomize the color of each line, so it functions as animation. */
		for i := range 360 {
			const (
				size = 30.
				x    = 320.
				y    = 95. - (size / 2.)
			)
			renderer.SetRenderDrawColor(color.RGBA{
				R: uint8(rand.IntN(256)),
				G: uint8(rand.IntN(256)),
				B: uint8(rand.IntN(256)),
				A: 255,
			})
			x2 := x + float32(math.Sin(float64(i))*size)
			y2 := y + float32(math.Cos(float64(i))*size)
			renderer.RenderLine(x, y, x2, y2)
		}

		renderer.RenderPresent() /* put it all on the screen! */
	}
}
