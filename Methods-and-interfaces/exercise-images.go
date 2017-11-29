package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct {
	v uint8
	w int
	h int
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
	m := Image{128, 200, 200}
	pic.ShowImage(m)
}
