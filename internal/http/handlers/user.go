package handlers

import (
	"docker_go/internal/repo"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u repo.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Failed to decode request body:", err)
		return
	}

	err := h.Users.CreateUser(r.Context(), &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Failed to create user:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(&u); err != nil {
		log.Panicln("Failed to encode response:", err)
	}

}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.Users.GetUserByID(r.Context(), id)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		log.Panicln(err)
	}
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var u repo.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.Users.UpdateUser(r.Context(), &u, uint(id))
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(u); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Panicln(err)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.Users.DeleteUser(r.Context(), uint(id))
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
