package usecase

import "jack-test/internal/dataservice"

type MockRepository struct{}

func (mockRepository *MockRepository) GetList() (dataservice.UserTodoList, error) {
	list := dataservice.UserTodoList{
		Users: map[string]dataservice.TodoList{
			"user1": {
				Todo:     []string{"123"},
				CreateAt: "123",
			},
		},
	}
	return list, nil
}
