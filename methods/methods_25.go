package main

// task
// http://127.0.0.1:3999/methods/25

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
	v uint8
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
	return color.RGBA{i.v, i.v, 255, 255}
}

func main() {
	m := Image{w: 100, h: 100, v: 0}
	pic.ShowImage(m)
}

