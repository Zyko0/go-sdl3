//go:build js

package sdl

import (
	"syscall/js"

	"github.com/Zyko0/go-sdl3/internal"
)

// SDL_GetKeyboardState - Get a snapshot of the current state of the keyboard.
// (https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState)
//
// SDL returns a pointer to its internal state array (one byte per scancode)
// that lives in the Emscripten heap; copy the snapshot into Go memory.
func GetKeyboardState() []bool {
	internal.StackSave()
	defer internal.StackRestore()
	_count := internal.StackAlloc(4)
	ret := js.Global().Get("Module").Call(
		"_SDL_GetKeyboardState",
		_count,
	)
	count := internal.GetValue(_count, "i32").Int()
	if ret.Int() == 0 || count <= 0 {
		return nil
	}
	raw := internal.GetByteSliceFromJSPtr(ret, count)
	state := make([]bool, count)
	for i, b := range raw {
		state[i] = b != 0
	}

	return state
}
