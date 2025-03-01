package main

import (
	"errors"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)
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

	renderer.SetDrawColor(255, 255, 255, 255)

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return errors.New("quit")
			}
		}

		renderer.DebugText(50, 50, "Hello world")
		renderer.Present()

		return nil
	})
}
