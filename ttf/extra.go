package ttf

import sdl "github.com/Zyko0/go-sdl3"

// Types

// union type
type DrawOperation struct {
	Cmd  DrawCommand
	data [60]byte
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
	Version     uint32
	Userdata    uintptr
	CreateText  uintptr
	DestroyText uintptr
}
