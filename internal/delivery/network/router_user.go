package network

import (
	"jack-test/internal/dataservice"
	"net/http"
)

type getUsers struct {
	Response
	Users []dataservice.User `json:"users""`
}

func (network *Network) GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		usernames, _ := network.usecaseapi.FindAllUsers()

		response := getUsers{
			NewSuccessResponse(),
			usernames,
		}

		w.Write(ToJson(response))
	}
}
