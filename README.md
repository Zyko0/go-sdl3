# go-sdl3

[![Go Reference](https://pkg.go.dev/badge/github.com/zyko0/go-sdl3.svg)](https://pkg.go.dev/github.com/zyko0/go-sdl3)

[SDL3](https://wiki.libsdl.org/SDL3/FrontPage) bindings for Go in pure Go (thanks to [ebitengine/purego](https://github.com/ebitengine/purego)).

## Status

> [!NOTE]
> The API is currently subject to many changes / refactors, many method functions are also exposed but not implemented yet (need human validation).

- Tested on Windows and Ubuntu at the moment.
- A lot of methods from [methods.go](methods.go) still `panic`, because they need to be refactored as idiomatic Go, and checked against documentation (necessity of `SDL_free()`, `error` return for example).
- Other functions might simply not be exposed, but most likely [generated](sdl_functions.gen_impl.go).
- Help is appreciated! (see [CONTRIBUTE.md](CONTRIBUTE.md).)

## Usage

The library is linked dynamically with [purego](https://github.com/ebitengine/purego) (does not require CGo).
- Automatic

If the platform is supported (if a binary exists for it), the code below will write the library to a temporary folder, and remove it when the `main` function returns.
```go
defer binsdl.Load().Unload()
```
- Manual

The binary is already installed and/or its location is known (e.g: same folder), so it can be loaded by its path.
```go
sdl.LoadLibrary("SDL3.dll") // "libSDL3.so", "libSDL3.dylib"
```

Example:
```go
package main

import (
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

	running := true
	for running {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				running = false
			}
		}

		renderer.DebugText(50, 50, "Hello world")
		renderer.Present()
	}
}
```

## Examples

The [examples](./examples/) folder contains the offical examples that can be found here https://examples.libsdl.org/SDL3, and a few more.

## Libraries

- SDL3
- SDL3_ttf (TBD)
- SDL3_mixer (TBD)
- SDL3_shadercross (TBD)