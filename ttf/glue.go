package ttf

import (
	"unsafe"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/internal"
)

// Utils

func colorToUint32(clr sdl.Color) uint32 {
	return uint32(clr.R)<<24 |
		uint32(clr.G)<<16 |
		uint32(clr.B)<<8 |
		uint32(clr.A)
}

// Types

type Pointer = internal.Pointer

// union type
type DrawOperation struct {
	Cmd  DrawCommand
	data [60]byte
}

func (d *DrawOperation) CopyOperation() *CopyOperation {
	return (*CopyOperation)(unsafe.Pointer(d))
}

func (d *DrawOperation) FillOperation() *FillOperation {
	return (*FillOperation)(unsafe.Pointer(d))
}

type TextEngine struct {
	Version         uint32
	Userdata        Pointer
	CreateTextFunc  Pointer
	DestroyTextFunc Pointer
}

type TextData struct {
	Font              *Font
	Color             sdl.FColor
	NeedsLayoutUpdate bool
	Layout            *TextLayout
	X                 int32
	Y                 int32
	W                 int32
	H                 int32
	NumOps            int32
	Ops               *DrawOperation
	NumClusters       int32
	Clusters          *SubString
	Props             sdl.PropertiesID
	NeedsEngineUpdate bool
	Engine            *TextEngine
	EngineText        Pointer
}

// Custom

type GlyphMetrics struct {
	MinX    int32
	MaxX    int32
	MinY    int32
	MaxY    int32
	Advance int32
}

type LibraryVersion struct {
	Minor int32
	Major int32
	Patch int32
}
