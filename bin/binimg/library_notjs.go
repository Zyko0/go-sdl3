//go:build !js

package binimg

import (
	"bytes"
	"compress/gzip"
	"io"
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

	r, err := gzip.NewReader(bytes.NewReader(imgBlob))
	if err != nil {
		log.Fatal("binimg: couldn't read compressed img binary: " + err.Error())
	}
	defer r.Close()

	f, err := os.Create(imgPath)
	if err != nil {
		log.Fatal("binimg: couldn't create img library file to disk: " + err.Error())
	}

	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		log.Fatal("binimg: couldn't decompress img library file: " + err.Error())
	}
	f.Close()

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
