// https://examples.libsdl.org/SDL3/renderer/18-debug-text/

package main

import (
	"fmt"
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
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

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/18-debug-text", WindowWidth, WindowHeight, 0)
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

		const charSize = sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetRenderDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.RenderClear()                    /* start with a blank canvas. */

		renderer.SetRenderDrawColor(255, 255, 255, 255) /* white, full alpha */
		renderer.RenderDebugText(272, 100, "Hello world!")
		renderer.RenderDebugText(224, 150, "This is some debug text.")

		renderer.SetRenderDrawColor(51, 102, 255, 255) /* light blue, full alpha */
		renderer.RenderDebugText(184, 200, "You can do it in different colors.")
		renderer.SetRenderDrawColor(255, 255, 255, 255) /* white, full alpha */

		renderer.SetRenderScale(4, 4)
		renderer.RenderDebugText(14, 65, "It can be scaled.")
		renderer.SetRenderScale(1, 1)
		renderer.RenderDebugText(64, 350, "This only does ASCII chars. So this laughing emoji won't draw: 🤣")

		renderer.RenderDebugText(float32(WindowWidth-charSize*46)/2, 400, fmt.Sprintf("(This program has been running for %d seconds.)", sdl.GetTicks()/1000))

		renderer.RenderPresent() /* put it all on the screen! */
	}
}
