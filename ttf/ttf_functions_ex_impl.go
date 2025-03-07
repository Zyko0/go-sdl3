//go:build windows || unix

package ttf

import (
	"unsafe"

	sdl "github.com/Zyko0/go-sdl3"
	puregogen "github.com/Zyko0/purego-gen"
	purego "github.com/ebitengine/purego"
)

var (
	// Symbols
	// ttf
	_addr_TTF_TagToString                uintptr
	_addr_TTF_RenderText_Solid           uintptr
	_addr_TTF_RenderText_Solid_Wrapped   uintptr
	_addr_TTF_RenderGlyph_Solid          uintptr
	_addr_TTF_RenderText_Shaded          uintptr
	_addr_TTF_RenderText_Shaded_Wrapped  uintptr
	_addr_TTF_RenderGlyph_Shaded         uintptr
	_addr_TTF_RenderText_Blended         uintptr
	_addr_TTF_RenderText_Blended_Wrapped uintptr
	_addr_TTF_RenderGlyph_Blended        uintptr
	_addr_TTF_RenderText_LCD             uintptr
	_addr_TTF_RenderText_LCD_Wrapped     uintptr
	_addr_TTF_RenderGlyph_LCD            uintptr
)

func initialize_ex() {
	var err error

	// Symbols ttf
	_addr_TTF_TagToString, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_TagToString")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_TagToString")
	}
	_addr_TTF_RenderText_Solid, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Solid")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Solid")
	}
	_addr_TTF_RenderText_Solid_Wrapped, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Solid_Wrapped")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Solid_Wrapped")
	}
	_addr_TTF_RenderGlyph_Solid, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderGlyph_Solid")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderGlyph_Solid")
	}
	_addr_TTF_RenderText_Shaded, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Shaded")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Shaded")
	}
	_addr_TTF_RenderText_Shaded_Wrapped, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Shaded_Wrapped")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Shaded_Wrapped")
	}
	_addr_TTF_RenderGlyph_Shaded, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderGlyph_Shaded")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderGlyph_Shaded")
	}
	_addr_TTF_RenderText_Blended, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Blended")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Blended")
	}
	_addr_TTF_RenderText_Blended_Wrapped, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_Blended_Wrapped")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_Blended_Wrapped")
	}
	_addr_TTF_RenderGlyph_Blended, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderGlyph_Blended")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderGlyph_Blended")
	}
	_addr_TTF_RenderText_LCD, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_LCD")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_LCD")
	}
	_addr_TTF_RenderText_LCD_Wrapped, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderText_LCD_Wrapped")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderText_LCD_Wrapped")
	}
	_addr_TTF_RenderGlyph_LCD, err = puregogen.OpenSymbol(_hnd_ttf, "TTF_RenderGlyph_LCD")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: TTF_RenderGlyph_LCD")
	}

	iTagToString = func(tag uint32, str *byte, size uintptr) {
		purego.SyscallN(_addr_TTF_TagToString, uintptr(tag), uintptr(unsafe.Pointer(str)), uintptr(size))
	}
	iRenderText_Solid = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Solid, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_Solid_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapLength int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Solid_Wrapped, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(wrapLength))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderGlyph_Solid = func(font *Font, ch uint32, fg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderGlyph_Solid, uintptr(unsafe.Pointer(font)), uintptr(ch), uintptr(fg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_Shaded = func(font *Font, str string, length uintptr, fg uint32, bg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Shaded, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(bg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_Shaded_Wrapped = func(font *Font, str string, length uintptr, fg uint32, bg uint32, wrapWidth int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Shaded_Wrapped, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(bg), uintptr(wrapWidth))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderGlyph_Shaded = func(font *Font, ch uint32, fg uint32, bg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderGlyph_Shaded, uintptr(unsafe.Pointer(font)), uintptr(ch), uintptr(fg), uintptr(bg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_Blended = func(font *Font, str string, length uintptr, fg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Blended, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_Blended_Wrapped = func(font *Font, str string, length uintptr, fg uint32, wrapWidth int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_Blended_Wrapped, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(wrapWidth))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderGlyph_Blended = func(font *Font, ch uint32, fg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderGlyph_Blended, uintptr(unsafe.Pointer(font)), uintptr(ch), uintptr(fg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_LCD = func(font *Font, str string, length uintptr, fg uint32, bg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_LCD, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(bg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderText_LCD_Wrapped = func(font *Font, str string, length uintptr, fg uint32, bg uint32, wrapWidth int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderText_LCD_Wrapped, uintptr(unsafe.Pointer(font)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(str))), uintptr(length), uintptr(fg), uintptr(bg), uintptr(wrapWidth))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
	iRenderGlyph_LCD = func(font *Font, ch uint32, fg uint32, bg uint32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_TTF_RenderGlyph_LCD, uintptr(unsafe.Pointer(font)), uintptr(ch), uintptr(fg), uintptr(bg))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		return __r0
	}
}
