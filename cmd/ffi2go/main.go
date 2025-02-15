package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/dave/jennifer/jen"
)

type Config struct {
	OutDir                  string   `json:"out_dir"`
	LibraryName             string   `json:"library_name"`
	Prefix                  string   `json:"prefix"`
	QuickAPIRefURL          string   `json:"quick_api_ref_url"`
	AllowedInclude          string   `json:"allowed_include"`
	IgnoredHeaders          []string `json:"ignored_headers"`
	IgnoredTypes            []string `json:"ignored_types"`
	IgnoredFunctions        []string `json:"ignored_functions"`
	AllowlistedFunctions    []string `json:"allowlisted_functions"`
	AllowlistedTypePrefixes []string `json:"allowlisted_type_prefixes"`
	BaseTypes               []string `json:"base_types"`
	SDLFreeFunctions        []string `json:"sdl_free_functions"`
}

type FFIEntry struct {
	Name         string      `json:"name"`
	Ns           int         `json:"ns"`
	Tag          string      `json:"tag"`
	Type         *FFIEntry   `json:"type"`
	Value        int         `json:"value"`
	Size         int         `json:"size"`
	Fields       []*FFIEntry `json:"fields"`
	StorageClass string      `json:"storage-class"`
	Variadic     bool        `json:"variadic"`
	Inline       bool        `json:"inline"`
	ReturnType   *FFIEntry   `json:"return-type"`
	Parameters   []*FFIEntry `json:"parameters"`
	BitOffset    int         `json:"bit-offset"`
	BitSize      int         `json:"bit-size"`
	BitAlignment int         `json:"bit-alignment"`
	ID           int         `json:"id"`
	Location     string      `json:"location"`

	symbolHasPrefix bool
}

var (
	cfg        Config
	ffiEntries []*FFIEntry
	apiRefCode string
)

const (
	genComment = "// Code generated by cmd/ffi2go. DO NOT EDIT.\n"
)

var (
	typesConversions = map[string]string{
		"_Bool":          "bool", // TODO: SDLBool
		"int":            "int",
		"Sint8":          "int8",
		"Sint16":         "int16",
		"Sint32":         "int32",
		"Sint64":         "int64",
		"Uint8":          "uint8",
		"Uint16":         "uint16",
		"Uint32":         "uint32",
		"Uint64":         "uint64",
		"long":           "int64",
		"float":          "float32",
		"double":         "float64",
		"unsigned-short": "uint16",
		"unsigned-int":   "uint",
		"unsigned-long":  "uint64",
		"long-long":      "int64", // TODO: sure?
		"ulong":          "uint64",
		"size_t":         "uintptr",
		"wchar_t":        "byte",
		"intptr_t":       "int64",
		"char":           "char", // Note: this allows replacing char* with string
		"unsigned-char":  "byte",
		"void":           "void",
		//"pointer":          "uintptr",
		"function-pointer": "uintptr",
	}
)

