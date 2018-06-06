package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (Image) At(x, y int) color.Color {

	return color.RGBA{uint8(x ^ y), uint8(x ^ y), 200, 100}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
