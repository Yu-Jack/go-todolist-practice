package network

import (
	"fmt"
	"jack-test/internal/dataservice"
	"net/http"
	"strconv"
)

type getListResponse struct {
	Response
	dataservice.TodoList
}

func (network *Network) GetList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		r.ParseForm()
		username := r.FormValue("username")
		todoList, err := network.usecaseapi.FindByUserName(username)

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

		w.Write(ToJson(response))
	}
}

func (network *Network) CreateList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		username := r.FormValue("username")
		todoList, err := network.usecaseapi.FindByUserName(username)

		if err != nil {
			fmt.Print(err.Error())
			resp := NewFailedResponse(100)
			w.Write(resp.convertToBytes())
			return
		}

		task := r.FormValue("task")
		todoList.UpdateTodoList(task)

		err = network.usecaseapi.UpdateUsersList(username, todoList)

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

func (network *Network) DeleteList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		username := r.FormValue("username")
		todoList, err := network.usecaseapi.FindByUserName(username)

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

		err = network.usecaseapi.UpdateUsersList(username, todoList)
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
