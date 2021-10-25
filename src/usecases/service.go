package usecases

type Service struct {
	db Database
	s  Storage
}

func NewService(db Database, s Storage) *Service {
	return &Service{
		db: db,
		s:  s,
	}
}
