package bin_test

import (
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/Zyko0/go-sdl3/bin/binimg"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

func Test_EmbeddedBinaries(t *testing.T) {
	debug.SetPanicOnFault(true)

	os.Setenv("SDL_AUDIODRIVER", "dummy")
	os.Setenv("SDL_VIDEODRIVER", "dummy")

	t.Log("OS:", runtime.GOOS, "Arch:", runtime.GOARCH)

	defer binsdl.Load().Unload()
	defer binttf.Load().Unload()
	// defer binmix.Load().Unload() // TODO:
	defer binimg.Load().Unload()

	//ttf.Quit()
	//sdl.Quit()

	t.Run("SDL", func(t *testing.T) {
		t.Run("Init", func(t *testing.T) {
			err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			v := sdl.GetVersion()
			t.Log("SDL version:", v.String())
		})
	})

	t.Run("SDL_ttf", func(t *testing.T) {
		t.Run("Init", func(t *testing.T) {
			err := ttf.Init()
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			v := ttf.GetVersion()
			t.Log("SDL_ttf version:", v.String())
		})
	})

	// TODO: mixer
	/*t.Run("SDL_mixer", func(t *testing.T) {
		t.Run("Init", func(t *testing.T) {
			defer binsdl.Load().Unload()
			defer binmix.Load().Unload()
			defer sdl.Quit()

			err := sdl.Init(0)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			err = mixer.Init()
			if err != nil {
				t.Log(err)
				t.FailNow()
			}
			v := mix.GetVersion()
			t.Log("SDL_mixer version:", v.String())
		})
	})*/

	t.Run("SDL_image", func(t *testing.T) {
		t.Run("Init", func(t *testing.T) {
			v := img.GetVersion()
			t.Log("SDL_image version:", v.String())
		})
	})
}
