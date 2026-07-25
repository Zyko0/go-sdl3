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

// RunLoop drives updateFunc at requestAnimationFrame pace, executing it on
// the goroutine that called RunLoop (normally the main goroutine).
//
// It deliberately does NOT use emscripten_set_main_loop: that invokes the
// callback through js.FuncOf, i.e. on a fresh event-handler goroutine, and
// when updateFunc blocks there (experiment-style code sleeps in input-wait
// loops across many frames) the Go scheduler never pauses back to the
// browser event loop — the tab's main thread wedges inside findRunnable.
// Blocking on the main goroutine pauses cleanly, so the RAF callback only
// signals a channel and all user code runs here.
func RunLoop(updateFunc func() error) error {
	frame := make(chan struct{}, 1)
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		select {
		case frame <- struct{}{}:
		default:
		}
		return nil
	})
	defer cb.Release()
	raf := js.Global().Get("requestAnimationFrame")

	for {
		raf.Invoke(cb)
		<-frame
		if err := updateFunc(); err != nil {
			if errors.Is(err, EndLoop) {
				return nil
			}
			return err
		}
	}
}
