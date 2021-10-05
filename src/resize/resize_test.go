package resize

import (
	"fmt"
	"image"
	"os"
	"testing"
)

var (
	dirname = "testdata"
)

func TestNaive(t *testing.T) {
	cs := []struct {
		name     string
		img      image.Image
		width    int
		height   int
		expected image.Image
	}{
		{
			"2x1-line",
			createLine(2, 1),
			3, 1,
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 3, 1))
				img.Pix = []uint8{
					255, 0, 0, 255,
					255, 0, 0, 255,
					0, 0, 255, 255,
				}
				return img
			}(),
		},
		{
			"2x2-line",
			createLine(2, 2),
			3, 2,
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 3, 2))
				img.Pix = []uint8{
					255, 0, 0, 255,
					255, 0, 0, 255,
					0, 0, 255, 255,
					255, 0, 0, 255,
					255, 0, 0, 255,
					0, 0, 255, 255,
				}
				return img
			}(),
		},
		{
			"checkers",
			createCheckers(2),
			4, 4,
			func() image.Image {
				img := image.NewGray(image.Rect(0, 0, 4, 4))
				img.Pix = []uint8{
					255, 255, 0, 0,
					255, 255, 0, 0,
					0, 0, 255, 255,
					0, 0, 255, 255,
				}
				return img
			}(),
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := Naive(c.img, c.width, c.height)
			if !equalImage(c.expected, got) {
				t.Errorf("expected: %v; got: %v;\n", c.expected, got)
			}
		})
	}
}

func TestBilinear(t *testing.T) {
	cs := []struct {
		name     string
		img      image.Image
		width    int
		height   int
		expected image.Image
	}{
		{
			"2x1-line",
			createLine(2, 1),
			3, 1,
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 3, 1))
				img.Pix = []uint8{
					255, 0, 0, 255,
					127, 0, 127, 255,
					0, 0, 255, 255,
				}
				return img
			}(),
		},
		{
			"2x2-line",
			createLine(2, 2),
			3, 2,
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 3, 2))
				img.Pix = []uint8{
					255, 0, 0, 255,
					127, 0, 127, 255,
					0, 0, 255, 255,
					255, 0, 0, 255,
					127, 0, 127, 255,
					0, 0, 255, 255,
				}
				return img
			}(),
		},
		{
			"checkers",
			createCheckers(2),
			3, 3,
			func() image.Image {
				img := image.NewRGBA(image.Rect(0, 0, 3, 3))
				img.Pix = []uint8{
					255, 255, 255, 255,
					127, 127, 127, 255,
					0, 0, 0, 255,
					127, 127, 127, 255,
					127, 127, 127, 255,
					127, 127, 127, 255,
					0, 0, 0, 255,
					127, 127, 127, 255,
					255, 255, 255, 255,
				}
				return img
			}(),
		},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := Bilinear(c.img, c.width, c.height)
			if !equalImage(c.expected, got) {
				t.Errorf("expected: %v; got: %v;\n", c.expected, got)
				writetestimage(got, "bilinear-got-"+c.name, "png")
				writetestimage(c.expected, "bilinear-expected-"+c.name, "png")
				writetestimage(c.img, "bilinear-input-"+c.name, "png")
			}
		})
	}
}

func writetestimage(img image.Image, name string, format string) error {
	_ = os.Mkdir(dirname, os.ModeDir) // ignore error if directory exists
	filename := fmt.Sprintf("./testdata/%s.%s", name, format)
	err := WriteFile(img, filename, format)
	if err != nil {
		return err
	}
	return nil
}
