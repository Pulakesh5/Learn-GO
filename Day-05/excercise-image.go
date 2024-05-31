package main

import (
	"image"
	"image/color"
)

// import "golang.org/x/tour/pic"

type Image struct {
	w, h  int
	Image *image.RGBA
}

func (i *Image) At(x, y int) color.Color {
	return i.Image.At(x, y)
}
func (i *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}
func (i *Image) colorModel() color.Model {
	return i.Image.ColorModel()
}
func main() {
	m := Image{w: 100, h: 100}
	m.Image = image.NewRGBA(image.Rect(0, 0, m.w, m.h))
	for x := 0; x < m.w; x++ {
		for y := 0; y < m.h; y++ {
			m.Image.Set(x, y, color.RGBA{uint8(x), uint8(y), uint8(x * y), 255})
		}
	}
	pic.ShowImage(m)
}
