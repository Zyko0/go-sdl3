package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/Zyko0/go-sdl3/cmd/internal/assets"
)

var (
	regDesktop = regexp.MustCompile(`i([A-Z][A-Za-z_0-9]+)\(`)
	regJsFunc  = regexp.MustCompile(`.*\s=\sfunc`)
	regJS      *regexp.Regexp

	cfg *assets.Config
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
	Category string
	Name     string
	URL      string

	// Position in the reference, so functions keep the order their author gave
	// them rather than an alphabetical one.
	Group int
	Order int

	Desktop coverage
	JS      coverage
}

var (
	// Display order of the sections. A header absent from this list is
	// reported rather than silently dropped.
	categoryOrder = map[string][]string{
		"sdl": {
			"SDL_init.h", "SDL_hints.h", "SDL_error.h", "SDL_version.h",
			"SDL_properties.h", "SDL_log.h", "SDL_video.h", "SDL_events.h",
			"SDL_keyboard.h", "SDL_mouse.h", "SDL_touch.h", "SDL_gamepad.h",
			"SDL_joystick.h", "SDL_haptic.h", "SDL_audio.h", "SDL_time.h",
			"SDL_timer.h", "SDL_render.h", "SDL_loadso.h", "SDL_thread.h",
			"SDL_mutex.h", "SDL_atomic.h", "SDL_filesystem.h", "SDL_iostream.h",
			"SDL_asyncio.h", "SDL_storage.h", "SDL_pixels.h", "SDL_surface.h",
			"SDL_blendmode.h", "SDL_rect.h", "SDL_camera.h", "SDL_messagebox.h",
			"SDL_clipboard.h", "SDL_dialog.h", "SDL_tray.h", "SDL_notification.h",
			"SDL_gpu.h", "SDL_vulkan.h", "SDL_metal.h", "SDL_power.h",
			"SDL_sensor.h", "SDL_process.h", "SDL_bits.h", "SDL_endian.h",
			"SDL_assert.h", "SDL_cpuinfo.h", "SDL_locale.h", "SDL_system.h",
			"SDL_misc.h", "SDL_guid.h", "SDL_stdinc.h",
		},
		"img":   {"SDL_image.h"},
		"ttf":   {"SDL_ttf.h", "SDL_textengine.h"},
		"mixer": {"SDL_mixer.h"},
		"midi":  {"SDL_native_midi.h"},
	}
	// Headers whose section title is not just the title-cased header stem.
	categoryLabels = map[string]string{
		"SDL_loadso.h":      "SharedObject",
		"SDL_iostream.h":    "IOStream",
		"SDL_asyncio.h":     "AsyncIO",
		"SDL_blendmode.h":   "BlendMode",
		"SDL_messagebox.h":  "MessageBox",
		"SDL_gpu.h":         "GPU",
		"SDL_cpuinfo.h":     "CPUInfo",
		"SDL_guid.h":        "GUID",
		"SDL_image.h":       "Image",
		"SDL_ttf.h":         "TTF",
		"SDL_textengine.h":  "TTF",
		"SDL_mixer.h":       "Mixer",
		"SDL_native_midi.h": "NativeMIDI",
	}
	collapsedCategories = map[string]struct{}{
		"Error":        {},
		"Version":      {},
		"Log":          {},
		"Time":         {},
		"SharedObject": {},
		"Thread":       {},
		"Mutex":        {},
		"Atomic":       {},
		"Filesystem":   {},
		"Vulkan":       {},
		"Metal":        {},
		"Process":      {},
		"Bits":         {},
		"Endian":       {},
		"Assert":       {},
		"CPUInfo":      {},
		"Intrinsics":   {},
		"Locale":       {},
		"System":       {},
		"Misc":         {},
		"GUID":         {},
		"Stdinc":       {},
	}
	uniqueAPIFunctions = map[string]*refFunc{}
	functions          []*refFunc
)

func label(header string) string {
	if l, ok := categoryLabels[header]; ok {
		return l
	}

	stem := strings.TrimSuffix(strings.TrimPrefix(header, "SDL_"), ".h")

	return strings.ToUpper(stem[:1]) + stem[1:]
}

