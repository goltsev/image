package resize

import (
	"image"
	"image/color"
)

// equalImage compares two images and returns true if they are the same
func equalImage(img0, img1 image.Image) bool {
	if img0 == nil || img1 == nil {
		return img0 == img1
	}
	if img0.Bounds() != img1.Bounds() {
		return false
	}
	r0 := img0.Bounds()
	for y := r0.Min.Y; y < r0.Max.Y; y++ {
		for x := r0.Min.X; x < r0.Max.X; x++ {
			if !equalColor(img0.At(x, y), img1.At(x, y)) {
				return false
			}
		}
	}
	return true
}

// equalColor compares two colors and returns true if they are the same
func equalColor(c0, c1 color.Color) bool {
	if c0 == nil || c1 == nil {
		return c0 == c1
	}
	r0, g0, b0, _ := c0.RGBA()
	r1, g1, b1, _ := c1.RGBA()
	if r0 == r1 && g0 == g1 && b0 == b1 {
		return true
	}
	return false
}
