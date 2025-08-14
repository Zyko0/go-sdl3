//go:build !js

package binmix

import (
	"bytes"
	"compress/gzip"
	"io"
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

	r, err := gzip.NewReader(bytes.NewReader(mixBlob))
	if err != nil {
		log.Fatal("binmix: couldn't read compressed mixer binary: " + err.Error())
	}
	defer r.Close()

	f, err := os.Create(mixPath)
	if err != nil {
		log.Fatal("binmix: couldn't create mixer library file to disk: " + err.Error())
	}

	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		log.Fatal("binmix: couldn't decompress mixer library file: " + err.Error())
	}
	f.Close()

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
