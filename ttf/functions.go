package ttf

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// TTF_Version - This function gets the version of the dynamically linked SDL_ttf library.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_Version)
func Version() int32 {
	return iVersion()
}

// TTF_GetFreeTypeVersion - Query the version of the FreeType library in use.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFreeTypeVersion)
func FreeTypeVersion() LibraryVersion {
	var version LibraryVersion

	iGetFreeTypeVersion(&version.Major, &version.Minor, &version.Patch)

	return version
}

// TTF_GetHarfBuzzVersion - Query the version of the HarfBuzz library in use.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetHarfBuzzVersion)
func HarfBuzzVersion() LibraryVersion {
	var version LibraryVersion

	iGetHarfBuzzVersion(&version.Major, &version.Minor, &version.Patch)

	return version
}

// TTF_Init - Initialize SDL_ttf.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_Init)
func Init() error {
	if !iInit() {
		return internal.LastErr()
	}

	return nil
}

// TTF_OpenFont - Create a font from a file, using a specified point size.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFont)
func OpenFont(file string, ptSize float32) (*Font, error) {
	font := iOpenFont(file, ptSize)
	if font == nil {
		return nil, internal.LastErr()
	}

	return font, nil
}

// TTF_OpenFontIO - Create a font from an SDL_IOStream, using a specified point size.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontIO)
func OpenFontIO(src *sdl.IOStream, closeio bool, ptSize float32) (*Font, error) {
	font := iOpenFontIO(src, closeio, ptSize)
	if font == nil {
		return nil, internal.LastErr()
	}

	return font, nil
}

// TTF_OpenFontWithProperties - Create a font with the specified properties.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_OpenFontWithProperties)
func OpenFontWithProperties(props sdl.PropertiesID) (*Font, error) {
	font := iOpenFontWithProperties(props)
	if font == nil {
		return nil, internal.LastErr()
	}

	return font, nil
}

// TTF_StringToTag - Convert from a 4 character string to a 32-bit tag.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_StringToTag)
func StringToTag(str string) uint32 {
	return iStringToTag(str)
}

// TTF_TagToString - Convert from a 32-bit tag to a 4 character string.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString)
func TagToString(tag uint32, size uintptr) string {
	// https://wiki.libsdl.org/SDL3_ttf/TTF_TagToString
	// Size should be at least 4
	if size < 4 {
		return ""
	}
	str := make([]byte, size+1)
	iTagToString(tag, unsafe.SliceData(str), size)

	return string(str)
}

// TTF_GetGlyphScript - Get the script used by a 32-bit codepoint.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphScript)
func GetGlyphScript(ch rune) (uint32, error) {
	script := iGetGlyphScript(uint32(ch))
	if script == 0 {
		return 0, internal.LastErr()
	}

	return script, nil
}

// TTF_CreateSurfaceTextEngine - Create a text engine for drawing text on SDL surfaces.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CreateSurfaceTextEngine)
func CreateSurfaceTextEngine() (*TextEngine, error) {
	engine := iCreateSurfaceTextEngine()
	if engine == nil {
		return nil, internal.LastErr()
	}

	return engine, nil
}

// TTF_CreateRendererTextEngine - Create a text engine for drawing text on an SDL renderer.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngine)
func CreateRendererTextEngine(renderer *sdl.Renderer) (*TextEngine, error) {
	engine := iCreateRendererTextEngine(renderer)
	if engine == nil {
		return nil, internal.LastErr()
	}

	return engine, nil
}

// TTF_CreateRendererTextEngineWithProperties - Create a text engine for drawing text on an SDL renderer, with the specified properties.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CreateRendererTextEngineWithProperties)
func CreateRendererTextEngineWithProperties(props sdl.PropertiesID) (*TextEngine, error) {
	engine := iCreateRendererTextEngineWithProperties(props)
	if engine == nil {
		return nil, internal.LastErr()
	}

	return engine, nil
}

// TTF_Quit - Deinitialize SDL_ttf.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_Quit)
func Quit() {
	iQuit()
}

// TTF_WasInit - Check if SDL_ttf is initialized.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_WasInit)
func WasInit() int32 {
	return iWasInit()
}
