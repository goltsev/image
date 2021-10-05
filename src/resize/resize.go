package resize

import (
	"image"
	"math"

	"image/color"
)

func Naive(src image.Image, width int, height int) image.Image {
	if src == nil {
		return nil
	}
	rect := src.Bounds()
	minx := rect.Min.X
	miny := rect.Min.Y
	dst := image.NewRGBA(image.Rect(minx, miny, width, height))
	wratio := float64(rect.Max.X) / float64(width)
	hratio := float64(rect.Max.Y) / float64(height)
	for y := miny; y < height; y++ {
		for x := minx; x < width; x++ {
			dst.Set(x, y, src.At(
				int(float64(x)*wratio),
				int(float64(y)*hratio)),
			)
		}
	}
	return dst
}

func Bilinear(src image.Image, width int, height int) image.Image {
	if src == nil {
		return nil
	}
	rect := src.Bounds()
	dst := image.NewRGBA(image.Rect(rect.Min.X, rect.Min.Y, width, height))
	widthRatio := ratio(float64(rect.Max.X-1), float64(width-1))
	heightRatio := ratio(float64(rect.Max.Y-1), float64(height-1))
	for y := rect.Min.Y; y < height; y++ {
		for x := rect.Min.X; x < width; x++ {
			x0 := float64(x) * widthRatio
			y0 := float64(y) * heightRatio
			dst.Set(x, y, coordColor(src, x0, y0))
		}
	}
	return dst
}

func coordColor(src image.Image, x0 float64, y0 float64) color.RGBA {
	xleft, xright := edges(x0)
	ytop, ybot := edges(y0)

	color0 := src.At(int(xleft), int(ytop))
	color1 := src.At(int(xright), int(ytop))
	color2 := src.At(int(xleft), int(ybot))
	color3 := src.At(int(xright), int(ybot))

	weightx := x0 - xleft
	weighty := y0 - ytop
	wtop := weightedAverageColor(color0, color1, weightx)
	wbot := weightedAverageColor(color2, color3, weightx)
	return weightedAverageColor(wtop, wbot, weighty)
}

func ratio(x, y float64) float64 {
	if x == 0 {
		x = 1
	}
	if y == 0 {
		y = 1
	}
	return x / y
}

func edges(i float64) (left float64, right float64) {
	return math.Floor(i), math.Ceil(i)
}

// --a-----x----------b-->
//   ---w---
// weight is a difference between x and a coordinates
func weightedAverageColor(a color.Color, b color.Color, weight float64) color.RGBA {
	r0, g0, b0, _ := a.RGBA()
	r1, g1, b1, _ := b.RGBA()
	c := color.RGBA{
		R: uint8((float64(r0)*(1-weight) + float64(r1)*weight) / 0x101),
		G: uint8((float64(g0)*(1-weight) + float64(g1)*weight) / 0x101),
		B: uint8((float64(b0)*(1-weight) + float64(b1)*weight) / 0x101),
		A: 255}
	return c
}
