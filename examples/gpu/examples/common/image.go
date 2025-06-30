package common

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"

	"github.com/Zyko0/go-sdl3/examples/gpu/content"
	"github.com/mdouchement/hdr"
	"github.com/mdouchement/hdr/codec/rgbe"
)

func LoadHDRImage(imageFilename string) ([]float32, int, int, error) {
	data, err := content.ReadFile("images/" + imageFilename)
	if err != nil {
		return nil, 0, 0, errors.New("failed to read data: " + err.Error())
	}

	img, err := rgbe.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, 0, 0, errors.New("failed to decode hdr image: " + err.Error())
	}

	hdrRGB, ok := img.(*hdr.RGB)
	if !ok {
		return nil, 0, 0, fmt.Errorf("failed to cast: %s", reflect.TypeOf(img))
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

	return rgba, w, h, nil
}
