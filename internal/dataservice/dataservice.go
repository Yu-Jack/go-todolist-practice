package dataservice

import (
	"errors"
	"time"
)

// TodoList Reference: https://golang.org/pkg/encoding/json/#Marshal

type Task struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	CreateAt int64 `json:"create_at"`
}

type TodoList struct {
	Sequence int64 `json:"-"`
	Todo     []Task `json:"todo"`
	CreateAt int64   `json:"create_at"`
}

type UserTodoList struct {
	Users map[string]TodoList
}

func (todoList *TodoList) UpdateTodoList(task string, sequence int64) {
	newTask := Task {
		Id: sequence,
		Name: task,
		CreateAt: time.Now().UnixNano() / int64(time.Millisecond),
	}
	todoList.Todo = append(todoList.Todo, newTask)
}

func (todoList *TodoList) DeleteTodoList(taskId int64) error {
	counter := 0

	for i := range todoList.Todo {
		if todoList.Todo[i].Id != taskId {
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
