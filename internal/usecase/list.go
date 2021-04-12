package usecase

import (
	"encoding/json"
	"errors"
	"jack-test/internal/repository"
)

// Reference: https://golang.org/pkg/encoding/json/#Marshal
type TodoList struct {
	Todo     []string `json:"todo"`
	CreateAt string   `json:"create_at"`
}

type UserTodoList struct {
	Users map[string]TodoList
}

func (userList *UserTodoList) FindByUserName(username string) (TodoList, error) {
	content, err := repository.GetList()
	if err != nil {
		return userList.Users[""], err
	}

	json.Unmarshal(content, &userList.Users)
	user, ok := userList.Users[username]
	if !ok {
		return userList.Users[""], errors.New("user is not found")
	}

	return user, nil
}

func (userList *UserTodoList) UpdateUsersList() error {
	jsonData, err := json.MarshalIndent(userList.Users, "", "    ")
	if err != nil {
		return err
	}
	repository.SaveList(jsonData)
	return nil
}

func (todoList *TodoList) UpdateTodoList(task string) *TodoList {
	todoList.Todo = append(todoList.Todo, task)
	return todoList
}
