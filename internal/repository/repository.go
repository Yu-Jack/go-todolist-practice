package repository

import "jack-test/internal/dataservice"

type Repository interface {
	GetList() (dataservice.UserTodoList, error)
}

type repository struct{}

func NewRepository() repository {
	return repository{}
}
