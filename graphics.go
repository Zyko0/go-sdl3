package sdl

import "github.com/Zyko0/go-sdl3/internal"

// Render

func GetNumRenderDrivers() int {
	return int(iGetNumRenderDrivers())
}

func GetRenderDriver(index int) string {
	return iGetRenderDriver(int32(index))
}

func CreateWindowAndRenderer(title string, width, height int, flags WindowFlags) (*Window, *Renderer, error) {
	var window *Window
	var renderer *Renderer

	if !iCreateWindowAndRenderer(title, int32(width), int32(height), flags, &window, &renderer) {
		return nil, nil, internal.LastErr()
	}

	return window, renderer, nil
}

func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	renderer := iCreateRendererWithProperties(props)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// Pixels

func GetPixelFormatForMasks(bpp int, rmask, gmask, bmask, amask uint32) PixelFormat {
	return iGetPixelFormatForMasks(int32(bpp), rmask, gmask, bmask, amask)
}

func CreatePalette(numColors int) (*Palette, error) {
	palette := iCreatePalette(int32(numColors))
	if palette == nil {
		return nil, internal.LastErr()
	}

	return palette, nil
}

func MapRGB(format *PixelFormatDetails, palette *Palette, r, g, b byte) uint32 {
	return iMapRGB(format, palette, r, g, b)
}

func MapRGBA(format *PixelFormatDetails, palette *Palette, r, g, b, a byte) uint32 {
	return iMapRGBA(format, palette, r, g, b, a)
}

func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b uint8) {
	iGetRGB(pixel, format, palette, &r, &g, &b)

	return r, g, b
}

func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b, a uint8) {
	iGetRGBA(pixel, format, palette, &r, &g, &b, &a)

	return r, g, b, a
}

// Surface

func CreateSurface(width, height int, format PixelFormat) (*Surface, error) {
	surface := iCreateSurface(int32(width), int32(height), format)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

/*
// TODO: idk about the void* pixels since the primitive type might depend on PixelFormat
func CreateSurfaceFrom(width, height int, format PixelFormat, pixels []uint32) (*Surface, error) {
	surface := iCreateSurface(width, height, format)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}*/

func LoadBMP_IO(src *IOStream, closeIO bool) (*Surface, error) {
	surface := iLoadBMP_IO(src, closeIO)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

func LoadBMP(file string) (*Surface, error) {
	surface := iLoadBMP(file)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// TODO: ConvertPixels => void* data
// TODO: ConvertPixelsAndColorspace => ^
// TODO: PremultiplyAlpha => ^

// Blend mode

func ComposeCustomBlendMode(srcFactor, dstFactor BlendFactor, colorOp BlendOperation, srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOp BlendOperation) BlendMode {
	return iComposeCustomBlendMode(srcFactor, dstFactor, colorOp, srcAlphaFactor, dstAlphaFactor, alphaOp)
}
