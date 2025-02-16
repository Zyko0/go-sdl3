package ttf

import (
	"unsafe"

	sdl "github.com/Zyko0/go-sdl3"
	"github.com/Zyko0/go-sdl3/internal"
)

// Text

func (text *Text) DrawSurface(x int32, y int32, surface *sdl.Surface) bool {
	panic("not implemented")
	return iDrawSurfaceText(text, x, y, surface)
}

func (text *Text) DrawRenderer(x float32, y float32) bool {
	panic("not implemented")
	return iDrawRendererText(text, x, y)
}

func (text *Text) GPUDrawData() *GPUAtlasDrawSequence {
	panic("not implemented")
	return iGetGPUTextDrawData(text)
}

func (text *Text) Properties() sdl.PropertiesID {
	return iGetTextProperties(text)
}

func (text *Text) SetEngine(engine *TextEngine) error {
	if !iSetTextEngine(text, engine) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Engine() (*TextEngine, error) {
	engine := iGetTextEngine(text)
	if engine == nil {
		return nil, internal.LastErr()
	}

	return engine, nil
}

func (text *Text) SetFont(font *Font) error {
	if !iSetTextFont(text, font) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Font() (*Font, error) {
	font := iGetTextFont(text)
	if font == nil {
		return nil, internal.LastErr()
	}

	return font, nil
}

func (text *Text) SetDirection(direction Direction) error {
	if !iSetTextDirection(text, direction) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Direction() Direction {
	return iGetTextDirection(text)
}

func (text *Text) SetScript(script uint32) error {
	if !iSetTextScript(text, script) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Script() uint32 {
	return iGetTextScript(text)
}

func (text *Text) SetColor(r, g, b, a uint8) error {
	if !iSetTextColor(text, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) SetColorFloat(r, g, b, a float32) error {
	if !iSetTextColorFloat(text, r, g, b, a) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Color() (sdl.Color, error) {
	var clr sdl.Color

	if !iGetTextColor(text, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

func (text *Text) ColorFloat() (sdl.FColor, error) {
	var clr sdl.FColor

	if !iGetTextColorFloat(text, &clr.R, &clr.G, &clr.B, &clr.A) {
		return clr, internal.LastErr()
	}

	return clr, nil
}

func (text *Text) SetPosition(x, y int32) {
	iSetTextPosition(text, x, y)
}

func (text *Text) Position() (int32, int32) {
	var x, y int32

	iGetTextPosition(text, &x, &y)

	return x, y
}

func (text *Text) SetWrapWidth(wrapWidth int32) error {
	if !iSetTextWrapWidth(text, wrapWidth) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) WrapWidth() (int32, error) {
	var wrapWidth int32

	if !iGetTextWrapWidth(text, &wrapWidth) {
		return 0, internal.LastErr()
	}

	return wrapWidth, nil
}

func (text *Text) SetWrapWhitespaceVisible(visible bool) error {
	if !iSetTextWrapWhitespaceVisible(text, visible) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) WrapWhitespaceVisible() bool {
	return iTextWrapWhitespaceVisible(text)
}

func (text *Text) SetString(str string) error {
	if !iSetTextString(text, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) InsertString(offset int32, str string) error {
	if !iInsertTextString(text, offset, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) AppendString(str string) error {
	if !iAppendTextString(text, str, uintptr(len(str))) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) DeleteString(offset int32, length int32) error {
	if !iDeleteTextString(text, offset, length) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Size() (int32, int32, error) {
	var w, h int32

	if !iGetTextSize(text, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

func (text *Text) SubString(offset int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubString(text, offset, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

func (text *Text) SubStringForLine(line int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubStringForLine(text, line, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

func (text *Text) SubStringsForRange(offset int32, length int32) ([]*SubString, error) {
	var count int32

	ptr := iGetTextSubStringsForRange(text, offset, length, &count)
	if ptr == nil {
		return nil, internal.LastErr()
	}
	defer internal.Free(uintptr(unsafe.Pointer(ptr)))

	return internal.ClonePtrSlice[*SubString](uintptr(unsafe.Pointer(ptr)), int(count)), nil
}

func (text *Text) SubStringForPoint(x, y int32) (*SubString, error) {
	var substring SubString

	if !iGetTextSubStringForPoint(text, x, y, &substring) {
		return nil, internal.LastErr()
	}

	return &substring, nil
}

func (text *Text) PreviousSubString(substring *SubString) (*SubString, error) {
	var previous SubString

	if !iGetPreviousTextSubString(text, substring, &previous) {
		return nil, internal.LastErr()
	}

	return &previous, nil
}

func (text *Text) NextSubString(substring *SubString) (*SubString, error) {
	var next SubString

	if !iGetPreviousTextSubString(text, substring, &next) {
		return nil, internal.LastErr()
	}

	return &next, nil
}

func (text *Text) Update() error {
	if !iUpdateText(text) {
		return internal.LastErr()
	}

	return nil
}

func (text *Text) Destroy() {
	iDestroyText(text)
}

// TextEngine

func (engine *TextEngine) DestroySurface() {
	iDestroySurfaceTextEngine(engine)
}

func (engine *TextEngine) DestroyRenderer() {
	iDestroyRendererTextEngine(engine)
}

func (engine *TextEngine) DestroyGPU() {
	iDestroyGPUTextEngine(engine)
}

func (engine *TextEngine) SetGPUWinding(winding GPUTextEngineWinding) {
	iSetGPUTextEngineWinding(engine, winding)
}

func (engine *TextEngine) GPUWinding() GPUTextEngineWinding {
	return iGetGPUTextEngineWinding(engine)
}

func (engine *TextEngine) CreateText(font *Font, text string) (*Text, error) {
	txt := iCreateText(engine, font, text, uintptr(len(text)))
	if txt == nil {
		return nil, internal.LastErr()
	}

	return txt, nil
}

// Font

func (font *Font) Copy() (*Font, error) {
	f := iCopyFont(font)
	if f == nil {
		return nil, internal.LastErr()
	}

	return f, nil
}

func (font *Font) Properties() sdl.PropertiesID {
	return iGetFontProperties(font)
}

func (font *Font) Generation() uint32 {
	return iGetFontGeneration(font)
}

func (font *Font) AddFallback(fallback *Font) error {
	if !iAddFallbackFont(font, fallback) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) RemoveFallback(fallback *Font) {
	iRemoveFallbackFont(font, fallback)
}

func (font *Font) ClearFallbacks() {
	iClearFallbackFonts(font)
}

func (font *Font) SetSize(ptsize float32) error {
	if !iSetFontSize(font, ptsize) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) SetSizeDPI(ptsize float32, hdpi, vdpi int32) error {
	if !iSetFontSizeDPI(font, ptsize, hdpi, vdpi) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) Size() (float32, error) {
	size := iGetFontSize(font)
	if size == 0 {
		return 0, internal.LastErr()
	}

	return size, nil
}

func (font *Font) DPI() (int32, int32, error) {
	var hdpi, vdpi int32

	if !iGetFontDPI(font, &hdpi, &vdpi) {
		return 0, 0, internal.LastErr()
	}

	return hdpi, vdpi, nil
}

func (font *Font) SetStyle(style FontStyleFlags) {
	iSetFontStyle(font, style)
}

func (font *Font) Style() FontStyleFlags {
	return iGetFontStyle(font)
}

func (font *Font) SetOutline(outline int32) error {
	if !iSetFontOutline(font, outline) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) Outline() int32 {
	return iGetFontOutline(font)
}

func (font *Font) SetHinting(hinting HintingFlags) {
	iSetFontHinting(font, hinting)
}

func (font *Font) NumFaces() int32 {
	return iGetNumFontFaces(font)
}

func (font *Font) Hinting() HintingFlags {
	return iGetFontHinting(font)
}

func (font *Font) SetWrapAlignment(align HorizontalAlignment) {
	iSetFontWrapAlignment(font, align)
}

func (font *Font) WrapAlignment() HorizontalAlignment {
	return iGetFontWrapAlignment(font)
}

func (font *Font) Height() int32 {
	return iGetFontHeight(font)
}

func (font *Font) Ascent() int32 {
	return iGetFontAscent(font)
}

func (font *Font) Descent() int32 {
	return iGetFontDescent(font)
}

func (font *Font) SetLineSkip(lineskip int32) {
	iSetFontLineSkip(font, lineskip)
}

func (font *Font) LineSkip() int32 {
	return iGetFontLineSkip(font)
}

func (font *Font) SetKerning(enabled bool) {
	iSetFontKerning(font, enabled)
}

func (font *Font) Kerning() bool {
	return iGetFontKerning(font)
}

func (font *Font) IsFixedWidth() bool {
	return iFontIsFixedWidth(font)
}

func (font *Font) FamilyName() string {
	return iGetFontFamilyName(font)
}

func (font *Font) StyleName() string {
	return iGetFontStyleName(font)
}

func (font *Font) SetDirection(direction Direction) error {
	if !iSetFontDirection(font, direction) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) Direction() Direction {
	return iGetFontDirection(font)
}

func (font *Font) SetScript(script uint32) error {
	if !iSetFontScript(font, script) {
		return internal.LastErr()
	}

	return nil
}

func (font *Font) Script() uint32 {
	return iGetFontScript(font)
}

func (font *Font) HasGlyph(ch uint32) bool {
	return iFontHasGlyph(font, ch)
}

func (font *Font) GlyphImage(ch uint32) (*sdl.Surface, ImageType, error) {
	var typ ImageType

	surface := iGetGlyphImage(font, ch, &typ)
	if surface == nil {
		return nil, 0, internal.LastErr()
	}

	return surface, typ, nil
}

func (font *Font) GlyphImageForIndex(glyphIndex uint32) (*sdl.Surface, ImageType, error) {
	var typ ImageType

	surface := iGetGlyphImageForIndex(font, glyphIndex, &typ)
	if surface == nil {
		return nil, 0, internal.LastErr()
	}

	return surface, typ, nil
}

func (font *Font) GlyphMetrics(ch uint32) (*GlyphMetrics, error) {
	var m GlyphMetrics
	
	if !iGetGlyphMetrics(font, ch, &m.MinX, &m.MaxX, &m.MinY, &m.MaxY, &m.Advance) {
		return nil, internal.LastErr()
	}

	return &m, nil
}

func (font *Font) StringSize(text string) (int32, int32, error) {
	var w, h int32

	if !iGetStringSize(font, text, uintptr(len(text)), &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return w, h, nil
}

func (font *Font) StringSizeWrapped(text string, wrapWidth int32) (int32, int32, error) {
	var w, h int32

	if !iGetStringSizeWrapped(font, text, uintptr(len(text)), wrapWidth, &w, &h) {
		return 0, 0, internal.LastErr()
	}

	return 0, 0, nil
}

func (font *Font) MeasureString(text string, maxWidth int32, measured_width *int32, measured_length *uintptr) bool {
	panic("not implemented")
	return iMeasureString(font, text, uintptr(len(text)), maxWidth, measured_width, measured_length)
}

func (font *Font) Close() {
	panic("not implemented")
	iCloseFont(font)
}
