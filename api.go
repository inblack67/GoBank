package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listAddress string
	storage     Storage
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error // ours

type APIError struct {
	Error string
}

func writeJSON(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func NewAPIServer(listAddress string, storage Storage) *APIServer {
	return &APIServer{
		listAddress: listAddress,
		storage:     storage,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandleFunc(s.handleGetAccount))

	println("Server running on port:", s.listAddress)

	http.ListenAndServe(s.listAddress, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	return fmt.Errorf("method Not allowed")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return err
	}
	account, err := s.storage.getAccount(id)
	if err != nil {
		return err
	}
	writeJSON(w, http.StatusOK, account)
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	payload := new(CreateAccountReq)
	if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
		return err
	}
	account := NewAccount(payload.Username)

	newAcc, err := s.storage.createAccount(account)

	if err != nil {
		return nil
	}

	return writeJSON(w, http.StatusCreated, newAcc)
}
