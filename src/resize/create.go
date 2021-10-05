package resize

import (
	"image"
	"image/color"
	"image/draw"
)

func createSquare(size int) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	draw.Draw(img, img.Bounds(), image.NewUniform(blue), image.Point{}, draw.Src)
	for i := 0; i < size; i++ {
		img.Set(i, i, red)
	}
	return img
}

func createLine(width int, height int) image.Image {
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

func createCheckers(size int) image.Image {
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
