package models

import (
	"time"

	"gorm.io/gorm"
)

type AccessLevel string

const (
	AccessLevelAdmin    AccessLevel = "admin"
	AccessLevelManager  AccessLevel = "manager"
	AccessLevelStaff    AccessLevel = "staff"
	AccessLevelReadOnly AccessLevel = "readonly"
)

type User struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Username    string         `gorm:"uniqueIndex;not null" json:"username"`
	Email       string         `gorm:"uniqueIndex;not null" json:"email"`
	Password    string         `gorm:"not null" json:"-"`
	FullName    string         `gorm:"not null" json:"full_name"`
	AccessLevel AccessLevel    `gorm:"type:varchar(20);not null;default:'staff'" json:"access_level"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateUserRequest struct {
	Username    string      `json:"username" validate:"required,min=3,max=50"`
	Email       string      `json:"email" validate:"required,email"`
	Password    string      `json:"password" validate:"required,min=6"`
	FullName    string      `json:"full_name" validate:"required"`
	AccessLevel AccessLevel `json:"access_level" validate:"required,oneof=admin manager staff readonly"`
}

type UpdateUserRequest struct {
	Username    string      `json:"username" validate:"omitempty,min=3,max=50"`
	Email       string      `json:"email" validate:"omitempty,email"`
	FullName    string      `json:"full_name" validate:"omitempty"`
	AccessLevel AccessLevel `json:"access_level" validate:"omitempty,oneof=admin manager staff readonly"`
	IsActive    *bool       `json:"is_active" validate:"omitempty"`
}
