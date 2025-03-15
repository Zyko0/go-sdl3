//go:build !js

package binmix

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/mixer"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binmix: couldn't create a temporary directory: " + err.Error())
	}

	mixPath := filepath.Join(tmp, mixLibName)
	err = os.WriteFile(mixPath, mixBlob, 0666)
	if err != nil {
		log.Fatal("binmix: couldn't write mixer library to disk: " + err.Error())
	}

	err = mixer.LoadLibrary(mixPath)
	if err != nil {
		log.Fatal("binmix: couldn't mixer.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := mixer.CloseLibrary()
	if err != nil {
		log.Fatal("binmix: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
