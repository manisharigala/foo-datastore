package api

import (
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router

	documents []Document
}

func NewServer() *Server {
	s := &Server{
		Router:    mux.NewRouter(),
		documents: []Document{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/foo", s.createDocument()).Methods("POST")
	s.HandleFunc("/foo/{id}", s.getDocument()).Methods("GET")
	s.HandleFunc("/foo/{id}", s.removeDocument()).Methods("DELETE")
}
