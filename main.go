package main

import (
	"jack-test/internal/delivery/network"
	"jack-test/internal/repository"
	"jack-test/internal/usecase"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	repo := repository.NewRepository()
	usecaseapi := usecase.NewUsecaseapi(&repo)
	network := network.NewNetwork(usecaseapi)
	network.BuildRouter(mux)

	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
