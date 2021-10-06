package resize

import (
	"image"
	"testing"
)

func _TestCreateLine(t *testing.T) {
	if err := writetestimage(CreateLine(10, 2), "line", "png"); err != nil {
		t.Fatal(err)
	}
}

func TestCreateCheckers(t *testing.T) {
	cs := []struct {
		name     string
		size     int
		expected image.Image
	}{
		{
			"0",
			0,
			nil,
		},
		{
			"1x1",
			1,
			func() *image.Gray {
				img := image.NewGray(image.Rect(0, 0, 1, 1))
				img.Pix = []uint8{255}
				return img
			}(),
		},
		{
			"2x2",
			2,
			func() *image.Gray {
				img := image.NewGray(image.Rect(0, 0, 2, 2))
				img.Pix = []uint8{255, 0, 0, 255}
				return img
			}(),
		},
		{
			"4x4",
			4,
			func() *image.Gray {
				img := image.NewGray(image.Rect(0, 0, 4, 4))
				img.Pix = []uint8{
					255, 0, 255, 0,
					0, 255, 0, 255,
					255, 0, 255, 0,
					0, 255, 0, 255,
				}
				return img
			}(),
		},
	}
	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := CreateCheckers(c.size)
			if !equalImage(got, c.expected) {
				t.Errorf("expected: %v; got: %v;\n", c.expected, got)
			}
		})
	}
}
