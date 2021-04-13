package usecase

import (
	"testing"
)

func Test_Usecaseapi_FindByUserName_should_return_the_todoList_of_user_with_correct_username(t *testing.T) {
	mockRepo := MockRepository{}
	usercaseapi := NewUsecaseapi(&mockRepo)

	userList, err := usercaseapi.FindByUserName("user1")

	if err != nil {
		t.Errorf("FindByUserName get error")
		return
	}

	if len(userList.Todo) == 0 {
		t.Errorf("FindByUserName should get one item in todo list")
		return
	}
}
