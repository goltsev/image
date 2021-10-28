//go:generate mockgen -source=database.go -destination=mocks/database_mock.go -package mocks
package usecases

import (
	"context"

	"github.com/goltsev/image/src/models"
)

type Database interface {
	Create(context.Context, *models.Image) (int64, error)
	CreateWithAction(context.Context, *models.Image, func(int64) error) (int64, error)
	GetByID(context.Context, int64) (*models.Image, error)
	GetByRelated(context.Context, int64) ([]*models.Image, error)
}
