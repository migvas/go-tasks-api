package services

import (
	"errors"
	"log"
	"time"

	"github.com/migvas/go-tasks-api/internal/models"
	"gorm.io/gorm"
)

var (
	ErrTaskNotFound    = errors.New("service: task not found")
	ErrInvalidTaskData = errors.New("service: invalid task data provided")
	ErrCreateTask      = errors.New("service: error creating task")
	ErrInvalidUser     = errors.New("service: invalid user for task")
)

type TaskServices interface {
	GetTask(id int) (*TaskResponse, error)
	CreateTask(task *TaskInput) (*TaskResponse, error)
	// UpdateTask(id int, name, email string) error
	// CompleteTask(id int) error
	// DeleteTask(id int) error
	// GetAllTasks() error
}

type TaskService struct {
	db *gorm.DB
}

type TaskInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    uint   `json:"priority"`
	AssigneeID  uint   `json:"assignee"`
	CreatedByID uint   `json:"created_by"`
	UpdatedById uint   `json:"updated_by"`
}
type TaskResponse struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    uint      `json:"priority"`
	Completed   bool      `json:"completed"`
	Assignee    string    `json:"assignee"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTaskService(db *gorm.DB) TaskServices {
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
		Completed:   task.Completed,
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
	result := s.db.Preload("Assignee").Preload("CreatedBy").Preload("UpdatedBy").First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrTaskNotFound
		}
		return nil, ErrInvalidTaskData
	}
	taskResponse := ConvertTaskToResponse(&task)
	return taskResponse, nil
}

func (s *TaskService) CreateTask(task *TaskInput) (*TaskResponse, error) {
	// Check if the user id which is creating the task and the
	// assignee are valid

	var creator, assignee models.User

	check_creator := s.db.First(&creator, task.CreatedByID)

	if check_creator.Error != nil {
		if errors.Is(check_creator.Error, gorm.ErrRecordNotFound) {
			log.Printf("Invalid creator user id: %v", check_creator.Error)
			return nil, ErrInvalidUser
		}
		return nil, ErrCreateTask
	}

	check_assignee := s.db.First(&assignee, task.AssigneeID)

	if check_assignee.Error != nil {
		if errors.Is(check_assignee.Error, gorm.ErrRecordNotFound) {
			log.Printf("Invalid assignee user id: %v", check_assignee.Error)
			return nil, ErrInvalidUser
		}
		return nil, ErrCreateTask
	}

	newTask := models.Task{
		Title:       task.Title,
		Description: task.Description,
		Priority:    task.Priority,
		Assignee:    assignee,
		CreatedBy:   creator,
		UpdatedBy:   creator,
	}

	result := s.db.Create(&newTask)

	if result.Error != nil {
		log.Printf("Error creating new task: %v", result.Error)
		return nil, ErrCreateTask
	}

	taskResponse := ConvertTaskToResponse(&newTask)
	return taskResponse, nil
}
