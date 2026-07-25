//go:build js

package sdl

import "syscall/js"

// WaitAnimationFrame parks the calling goroutine until the browser's next
// requestAnimationFrame callback fires, and returns that callback's
// DOMHighResTimeStamp (milliseconds on the performance.now clock — the time
// the browser started producing the upcoming frame).
//
// This is the browser's VSYNC equivalent: RAF callbacks run once per display
// refresh, right before the compositor paints. Presenting to the canvas and
// then waiting here (a) paces a render loop to the display refresh rate and
// (b) yields to the browser event loop, without which the canvas is never
// composited at all (canvas updates only become visible when the page yields).
//
// Browsers throttle or suspend RAF in background tabs; to avoid blocking
// forever there, a 250 ms setTimeout acts as a fallback tick (returning the
// current performance.now instead of a frame timestamp).
func WaitAnimationFrame() float64 {
	ch := make(chan float64, 1)
	var rafID, timerID js.Value
	raf := js.FuncOf(func(this js.Value, args []js.Value) any {
		js.Global().Call("clearTimeout", timerID)
		ts := 0.0
		if len(args) > 0 {
			ts = args[0].Float()
		}
		ch <- ts
		return nil
	})
	timer := js.FuncOf(func(this js.Value, args []js.Value) any {
		js.Global().Call("cancelAnimationFrame", rafID)
		ch <- js.Global().Get("performance").Call("now").Float()
		return nil
	})
	rafID = js.Global().Call("requestAnimationFrame", raf)
	timerID = js.Global().Call("setTimeout", timer, 250)
	ts := <-ch
	raf.Release()
	timer.Release()
	return ts
}
