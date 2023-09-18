package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	dao TodoDao
}

func NewTodoHandler(dao TodoDao, r *mux.Router) *TodoHandler {
	h := &TodoHandler{dao: dao}
	r.HandleFunc("/api/v1/owners/{ownerID}/todos", h.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/owners/{ownerID}/todos/{todoID}", h.Get).Methods("GET")
	r.HandleFunc("/api/v1/owners/{ownerID}/todos", h.Create).Methods("POST")
	r.HandleFunc("/api/v1/owners/{ownerID}/todos/{todoID}", h.Update).Methods("PUT")
	r.HandleFunc("/api/v1/owners/{ownerID}/todos/{todoID}/complete", h.Complete).Methods("PUT")
	r.HandleFunc("/api/v1/owners/{ownerID}/todos/{todoID}", h.Delete).Methods("DELETE")
	return h
}

func (h *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ownerID := mux.Vars(r)["ownerID"]
	todos, err := h.dao.GetAll(ownerID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (h *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["todoID"]
	todo, err := h.dao.Get(todoID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {
	ownerID := mux.Vars(r)["ownerID"]
	todo := &Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.Owner_ID = ownerID
	err = h.dao.Create(todo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TodoHandler) Update(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["todoID"]
	todo := &Todo{}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo.ID = todoID
	err = h.dao.Update(todo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) Complete(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["todoID"]
	err := h.dao.Done(todoID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	todoID := mux.Vars(r)["todoID"]
	err := h.dao.Delete(todoID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}