package usecase

import "jack-test/internal/dataservice"

func (usecaseapi *usecaseapi) FindAllUsers() (username []dataservice.User, error error) {
	var users []dataservice.User
	userList, _ := usecaseapi.repo.GetList()

	for key, user := range userList.Users {
		user := dataservice.User{
			Name:     key,
			CreateAt: user.CreateAt,
		}
		users = append(users, user)
	}

	return users, nil
}
