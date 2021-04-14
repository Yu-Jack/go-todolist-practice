package dataservice

import "errors"

// TodoList Reference: https://golang.org/pkg/encoding/json/#Marshal
type TodoList struct {
	Todo     []string `json:"todo" :"todo"`
	CreateAt string   `json:"create_at" :"create_at"`
}

type UserTodoList struct {
	Users map[string]TodoList
}

func (todoList *TodoList) UpdateTodoList(task string) {
	todoList.Todo = append(todoList.Todo, task)
}

func (todoList *TodoList) DeleteTodoList(task string) error {
	counter := 0

	for i := range todoList.Todo {
		if todoList.Todo[i] != task {
			todoList.Todo[counter] = todoList.Todo[i]
			counter++
		}
	}

	if counter == len(todoList.Todo) {
		return errors.New("not found task")
	}

	todoList.Todo = todoList.Todo[:counter]
	return nil
}
