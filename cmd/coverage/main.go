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
	regDesktop = regexp.MustCompile(`i([A-Z][A-Za-z_0-9]+)\(`)
	regJsFunc  = regexp.MustCompile(`.*\s=\sfunc`)
	regJS      *regexp.Regexp

	cfg        assets.Config
	apiRefCode string
)

func True() *bool {
	b := true
	return &b
}

func False() *bool {
	b := false
	return &b
}

type coverage struct {
	Exposed  *bool
	Filename string
	Line     int
}

type refFunc struct {
	CategoryIndex int
	Name          string
	URL           string

	Desktop coverage
	JS      coverage
}

var (
	categories = map[string][]string{
		"sdl": {
			"Init", "Hints", "Error", "Properties", "Log", "Video",
			"Events", "Keyboard", "Mouse", "Touch", "Gamepad", "Joystick",
			"Haptic", "Audio", "Time", "Timer", "Render", "SharedObject",
			"Thread", "Mutex", "Atomic", "Filesystem", "IOStream", "AsyncIO",
			"Storage", "Pixels", "Surface", "BlendMode", "Rect", "Camera",
			"Clipboard", "Dialog", "GPU", "MessageBox", "Vulkan", "Metal",
			"Platform", "Power", "Sensor", "Process", "Bits", "Endian",
			"Assert", "CPUInfo", "Intrinsics", "Locale", "System", "Misc",
			"GUID", "Main", "Stdinc",
		},
		"img":   {"Image"},
		"ttf":   {"TTF"},
		"mixer": {"Mixer"},
	}
	uniqueAPIFunctions = map[string]*refFunc{}
	functions          []*refFunc
)

