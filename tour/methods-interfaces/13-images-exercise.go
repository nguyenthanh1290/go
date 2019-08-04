package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w, h int
	c    uint8
}

func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img *Image) At(x, y int) color.Color {
	return color.RGBA{img.c + uint8(x^y), img.c + uint8(x^y), 255, 255}
	// return color.RGBA{uint8(x^y), 255, uint8(x^y), 255}
	// return color.RGBA{img.c + uint8(x^y), 255, img.c + uint8(x^y), 255}
	// return color.RGBA{uint8(x), 255, uint8(y), 255}
}

func main() {
	m := Image{255, 255, 255}
	pic.ShowImage(&m)
}
