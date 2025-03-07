package img

import "github.com/Zyko0/go-sdl3/internal"

// Types

type Pointer = internal.Pointer

type Animation struct {
	W      int32
	H      int32
	count  int32
	frames Pointer
	delays Pointer
}
