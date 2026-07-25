//go:build js

package sdl

import (
	"errors"
	"syscall/js"
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
)

// SDL_LoadWAV_IO - Load the audio data of a WAVE file into memory.
// (https://wiki.libsdl.org/SDL3/SDL_LoadWAV_IO)
//
// SDL allocates the decoded buffer in the Emscripten heap, so we copy the bytes
// into Go memory and free the SDL buffer rather than handing back a raw pointer.
func LoadWAV_IO(src *IOStream, closeIO bool, spec *AudioSpec) ([]byte, error) {
	_src, ok := internal.GetJSPointer(src)
	if !ok {
		panic("nil stream")
	}

	internal.StackSave()
	defer internal.StackRestore()
	_spec := internal.StackAlloc(int(unsafe.Sizeof(*spec)))
	_buf := internal.StackAlloc(4)
	_len := internal.StackAlloc(4)

	ret := js.Global().Get("Module").Call(
		"_SDL_LoadWAV_IO",
		_src,
		internal.NewBoolean(closeIO),
		_spec,
		_buf,
		_len,
	)
	if !internal.GetBool(ret) {
		return nil, internal.LastErr()
	}

	internal.CopyJSToObject(spec, _spec)
	bufAddr := internal.GetValue(_buf, "*")
	count := int(internal.GetValue(_len, "i32").Int())
	data := internal.GetByteSliceFromJSPtr(bufAddr, count)
	js.Global().Get("Module").Call("_SDL_free", bufAddr)

	return data, nil
}

// SDL_LoadWAV - Loads a WAV from a file path.
//
// Not supported on js/wasm: there is no real filesystem in the browser. Load the
// bytes yourself (e.g. go:embed or fetch) and use LoadWAV_IO with IOFromConstMem.
func LoadWAV(path string, spec *AudioSpec) ([]byte, error) {
	return nil, errors.New("sdl.LoadWAV is not supported on js/wasm; use LoadWAV_IO with sdl.IOFromConstMem")
}
