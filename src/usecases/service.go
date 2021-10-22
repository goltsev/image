package usecases

type Service struct {
	Upload   *UploadImageHandler
	Download *DownloadImageHandler
}

func NewService(db Database, s Storage) *Service {
	return &Service{
		Upload:   NewUploadImageHandler(db, s),
		Download: NewDownloadImageHandler(db, s),
	}
}
