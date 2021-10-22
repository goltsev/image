package resize

import (
	"bytes"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

func ReadFile(filename string) (image.Image, string, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer infile.Close()
	img, format, err := image.Decode(infile)
	if err != nil {
		return nil, "", err
	}
	return img, format, nil
}

// WriteFile writes image to a file using specified file format.
// png and jpeg are supported.
func WriteFile(img image.Image, filename string, format string) error {
	if img == nil {
		return errors.New("image is nil")
	}
	buf := &bytes.Buffer{}
	switch format {
	case "png":
		if err := png.Encode(buf, img); err != nil {
			return err
		}
	case "jpeg":
		if err := jpeg.Encode(buf, img, &jpeg.Options{Quality: 100}); err != nil {
			return err
		}
	}
	return writefile(buf, filename)
}

func writefile(buf io.Reader, filename string) error {
	outfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	io.Copy(outfile, buf)
	return outfile.Close()
}
