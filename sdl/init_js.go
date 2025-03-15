//go:build js

package sdl

import (
	"errors"
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
)

// We can just initialize everything here in js/wasm env
func init() {
	initialize()

	// Set free, error functions
	internal.SetSDLFreeFunc(func(mem uintptr) {
		ifree(mem)
		internal.DeleteJSPointer(mem)
	})
	internal.SetSDLLastErrFunc(func() error {
		if msg := iGetError(); msg != "" {
			return errors.New(msg)
		}
		return nil
	})
}

// Path returns an empty string in js/wasm environment.
func Path() string {
	return ""
}

// LoadLibrary does nothing in js/wasm environment.
func LoadLibrary(path string) error {
	return nil
}

// CloseLibrary does nothing in js/wasm environment.
func CloseLibrary() error {
	return nil
}

func RunLoop(updateFunc func() error) error {
	ch := make(chan error)
	fn := js.FuncOf(func(this js.Value, args []js.Value) any {
		if err := updateFunc(); err != nil {
			js.Global().Call("_emscripten_cancel_main_loop")
			ch <- err
		}

		return nil
	})
	fnAddr := js.Global().Get("Module").Call("addFunction", fn, "v")
	js.Global().Call("_emscripten_set_main_loop", fnAddr, -1, 0)

	err := <-ch

	if errors.Is(err, EndLoop) {
		return nil
	}
	return err
}
