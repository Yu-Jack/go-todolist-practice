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
		username := r.FormValue("username")

		usecaseapi := usecase.NewUsecaseapi()
		todoList, err := usecaseapi.FindByUserName(username)

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
		r.ParseForm()

		username := r.FormValue("username")
		usecaseapi := usecase.NewUsecaseapi()
		todoList, err := usecaseapi.FindByUserName(username)

		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		task := r.FormValue("task")
		todoList.UpdateTodoList(task)

		err = usecaseapi.UpdateUsersList(username, todoList)

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
