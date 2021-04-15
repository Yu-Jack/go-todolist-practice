package usecase

import "jack-test/internal/dataservice"

type MockRepository struct{}

func (mockRepository *MockRepository) GetList() (dataservice.UserTodoList, error) {
	list := dataservice.UserTodoList{
		Users: map[string]dataservice.TodoList{
			"user1": {
				Todo:     []dataservice.Task{{
					Id: 1,
					Name: "task",
					CreateAt: 123,
				}},
				CreateAt: 123,
			},
		},
	}
	return list, nil
}

func (mockRepository *MockRepository) SaveList(dataservice.UserTodoList) error {
// Do nothing
	return nil
}
