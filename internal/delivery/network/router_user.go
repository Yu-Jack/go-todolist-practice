package network

import (
	"fmt"
	"jack-test/internal/dataservice"
	"jack-test/internal/repository"
	"jack-test/internal/usecase"
	"net/http"
)

type getUsers struct {
	Response
	Users []dataservice.User `json:"users""`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		repo := repository.NewRepository()
		usecaseapi := usecase.NewUsecaseapi(&repo)
		usernames, _ := usecaseapi.FindAllUsers()

		response := getUsers{
			NewSuccessResponse(),
			usernames,
		}

		fmt.Println(response)

		w.Write(ToJson(response))
	}
}
