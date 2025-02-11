package sdl

// Render

func GetNumRenderDrivers() int {
	return iGetNumRenderDrivers()
}

func GetRenderDriver(index int) string {
	return iGetRenderDriver(index)
}

func CreateWindowAndRenderer(title string, width, height int, flags WindowFlags) (*Window, *Renderer, error) {
	var window *Window
	var renderer *Renderer

	if !iCreateWindowAndRenderer(title, width, height, flags, &window, &renderer) {
		return nil, nil, lastError()
	}

	return window, renderer, nil
}

func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	renderer := iCreateRendererWithProperties(props)
	if renderer == nil {
		return nil, lastError()
	}

	return renderer, nil
}

// Pixels

func GetPixelFormatForMasks(bpp int, rmask, gmask, bmask, amask uint32) PixelFormat {
	return iGetPixelFormatForMasks(bpp, rmask, gmask, bmask, amask)
}

func CreatePalette(numColors int) (*Palette, error) {
	palette := iCreatePalette(numColors)
	if palette == nil {
		return nil, lastError()
	}

	return palette, nil
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
	surface := iCreateSurface(width, height, format)
	if surface == nil {
		return nil, lastError()
	}

	return surface, nil
}

/*
// TODO: idk about the void* pixels since the primitive type might depend on PixelFormat
func CreateSurfaceFrom(width, height int, format PixelFormat, pixels []uint32) (*Surface, error) {
	surface := iCreateSurface(width, height, format)
	if surface == nil {
		return nil, lastError()
	}

	return surface, nil
}*/

func LoadBMP_IO(src *IOStream, closeIO bool) (*Surface, error) {
	surface := iLoadBMP_IO(src, closeIO)
	if surface == nil {
		return nil, lastError()
	}

	return surface, nil
}

func LoadBMP(file string) (*Surface, error) {
	surface := iLoadBMP(file)
	if surface == nil {
		return nil, lastError()
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
