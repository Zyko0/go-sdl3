package assets

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
)

type WikiEntry struct {
	Name        string
	Description string
	URL         string
}

func LoadWikiAnnotations(path string) (map[string]*WikiEntry, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	csvr := csv.NewReader(bytes.NewReader(b))
	records, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}

	entries := make(map[string]*WikiEntry)
	for _, record := range records[1:] { // Skip header
		entries[record[0]] = &WikiEntry{
			Name:        record[0],
			Description: strings.ReplaceAll(record[1], "\n", ""),
			URL:         record[2],
		}
	}

	return entries, nil
}
