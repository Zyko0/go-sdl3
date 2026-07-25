//go:build js

package sdl

import (
	"strconv"
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
)

// sizeCanvas sets the CSS pixel size of the Emscripten canvas. SDL's web
// backend derives the window size from the canvas CSS layout size, so this must
// run before SDL_CreateWindow* for the requested dimensions to take effect.
// Without it the canvas can collapse to 0x0 and nothing is rendered.
func sizeCanvas(w, h int32) {
	canvas := js.Global().Get("Module").Get("canvas")
	if !canvas.Truthy() {
		return
	}
	style := canvas.Get("style")
	if !style.Truthy() {
		return
	}
	style.Set("width", strconv.Itoa(int(w))+"px")
	style.Set("height", strconv.Itoa(int(h))+"px")
}

func (s *Surface) Pixels() []byte {
	return internal.GetByteSliceFromJSPtr(js.ValueOf(s.pixels), int(s.H*s.Pitch))
}

// Callbacks

func NewCleanupPropertyCallback(fn func(value uintptr)) CleanupPropertyCallback {
	panic("not implemented in js/wasm environment")
}

func NewEnumeratePropertiesCallback(fn func(props PropertiesID, name string)) EnumeratePropertiesCallback {
	jsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		// Userdata is at index 0
		props := PropertiesID(args[1].Int())
		name := internal.UTF8JSToString(args[2])
		fn(props, name)

		return nil
	})
	fnAddr := js.Global().Get("Module").Call("addFunction", jsFunc, "vpip")
	return EnumeratePropertiesCallback(fnAddr.Int())
}
