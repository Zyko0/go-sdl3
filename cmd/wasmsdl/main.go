package main

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"go/version"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	//go:embed assets/index.html
	indexHTML []byte

	//go:embed assets/sdl.js
	sdlJS []byte

	//go:embed assets/sdl.wasm
	sdlWASM []byte
)

// Get the go version of the user module
func goVersion(buildDir string) (string, error) {
	cmd := exec.Command("go", "list", "-f", "go{{.Module.GoVersion}}")
	cmd.Dir = buildDir
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")
	out, err := cmd.Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return "", fmt.Errorf("couldn't get go version: %s", exitError.Stderr)
		}
		return "", err
	}
	v := strings.TrimSpace(string(out))
	if !version.IsValid(v) {
		return "", errors.New("could not determine go version")
	}

	return v, nil
}

func getWasmExecJS(buildDir string) ([]byte, error) {
	// Try finding it locally
	dir, ok := os.LookupEnv("GOROOT")
	if ok {
		// go1.23.0 and before were under "misc" folder
		for _, d := range []string{"lib", "misc"} {
			b, err := os.ReadFile(filepath.Join(dir, d, "wasm/wasm_exec.js"))
			if err == nil {
				return b, nil
			}
			if !errors.Is(err, os.ErrNotExist) {
				return nil, err
			}
		}
	}
	// Try fetching it online
	v, err := goVersion(buildDir)
	if err != nil {
		return nil, err
	}
	// Build path
	path := "lib/wasm/wasm_exec.js"
	if version.Compare(v, "go1.24.0") < 0 {
		path = "misc/wasm/wasm_exec.js"
	}
	// Fetch wasm_exec.js content
	resp, err := http.Get(fmt.Sprintf("https://go.googlesource.com/go/+/refs/tags/%s/%s?format=TEXT", v, path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(base64.NewDecoder(base64.StdEncoding, resp.Body))
	if err != nil {
		return nil, err
	}

	return content, nil
}

type Files map[string][]byte

func GetFiles(buildDir, htmlPath string) (Files, error) {
	// index.html content
	htmlContent := indexHTML
	if htmlPath != "" {
		b, err := os.ReadFile(htmlPath)
		if err != nil {
			log.Fatal("couldn't read index.html file: ", err)
		}
		htmlContent = b
	}
	// wasm_exec.js content
	wasmExecContent, err := getWasmExecJS(buildDir)
	if err != nil {
		log.Fatal("couldn't get wasm_exec.js content: ", err)
	}
	// Build
	wasmFileName := fmt.Sprintf("main_%d.wasm", time.Now().UnixNano())
	cmd := exec.Command("go", "build", "-o", wasmFileName)
	cmd.Dir = buildDir
	cmd.Env = append(os.Environ(), "GOOS=js", "GOARCH=wasm")

	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("%s%w", stderr.String(), err)
	}
	// Read wasm binary and remove it from disk
	wasmBytes, err := os.ReadFile(filepath.Join(buildDir, wasmFileName))
	if err != nil {
		return nil, err
	}
	if err := os.Remove(filepath.Join(buildDir, wasmFileName)); err != nil {
		fmt.Fprintf(os.Stderr, "warn: couldn't remove wasm file '%s': %v\n", wasmFileName, err)
	}

	return Files{
		"index.html":   htmlContent,
		"wasm_exec.js": wasmExecContent,
		"main.wasm":    wasmBytes,
		"sdl.js":       sdlJS,
		"sdl.wasm":     sdlWASM,
	}, nil
}

func main() {
	const usage = `
Usage of wasmsdl:

	wasmsdl <command> [arguments]

Commands:

	build           build the application on disk
	build-archive   build the application and pack all files in an archive
	serve           serve the application on localhost:8080
`
	var htmlPath string

	if len(os.Args) <= 1 {
		log.Fatal(usage)
	}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cmd := os.Args[1]
	cmdFlags := flag.NewFlagSet("wasmsdl "+cmd, flag.ContinueOnError)
	cmdFlags.StringVar(&htmlPath, "html", "", "optional index.html file")
	switch cmd {
	case "build":
		var out string

		cmdFlags.StringVar(&out, "out", "", "output directory")
		err := cmdFlags.Parse(os.Args[2:])
		if err != nil {
			os.Exit(1)
		}
		if cmdFlags.Arg(0) != "" {
			dir, _ = filepath.Abs(cmdFlags.Arg(0))
		}
		files, err := GetFiles(dir, htmlPath)
		if err != nil {
			log.Fatal(err)
		}
		if out == "" {
			out = fmt.Sprintf("out_%d", time.Now().UnixNano())
		}
		// Create out directory
		err = os.MkdirAll(out, 0700)
		if err != nil {
			log.Fatal("couldn't create output directory: ", err)
		}
		// Write files
		for name, data := range files {
			if err := os.WriteFile(filepath.Join(out, name), data, 0666); err != nil {
				log.Fatalf("couldn't write '%s': %v", name, err)
			}
		}
		fmt.Printf("Created directory '%s'\n", out)
	case "build-archive":
		var out string

		cmdFlags.StringVar(&out, "out", "", "output filename")
		err := cmdFlags.Parse(os.Args[2:])
		if err != nil {
			os.Exit(1)
		}
		if cmdFlags.Arg(0) != "" {
			dir, _ = filepath.Abs(cmdFlags.Arg(0))
		}
		files, err := GetFiles(dir, htmlPath)
		if err != nil {
			log.Fatal(err)
		}
		if out == "" {
			out = fmt.Sprintf("out_%d.zip", time.Now().UnixNano())
		}
		// Create out zip f
		f, err := os.Create(out)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		// Write files to archive
		zipw := zip.NewWriter(f)
		for name, data := range files {
			w, err := zipw.Create(name)
			if err != nil {
				log.Fatalf("couldn't add '%s' to archive: %v", name, err)
			}
			_, err = w.Write(data)
			if err != nil {
				log.Fatalf("couldn't write '%s' data to archive: %v", name, err)
			}
		}
		err = zipw.Flush()
		if err != nil {
			log.Fatalf("couldn't flush remaining data to archive: %v", err)
		}
		fmt.Printf("Created archive '%s'\n", out)
	case "serve":
		err := cmdFlags.Parse(os.Args[2:])
		if err != nil {
			os.Exit(1)
		}
		if cmdFlags.Arg(0) != "" {
			dir, _ = filepath.Abs(cmdFlags.Arg(0))
		}
		files, err := GetFiles(dir, htmlPath)
		if err != nil {
			log.Fatal(err)
		}

		var server http.Server

		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			name := path.Base(r.URL.Path)

			switch name {
			case "/", "", ".", "index.html":
				name = "index.html"
				fallthrough
			default:
				if data, ok := files[name]; ok {
					http.ServeContent(w, r, name, time.Now(), bytes.NewReader(data))
					return
				}
			}
			http.ServeFile(w, r, filepath.Join(".", r.URL.Path))
		})
		server.Handler = mux
		server.Addr = "localhost:8080"

		fmt.Println("Listening on localhost:8080...")
		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("couldn't listen and serve: %v\n", err)
		}
	default:
		fmt.Fprint(os.Stderr, usage)
		os.Exit(2)
	}

	fmt.Println("Done")
}
