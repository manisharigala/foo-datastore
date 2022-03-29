package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// API to check if server is up
func (s *Server) homepage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, welcome to Foo Datastore"))
	}
}

// Create new foo record and store in memory
func (s *Server) createFooRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f Foo
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// create new uuid for new record
		f.ID = uuid.New().String()
		s.records[f.ID] = f

		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Fetching foo record from memory
func (s *Server) getFooRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		// finding record from memory
		f, found := s.records[id]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Removing Foo record from memory
func (s *Server) removeFooRecord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		// finding record from memory
		_, found := s.records[id]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		// deleting record from memory
		delete(s.records, id)
		w.WriteHeader(http.StatusNoContent)
	}
}
