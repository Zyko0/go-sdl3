package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/Zyko0/go-sdl3/cmd/internal/assets"
)

var (
	reg = regexp.MustCompile(`i([A-Z][A-Za-z_0-9]+)\(`)

	cfg        assets.Config
	apiRefCode string
)

type refFunc struct {
	Name        string
	Description string
}

var (
	uniqueTypes        = map[string]struct{}{}
	uniqueAPIFunctions = map[string]refFunc{}
)

func AllTypesFromAPIRef() {
	for _, l := range strings.Split(apiRefCode, "\n") {
		l = strings.TrimSpace(l)
		l = strings.ReplaceAll(l, "const ", "")
		l = strings.ReplaceAll(l, " * ", "* ")
		l = strings.ReplaceAll(l, " ** ", "** ")
		l = strings.ReplaceAll(l, "* * ", "** ")
		switch {
		case strings.HasPrefix(l, "//"):
			continue
		case strings.HasPrefix(l, "#"):
			continue
		case l == "":
			continue
		default:
			idx := strings.Index(l, "//")
			comment := ""
			if idx != -1 {
				comment = strings.TrimSpace(l[idx+2:])
				comment = strings.TrimSuffix(comment, "\n")
				l = l[:idx]
			}
			// Parse function prototype
			nameIdx := strings.Index(l[1:], cfg.Prefix)
			// Function name
			name := l[nameIdx+1 : strings.Index(l, "(")]
			// Function comment
			uniqueAPIFunctions[name] = refFunc{
				Name:        name,
				Description: comment,
			}
			// Return type
			typ := strings.TrimSpace(l[:nameIdx])
			typ = strings.ReplaceAll(typ, "*", "")
			uniqueTypes[typ] = struct{}{}
			// Argument types
			args := l[strings.Index(l, "(")+1 : strings.Index(l, ")")]
			argsParts := strings.Split(args, ", ")
			for _, at := range argsParts {
				argParts := strings.Split(at, " ")
				typ = strings.TrimSpace(strings.Join(argParts[:len(argParts)-1], " "))
				typ = strings.ReplaceAll(typ, "*", "")
				typ = strings.ReplaceAll(typ, "struct ", "")
				uniqueTypes[typ] = struct{}{}
			}
		}
	}
	for tp := range uniqueTypes {
		var found bool
		for _, prefix := range cfg.AllowlistedTypePrefixes {
			if strings.HasPrefix(tp, prefix) {
				found = true
				break
			}
		}
		if !found {
			delete(uniqueTypes, tp)
		}
	}
}

func main() {
	var (
		configPath string
		dir        string
	)

	flag.StringVar(&configPath, "config", "", "path to config.json file")
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

	// Download API ref code
	resp, err := http.Get(cfg.QuickAPIRefURL)
	if err != nil {
		log.Fatal("couldn't download api ref: ", err)
	}
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("couldn't read http response body: ", err)
	}
	apiRefCode = string(b)
	_, apiRefCode, _ = strings.Cut(apiRefCode, "```c")
	apiRefCode, _, _ = strings.Cut(apiRefCode, "```")

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	path = filepath.Join(path, dir)

	AllTypesFromAPIRef()

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("err: ", err)
	}

	var urlLibrarySuffix string
	if cfg.LibraryName != "sdl" {
		urlLibrarySuffix = "_" + cfg.LibraryName
	}

	var annotations int

	for _, e := range entries {
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
		var funcName string
		var lineIndex int

		lines := strings.Split(string(b), "\n")
		for i, l := range lines {
			if inFunc {
				braces += strings.Count(l, "{")
				braces -= strings.Count(l, "}")
				if braces > 0 {
					if reg.MatchString(l) {
						matches := reg.FindAll([]byte(l), -1)
						for _, m := range matches {
							name := string(m[1 : len(m)-1])
							_, found := uniqueAPIFunctions[name]
							if !found {
								name = cfg.Prefix + name
								_, found = uniqueAPIFunctions[name]
							}
							if found {
								funcName = name
							}
						}
					}
				} else {
					inFunc = false
					braces = 0
					// Add comments + whole function
					if funcName != "" {
						outLines = append(outLines, fmt.Sprintf(
							"// %s - %s", funcName, uniqueAPIFunctions[funcName].Description,
						))
						outLines = append(outLines, fmt.Sprintf(
							"// (https://wiki.libsdl.org/SDL3%s/%s)", urlLibrarySuffix, funcName,
						))
						annotations++
						edited = true
					}
					// Function lines
					for lineIndex <= i {
						outLines = append(outLines, lines[lineIndex])
						lineIndex++
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
			funcName = ""
			lineIndex = i
		}

		// Write file if there has been some changes
		if edited {
			err = os.WriteFile(filepath.Join(dir, e.Name()), []byte(strings.Join(outLines, "\n")), 0666)
			if err != nil {
				log.Fatalf("couldn't write file %s: %v\n", e.Name(), err)
			}
		}
	}

	fmt.Println("Total functions annotated:", annotations)
}
