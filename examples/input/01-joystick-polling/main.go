// https://examples.libsdl.org/SDL3/input/01-joystick-polling/

package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"os"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
)

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(sdl.Path())
	defer sdl.Quit()
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_JOYSTICK)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/input/01-joystick-polling", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	var joystick *sdl.Joystick

	var colors [64]sdl.Color
	for i := range colors {
		colors[i].R = uint8(rand.IntN(255))
		colors[i].G = uint8(rand.IntN(255))
		colors[i].B = uint8(rand.IntN(255))
		colors[i].A = 255
	}

	sdl.RunLoop(func() error {
		var event sdl.Event
		var err error

		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_QUIT:
				return sdl.EndLoop
			case sdl.EVENT_JOYSTICK_ADDED:
				/* this event is sent for each hotplugged stick, but also each already-connected joystick during SDL_Init(). */
				evt := event.JoyDeviceEvent()
				if joystick == nil { /* we don't have a stick yet and one was added, open it! */
					joystick, err = evt.Which.OpenJoystick()
					if err != nil {
						fmt.Fprintf(os.Stderr, "failed to open joystick ID %d: %s\n", evt.Which, err.Error())
					}
				}
			case sdl.EVENT_JOYSTICK_REMOVED:
				evt := event.JoyDeviceEvent()
				if joystick != nil {
					id, _ := joystick.ID()
					if id == evt.Which {
						joystick.Close() /* our joystick was unplugged. */
						joystick = nil
					}
				}
			}
		}

		// Rendering

		text := "Plug in a joystick, please."
		if joystick != nil { /* we have a stick opened? */
			text, _ = joystick.Name()
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		winw, winh, _ := window.Size()

		/* note that you can get input as events, instead of polling, which is
		   better since it won't miss button presses if the system is lagging,
		   but often times checking the current state per-frame is good enough,
		   and maybe better if you'd rather _drop_ inputs due to lag. */

		if joystick != nil { /* we have a stick opened? */
			size := float32(30)

			/* draw axes as bars going across middle of screen. We don't know if it's an X or Y or whatever axis, so we can't do more than this. */
			total, _ := joystick.NumAxes()
			y := (float32(winh) - (float32(total) * size)) / 2
			x := float32(winw) / 2
			for i := 0; i < int(total); i++ {
				color := colors[i%len(colors)]
				axis, _ := joystick.Axis(int32(i))
				dx := x + (float32(axis) / 32767 * x) /* make it -1.0f to 1.0f */
				dst := sdl.FRect{
					X: dx,
					Y: y,
					W: x - float32(math.Abs(float64(dx))),
					H: size,
				}
				renderer.SetDrawColor(color.R, color.G, color.B, color.A)
				renderer.RenderFillRect(&dst)
				y += size
			}

			/* draw buttons as blocks across top of window. We only know the button numbers, but not where they are on the device. */
			total, _ = joystick.NumButtons()
			x = float32((float32(winw) - (float32(total) * size)) / 2)
			for i := 0; i < int(total); i++ {
				color := colors[i%len(colors)]
				dst := sdl.FRect{
					X: x,
					Y: 0,
					W: size,
					H: size,
				}
				if joystick.Button(int32(i)) {
					renderer.SetDrawColor(color.R, color.G, color.B, color.A)
				} else {
					renderer.SetDrawColor(0, 0, 0, 255)
				}
				renderer.RenderFillRect(&dst)
				renderer.SetDrawColor(255, 255, 255, color.A)
				renderer.RenderRect(&dst) /* outline it */
				x += size
			}

			/* draw hats across the bottom of the screen. */
			total, _ = joystick.NumHats()
			x = (float32(winw) - (float32(total)*(size*2))/2) + size/2
			y = float32(winh) - size
			for i := 0; i < int(total); i++ {
				color := colors[i%len(colors)]
				thirdsize := size / 3
				cross := []sdl.FRect{
					{
						X: x,
						Y: y + float32(thirdsize),
						W: float32(size),
						H: float32(thirdsize),
					},
					{
						X: x + float32(thirdsize),
						Y: y,
						W: float32(thirdsize),
						H: float32(size),
					},
				}
				hat := joystick.Hat(int32(i))

				renderer.SetDrawColor(90, 90, 90, 255)
				renderer.RenderFillRects(cross)

				renderer.SetDrawColor(color.R, color.G, color.B, color.A)

				if hat&sdl.HAT_UP != 0 {
					dst := sdl.FRect{
						X: x + float32(thirdsize),
						Y: y,
						W: float32(thirdsize),
						H: float32(thirdsize),
					}
					renderer.RenderFillRect(&dst)
				}
				if hat&sdl.HAT_RIGHT != 0 {
					dst := sdl.FRect{
						X: x + float32(thirdsize*2),
						Y: y + float32(thirdsize),
						W: float32(thirdsize),
						H: float32(thirdsize),
					}
					renderer.RenderFillRect(&dst)
				}

				if hat&sdl.HAT_DOWN != 0 {
					dst := sdl.FRect{
						X: x + float32(thirdsize),
						Y: y + float32(thirdsize*2),
						W: float32(thirdsize),
						H: float32(thirdsize),
					}
					renderer.RenderFillRect(&dst)
				}

				if hat&sdl.HAT_LEFT != 0 {
					dst := sdl.FRect{
						X: x,
						Y: y + float32(thirdsize),
						W: float32(thirdsize),
						H: float32(thirdsize),
					}
					renderer.RenderFillRect(&dst)
				}

				x += size * 2
			}
		}

		x := (float32(winw) - float32(len(text)*sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE)) / 2
		y := (float32(winh) - sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE) / 2
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DebugText(x, y, text)

		/* put the newly-cleared rendering on the screen. */
		renderer.Present()

		return nil
	})
}
