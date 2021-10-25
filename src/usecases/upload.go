package usecases

import (
	"context"
	"fmt"
	"io"

	"github.com/goltsev/image/src/models"
)

// UploadImageParams is used to provide parameters for invoking
// UploadImageHandler command.
type UploadImageParams struct {
	ID       int64
	Reader   io.Reader
	Format   string
	Filename string
}

// Key creates key from ID and Filename
func (p *UploadImageParams) Key() string {
	return fmt.Sprintf("%d/%s", p.ID, p.Filename)
}

// Image creates image
func (p *UploadImageParams) Image() *models.Image {
	return &models.Image{
		ID:       p.ID,
		Filename: p.Filename,
		Format:   p.Format,
	}
}

// UploadImageHandler is used to upload files to application.
type UploadImageHandler struct {
	db      Database
	storage Storage
}

func NewUploadImageHandler(db Database, s Storage) *UploadImageHandler {
	return &UploadImageHandler{
		db:      db,
		storage: s,
	}
}

// Do calls Save with casted parameters.
// It receives context.Context and UploadImageParams as parameters.
func (h *UploadImageHandler) Do(ctx context.Context, v interface{}) error {
	params, ok := v.(*UploadImageParams)
	if !ok {
		return NewErrWrongType(params, v)
	}
	return h.Save(ctx, params)
}

// Save saves data from reader to application.
func (h *UploadImageHandler) Save(ctx context.Context, params *UploadImageParams) error {
	img := params.Image()
	id, err := h.storeDB(ctx, img)
	if err != nil {
		return err
	}
	params.ID = id
	if err := h.storeS3(ctx, params); err != nil {
		return err
	}
	return nil
}

func (u *UploadImageHandler) storeDB(ctx context.Context, img *models.Image) (int64, error) {
	return u.db.Create(ctx, img)
}

func (u *UploadImageHandler) storeS3(ctx context.Context, params *UploadImageParams) error {
	key := params.Key()
	return u.storage.Put(ctx, key, params.Format, params.Reader)
}
