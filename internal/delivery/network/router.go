package network

import (
	"net/http"
)

func BuildRouter(mux *http.ServeMux) {
	mux.HandleFunc("/todo-list/get", GetList)
	mux.HandleFunc("/todo-list/create", CreateList)
	//mux.HandleFunc("/todo-list/delete", DeleteList)
}
