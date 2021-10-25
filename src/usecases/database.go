//go:generate mockgen -source=database.go -destination=mocks/database_mock.go -package mocks
package usecases

import (
	"context"

	"github.com/goltsev/image/src/models"
)

type Database interface {
	Create(context.Context, *models.Image) (int64, error)
	GetID(context.Context, int64) (*models.Image, error)
	GetRelated(context.Context, int64) ([]*models.Image, error)
}
