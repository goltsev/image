// Package resize implements algorithms for resizing images
// and contains some helpers for working with files and creating simple images.
package resize

import (
	"image"
	"math"

	"image/color"
	"image/draw"
)

// Naive changes image size to specified width and height.
// Function uses an algorithm similar to nearest-neighbour interpolation width similar results.
// Width and height have to be more than zero.
func Naive(src image.Image, width int, height int) image.Image {
	if src == nil {
		return nil
	}
	dst := createCanvas(src, width, height)
	naiveResize(src, dst, width, height)
	return dst
}

func naiveResize(src image.Image, dst draw.Image, width int, height int) {
	rect := src.Bounds()
	minx := rect.Min.X
	miny := rect.Min.Y
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
}

// Bilinear changes image size to specified width and height.
// Function uses a bilinear interpolation algorithm.
// Width and height have to be more than zero.
func Bilinear(src image.Image, width int, height int) image.Image {
	if src == nil {
		return nil
	}
	dst := createCanvas(src, width, height)
	bilinearResize(src, dst, width, height)
	return dst
}

func bilinearResize(src image.Image, dst draw.Image, width int, height int) {
	if src == nil || dst == nil {
		return
	}
	rect := src.Bounds()
	widthRatio := ratio(float64(rect.Max.X-1), float64(width-1))
	heightRatio := ratio(float64(rect.Max.Y-1), float64(height-1))
	for y := rect.Min.Y; y < height; y++ {
		for x := rect.Min.X; x < width; x++ {
			x0 := float64(x) * widthRatio
			y0 := float64(y) * heightRatio
			dst.Set(x, y, coordColor(src, x0, y0))
		}
	}
}

// coordColor returns color of a pixel at (x0, y0) coordinate
// using bilinear interpolation algorithm.
func coordColor(src image.Image, x0 float64, y0 float64) color.Color {
	if src == nil {
		return nil
	}
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
	r0, g0, b0, a0 := a.RGBA()
	r1, g1, b1, a1 := b.RGBA()
	c := color.RGBA{
		R: uint8((float64(r0)*(1-weight) + float64(r1)*weight) / 0x101),
		G: uint8((float64(g0)*(1-weight) + float64(g1)*weight) / 0x101),
		B: uint8((float64(b0)*(1-weight) + float64(b1)*weight) / 0x101),
		A: uint8((float64(a0)*(1-weight) + float64(a1)*weight) / 0x101)}
	return c
}

// createCanvas returns blank image of the same color type as specified image
// but width specified width and height.
func createCanvas(src image.Image, width int, height int) draw.Image {
	if src == nil {
		return nil
	}
	srcrect := src.Bounds()
	dstrect := image.Rect(srcrect.Min.X, srcrect.Min.Y, width, height)
	var dst draw.Image
	switch src.(type) {
	case (*image.Gray):
		dst = image.NewGray(dstrect)
	case (*image.Gray16):
		dst = image.NewGray16(dstrect)
	case (*image.CMYK):
		dst = image.NewCMYK(dstrect)
	case (*image.RGBA):
		dst = image.NewRGBA(dstrect)
	case (*image.RGBA64):
		dst = image.NewRGBA64(dstrect)
	case (*image.NRGBA):
		dst = image.NewNRGBA(dstrect)
	case (*image.NRGBA64):
		dst = image.NewNRGBA64(dstrect)
	default:
		dst = image.NewRGBA(dstrect)
	}
	return dst
}
