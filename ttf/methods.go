package ttf

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/sdl"
	"github.com/Zyko0/go-sdl3/internal"
)

// Text

// TTF_DrawSurfaceText - Draw text to an SDL surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DrawSurfaceText)
func (text *Text) DrawSurface(x, y int32, surface *sdl.Surface) error {
	if !iDrawSurfaceText(text, x, y, surface) {
		return internal.LastErr()
	}

	return nil
}

// TTF_DrawRendererText - Draw text to an SDL renderer.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DrawRendererText)
func (text *Text) DrawRenderer(x, y float32) error {
	if !iDrawRendererText(text, x, y) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetGPUTextDrawData - Get the geometry data needed for drawing the text.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextDrawData)
func (text *Text) GPUDrawData() (*GPUAtlasDrawSequence, error) {
	seq := iGetGPUTextDrawData(text)
	if seq == nil {
		return nil, internal.LastErr()
	}

	return seq, nil
}

// TTF_GetTextProperties - Get the properties associated with a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextProperties)
func (text *Text) Properties() sdl.PropertiesID {
	return iGetTextProperties(text)
}

// TTF_SetTextEngine - Set the text engine used by a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextEngine)
func (text *Text) SetEngine(engine *TextEngine) error {
	if !iSetTextEngine(text, engine) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextEngine - Get the text engine used by a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextEngine)
func (text *Text) Engine() (*TextEngine, error) {
	engine := iGetTextEngine(text)
	if engine == nil {
		return nil, internal.LastErr()
	}

	return engine, nil
}

// TTF_SetTextFont - Set the font used by a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextFont)
func (text *Text) SetFont(font *Font) error {
	if !iSetTextFont(text, font) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextFont - Get the font used by a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextFont)
func (text *Text) Font() (*Font, error) {
	font := iGetTextFont(text)
	if font == nil {
		return nil, internal.LastErr()
	}

	return font, nil
}

// TTF_SetTextDirection - Set the direction to be used for text shaping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextDirection)
func (text *Text) SetDirection(direction Direction) error {
	if !iSetTextDirection(text, direction) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextDirection - Get the direction to be used for text shaping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextDirection)
func (text *Text) Direction() Direction {
	return iGetTextDirection(text)
}

// TTF_SetTextScript - Set the script to be used for text shaping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextScript)
func (text *Text) SetScript(script uint32) error {
	if !iSetTextScript(text, script) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextScript - Get the script used for text shaping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextScript)
func (text *Text) Script() uint32 {
	return iGetTextScript(text)
}

