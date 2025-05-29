package ttf

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// Utils

func colorToUint32(clr sdl.Color) uint32 {
	return *(*uint32)(unsafe.Pointer(&clr))
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

// TTF_TextEngine - A text engine used to create text objects.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_TextEngine)
type TextEngine struct {
	Version         uint32
	Userdata        Pointer
	CreateTextFunc  Pointer
	DestroyTextFunc Pointer
}

// TTF_TextData - Internal data for [TTF_Text](TTF_Text)
// (https://wiki.libsdl.org/SDL3_ttf/TTF_TextData)
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
