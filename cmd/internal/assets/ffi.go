package assets

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
	TagHasPrefix    bool
}

func (e FFIEntry) PrefixedName(prefix string) string {
	if e.SymbolHasPrefix {
		return prefix + e.Name
	}

	return e.Name
}
