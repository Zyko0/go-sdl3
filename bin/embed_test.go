package bin_test

import (
	"runtime/debug"
	"testing"

	"github.com/Zyko0/go-sdl3/bin/binimg"
	"github.com/Zyko0/go-sdl3/bin/binsdl"
	"github.com/Zyko0/go-sdl3/bin/binttf"
	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/ttf"
)

func init() {
	debug.SetPanicOnFault(true)
}

func Test_SDL(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		defer binsdl.Load().Unload()
		defer sdl.Quit()

		err := sdl.Init(0)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		v := sdl.GetVersion()
		t.Log("SDL version:", v.String())
	})
}

func Test_SDL_ttf(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		defer binsdl.Load().Unload()
		defer binttf.Load().Unload()
		defer sdl.Quit()
		defer ttf.Quit()

		err := sdl.Init(0)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		err = ttf.Init()
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		v := ttf.GetVersion()
		t.Log("SDL_ttf version:", v.String())
	})
}

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

func Test_SDL_Image(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		t.Log("before load unload")
		defer binsdl.Load().Unload()
		defer binimg.Load().Unload()
		defer sdl.Quit()

		t.Log("before init")
		err := sdl.Init(0)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}

		t.Log("before version")
		v := img.GetVersion()
		t.Log("after version")
		t.Log("SDL_image version:", v.String())
	})
}
