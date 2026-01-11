package assets

import (
	"encoding/json"
	"os"
)

type Config struct {
	OutDir                  string   `json:"out_dir"`
	LibraryName             string   `json:"library_name"`
	URLLibrarySuffix        string   `json:"url_library_suffix"`
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
	NoAutoStringFunctions   []string `json:"no_auto_string_functions"`
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
