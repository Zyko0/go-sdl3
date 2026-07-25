// Verification harness for the SDL3_image and SDL3_ttf bindings on the
// WebAssembly (GOOS=js/GOARCH=wasm) path.
//
// Assets are passed to SDL as in-memory IOStreams (sdl.IOFromConstMem), which
// works identically on native and in the browser - no Emscripten virtual
// filesystem or preloading required. The image is generated at runtime; the
// font is embedded.
//
// Run natively:  go run ./examples/imgttf
// Run in browser: go run ./cmd/wasmsdl serve ./examples/imgttf
package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

//go:embed assets/Go-Medium.ttf
var fontData []byte

// makePNG builds a small RGBA gradient image and encodes it as PNG bytes.
func makePNG(w, h int) []byte {
	im := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			im.Set(x, y, color.RGBA{
				R: uint8(x * 255 / w),
				G: uint8(y * 255 / h),
				B: 128,
				A: 255,
			})
		}
	}
	var buf bytes.Buffer
	if err := png.Encode(&buf, im); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func loadTextureFromBytes(renderer *sdl.Renderer, data []byte) *sdl.Texture {
	stream, err := sdl.IOFromConstMem(data)
	if err != nil {
		panic(err)
	}
	// closeio=true: SDL closes the stream once the image is decoded.
	tex, err := img.LoadTextureIO(renderer, stream, true)
	if err != nil {
		panic(err)
	}
	return tex
}

func main() {
	defer binsdl.Load().Unload()
	defer sdl.Quit()

	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/imgttf", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	// --- SDL3_image: decode an in-memory PNG into a texture ---
	imgTex := loadTextureFromBytes(renderer, makePNG(256, 256))
	defer imgTex.Destroy()
	iw, ih, _ := imgTex.Size()
	fmt.Printf("image texture: %vx%v\n", iw, ih)

	// --- SDL3_ttf: open the embedded font and render text into a texture ---
	if err := ttf.Init(); err != nil {
		panic(err)
	}
	defer ttf.Quit()

	fontStream, err := sdl.IOFromConstMem(fontData)
	if err != nil {
		panic(err)
	}
	font, err := ttf.OpenFontIO(fontStream, true, 48)
	if err != nil {
		panic(err)
	}
	defer font.Close()

	// Blended (single color).
	surface, err := font.RenderTextBlended("Hello SDL3_ttf!", sdl.Color{R: 255, G: 255, B: 0, A: 255})
	if err != nil {
		panic(err)
	}
	textTex, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	surface.Destroy()
	defer textTex.Destroy()
	tw, th, _ := textTex.Size()
	fmt.Printf("text texture: %vx%v\n", tw, th)

	// Shaded (two colors: foreground + background) - exercises the bg color path.
	shadedSurf, err := font.RenderTextShaded(
		"Shaded text",
		sdl.Color{R: 0, G: 0, B: 0, A: 255},
		sdl.Color{R: 0, G: 200, B: 255, A: 255},
	)
	if err != nil {
		panic(err)
	}
	shadedTex, err := renderer.CreateTextureFromSurface(shadedSurf)
	if err != nil {
		panic(err)
	}
	shadedSurf.Destroy()
	defer shadedTex.Destroy()
	sw, sh, _ := shadedTex.Size()

	sdl.RunLoop(func() error {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			if event.Type == sdl.EVENT_QUIT {
				return sdl.EndLoop
			}
		}

		renderer.SetDrawColor(30, 30, 40, 255)
		renderer.Clear()

		// Blit the decoded image.
		imgDst := sdl.FRect{X: 40, Y: 40, W: iw, H: ih}
		renderer.RenderTexture(imgTex, nil, &imgDst)

		// Blit the blended text below it.
		textDst := sdl.FRect{X: 40, Y: 330, W: tw, H: th}
		renderer.RenderTexture(textTex, nil, &textDst)

		// Blit the shaded (two-color) text below that.
		shadedDst := sdl.FRect{X: 40, Y: 400, W: sw, H: sh}
		renderer.RenderTexture(shadedTex, nil, &shadedDst)

		renderer.Present()
		return nil
	})
}
