package img

import "github.com/Zyko0/go-sdl3/internal"

// Utils

// cStringArray returns strs as a NULL-terminated char ** array, which the
// bindings declare as *string.
func cStringArray(strs []string) []*byte {
	ptrs := make([]*byte, len(strs)+1)
	for i, s := range strs {
		ptrs[i] = internal.StringToPtr(s)
	}
	return ptrs
}

// Types

type Pointer = internal.Pointer

// IMG_Animation - Animated image support
// (https://wiki.libsdl.org/SDL3_image/IMG_Animation)
type Animation struct {
	W      int32
	H      int32
	count  int32
	frames Pointer
	delays Pointer
}
