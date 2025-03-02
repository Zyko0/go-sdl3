// https://examples.libsdl.org/SDL3/input/02-joystick-events/

package main

import (
	"fmt"
	"math/rand/v2"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/binsdl"
)

const (
	MotionEventCooldown = 40
)

type EventMessage struct {
	str        string
	color      sdl.Color
	startTicks uint64
}

var (
	colors   [64]sdl.Color
	messages []EventMessage
)

func hatStateString(state uint8) string {
	switch state {
	case sdl.HAT_CENTERED:
		return "CENTERED"
	case sdl.HAT_UP:
		return "UP"
	case sdl.HAT_RIGHT:
		return "RIGHT"
	case sdl.HAT_DOWN:
		return "DOWN"
	case sdl.HAT_LEFT:
		return "LEFT"
	case sdl.HAT_RIGHTUP:
		return "RIGHT+UP"
	case sdl.HAT_RIGHTDOWN:
		return "RIGHT+DOWN"
	case sdl.HAT_LEFTUP:
		return "LEFT+UP"
	case sdl.HAT_LEFTDOWN:
		return "LEFT+DOWN"
	default:
		return "UNKNOWN"
	}
}

func batteryStateString(state sdl.PowerState) string {
	switch state {
	case sdl.POWERSTATE_ERROR:
		return "ERROR"
	case sdl.POWERSTATE_UNKNOWN:
		return "UNKNOWN"
	case sdl.POWERSTATE_ON_BATTERY:
		return "ON BATTERY"
	case sdl.POWERSTATE_NO_BATTERY:
		return "NO BATTERY"
	case sdl.POWERSTATE_CHARGING:
		return "CHARGING"
	case sdl.POWERSTATE_CHARGED:
		return "CHARGED"
	default:
		return "UNKNOWN"
	}
}

func AddMessage(jid sdl.JoystickID, format string, values ...any) {
	var msg EventMessage

	color := colors[int(jid)%len(colors)]
	msg.str = fmt.Sprintf(format, values...)
	msg.color = color
	msg.startTicks = sdl.Ticks()
	messages = append(messages, msg)
}

func main() {
	defer binsdl.Load().Unload() // sdl.LoadLibrary(pathToSDLBinary)
	defer sdl.Quit()
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_JOYSTICK)
	if err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("examples/input/02-joystick-events", 640, 480, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	defer renderer.Destroy()

	for i := range colors {
		colors[i].R = uint8(rand.IntN(255))
		colors[i].G = uint8(rand.IntN(255))
		colors[i].B = uint8(rand.IntN(255))
		colors[i].A = 255
	}
	colors[0].R, colors[0].G, colors[0].B, colors[0].A = 255, 255, 255, 255

	AddMessage(0, "Please plug in a joystick.")

	var (
		axisMotionCooldownTime uint64
		ballMotionCooldownTime uint64
	)

	sdl.RunLoop(func() error {
		var event sdl.Event

		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_QUIT:
				return sdl.EndLoop
			case sdl.EVENT_JOYSTICK_ADDED:
				/* this event is sent for each hotplugged stick, but also each already-connected joystick during SDL_Init(). */
				evt := event.JoyDeviceEvent()
				joystick, err := evt.Which.OpenJoystick()
				if err != nil {
					AddMessage(evt.Which, "Joystick %d add, but not opened: %v", evt.Which, err)
				} else {
					name, _ := joystick.Name()
					AddMessage(evt.Which, "Joystick %d ('%s') added", evt.Which, name)
				}
			case sdl.EVENT_JOYSTICK_REMOVED:
				evt := event.JoyDeviceEvent()
				joystick, _ := evt.Which.Joystick()
				if joystick != nil {
					joystick.Close() /* our joystick was unplugged. */
				}
				AddMessage(evt.Which, "Joystick %d removed", evt.Which)
			case sdl.EVENT_JOYSTICK_AXIS_MOTION:
				evt := event.JoyAxisEvent()
				now := sdl.Ticks()
				if now >= axisMotionCooldownTime {
					axisMotionCooldownTime = now + MotionEventCooldown
					AddMessage(evt.Which, "Joystick %d axis %d -> %d", evt.Which, evt.Axis, evt.Value)
				}
			case sdl.EVENT_JOYSTICK_BALL_MOTION:
				evt := event.JoyBallEvent()
				now := sdl.Ticks()
				if now >= ballMotionCooldownTime {
					ballMotionCooldownTime = now + MotionEventCooldown
					AddMessage(evt.Which, "Joystick %d ball %d -> %d, %d", evt.Which, evt.Ball, evt.Xrel, evt.Yrel)
				}
			case sdl.EVENT_JOYSTICK_HAT_MOTION:
				evt := event.JoyHatEvent()
				AddMessage(evt.Which, "Joystick %d hat %d -> %s", evt.Which, evt.Hat, hatStateString(evt.Value))
			case sdl.EVENT_JOYSTICK_BUTTON_UP, sdl.EVENT_JOYSTICK_BUTTON_DOWN:
				evt := event.JoyButtonEvent()
				state := "RELEASED"
				if evt.Down {
					state = "PRESSED"
				}
				AddMessage(evt.Which, "Joystick %d button %d -> %s", evt.Which, evt.Button, state)
			case sdl.EVENT_JOYSTICK_BATTERY_UPDATED:
				evt := event.JoyBatteryEvent()
				AddMessage(evt.Which, "Joystick %d batter -> %s - %d%%", evt.Which, batteryStateString(evt.State), evt.Percent)
			}
		}

		// Rendering

		now := sdl.Ticks()
		msgLifetime := 3500.
		prevY := float32(0)
		winw := int32(640)
		winh := int32(480)

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()
		winw, winh, _ = window.Size()

		start := 0
		for i, msg := range messages {
			lifePercent := float32(now-msg.startTicks) / float32(msgLifetime)
			if lifePercent >= 1 { /* msg is done. */
				start++
				continue
			}
			x := (float32(winw) - float32(len(msg.str)*sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE)) / 2
			y := float32(winh) * lifePercent
			if prevY != 0 && (prevY-y) < sdl.DEBUG_TEXT_FONT_CHARACTER_SIZE {
				messages[i].startTicks = now
				break // wait for the previous message to tick up a little.
			}

			renderer.SetDrawColor(msg.color.R, msg.color.G, msg.color.B, uint8(float32(msg.color.A)*(1-lifePercent)))
			renderer.DebugText(x, y, msg.str)

			prevY = y
		}
		messages = messages[start:]

		/* put the newly-cleared rendering on the screen. */
		renderer.Present()

		return nil
	})
}
