package internal

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
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
		if dir.Err == nil {
			// Ensure the temporary directory is removed if program
			// exits outside main function
			channel := make(chan os.Signal, 1)
			signal.Notify(channel,
				syscall.SIGTERM,
				syscall.SIGINT,
				syscall.SIGQUIT,
			)
			go func() {
				<-channel
				RemoveTmpDir()
			}()
		}
	})

	return dir.Dir, dir.Err
}

func RemoveTmpDir() {
	dir.onceRemove.Do(func() {
		os.RemoveAll(dir.Dir)
	})
	// Clear tmpDir once entry after removal, so that it can be created again
	dir = tmpDir{}
}
