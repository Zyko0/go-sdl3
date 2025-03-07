// https://examples.libsdl.org/SDL3/renderer/05-rectangles/

package main

import (
	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
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

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/05-rectangles", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		// Rendering

		var rects [16]sdl.FRect
		now := sdl.Ticks()

		/* we'll have the rectangles grow and shrink over a few seconds. */
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

		/* Rectangles are comprised of set of X and Y coordinates, plus width and
		   height. (0, 0) is the top left of the window, and larger numbers go
		   down and to the right. This isn't how geometry works, but this is
		   pretty standard in 2D graphics. */

		/* Let's draw a single rectangle (square, really). */
		rects[0].X, rects[0].Y = 100, 100
		rects[0].W, rects[0].H = 100+(100*scale), 100+(100*scale)
		renderer.SetDrawColor(255, 0, 0, 255) /* red, full alpha */
		renderer.RenderRect(&rects[0])

		/* Now let's draw several rectangles with one function call. */
		for i := 0; i < 3; i++ {
			size := float32(i+1) * 50
			rects[i].W, rects[i].H = size+(size*scale), size+(size*scale)
			rects[i].X = (WindowWidth - rects[i].W) / 2  /* center it. */
			rects[i].Y = (WindowHeight - rects[i].H) / 2 /* center it. */
		}
		renderer.SetDrawColor(0, 255, 0, 255) /* green, full alpha */
		renderer.RenderRects(rects[:3])       /* draw three rectangles at once */

		/* those were rectangle _outlines_, really. You can also draw _filled_ rectangles! */
		rects[0].X = 400
		rects[0].Y = 50
		rects[0].W = 100 + (100 * scale)
		rects[0].H = 50 + (50 * scale)
		renderer.SetDrawColor(0, 0, 255, 255) /* blue, full alpha */
		renderer.RenderFillRect(&rects[0])

		/* ...and also fill a bunch of rectangles at once... */
		for i := 0; i < len(rects); i++ {
			w := float32(WindowWidth / len(rects))
			h := float32(i * 8)
			rects[i].X = float32(i) * w
			rects[i].Y = WindowHeight - h
			rects[i].W = w
			rects[i].H = h
		}
		renderer.SetDrawColor(255, 255, 255, 255) /* white, full alpha */
		renderer.RenderFillRects(rects[:])

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
