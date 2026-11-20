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

// RunLoop calls updateFunc once per browser animation frame, on the calling
// goroutine. A js.FuncOf callback must not block, as it holds the browser
// event loop: the frame callback only wakes this goroutine, so updateFunc may
// block (net/http, channels, time.Sleep).
func RunLoop(updateFunc func() error) error {
	frame := make(chan struct{}, 1)
	onFrame := js.FuncOf(func(this js.Value, args []js.Value) any {
		frame <- struct{}{}
		return nil
	})
	defer onFrame.Release()

	requestAnimationFrame := js.Global().Get("requestAnimationFrame")
	for {
		requestAnimationFrame.Invoke(onFrame)
		<-frame
		if err := updateFunc(); err != nil {
			if errors.Is(err, EndLoop) {
				return nil
			}
			return err
		}
	}
}
