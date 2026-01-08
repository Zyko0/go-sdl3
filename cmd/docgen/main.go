package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Zyko0/go-sdl3/cmd/internal/assets"
)

var (
	regFunc = regexp.MustCompile(`i([A-Za-z][A-Za-z_0-9]+)\(`)

	cfg assets.Config
)

type Entry struct {
	ID          string
	Description string
	URL         string
}

var (
	uniqueAPIIdentifiers = make(map[string]*Entry)
)

func main() {
	var (
		configPath      string
		annotationsPath string
		dir             string
	)

	flag.StringVar(&configPath, "config", "", "path to config.json file")
	flag.StringVar(&annotationsPath, "annotations", "", "path to annotations.csv file")
	flag.StringVar(&dir, "dir", "", "base directory to generate from/to")
	flag.Parse()

	// Parse config
	b, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("couldn't read config.json file: ", err)
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal("couldn't unmarshal config file: ", err)
	}

	// Parse annotations
	b, err = os.ReadFile(annotationsPath)
	if err != nil {
		log.Fatal("couldn't read annotations.csv file: ", err)
	}
	csvr := csv.NewReader(bytes.NewReader(b))
	records, err := csvr.ReadAll()
	if err != nil {
		log.Fatal("couldn't read csv records: ", err)
	}

	for _, record := range records[1:] { // Skip header
		uniqueAPIIdentifiers[record[0]] = &Entry{
			ID:          record[0],
			Description: strings.ReplaceAll(record[1], "\n", ""),
			URL:         record[2],
		}
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	path = filepath.Join(path, dir)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("err: ", err)
	}

	var fnAnnotations int
	var typesAnnotations int
	var enumsAnnotations int

	for _, e := range files {
		if !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		b, err := os.ReadFile(filepath.Join(path, e.Name()))
		if err != nil {
			log.Fatalf("couldn't read file %s: %v\n", e.Name(), err)
		}

		var outLines []string
		var edited bool

		var inFunc bool
		var braces int
		var fnName string
		var fnLineIndex int
		var typeName string
		var regEnumDef *regexp.Regexp

		lines := strings.Split(string(b), "\n")
		for i, l := range lines {
			// Type
			if strings.HasPrefix(l, "type ") {
				parts := strings.Split(l, " ")
				if len(parts) > 2 {
					typeName = parts[1]
					entry, found := uniqueAPIIdentifiers[cfg.Prefix+typeName]
					if found {
						// Define regexp for enum values matching
						regEnumDef = regexp.MustCompile(`([A-Za-z0-9_]*)\s+` + typeName + `\s=\s`)
						// Only append comment lines if there are none
						if i > 0 && strings.TrimSpace(lines[i-1]) == "" {
							// Add comments on top of the type
							outLines = append(outLines, fmt.Sprintf(
								"// %s - %s", entry.ID, entry.Description,
							))
							outLines = append(outLines, fmt.Sprintf(
								"// (%s)", entry.URL,
							))
							typesAnnotations++
							edited = true
						}
					}
				}
				outLines = append(outLines, l)
				continue
			}
			// Enum values
			if regEnumDef != nil && regEnumDef.Match([]byte(l)) {
				matches := regEnumDef.FindAllStringSubmatch(l, -1)
				if len(matches) > 0 && len(matches[0]) == 2 {
					enum := matches[0][1]
					entry, found := uniqueAPIIdentifiers[cfg.Prefix+enum]
					if found {
						// Don't comment twice
						if !strings.Contains(l, " // ") {
							desc := " // " + entry.Description
							outLines = append(outLines, l+desc)
							enumsAnnotations++
							edited = true
							continue
						}
					}
				}
			}
			// Function
			if inFunc {
				braces += strings.Count(l, "{")
				braces -= strings.Count(l, "}")
				if braces > 0 {
					if regFunc.MatchString(l) {
						matches := regFunc.FindAll([]byte(l), -1)
						for _, m := range matches {
							name := string(m[1 : len(m)-1])
							_, found := uniqueAPIIdentifiers[name]
							if !found {
								name = cfg.Prefix + name
								_, found = uniqueAPIIdentifiers[name]
							}
							if found {
								fnName = name
							}
						}
					}
				} else {
					inFunc = false
					braces = 0
					// Add comments + whole function
					entry, found := uniqueAPIIdentifiers[fnName]
					if found {
						outLines = append(outLines, fmt.Sprintf(
							"// %s - %s", entry.ID, entry.Description,
						))
						outLines = append(outLines, fmt.Sprintf(
							"// (%s)", entry.URL,
						))
						fnAnnotations++
						edited = true
					}
					// Function lines
					for fnLineIndex <= i {
						outLines = append(outLines, lines[fnLineIndex])
						fnLineIndex++
					}
				}
				continue
			}

			if !strings.HasPrefix(l, "func ") {
				outLines = append(outLines, l)
				continue
			}
			// Skip if there's something written on top of the function
			// (assuming a comment already)
			if strings.TrimSpace(lines[i-1]) != "" {
				outLines = append(outLines, l)
				continue
			}
			inFunc = true
			braces = 1
			fnName = ""
			fnLineIndex = i
		}

		// Write file if there has been some changes
		if edited {
			err = os.WriteFile(filepath.Join(dir, e.Name()), []byte(strings.Join(outLines, "\n")), 0666)
			if err != nil {
				log.Fatalf("couldn't write file %s: %v\n", e.Name(), err)
			}
		}
	}

	fmt.Println("Total functions annotated:", fnAnnotations)
	fmt.Println("Total types annotated:", typesAnnotations)
	fmt.Println("Total enums annotated:", enumsAnnotations)
}
