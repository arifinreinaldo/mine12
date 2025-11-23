package models

import (
	"time"

	"gorm.io/gorm"
)

type CustomerType string

const (
	CustomerTypeRetail    CustomerType = "retail"
	CustomerTypeWholesale CustomerType = "wholesale"
	CustomerTypeCorporate CustomerType = "corporate"
)

type Customer struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Email       string         `gorm:"uniqueIndex" json:"email"`
	Phone       string         `json:"phone"`
	Address     string         `json:"address"`
	City        string         `json:"city"`
	Country     string         `json:"country"`
	Type        CustomerType   `gorm:"type:varchar(20);not null;default:'retail'" json:"type"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateCustomerRequest struct {
	Name    string       `json:"name" validate:"required"`
	Email   string       `json:"email" validate:"required,email"`
	Phone   string       `json:"phone" validate:"required"`
	Address string       `json:"address"`
	City    string       `json:"city"`
	Country string       `json:"country"`
	Type    CustomerType `json:"type" validate:"required,oneof=retail wholesale corporate"`
}

type UpdateCustomerRequest struct {
	Name     string       `json:"name" validate:"omitempty"`
	Email    string       `json:"email" validate:"omitempty,email"`
	Phone    string       `json:"phone" validate:"omitempty"`
	Address  string       `json:"address"`
	City     string       `json:"city"`
	Country  string       `json:"country"`
	Type     CustomerType `json:"type" validate:"omitempty,oneof=retail wholesale corporate"`
	IsActive *bool        `json:"is_active" validate:"omitempty"`
}
