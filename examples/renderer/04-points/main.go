// https://examples.libsdl.org/SDL3/renderer/02-primitives/

package main

import (
	"math/rand/v2"
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binary"
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
	defer binary.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)

	runtime.LockOSThread()

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

	lastTime = sdl.GetTicks()

	running := true
	for running {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				running = false
			}
		}

		// Rendering

		now := sdl.GetTicks()
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
		renderer.SetRenderDrawColor(0, 0, 0, 255)       /* black, full alpha */
		renderer.RenderClear()                          /* start with a blank canvas. */
		renderer.SetRenderDrawColor(255, 255, 255, 255) /* white, full alpha */
		renderer.RenderPoints(points[:])                /* draw all the points! */

		/* You can also draw single points with SDL_RenderPoint(), but it's
		   cheaper (sometimes significantly so) to do them all at once. */

		renderer.RenderPresent() /* put it all on the screen! */
	}
}
