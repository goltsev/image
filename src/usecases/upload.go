package usecases

import (
	"context"
	"fmt"
	"io"

	"github.com/goltsev/image/src/models"
)

// UploadImageParams is used to provide parameters for invoking
// UploadImageHandler command.
type UploadParams struct {
	ID       int64
	Reader   io.Reader
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
	img := params.Image()
	id, err := s.storeDB(ctx, img)
	if err != nil {
		return err
	}
	params.ID = id
	if err := s.storeS3(ctx, params); err != nil {
		return err
	}
	return nil
}

func (s *Service) storeDB(ctx context.Context, img *models.Image) (int64, error) {
	return s.db.Create(ctx, img)
}

func (s *Service) storeS3(ctx context.Context, params *UploadParams) error {
	key := params.Key()
	return s.s.Put(ctx, key, params.Format, params.Reader)
}
