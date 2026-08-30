// Package testing holds cross-cutting tests that belong to no single binding
// package.
package testing

import (
	"encoding/json"
	"go/importer"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"
)

type ffiEntry struct {
	Tag       string     `json:"tag"`
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	BitSize   int        `json:"bit-size"`
	BitOffset int        `json:"bit-offset"`
	Value     int64      `json:"value"`
	Fields    []ffiEntry `json:"fields"`
}

type ffiConfig struct {
	OutDir         string   `json:"out_dir"`
	Prefix         string   `json:"prefix"`
	AllowedInclude string   `json:"allowed_include"`
	IgnoredHeaders []string `json:"ignored_headers"`
}

// includes reports whether a c2ffi location names a header the generators read.
// The dumps cover the whole translation unit, so without this every check would
// also see libc and, for the satellite libraries, all of SDL3.
func (c ffiConfig) includes(location string) bool {
	if !strings.HasPrefix(location, c.AllowedInclude) {
		return false
	}
	header, _, _ := strings.Cut(filepath.Base(location), ":")
	return !slices.Contains(c.IgnoredHeaders, header)
}

const modulePath = "github.com/Zyko0/go-sdl3"

var libraries = []string{"sdl", "ttf", "mixer", "img", "midi"}

func readJSON[T any](t *testing.T, path string) T {
	t.Helper()

	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		t.Fatal(path, err)
	}
	return v
}

// assetPath resolves a file under cmd/internal/assets, which cannot be
// imported from outside cmd/ and so is read directly.
func assetPath(kind, library string) string {
	return filepath.Join("..", "..", "cmd", "internal", "assets", kind, library+".json")
}

var packages = map[string]*types.Package{}

// loadPackage type-checks a binding package from source, which is the only way
// to reach constant values and field offsets: reflect can enumerate neither.
// The result is cached because the import costs a couple of seconds and every
// test here walks all five libraries.
func loadPackage(t *testing.T, importPath string) *types.Package {
	t.Helper()

	if pkg, ok := packages[importPath]; ok {
		return pkg
	}

	start := time.Now()
	pkg, err := importer.ForCompiler(token.NewFileSet(), "source", nil).Import(importPath)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("type-checked", importPath, "in", time.Since(start).Round(time.Millisecond))

	packages[importPath] = pkg
	return pkg
}
