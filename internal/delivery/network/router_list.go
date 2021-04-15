package network

import (
	"encoding/json"
	"fmt"
	"jack-test/internal/dataservice"
	"jack-test/internal/repository"
	"jack-test/internal/usecase"
	"net/http"
	"strconv"
)

type getListResponse struct {
	Response
	dataservice.TodoList
}

func GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		username := r.FormValue("username")

		repo := repository.NewRepository()
		usecaseapi := usecase.NewUsecaseapi(&repo)
		todoList, err := usecaseapi.FindByUserName(username)

		if err != nil {
			fmt.Print(err.Error())
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		response := getListResponse{
			NewSuccessResponse(),
			todoList,
		}

		responseJson, _ := json.MarshalIndent(response, "", "")
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
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		task := r.FormValue("task")
		todoList.UpdateTodoList(task)

		err = usecaseapi.UpdateUsersList(username, todoList)

		if err != nil {
			fmt.Print(err.Error())
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		resp := NewSuccessResponse()
		w.Write(resp.convertToBytes())
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
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		taskId, _ := strconv.ParseInt(r.FormValue("id"), 10, 0)
		err = todoList.DeleteTodoList(taskId)
		if err != nil {
			fmt.Print(err.Error())
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		err = usecaseapi.UpdateUsersList(username, todoList)
		if err != nil {
			fmt.Print(err.Error())
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		resp := NewSuccessResponse()
		w.Write(resp.convertToBytes())
	}
}
