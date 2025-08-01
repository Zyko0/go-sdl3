//go:build windows || unix

package img

import (
	sdl "github.com/Zyko0/go-sdl3/sdl"
	puregogen "github.com/Zyko0/purego-gen"
	purego "github.com/ebitengine/purego"
	"runtime"
	"unsafe"
)

// File generated by github.com/Zyko0/purego-gen. DO NOT EDIT.

var (
	// Library handles
	_hnd_img uintptr
	// Symbols
	// img
	_addr_IMG_Version                  uintptr
	_addr_IMG_LoadTyped_IO             uintptr
	_addr_IMG_Load                     uintptr
	_addr_IMG_Load_IO                  uintptr
	_addr_IMG_LoadTexture              uintptr
	_addr_IMG_LoadTexture_IO           uintptr
	_addr_IMG_LoadTextureTyped_IO      uintptr
	_addr_IMG_isAVIF                   uintptr
	_addr_IMG_isICO                    uintptr
	_addr_IMG_isCUR                    uintptr
	_addr_IMG_isBMP                    uintptr
	_addr_IMG_isGIF                    uintptr
	_addr_IMG_isJPG                    uintptr
	_addr_IMG_isJXL                    uintptr
	_addr_IMG_isLBM                    uintptr
	_addr_IMG_isPCX                    uintptr
	_addr_IMG_isPNG                    uintptr
	_addr_IMG_isPNM                    uintptr
	_addr_IMG_isSVG                    uintptr
	_addr_IMG_isQOI                    uintptr
	_addr_IMG_isTIF                    uintptr
	_addr_IMG_isXCF                    uintptr
	_addr_IMG_isXPM                    uintptr
	_addr_IMG_isXV                     uintptr
	_addr_IMG_isWEBP                   uintptr
	_addr_IMG_LoadAVIF_IO              uintptr
	_addr_IMG_LoadICO_IO               uintptr
	_addr_IMG_LoadCUR_IO               uintptr
	_addr_IMG_LoadBMP_IO               uintptr
	_addr_IMG_LoadGIF_IO               uintptr
	_addr_IMG_LoadJPG_IO               uintptr
	_addr_IMG_LoadJXL_IO               uintptr
	_addr_IMG_LoadLBM_IO               uintptr
	_addr_IMG_LoadPCX_IO               uintptr
	_addr_IMG_LoadPNG_IO               uintptr
	_addr_IMG_LoadPNM_IO               uintptr
	_addr_IMG_LoadSVG_IO               uintptr
	_addr_IMG_LoadQOI_IO               uintptr
	_addr_IMG_LoadTGA_IO               uintptr
	_addr_IMG_LoadTIF_IO               uintptr
	_addr_IMG_LoadXCF_IO               uintptr
	_addr_IMG_LoadXPM_IO               uintptr
	_addr_IMG_LoadXV_IO                uintptr
	_addr_IMG_LoadWEBP_IO              uintptr
	_addr_IMG_LoadSizedSVG_IO          uintptr
	_addr_IMG_ReadXPMFromArray         uintptr
	_addr_IMG_ReadXPMFromArrayToRGB888 uintptr
	_addr_IMG_SaveAVIF                 uintptr
	_addr_IMG_SaveAVIF_IO              uintptr
	_addr_IMG_SavePNG                  uintptr
	_addr_IMG_SavePNG_IO               uintptr
	_addr_IMG_SaveJPG                  uintptr
	_addr_IMG_SaveJPG_IO               uintptr
	_addr_IMG_LoadAnimation            uintptr
	_addr_IMG_LoadAnimation_IO         uintptr
	_addr_IMG_LoadAnimationTyped_IO    uintptr
	_addr_IMG_FreeAnimation            uintptr
	_addr_IMG_LoadGIFAnimation_IO      uintptr
	_addr_IMG_LoadWEBPAnimation_IO     uintptr
)

