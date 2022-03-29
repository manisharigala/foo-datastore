package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Document struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

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

func (s *Server) createDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var d Document
		if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		d.ID = uuid.New().String()
		s.documents = append(s.documents, d)

		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(d); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var d Document

		for _, doc := range s.documents {
			if doc.ID == id {
				d = doc
			}
		}

		if (d == Document{}) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(d); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		for i, doc := range s.documents {
			if doc.ID == id {
				s.documents = append(s.documents[:i], s.documents[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
	}
}
