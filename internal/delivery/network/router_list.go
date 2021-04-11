package network

import (
	"encoding/json"
	"fmt"
	"jack-test/internal/usecase"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()

		var userTodoLists usecase.UserTodoList
		username := r.FormValue("username")
		todoList, err := userTodoLists.FindByUserName(username)
		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		responseJson, _ := json.MarshalIndent(todoList, "", "")
		w.Write(responseJson)
	}
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var userTodoLists usecase.UserTodoList

		r.ParseForm()
		username := r.FormValue("username")
		todoList, err := userTodoLists.FindByUserName(username)
		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		task := r.FormValue("task")
		userTodoLists.Users[username] = *todoList.UpdateTodoList(task)

		err = userTodoLists.UpdateUsersList()
		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "success")
			return
		}

		fmt.Fprintf(w, "success")
	}
}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// TODO: Implement delete list
		fmt.Fprintf(w, "this is delete list")
	}
}
