package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

func rand(x, y int) uint8 {
	return uint8((x+y)/2 | x*y | x ^ y)
}

func Pic(dx, dy int) [][]uint8 {

	pic := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		row := make([]uint8, dy)
		for y := 0; y < dy; y++ {
			row[y] = rand(x, y)
		}
		pic[x] = row
	}

	return pic
}

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
	v := rand(x, y)
	return color.RGBA{
		v, v, 255, 255,
	}
}

func NewImage(w, h int) *Image {
	return &Image{
		rec:       image.Rect(0, 0, w, h),
		colorMode: color.RGBAModel,
	}
}

func main() {
	m := NewImage(10, 10)
	pic.ShowImage(m)
}
