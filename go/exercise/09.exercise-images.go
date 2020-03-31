package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{R: uint8(x % 255), G: uint8(y % 255), B: 255, A: 255}
}

func main() {
	m := Image{
		w: 100,
		h: 100,
	}
	pic.ShowImage(m)
}
