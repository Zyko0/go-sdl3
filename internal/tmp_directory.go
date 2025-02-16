package internal

import (
	"os"
	"sync"
)

type tmpDir struct {
	onceCreate sync.Once
	onceRemove sync.Once

	Dir string
	Err error
}

var dir tmpDir

func TmpDir() (string, error) {
	dir.onceCreate.Do(func() {
		dir.Dir, dir.Err = os.MkdirTemp("", "")
	})

	return dir.Dir, dir.Err
}

func RemoveTmpDir() {
	dir.onceRemove.Do(func() {
		os.RemoveAll(dir.Dir)
	})
}
