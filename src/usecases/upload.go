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

func (h *UploadImageHandler) Do(ctx context.Context, v interface{}) error {
	params, ok := v.(*UploadImageParams)
	if !ok {
		return NewErrWrongType(params, v)
	}
	return h.save(ctx, params)
}

func (h *UploadImageHandler) save(ctx context.Context, params *UploadImageParams) error {
	img := h.imgFromParams(params)
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

func (u *UploadImageHandler) imgFromParams(params *UploadImageParams) *models.Image {
	return &models.Image{
		Filename: params.Filename,
		Format:   params.Format,
	}
}

func (u *UploadImageHandler) storeS3(ctx context.Context, params *UploadImageParams) error {
	key := u.key(params)
	return u.storage.Put(ctx, key, params.Format, params.Reader)
}

func (h *UploadImageHandler) key(params *UploadImageParams) string {
	return fmt.Sprintf("%d/%s", params.ID, params.Filename)
}
