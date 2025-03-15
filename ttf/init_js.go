//go:build js

package ttf

// We can just initialize everything here in js/wasm env
func init() {
	initialize()
}

// Path returns an empty string in js/wasm environment.
func Path() string {
	return ""
}

// LoadLibrary does nothing in js/wasm environment.
func LoadLibrary(path string) error {
	return nil
}

// CloseLibrary does nothing in js/wasm environment.
func CloseLibrary() error {
	return nil
}
