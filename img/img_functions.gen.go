package img

import sdl "github.com/Zyko0/go-sdl3/sdl"

// Code generated by cmd/ffi2go. DO NOT EDIT.

var (
	// IMG_Version => This function gets the version of the dynamically linked SDL_image library.
	//
	//puregogen:library path:windows=img.dll path:unix=img.so alias=img
	//puregogen:function symbol=IMG_Version
	iVersion func() int32

	// IMG_LoadTyped_IO => Load an image from an SDL data source into a software surface.
	//
	//puregogen:function symbol=IMG_LoadTyped_IO
	iLoadTyped_IO func(src *sdl.IOStream, closeio bool, typ string) *sdl.Surface

	// IMG_Load => Load an image from a filesystem path into a software surface.
	//
	//puregogen:function symbol=IMG_Load
	iLoad func(file string) *sdl.Surface

	// IMG_Load_IO => Load an image from an SDL data source into a software surface.
	//
	//puregogen:function symbol=IMG_Load_IO
	iLoad_IO func(src *sdl.IOStream, closeio bool) *sdl.Surface

	// IMG_LoadTexture => Load an image from a filesystem path into a GPU texture.
	//
	//puregogen:function symbol=IMG_LoadTexture
	iLoadTexture func(renderer *sdl.Renderer, file string) *sdl.Texture

	// IMG_LoadTexture_IO => Load an image from an SDL data source into a GPU texture.
	//
	//puregogen:function symbol=IMG_LoadTexture_IO
	iLoadTexture_IO func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) *sdl.Texture

	// IMG_LoadTextureTyped_IO => Load an image from an SDL data source into a GPU texture.
	//
	//puregogen:function symbol=IMG_LoadTextureTyped_IO
	iLoadTextureTyped_IO func(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) *sdl.Texture

	// IMG_isAVIF => Detect AVIF image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isAVIF
	iisAVIF func(src *sdl.IOStream) bool

	// IMG_isICO => Detect ICO image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isICO
	iisICO func(src *sdl.IOStream) bool

	// IMG_isCUR => Detect CUR image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isCUR
	iisCUR func(src *sdl.IOStream) bool

	// IMG_isBMP => Detect BMP image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isBMP
	iisBMP func(src *sdl.IOStream) bool

	// IMG_isGIF => Detect GIF image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isGIF
	iisGIF func(src *sdl.IOStream) bool

	// IMG_isJPG => Detect JPG image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isJPG
	iisJPG func(src *sdl.IOStream) bool

	// IMG_isJXL => Detect JXL image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isJXL
	iisJXL func(src *sdl.IOStream) bool

	// IMG_isLBM => Detect LBM image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isLBM
	iisLBM func(src *sdl.IOStream) bool

	// IMG_isPCX => Detect PCX image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isPCX
	iisPCX func(src *sdl.IOStream) bool

	// IMG_isPNG => Detect PNG image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isPNG
	iisPNG func(src *sdl.IOStream) bool

	// IMG_isPNM => Detect PNM image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isPNM
	iisPNM func(src *sdl.IOStream) bool

	// IMG_isSVG => Detect SVG image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isSVG
	iisSVG func(src *sdl.IOStream) bool

	// IMG_isQOI => Detect QOI image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isQOI
	iisQOI func(src *sdl.IOStream) bool

	// IMG_isTIF => Detect TIFF image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isTIF
	iisTIF func(src *sdl.IOStream) bool

	// IMG_isXCF => Detect XCF image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isXCF
	iisXCF func(src *sdl.IOStream) bool

	// IMG_isXPM => Detect XPM image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isXPM
	iisXPM func(src *sdl.IOStream) bool

	// IMG_isXV => Detect XV image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isXV
	iisXV func(src *sdl.IOStream) bool

	// IMG_isWEBP => Detect WEBP image data on a readable/seekable SDL_IOStream.
	//
	//puregogen:function symbol=IMG_isWEBP
	iisWEBP func(src *sdl.IOStream) bool

	// IMG_LoadAVIF_IO => Load a AVIF image directly.
	//
	//puregogen:function symbol=IMG_LoadAVIF_IO
	iLoadAVIF_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadICO_IO => Load a ICO image directly.
	//
	//puregogen:function symbol=IMG_LoadICO_IO
	iLoadICO_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadCUR_IO => Load a CUR image directly.
	//
	//puregogen:function symbol=IMG_LoadCUR_IO
	iLoadCUR_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadBMP_IO => Load a BMP image directly.
	//
	//puregogen:function symbol=IMG_LoadBMP_IO
	iLoadBMP_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadGIF_IO => Load a GIF image directly.
	//
	//puregogen:function symbol=IMG_LoadGIF_IO
	iLoadGIF_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadJPG_IO => Load a JPG image directly.
	//
	//puregogen:function symbol=IMG_LoadJPG_IO
	iLoadJPG_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadJXL_IO => Load a JXL image directly.
	//
	//puregogen:function symbol=IMG_LoadJXL_IO
	iLoadJXL_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadLBM_IO => Load a LBM image directly.
	//
	//puregogen:function symbol=IMG_LoadLBM_IO
	iLoadLBM_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadPCX_IO => Load a PCX image directly.
	//
	//puregogen:function symbol=IMG_LoadPCX_IO
	iLoadPCX_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadPNG_IO => Load a PNG image directly.
	//
	//puregogen:function symbol=IMG_LoadPNG_IO
	iLoadPNG_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadPNM_IO => Load a PNM image directly.
	//
	//puregogen:function symbol=IMG_LoadPNM_IO
	iLoadPNM_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadSVG_IO => Load a SVG image directly.
	//
	//puregogen:function symbol=IMG_LoadSVG_IO
	iLoadSVG_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadQOI_IO => Load a QOI image directly.
	//
	//puregogen:function symbol=IMG_LoadQOI_IO
	iLoadQOI_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadTGA_IO => Load a TGA image directly.
	//
	//puregogen:function symbol=IMG_LoadTGA_IO
	iLoadTGA_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadTIF_IO => Load a TIFF image directly.
	//
	//puregogen:function symbol=IMG_LoadTIF_IO
	iLoadTIF_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadXCF_IO => Load a XCF image directly.
	//
	//puregogen:function symbol=IMG_LoadXCF_IO
	iLoadXCF_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadXPM_IO => Load a XPM image directly.
	//
	//puregogen:function symbol=IMG_LoadXPM_IO
	iLoadXPM_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadXV_IO => Load a XV image directly.
	//
	//puregogen:function symbol=IMG_LoadXV_IO
	iLoadXV_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadWEBP_IO => Load a WEBP image directly.
	//
	//puregogen:function symbol=IMG_LoadWEBP_IO
	iLoadWEBP_IO func(src *sdl.IOStream) *sdl.Surface

	// IMG_LoadSizedSVG_IO => Load an SVG image, scaled to a specific size.
	//
	//puregogen:function symbol=IMG_LoadSizedSVG_IO
	iLoadSizedSVG_IO func(src *sdl.IOStream, width int32, height int32) *sdl.Surface

	// IMG_ReadXPMFromArray => Load an XPM image from a memory array.
	//
	//puregogen:function symbol=IMG_ReadXPMFromArray
	iReadXPMFromArray func(xpm *string) *sdl.Surface

	// IMG_ReadXPMFromArrayToRGB888 => Load an XPM image from a memory array.
	//
	//puregogen:function symbol=IMG_ReadXPMFromArrayToRGB888
	iReadXPMFromArrayToRGB888 func(xpm *string) *sdl.Surface

	// IMG_SaveAVIF => Save an SDL_Surface into a AVIF image file.
	//
	//puregogen:function symbol=IMG_SaveAVIF
	iSaveAVIF func(surface *sdl.Surface, file string, quality int32) bool

	// IMG_SaveAVIF_IO => Save an SDL_Surface into AVIF image data, via an SDL_IOStream.
	//
	//puregogen:function symbol=IMG_SaveAVIF_IO
	iSaveAVIF_IO func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool

	// IMG_SavePNG => Save an SDL_Surface into a PNG image file.
	//
	//puregogen:function symbol=IMG_SavePNG
	iSavePNG func(surface *sdl.Surface, file string) bool

	// IMG_SavePNG_IO => Save an SDL_Surface into PNG image data, via an SDL_IOStream.
	//
	//puregogen:function symbol=IMG_SavePNG_IO
	iSavePNG_IO func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) bool

	// IMG_SaveJPG => Save an SDL_Surface into a JPEG image file.
	//
	//puregogen:function symbol=IMG_SaveJPG
	iSaveJPG func(surface *sdl.Surface, file string, quality int32) bool

	// IMG_SaveJPG_IO => Save an SDL_Surface into JPEG image data, via an SDL_IOStream.
	//
	//puregogen:function symbol=IMG_SaveJPG_IO
	iSaveJPG_IO func(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) bool

	// IMG_LoadAnimation => Load an animation from a file.
	//
	//puregogen:function symbol=IMG_LoadAnimation
	iLoadAnimation func(file string) *Animation

	// IMG_LoadAnimation_IO => Load an animation from an SDL_IOStream.
	//
	//puregogen:function symbol=IMG_LoadAnimation_IO
	iLoadAnimation_IO func(src *sdl.IOStream, closeio bool) *Animation

	// IMG_LoadAnimationTyped_IO => Load an animation from an SDL datasource
	//
	//puregogen:function symbol=IMG_LoadAnimationTyped_IO
	iLoadAnimationTyped_IO func(src *sdl.IOStream, closeio bool, typ string) *Animation

	// IMG_FreeAnimation => Dispose of an IMG_Animation and free its resources.
	//
	//puregogen:function symbol=IMG_FreeAnimation
	iFreeAnimation func(anim *Animation)

	// IMG_LoadGIFAnimation_IO => Load a GIF animation directly.
	//
	//puregogen:function symbol=IMG_LoadGIFAnimation_IO
	iLoadGIFAnimation_IO func(src *sdl.IOStream) *Animation

	// IMG_LoadWEBPAnimation_IO => Load a WEBP animation directly.
	//
	//puregogen:function symbol=IMG_LoadWEBPAnimation_IO
	iLoadWEBPAnimation_IO func(src *sdl.IOStream) *Animation
)
