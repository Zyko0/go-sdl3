//go:build js

package img

import (
	js "syscall/js"
	"unsafe"

	internal "github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

func initialize() {
	iVersion = func() int32 {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		ret := js.Global().Get("Module").Call(
			"_IMG_Version",
		)

		return int32(ret.Int())
	}

	iLoadTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTyped_IO",
			_src,
			_closeio,
			_typ,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoad = func(file string) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_Load",
			_file,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoad_IO = func(src *sdl.IOStream, closeio bool) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_Load_IO",
			_src,
			_closeio,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadTexture = func(renderer *sdl.Renderer, file string) *sdl.Texture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTexture",
			_renderer,
			_file,
		)

		_obj := internal.NewObject[sdl.Texture](ret)
		return _obj
	}

	iLoadTexture_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTexture_IO",
			_renderer,
			_src,
			_closeio,
		)

		_obj := internal.NewObject[sdl.Texture](ret)
		return _obj
	}

	iLoadTextureTyped_IO = func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) *sdl.Texture {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_renderer, ok := internal.GetJSPointer(renderer)
		if !ok {
			_renderer = internal.StackAlloc(int(unsafe.Sizeof(*renderer)))
		}
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTextureTyped_IO",
			_renderer,
			_src,
			_closeio,
			_typ,
		)

		_obj := internal.NewObject[sdl.Texture](ret)
		return _obj
	}

	iLoadAVIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAVIF_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadICO_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadICO_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadCUR_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadCUR_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadBMP_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadBMP_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadGIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadGIF_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadJPG_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadJPG_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadJXL_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadJXL_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadLBM_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadLBM_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadPCX_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPCX_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadPNG_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPNG_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadPNM_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadPNM_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadSVG_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadSVG_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadQOI_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadQOI_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadTGA_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTGA_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadTIF_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadTIF_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadXCF_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXCF_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadXPM_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXPM_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadXV_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadXV_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadWEBP_IO = func(src *sdl.IOStream) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadWEBP_IO",
			_src,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	iLoadSizedSVG_IO = func(src *sdl.IOStream, width int32, height int32) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_width := int32(width)
		_height := int32(height)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadSizedSVG_IO",
			_src,
			_width,
			_height,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}

	/*iReadXPMFromArray = func(xpm *string) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_xpm, ok := internal.GetJSPointer(xpm)
		if !ok {
			_xpm = internal.StackAlloc()
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_ReadXPMFromArray",
			_xpm,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	/*iReadXPMFromArrayToRGB888 = func(xpm *string) *sdl.Surface {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_xpm, ok := internal.GetJSPointer(xpm)
		if !ok {
			_xpm = internal.StackAlloc()
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_ReadXPMFromArrayToRGB888",
			_xpm,
		)

		_obj := internal.NewObject[sdl.Surface](ret)
		return _obj
	}*/

	iSaveAVIF = func(surface *sdl.Surface, file string, quality int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_file := internal.StringOnJSStack(file)
		_quality := int32(quality)
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveAVIF",
			_surface,
			_file,
			_quality,
		)

		return internal.GetBool(ret)
	}

	iSaveAVIF_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_closeio := internal.NewBoolean(closeio)
		_quality := int32(quality)
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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_SavePNG",
			_surface,
			_file,
		)

		return internal.GetBool(ret)
	}

	iSavePNG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_file := internal.StringOnJSStack(file)
		_quality := int32(quality)
		ret := js.Global().Get("Module").Call(
			"_IMG_SaveJPG",
			_surface,
			_file,
			_quality,
		)

		return internal.GetBool(ret)
	}

	iSaveJPG_IO = func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_surface, ok := internal.GetJSPointer(surface)
		if !ok {
			_surface = internal.StackAlloc(int(unsafe.Sizeof(*surface)))
		}
		_dst, ok := internal.GetJSPointer(dst)
		if !ok {
			_dst = internal.StackAlloc(int(unsafe.Sizeof(*dst)))
		}
		_closeio := internal.NewBoolean(closeio)
		_quality := int32(quality)
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
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_file := internal.StringOnJSStack(file)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimation",
			_file,
		)

		_obj := internal.NewObject[Animation](ret)
		return _obj
	}

	iLoadAnimation_IO = func(src *sdl.IOStream, closeio bool) *Animation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimation_IO",
			_src,
			_closeio,
		)

		_obj := internal.NewObject[Animation](ret)
		return _obj
	}

	iLoadAnimationTyped_IO = func(src *sdl.IOStream, closeio bool, typ string) *Animation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		_closeio := internal.NewBoolean(closeio)
		_typ := internal.StringOnJSStack(typ)
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadAnimationTyped_IO",
			_src,
			_closeio,
			_typ,
		)

		_obj := internal.NewObject[Animation](ret)
		return _obj
	}

	iFreeAnimation = func(anim *Animation) {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_anim, ok := internal.GetJSPointer(anim)
		if !ok {
			_anim = internal.StackAlloc(int(unsafe.Sizeof(*anim)))
		}
		js.Global().Get("Module").Call(
			"_IMG_FreeAnimation",
			_anim,
		)
	}

	iLoadGIFAnimation_IO = func(src *sdl.IOStream) *Animation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadGIFAnimation_IO",
			_src,
		)

		_obj := internal.NewObject[Animation](ret)
		return _obj
	}

	iLoadWEBPAnimation_IO = func(src *sdl.IOStream) *Animation {
		panic("not implemented on js")
		internal.StackSave()
		defer internal.StackRestore()
		_src, ok := internal.GetJSPointer(src)
		if !ok {
			_src = internal.StackAlloc(int(unsafe.Sizeof(*src)))
		}
		ret := js.Global().Get("Module").Call(
			"_IMG_LoadWEBPAnimation_IO",
			_src,
		)

		_obj := internal.NewObject[Animation](ret)
		return _obj
	}

}
