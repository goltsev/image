package resize

import (
	"image"
	"image/color"
	"image/draw"
)

// CreateSquare returns a square with a diagonal line of pixels
// from top-left to bottom-right, where image width = size and image height = size.
// Background is blue and the line is red.
func CreateSquare(size int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	draw.Draw(img, img.Bounds(), image.NewUniform(blue), image.Point{}, draw.Src)
	for i := 0; i < size; i++ {
		img.Set(i, i, red)
	}
	return img
}

// CreateLine creates a dotted line with specified width and height.
// Background is blue and dots is red.
func CreateLine(width int, height int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	draw.Draw(img, img.Bounds(), image.NewUniform(blue), image.Point{}, draw.Src)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j += 2 {
			img.Set(j, i, red)
		}
	}
	return img
}

// CreateCheckers creates a checker board,
// where image width = size and image height = size.
// Step is 1 pixel. Colors are black and white.
func CreateCheckers(size int) image.Image {
	if size < 1 {
		return nil
	}
	pix := make([]uint8, size*size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pix[i*size+j] = uint8(((i + j + 1) % 2) * 255)
		}
	}
	img := image.NewGray(image.Rect(0, 0, size, size))
	img.Pix = pix
	return img
}
