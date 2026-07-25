// Verification harness for the WebAssembly (GOOS=js/GOARCH=wasm) path.
//
// It exercises the SDL functions that were un-stubbed for js: creating a
// window and a renderer as separate calls (CreateWindow + Window.CreateRenderer),
// reading back the render output size, drawing points/primitives, and creating a
// render-target texture with blend/scale/color/alpha modulation.
//
// Run natively:  go run ./examples/wasmcheck
// Run in browser: go run ./cmd/wasmsdl serve ./examples/wasmcheck
//
//	(then open the printed URL and watch the JS console)
package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/sdl"
)

var points = [256]sdl.FPoint{}

func main() {
	defer binsdl.Load().Unload() // no-op on js; loads the native lib otherwise
	defer sdl.Quit()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	// CreateWindow + CreateRenderer as separate calls (not the combined helper).
	window, err := sdl.CreateWindow("examples/wasmcheck", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	renderer, err := window.CreateRenderer("")
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// Read-back out-parameter functions.
	if ww, wh, err := window.Size(); err == nil {
		fmt.Printf("window size: %dx%d\n", ww, wh)
	}
	ow, oh, err := renderer.RenderOutputSize()
	if err != nil {
		panic(err)
	}
	fmt.Printf("render output size (before loop): %dx%d\n", ow, oh)

	// A render-target texture exercises CreateTexture + the texture modulation setters.
	tex, err := renderer.CreateTexture(
		sdl.PIXELFORMAT_RGBA8888,
		sdl.TEXTUREACCESS_TARGET,
		64, 64,
	)
	if err != nil {
		panic(err)
	}
	defer tex.Destroy()
	if err := tex.SetBlendMode(sdl.BLENDMODE_BLEND); err != nil {
		panic(err)
	}
	if err := tex.SetScaleMode(sdl.SCALEMODE_NEAREST); err != nil {
		panic(err)
	}
	if err := tex.SetColorMod(255, 128, 64); err != nil {
		panic(err)
	}
	if err := tex.SetAlphaMod(200); err != nil {
		panic(err)
	}
	tw, th, err := tex.Size()
	if err != nil {
		panic(err)
	}
	fmt.Printf("texture size: %vx%v\n", tw, th)

	for i := range points {
		points[i].X = rand.Float32() * 640
		points[i].Y = rand.Float32() * 480
	}

	logged := false
	sdl.RunLoop(func() error {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		if !logged {
			logged = true
			if w, h, err := renderer.RenderOutputSize(); err == nil {
				fmt.Printf("render output size (in loop): %dx%d\n", w, h)
			}
		}

		renderer.SetDrawColor(20, 20, 30, 255)
		renderer.Clear()

		// RenderPoint (single, float args) + RenderPoints (array).
		renderer.SetDrawColor(255, 0, 0, 255)
		for i := range points {
			renderer.RenderPoint(points[i].X, points[i].Y)
		}

		var rect sdl.FRect
		rect.X, rect.Y, rect.W, rect.H = 220, 140, 200, 200
		renderer.SetDrawColor(0, 200, 80, 255)
		renderer.RenderRect(&rect)

		renderer.Present()
		return nil
	})
}