// TTF_SetTextColor - Set the color of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColor)
func (text *Text) SetColor(r, g, b, a uint8) error {
	if !iSetTextColor(text, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// TTF_SetTextColorFloat - Set the color of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextColorFloat)
func (text *Text) SetColorFloat(r, g, b, a float32) error {
	if !iSetTextColorFloat(text, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextColor - Get the color of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColor)
func (text *Text) Color() (sdl.Color, error) {
	var clr sdl.Color

	if !iGetTextColor(text, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

// TTF_GetTextColorFloat - Get the color of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextColorFloat)
func (text *Text) ColorFloat() (sdl.FColor, error) {
	var clr sdl.FColor

	if !iGetTextColorFloat(text, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

// TTF_SetTextPosition - Set the position of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextPosition)
func (text *Text) SetPosition(x, y int32) {
	iSetTextPosition(text, x, y)
}

// TTF_GetTextPosition - Get the position of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextPosition)
func (text *Text) Position() (int32, int32) {
	var x, y int32

	iGetTextPosition(text, &x, &y)

	return x, y
}

// TTF_SetTextWrapWidth - Set whether wrapping is enabled on a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWidth)
func (text *Text) SetWrapWidth(wrapWidth int32) error {
	if !iSetTextWrapWidth(text, wrapWidth) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextWrapWidth - Get whether wrapping is enabled on a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextWrapWidth)
func (text *Text) WrapWidth() (int32, error) {
	var wrapWidth int32

	if !iGetTextWrapWidth(text, &wrapWidth) {
		return 0, internal.LastErr()
	}

	return wrapWidth, nil
}

// TTF_SetTextWrapWhitespaceVisible - Set whether whitespace should be visible when wrapping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextWrapWhitespaceVisible)
func (text *Text) SetWrapWhitespaceVisible(visible bool) error {
	if !iSetTextWrapWhitespaceVisible(text, visible) {
		return internal.LastErr()
	}

	return nil
}

// TTF_TextWrapWhitespaceVisible - Return whether whitespace is shown when wrapping a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_TextWrapWhitespaceVisible)
func (text *Text) WrapWhitespaceVisible() bool {
	return iTextWrapWhitespaceVisible(text)
}

// TTF_SetTextString - Set the UTF-8 text used by a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetTextString)
func (text *Text) SetString(str string) error {
	if !iSetTextString(text, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

// TTF_InsertTextString - Insert UTF-8 text into a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_InsertTextString)
func (text *Text) InsertString(offset int32, str string) error {
	if !iInsertTextString(text, offset, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

// TTF_AppendTextString - Append UTF-8 text to a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_AppendTextString)
func (text *Text) AppendString(str string) error {
	if !iAppendTextString(text, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

// TTF_DeleteTextString - Delete UTF-8 text from a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DeleteTextString)
func (text *Text) DeleteString(offset int32, length int32) error {
	if !iDeleteTextString(text, offset, length) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetTextSize - Get the size of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSize)
func (text *Text) Size() (int32, int32, error) {
	var w, h int32

	if !iGetTextSize(text, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// TTF_GetTextSubString - Get the substring of a text object that surrounds a text offset.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubString)
func (text *Text) SubString(offset int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubString(text, offset, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

// TTF_GetTextSubStringForLine - Get the substring of a text object that contains the given line.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForLine)
func (text *Text) SubStringForLine(line int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubStringForLine(text, line, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

// TTF_GetTextSubStringsForRange - Get the substrings of a text object that contain a range of text.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringsForRange)
func (text *Text) SubStringsForRange(offset int32, length int32) ([]*SubString, error) {
	var count int32

	ptr := iGetTextSubStringsForRange(text, offset, length, &count)
	if ptr == nil {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[*SubString](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

// TTF_GetTextSubStringForPoint - Get the portion of a text string that is closest to a point.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetTextSubStringForPoint)
func (text *Text) SubStringForPoint(x, y int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubStringForPoint(text, x, y, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

// TTF_GetPreviousTextSubString - Get the previous substring in a text object
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetPreviousTextSubString)
func (text *Text) PreviousSubString(substring *SubString) (*SubString, error) {
	var previous SubString

	if !iGetPreviousTextSubString(text, substring, &previous) {
		return nil, internal.LastErr()
	}

	return &previous, nil
}

// TTF_GetNextTextSubString - Get the previous substring in a text object
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetNextTextSubString)
func (text *Text) NextSubString(substring *SubString) (*SubString, error) {
	var next SubString

	if !iGetNextTextSubString(text, substring, &next) {
		return nil, internal.LastErr()
	}

	return &next, nil
}

// TTF_UpdateText - Update the layout of a text object.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_UpdateText)
func (text *Text) Update() error {
	if !iUpdateText(text) {
		return internal.LastErr()
	}

	return nil
}

// TTF_DestroyText - Destroy a text object created by a text engine.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyText)
func (text *Text) Destroy() {
	iDestroyText(text)
}

// TextEngine

// TTF_DestroySurfaceTextEngine - Destroy a text engine created for drawing text on SDL surfaces.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DestroySurfaceTextEngine)
func (engine *TextEngine) DestroySurface() {
	iDestroySurfaceTextEngine(engine)
}

// TTF_DestroyRendererTextEngine - Destroy a text engine created for drawing text on an SDL renderer.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyRendererTextEngine)
func (engine *TextEngine) DestroyRenderer() {
	iDestroyRendererTextEngine(engine)
}

// TTF_DestroyGPUTextEngine - Destroy a text engine created for drawing text with the SDL GPU API.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_DestroyGPUTextEngine)
func (engine *TextEngine) DestroyGPU() {
	iDestroyGPUTextEngine(engine)
}

// TTF_SetGPUTextEngineWinding - Sets the winding order of the vertices returned by TTF_GetGPUTextDrawData for a particular GPU text engine.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetGPUTextEngineWinding)
func (engine *TextEngine) SetGPUWinding(winding GPUTextEngineWinding) {
	iSetGPUTextEngineWinding(engine, winding)
}

// TTF_GetGPUTextEngineWinding - Get the winding order of the vertices returned by TTF_GetGPUTextDrawData for a particular GPU text engine
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGPUTextEngineWinding)
func (engine *TextEngine) GPUWinding() GPUTextEngineWinding {
	return iGetGPUTextEngineWinding(engine)
}

// TTF_CreateText - Create a text object from UTF-8 text and a text engine.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CreateText)
func (engine *TextEngine) CreateText(font *Font, text string) (*Text, error) {
	txt := iCreateText(engine, font, text, uintptr(len(text)))
	if txt == nil {
		return nil, internal.LastErr()
	}

	return txt, nil
}

// Font

// TTF_CopyFont - Create a copy of an existing font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CopyFont)
func (font *Font) Copy() (*Font, error) {
	f := iCopyFont(font)
	if f == nil {
		return nil, internal.LastErr()
	}

	return f, nil
}

// TTF_GetFontProperties - Get the properties associated with a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontProperties)
func (font *Font) Properties() sdl.PropertiesID {
	return iGetFontProperties(font)
}

// TTF_GetFontGeneration - Get the font generation.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontGeneration)
func (font *Font) Generation() uint32 {
	return iGetFontGeneration(font)
}

// TTF_AddFallbackFont - Add a fallback font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_AddFallbackFont)
func (font *Font) AddFallback(fallback *Font) error {
	if !iAddFallbackFont(font, fallback) {
		return internal.LastErr()
	}

	return nil
}

// TTF_RemoveFallbackFont - Remove a fallback font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RemoveFallbackFont)
func (font *Font) RemoveFallback(fallback *Font) {
	iRemoveFallbackFont(font, fallback)
}

// TTF_ClearFallbackFonts - Remove all fallback fonts.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_ClearFallbackFonts)
func (font *Font) ClearFallbacks() {
	iClearFallbackFonts(font)
}

// TTF_SetFontSize - Set a font's size dynamically.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSize)
func (font *Font) SetSize(ptsize float32) error {
	if !iSetFontSize(font, ptsize) {
		return internal.LastErr()
	}

	return nil
}

// TTF_SetFontSizeDPI - Set font size dynamically with target resolutions, in dots per inch.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSizeDPI)
func (font *Font) SetSizeDPI(ptsize float32, hdpi, vdpi int32) error {
	if !iSetFontSizeDPI(font, ptsize, hdpi, vdpi) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetFontSize - Get the size of a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSize)
func (font *Font) Size() (float32, error) {
	size := iGetFontSize(font)
	if size == 0 {
		return 0, internal.LastErr()
	}

	return size, nil
}

// TTF_GetFontDPI - Get font target resolutions, in dots per inch.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDPI)
func (font *Font) DPI() (int32, int32, error) {
	var hdpi, vdpi int32

	if !iGetFontDPI(font, &hdpi, &vdpi) {
		return 0, 0, internal.LastErr()
	}

	return hdpi, vdpi, nil
}

// TTF_SetFontStyle - Set a font's current style.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontStyle)
func (font *Font) SetStyle(style FontStyleFlags) {
	iSetFontStyle(font, style)
}

// TTF_GetFontStyle - Query a font's current style.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyle)
func (font *Font) Style() FontStyleFlags {
	return iGetFontStyle(font)
}

// TTF_SetFontOutline - Set a font's current outline.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontOutline)
func (font *Font) SetOutline(outline int32) error {
	if !iSetFontOutline(font, outline) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetFontOutline - Query a font's current outline.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontOutline)
func (font *Font) Outline() int32 {
	return iGetFontOutline(font)
}

// TTF_SetFontHinting - Set a font's current hinter setting.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontHinting)
func (font *Font) SetHinting(hinting HintingFlags) {
	iSetFontHinting(font, hinting)
}

// TTF_GetNumFontFaces - Query the number of faces of a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetNumFontFaces)
func (font *Font) NumFaces() int32 {
	return iGetNumFontFaces(font)
}

// TTF_GetFontHinting - Query a font's current FreeType hinter setting.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHinting)
func (font *Font) Hinting() HintingFlags {
	return iGetFontHinting(font)
}

// TTF_SetFontSDF - Enable Signed Distance Field rendering for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontSDF)
func (font *Font) SetSDF(enabled bool) error {
	if !iSetFontSDF(font, enabled) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetFontSDF - Query whether Signed Distance Field rendering is enabled for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontSDF)
func (font *Font) SDF() bool {
	return iGetFontSDF(font)
}

// TTF_SetFontWrapAlignment - Set a font's current wrap alignment option.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontWrapAlignment)
func (font *Font) SetWrapAlignment(align HorizontalAlignment) {
	iSetFontWrapAlignment(font, align)
}

// TTF_GetFontWrapAlignment - Query a font's current wrap alignment option.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontWrapAlignment)
func (font *Font) WrapAlignment() HorizontalAlignment {
	return iGetFontWrapAlignment(font)
}

// TTF_GetFontHeight - Query the total height of a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontHeight)
func (font *Font) Height() int32 {
	return iGetFontHeight(font)
}

// TTF_GetFontAscent - Query the offset from the baseline to the top of a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontAscent)
func (font *Font) Ascent() int32 {
	return iGetFontAscent(font)
}

// TTF_GetFontDescent - Query the offset from the baseline to the bottom of a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDescent)
func (font *Font) Descent() int32 {
	return iGetFontDescent(font)
}

// TTF_SetFontLineSkip - Set the spacing between lines of text for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLineSkip)
func (font *Font) SetLineSkip(lineskip int32) {
	iSetFontLineSkip(font, lineskip)
}

// TTF_GetFontLineSkip - Query the spacing between lines of text for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontLineSkip)
func (font *Font) LineSkip() int32 {
	return iGetFontLineSkip(font)
}

// TTF_SetFontKerning - Set if kerning is enabled for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontKerning)
func (font *Font) SetKerning(enabled bool) {
	iSetFontKerning(font, enabled)
}

// TTF_GetFontKerning - Query whether or not kerning is enabled for a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontKerning)
func (font *Font) Kerning() bool {
	return iGetFontKerning(font)
}

// TTF_FontIsFixedWidth - Query whether a font is fixed-width.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsFixedWidth)
func (font *Font) IsFixedWidth() bool {
	return iFontIsFixedWidth(font)
}

// TTF_FontIsScalable - Query whether a font is scalable or not.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_FontIsScalable)
func (font *Font) IsScalable() bool {
	return iFontIsScalable(font)
}

// TTF_GetFontFamilyName - Query a font's family name.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontFamilyName)
func (font *Font) FamilyName() string {
	return iGetFontFamilyName(font)
}

// TTF_GetFontStyleName - Query a font's style name.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontStyleName)
func (font *Font) StyleName() string {
	return iGetFontStyleName(font)
}

// TTF_SetFontDirection - Set the direction to be used for text shaping by a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontDirection)
func (font *Font) SetDirection(direction Direction) error {
	if !iSetFontDirection(font, direction) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetFontDirection - Get the direction to be used for text shaping by a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontDirection)
func (font *Font) Direction() Direction {
	return iGetFontDirection(font)
}

// TTF_SetFontScript - Set the script to be used for text shaping by a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontScript)
func (font *Font) SetScript(script uint32) error {
	if !iSetFontScript(font, script) {
		return internal.LastErr()
	}

	return nil
}

// TTF_GetFontScript - Get the script used for text shaping a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetFontScript)
func (font *Font) Script() uint32 {
	return iGetFontScript(font)
}

// TTF_SetFontLanguage - Set language to be used for text shaping by a font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_SetFontLanguage)
func (font *Font) SetLanguage(languageBCP47 string) error {
	if !iSetFontLanguage(font, languageBCP47) {
		return internal.LastErr()
	}

	return nil
}

// TTF_FontHasGlyph - Check whether a glyph is provided by the font for a UNICODE codepoint.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_FontHasGlyph)
func (font *Font) HasGlyph(ch uint32) bool {
	return iFontHasGlyph(font, ch)
}

// TTF_GetGlyphImage - Get the pixel image for a UNICODE codepoint.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImage)
func (font *Font) GlyphImage(ch uint32) (*sdl.Surface, ImageType, error) {
	var typ ImageType

	surface := iGetGlyphImage(font, ch, &typ)
	if surface == nil {
		return nil, 0, internal.LastErr()
	}

	return surface, typ, nil
}

// TTF_GetGlyphImageForIndex - Get the pixel image for a character index.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphImageForIndex)
func (font *Font) GlyphImageForIndex(glyphIndex uint32) (*sdl.Surface, ImageType, error) {
	var typ ImageType

	surface := iGetGlyphImageForIndex(font, glyphIndex, &typ)
	if surface == nil {
		return nil, 0, internal.LastErr()
	}

	return surface, typ, nil
}

// TTF_GetGlyphMetrics - Query the metrics (dimensions) of a font's glyph for a UNICODE codepoint.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphMetrics)
func (font *Font) GlyphMetrics(ch uint32) (*GlyphMetrics, error) {
	var m GlyphMetrics

	if !iGetGlyphMetrics(font, ch, &m.MinX, &m.MaxX, &m.MinY, &m.MaxY, &m.Advance) {
		return nil, internal.LastErr()
	}

	return &m, nil
}

// TTF_GetGlyphKerning - Query the kerning size between the glyphs of two UNICODE codepoints.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetGlyphKerning)
func (font *Font) GlyphKerning(previousCh, ch rune) (int32, error) {
	var kerning int32

	if !iGetGlyphKerning(font, uint32(previousCh), uint32(ch), &kerning) {
		return 0, internal.LastErr()
	}

	return kerning, nil
}

// TTF_GetStringSize - Calculate the dimensions of a rendered string of UTF-8 text.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSize)
func (font *Font) StringSize(text string) (int32, int32, error) {
	var w, h int32

	if !iGetStringSize(font, text, uintptr(len(text)), &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

// TTF_GetStringSizeWrapped - Calculate the dimensions of a rendered string of UTF-8 text.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_GetStringSizeWrapped)
func (font *Font) StringSizeWrapped(text string, wrapWidth int32) (int32, int32, error) {
	var w, h int32

	if !iGetStringSizeWrapped(font, text, uintptr(len(text)), wrapWidth, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return 0, 0, nil
}

// TTF_MeasureString - Calculate how much of a UTF-8 string will fit in a given width.
// Returns the measured_width, measured_length
// (https://wiki.libsdl.org/SDL3_ttf/TTF_MeasureString)
func (font *Font) MeasureString(text string, maxWidth int32) (int32, uintptr, error) {
	var width int32
	var length uintptr

	if !iMeasureString(font, text, uintptr(len(text)), maxWidth, &width, &length) {
		return 0, 0, internal.LastErr()
	}

	return width, length, nil
}

// TTF_RenderText_Solid - Render UTF-8 text at fast quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid)
func (font *Font) RenderTextSolid(text string, fg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderText_Solid(font, text, uintptr(len(text)), colorToUint32(fg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_Solid_Wrapped - Render word-wrapped UTF-8 text at fast quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Solid_Wrapped)
func (font *Font) RenderTextSolidWrapped(text string, fg sdl.Color, wrapLength int32) (*sdl.Surface, error) {
	surface := iRenderText_Solid_Wrapped(font, text, uintptr(len(text)), colorToUint32(fg), wrapLength)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderGlyph_Solid - Render a single 32-bit glyph at fast quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Solid)
func (font *Font) RenderGlyphSolid(glyph rune, fg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderGlyph_Solid(font, uint32(glyph), colorToUint32(fg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_Shaded - Render UTF-8 text at high quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded)
func (font *Font) RenderTextShaded(text string, fg, bg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderText_Shaded(font, text, uintptr(len(text)), colorToUint32(fg), colorToUint32(bg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_Shaded_Wrapped - Render word-wrapped UTF-8 text at high quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Shaded_Wrapped)
func (font *Font) RenderTextShadedWrapped(text string, fg, bg sdl.Color, wrapWidth int32) (*sdl.Surface, error) {
	surface := iRenderText_Shaded_Wrapped(font, text, uintptr(len(text)), colorToUint32(fg), colorToUint32(bg), wrapWidth)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderGlyph_Shaded - Render a single UNICODE codepoint at high quality to a new 8-bit surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Shaded)
func (font *Font) RenderGlyphShaded(glyph rune, fg, bg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderGlyph_Shaded(font, uint32(glyph), colorToUint32(fg), colorToUint32(bg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_Blended - Render UTF-8 text at high quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended)
func (font *Font) RenderTextBlended(text string, fg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderText_Blended(font, text, uintptr(len(text)), colorToUint32(fg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_Blended_Wrapped - Render word-wrapped UTF-8 text at high quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_Blended_Wrapped)
func (font *Font) RenderTextBlendedWrapped(text string, fg sdl.Color, wrapWidth int32) (*sdl.Surface, error) {
	surface := iRenderText_Blended_Wrapped(font, text, uintptr(len(text)), colorToUint32(fg), wrapWidth)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderGlyph_Blended - Render a single UNICODE codepoint at high quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_Blended)
func (font *Font) RenderGlyphBlended(glyph rune, fg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderGlyph_Blended(font, uint32(glyph), colorToUint32(fg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_LCD - Render UTF-8 text at LCD subpixel quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD)
func (font *Font) RenderTextLCD(text string, fg, bg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderText_LCD(font, text, uintptr(len(text)), colorToUint32(fg), colorToUint32(bg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderText_LCD_Wrapped - Render word-wrapped UTF-8 text at LCD subpixel quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderText_LCD_Wrapped)
func (font *Font) RenderTextLCDWrapped(text string, fg, bg sdl.Color, wrapWidth int32) (*sdl.Surface, error) {
	surface := iRenderText_LCD_Wrapped(font, text, uintptr(len(text)), colorToUint32(fg), colorToUint32(bg), wrapWidth)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_RenderGlyph_LCD - Render a single UNICODE codepoint at LCD subpixel quality to a new ARGB surface.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_RenderGlyph_LCD)
func (font *Font) RenderGlyphLCD(glyph rune, fg, bg sdl.Color) (*sdl.Surface, error) {
	surface := iRenderGlyph_LCD(font, uint32(glyph), colorToUint32(fg), colorToUint32(bg))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TTF_CloseFont - Dispose of a previously-created font.
// (https://wiki.libsdl.org/SDL3_ttf/TTF_CloseFont)
func (font *Font) Close() {
	iCloseFont(font)
}
