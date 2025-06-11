package api

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handlers *APIHandlers) {
	mux.HandleFunc("GET /tasks/{id}", handlers.GetTask)
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("GET /users/{id}", handlers.GetUser)
}
