package services

import (
	"errors"
	"time"

	"github.com/migvas/go-tasks-api/internal/models"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound    = errors.New("service: task not found")
	ErrInvalidTaskData = errors.New("service: invalid task data provided")
)

type TaskServices interface {
	GetTask(id int) (*models.Task, error)
	// CreateTask(title, description, dueDate, email string) error
	// UpdateTask(id int, name, email string) error
	// CompleteTask(id int) error
	// DeleteTask(id int) error
	// GetAllTasks() error
}

type TaskService struct {
	db *gorm.DB
}

type TaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    uint      `json:"priority"`
	Assignee    string    `json:"assignee"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTaskService(db *gorm.DB) *TaskServices {
	return &TaskService{db: db}
}

func ConvertTaskToResponse(task *models.Task) *TaskResponse {
	if task == nil {
		return nil
	}

	response := &TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	if task.Assignee.ID != 0 || task.Assignee.Name != "" {
		response.Assignee = task.Assignee.Name
	}

	if task.CreatedBy.ID != 0 || task.CreatedBy.Name != "" {
		response.CreatedBy = task.CreatedBy.Name
	}

	if task.UpdatedBy.ID != 0 || task.UpdatedBy.Name != "" {
		response.UpdatedBy = task.UpdatedBy.Name
	}

	return response
}

func (s *TaskService) GetTask(id int) (*TaskResponse, error) {
	if id <= 0 {
		return nil, ErrInvalidTaskData
	}
	var task models.Task
	result := s.db.Where("ID = ?", id).First(&task)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrTaskNotFound
		}
		return nil, ErrInvalidTaskData
	}
	taskResponse := ConvertTaskToResponse(&task)
	return taskResponse, nil
}