// groupHeaders names every section of the reference. A section is named after
// the header its functions are declared in; sections whose functions are all
// newer than the ffi entries keep no such evidence, so they take what is left
// of the expected headers, in order.
func groupHeaders(apiref map[string]*assets.APIRefEntry, ffiEntries []*assets.FFIEntry) map[int]string {
	votes := map[int]map[string]int{}
	groups := map[int]struct{}{}

	for _, e := range apiref {
		groups[e.Group] = struct{}{}
	}
	for _, e := range ffiEntries {
		if e.Tag != "function" || !strings.HasPrefix(e.Location, cfg.AllowedInclude) {
			continue
		}
		ref, ok := apiref[e.Name]
		if !ok {
			continue
		}
		header, _, _ := strings.Cut(filepath.Base(e.Location), ":")
		if votes[ref.Group] == nil {
			votes[ref.Group] = map[string]int{}
		}
		votes[ref.Group][header]++
	}

	headers := map[int]string{}
	for group, v := range votes {
		var best string
		for header, n := range v {
			if n > v[best] {
				best = header
			}
		}
		headers[group] = best
	}

	// Sections left without evidence are matched against the headers no other
	// section claimed, both taken in order.
	var orphans []int
	for group := range groups {
		if _, ok := headers[group]; !ok {
			orphans = append(orphans, group)
		}
	}
	slices.Sort(orphans)
	var free []string
	for _, header := range categoryOrder[cfg.LibraryName] {
		if !slices.Contains(slices.Collect(maps.Values(headers)), header) {
			free = append(free, header)
		}
	}
	for i, group := range orphans {
		if i >= len(free) {
			log.Printf("%s: section %d has no header to name it", cfg.LibraryName, group)
			break
		}
		headers[group] = free[i]
	}

	return headers
}

// AllFunctions lists the documented public API in the order the reference
// gives it. Functions the ffi entries do not know about are kept: they are the
// ones a newer upstream added, and they belong in the table as unimplemented.
func AllFunctions(apiref map[string]*assets.APIRefEntry, ffiEntries []*assets.FFIEntry) {
	headers := groupHeaders(apiref, ffiEntries)

	for _, e := range apiref {
		fn := &refFunc{
			Category: label(headers[e.Group]),
			Name:     e.Name,
			Group:    e.Group,
			Order:    e.Order,
		}
		uniqueAPIFunctions[e.Name] = fn
		functions = append(functions, fn)
	}

	slices.SortFunc(functions, func(a, b *refFunc) int {
		return a.Order - b.Order
	})
}

func main() {
	var (
		configPath string
		ffiPath    string
		apirefPath string
		dir        string
	)

	flag.StringVar(&configPath, "config", "", "path to config.json file")
	flag.StringVar(&ffiPath, "ffi", "", "path to ffi.json file")
	flag.StringVar(&apirefPath, "apiref", "", "path to apiref csv file")
	flag.StringVar(&dir, "dir", "", "base directory to generate from/to")
	flag.Parse()

	// Load config
	var err error
	cfg, err = assets.LoadConfig(configPath)
	if err != nil {
		log.Fatal("couldn't parse config file: ", err)
	}

	regJS, err = regexp.Compile(fmt.Sprintf(`"_%s([A-Z][A-Za-z_0-9]+)",`, cfg.Prefix))
	if err != nil {
		log.Fatal(err)
	}

	// Load the public API surface
	apiref, err := assets.LoadAPIRef(apirefPath)
	if err != nil {
		log.Fatal("couldn't load apiref file: ", err)
	}

	// Parse FFI
	var ffiEntries []*assets.FFIEntry
	b, err := os.ReadFile(ffiPath)
	if err != nil {
		log.Fatal("couldn't read ffi.json file: ", err)
	}
	err = json.Unmarshal(b, &ffiEntries)
	if err != nil {
		log.Fatal("couldn't unmarshal ffi file: ", err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	path = filepath.Join(path, dir)

	AllFunctions(apiref, ffiEntries)

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
	var category string

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
	sb.WriteString("<details open>\n")
	sb.WriteString("<summary>")
	sb.WriteString("<h2>" + strings.ToUpper(cfg.LibraryName) + "</h2>")
	sb.WriteString("</summary>\n")
	for _, fn := range functions {
		if fn.Category != category {
			if category != "" {
				// Close the previous details category
				sb.WriteString("</details>\n")
			}
			category = fn.Category
			if _, ok := collapsedCategories[fn.Category]; ok {
				sb.WriteString("<details>\n")
			} else {
				sb.WriteString("<details open>\n")
			}
			sb.WriteString("<summary>")
			sb.WriteString("<h3>" + fn.Category + "</h3>")
			sb.WriteString("</summary>\n\n")
			sb.WriteString("|Function|Desktop|WASM/js|\n")
			sb.WriteString("|:--|:--:|:--:|\n")
		}

		// A library with no quick reference page has no wiki at all, so there
		// is nothing to link its functions to.
		if cfg.QuickAPIRefURL != "" {
			fn.URL = fmt.Sprintf("https://wiki.libsdl.org/SDL3%s/%s", cfg.URLLibrarySuffix, fn.Name)
		}

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
		} else {
			js = ":question:"
		}
		name := fn.Name
		if fn.URL != "" {
			name = fmt.Sprintf("[%s](%s)", fn.Name, fn.URL)
		}
		sb.WriteString(fmt.Sprintf(
			"| %s | [%s](%s) | [%s](%s) |\n",
			name,
			desktop, urlDesktop,
			js, urlJS,
		))
	}
	sb.WriteString("</details>\n")
	sb.WriteString("</details>\n")

	var f *os.File
	if cfg.LibraryName == "sdl" {
		f, err = os.Create("COVERAGE.md")
		if err != nil {
			log.Fatal("couldn't create file: ", err)
		}
	} else {
		f, err = os.OpenFile("COVERAGE.md", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
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
