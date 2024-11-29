package handlers

import (
	"docker_go/internal/repo"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	Router *mux.Router
	Users  repo.UserRepo
	Server *http.Server
}

func NewHandler(users repo.UserRepo) *Handler {
	log.Println("Start up handler")

	h := &Handler{
		Users: users,
	}

	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/api/v1/user/create", h.CreateUser).Methods("POST")
	h.Router.HandleFunc("/api/v1/users/{id}", h.GetUserByID).Methods("GET")
	h.Router.HandleFunc("/api/v1/users/{id}", h.UpdateUser).Methods("PUT")
	h.Router.HandleFunc("/api/v1/users/{id}", h.DeleteUser).Methods("DELETE")
}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		log.Println(err)
	}
	return nil
}
