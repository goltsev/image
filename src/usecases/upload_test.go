package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/goltsev/image/src/models"
	"github.com/goltsev/image/src/usecases/mocks"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	dbmock := mocks.NewMockDatabase(ctrl)
	stmock := mocks.NewMockStorage(ctrl)
	s := NewService(dbmock, stmock)

	tests := []struct {
		name       string
		params     *UploadParams
		dbmockcall *gomock.Call
		errnil     bool
	}{
		{
			"nil params",
			nil,
			nil,
			false,
		},
		{
			"error",
			&UploadParams{
				Format:   "text/plain",
				Filename: "file.txt",
			},
			dbmock.
				EXPECT().
				CreateWithAction(
					ctx,
					&models.Image{
						Format:   "text/plain",
						Filename: "file.txt",
					},
					gomock.Any()).
				Return(int64(0), errors.New("error")),
			false,
		},
		{
			"ok",
			&UploadParams{
				Format:   "text/plain",
				Filename: "file.txt",
			},
			dbmock.
				EXPECT().
				CreateWithAction(
					ctx,
					&models.Image{
						Format:   "text/plain",
						Filename: "file.txt",
					},
					gomock.Any()).
				Return(int64(1), nil),
			true,
		},
	}
	// tc = test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if err := s.Upload(ctx, tc.params); (err == nil) != tc.errnil {
				t.Error(err)
			}
		})
	}
}
