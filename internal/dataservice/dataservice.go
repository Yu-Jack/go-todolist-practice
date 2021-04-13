package dataservice

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
