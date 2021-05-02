package network

import (
	"net/http"
)

func (network *Network) BuildRouter(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("./public"))
	mux.Handle("/", fs)
	mux.HandleFunc("/todo-list/get", network.GetList)
	mux.HandleFunc("/todo-list/create", network.CreateList)
	mux.HandleFunc("/todo-list/delete", network.DeleteList)
	mux.HandleFunc("/users", network.GetUsers)
}
