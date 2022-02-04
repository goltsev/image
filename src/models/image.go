package models

import (
	"fmt"
	"image"
)

type Image struct {
	ID       int64
	Filename string
	Format   string
	Image    image.Image
	Related  *Image
}

func (img *Image) RelatedID() (int64, error) {
	if img.Related == nil {
		return 0, fmt.Errorf("no related image")
	}
	return img.Related.ID, nil
}
