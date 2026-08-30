package testing

import (
	"go/constant"
	"go/token"
	"go/types"
	"sort"
	"strings"
	"testing"
)

// Test_EnumValues compares every generated constant against the value c2ffi
// recorded for the C enumerator it came from.
func Test_EnumValues(t *testing.T) {
	for _, library := range libraries {
		t.Run(library, func(t *testing.T) {
			cfg := readJSON[ffiConfig](t, assetPath("config", library))
			entries := readJSON[[]ffiEntry](t, assetPath("ffi", library))
			scope := loadPackage(t, modulePath+"/"+cfg.OutDir).Scope()

			var checked int
			var untested []string
			for _, e := range entries {
				if e.Tag != "enum" || !cfg.includes(e.Location) {
					continue
				}

				var found int
				for _, f := range e.Fields {
					// Enumerators carry the library prefix, the enum type name
					// does not take part: SDL_ASYNCIO_TASK_READ is ASYNCIO_TASK_READ.
					c, ok := scope.Lookup(strings.TrimPrefix(f.Name, cfg.Prefix)).(*types.Const)
					if !ok {
						continue
					}

					found++
					want := constant.MakeInt64(f.Value)
					if constant.Compare(c.Val(), token.NEQ, want) {
						t.Errorf("%s (Go %s): %s, C says %s", f.Name, c.Name(), c.Val(), want)
					}
				}

				// The generators emit an enum whole or not at all, so a partial
				// hit means an enumerator was renamed out from under the bindings.
				switch {
				case found == 0:
					untested = append(untested, e.Name)
				case found < len(e.Fields):
					t.Errorf("%s: %d of %d enumerators found", e.Name, found, len(e.Fields))
				}
				checked += found
			}

			sort.Strings(untested)
			t.Log("checked:", checked, "untested:", len(untested), untested)
		})
	}
}
