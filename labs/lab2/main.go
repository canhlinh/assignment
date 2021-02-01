package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	rec       image.Rectangle
	colorMode color.Model
}

func (i Image) Bounds() image.Rectangle {
	return i.rec
}

func (i Image) ColorModel() color.Model {
	return i.colorMode
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{
		0, 0, 255, 255,
	}
}

func NewImage() *Image {
	return &Image{
		rec:       image.Rect(0, 0, 10, 10),
		colorMode: color.RGBAModel,
	}
}

func main() {
	m := NewImage()
	pic.ShowImage(m)
}
