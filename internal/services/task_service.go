package services

import (
	"errors"

	"github.com/migvas/go-tasks-api/internal/models"
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

type TaskService struct{}

func NewTaskService() TaskServices {
	return &TaskService{}
}

func (s *TaskService) GetTask(id int) (*models.Task, error) {
	if id <= 0 {
		return nil, ErrInvalidTaskData
	}
	task := &models.Task{ID: int64(id)}
	return task, nil
}
