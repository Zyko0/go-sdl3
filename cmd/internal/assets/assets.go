package assets

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

	SymbolHasPrefix bool
}
