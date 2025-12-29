package img

import (
	"unsafe"

	"github.com/Zyko0/go-sdl3/internal"
	"github.com/Zyko0/go-sdl3/sdl"
)

// IMG_Version - This function gets the version of the dynamically linked SDL_image library.
// (https://wiki.libsdl.org/SDL3_img/IMG_Version)
func GetVersion() sdl.Version {
	return sdl.Version(iVersion())
}

// IMG_LoadTyped_IO - Load an image from an SDL data source into a software surface.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTyped_IO)
func LoadTypedIO(src *sdl.IOStream, closeio bool, typ string) (*sdl.Surface, error) {
	surface := iLoadTyped_IO(src, closeio, typ)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_Load - Load an image from a filesystem path into a software surface.
// (https://wiki.libsdl.org/SDL3_img/IMG_Load)
func Load(file string) (*sdl.Surface, error) {
	surface := iLoad(file)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_Load_IO - Load an image from an SDL data source into a software surface.
// (https://wiki.libsdl.org/SDL3_img/IMG_Load_IO)
func LoadIO(src *sdl.IOStream, closeio bool) (*sdl.Surface, error) {
	surface := iLoad_IO(src, closeio)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadTexture - Load an image from a filesystem path into a GPU texture.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTexture)
func LoadTexture(renderer *sdl.Renderer, file string) (*sdl.Texture, error) {
	texture := iLoadTexture(renderer, file)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// IMG_LoadTexture_IO - Load an image from an SDL data source into a GPU texture.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTexture_IO)
func LoadTextureIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool) (*sdl.Texture, error) {
	texture := iLoadTexture_IO(renderer, src, closeio)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// IMG_LoadTextureTyped_IO - Load an image from an SDL data source into a GPU texture.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTextureTyped_IO)
func LoadTextureTypedIO(renderer *sdl.Renderer, src *sdl.IOStream, closeio bool, typ string) (*sdl.Texture, error) {
	texture := iLoadTextureTyped_IO(renderer, src, closeio, typ)
	if texture == nil {
		return nil, internal.LastErr()
	}

	return texture, nil
}

// IMG_isAVIF - Detect AVIF image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isAVIF)
func IsAVIF(src *sdl.IOStream) bool {
	return iisAVIF(src)
}

// IMG_isICO - Detect ICO image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isICO)
func IsICO(src *sdl.IOStream) bool {
	return iisICO(src)
}

// IMG_isCUR - Detect CUR image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isCUR)
func IsCUR(src *sdl.IOStream) bool {
	return iisCUR(src)
}

// IMG_isBMP - Detect BMP image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isBMP)
func IsBMP(src *sdl.IOStream) bool {
	return iisBMP(src)
}

// IMG_isGIF - Detect GIF image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isGIF)
func IsGIF(src *sdl.IOStream) bool {
	return iisGIF(src)
}

// IMG_isJPG - Detect JPG image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isJPG)
func IsJPG(src *sdl.IOStream) bool {
	return iisJPG(src)
}

// IMG_isJXL - Detect JXL image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isJXL)
func IsJXL(src *sdl.IOStream) bool {
	return iisJXL(src)
}

// IMG_isLBM - Detect LBM image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isLBM)
func IsLBM(src *sdl.IOStream) bool {
	return iisLBM(src)
}

// IMG_isPCX - Detect PCX image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isPCX)
func IsPCX(src *sdl.IOStream) bool {
	return iisPCX(src)
}

// IMG_isPNG - Detect PNG image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isPNG)
func IsPNG(src *sdl.IOStream) bool {
	return iisPNG(src)
}

// IMG_isPNM - Detect PNM image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isPNM)
func IsPNM(src *sdl.IOStream) bool {
	return iisPNM(src)
}

// IMG_isSVG - Detect SVG image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isSVG)
func IsSVG(src *sdl.IOStream) bool {
	return iisSVG(src)
}

// IMG_isQOI - Detect QOI image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isQOI)
func IsQOI(src *sdl.IOStream) bool {
	return iisQOI(src)
}

// IMG_isTIF - Detect TIFF image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isTIF)
func IsTIF(src *sdl.IOStream) bool {
	return iisTIF(src)
}

// IMG_isXCF - Detect XCF image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isXCF)
func IsXCF(src *sdl.IOStream) bool {
	return iisXCF(src)
}

// IMG_isXPM - Detect XPM image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isXPM)
func IsXPM(src *sdl.IOStream) bool {
	return iisXPM(src)
}

// IMG_isXV - Detect XV image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isXV)
func IsXV(src *sdl.IOStream) bool {
	return iisXV(src)
}

// IMG_isWEBP - Detect WEBP image data on a readable/seekable SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_image/IMG_isWEBP)
func IsWEBP(src *sdl.IOStream) bool {
	return iisWEBP(src)
}

// IMG_LoadAVIF_IO - Load a AVIF image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadAVIF_IO)
func LoadAVIF_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadAVIF_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadICO_IO - Load a ICO image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadICO_IO)
func LoadICO_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadICO_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadCUR_IO - Load a CUR image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadCUR_IO)
func LoadCUR_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadCUR_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadBMP_IO - Load a BMP image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadBMP_IO)
func LoadBMP_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadBMP_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadGIF_IO - Load a GIF image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadGIF_IO)
func LoadGIF_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadGIF_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadJPG_IO - Load a JPG image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadJPG_IO)
func LoadJPG_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadJPG_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadJXL_IO - Load a JXL image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadJXL_IO)
func LoadJXL_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadJXL_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadLBM_IO - Load a LBM image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadLBM_IO)
func LoadLBM_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadLBM_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadPCX_IO - Load a PCX image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadPCX_IO)
func LoadPCX_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadPCX_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadPNG_IO - Load a PNG image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadPNG_IO)
func LoadPNG_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadPNG_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadPNM_IO - Load a PNM image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadPNM_IO)
func LoadPNM_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadPNM_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadSVG_IO - Load a SVG image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadSVG_IO)
func LoadSVG_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadSVG_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadQOI_IO - Load a QOI image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadQOI_IO)
func LoadQOI_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadQOI_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadTGA_IO - Load a TGA image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTGA_IO)
func LoadTGA_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadTGA_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadTIF_IO - Load a TIFF image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadTIF_IO)
func LoadTIF_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadTIF_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadXCF_IO - Load a XCF image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadXCF_IO)
func LoadXCF_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadXCF_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadXPM_IO - Load a XPM image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadXPM_IO)
func LoadXPM_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadXPM_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadXV_IO - Load a XV image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadXV_IO)
func LoadXV_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadXV_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadWEBP_IO - Load a WEBP image directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadWEBP_IO)
func LoadWEBP_IO(src *sdl.IOStream) (*sdl.Surface, error) {
	surface := iLoadWEBP_IO(src)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_LoadSizedSVG_IO - Load an SVG image, scaled to a specific size.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadSizedSVG_IO)
