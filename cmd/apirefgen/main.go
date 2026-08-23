package main

import (
	"encoding/csv"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/Zyko0/go-sdl3/cmd/internal/assets"
)

var identifier = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

type entry struct {
	Name        string
	ReturnType  string
	Parameters  string
	Types       []string
	Description string
}

// normalize collapses the pointer spellings used across the reference so a
// type reads the same whether the page writes "SDL_Window *", "SDL_Window*"
// or "SDL_Window * *".
func normalize(l string) string {
	l = strings.ReplaceAll(l, "const ", "")
	l = strings.ReplaceAll(l, " * ", "* ")
	l = strings.ReplaceAll(l, " ** ", "** ")
	l = strings.ReplaceAll(l, "* * ", "** ")

	return l
}

// typeOf reduces a declaration fragment to the bare type it references, or
// an empty string when it references none ("void", "...").
func typeOf(s string) string {
	s = strings.ReplaceAll(s, "*", "")
	s = strings.ReplaceAll(s, "struct ", "")
	s = strings.TrimSpace(s)
	if s == "void" || !identifier.MatchString(s) {
		return ""
	}

	return s
}

// paramTypes returns the bare type of every argument, so consumers never have
// to re-split a signature.
func paramTypes(params string) []string {
	var types []string

	params = strings.TrimSpace(params)
	if params == "" || params == "void" {
		return nil
	}

	for p := range strings.SplitSeq(params, ", ") {
		parts := strings.Fields(p)
		if len(parts) < 2 {
			continue
		}
		t := typeOf(strings.Join(parts[:len(parts)-1], " "))
		if t != "" && !slices.Contains(types, t) {
			types = append(types, t)
		}
	}

	return types
}

func parse(prefix, src string) []entry {
	var entries []entry

	for l := range strings.SplitSeq(src, "\n") {
		l = strings.TrimSpace(l)
		switch {
		case l == "":
			continue
		case strings.HasPrefix(l, "//"): // section banners are ASCII art, nothing to read
			continue
		case strings.HasPrefix(l, "#"): // macro definitions are not part of the callable API
			continue
		}

		proto, description, _ := strings.Cut(l, "//")
		description = strings.TrimSpace(description)
		// The raw prototype keeps qualifiers the reader expects to see in a
		// diff; the normalized one is what types are read from.
		raw := strings.TrimSuffix(strings.TrimSpace(proto), ";")
		norm := normalize(raw)

		open, close := strings.Index(norm, "("), strings.LastIndex(norm, ")")
		rawOpen, rawClose := strings.Index(raw, "("), strings.LastIndex(raw, ")")
		if open == -1 || close < open {
			log.Fatal("couldn't parse prototype: ", l)
		}
		// Search past the first character so a return type carrying the same
		// prefix (e.g. "SDL_Window* SDL_CreateWindow") is skipped.
		nameIdx := strings.Index(norm[1:], prefix)
		if nameIdx == -1 {
			log.Fatal("couldn't find the library prefix in prototype: ", l)
		}
		nameIdx++

		e := entry{
			Name:        strings.TrimSpace(norm[nameIdx:open]),
			ReturnType:  strings.TrimSpace(norm[:nameIdx]),
			Parameters:  strings.TrimSpace(raw[rawOpen+1 : rawClose]),
			Types:       paramTypes(norm[open+1 : close]),
			Description: description,
		}
		if e.Parameters == "void" {
			e.Parameters = ""
		}
		if t := typeOf(e.ReturnType); t != "" && !slices.Contains(e.Types, t) {
			e.Types = append(e.Types, t)
		}

		entries = append(entries, e)
	}

	// Sorted by name so an upstream reordering of the reference produces no diff.
	slices.SortFunc(entries, func(a, b entry) int {
		return strings.Compare(a.Name, b.Name)
	})

	return entries
}

func download(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("couldn't download api ref: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("couldn't download api ref: unexpected status ", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("couldn't read http response body: ", err)
	}
	src := string(b)

	_, src, found := strings.Cut(src, "```c")
	if !found {
		log.Fatal("couldn't find a c code fence in the api ref")
	}
	src, _, found = strings.Cut(src, "```")
	if !found {
		log.Fatal("couldn't find the end of the c code fence in the api ref")
	}

	return src
}

func write(path string, entries []entry) {
	err := os.MkdirAll(filepath.Dir(path), 0o755)
	if err != nil {
		log.Fatal("couldn't create output directory: ", err)
	}

	f, err := os.Create(path)
	if err != nil {
		log.Fatal("couldn't create output file: ", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	err = w.Write([]string{"name", "return_type", "parameters", "types", "description"})
	if err != nil {
		log.Fatal("couldn't write csv header: ", err)
	}
	for _, e := range entries {
		err = w.Write([]string{
			e.Name,
			e.ReturnType,
			e.Parameters,
			strings.Join(e.Types, " "),
			e.Description,
		})
		if err != nil {
			log.Fatal("couldn't write csv record: ", err)
		}
	}
}

func main() {
	var configPath, outPath string

	flag.StringVar(&configPath, "config", "", "path to config.json file")
	flag.StringVar(&outPath, "out", "", "path to the generated apiref csv file")
	flag.Parse()

	cfg, err := assets.LoadConfig(configPath)
	if err != nil {
		log.Fatal("couldn't parse config file: ", err)
	}
	if cfg.QuickAPIRefURL == "" {
		log.Fatal("config has no quick_api_ref_url")
	}

	entries := parse(cfg.Prefix, download(cfg.QuickAPIRefURL))
	if len(entries) == 0 {
		log.Fatal("no function found in the api ref")
	}
	write(outPath, entries)

	log.Printf("%s: wrote %d functions to %s", cfg.LibraryName, len(entries), outPath)
}
