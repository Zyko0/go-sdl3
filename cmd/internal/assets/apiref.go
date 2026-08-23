package assets

import (
	"bytes"
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

// APIRefEntry is a function documented by a library's quick reference page,
// which is the source of truth for what belongs to its public API.
type APIRefEntry struct {
	Name string
	// Group and Order place the function back where the reference had it:
	// Group is the banner-delimited section it belongs to, Order its rank in
	// the page. Both are needed because the CSV is sorted by name.
	Group       int
	Order       int
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
		group, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		order, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		entries[record[0]] = &APIRefEntry{
			Name:        record[0],
			Group:       group,
			Order:       order,
			Types:       strings.Fields(record[3]),
			Description: record[4],
		}
	}

	return entries, nil
}
