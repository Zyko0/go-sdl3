package binary

import (
	"log"
	"os"
	"path/filepath"

	sdl "github.com/Zyko0/go-sdl3"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := os.MkdirTemp("", "")
	if err != nil {
		log.Fatal("binary: couldn't create a temporary directory: " + err.Error())
	}

	sdlPath := filepath.Join(tmp, "sdl.a")
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
	err = os.RemoveAll(l.dir)
	if err != nil {
		log.Fatal("binary: couldn't remove temporary directory "+l.dir+": ", err)
	}
}
