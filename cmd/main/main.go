package main

import (
	"log"
	"net/http"

	"github.com/manisharigala/Foo-Datastore/api"
)

func main() {
	// Get Server
	srv := api.NewServer()
	log.Println("Starting server on Port 8080...")

	// Listen and Serve
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		panic(err)
	}
}
