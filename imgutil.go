package main

import (
	"io"
	"image/png"
	"github.com/disintegration/imaging"
	"image/color"
	"image"
	"os"
	"golang.org/x/image/bmp"
)

func convertPngToBmp(reader io.Reader, fileName string, invertedImage bool) error {
	img, decodeErr := png.Decode(reader)
	if decodeErr != nil {
		return decodeErr
	}
	if invertedImage {
		img = imaging.Invert(img)
	}

	img = convertToGrayScale(img)

	f, createErr := os.Create(fileName)
	defer f.Close()
	if createErr != nil {
		return createErr
	}
	bmp.Encode(f, img)
	return nil
}

func convertToGrayScale(img image.Image) image.Image {
	b := img.Bounds()
	imgSet := image.NewGray(b)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			imgSet.Set(x, y, color.GrayModel.Convert(img.At(x, y)))
		}
	}
	return imgSet
}