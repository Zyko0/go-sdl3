package ttf

import (
	"unsafe"

	sdl "github.com/Zyko0/go-sdl3"
)

// Types

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
	EngineText        uintptr
}

type TextEngine struct {
	Version         uint32
	Userdata        uintptr
	CreateTextFunc  uintptr
	DestroyTextFunc uintptr
}

// Custom

type GlyphMetrics struct {
	MinX    int32
	MaxX    int32
	MinY    int32
	MaxY    int32
	Advance int32
}
