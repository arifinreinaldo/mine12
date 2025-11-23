package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Code        string         `gorm:"uniqueIndex;not null" json:"code"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price"`
	Stock       int            `gorm:"default:0" json:"stock"`
	Unit        string         `gorm:"type:varchar(20);not null" json:"unit"`
	IsActive    bool           `gorm:"default:true" json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateItemRequest struct {
	Code        string  `json:"code" validate:"required,min=2,max=50"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
	Unit        string  `json:"unit" validate:"required"`
}

type UpdateItemRequest struct {
	Code        string   `json:"code" validate:"omitempty,min=2,max=50"`
	Name        string   `json:"name" validate:"omitempty"`
	Description string   `json:"description"`
	Price       *float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       *int     `json:"stock" validate:"omitempty,gte=0"`
	Unit        string   `json:"unit" validate:"omitempty"`
	IsActive    *bool    `json:"is_active" validate:"omitempty"`
}
