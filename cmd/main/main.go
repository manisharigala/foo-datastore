package main

import (
	"log"
	"net/http"

	"github.com/manisharigala/Foo-Datastore/api"
)

func main() {
	srv := api.NewServer()
	err := http.ListenAndServe(":8080", srv)
	if err != nil {
		panic(err)
	}
	log.Printf("Listening on Port 8080...")
}