func LoadSizedSVG_IO(src *sdl.IOStream, width, height int32) (*sdl.Surface, error) {
	surface := iLoadSizedSVG_IO(src, width, height)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_ReadXPMFromArray - Load an XPM image from a memory array.
// (https://wiki.libsdl.org/SDL3_img/IMG_ReadXPMFromArray)
func ReadXPMFFromArray(xpm []string) (*sdl.Surface, error) {
	xpm = append(xpm, "") // TODO: not sure that counts for a NULL-terminated array, see: https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArray
	surface := iReadXPMFromArray(unsafe.SliceData(xpm))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_ReadXPMFromArrayToRGB888 - Load an XPM image from a memory array.
// (https://wiki.libsdl.org/SDL3_img/IMG_ReadXPMFromArrayToRGB888)
func ReadXPMFFromArrayToRGB888(xpm []string) (*sdl.Surface, error) {
	xpm = append(xpm, "") // TODO: not sure that counts for a NULL-terminated array, see: https://wiki.libsdl.org/SDL3_image/IMG_ReadXPMFromArrayToRGB888
	surface := iReadXPMFromArrayToRGB888(unsafe.SliceData(xpm))
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// IMG_SaveAVIF - Save an SDL_Surface into a AVIF image file.
// (https://wiki.libsdl.org/SDL3_img/IMG_SaveAVIF)
func SaveAVIF(surface *sdl.Surface, file string, quality int32) error {
	if !iSaveAVIF(surface, file, quality) {
		return internal.LastErr()
	}

	return nil
}

// IMG_SaveAVIF_IO - Save an SDL_Surface into AVIF image data, via an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_img/IMG_SaveAVIF_IO)
func SaveAVIF_IO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) error {
	if !iSaveAVIF_IO(surface, dst, closeio, quality) {
		return internal.LastErr()
	}

	return nil
}

// IMG_SavePNG - Save an SDL_Surface into a PNG image file.
// (https://wiki.libsdl.org/SDL3_img/IMG_SavePNG)
func SavePNG(surface *sdl.Surface, file string) error {
	if !iSavePNG(surface, file) {
		return internal.LastErr()
	}

	return nil
}

// IMG_SavePNG_IO - Save an SDL_Surface into PNG image data, via an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_img/IMG_SavePNG_IO)
func SavePNG_IO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool) error {
	if !iSavePNG_IO(surface, dst, closeio) {
		return internal.LastErr()
	}

	return nil
}

// IMG_SaveJPG - Save an SDL_Surface into a JPEG image file.
// (https://wiki.libsdl.org/SDL3_img/IMG_SaveJPG)
func SaveJPG(surface *sdl.Surface, file string, quality int32) error {
	if !iSaveJPG(surface, file, quality) {
		return internal.LastErr()
	}

	return nil
}

// IMG_SaveJPG_IO - Save an SDL_Surface into JPEG image data, via an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_img/IMG_SaveJPG_IO)
func SaveJPG_IO(surface *sdl.Surface, dst *sdl.IOStream, closeio bool, quality int32) error {
	if !iSaveJPG_IO(surface, dst, closeio, quality) {
		return internal.LastErr()
	}

	return nil
}

// IMG_LoadAnimation - Load an animation from a file.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadAnimation)
func LoadAnimation(file string) (*Animation, error) {
	anim := iLoadAnimation(file)
	if anim == nil {
		return nil, internal.LastErr()
	}

	return anim, nil
}

// IMG_LoadAnimation_IO - Load an animation from an SDL_IOStream.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadAnimation_IO)
func LoadAnimation_IO(src *sdl.IOStream, closeio bool) (*Animation, error) {
	anim := iLoadAnimation_IO(src, closeio)
	if anim == nil {
		return nil, internal.LastErr()
	}

	return anim, nil
}

// IMG_LoadAnimationTyped_IO - Load an animation from an SDL datasource
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadAnimationTyped_IO)
func LoadAnimationTyped_IO(src *sdl.IOStream, closeio bool, typ string) (*Animation, error) {
	anim := iLoadAnimationTyped_IO(src, closeio, typ)
	if anim == nil {
		return nil, internal.LastErr()
	}

	return anim, nil
}

// IMG_LoadGIFAnimation_IO - Load a GIF animation directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadGIFAnimation_IO)
func LoadGIFAnimation_IO(src *sdl.IOStream) (*Animation, error) {
	anim := iLoadGIFAnimation_IO(src)
	if anim == nil {
		return nil, internal.LastErr()
	}

	return anim, nil
}

// IMG_LoadWEBPAnimation_IO - Load a WEBP animation directly.
// (https://wiki.libsdl.org/SDL3_img/IMG_LoadWEBPAnimation_IO)
func LoadWEBPAnimation_IO(src *sdl.IOStream) (*Animation, error) {
	anim := iLoadWEBPAnimation_IO(src)
	if anim == nil {
		return nil, internal.LastErr()
	}

	return anim, nil
}
