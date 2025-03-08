//go:build js

package ttf

import (
	js "syscall/js"
	"unsafe"

	sdl "github.com/Zyko0/go-sdl3"
	internal "github.com/Zyko0/go-sdl3/internal"
)

func initialize() {
	iVersion = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_Version",
		)

		return int32(ret.Int())
	}

	iGetFreeTypeVersion = func(major *int32, minor *int32, patch *int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_major, ok := internal.GetJSPointer(major)
		if !ok {
			_major = internal.StackAlloc(int(unsafe.Sizeof(*major)))
		}
		_minor, ok := internal.GetJSPointer(minor)
		if !ok {
			_minor = internal.StackAlloc(int(unsafe.Sizeof(*minor)))
		}
		_patch, ok := internal.GetJSPointer(patch)
		if !ok {
			_patch = internal.StackAlloc(int(unsafe.Sizeof(*patch)))
		}
		js.Global().Get("Module").Call(
			"_SDL_GetFreeTypeVersion",
			_major,
			_minor,
			_patch,
		)
	}

	iGetHarfBuzzVersion = func(major *int32, minor *int32, patch *int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_major, ok := internal.GetJSPointer(major)
		if !ok {
			_major = internal.StackAlloc(int(unsafe.Sizeof(*major)))
		}
		_minor, ok := internal.GetJSPointer(minor)
		if !ok {
			_minor = internal.StackAlloc(int(unsafe.Sizeof(*minor)))
		}
		_patch, ok := internal.GetJSPointer(patch)
		if !ok {
			_patch = internal.StackAlloc(int(unsafe.Sizeof(*patch)))
		}
		js.Global().Get("Module").Call(
			"_SDL_GetHarfBuzzVersion",
			_major,
			_minor,
			_patch,
		)
	}

	iInit = func() bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_Init",
		)

		return internal.GetBool(ret)
	}

	iOpenFont = func(file string, ptsize float32) *Font {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		_ptsize := int32(ptsize)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenFont",
			_file,
			_ptsize,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}

	iOpenFontIO = func(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		_ptsize := int32(ptsize)
		ret := js.Global().Get("Module").Call(
			"_SDL_OpenFontIO",
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
			"_SDL_OpenFontWithProperties",
			_props,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}*/

	iCopyFont = func(existing_font *Font) *Font {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_existing_font, ok := internal.GetJSPointer(existing_font)
		if !ok {
			_existing_font = internal.StackAlloc(int(unsafe.Sizeof(*existing_font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CopyFont",
			_existing_font,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}

	iGetFontProperties = func(font *Font) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontProperties",
			_font,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iGetFontGeneration = func(font *Font) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontGeneration",
			_font,
		)

		return uint32(ret.Int())
	}

	iAddFallbackFont = func(font *Font, fallback *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_fallback, ok := internal.GetJSPointer(fallback)
		if !ok {
			_fallback = internal.StackAlloc(int(unsafe.Sizeof(*fallback)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_AddFallbackFont",
			_font,
			_fallback,
		)

		return internal.GetBool(ret)
	}

	iRemoveFallbackFont = func(font *Font, fallback *Font) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_fallback, ok := internal.GetJSPointer(fallback)
		if !ok {
			_fallback = internal.StackAlloc(int(unsafe.Sizeof(*fallback)))
		}
		js.Global().Get("Module").Call(
			"_SDL_RemoveFallbackFont",
			_font,
			_fallback,
		)
	}

	iClearFallbackFonts = func(font *Font) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		js.Global().Get("Module").Call(
			"_SDL_ClearFallbackFonts",
			_font,
		)
	}

	iSetFontSize = func(font *Font, ptsize float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ptsize := int32(ptsize)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontSize",
			_font,
			_ptsize,
		)

		return internal.GetBool(ret)
	}

	iSetFontSizeDPI = func(font *Font, ptsize float32, hdpi int32, vdpi int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ptsize := int32(ptsize)
		_hdpi := int32(hdpi)
		_vdpi := int32(vdpi)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontSizeDPI",
			_font,
			_ptsize,
			_hdpi,
			_vdpi,
		)

		return internal.GetBool(ret)
	}

	iGetFontSize = func(font *Font) float32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontSize",
			_font,
		)

		return float32(ret.Int())
	}

	iGetFontDPI = func(font *Font, hdpi *int32, vdpi *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_hdpi, ok := internal.GetJSPointer(hdpi)
		if !ok {
			_hdpi = internal.StackAlloc(int(unsafe.Sizeof(*hdpi)))
		}
		_vdpi, ok := internal.GetJSPointer(vdpi)
		if !ok {
			_vdpi = internal.StackAlloc(int(unsafe.Sizeof(*vdpi)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontDPI",
			_font,
			_hdpi,
			_vdpi,
		)

		return internal.GetBool(ret)
	}

	iSetFontStyle = func(font *Font, style FontStyleFlags) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_style := int32(style)
		js.Global().Get("Module").Call(
			"_SDL_SetFontStyle",
			_font,
			_style,
		)
	}

	iGetFontStyle = func(font *Font) FontStyleFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontStyle",
			_font,
		)

		return FontStyleFlags(ret.Int())
	}

	iSetFontOutline = func(font *Font, outline int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_outline := int32(outline)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontOutline",
			_font,
			_outline,
		)

		return internal.GetBool(ret)
	}

	iGetFontOutline = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontOutline",
			_font,
		)

		return int32(ret.Int())
	}

	iSetFontHinting = func(font *Font, hinting HintingFlags) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_hinting := int32(hinting)
		js.Global().Get("Module").Call(
			"_SDL_SetFontHinting",
			_font,
			_hinting,
		)
	}

	iGetNumFontFaces = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNumFontFaces",
			_font,
		)

		return int32(ret.Int())
	}

	iGetFontHinting = func(font *Font) HintingFlags {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontHinting",
			_font,
		)

		return HintingFlags(ret.Int())
	}

	iSetFontSDF = func(font *Font, enabled bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_enabled := internal.NewBoolean(enabled)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontSDF",
			_font,
			_enabled,
		)

		return internal.GetBool(ret)
	}

	iGetFontSDF = func(font *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontSDF",
			_font,
		)

		return internal.GetBool(ret)
	}

	iSetFontWrapAlignment = func(font *Font, align HorizontalAlignment) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_align := int32(align)
		js.Global().Get("Module").Call(
			"_SDL_SetFontWrapAlignment",
			_font,
			_align,
		)
	}

	iGetFontWrapAlignment = func(font *Font) HorizontalAlignment {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontWrapAlignment",
			_font,
		)

		return HorizontalAlignment(ret.Int())
	}

	iGetFontHeight = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontHeight",
			_font,
		)

		return int32(ret.Int())
	}

	iGetFontAscent = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontAscent",
			_font,
		)

		return int32(ret.Int())
	}

	iGetFontDescent = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontDescent",
			_font,
		)

		return int32(ret.Int())
	}

	iSetFontLineSkip = func(font *Font, lineskip int32) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_lineskip := int32(lineskip)
		js.Global().Get("Module").Call(
			"_SDL_SetFontLineSkip",
			_font,
			_lineskip,
		)
	}

	iGetFontLineSkip = func(font *Font) int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontLineSkip",
			_font,
		)

		return int32(ret.Int())
	}

	iSetFontKerning = func(font *Font, enabled bool) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_enabled := internal.NewBoolean(enabled)
		js.Global().Get("Module").Call(
			"_SDL_SetFontKerning",
			_font,
			_enabled,
		)
	}

	iGetFontKerning = func(font *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontKerning",
			_font,
		)

		return internal.GetBool(ret)
	}

	iFontIsFixedWidth = func(font *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_FontIsFixedWidth",
			_font,
		)

		return internal.GetBool(ret)
	}

	iFontIsScalable = func(font *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_FontIsScalable",
			_font,
		)

		return internal.GetBool(ret)
	}

	iGetFontFamilyName = func(font *Font) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontFamilyName",
			_font,
		)

		return internal.UTF8JSToString(ret)
	}

	iGetFontStyleName = func(font *Font) string {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontStyleName",
			_font,
		)

		return internal.UTF8JSToString(ret)
	}

	iSetFontDirection = func(font *Font, direction Direction) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_direction := int32(direction)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontDirection",
			_font,
			_direction,
		)

		return internal.GetBool(ret)
	}

	iGetFontDirection = func(font *Font) Direction {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontDirection",
			_font,
		)

		return Direction(ret.Int())
	}

	iStringToTag = func(string string) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_string := internal.StringOnJSStack(string)
		ret := js.Global().Get("Module").Call(
			"_SDL_StringToTag",
			_string,
		)

		return uint32(ret.Int())
	}

	iSetFontScript = func(font *Font, script uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_script := int32(script)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontScript",
			_font,
			_script,
		)

		return internal.GetBool(ret)
	}

	iGetFontScript = func(font *Font) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetFontScript",
			_font,
		)

		return uint32(ret.Int())
	}

	iGetGlyphScript = func(ch uint32) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlyphScript",
			_ch,
		)

		return uint32(ret.Int())
	}

	iSetFontLanguage = func(font *Font, language_bcp47 string) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_language_bcp47 := internal.StringOnJSStack(language_bcp47)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetFontLanguage",
			_font,
			_language_bcp47,
		)

		return internal.GetBool(ret)
	}

	iFontHasGlyph = func(font *Font, ch uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		ret := js.Global().Get("Module").Call(
			"_SDL_FontHasGlyph",
			_font,
			_ch,
		)

		return internal.GetBool(ret)
	}

	iGetGlyphImage = func(font *Font, ch uint32, image_type *ImageType) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		_image_type, ok := internal.GetJSPointer(image_type)
		if !ok {
			_image_type = internal.StackAlloc(int(unsafe.Sizeof(*image_type)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlyphImage",
			_font,
			_ch,
			_image_type,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iGetGlyphImageForIndex = func(font *Font, glyph_index uint32, image_type *ImageType) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_glyph_index := int32(glyph_index)
		_image_type, ok := internal.GetJSPointer(image_type)
		if !ok {
			_image_type = internal.StackAlloc(int(unsafe.Sizeof(*image_type)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlyphImageForIndex",
			_font,
			_glyph_index,
			_image_type,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iGetGlyphMetrics = func(font *Font, ch uint32, minx *int32, maxx *int32, miny *int32, maxy *int32, advance *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_ch := int32(ch)
		_minx, ok := internal.GetJSPointer(minx)
		if !ok {
			_minx = internal.StackAlloc(int(unsafe.Sizeof(*minx)))
		}
		_maxx, ok := internal.GetJSPointer(maxx)
		if !ok {
			_maxx = internal.StackAlloc(int(unsafe.Sizeof(*maxx)))
		}
		_miny, ok := internal.GetJSPointer(miny)
		if !ok {
			_miny = internal.StackAlloc(int(unsafe.Sizeof(*miny)))
		}
		_maxy, ok := internal.GetJSPointer(maxy)
		if !ok {
			_maxy = internal.StackAlloc(int(unsafe.Sizeof(*maxy)))
		}
		_advance, ok := internal.GetJSPointer(advance)
		if !ok {
			_advance = internal.StackAlloc(int(unsafe.Sizeof(*advance)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlyphMetrics",
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
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_previous_ch := int32(previous_ch)
		_ch := int32(ch)
		_kerning, ok := internal.GetJSPointer(kerning)
		if !ok {
			_kerning = internal.StackAlloc(int(unsafe.Sizeof(*kerning)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGlyphKerning",
			_font,
			_previous_ch,
			_ch,
			_kerning,
		)

		return internal.GetBool(ret)
	}

	iGetStringSize = func(font *Font, text string, length uintptr, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_text := internal.StringOnJSStack(text)
		_length := internal.NewBigInt(length)
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStringSize",
			_font,
			_text,
			_length,
			_w,
			_h,
		)

		return internal.GetBool(ret)
	}

	iGetStringSizeWrapped = func(font *Font, text string, length uintptr, wrap_width int32, w *int32, h *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_text := internal.StringOnJSStack(text)
		_length := internal.NewBigInt(length)
		_wrap_width := int32(wrap_width)
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetStringSizeWrapped",
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
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_text := internal.StringOnJSStack(text)
		_length := internal.NewBigInt(length)
		_max_width := int32(max_width)
		_measured_width, ok := internal.GetJSPointer(measured_width)
		if !ok {
			_measured_width = internal.StackAlloc(int(unsafe.Sizeof(*measured_width)))
		}
		_measured_length, ok := internal.GetJSPointer(measured_length)
		if !ok {
			_measured_length = internal.StackAlloc(int(unsafe.Sizeof(*measured_length)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_MeasureString",
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
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateSurfaceTextEngine",
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}

	iDrawSurfaceText = func(text *Text, x int32, y int32, surface *sdl.Surface) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_x := int32(x)
		_y := int32(y)
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_DrawSurfaceText",
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
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroySurfaceTextEngine",
			_engine,
		)
	}

	iCreateRendererTextEngine = func(renderer *sdl.Renderer) *TextEngine {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateRendererTextEngine",
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
			"_SDL_CreateRendererTextEngineWithProperties",
			_props,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}*/

	iDrawRendererText = func(text *Text, x float32, y float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_DrawRendererText",
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
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyRendererTextEngine",
			_engine,
		)
	}

	iCreateGPUTextEngine = func(device *sdl.GPUDevice) *TextEngine {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_device, ok := internal.GetJSPointer(device)
		if !ok {
			_device = internal.StackAlloc(int(unsafe.Sizeof(*device)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateGPUTextEngine",
			_device,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
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
			"_SDL_CreateGPUTextEngineWithProperties",
			_props,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}*/

	iGetGPUTextDrawData = func(text *Text) *GPUAtlasDrawSequence {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUTextDrawData",
			_text,
		)

		_obj := internal.NewObject[GPUAtlasDrawSequence](ret)
		return _obj
	}

	iDestroyGPUTextEngine = func(engine *TextEngine) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyGPUTextEngine",
			_engine,
		)
	}

	iSetGPUTextEngineWinding = func(engine *TextEngine, winding GPUTextEngineWinding) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		_winding := int32(winding)
		js.Global().Get("Module").Call(
			"_SDL_SetGPUTextEngineWinding",
			_engine,
			_winding,
		)
	}

	iGetGPUTextEngineWinding = func(engine *TextEngine) GPUTextEngineWinding {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetGPUTextEngineWinding",
			_engine,
		)

		return GPUTextEngineWinding(ret.Int())
	}

	iCreateText = func(engine *TextEngine, font *Font, text string, length uintptr) *Text {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_text := internal.StringOnJSStack(text)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_CreateText",
			_engine,
			_font,
			_text,
			_length,
		)

		_obj := internal.NewObject[Text](ret)
		return _obj
	}

	iGetTextProperties = func(text *Text) sdl.PropertiesID {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextProperties",
			_text,
		)

		return sdl.PropertiesID(ret.Int())
	}

	iSetTextEngine = func(text *Text, engine *TextEngine) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_engine, ok := internal.GetJSPointer(engine)
		if !ok {
			_engine = internal.StackAlloc(int(unsafe.Sizeof(*engine)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextEngine",
			_text,
			_engine,
		)

		return internal.GetBool(ret)
	}

	iGetTextEngine = func(text *Text) *TextEngine {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextEngine",
			_text,
		)

		_obj := internal.NewObject[TextEngine](ret)
		return _obj
	}

	iSetTextFont = func(text *Text, font *Font) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextFont",
			_text,
			_font,
		)

		return internal.GetBool(ret)
	}

	iGetTextFont = func(text *Text) *Font {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextFont",
			_text,
		)

		_obj := internal.NewObject[Font](ret)
		return _obj
	}

	iSetTextDirection = func(text *Text, direction Direction) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_direction := int32(direction)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextDirection",
			_text,
			_direction,
		)

		return internal.GetBool(ret)
	}

	iGetTextDirection = func(text *Text) Direction {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextDirection",
			_text,
		)

		return Direction(ret.Int())
	}

	iSetTextScript = func(text *Text, script uint32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_script := int32(script)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextScript",
			_text,
			_script,
		)

		return internal.GetBool(ret)
	}

	iGetTextScript = func(text *Text) uint32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextScript",
			_text,
		)

		return uint32(ret.Int())
	}

	iSetTextColor = func(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextColor",
			_text,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iSetTextColorFloat = func(text *Text, r float32, g float32, b float32, a float32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_r := int32(r)
		_g := int32(g)
		_b := int32(b)
		_a := int32(a)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextColorFloat",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextColor",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_r, ok := internal.GetJSPointer(r)
		if !ok {
			_r = internal.StackAlloc(int(unsafe.Sizeof(*r)))
		}
		_g, ok := internal.GetJSPointer(g)
		if !ok {
			_g = internal.StackAlloc(int(unsafe.Sizeof(*g)))
		}
		_b, ok := internal.GetJSPointer(b)
		if !ok {
			_b = internal.StackAlloc(int(unsafe.Sizeof(*b)))
		}
		_a, ok := internal.GetJSPointer(a)
		if !ok {
			_a = internal.StackAlloc(int(unsafe.Sizeof(*a)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextColorFloat",
			_text,
			_r,
			_g,
			_b,
			_a,
		)

		return internal.GetBool(ret)
	}

	iSetTextPosition = func(text *Text, x int32, y int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_x := int32(x)
		_y := int32(y)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextPosition",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_x, ok := internal.GetJSPointer(x)
		if !ok {
			_x = internal.StackAlloc(int(unsafe.Sizeof(*x)))
		}
		_y, ok := internal.GetJSPointer(y)
		if !ok {
			_y = internal.StackAlloc(int(unsafe.Sizeof(*y)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextPosition",
			_text,
			_x,
			_y,
		)

		return internal.GetBool(ret)
	}

	iSetTextWrapWidth = func(text *Text, wrap_width int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_wrap_width := int32(wrap_width)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextWrapWidth",
			_text,
			_wrap_width,
		)

		return internal.GetBool(ret)
	}

	iGetTextWrapWidth = func(text *Text, wrap_width *int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_wrap_width, ok := internal.GetJSPointer(wrap_width)
		if !ok {
			_wrap_width = internal.StackAlloc(int(unsafe.Sizeof(*wrap_width)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextWrapWidth",
			_text,
			_wrap_width,
		)

		return internal.GetBool(ret)
	}

	iSetTextWrapWhitespaceVisible = func(text *Text, visible bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_visible := internal.NewBoolean(visible)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextWrapWhitespaceVisible",
			_text,
			_visible,
		)

		return internal.GetBool(ret)
	}

	iTextWrapWhitespaceVisible = func(text *Text) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_TextWrapWhitespaceVisible",
			_text,
		)

		return internal.GetBool(ret)
	}

	iSetTextString = func(text *Text, string string, length uintptr) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_string := internal.StringOnJSStack(string)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_SetTextString",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_offset := int32(offset)
		_string := internal.StringOnJSStack(string)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_InsertTextString",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_string := internal.StringOnJSStack(string)
		_length := internal.NewBigInt(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_AppendTextString",
			_text,
			_string,
			_length,
		)

		return internal.GetBool(ret)
	}

	iDeleteTextString = func(text *Text, offset int32, length int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_offset := int32(offset)
		_length := int32(length)
		ret := js.Global().Get("Module").Call(
			"_SDL_DeleteTextString",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_w, ok := internal.GetJSPointer(w)
		if !ok {
			_w = internal.StackAlloc(int(unsafe.Sizeof(*w)))
		}
		_h, ok := internal.GetJSPointer(h)
		if !ok {
			_h = internal.StackAlloc(int(unsafe.Sizeof(*h)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextSize",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_offset := int32(offset)
		_substring, ok := internal.GetJSPointer(substring)
		if !ok {
			_substring = internal.StackAlloc(int(unsafe.Sizeof(*substring)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextSubString",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_line := int32(line)
		_substring, ok := internal.GetJSPointer(substring)
		if !ok {
			_substring = internal.StackAlloc(int(unsafe.Sizeof(*substring)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextSubStringForLine",
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
			"_SDL_GetTextSubStringsForRange",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_x := int32(x)
		_y := int32(y)
		_substring, ok := internal.GetJSPointer(substring)
		if !ok {
			_substring = internal.StackAlloc(int(unsafe.Sizeof(*substring)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetTextSubStringForPoint",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_substring, ok := internal.GetJSPointer(substring)
		if !ok {
			_substring = internal.StackAlloc(int(unsafe.Sizeof(*substring)))
		}
		_previous, ok := internal.GetJSPointer(previous)
		if !ok {
			_previous = internal.StackAlloc(int(unsafe.Sizeof(*previous)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetPreviousTextSubString",
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
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		_substring, ok := internal.GetJSPointer(substring)
		if !ok {
			_substring = internal.StackAlloc(int(unsafe.Sizeof(*substring)))
		}
		_next, ok := internal.GetJSPointer(next)
		if !ok {
			_next = internal.StackAlloc(int(unsafe.Sizeof(*next)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_GetNextTextSubString",
			_text,
			_substring,
			_next,
		)

		return internal.GetBool(ret)
	}

	iUpdateText = func(text *Text) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		ret := js.Global().Get("Module").Call(
			"_SDL_UpdateText",
			_text,
		)

		return internal.GetBool(ret)
	}

	iDestroyText = func(text *Text) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_text, ok := internal.GetJSPointer(text)
		if !ok {
			_text = internal.StackAlloc(int(unsafe.Sizeof(*text)))
		}
		js.Global().Get("Module").Call(
			"_SDL_DestroyText",
			_text,
		)
	}

	iCloseFont = func(font *Font) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		js.Global().Get("Module").Call(
			"_SDL_CloseFont",
			_font,
		)
	}

	iQuit = func() {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		js.Global().Get("Module").Call(
			"_SDL_Quit",
		)
	}

	iWasInit = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_SDL_WasInit",
		)

		return int32(ret.Int())
	}

	iTagToString = func(tag uint32, str *byte, size uintptr) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_tag := int32(tag)
		_str, ok := internal.GetJSPointer(str)
		if !ok {
			_str = internal.StackAlloc(int(unsafe.Sizeof(*str)))
		}
		_size := internal.NewBigInt(size)
		js.Global().Get("Module").Call(
			"_SDL_TagToString",
			_tag,
			_str,
			_size,
		)
	}

	iRenderText_Solid = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
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
			"_SDL_RenderText_Solid",
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
		internal.StackSave()
		defer internal.StackRestore()
		_font, ok := internal.GetJSPointer(font)
		if !ok {
			_font = internal.StackAlloc(int(unsafe.Sizeof(*font)))
		}
		_str := internal.StringOnJSStack(str)
		_length := internal.NewBigInt(length)
		_fg := int32(fg)
		_wrapLength := int32(wrapLength)
		ret := js.Global().Get("Module").Call(
			"_SDL_RenderText_Solid_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapLength,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
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
			"_SDL_RenderGlyph_Solid",
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
			"_SDL_RenderText_Shaded",
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
			"_SDL_RenderText_Shaded_Wrapped",
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
			"_SDL_RenderGlyph_Shaded",
			_font,
			_ch,
			_fg,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	iRenderText_Blended = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
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
			"_SDL_RenderText_Blended",
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
			"_SDL_RenderText_Blended_Wrapped",
			_font,
			_str,
			_length,
			_fg,
			_wrapWidth,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
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
			"_SDL_RenderGlyph_Blended",
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
			"_SDL_RenderText_LCD",
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
			"_SDL_RenderText_LCD_Wrapped",
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
			"_SDL_RenderGlyph_LCD",
			_font,
			_ch,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

}
