//go:build js

package sdl

import (
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
)

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
