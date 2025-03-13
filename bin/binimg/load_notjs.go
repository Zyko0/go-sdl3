//go:build !js

package binimg

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Zyko0/go-sdl3/img"
	"github.com/Zyko0/go-sdl3/internal"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binimg: couldn't create a temporary directory: " + err.Error())
	}

	imgPath := filepath.Join(tmp, imgLibName)
	err = os.WriteFile(imgPath, imgBlob, 0666)
	if err != nil {
		log.Fatal("binimg: couldn't write img library to disk: " + err.Error())
	}

	err = img.LoadLibrary(imgPath)
	if err != nil {
		log.Fatal("binimg: couldn't img.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := img.CloseLibrary()
	if err != nil {
		log.Fatal("binimg: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
