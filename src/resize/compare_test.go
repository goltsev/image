package resize

import (
	"image"
	"image/color"
	"testing"
)

func TestEqualImage(t *testing.T) {
	cs := []struct {
		name     string
		img0     image.Image
		img1     image.Image
		expected bool
	}{
		{
			"both nil",
			nil, nil,
			true,
		},
		{
			"nil notnil",
			nil, image.NewGray(image.Rect(0, 0, 1, 1)),
			false,
		},
		{
			"notnil nil",
			image.NewGray(image.Rect(0, 0, 1, 1)), nil,
			false,
		},
		{
			"same",
			func() image.Image {
				img := image.NewGray(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.Gray{127})
				return img
			}(),
			func() image.Image {
				img := image.NewGray(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.Gray{127})
				return img
			}(),
			true,
		},
		{
			"different color",
			func() image.Image {
				img := image.NewGray(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.Gray{127})
				return img
			}(),
			func() image.Image {
				img := image.NewGray(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.Gray{200})
				return img
			}(),
			false,
		},
		{
			"different size",
			image.NewGray(image.Rect(0, 0, 1, 1)),
			image.NewGray(image.Rect(0, 0, 1, 2)),
			false,
		},
		{
			"different color scale",
			func() image.Image {
				img := image.NewCMYK(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.RGBA{A: 255})
				return img
			}(),
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 1, 1))
				img.Set(0, 0, color.RGBA{A: 255})
				return img
			}(),
			true,
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := equalImage(c.img0, c.img1)
			if got != c.expected {
				t.Errorf("expected: %v; got: %v;\n", c.expected, got)
			}
		})
	}
}
