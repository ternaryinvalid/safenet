package service

type ApiService struct {
	repository repository
}

type repository interface {
}

func New(repo repository) *ApiService {
	return &ApiService{
		repository: repo,
	}
}
