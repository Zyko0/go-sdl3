// https://examples.libsdl.org/SDL3/renderer/04-points/

package main

import (
	"math/rand/v2"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

const (
	WindowWidth  = 640
	WindowHeight = 480

	NumPoints          = 500
	MinPixelsPerSecond = 30
	MaxPixelsPerSecond = 60
)

var (
	points      [NumPoints]sdl.FPoint
	pointSpeeds [NumPoints]float32
	lastTime    uint64
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/04-points", WindowWidth, WindowHeight, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	/* set up some random points */
	for i := range len(points) {
		points[i].X = (rand.Float32() * WindowWidth) + 100
		points[i].Y = (rand.Float32() * WindowHeight) + 100
		pointSpeeds[i] = MinPixelsPerSecond + rand.Float32()*(MaxPixelsPerSecond-MinPixelsPerSecond)
	}

	lastTime = sdl.Ticks()

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		// Rendering

		now := sdl.Ticks()
		elapsed := float32(now-lastTime) / 1000

		/* let's move all our points a little for a new frame. */
		for i := range len(points) {
			distance := elapsed * pointSpeeds[i]
			points[i].X += distance
			points[i].Y += distance
			if points[i].X >= WindowWidth || points[i].Y >= WindowHeight {
				/* off the screen; restart it elsewhere! */
				if rand.IntN(2) != 0 {
					points[i].X = rand.Float32() * float32(WindowWidth)
					points[i].Y = 0
				} else {
					points[i].X = 0
					points[i].Y = rand.Float32() * float32(WindowHeight)
				}
				pointSpeeds[i] = MinPixelsPerSecond + rand.Float32()*(MaxPixelsPerSecond-MinPixelsPerSecond)
			}
		}

		lastTime = now

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(0, 0, 0, 255)       /* black, full alpha */
		renderer.Clear()                          /* start with a blank canvas. */
		renderer.SetDrawColor(255, 255, 255, 255) /* white, full alpha */
		renderer.RenderPoints(points[:])          /* draw all the points! */

		/* You can also draw single points with SDL_RenderPoint(), but it's
		   cheaper (sometimes significantly so) to do them all at once. */

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
