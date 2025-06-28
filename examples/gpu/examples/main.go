package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/sdl"
)

var examples = []Example{
	ClearScreenExample,
}

func main() {
	var context Context
	var exampleIndex int = -1
	var gotoExampleIndex int
	var quit bool
	var lastTime time.Time

	if len(os.Args) > 1 {
		exampleName := os.Args[1]
		foundExample := false

		for i, example := range examples {
			if example.Name == exampleName {
				gotoExampleIndex = i
				foundExample = true
				break
			}
		}

		if !foundExample {
			fmt.Printf("No example named '%s' exists.\n", exampleName)
			os.Exit(1)
		}
	}

	defer binsdl.Load().Unload()

	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_GAMEPAD)
	if err != nil {
		panic("Failed to initialize SDL: " + err.Error())
	}

	// InitializeAssetLoader()
	// SDL_AddEventWatch(AppLifecycleWatcher, NULL);

	fmt.Println("Welcome to the SDL_GPU example suite!")
	fmt.Println("Press A/D (or LB/RB) to move between examples!")

	// gamepad

	var gamepad *sdl.Gamepad
	var canDraw bool = true

	// sdl.RunLoop(func() error {

	for !quit {

		context.LeftPressed = false
		context.RightPressed = false
		context.DownPressed = false
		context.UpPressed = false

		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type {
			case sdl.EVENT_QUIT:
				if exampleIndex != -1 {
					examples[exampleIndex].Quit(&context)
				}
				quit = true
			case sdl.EVENT_GAMEPAD_ADDED:
				if gamepad == nil {
					deviceEvent := event.GamepadDeviceEvent()
					gamepad, err = deviceEvent.Which.OpenGamepad()
					if err != nil {
						panic("Failed to open gamepad: " + err.Error())
					}
				}
			case sdl.EVENT_GAMEPAD_REMOVED:
				if gamepad == nil {
					deviceEvent := event.GamepadDeviceEvent()
					gamepadID, err := gamepad.ID()
					if err != nil {
						panic("Failed to get gamepad ID: " + err.Error())
					}
					if deviceEvent.Which == gamepadID {
						gamepad.Close()
					}
				}
			// case sdl.EVENT_USER:
			// 	// implement
			case sdl.EVENT_KEY_DOWN:
				keyEvent := event.KeyboardEvent()
				switch keyEvent.Key {
				case sdl.K_D:
					gotoExampleIndex = exampleIndex + 1
					if gotoExampleIndex >= len(examples) {
						gotoExampleIndex = 0
					}
				case sdl.K_A:
					gotoExampleIndex = exampleIndex - 1
					if gotoExampleIndex < 0 {
						gotoExampleIndex = len(examples) - 1
					}
				case sdl.K_LEFT:
					context.LeftPressed = true
				case sdl.K_RIGHT:
					context.RightPressed = true
				case sdl.K_DOWN:
					context.DownPressed = true
				case sdl.K_UP:
					context.UpPressed = true
				}
				// case sdl.EVENT_GAMEPAD_BUTTON_DOWN:
				// 	// implement
				// case sdl.EVENT_GAMEPAD_BUTTON_DOWN:
				// 	// implement
			}
		}

		if quit {
			break
		}

		if gotoExampleIndex != -1 {
			if exampleIndex != -1 {
				examples[exampleIndex].Quit(&context)
				context = Context{}
			}

			exampleIndex = gotoExampleIndex
			context.ExampleName = examples[exampleIndex].Name
			fmt.Println("STARTING EXAMPLE: " + context.ExampleName)
			err = examples[exampleIndex].Init(&context)
			if err != nil {
				panic("Init failed! " + err.Error())
			}

			gotoExampleIndex = -1
		}

		newTime := time.Now()
		context.DeltaTime = float32(
			newTime.Sub(lastTime).Microseconds(),
		) * 0.001 * 0.001
		lastTime = newTime

		err = examples[exampleIndex].Update(&context)
		if err != nil {
			panic("Update failed! " + err.Error())
		}

		if canDraw {
			err = examples[exampleIndex].Draw(&context)
			if err != nil {
				panic("Draw failed! " + err.Error())
			}
		}
	}
}
