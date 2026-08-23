package assets

import (
	"bytes"
	"encoding/csv"
	"os"
	"strings"
)

// APIRefEntry is a function documented by a library's quick reference page,
// which is the source of truth for what belongs to its public API.
type APIRefEntry struct {
	Name        string
	Types       []string
	Description string
}

func LoadAPIRef(path string) (map[string]*APIRefEntry, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	csvr := csv.NewReader(bytes.NewReader(b))
	records, err := csvr.ReadAll()
	if err != nil {
		return nil, err
	}

	entries := make(map[string]*APIRefEntry)
	for _, record := range records[1:] { // Skip header
		entries[record[0]] = &APIRefEntry{
			Name:        record[0],
			Types:       strings.Fields(record[1]),
			Description: record[2],
		}
	}

	return entries, nil
}
