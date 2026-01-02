package main

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Entry struct {
	Id          string
	Description string
	URL         string
}

func main() {
	type config struct {
		Library string
		Prefix  string
	}
	cfgs := []*config{
		{"SDL3", "SDL_"},
		{"SDL3_image", "IMG_"},
		{"SDL3_mixer", "MIX_"},
		{"SDL3_ttf", "TTF_"},
	}

	var entries []*Entry
	for _, c := range cfgs {
		path := filepath.Join("sdlwiki", c.Library)
		files, err := fs.ReadDir(os.DirFS("."), path)
		if err != nil {
			log.Fatalf("couldn't read dir sdlwiki: %v\n", err)
		}
		for _, file := range files {
			if file.IsDir() || !strings.HasPrefix(file.Name(), c.Prefix) {
				continue
			}

			fullpath := filepath.Join(path, file.Name())
			content, err := os.ReadFile(fullpath)
			if err != nil {
				log.Fatalf("couldn't read file '%s': %v\n", fullpath, err)
			}

			lines := strings.Split(string(content), "\n")
			var description string
			var hasTitle bool
			for i := 0; i < len(lines); i++ {
				if lines[i] == "" {
					// Documentation content ended
					if description != "" {
						break
					}
					continue
				}
				if strings.HasPrefix(lines[i], "# ") {
					hasTitle = true
					continue
				}
				if hasTitle {
					if strings.HasPrefix(lines[i], "Please refer to [") {
						break
					}
					// Documentation line
					description += lines[i] + "\n"
				}
			}
			// If no content was found, skip this file
			if description == "" {
				continue
			}

			id := strings.TrimSuffix(file.Name(), ".md")
			entries = append(entries, &Entry{
				Id:          id,
				Description: description,
				URL:         fmt.Sprintf("https://wiki.libsdl.org/%s/%s", c.Library, id),
			})
		}
	}

	f, err := os.Create("annotations.csv")
	if err != nil {
		log.Fatalf("couldn't create 'annotations.csv': %v\n", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{"id", "description", "url"})
	for _, e := range entries {
		if err := w.Write([]string{e.Id, e.Description, e.URL}); err != nil {
			log.Fatalf("couldn't write entry to csv writer: %v\n", err)
		}
	}
	w.Flush()

	err = w.Error()
	if err != nil {
		log.Fatalf("error flushing csv writer: %v\n", err)
	}
}
