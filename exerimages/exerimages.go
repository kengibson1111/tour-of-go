package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type MyImage struct{}

func (im MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (im MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 300, 300)
}

func (im MyImage) At(x, y int) color.Color {
	v := uint8(((x*y)/2) % (1<<8-1))
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := MyImage{}
	pic.ShowImage(m)
}

