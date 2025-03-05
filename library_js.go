//go:build js

package sdl

import (
	"errors"
	"syscall/js"
	"time"

	"github.com/Zyko0/go-sdl3/internal"
)

// SDL_GetError - Retrieve a message about the last error that occurred on the current thread.
// (https://wiki.libsdl.org/SDL3/SDL_GetError)
func init() {
	// Wait for runtime initialization
	// TODO: find a cleaner way than loop+sleep?
	for !js.Global().Get("runtimeInitialized").Bool() {
		time.Sleep(5 * time.Millisecond)
	}

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
