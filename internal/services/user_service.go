package services

import (
	"errors"
	"time"

	"github.com/migvas/go-tasks-api/internal/models"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound    = errors.New("service: user not found")
	ErrInvalidUserData = errors.New("service: invalid user data provided")
)

type UserServices interface {
	GetUser(id int) (*UserResponse, error)
	// CreateTask(title, description, dueDate, email string) error
	// UpdateTask(id int, name, email string) error
	// CompleteTask(id int) error
	// DeleteTask(id int) error
	// GetAllTasks() error
}

type UserService struct {
	db *gorm.DB
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUserService(db *gorm.DB) UserServices {
	return &UserService{db: db}
}

func ConvertUserToResponse(user *models.User) *UserResponse {
	if user == nil {
		return nil
	}

	response := &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return response
}

func (s *UserService) GetUser(id int) (*UserResponse, error) {
	if id <= 0 {
		return nil, ErrInvalidTaskData
	}
	var user models.User
	result := s.db.Where("ID = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		return nil, ErrInvalidUserData
	}
	userResponse := ConvertUserToResponse(&user)
	return userResponse, nil
}
