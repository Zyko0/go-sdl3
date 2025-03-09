package ttf

import "github.com/Zyko0/go-sdl3/sdl"

var (
	//puregogen:library path:windows=ttf.dll path:unix=ttf.so alias=ttf
	//puregogen:function symbol=TTF_TagToString
	iTagToString func(tag uint32, str *byte, size uintptr)

	//puregogen:function symbol=TTF_RenderText_Solid
	iRenderText_Solid func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_Solid_Wrapped
	iRenderText_Solid_Wrapped func(font *Font, str string, length uintptr, fg uint32, wrapLength int32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderGlyph_Solid
	iRenderGlyph_Solid func(font *Font, ch, fg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_Shaded
	iRenderText_Shaded func(font *Font, str string, length uintptr, fg, bg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_Shaded_Wrapped
	iRenderText_Shaded_Wrapped func(font *Font, str string, length uintptr, fg, bg uint32, wrapWidth int32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderGlyph_Shaded
	iRenderGlyph_Shaded func(font *Font, ch uint32, fg, bg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_Blended
	iRenderText_Blended func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_Blended_Wrapped
	iRenderText_Blended_Wrapped func(font *Font, str string, length uintptr, fg uint32, wrapWidth int32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderGlyph_Blended
	iRenderGlyph_Blended func(font *Font, ch, fg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_LCD
	iRenderText_LCD func(font *Font, str string, length uintptr, fg, bg uint32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderText_LCD_Wrapped
	iRenderText_LCD_Wrapped func(font *Font, str string, length uintptr, fg, bg uint32, wrapWidth int32) *sdl.Surface

	//puregogen:function symbol=TTF_RenderGlyph_LCD
	iRenderGlyph_LCD func(font *Font, ch, fg, bg uint32) *sdl.Surface
)
