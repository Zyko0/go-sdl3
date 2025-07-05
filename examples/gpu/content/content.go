package content

import "embed"

//go:embed shaders/compiled images
var content embed.FS

func ReadFile(path string) ([]byte, error) {
	return content.ReadFile(path)
}
