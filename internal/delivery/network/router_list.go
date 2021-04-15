package network

import (
	"encoding/json"
	"fmt"
	"jack-test/internal/repository"
	"jack-test/internal/usecase"
	"net/http"
	"strconv"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		username := r.FormValue("username")

		repo := repository.NewRepository()
		usecaseapi := usecase.NewUsecaseapi(&repo)
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
		repo := repository.NewRepository()
		usecaseapi := usecase.NewUsecaseapi(&repo)
		todoList, err := usecaseapi.FindByUserName(username)

		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		task := r.FormValue("task")
		todoList.Sequence++
		todoList.UpdateTodoList(task, todoList.Sequence)

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
		r.ParseForm()

		username := r.FormValue("username")
		repo := repository.NewRepository()
		usecaseapi := usecase.NewUsecaseapi(&repo)
		todoList, err := usecaseapi.FindByUserName(username)

		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		taskId, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
		err = todoList.DeleteTodoList(taskId)
		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		err = usecaseapi.UpdateUsersList(username, todoList)
		if err != nil {
			fmt.Print(err.Error())
			fmt.Fprintf(w, "failed")
			return
		}

		fmt.Fprintf(w, "success")
	}
}
