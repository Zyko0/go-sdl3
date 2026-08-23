//go:build !js

package binmidi

import (
	"bytes"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/midi"
)

type library struct {
	dir string
}

func Load() library {
	tmp, err := internal.TmpDir()
	if err != nil {
		log.Fatal("binmidi: couldn't create a temporary directory: " + err.Error())
	}
	midiPath := filepath.Join(tmp, midiLibName)

	r, err := gzip.NewReader(bytes.NewReader(midiBlob))
	if err != nil {
		log.Fatal("binmidi: couldn't read compressed native_midi binary: " + err.Error())
	}
	defer r.Close()

	f, err := os.Create(midiPath)
	if err != nil {
		log.Fatal("binmidi: couldn't create native_midi library file to disk: " + err.Error())
	}

	_, err = io.Copy(f, r)
	if err != nil {
		f.Close()
		log.Fatal("binmidi: couldn't decompress native_midi library file: " + err.Error())
	}
	f.Close()

	err = midi.LoadLibrary(midiPath)
	if err != nil {
		log.Fatal("binmidi: couldn't midi.LoadLibrary: ", err.Error())
	}

	return library{
		dir: tmp,
	}
}

func (l library) Unload() {
	err := midi.CloseLibrary()
	if err != nil {
		log.Fatal("binmidi: couldn't close library: ", err.Error())
	}
	internal.RemoveTmpDir()
}
