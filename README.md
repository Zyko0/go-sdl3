# go-sdl3

[![Go Reference](https://pkg.go.dev/badge/github.com/zyko0/go-sdl3.svg)](https://pkg.go.dev/github.com/zyko0/go-sdl3)

[SDL3](https://wiki.libsdl.org/SDL3/FrontPage) bindings for Go in pure Go (thanks to [ebitengine/purego](https://github.com/ebitengine/purego)).

# Status

- Only tested on Windows at the moment.
- A lot of methods from [methods.go](methods.go) still `panic`, because they need to be refactored as idiomatic Go, and checked against documentation (necessity of `SDL_free()`, `error` return for example).
- Other functions might simply not be exposed, but most likely [generated](sdl_functions.gen_impl.go).
- Help is appreciated! (see [CONTRIBUTE.md](CONTRIBUTE.md).)

# Usage

```go
package main

import (
	"runtime"

	sdl "github.com/Zyko0/go-sdl3"
)

func main() {
	sdl.LoadLibrary("SDL3.dll")

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
```