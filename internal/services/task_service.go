package services

import (
	"errors"
	"fmt"
)

var (
	ErrTaskNotFound    = errors.New("service: task not found")
	ErrInvalidTaskData = errors.New("service: invalid task data provided")
)

type TaskServices interface {
	GetTask(id int) (*MockTask, error)
	// CreateTask(title, description, dueDate, email string) error
	// UpdateTask(id int, name, email string) error
	// CompleteTask(id int) error
	// DeleteTask(id int) error
	// GetAllTasks() error
}

type TaskService struct{}

type MockTask struct {
	ID int `json:"id"`
}

func NewTaskService() TaskServices {
	return &TaskService{}
}

func (s *TaskService) GetTask(id int) (*MockTask, error) {
	if id <= 0 {
		return nil, ErrInvalidTaskData
	}
	fmt.Printf("This is task with id %d\n", id)
	task := &MockTask{ID: id}
	return task, nil
}
