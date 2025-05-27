package models

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string
	Description string
	Priority    uint
	AssigneeID  uint // Foreign key for assigned User
	Assignee    User `gorm:"foreignKey:AssigneeID"`
	CreatedByID uint // Foreign key for User who created the task
	CreatedBy   User `gorm:"foreignKey:CreatedByID"`
	UpdatedById uint // Foreign key for User who updated the task
	UpdatedBy   User `gorm:"foreignKey:UpdatedById"`
}
