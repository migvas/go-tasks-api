package api

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handlers *APIHandlers) {
	// Task routes
	mux.HandleFunc("GET /tasks/{id}", handlers.GetTask)
	mux.HandleFunc("POST /tasks", handlers.CreateTask)

	// User routes
	mux.HandleFunc("POST /users", handlers.CreateUser)
	mux.HandleFunc("GET /users/{id}", handlers.GetUser)
}
