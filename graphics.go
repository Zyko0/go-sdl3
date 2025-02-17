package sdl

import "github.com/Zyko0/go-sdl3/internal"

// Render

// SDL_GetNumRenderDrivers - Get the number of 2D rendering drivers available for the current display.
// (https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers)
func GetNumRenderDrivers() int {
	return int(iGetNumRenderDrivers())
}

// SDL_GetRenderDriver - Use this function to get the name of a built in 2D rendering driver.
// (https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver)
func GetRenderDriver(index int) string {
	return iGetRenderDriver(int32(index))
}

// SDL_CreateWindowAndRenderer - Create a window and default renderer.
// (https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer)
func CreateWindowAndRenderer(title string, width, height int, flags WindowFlags) (*Window, *Renderer, error) {
	var window *Window
	var renderer *Renderer

	if !iCreateWindowAndRenderer(title, int32(width), int32(height), flags, &window, &renderer) {
		return nil, nil, internal.LastErr()
	}

	return window, renderer, nil
}

// SDL_CreateRendererWithProperties - Create a 2D rendering context for a window, with the specified properties.
// (https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties)
func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	renderer := iCreateRendererWithProperties(props)
	if renderer == nil {
		return nil, internal.LastErr()
	}

	return renderer, nil
}

// Pixels

// SDL_GetPixelFormatForMasks - Convert a bpp value and RGBA masks to an enumerated pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks)
func GetPixelFormatForMasks(bpp int, rmask, gmask, bmask, amask uint32) PixelFormat {
	return iGetPixelFormatForMasks(int32(bpp), rmask, gmask, bmask, amask)
}

// SDL_CreatePalette - Create a palette structure with the specified number of color entries.
// (https://wiki.libsdl.org/SDL3/SDL_CreatePalette)
func CreatePalette(numColors int) (*Palette, error) {
	palette := iCreatePalette(int32(numColors))
	if palette == nil {
		return nil, internal.LastErr()
	}

	return palette, nil
}

// SDL_MapRGB - Map an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGB)
func MapRGB(format *PixelFormatDetails, palette *Palette, r, g, b byte) uint32 {
	return iMapRGB(format, palette, r, g, b)
}

// SDL_MapRGBA - Map an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_MapRGBA)
func MapRGBA(format *PixelFormatDetails, palette *Palette, r, g, b, a byte) uint32 {
	return iMapRGBA(format, palette, r, g, b, a)
}

// SDL_GetRGB - Get RGB values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGB)
func GetRGB(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b uint8) {
	iGetRGB(pixel, format, palette, &r, &g, &b)

	return r, g, b
}

// SDL_GetRGBA - Get RGBA values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL3/SDL_GetRGBA)
func GetRGBA(pixel uint32, format *PixelFormatDetails, palette *Palette) (r, g, b, a uint8) {
	iGetRGBA(pixel, format, palette, &r, &g, &b, &a)

	return r, g, b, a
}

// Surface

// SDL_CreateSurface - Allocate a new surface with a specific pixel format.
// (https://wiki.libsdl.org/SDL3/SDL_CreateSurface)
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

// SDL_LoadBMP_IO - Load a BMP image from a seekable SDL data stream.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO)
func LoadBMP_IO(src *IOStream, closeIO bool) (*Surface, error) {
	surface := iLoadBMP_IO(src, closeIO)
	if surface == nil {
		return nil, internal.LastErr()
	}

	return surface, nil
}

// SDL_LoadBMP - Load a BMP image from a file.
// (https://wiki.libsdl.org/SDL3/SDL_LoadBMP)
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

// SDL_ComposeCustomBlendMode - Compose a custom blend mode for renderers.
// (https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode)
func ComposeCustomBlendMode(srcFactor, dstFactor BlendFactor, colorOp BlendOperation, srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOp BlendOperation) BlendMode {
	return iComposeCustomBlendMode(srcFactor, dstFactor, colorOp, srcAlphaFactor, dstAlphaFactor, alphaOp)
}

// GPU

// SDL_GPUSupportsShaderFormats - Checks for GPU runtime support.
// (https://wiki.libsdl.org/SDL3/SDL_GPUSupportsShaderFormats)
func GPUSupportShaderFormats(formatFlags GPUShaderFormat, name string) bool {
	return iGPUSupportsShaderFormats(formatFlags, name)
}

// SDL_CreateGPUDevice - Creates a GPU context.
// (https://wiki.libsdl.org/SDL3/SDL_CreateGPUDevice)
func CreateGPUDevice(formatFlags GPUShaderFormat, debugMode bool, name string) (*GPUDevice, error) {
	device := iCreateGPUDevice(formatFlags, debugMode, name)
	if device == nil {
		return nil, internal.LastErr()
	}

	return device, nil
}
