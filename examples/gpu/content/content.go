package content

import "embed"

//go:embed shaders/compiled
var content embed.FS

func ReadFile(path string) ([]byte, error) {
	return content.ReadFile(path)
}
