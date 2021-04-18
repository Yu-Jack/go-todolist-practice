package usecase

import "jack-test/internal/repository"

type Usecaseapi struct {
	repo repository.Repository
}

func NewUsecaseapi(repo repository.Repository) Usecaseapi {
	return Usecaseapi{
		repo: repo,
	}
}