func sanitizeVarName(s string) string {
	var b strings.Builder
	nextCap := true
	for _, r := range s {
		if nextCap && r >= 'a' && r <= 'z' {
			r -= 32
		}
		nextCap = false
		if r == '_' {
			nextCap = true
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}

func sanitizeArgName(s string) string {
	switch {
	case s == "func":
		return "function"
	case s == "type":
		return "typ"
	default:
		return s
	}
}

func convType(s string, bitSize int) string {
	s = strings.ReplaceAll(s, ":", "")
	if ret, ok := typesConversions[s]; ok {
		switch ret {
		case "int", "uint":
			return fmt.Sprintf("%s%d", s, bitSize)
		default:
			return ret
		}
	}
	return s
}

func postConvertType(s string) string {
	switch s {
	case "*char":
		return "string"
	case "*void":
		return "uintptr"
	default:
		return s
	}
}

func extractType(e *FFIEntry) string {
	var typ string
	switch e.Tag {
	case ":array":
		typ = fmt.Sprintf("[%d]", e.Size) + extractType(e.Type)
	case ":function-pointer":
		typ = "uintptr"
	case ":pointer":
		typ = "*" + extractType(e.Type)
	case ":struct", "struct":
		typ = e.Name
	default:
		typ = convType(e.Tag, e.BitSize)
	}

	return postConvertType(typ)
}

func trimPrefix(e *FFIEntry) {
	e.Tag = strings.TrimPrefix(e.Tag, cfg.Prefix)
	e.Name = strings.TrimPrefix(e.Name, cfg.Prefix)
	if e.Type != nil {
		trimPrefix(e.Type)
	}
	if e.ReturnType != nil {
		trimPrefix(e.ReturnType)
	}
	for _, ee := range e.Fields {
		trimPrefix(ee)
	}
	for _, ee := range e.Parameters {
		trimPrefix(ee)
	}
}

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
			comment = name + " => " + comment
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
	var configPath, ffiPath string

	flag.StringVar(&configPath, "config", "", "path to config.json file")
	flag.StringVar(&ffiPath, "ffi", "", "path to ffi.json file")
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

	// Parse FFI
	b, err = os.ReadFile(ffiPath)
	if err != nil {
		log.Fatal("couldn't read ffi.json file: ", err)
	}
	err = json.Unmarshal(b, &ffiEntries)
	if err != nil {
		log.Fatal("couldn't unmarshal ffi file: ", err)
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("err: ", err)
	}
	dir = filepath.Join(dir, cfg.OutDir)

	// Init all types and functions from SDL3 quick reference
	AllTypesFromAPIRef()

	// Filter and sanitize sdl entries
	n := 0
	for _, e := range ffiEntries {
		if !strings.HasPrefix(e.Location, cfg.AllowedInclude) {
			continue
		}
		// Trim prefixes
		if strings.HasPrefix(e.Name, cfg.Prefix) {
			e.symbolHasPrefix = true
		}
		trimPrefix(e)
		var allowlisted bool
		if e.Tag == "function" {
			allowlisted = slices.Contains(cfg.AllowlistedFunctions, cfg.Prefix+e.Name)
			if !allowlisted {
				if slices.Contains(cfg.IgnoredFunctions, cfg.Prefix+e.Name) {
					continue
				}
				// Skip lower case functions
				if e.Name[0] >= 'a' && e.Name[0] <= 'z' {
					continue
				}
			}
		}
		// Skip some headers
		if !allowlisted {
			base, _, _ := strings.Cut(filepath.Base(e.Location), ":")
			if slices.Contains(cfg.IgnoredHeaders, base) {
				continue
			}
		}

		ffiEntries[n] = e
		n++
	}
	ffiEntries = ffiEntries[:n]

	// Register used types
	var regUsedTypes = func(e *FFIEntry) bool {
		var registration bool
		for _, f := range e.Fields {
			// Remove pointer or array potential prefix
			t := strings.ReplaceAll(extractType(f.Type), "*", "")
			if idx := strings.Index(t, "]"); idx != -1 {
				t = t[idx+1:]
			}
			if _, ok := uniqueTypes[cfg.Prefix+t]; !ok {
				uniqueTypes[cfg.Prefix+t] = struct{}{}
				registration = true
			}
		}
		return registration
	}
	var done bool
	for !done {
		done = true
		for _, e := range ffiEntries {
			_, ok := uniqueTypes[cfg.Prefix+e.Name]
			ok = ok || slices.Contains(cfg.BaseTypes, cfg.Prefix+e.Name)
			if ok {
				switch e.Tag {
				case "struct", "union":
					if registration := regUsedTypes(e); registration {
						done = false
					}
				case "typedef":
					if _, ok := uniqueTypes[cfg.Prefix+e.Tag]; !ok {
						t := strings.ReplaceAll(extractType(e.Type), "*", "")
						if idx := strings.Index(t, "]"); idx != -1 {
							t = t[idx+1:]
						}
						if _, ok := uniqueTypes[cfg.Prefix+t]; !ok {
							uniqueTypes[cfg.Prefix+t] = struct{}{}
							done = false
						}
					}
				}
			}
		}
	}

	fmt.Println("count entries:", len(ffiEntries))

	found := 0
	// Functions
	apifunc := map[string]struct{}{}
	f := jen.NewFile(cfg.LibraryName)
	f.Comment(genComment)
	f.Var().DefsFunc(func(g *jen.Group) {
		g.Comment(fmt.Sprintf("//puregogen:library path:windows=%s.dll path:unix=%s.so alias=%s",
			cfg.LibraryName, cfg.LibraryName, cfg.LibraryName,
		))
		for _, e := range ffiEntries {
			if e.Tag != "function" || e.Inline {
				continue
			}
			if ref, ok := uniqueAPIFunctions[cfg.Prefix+e.Name]; ok {
				found++
				apifunc[e.Name] = struct{}{}
				// Only add once
				fn := jen.Id("i" + e.Name).Func()
				fn.ParamsFunc(func(h *jen.Group) {
					for _, ee := range e.Parameters {
						h.Add(
							jen.Id(sanitizeArgName(ee.Name)).Id(extractType(ee.Type)),
						)
					}
				})
				if e.ReturnType != nil && e.ReturnType.Tag != ":void" {
					if slices.Contains(cfg.SDLFreeFunctions, cfg.Prefix+e.Name) {
						fn.Uintptr()
					} else {
						fn.Id(extractType(e.ReturnType))
					}
				}
				if e.symbolHasPrefix {
					g.Add(jen.Comment(
						"//puregogen:function symbol=" + cfg.Prefix + e.Name,
					))
				}
				g.Add(jen.Comment(
					"// " + ref.Description,
				))
				if slices.Contains(cfg.SDLFreeFunctions, cfg.Prefix+e.Name) {
					g.Add(jen.Comment(
						"// SDL_free() must be called on the returned pointer.",
					))
				}
				g.Add(fn.Line())
			}
		}
	})
	fmt.Println("count api functions:", len(uniqueAPIFunctions), len(apifunc))
	os.WriteFile(filepath.Join(dir, cfg.LibraryName+"_functions.gen.go"), []byte(f.GoString()), 0666)

	// Enums
	f = jen.NewFile(cfg.LibraryName)
	f.Comment(genComment)
	for _, e := range ffiEntries {
		if e.Tag != "enum" {
			continue
		}
		if _, ok := uniqueTypes[cfg.Prefix+e.Name]; !ok {
			continue
		}
		f.Type().Id(e.Name).Uint32()
		f.Const().DefsFunc(func(g *jen.Group) {
			for _, ee := range e.Fields {
				g.Add(
					jen.Id(ee.Name).Id(e.Name).Op("=").Lit(ee.Value),
				)
			}
		})
	}
	os.WriteFile(filepath.Join(dir, cfg.LibraryName+"_enums.gen.go"), []byte(f.GoString()), 0666)

	// Structs
	f = jen.NewFile(cfg.LibraryName)
	f.Comment(genComment)
	for _, e := range ffiEntries {
		if e.Tag != "struct" {
			continue
		}
		if slices.Contains(cfg.IgnoredTypes, e.Name) {
			continue
		}
		if _, ok := uniqueTypes[cfg.Prefix+e.Name]; !ok {
			continue
		}
		f.Type().Id(e.Name).StructFunc(func(g *jen.Group) {
			for _, ee := range e.Fields {
				g.Add(
					jen.Id(sanitizeVarName(ee.Name)).Id(extractType(ee.Type)),
				)
			}
		}).Line()
	}
	os.WriteFile(filepath.Join(dir, cfg.LibraryName+"_structs.gen.go"), []byte(f.GoString()), 0666)

	// Types
	f = jen.NewFile(cfg.LibraryName)
	f.Comment(genComment)
	f.Type().DefsFunc(func(g *jen.Group) {
		for _, e := range ffiEntries {
			switch {
			case e.Tag != "typedef",
				e.Type.Tag == ":struct", e.Type.Tag == "struct",
				e.Type.Tag == ":enum",
				e.Type.Tag == ":union", e.Type.Tag == "union",
				e.Type.Tag == ":function-pointer":
				continue
			}
			if _, ok := uniqueTypes[cfg.Prefix+e.Name]; !ok {
				continue
			}
			if _, handled := typesConversions[e.Name]; handled {
				continue
			}
			typ := extractType(e.Type)
			if e.Name == typ {
				continue
			}
			g.Id(e.Name).Id(typ)
		}
	})
	os.WriteFile(filepath.Join(dir, cfg.LibraryName+"_types.gen.go"), []byte(f.GoString()), 0666)
}
