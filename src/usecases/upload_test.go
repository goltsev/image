package usecases_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goltsev/image/src/usecases"
	"github.com/goltsev/image/src/usecases/mocks"
)

func TestUploadImage(t *testing.T) {
	ctrl := gomock.NewController(t)
	dbmock := mocks.NewMockDatabase(ctrl)
	stmock := mocks.NewMockStorage(ctrl)

	r := bytes.NewBuffer([]byte("data"))
	params := &usecases.UploadImageParams{
		ID:       10,
		Reader:   r,
		Format:   "jpeg",
		Filename: "img.jpg",
	}

	ctx := context.Background()
	dbmock.
		EXPECT().
		Create(ctx, params.Image()).
		Return(params.ID, nil)
	stmock.
		EXPECT().
		Put(ctx, params.Key(), params.Format, r).
		Return(nil)

	handler := usecases.NewUploadImageHandler(dbmock, stmock)
	if err := handler.Save(ctx, params); err != nil {
		t.Error(err)
	}
}
