//go:build js

package ttf

import (
	js "syscall/js"
	"unsafe"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

func initialize() {
	iVersion = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_TTF_Version",
		)
		return int32(ret.Int())
	}

	iGetFreeTypeVersion = func(major *int32, minor *int32, patch *int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_major := major // TODO: input, output or array?
		_minor := minor // TODO: input, output or array?
		_patch := patch // TODO: input, output or array?
		js.Global().Get("Module").Call(
			"_TTF_GetFreeTypeVersion",
			_major,
			_minor,
			_patch,
		)
	}

	iGetHarfBuzzVersion = func(major *int32, minor *int32, patch *int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_major := major // TODO: input, output or array?
		_minor := minor // TODO: input, output or array?
		_patch := patch // TODO: input, output or array?
		js.Global().Get("Module").Call(
			"_TTF_GetHarfBuzzVersion",
			_major,
			_minor,
			_patch,
		)
	}

	iInit = func() bool {
		ret := js.Global().Get("Module").Call(
			"_TTF_Init",
		)

		return internal.GetBool(ret)
	}

	iOpenFont = func(file string, ptsize float32) *Font {
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_ptsize := ptsize
		ret := js.Global().Get("Module").Call(
			"_TTF_OpenFont",
			_file,
			_ptsize,
		)
		return internal.NewObject[Font](ret)
	}

	iOpenFontIO = func(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			panic("nil stream")
		}
		_closeio := internal.NewBoolean(closeio)
		_ptsize := int32(ptsize)
		ret := js.Global().Get("Module").Call(
			"_TTF_OpenFontIO",
			_src,
			_closeio,
			_ptsize,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}

	/*iOpenFontWithProperties = func(props *sdl.PropertiesID) *Font {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props, ok := internal.GetJSPointer(props)
		if !ok {
			_props = internal.StackAlloc(int(unsafe.Sizeof(*props)))
		}
		ret := js.Global().Get("Module").Call(
			"_TTF_OpenFontWithProperties",
			_props,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}*/

	iCopyFont = func(existing_font *Font) *Font {
		_existing_font, _ := internal.GetJSPointer(existing_font)
		ret := js.Global().Get("Module").Call(
			"_TTF_CopyFont",
			_existing_font,
		)
		return internal.NewObject[Font](ret)
	}

	iGetFontProperties = func(font *Font) sdl.PropertiesID {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontProperties",
			_font,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iGetFontGeneration = func(font *Font) uint32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontGeneration",
			_font,
		)
		return uint32(ret.Int())
	}

	iAddFallbackFont = func(font *Font, fallback *Font) bool {
		_font, _ := internal.GetJSPointer(font)
		_fallback, _ := internal.GetJSPointer(fallback)
		ret := js.Global().Get("Module").Call(
			"_TTF_AddFallbackFont",
			_font,
			_fallback,
		)
		return internal.GetBool(ret)
	}

	iRemoveFallbackFont = func(font *Font, fallback *Font) {
		_font, _ := internal.GetJSPointer(font)
		_fallback, _ := internal.GetJSPointer(fallback)
		js.Global().Get("Module").Call(
			"_TTF_RemoveFallbackFont",
			_font,
			_fallback,
		)
	}

	iClearFallbackFonts = func(font *Font) {
		_font, _ := internal.GetJSPointer(font)
		js.Global().Get("Module").Call(
			"_TTF_ClearFallbackFonts",
			_font,
		)
	}

	iSetFontSize = func(font *Font, ptsize float32) bool {
		_font, _ := internal.GetJSPointer(font)
		_ptsize := ptsize
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontSize",
			_font,
			_ptsize,
		)
		return internal.GetBool(ret)
	}

	iSetFontSizeDPI = func(font *Font, ptsize float32, hdpi int32, vdpi int32) bool {
		_font, _ := internal.GetJSPointer(font)
		_ptsize := ptsize
		_hdpi := hdpi
		_vdpi := vdpi
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontSizeDPI",
			_font,
			_ptsize,
			_hdpi,
			_vdpi,
		)
		return internal.GetBool(ret)
	}

	iGetFontSize = func(font *Font) float32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontSize",
			_font,
		)
		return float32(ret.Float())
	}

	iGetFontDPI = func(font *Font, hdpi *int32, vdpi *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_hdpi := hdpi // TODO: input, output or array?
		_vdpi := vdpi // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontDPI",
			_font,
			_hdpi,
			_vdpi,
		)
		return internal.GetBool(ret)
	}

	iSetFontStyle = func(font *Font, style FontStyleFlags) {
		_font, _ := internal.GetJSPointer(font)
		_style := int32(style)
		js.Global().Get("Module").Call(
			"_TTF_SetFontStyle",
			_font,
			_style,
		)
	}

	iGetFontStyle = func(font *Font) FontStyleFlags {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontStyle",
			_font,
		)
		return FontStyleFlags(ret.Int())
	}

	iSetFontOutline = func(font *Font, outline int32) bool {
		_font, _ := internal.GetJSPointer(font)
		_outline := outline
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontOutline",
			_font,
			_outline,
		)
		return internal.GetBool(ret)
	}

	iGetFontOutline = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontOutline",
			_font,
		)
		return int32(ret.Int())
	}

	iSetFontHinting = func(font *Font, hinting HintingFlags) {
		_font, _ := internal.GetJSPointer(font)
		_hinting := int32(hinting)
		js.Global().Get("Module").Call(
			"_TTF_SetFontHinting",
			_font,
			_hinting,
		)
	}

	iGetNumFontFaces = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetNumFontFaces",
			_font,
		)
		return int32(ret.Int())
	}

	iGetFontHinting = func(font *Font) HintingFlags {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontHinting",
			_font,
		)
		return HintingFlags(ret.Int())
	}

	iSetFontSDF = func(font *Font, enabled bool) bool {
		_font, _ := internal.GetJSPointer(font)
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontSDF",
			_font,
			_enabled,
		)
		return internal.GetBool(ret)
	}

	iGetFontSDF = func(font *Font) bool {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontSDF",
			_font,
		)
		return internal.GetBool(ret)
	}

	iSetFontWrapAlignment = func(font *Font, align HorizontalAlignment) {
		_font, _ := internal.GetJSPointer(font)
		_align := int32(align)
		js.Global().Get("Module").Call(
			"_TTF_SetFontWrapAlignment",
			_font,
			_align,
		)
	}

	iGetFontWrapAlignment = func(font *Font) HorizontalAlignment {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontWrapAlignment",
			_font,
		)
		return HorizontalAlignment(ret.Int())
	}

	iGetFontHeight = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontHeight",
			_font,
		)
		return int32(ret.Int())
	}

	iGetFontAscent = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontAscent",
			_font,
		)
		return int32(ret.Int())
	}

	iGetFontDescent = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontDescent",
			_font,
		)
		return int32(ret.Int())
	}

	iSetFontLineSkip = func(font *Font, lineskip int32) {
		_font, _ := internal.GetJSPointer(font)
		_lineskip := lineskip
		js.Global().Get("Module").Call(
			"_TTF_SetFontLineSkip",
			_font,
			_lineskip,
		)
	}

	iGetFontLineSkip = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontLineSkip",
			_font,
		)
		return int32(ret.Int())
	}

	iSetFontKerning = func(font *Font, enabled bool) {
		_font, _ := internal.GetJSPointer(font)
		_enabled := internal.NewBoolean(enabled)
		js.Global().Get("Module").Call(
			"_TTF_SetFontKerning",
			_font,
			_enabled,
		)
	}

	iGetFontKerning = func(font *Font) bool {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontKerning",
			_font,
		)
		return internal.GetBool(ret)
	}

	iFontIsFixedWidth = func(font *Font) bool {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_FontIsFixedWidth",
			_font,
		)
		return internal.GetBool(ret)
	}

	iFontIsScalable = func(font *Font) bool {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_FontIsScalable",
			_font,
		)
		return internal.GetBool(ret)
	}

	iGetFontFamilyName = func(font *Font) string {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontFamilyName",
			_font,
		)
		return internal.UTF8JSToString(ret)
	}

	iGetFontStyleName = func(font *Font) string {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontStyleName",
			_font,
		)
		return internal.UTF8JSToString(ret)
	}

	iSetFontDirection = func(font *Font, direction Direction) bool {
		_font, _ := internal.GetJSPointer(font)
		_direction := int32(direction)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontDirection",
			_font,
			_direction,
		)
		return internal.GetBool(ret)
	}

	iGetFontDirection = func(font *Font) Direction {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontDirection",
			_font,
		)
		return Direction(ret.Int())
	}

	iStringToTag = func(string string) uint32 {
		internal.StackSave()
		defer internal.StackRestore()
		_string := internal.StringOnJSStack(string)
		ret := js.Global().Get("Module").Call(
			"_TTF_StringToTag",
			_string,
		)
		return uint32(ret.Int())
	}

	iSetFontScript = func(font *Font, script uint32) bool {
		_font, _ := internal.GetJSPointer(font)
		_script := int32(script)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontScript",
			_font,
			_script,
		)
		return internal.GetBool(ret)
	}

	iGetFontScript = func(font *Font) uint32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontScript",
			_font,
		)
		return uint32(ret.Int())
	}

	iGetGlyphScript = func(ch uint32) uint32 {
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGlyphScript",
			_ch,
		)
		return uint32(ret.Int())
	}

	iSetFontLanguage = func(font *Font, language_bcp47 string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_language_bcp47 := internal.StringOnJSStack(language_bcp47)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetFontLanguage",
			_font,
			_language_bcp47,
		)
		return internal.GetBool(ret)
	}

	iFontHasGlyph = func(font *Font, ch uint32) bool {
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_TTF_FontHasGlyph",
			_font,
			_ch,
		)
		return internal.GetBool(ret)
	}

	iGetGlyphImage = func(font *Font, ch uint32, image_type *ImageType) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_image_type := image_type // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGlyphImage",
			_font,
			_ch,
			_image_type,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iGetGlyphImageForIndex = func(font *Font, glyph_index uint32, image_type *ImageType) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_glyph_index := int32(glyph_index)
		_image_type := image_type // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGlyphImageForIndex",
			_font,
			_glyph_index,
			_image_type,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iGetGlyphMetrics = func(font *Font, ch uint32, minx *int32, maxx *int32, miny *int32, maxy *int32, advance *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_minx := minx       // TODO: input, output or array?
		_maxx := maxx       // TODO: input, output or array?
		_miny := miny       // TODO: input, output or array?
		_maxy := maxy       // TODO: input, output or array?
		_advance := advance // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGlyphMetrics",
			_font,
			_ch,
			_minx,
			_maxx,
			_miny,
			_maxy,
			_advance,
		)
		return internal.GetBool(ret)
	}

	iGetGlyphKerning = func(font *Font, previous_ch uint32, ch uint32, kerning *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_previous_ch := int32(previous_ch)
		_ch := int32(ch)
		_kerning := kerning // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGlyphKerning",
			_font,
			_previous_ch,
			_ch,
			_kerning,
		)
		return internal.GetBool(ret)
	}

	iGetStringSize = func(font *Font, text string, length uintptr, w *int32, h *int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			panic("nil font")
		}
		_text := internal.StringOnJSStack(text)
		_length := int32(length)
		_w := internal.StackAlloc(4)
		_h := internal.StackAlloc(4)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetStringSize",
			_font,
			_text,
			_length,
			_w,
			_h,
		)
		*w = int32(internal.GetValue(_w, "i32").Int())
		*h = int32(internal.GetValue(_h, "i32").Int())

		return internal.GetBool(ret)
	}

	iGetStringSizeWrapped = func(font *Font, text string, length uintptr, wrap_width int32, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_text := internal.StringOnJSStack(text)
		_length := length // TODO: size, buffer, userdata or callback?
		_wrap_width := wrap_width
		_w := w // TODO: input, output or array?
		_h := h // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetStringSizeWrapped",
			_font,
			_text,
			_length,
			_wrap_width,
			_w,
			_h,
		)
		return internal.GetBool(ret)
	}

	iMeasureString = func(font *Font, text string, length uintptr, max_width int32, measured_width *int32, measured_length *uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_text := internal.StringOnJSStack(text)
		_length := length // TODO: size, buffer, userdata or callback?
		_max_width := max_width
		_measured_width := measured_width   // TODO: input, output or array?
		_measured_length := measured_length // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_MeasureString",
			_font,
			_text,
			_length,
			_max_width,
			_measured_width,
			_measured_length,
		)
		return internal.GetBool(ret)
	}

	iCreateSurfaceTextEngine = func() *TextEngine {
		panic("not implemented on js")
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateSurfaceTextEngine",
		)
		return internal.NewObject[TextEngine](ret) // TODO: single value or array?
	}

	iDrawSurfaceText = func(text *Text, x int32, y int32, surface *sdl.Surface) bool {
		_text, _ := internal.GetJSPointer(text)
		_x := x
		_y := y
		_surface, _ := internal.GetJSPointer(surface)
		ret := js.Global().Get("Module").Call(
			"_TTF_DrawSurfaceText",
			_text,
			_x,
			_y,
			_surface,
		)
		return internal.GetBool(ret)
	}

	iDestroySurfaceTextEngine = func(engine *TextEngine) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine := engine // TODO: input, output or array?
		js.Global().Get("Module").Call(
			"_TTF_DestroySurfaceTextEngine",
			_engine,
		)
	}

	iCreateRendererTextEngine = func(renderer *sdl.Renderer) *TextEngine {
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			panic("nil renderer")
		}
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateRendererTextEngine",
			_renderer,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}

	/*iCreateRendererTextEngineWithProperties = func(props *sdl.PropertiesID) *TextEngine {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props, ok := internal.GetJSPointer(props)
		if !ok {
			_props = internal.StackAlloc(int(unsafe.Sizeof(*props)))
		}
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateRendererTextEngineWithProperties",
			_props,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}*/

	iDrawRendererText = func(text *Text, x float32, y float32) bool {
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			panic("nil text")
		}
		_x := int32(x) // TODO: find out if int32 or float32
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_TTF_DrawRendererText",
			_text,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iDestroyRendererTextEngine = func(engine *TextEngine) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine := engine // TODO: input, output or array?
		js.Global().Get("Module").Call(
			"_TTF_DestroyRendererTextEngine",
			_engine,
		)
	}

	iCreateGPUTextEngine = func(device *sdl.GPUDevice) *TextEngine {
		panic("not implemented on js")
		_device, _ := internal.GetJSPointer(device)
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateGPUTextEngine",
			_device,
		)
		return internal.NewObject[TextEngine](ret) // TODO: single value or array?
	}

	/*iCreateGPUTextEngineWithProperties = func(props *sdl.PropertiesID) *TextEngine {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_props, ok := internal.GetJSPointer(props)
		if !ok {
			_props = internal.StackAlloc(int(unsafe.Sizeof(*props)))
		}
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateGPUTextEngineWithProperties",
			_props,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}*/

	iGetGPUTextDrawData = func(text *Text) *GPUAtlasDrawSequence {
		panic("not implemented on js")
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGPUTextDrawData",
			_text,
		)
		return internal.NewObject[GPUAtlasDrawSequence](ret) // TODO: single value or array?
	}

	iDestroyGPUTextEngine = func(engine *TextEngine) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine := engine // TODO: input, output or array?
		js.Global().Get("Module").Call(
			"_TTF_DestroyGPUTextEngine",
			_engine,
		)
	}

	iSetGPUTextEngineWinding = func(engine *TextEngine, winding GPUTextEngineWinding) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine := engine // TODO: input, output or array?
		_winding := int32(winding)
		js.Global().Get("Module").Call(
			"_TTF_SetGPUTextEngineWinding",
			_engine,
			_winding,
		)
	}

	iGetGPUTextEngineWinding = func(engine *TextEngine) GPUTextEngineWinding {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine := engine // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetGPUTextEngineWinding",
			_engine,
		)
		return GPUTextEngineWinding(ret.Int())
	}

	iCreateText = func(engine *TextEngine, font *Font, text string, length uintptr) *Text {
		internal.StackSave()
		defer internal.StackRestore()
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			panic("nil engine")
		}
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			panic("nil font")
		}
		_text := internal.StringOnJSStack(text)
		_length := int32(length)
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateText",
			_engine,
			_font,
			_text,
			_length,
		)

		_obj := internal.NewObject[Text](ret)

		return _obj
	}

	iGetTextProperties = func(text *Text) sdl.PropertiesID {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextProperties",
			_text,
		)
		return sdl.PropertiesID(ret.Int())
	}

	iSetTextEngine = func(text *Text, engine *TextEngine) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_engine := engine // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextEngine",
			_text,
			_engine,
		)
		return internal.GetBool(ret)
	}

	iGetTextEngine = func(text *Text) *TextEngine {
		panic("not implemented on js")
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextEngine",
			_text,
		)
		return internal.NewObject[TextEngine](ret) // TODO: single value or array?
	}

	iSetTextFont = func(text *Text, font *Font) bool {
		_text, _ := internal.GetJSPointer(text)
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextFont",
			_text,
			_font,
		)
		return internal.GetBool(ret)
	}

	iGetTextFont = func(text *Text) *Font {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextFont",
			_text,
		)
		return internal.NewObject[Font](ret)
	}

	iSetTextDirection = func(text *Text, direction Direction) bool {
		_text, _ := internal.GetJSPointer(text)
		_direction := int32(direction)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextDirection",
			_text,
			_direction,
		)
		return internal.GetBool(ret)
	}

	iGetTextDirection = func(text *Text) Direction {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextDirection",
			_text,
		)
		return Direction(ret.Int())
	}

	iSetTextScript = func(text *Text, script uint32) bool {
		_text, _ := internal.GetJSPointer(text)
		_script := int32(script)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextScript",
			_text,
			_script,
		)
		return internal.GetBool(ret)
	}

	iGetTextScript = func(text *Text) uint32 {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextScript",
			_text,
		)
		return uint32(ret.Int())
	}

	iSetTextColor = func(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			panic("nil text")
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextColor",
			_text,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iSetTextColorFloat = func(text *Text, r float32, g float32, b float32, a float32) bool {
		_text, _ := internal.GetJSPointer(text)
		_r := r
		_g := g
		_b := b
		_a := a
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextColorFloat",
			_text,
			_r,
			_g,
			_b,
			_a,
		)
		return internal.GetBool(ret)
	}

	iGetTextColor = func(text *Text, r *uint8, g *uint8, b *uint8, a *uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_r := r // TODO: input, output or array?
		_g := g // TODO: input, output or array?
		_b := b // TODO: input, output or array?
		_a := a // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextColor",
			_text,
			_r,
			_g,
			_b,
			_a,
		)
		return internal.GetBool(ret)
	}

	iGetTextColorFloat = func(text *Text, r *float32, g *float32, b *float32, a *float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_r := r // TODO: input, output or array?
		_g := g // TODO: input, output or array?
		_b := b // TODO: input, output or array?
		_a := a // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextColorFloat",
			_text,
			_r,
			_g,
			_b,
			_a,
		)
		return internal.GetBool(ret)
	}

	iSetTextPosition = func(text *Text, x int32, y int32) bool {
		_text, _ := internal.GetJSPointer(text)
		_x := x
		_y := y
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextPosition",
			_text,
			_x,
			_y,
		)
		return internal.GetBool(ret)
	}

	iGetTextPosition = func(text *Text, x *int32, y *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_x := x // TODO: input, output or array?
		_y := y // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextPosition",
			_text,
			_x,
			_y,
		)
		return internal.GetBool(ret)
	}

	iSetTextWrapWidth = func(text *Text, wrap_width int32) bool {
		_text, _ := internal.GetJSPointer(text)
		_wrap_width := wrap_width
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextWrapWidth",
			_text,
			_wrap_width,
		)
		return internal.GetBool(ret)
	}

	iGetTextWrapWidth = func(text *Text, wrap_width *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_wrap_width := wrap_width // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextWrapWidth",
			_text,
			_wrap_width,
		)
		return internal.GetBool(ret)
	}

	iSetTextWrapWhitespaceVisible = func(text *Text, visible bool) bool {
		_text, _ := internal.GetJSPointer(text)
		_visible := internal.NewBoolean(visible)
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextWrapWhitespaceVisible",
			_text,
			_visible,
		)
		return internal.GetBool(ret)
	}

	iTextWrapWhitespaceVisible = func(text *Text) bool {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_TextWrapWhitespaceVisible",
			_text,
		)
		return internal.GetBool(ret)
	}

	iSetTextString = func(text *Text, string string, length uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_string := internal.StringOnJSStack(string)
		_length := length // TODO: size, buffer, userdata or callback?
		ret := js.Global().Get("Module").Call(
			"_TTF_SetTextString",
			_text,
			_string,
			_length,
		)
		return internal.GetBool(ret)
	}

	iInsertTextString = func(text *Text, offset int32, string string, length uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_offset := offset
		_string := internal.StringOnJSStack(string)
		_length := length // TODO: size, buffer, userdata or callback?
		ret := js.Global().Get("Module").Call(
			"_TTF_InsertTextString",
			_text,
			_offset,
			_string,
			_length,
		)
		return internal.GetBool(ret)
	}

	iAppendTextString = func(text *Text, string string, length uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_string := internal.StringOnJSStack(string)
		_length := length // TODO: size, buffer, userdata or callback?
		ret := js.Global().Get("Module").Call(
			"_TTF_AppendTextString",
			_text,
			_string,
			_length,
		)
		return internal.GetBool(ret)
	}

	iDeleteTextString = func(text *Text, offset int32, length int32) bool {
		_text, _ := internal.GetJSPointer(text)
		_offset := offset
		_length := length
		ret := js.Global().Get("Module").Call(
			"_TTF_DeleteTextString",
			_text,
			_offset,
			_length,
		)
		return internal.GetBool(ret)
	}

	iGetTextSize = func(text *Text, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_w := w // TODO: input, output or array?
		_h := h // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSize",
			_text,
			_w,
			_h,
		)
		return internal.GetBool(ret)
	}

	iGetTextSubString = func(text *Text, offset int32, substring *SubString) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_offset := offset
		_substring := substring // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSubString",
			_text,
			_offset,
			_substring,
		)
		return internal.GetBool(ret)
	}

	iGetTextSubStringForLine = func(text *Text, line int32, substring *SubString) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_line := line
		_substring := substring // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSubStringForLine",
			_text,
			_line,
			_substring,
		)
		return internal.GetBool(ret)
	}

	/*iGetTextSubStringsForRange = func(text *Text, offset int32, length int32, count *int32) **SubString {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_offset := int32(offset)
		_length := int32(length)
		_count, ok := internal.GetJSPointer(count)
		if !ok {
			_count = internal.StackAlloc(int(unsafe.Sizeof(*count)))
		}
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSubStringsForRange",
			_text,
			_offset,
			_length,
			_count,
		)

		_obj := internal.NewObject[SubString](ret)
		return _obj
	}*/

	iGetTextSubStringForPoint = func(text *Text, x int32, y int32, substring *SubString) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_x := x
		_y := y
		_substring := substring // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSubStringForPoint",
			_text,
			_x,
			_y,
			_substring,
		)
		return internal.GetBool(ret)
	}

	iGetPreviousTextSubString = func(text *Text, substring *SubString, previous *SubString) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_substring := substring // TODO: input, output or array?
		_previous := previous   // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetPreviousTextSubString",
			_text,
			_substring,
			_previous,
		)
		return internal.GetBool(ret)
	}

	iGetNextTextSubString = func(text *Text, substring *SubString, next *SubString) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_substring := substring // TODO: input, output or array?
		_next := next           // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetNextTextSubString",
			_text,
			_substring,
			_next,
		)
		return internal.GetBool(ret)
	}

	iUpdateText = func(text *Text) bool {
		_text, _ := internal.GetJSPointer(text)
		ret := js.Global().Get("Module").Call(
			"_TTF_UpdateText",
			_text,
		)
		return internal.GetBool(ret)
	}

	iDestroyText = func(text *Text) {
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			panic("nil text")
		}
		js.Global().Get("Module").Call(
			"_TTF_DestroyText",
			_text,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(text)))
	}

	iCloseFont = func(font *Font) {
		_font, _ := internal.GetJSPointer(font)
		js.Global().Get("Module").Call(
			"_TTF_CloseFont",
			_font,
		)
		internal.DeleteJSPointer(uintptr(unsafe.Pointer(font)))
	}

	iQuit = func() {
		js.Global().Get("Module").Call(
			"_TTF_Quit",
		)
	}

	iWasInit = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_TTF_WasInit",
		)
		return int32(ret.Int())
	}

	iTagToString = func(tag uint32, str *byte, size uintptr) {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		_str := str   // TODO: input, output or array?
		_size := size // TODO: size, buffer, userdata or callback?
		js.Global().Get("Module").Call(
			"_TTF_TagToString",
			_tag,
			_str,
			_size,
		)
	}

	iRenderText_Solid = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			panic("nil font")
		}
		_str := internal.StringOnJSStack(str)
		_length := int32(length)
		// wasm32 passes SDL_Color by value as a pointer to a copy.
		_fg := internal.CloneObjectToJSStack(&fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Solid",
			_font,
			_str,
			_length,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iRenderText_Solid_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapLength int32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_wrapLength := wrapLength
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Solid_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapLength,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	/*iRenderGlyph_Solid = func(font *Font, ch uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Solid",
			_font,
			_ch,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderText_Shaded = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_str := internal.StringOnJSStack(str)
		_length := internal.NewBigInt(length)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Shaded",
			_font,
			_str,
			_length,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderText_Shaded_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapWidth int32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_str := internal.StringOnJSStack(str)
		_length := internal.NewBigInt(length)
		_fg := int32(fg)
		_wrapWidth := int32(wrapWidth)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Shaded_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapWidth,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderGlyph_Shaded = func(font *Font, ch uint32, fg uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Shaded",
			_font,
			_ch,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	iRenderText_Blended = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			panic("nil font")
		}
		_str := internal.StringOnJSStack(str)
		_length := int32(length)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Blended",
			_font,
			_str,
			_length,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iRenderText_Blended_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapWidth int32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_wrapWidth := wrapWidth
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Blended_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapWidth,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	/*iRenderGlyph_Blended = func(font *Font, ch uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Blended",
			_font,
			_ch,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderText_LCD = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_str := internal.StringOnJSStack(str)
		_length := internal.NewBigInt(length)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_LCD",
			_font,
			_str,
			_length,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderText_LCD_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapWidth int32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_str := internal.StringOnJSStack(str)
		_length := internal.NewBigInt(length)
		_fg := int32(fg)
		_wrapWidth := int32(wrapWidth)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_LCD_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapWidth,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iRenderGlyph_LCD = func(font *Font, ch uint32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_LCD",
			_font,
			_ch,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	iOpenFontWithProperties = func(props sdl.PropertiesID) *Font {
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_TTF_OpenFontWithProperties",
			_props,
		)
		return internal.NewObject[Font](ret)
	}

	iGetFontWeight = func(font *Font) int32 {
		_font, _ := internal.GetJSPointer(font)
		ret := js.Global().Get("Module").Call(
			"_TTF_GetFontWeight",
			_font,
		)
		return int32(ret.Int())
	}

	iCreateRendererTextEngineWithProperties = func(props sdl.PropertiesID) *TextEngine {
		panic("not implemented on js")
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateRendererTextEngineWithProperties",
			_props,
		)
		return internal.NewObject[TextEngine](ret) // TODO: single value or array?
	}

	iCreateGPUTextEngineWithProperties = func(props sdl.PropertiesID) *TextEngine {
		panic("not implemented on js")
		_props := int32(props)
		ret := js.Global().Get("Module").Call(
			"_TTF_CreateGPUTextEngineWithProperties",
			_props,
		)
		return internal.NewObject[TextEngine](ret) // TODO: single value or array?
	}

	iGetTextSubStringsForRange = func(text *Text, offset int32, length int32, count *int32) **SubString {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, _ := internal.GetJSPointer(text)
		_offset := offset
		_length := length
		_count := count // TODO: input, output or array?
		ret := js.Global().Get("Module").Call(
			"_TTF_GetTextSubStringsForRange",
			_text,
			_offset,
			_length,
			_count,
		)
		return internal.NewObject[*SubString](ret) // TODO: single value or array?
	}

	iRenderGlyph_Solid = func(font *Font, ch uint32, fg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Solid",
			_font,
			_ch,
			_fg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderText_Shaded = func(font *Font, str string, length uintptr, fg uint32, bg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_bg := int32(bg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Shaded",
			_font,
			_str,
			_length,
			_fg,
			_bg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderText_Shaded_Wrapped = func(font *Font, str string, length uintptr, fg uint32, bg uint32, wrapWidth int32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_bg := int32(bg)
		_wrapWidth := wrapWidth
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_Shaded_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_bg,
			_wrapWidth,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderGlyph_Shaded = func(font *Font, ch uint32, fg uint32, bg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_fg := int32(fg)
		_bg := int32(bg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Shaded",
			_font,
			_ch,
			_fg,
			_bg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderGlyph_Blended = func(font *Font, ch uint32, fg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_fg := int32(fg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_Blended",
			_font,
			_ch,
			_fg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderText_LCD = func(font *Font, str string, length uintptr, fg uint32, bg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_bg := int32(bg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_LCD",
			_font,
			_str,
			_length,
			_fg,
			_bg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderText_LCD_Wrapped = func(font *Font, str string, length uintptr, fg uint32, bg uint32, wrapWidth int32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		internal.StackSave()
		defer internal.StackRestore()
		_font, _ := internal.GetJSPointer(font)
		_str := internal.StringOnJSStack(str)
		_length := length // TODO: size, buffer, userdata or callback?
		_fg := int32(fg)
		_bg := int32(bg)
		_wrapWidth := wrapWidth
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderText_LCD_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_bg,
			_wrapWidth,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iRenderGlyph_LCD = func(font *Font, ch uint32, fg uint32, bg uint32) *sdl.Surface {
		panic("not implemented on js")
		// TODO: this signature follows the native ABI, check the C prototype.
		_font, _ := internal.GetJSPointer(font)
		_ch := int32(ch)
		_fg := int32(fg)
		_bg := int32(bg)
		ret := js.Global().Get("Module").Call(
			"_TTF_RenderGlyph_LCD",
			_font,
			_ch,
			_fg,
			_bg,
		)
		return internal.NewObject[sdl.Surface](ret)
	}
}
