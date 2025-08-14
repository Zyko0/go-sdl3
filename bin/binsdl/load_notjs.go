//go:build !js

package binsdl

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binsdl: couldn't create a temporary directory: " + err.Error())
	}
	sdlPath := filepath.Join(tmp, sdlLibName)

	r, err := gzip.NewReader(bytes.NewReader(sdlBlob))
	if err != nil {
		log.Fatal("binsdl: couldn't read compressed sdl binary: " + err.Error())
	}
	defer r.Close()

	f, err := os.Create(sdlPath)
	if err != nil {
		log.Fatal("binsdl: couldn't create sdl library file to disk: " + err.Error())
	}

	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		log.Fatal("binsdl: couldn't decompress sdl library file: " + err.Error())
	}
	f.Close()

	err = sdl.LoadLibrary(sdlPath)
	if err != nil {
		log.Fatal("binsdl: couldn't sdl.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := sdl.CloseLibrary()
	if err != nil {
		log.Fatal("binsdl: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
