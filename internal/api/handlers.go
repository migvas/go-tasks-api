package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	fmt.Printf("Task: %v\n", task)
	jsonutil.JSONResponse(w, task, http.StatusOK)
}

func (h *APIHandlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
		return
	}

	body, err := ioutil.ReadAll(r.Body) // For Go 1.16+, use io.ReadAll(r.Body)
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

}
