//go:build !js

package binttf

import (
	"bytes"
	"compress/gzip"
	"io"
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

	r, err := gzip.NewReader(bytes.NewReader(ttfBlob))
	if err != nil {
		log.Fatal("binttf: couldn't read compressed ttf binary: " + err.Error())
	}
	defer r.Close()

	f, err := os.Create(ttfPath)
	if err != nil {
		log.Fatal("binttf: couldn't create ttf library file to disk: " + err.Error())
	}

	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		log.Fatal("binttf: couldn't decompress ttf library file: " + err.Error())
	}
	f.Close()

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
