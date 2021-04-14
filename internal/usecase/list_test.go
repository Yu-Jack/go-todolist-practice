package usecase

import (
	"testing"
)

func Test_Usecaseapi_FindByUserName_should_return_the_todoList_of_user_with_correct_username(t *testing.T) {
	mockRepo := MockRepository{}
	usercaseapi := NewUsecaseapi(&mockRepo)

	userList, err := usercaseapi.FindByUserName("user1")

	if err != nil {
		t.Errorf("Shoud get the nil error, but got error")
		return
	}

	if len(userList.Todo) == 0 {
		t.Errorf("Should get one item in todo list, but get zero item")
		return
	}
}

func Test_Usecaseapi_FindByUserName_should_return_user_not_found_error_with_incorrect_username(t *testing.T) {
	mockRepo := MockRepository{}
	usercaseapi := NewUsecaseapi(&mockRepo)

	_, err := usercaseapi.FindByUserName("user2")

	if err.Error() != "user is not found" {
		t.Errorf("Shoud get the error \"user is not found\" message, but got wrong error message")
		return
	}
}
