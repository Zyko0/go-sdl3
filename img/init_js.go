//go:build js

package img

import (
	"syscall/js"
	"time"
)

// We can just initialize everything here in js/wasm env
func init() {
	// Wait for runtime initialization
	// TODO: find a cleaner way than loop+sleep?
	for !js.Global().Get("runtimeInitialized").Bool() {
		time.Sleep(5 * time.Millisecond)
	}

	initialize()
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
