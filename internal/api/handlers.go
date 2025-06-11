package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/migvas/go-tasks-api/internal/services"
	"github.com/migvas/go-tasks-api/pkg/jsonutil"
)

func (h *APIHandlers) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // Go 1.22+
	if idStr == "" {
		jsonutil.ErrorResponse(w, "Task ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonutil.ErrorResponse(w, "Invalid task ID format", http.StatusBadRequest)
		return
	}

	task, err := h.TaskServices.GetTask(id) // Call the service layer
	if err != nil {
		if errors.Is(err, services.ErrTaskNotFound) {
			jsonutil.ErrorResponse(w, "Task not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, services.ErrInvalidTaskData) {
			jsonutil.ErrorResponse(w, "Invalid task ID", http.StatusBadRequest)
			return
		}
		jsonutil.ErrorResponse(w, "Failed to retrieve task", http.StatusInternalServerError)
		return
	}

	jsonutil.JSONResponse(w, task, http.StatusOK)
}

func (h *APIHandlers) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		log.Printf("Error reading JSON body: %v", err)
		return
	}
	defer r.Body.Close()

	var task services.TaskInput

	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	taskResponse, err := h.TaskServices.CreateTask(&task)

	if err != nil {
		if errors.Is(err, services.ErrInvalidUser) {
			jsonutil.ErrorResponse(w, "Invalid user", http.StatusBadRequest)
			return
		}
		if errors.Is(err, services.ErrCreateTask) {
			jsonutil.ErrorResponse(w, "Error creating task", http.StatusInternalServerError)
			return
		}
	}
	jsonutil.JSONResponse(w, taskResponse, http.StatusOK)
}

func (h *APIHandlers) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // Go 1.22+
	if idStr == "" {
		jsonutil.ErrorResponse(w, "User ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonutil.ErrorResponse(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	user, err := h.UserServices.GetUser(id) // Call the service layer
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			jsonutil.ErrorResponse(w, "User not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, services.ErrInvalidUserData) {
			jsonutil.ErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		jsonutil.ErrorResponse(w, "Failed to retrieve user", http.StatusInternalServerError)
		return
	}

	jsonutil.JSONResponse(w, user, http.StatusOK)
}

func (h *APIHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		log.Printf("Error reading JSON body: %v", err)
		return
	}
	defer r.Body.Close()

	var user services.UserInput

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		log.Printf("Error unmarshalling JSON: %v", err)
		return
	}

	userResponse, err := h.UserServices.CreateUser(&user)

	if err != nil {
		if errors.Is(err, services.ErrInvalidEmail) {
			jsonutil.ErrorResponse(w, "Invalid email", http.StatusBadRequest)
			return
		}
		if errors.Is(err, services.ErrCreateUser) {
			jsonutil.ErrorResponse(w, "Error creating user", http.StatusInternalServerError)
			return
		}
	}
	jsonutil.JSONResponse(w, userResponse, http.StatusOK)
}
