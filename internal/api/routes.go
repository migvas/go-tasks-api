package api

import (
	"net/http"
)

func SetupRoutes(mux *http.ServeMux, handlers *APIHandlers) {
	// Here, we're not 'calling' GetTask, we're providing it as a function value
	// to mux.HandleFunc. The 'mux' will call it later when a request matches.
	mux.HandleFunc("GET /tasks/{id}", handlers.GetTask)
}
