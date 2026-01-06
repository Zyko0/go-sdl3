package common

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"reflect"
	"unsafe"

	"github.com/Zyko0/go-sdl3/examples/gpu/assets"
	"github.com/mdouchement/hdr"
	"github.com/mdouchement/hdr/codec/rgbe"
	"golang.org/x/image/bmp"
)

type HDRImage struct {
	Data []float32
	W    int
	H    int
}

func LoadHDR(filename string) (HDRImage, error) {
	data, err := assets.ReadFile("images/" + filename)
	if err != nil {
		return HDRImage{}, errors.New("failed to read data: " + err.Error())
	}

	img, err := rgbe.Decode(bytes.NewReader(data))
	if err != nil {
		return HDRImage{}, errors.New("failed to decode hdr image: " + err.Error())
	}

	hdrRGB, ok := img.(*hdr.RGB)
	if !ok {
		return HDRImage{}, fmt.Errorf("failed to cast: %s", reflect.TypeOf(img))
	}

	w := hdrRGB.Rect.Size().X
	h := hdrRGB.Rect.Size().Y

	rgba := make([]float32, w*h*4)

	for i := range w * h {
		rgba[i*4+0] = hdrRGB.Pix[i*3+0]
		rgba[i*4+1] = hdrRGB.Pix[i*3+1]
		rgba[i*4+2] = hdrRGB.Pix[i*3+2]
		rgba[i*4+3] = 1
	}

	return HDRImage{
		Data: rgba,
		W:    w,
		H:    h,
	}, nil
}

type Image struct {
	Data []byte
	W    int
	H    int
}

func LoadBMP(filename string) (Image, error) {
	imgBytes, err := assets.ReadFile("images/" + filename)
	if err != nil {
		return Image{}, errors.New("failed to read file: " + err.Error())
	}

	img, err := bmp.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return Image{}, errors.New("failed to decode bmp: " + err.Error())
	}

	imgNRGBA, ok := img.(*image.NRGBA)
	if ok {
		return Image{
			Data: imgNRGBA.Pix,
			W:    imgNRGBA.Rect.Size().X,
			H:    imgNRGBA.Rect.Size().Y,
		}, nil
	}

	imgRGBA, ok := img.(*image.RGBA)
	if ok {
		return Image{
			Data: imgRGBA.Pix,
			W:    imgRGBA.Rect.Size().X,
			H:    imgRGBA.Rect.Size().Y,
		}, nil
	}

	return Image{}, fmt.Errorf("unknown type: %s", reflect.TypeOf(img))
}

type ASTCHeader struct {
	magic  [4]uint8
	blockX uint8
	blockY uint8
	blockZ uint8
	dimX   [3]uint8
	dimY   [3]uint8
	dimZ   [3]uint8
}

func LoadASTC(filename string) (Image, error) {
	fileContents, err := assets.ReadFile("images/" + filename)
	if err != nil {
		return Image{}, errors.New("failed to read file: " + err.Error())
	}

	header := (*ASTCHeader)(unsafe.Pointer(&fileContents[0]))
	if header.magic[0] != 0x13 || header.magic[1] != 0xAB ||
		header.magic[2] != 0xA1 || header.magic[3] != 0x5C {
		return Image{}, errors.New("bad magic number")
	}

	width := uint32(header.dimX[0]) + (uint32(header.dimX[1]) << 8) + (uint32(header.dimX[2]) << 16)
	height := uint32(header.dimY[0]) + (uint32(header.dimY[1]) << 8) + (uint32(header.dimY[2]) << 16)

	blockCountX := (width + uint32(header.blockX) - 1) / uint32(header.blockX)
	blockCountY := (height + uint32(header.blockY) - 1) / uint32(header.blockY)
	imageDataLength := blockCountX * blockCountY * 16

	data := make([]byte, imageDataLength)
	copy(data, unsafe.Slice(
		(*byte)(unsafe.Pointer(&fileContents[unsafe.Sizeof(ASTCHeader{})])),
		imageDataLength,
	))

	return Image{
		Data: data,
		W:    int(width),
		H:    int(height),
	}, nil
}

type DDSPixelFormat struct {
	dwSize        int32
	dwFlags       int32
	dwFourCC      int32
	dwRGBBitCount int32
	dwRBitMask    int32
	dwGBitMask    int32
	dwBBitMask    int32
	dwABitMask    int32
}

type DDSHeader struct {
	dwMagic             int32
	dwSize              int32
	dwFlags             int32
	dwHeight            int32
	dwWidth             int32
	dwPitchOrLinearSize int32
	dwDepth             int32
	dwMipMapCount       int32
	dwReserved1         [11]int32
	ddspf               DDSPixelFormat
	dwCaps              int32
	dwCaps2             int32
	dwCaps3             int32
	dwCaps4             int32
	dwReserved2         int32
}

type DDSHeaderDXT10 struct {
	dxgiFormat        int32
	resourceDimension int32
	miscFlag          uint32
	arraySize         uint32
	miscFlags2        uint32
}

func LoadDDS(filename string) (Image, error) {
	fileContents, err := assets.ReadFile("images/" + filename)
	if err != nil {
		return Image{}, errors.New("failed to read file: " + err.Error())
	}

	header := (*DDSHeader)(unsafe.Pointer(&fileContents[0]))
	if header.dwMagic != 0x20534444 {
		return Image{}, errors.New("bad magic number")
	}

	hasDX10Header := header.ddspf.dwFlags == 0x4 && header.ddspf.dwFourCC == 0x30315844

	width := header.dwWidth
	height := header.dwHeight
	imageDataLength := header.dwPitchOrLinearSize

	data := make([]byte, imageDataLength)
	offset := unsafe.Sizeof(DDSHeader{})
	if hasDX10Header {
		offset += unsafe.Sizeof(DDSHeaderDXT10{})
	}
	copy(data, unsafe.Slice(
		(*byte)(unsafe.Pointer(&fileContents[offset])),
		imageDataLength,
	))

	return Image{
		Data: data,
		W:    int(width),
		H:    int(height),
	}, nil
}
