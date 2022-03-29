package main

import (
	"net/http"

	"github.com/manisharigala/Foo-Datastore/api"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
