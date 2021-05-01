package usecase

import (
	"jack-test/internal/dataservice"
	"jack-test/internal/repository"
)

type Usecaseapi interface {
	FindByUserName(username string) (dataservice.TodoList, error)
	UpdateUsersList(username string, todoList dataservice.TodoList) error
	FindAllUsers() (username []dataservice.User, error error)
}

type usecaseapi struct {
	repo repository.Repository
}

func NewUsecaseapi(repo repository.Repository) Usecaseapi {
	return &usecaseapi{
		repo: repo,
	}
}
