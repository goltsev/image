//go:generate mockgen -source=storage.go -destination=mocks/storage_mock.go -package mocks
package usecases

import (
	"context"
	"io"
)

type Storage interface {
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Put(ctx context.Context, key string, format string, r io.Reader) error
}