func AllFunctions() {
	inComments := false
	categoryIndex := -1
	for _, l := range strings.Split(apiRefCode, "\n") {
		l = strings.TrimSpace(l)
		l = strings.ReplaceAll(l, "const ", "")
		l = strings.ReplaceAll(l, " * ", "* ")
		l = strings.ReplaceAll(l, " ** ", "** ")
		l = strings.ReplaceAll(l, "* * ", "** ")
		switch {
		case strings.HasPrefix(l, "//"):
			if !inComments {
				categoryIndex++
				inComments = true
			}
			continue
		case strings.HasPrefix(l, "#"):
			continue
		case l == "":
			continue
		default:
			inComments = false
			idx := strings.Index(l, "//")
			if idx != -1 {
				l = l[:idx]
			}
			// Parse function name
			nameIdx := strings.Index(l[1:], cfg.Prefix)
			name := l[nameIdx+1 : strings.Index(l, "(")]
			fn := &refFunc{
				CategoryIndex: categoryIndex,
				Name:          name,
			}
			uniqueAPIFunctions[name] = fn
			functions = append(functions, fn)
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
	regJS, err = regexp.Compile(fmt.Sprintf(`"_%s([A-Z][A-Za-z_0-9]+)",`, cfg.Prefix))
	if err != nil {
		log.Fatal(err)
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

	AllFunctions()

	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("err: ", err)
	}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".go") {
			continue
		}

		b, err := os.ReadFile(filepath.Join(path, e.Name()))
		if err != nil {
			log.Fatalf("couldn't read file %s: %v\n", e.Name(), err)
		}

		var inFunc bool
		var braces int
		var funcName string
		var lineIndex int

		var impl *bool

		lines := strings.Split(string(b), "\n")
		isJS := strings.HasPrefix(lines[0], "//go:build js")
		for i, l := range lines {
			if inFunc {
				braces += strings.Count(l, "{")
				braces -= strings.Count(l, "}")
				if braces > 0 {
					switch {
					case !isJS && regDesktop.MatchString(l):
						matches := regDesktop.FindAll([]byte(l), -1)
						for _, m := range matches {
							name := string(m[1 : len(m)-1])
							fn, found := uniqueAPIFunctions[name]
							if !found {
								name = cfg.Prefix + name
								fn, found = uniqueAPIFunctions[name]
							}
							if found {
								funcName = name
								fn.Desktop.Exposed = True()
								fn.Desktop.Line = lineIndex
								fn.Desktop.Filename = dir + "/" + e.Name()
							}
						}
					case isJS && regJS.MatchString(l):
						matches := regJS.FindAll([]byte(l), -1)
						for _, m := range matches {
							name := string(m[6 : len(m)-2])
							fn, found := uniqueAPIFunctions[name]
							if !found {
								name = cfg.Prefix + name
								fn, found = uniqueAPIFunctions[name]
							}
							if found {
								funcName = name
								fn.JS.Line = lineIndex
								fn.JS.Filename = dir + "/" + e.Name()
							}
						}
					case isJS && strings.Contains(l, "panic(\"not implemented on js\")"):
						impl = False()
					case !isJS && strings.Contains(l, "panic(\"not implemented\")"):
						impl = False()
					}
				} else {
					inFunc = false
					braces = 0
					if funcName != "" {
						fn := uniqueAPIFunctions[funcName]
						if impl != nil {
							if isJS {
								fn.JS.Exposed = impl
							} else {
								fn.Desktop.Exposed = impl
							}
						}
					}
					impl = nil
				}
				continue
			}

			if isJS {
				if !regJsFunc.Match([]byte(l)) {
					continue
				}
			} else {
				if !strings.HasPrefix(l, "func ") {
					continue
				}
			}
			inFunc = true
			braces = 1
			funcName = ""
			lineIndex = i
		}
	}
	// Output coverage
	var sb strings.Builder
	categoryIndex := -1

	if cfg.LibraryName == "sdl" {
		sb.WriteString("# API Coverage\n\n")
		sb.WriteString(`
This file tracks the functions that have been wrapped.<br>
The following emojis mean (they are clickable and should link to the code implementation):
- :heavy_check_mark: = implemented
- :x: = not implemented yet
- :question: = not planned / don't know about integrating it or not
`)
	}
	sb.WriteString("\n## " + strings.ToUpper(cfg.LibraryName) + "\n\n")
	for _, fn := range functions {
		if fn.CategoryIndex != categoryIndex {
			categoryIndex = fn.CategoryIndex
			sb.WriteString("### " + categories[cfg.LibraryName][fn.CategoryIndex] + "\n\n")
			sb.WriteString("|Function|Desktop|WASM/js|\n")
			sb.WriteString("|:--|:--:|:--:|\n")
		}

		fn.URL = fmt.Sprintf("https://wiki.libsdl.org/SDL3%s/%s", cfg.URLLibrarySuffix, fn.Name)

		desktop := ":question:"
		js := ":question:"
		if fn.Desktop.Exposed != nil {
			exposedDesktop := *fn.Desktop.Exposed
			if exposedDesktop {
				desktop = ":heavy_check_mark:"
				if fn.JS.Exposed == nil {
					js = ":heavy_check_mark:"
				} else {
					js = ":x:"
				}
			} else {
				desktop = ":x:"
				js = ":x:"
			}
		}
		var urlDesktop, urlJS string
		if fn.Desktop.Filename != "" && fn.Desktop.Line != 0 {
			urlDesktop = fmt.Sprintf("%s#L%d", fn.Desktop.Filename, fn.Desktop.Line)
		}
		if fn.JS.Filename != "" && fn.JS.Line != 0 {
			urlJS = fmt.Sprintf("%s#L%d", fn.JS.Filename, fn.JS.Line)
		}
		sb.WriteString(fmt.Sprintf(
			"| [%s](%s) | [%s](%s) | [%s](%s) |\n",
			fn.Name, fn.URL,
			desktop, urlDesktop,
			js, urlJS,
		))
	}

	var f *os.File
	if cfg.LibraryName == "sdl" {
		f, err = os.Create("COVERAGE.md")
		if err != nil {
			log.Fatal("couldn't create file: ", err)
		}
	} else {
		f, err = os.OpenFile("COVERAGE.md", os.O_APPEND, os.ModeAppend)
		if err != nil {
			log.Fatal("couldn't open file: ", err)
		}
	}
	defer f.Close()
	_, err = f.Write([]byte(sb.String()))
	if err != nil {
		log.Fatal("couldn't write file: ", err)
	}
}
