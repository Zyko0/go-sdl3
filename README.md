# go-sdl3

[![Go Reference](https://pkg.go.dev/badge/github.com/zyko0/go-sdl3.svg)](https://pkg.go.dev/github.com/zyko0/go-sdl3)
[![Go Report Card](https://goreportcard.com/badge/github.com/Zyko0/go-sdl3)](https://goreportcard.com/report/github.com/Zyko0/go-sdl3)

[SDL3](https://wiki.libsdl.org/SDL3/FrontPage) bindings for Go in pure Go (thanks to [ebitengine/purego](https://github.com/ebitengine/purego)).

## About

The API is currently not stable, I might make some renames and packages refactor.

This library wraps SDL3 to more idiomatic go and:
- Changes return values from `bool` to `error` (`SDL_GetError`).
- Trims `SDL_` prefix from all types, variables, function names.
- Make methods from global functions when it is possible.
- Turn some pointer function parameters into return values.

If you are looking for pure Go bindings that are closer to the original API, please have a look at https://github.com/JupiterRider/purego-sdl3. 

## Status

> [!NOTE]
> The API is currently subject to many changes / refactors, many method functions are also exposed but not implemented yet (need human validation).

The list of currently supported functions can be found in [COVERAGE.md](COVERAGE.md).

- Tested on Windows, Ubuntu, WebAssembly (wasm/js) at the moment.
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

**Example:**
> Note that you do not have to pass your update function `sdl.RunLoop`, however doing so allows you to target `GOOS=js`/`GOARCH=wasm`, see [wasmsdl](cmd/wasmsdl/). 
```go
package main

import (
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
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
				return sdl.EndLoop
			}
		}

		renderer.DebugText(50, 50, "Hello world")
		renderer.Present()

		return nil
	})
}
```

## Examples

The [examples](./examples/) folder contains the offical examples that can be found here https://examples.libsdl.org/SDL3, and a few more.

## Libraries

- SDL3
- SDL3_ttf
- SDL3_image
- SDL3_mixer
- SDL3_shadercross (TBD)