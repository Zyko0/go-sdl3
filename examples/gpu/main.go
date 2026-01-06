package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/examples/gpu/common"
	"github.com/Zyko0/go-sdl3/examples/gpu/examples"
	"github.com/Zyko0/go-sdl3/sdl"
)

type ExampleInterface interface {
	String() string
	Init(context *common.Context) error
	Update(context *common.Context) error
	Draw(context *common.Context) error
	Quit(context *common.Context)
}

var allExamples = []ExampleInterface{
	examples.ClearScreenExample,
	examples.ClearScreenMultiWindowExample,
	examples.BasicTriangleExample,
	examples.BasicVertexBufferExample,
	examples.CullModeExample,
	examples.BasicStencilExample,
	examples.InstancedIndexedExample,
	examples.TexturedQuadExample,
	examples.TexturedAnimatedQuadExample,
	examples.Clear3DSliceExample,
	examples.BasicComputeExample,
	examples.ComputeUniformsExample,
	examples.ToneMappingExample,
	examples.CustomSamplingExample,
	examples.DrawIndirectExample,
	examples.ComputeSamplerExample,
	examples.CopyAndReadbackExample,
	examples.CopyConsistencyExample,
	examples.Texture2DArrayExample,
	examples.TriangleMSAAExample,
	examples.CubemapExample,
	examples.WindowResizeExample,
	examples.Blit2DArrayExample,
	examples.BlitCubeExample,
	examples.BlitMirrorExample,
	examples.GenerateMipmapsExample,
	examples.LatencyExample,
	examples.DepthSamplerExample,
	examples.ComputeSpriteBatchExample,
	examples.PullSpriteBatchExample,
	examples.TextureTypeTestExample,
	examples.CompressedTexturesExample,
}

func main() {
	var context common.Context
	var exampleIndex int = -1
	var gotoExampleIndex int
	var quit bool
	var lastTime = time.Now()

	if len(os.Args) > 1 {
		exampleName := os.Args[1]
		exampleNameLower := strings.ToLower(os.Args[1])
		foundExample := false

		for i, example := range allExamples {
			if strings.ToLower(example.String()) == exampleNameLower {
				gotoExampleIndex = i
				foundExample = true
				break
			}
		}

		if !foundExample {
			fmt.Printf("No example named \"%s\" exists\n", exampleName)
			os.Exit(1)
		}
	}

	err := sdl.LoadLibrary(sdl.Path())
	if err != nil {
		fmt.Println("Failed to load " + sdl.Path())
		fmt.Println("Will load embedded instead")
		defer binsdl.Load().Unload()
	}

	err = sdl.Init(sdl.INIT_VIDEO | sdl.INIT_GAMEPAD)
	if err != nil {
		panic("failed to initialize SDL: " + err.Error())
	}

	fmt.Println("SDL: " + sdl.GetVersion().String())

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
					allExamples[exampleIndex].Quit(&context)
				}
				quit = true
			case sdl.EVENT_GAMEPAD_ADDED:
				if gamepad == nil {
					deviceEvent := event.GamepadDeviceEvent()
					gamepad, err = deviceEvent.Which.OpenGamepad()
					if err != nil {
						panic("failed to open gamepad: " + err.Error())
					}
				}
			case sdl.EVENT_GAMEPAD_REMOVED:
				if gamepad == nil {
					deviceEvent := event.GamepadDeviceEvent()
					gamepadID, err := gamepad.ID()
					if err != nil {
						panic("failed to get gamepad id: " + err.Error())
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
					if gotoExampleIndex >= len(allExamples) {
						gotoExampleIndex = 0
					}
				case sdl.K_A:
					gotoExampleIndex = exampleIndex - 1
					if gotoExampleIndex < 0 {
						gotoExampleIndex = len(allExamples) - 1
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
				allExamples[exampleIndex].Quit(&context)
				context = common.Context{}
			}

			exampleIndex = gotoExampleIndex
			context.ExampleName = allExamples[exampleIndex].String()
			fmt.Println("STARTING EXAMPLE: " + context.ExampleName)
			err = allExamples[exampleIndex].Init(&context)
			if err != nil {
				panic("failed to initialize: " + err.Error())
			}

			gotoExampleIndex = -1
		}

		newTime := time.Now()
		context.DeltaTime = float32(
			newTime.Sub(lastTime).Seconds(),
		)
		lastTime = newTime

		err = allExamples[exampleIndex].Update(&context)
		if err != nil {
			panic("failed to update: " + err.Error())
		}

		if canDraw {
			err = allExamples[exampleIndex].Draw(&context)
			if err != nil {
				panic("failed to draw: " + err.Error())
			}
		}
	}
}
