package main

import (
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binary"
)

func main() {
	defer binary.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)

	runtime.LockOSThread()

	if err := sdl.SetHint("SDL_RENDER_VSYNC", "1"); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("Hello world", 500, 500, 0)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	defer window.Destroy()

	err = renderer.SetRenderDrawColor(255, 255, 255, 255)
	if err != nil {
		panic(err)
	}

	running := true
	for running {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				running = false
			}
		}

		renderer.RenderDebugText(50, 50, "Hello world")
		renderer.RenderPresent()
	}
}
