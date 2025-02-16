package binsdl

import (
	"log"
	"os"
	"path/filepath"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/internal"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binary: couldn't create a temporary directory: " + err.Error())
	}

	sdlPath := filepath.Join(tmp, sdlLibName)
	err = os.WriteFile(sdlPath, sdlBlob, 0666)
	if err != nil {
		log.Fatal("binary: couldn't write sdl library to disk: " + err.Error())
	}

	err = sdl.LoadLibrary(sdlPath)
	if err != nil {
		log.Fatal("binary: couldn't sdl.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := sdl.CloseLibrary()
	if err != nil {
		log.Fatal("binary: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
