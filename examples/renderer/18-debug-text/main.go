// https://examples.libsdl.org/SDL3/renderer/18-debug-text/

package main

import (
	"fmt"

	"github.com/Zyko0/go-sdl3/sdl"
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

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/renderer/18-debug-text", WindowWidth, WindowHeight, 0)
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

		const charSize = sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE

		/* as you can see from this, rendering draws over whatever was drawn before it. */
		renderer.SetDrawColor(0, 0, 0, 255) /* black, full alpha */
		renderer.Clear()                    /* start with a blank canvas. */

		renderer.SetDrawColor(255, 255, 255, 255) /* white, full alpha */
		renderer.DebugText(272, 100, "Hello world!")
		renderer.DebugText(224, 150, "This is some debug text.")

		renderer.SetDrawColor(51, 102, 255, 255) /* light blue, full alpha */
		renderer.DebugText(184, 200, "You can do it in different colors.")
		renderer.SetDrawColor(255, 255, 255, 255) /* white, full alpha */

		renderer.SetScale(4, 4)
		renderer.DebugText(14, 65, "It can be scaled.")
		renderer.SetScale(1, 1)
		renderer.DebugText(64, 350, "This only does ASCII chars. So this laughing emoji won't draw: 🤣")

		renderer.DebugText(float32(WindowWidth-charSize*46)/2, 400, fmt.Sprintf("(This program has been running for %d seconds.)", sdl.Ticks()/1000))

		renderer.Present() /* put it all on the screen! */

		return nil
	})
}
