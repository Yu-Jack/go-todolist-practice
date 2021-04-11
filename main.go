package main

import (
	"jack-test/internal/delivery/network"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	network.BuildRouter(mux)
	err := http.ListenAndServe(":7070", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
