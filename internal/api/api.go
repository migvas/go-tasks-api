package api

import "github.com/migvas/go-tasks-api/internal/services"

// APIHandlers holds all the HTTP handlers for your API.
// It receives service interfaces as dependencies.
type APIHandlers struct {
	TaskServices services.TaskServices // Use the interface
	// Add other services if you have them (e.g., ProductService, OrderService)
}

// NewAPIHandlers creates a new instance of APIHandlers.
// It receives initialized service implementations.
func NewAPIHandlers(taskService services.TaskServices) *APIHandlers {
	return &APIHandlers{
		TaskServices: taskService,
	}
}
