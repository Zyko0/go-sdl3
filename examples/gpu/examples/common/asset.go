package common

import (
	"path/filepath"
	"runtime"
)

// TODO: use go:embed?

var (
	BasePath string
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	BasePath = filepath.Join(filepath.Dir(file), "../..")
}
