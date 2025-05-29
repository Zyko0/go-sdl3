package img

import "github.com/Zyko0/go-sdl3/internal"

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
