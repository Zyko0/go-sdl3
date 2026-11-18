//go:build js

package img

import (
	js "syscall/js"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

func initialize() {
	iVersion = func() int32 {
		ret := js.Global().Get("Module").Call(
			"_IMG_Version",
		)
		return int32(ret.Int())
	}

	iLoadTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *sdl.Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTyped_IO",
			_src,
			_closeio,
			_typ,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoad = func(file string) *sdl.Surface {
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_Load",
			_file,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoad_IO = func(src *sdl.IOStream, closeio bool) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_Load_IO",
			_src,
			_closeio,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadTexture = func(renderer *sdl.Renderer, file string) *sdl.Texture {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, _ := internal.GetJSPointer(renderer)
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTexture",
			_renderer,
			_file,
		)
		return internal.NewObject[sdl.Texture](ret)
	}

	iLoadTexture_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
		_renderer, _ := internal.GetJSPointer(renderer)
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTexture_IO",
			_renderer,
			_src,
			_closeio,
		)
		return internal.NewObject[sdl.Texture](ret)
	}

	iLoadTextureTyped_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) *sdl.Texture {
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, _ := internal.GetJSPointer(renderer)
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTextureTyped_IO",
			_renderer,
			_src,
			_closeio,
			_typ,
		)
		return internal.NewObject[sdl.Texture](ret)
	}

	iLoadAVIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAVIF_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadICO_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadICO_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadCUR_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadCUR_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadBMP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadBMP_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadGIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadGIF_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadJPG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadJPG_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadJXL_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadJXL_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadLBM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadLBM_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadPCX_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPCX_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadPNG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPNG_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadPNM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPNM_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadSVG_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadSVG_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadQOI_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadQOI_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadTGA_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTGA_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadTIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTIF_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadXCF_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXCF_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadXPM_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXPM_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadXV_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXV_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadWEBP_IO = func(src *sdl.IOStream) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadWEBP_IO",
			_src,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iLoadSizedSVG_IO = func(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
		_src, _ := internal.GetJSPointer(src)
		_width := width
		_height := height
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadSizedSVG_IO",
			_src,
			_width,
			_height,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iSaveAVIF = func(surface *sdl.Surface, file string, quality int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, _ := internal.GetJSPointer(surface)
		_file := internal.StringOnJSStack(file)
		_quality := quality
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveAVIF",
			_surface,
			_file,
			_quality,
		)
		return internal.GetBool(ret)
	}

	iSaveAVIF_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_surface, _ := internal.GetJSPointer(surface)
		_dst, _ := internal.GetJSPointer(dst)
		_closeio := internal.NewBoolean(closeio)
		_quality := quality
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveAVIF_IO",
			_surface,
			_dst,
			_closeio,
			_quality,
		)
		return internal.GetBool(ret)
	}

	iSavePNG = func(surface *sdl.Surface, file string) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, _ := internal.GetJSPointer(surface)
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_SavePNG",
			_surface,
			_file,
		)
		return internal.GetBool(ret)
	}

	iSavePNG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
		_surface, _ := internal.GetJSPointer(surface)
		_dst, _ := internal.GetJSPointer(dst)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_SavePNG_IO",
			_surface,
			_dst,
			_closeio,
		)
		return internal.GetBool(ret)
	}

	iSaveJPG = func(surface *sdl.Surface, file string, quality int32) bool {
		internal.StackSave()
		defer internal.StackRestore()
		_surface, _ := internal.GetJSPointer(surface)
		_file := internal.StringOnJSStack(file)
		_quality := quality
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveJPG",
			_surface,
			_file,
			_quality,
		)
		return internal.GetBool(ret)
	}

	iSaveJPG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		_surface, _ := internal.GetJSPointer(surface)
		_dst, _ := internal.GetJSPointer(dst)
		_closeio := internal.NewBoolean(closeio)
		_quality := quality
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveJPG_IO",
			_surface,
			_dst,
			_closeio,
			_quality,
		)
		return internal.GetBool(ret)
	}

	iLoadAnimation = func(file string) *Animation {
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimation",
			_file,
		)
		return internal.NewObject[Animation](ret)
	}

	iLoadAnimation_IO = func(src *sdl.IOStream, closeio bool) *Animation {
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimation_IO",
			_src,
			_closeio,
		)
		return internal.NewObject[Animation](ret)
	}

	iLoadAnimationTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *Animation {
		internal.StackSave()
		defer internal.StackRestore()
		_src, _ := internal.GetJSPointer(src)
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimationTyped_IO",
			_src,
			_closeio,
			_typ,
		)
		return internal.NewObject[Animation](ret)
	}

	iFreeAnimation = func(anim *Animation) {
		_anim, _ := internal.GetJSPointer(anim)
		js.Global().Get("Module").Call(
			"_IMG_FreeAnimation",
			_anim,
		)
	}

	iLoadGIFAnimation_IO = func(src *sdl.IOStream) *Animation {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadGIFAnimation_IO",
			_src,
		)
		return internal.NewObject[Animation](ret)
	}

	iLoadWEBPAnimation_IO = func(src *sdl.IOStream) *Animation {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadWEBPAnimation_IO",
			_src,
		)
		return internal.NewObject[Animation](ret)
	}

	iisAVIF = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isAVIF",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisICO = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isICO",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisCUR = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isCUR",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisBMP = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isBMP",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisGIF = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isGIF",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisJPG = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isJPG",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisJXL = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isJXL",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisLBM = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isLBM",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisPCX = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isPCX",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisPNG = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isPNG",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisPNM = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isPNM",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisSVG = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isSVG",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisQOI = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isQOI",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisTIF = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isTIF",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisXCF = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isXCF",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisXPM = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isXPM",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisXV = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isXV",
			_src,
		)
		return internal.GetBool(ret)
	}

	iisWEBP = func(src *sdl.IOStream) bool {
		_src, _ := internal.GetJSPointer(src)
		ret := js.Global().Get("Module").Call(
			"_IMG_isWEBP",
			_src,
		)
		return internal.GetBool(ret)
	}

	iReadXPMFromArray = func(xpm *string) *sdl.Surface {
		_xpm, free := cStringArrayToJSHeap(xpm)
		defer free()
		ret := js.Global().Get("Module").Call(
			"_IMG_ReadXPMFromArray",
			_xpm,
		)
		return internal.NewObject[sdl.Surface](ret)
	}

	iReadXPMFromArrayToRGB888 = func(xpm *string) *sdl.Surface {
		_xpm, free := cStringArrayToJSHeap(xpm)
		defer free()
		ret := js.Global().Get("Module").Call(
			"_IMG_ReadXPMFromArrayToRGB888",
			_xpm,
		)
		return internal.NewObject[sdl.Surface](ret)
	}
}
