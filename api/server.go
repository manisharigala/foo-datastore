package api

import (
	"github.com/gorilla/mux"
)

// Server data structure
type Server struct {
	// mux router matches incoming requests to their respective handler
	*mux.Router

	// Hash map to store records in memory
	records map[string]Foo
}

// Get new server
func NewServer() *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		records: map[string]Foo{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	// Test API
	s.HandleFunc("/", s.homepage()).Methods("GET")

	// POST API to create new Foo record
	s.HandleFunc("/foo", s.createFooRecord()).Methods("POST")

	// GET API to fetch Foo record from memory
	s.HandleFunc("/foo/{id}", s.getFooRecord()).Methods("GET")

	// DELETE API to delete Foo record from memory
	s.HandleFunc("/foo/{id}", s.removeFooRecord()).Methods("DELETE")
}
