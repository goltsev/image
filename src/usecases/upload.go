package usecases

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/goltsev/image/src/models"
)

// UploadImageParams is used to provide parameters for invoking
// UploadImageHandler command.
type UploadParams struct {
	ID       int64
	Source   io.Reader
	Format   string
	Filename string
}

// Key creates key from ID and Filename
func (p *UploadParams) Key() string {
	return fmt.Sprintf("%d/%s", p.ID, p.Filename)
}

// Image creates image
func (p *UploadParams) Image() *models.Image {
	return &models.Image{
		ID:       p.ID,
		Filename: p.Filename,
		Format:   p.Format,
	}
}

// Upload saves data from reader to application.
func (s *Service) Upload(ctx context.Context, params *UploadParams) error {
	if params == nil {
		return errors.New("params cannot be nil")
	}
	_, err := s.db.CreateWithAction(ctx, params.Image(), func(id int64) error {
		params.ID = id
		return s.s.Put(ctx, params.Key(), params.Format, params.Source)
	})
	return err
}
