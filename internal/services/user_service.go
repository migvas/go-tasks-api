package services

import (
	"errors"
	"log"
	"time"

	verifier "github.com/AfterShip/email-verifier"
	"github.com/migvas/go-tasks-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	ErrUserNotFound    = errors.New("service: user not found")
	ErrInvalidUserData = errors.New("service: invalid user data provided")
	ErrInvalidEmail    = errors.New("service: invalid email")
	ErrCreateUser      = errors.New("service: error creating user")
)

type UserServices interface {
	GetUser(id int) (*UserResponse, error)
	CreateUser(user *UserInput) (*UserResponse, error)
	// UpdateTask(id int, name, email string) error
	// CompleteTask(id int) error
	// DeleteTask(id int) error
	// GetAllTasks() error
}

type UserService struct {
	db *gorm.DB
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func (s *UserService) CreateUser(user *UserInput) (*UserResponse, error) {
	// Check payload

	// Verify email
	v := verifier.NewVerifier()
	ret1, err := v.Verify(user.Email)

	if err != nil {
		log.Printf("Error validating email: %v", err)
		return nil, ErrCreateUser
	} else {
		if !ret1.Syntax.Valid {
			return nil, ErrInvalidEmail
		}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, ErrCreateUser
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	result := s.db.Create(&newUser)

	if result.Error != nil {
		log.Printf("Error creating new user: %v", result.Error)
		return nil, ErrCreateUser
	}

	userResponse := ConvertUserToResponse(&newUser)
	return userResponse, nil
}
