package usecase

import "jack-test/internal/repository"

type usecaseapi struct {
	repo repository.Repository
}

func NewUsecaseapi(repo repository.Repository) usecaseapi {
	return usecaseapi{
		repo: repo,
	}
}
