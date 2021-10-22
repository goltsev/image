package usecases

import (
	"context"
	"io"

	"github.com/goltsev/image/src/models"
)

type Database interface {
	Create(context.Context, *models.Image) (int64, error)
	ReadID(context.Context, int64) (*models.Image, error)
	ReadRelated(context.Context, int64) ([]*models.Image, error)
}

type Storage interface {
	Get(context.Context, string) (io.ReadCloser, error)
	Put(context.Context, string, string, io.Reader) error
}
