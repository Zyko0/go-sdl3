//go:build !js

package binttf

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/ttf"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binttf: couldn't create a temporary directory: " + err.Error())
	}

	ttfPath := filepath.Join(tmp, ttfLibName)
	err = os.WriteFile(ttfPath, ttfBlob, 0666)
	if err != nil {
		log.Fatal("binttf: couldn't write ttf library to disk: " + err.Error())
	}

	err = ttf.LoadLibrary(ttfPath)
	if err != nil {
		log.Fatal("binttf: couldn't ttf.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := ttf.CloseLibrary()
	if err != nil {
		log.Fatal("binttf: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