func initialize() {
	var err error

	// Symbols img
	_addr_IMG_Version, err = puregogen.OpenSymbol(_hnd_img, "IMG_Version")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Version")
	}
	_addr_IMG_LoadTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTyped_IO")
	}
	_addr_IMG_Load, err = puregogen.OpenSymbol(_hnd_img, "IMG_Load")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Load")
	}
	_addr_IMG_Load_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_Load_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_Load_IO")
	}
	_addr_IMG_LoadTexture, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTexture")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTexture")
	}
	_addr_IMG_LoadTexture_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTexture_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTexture_IO")
	}
	_addr_IMG_LoadTextureTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTextureTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTextureTyped_IO")
	}
	_addr_IMG_isAVIF, err = puregogen.OpenSymbol(_hnd_img, "IMG_isAVIF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isAVIF")
	}
	_addr_IMG_isICO, err = puregogen.OpenSymbol(_hnd_img, "IMG_isICO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isICO")
	}
	_addr_IMG_isCUR, err = puregogen.OpenSymbol(_hnd_img, "IMG_isCUR")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isCUR")
	}
	_addr_IMG_isBMP, err = puregogen.OpenSymbol(_hnd_img, "IMG_isBMP")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isBMP")
	}
	_addr_IMG_isGIF, err = puregogen.OpenSymbol(_hnd_img, "IMG_isGIF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isGIF")
	}
	_addr_IMG_isJPG, err = puregogen.OpenSymbol(_hnd_img, "IMG_isJPG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isJPG")
	}
	_addr_IMG_isJXL, err = puregogen.OpenSymbol(_hnd_img, "IMG_isJXL")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isJXL")
	}
	_addr_IMG_isLBM, err = puregogen.OpenSymbol(_hnd_img, "IMG_isLBM")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isLBM")
	}
	_addr_IMG_isPCX, err = puregogen.OpenSymbol(_hnd_img, "IMG_isPCX")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isPCX")
	}
	_addr_IMG_isPNG, err = puregogen.OpenSymbol(_hnd_img, "IMG_isPNG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isPNG")
	}
	_addr_IMG_isPNM, err = puregogen.OpenSymbol(_hnd_img, "IMG_isPNM")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isPNM")
	}
	_addr_IMG_isSVG, err = puregogen.OpenSymbol(_hnd_img, "IMG_isSVG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isSVG")
	}
	_addr_IMG_isQOI, err = puregogen.OpenSymbol(_hnd_img, "IMG_isQOI")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isQOI")
	}
	_addr_IMG_isTIF, err = puregogen.OpenSymbol(_hnd_img, "IMG_isTIF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isTIF")
	}
	_addr_IMG_isXCF, err = puregogen.OpenSymbol(_hnd_img, "IMG_isXCF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isXCF")
	}
	_addr_IMG_isXPM, err = puregogen.OpenSymbol(_hnd_img, "IMG_isXPM")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isXPM")
	}
	_addr_IMG_isXV, err = puregogen.OpenSymbol(_hnd_img, "IMG_isXV")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isXV")
	}
	_addr_IMG_isWEBP, err = puregogen.OpenSymbol(_hnd_img, "IMG_isWEBP")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_isWEBP")
	}
	_addr_IMG_LoadAVIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAVIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAVIF_IO")
	}
	_addr_IMG_LoadICO_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadICO_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadICO_IO")
	}
	_addr_IMG_LoadCUR_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadCUR_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadCUR_IO")
	}
	_addr_IMG_LoadBMP_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadBMP_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadBMP_IO")
	}
	_addr_IMG_LoadGIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadGIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadGIF_IO")
	}
	_addr_IMG_LoadJPG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadJPG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadJPG_IO")
	}
	_addr_IMG_LoadJXL_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadJXL_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadJXL_IO")
	}
	_addr_IMG_LoadLBM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadLBM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadLBM_IO")
	}
	_addr_IMG_LoadPCX_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPCX_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPCX_IO")
	}
	_addr_IMG_LoadPNG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPNG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPNG_IO")
	}
	_addr_IMG_LoadPNM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadPNM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadPNM_IO")
	}
	_addr_IMG_LoadSVG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadSVG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadSVG_IO")
	}
	_addr_IMG_LoadQOI_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadQOI_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadQOI_IO")
	}
	_addr_IMG_LoadTGA_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTGA_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTGA_IO")
	}
	_addr_IMG_LoadTIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadTIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadTIF_IO")
	}
	_addr_IMG_LoadXCF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXCF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXCF_IO")
	}
	_addr_IMG_LoadXPM_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXPM_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXPM_IO")
	}
	_addr_IMG_LoadXV_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadXV_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadXV_IO")
	}
	_addr_IMG_LoadWEBP_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadWEBP_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadWEBP_IO")
	}
	_addr_IMG_LoadSizedSVG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadSizedSVG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadSizedSVG_IO")
	}
	_addr_IMG_ReadXPMFromArray, err = puregogen.OpenSymbol(_hnd_img, "IMG_ReadXPMFromArray")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_ReadXPMFromArray")
	}
	_addr_IMG_ReadXPMFromArrayToRGB888, err = puregogen.OpenSymbol(_hnd_img, "IMG_ReadXPMFromArrayToRGB888")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_ReadXPMFromArrayToRGB888")
	}
	_addr_IMG_SaveAVIF, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveAVIF")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveAVIF")
	}
	_addr_IMG_SaveAVIF_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveAVIF_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveAVIF_IO")
	}
	_addr_IMG_SavePNG, err = puregogen.OpenSymbol(_hnd_img, "IMG_SavePNG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SavePNG")
	}
	_addr_IMG_SavePNG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SavePNG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SavePNG_IO")
	}
	_addr_IMG_SaveJPG, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveJPG")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveJPG")
	}
	_addr_IMG_SaveJPG_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_SaveJPG_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_SaveJPG_IO")
	}
	_addr_IMG_LoadAnimation, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimation")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimation")
	}
	_addr_IMG_LoadAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimation_IO")
	}
	_addr_IMG_LoadAnimationTyped_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadAnimationTyped_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadAnimationTyped_IO")
	}
	_addr_IMG_FreeAnimation, err = puregogen.OpenSymbol(_hnd_img, "IMG_FreeAnimation")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_FreeAnimation")
	}
	_addr_IMG_LoadGIFAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadGIFAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadGIFAnimation_IO")
	}
	_addr_IMG_LoadWEBPAnimation_IO, err = puregogen.OpenSymbol(_hnd_img, "IMG_LoadWEBPAnimation_IO")
	if err != nil {
		panic("cannot puregogen.OpenSymbol: IMG_LoadWEBPAnimation_IO")
	}

	iVersion = func() int32 {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Version)
		__r0 := int32(_r0)
		return __r0
	}
	iLoadTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTyped_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		runtime.KeepAlive(typ)
		return __r0
	}
	iLoad = func(file string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Load, uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(file)
		return __r0
	}
	iLoad_IO = func(src *sdl.IOStream, closeio bool) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_Load_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadTexture = func(renderer *sdl.Renderer, file string) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTexture, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(renderer)
		runtime.KeepAlive(file)
		return __r0
	}
	iLoadTexture_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTexture_IO, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(renderer)
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadTextureTyped_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) *sdl.Texture {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTextureTyped_IO, uintptr(unsafe.Pointer(renderer)), uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*sdl.Texture)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(renderer)
		runtime.KeepAlive(src)
		runtime.KeepAlive(typ)
		return __r0
	}
	iisAVIF = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isAVIF, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisICO = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isICO, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisCUR = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isCUR, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisBMP = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isBMP, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisGIF = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isGIF, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisJPG = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isJPG, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisJXL = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isJXL, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisLBM = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isLBM, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisPCX = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isPCX, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisPNG = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isPNG, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisPNM = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isPNM, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisSVG = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isSVG, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisQOI = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isQOI, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisTIF = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isTIF, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisXCF = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isXCF, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisXPM = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isXPM, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisXV = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isXV, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iisWEBP = func(src *sdl.IOStream) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_isWEBP, uintptr(unsafe.Pointer(src)))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadAVIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAVIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadICO_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadICO_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadCUR_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadCUR_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadBMP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadBMP_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadGIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadGIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadJPG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadJPG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadJXL_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadJXL_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadLBM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadLBM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadPCX_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPCX_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadPNG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPNG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadPNM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadPNM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadSVG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadSVG_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadQOI_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadQOI_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadTGA_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTGA_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadTIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadTIF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadXCF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXCF_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadXPM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXPM_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadXV_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadXV_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadWEBP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadWEBP_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadSizedSVG_IO = func(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadSizedSVG_IO, uintptr(unsafe.Pointer(src)), uintptr(width), uintptr(height))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iReadXPMFromArray = func(xpm *string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_ReadXPMFromArray, uintptr(unsafe.Pointer(xpm)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(xpm)
		return __r0
	}
	iReadXPMFromArrayToRGB888 = func(xpm *string) *sdl.Surface {
		_r0, _, _ := purego.SyscallN(_addr_IMG_ReadXPMFromArrayToRGB888, uintptr(unsafe.Pointer(xpm)))
		__r0 := (*sdl.Surface)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(xpm)
		return __r0
	}
	iSaveAVIF = func(surface *sdl.Surface, file string, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveAVIF, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))), uintptr(quality))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(file)
		return __r0
	}
	iSaveAVIF_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveAVIF_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio), uintptr(quality))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(dst)
		return __r0
	}
	iSavePNG = func(surface *sdl.Surface, file string) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SavePNG, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(file)
		return __r0
	}
	iSavePNG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SavePNG_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(dst)
		return __r0
	}
	iSaveJPG = func(surface *sdl.Surface, file string, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveJPG, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))), uintptr(quality))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(file)
		return __r0
	}
	iSaveJPG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_r0, _, _ := purego.SyscallN(_addr_IMG_SaveJPG_IO, uintptr(unsafe.Pointer(surface)), uintptr(unsafe.Pointer(dst)), puregogen.BoolToUintptr(closeio), uintptr(quality))
		__r0 := uint8(_r0) != 0
		runtime.KeepAlive(surface)
		runtime.KeepAlive(dst)
		return __r0
	}
	iLoadAnimation = func(file string) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimation, uintptr(unsafe.Pointer(puregogen.BytePtrFromString(file))))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(file)
		return __r0
	}
	iLoadAnimation_IO = func(src *sdl.IOStream, closeio bool) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimation_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadAnimationTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadAnimationTyped_IO, uintptr(unsafe.Pointer(src)), puregogen.BoolToUintptr(closeio), uintptr(unsafe.Pointer(puregogen.BytePtrFromString(typ))))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		runtime.KeepAlive(typ)
		return __r0
	}
	iFreeAnimation = func(anim *Animation) {
		purego.SyscallN(_addr_IMG_FreeAnimation, uintptr(unsafe.Pointer(anim)))
		runtime.KeepAlive(anim)
	}
	iLoadGIFAnimation_IO = func(src *sdl.IOStream) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadGIFAnimation_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
	iLoadWEBPAnimation_IO = func(src *sdl.IOStream) *Animation {
		_r0, _, _ := purego.SyscallN(_addr_IMG_LoadWEBPAnimation_IO, uintptr(unsafe.Pointer(src)))
		__r0 := (*Animation)(*(*unsafe.Pointer)(unsafe.Pointer(&_r0)))
		runtime.KeepAlive(src)
		return __r0
	}
}
