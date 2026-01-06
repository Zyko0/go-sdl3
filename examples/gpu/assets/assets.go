package assets

import "embed"

//go:embed shaders/compiled images
var assets embed.FS

func ReadFile(path string) ([]byte, error) {
	return assets.ReadFile(path)
}
