package api

import "github.com/migvas/go-tasks-api/internal/services"

// APIHandlers holds all the HTTP handlers for your API.
// It receives service interfaces as dependencies.
type APIHandlers struct {
	TaskServices services.TaskServices
	UserServices services.UserServices
}

// NewAPIHandlers creates a new instance of APIHandlers.
// It receives initialized service implementations.
func NewAPIHandlers(taskServices services.TaskServices, userServices services.UserServices) *APIHandlers {
	return &APIHandlers{
		TaskServices: taskServices,
		UserServices: userServices,
	}
}
