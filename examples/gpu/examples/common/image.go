package common

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"reflect"

	"github.com/Zyko0/go-sdl3/examples/gpu/content"
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
	data, err := content.ReadFile("images/" + filename)
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
	imgBytes, err := content.ReadFile("images/" + filename)
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
