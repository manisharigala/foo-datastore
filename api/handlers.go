package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *Server) homepage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, welcome to Foo Datastore"))
	}
}

func (s *Server) createDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f foo
		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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

func (s *Server) getDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

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

func (s *Server) removeDocument() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		_, found := s.records[id]
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		delete(s.records, id)
		w.WriteHeader(http.StatusNoContent)
	}
}
