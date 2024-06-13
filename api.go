package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	// "github.com/elastic/go-elasticsearch/v7"
)

func server(server string) *APIServer {
	return &APIServer{
		server: server,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/fuck", httpHandler(s.masterHandler))
	log.Println("Server is ready on port: ", s.server)

	http.ListenAndServe(s.server, router)

}

func (s *APIServer) masterHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetRequest(w, r)
	}
	if r.Method == "POST" {
		return s.handlePostRequest(w, r)
	}
	if r.Method == "PUT" {
		return s.handlePutRequest(w, r)
	}
	if r.Method == "PATCH" {
		return s.handlePatchRequest(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteRequest(w, r) 
	}
	return fmt.Errorf("Method not supported yet %s", r.Method)
}

func (s *APIServer) handleGetRequest(w http.ResponseWriter, r *http.Request) error {
	doc := createDocument("20/06/2022", "SBIN", 12, 34, 45, 75)
	fmt.Println(doc)
	return WriteJson(w, http.StatusOK, doc)
}

func (s *APIServer) handlePostRequest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteRequest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlePatchRequest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlePutRequest(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// ===== Utility Handlers =======

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type APIServer struct{ server string }

type ApiError struct {
	Error string
}

func httpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {

			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
