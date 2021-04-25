package usecase

import (
	"errors"
	"jack-test/internal/dataservice"
	"jack-test/internal/repository"
	"sync"
)

var mu sync.Mutex

func (usecaseapi *Usecaseapi) FindByUserName(username string) (dataservice.TodoList, error) {
	mu.Lock()
	defer mu.Unlock()
	userList, err := usecaseapi.repo.GetList()

	if err != nil {
		return userList.Users[""], err
	}

	user, ok := userList.Users[username]
	if !ok {
		return userList.Users[""], errors.New("user is not found")
	}

	return user, nil
}

func (usecaseapi *Usecaseapi) UpdateUsersList(username string, todoList dataservice.TodoList) error {
	mu.Lock()
	defer mu.Unlock()
	repo := repository.NewRepository()
	userList, err := repo.GetList()

	if err != nil {
		return err
	}

	userList.Users[username] = todoList
	err = repo.SaveList(userList)

	if err != nil {
		return err
	}

	return nil
}
