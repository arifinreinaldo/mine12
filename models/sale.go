package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCompleted OrderStatus = "completed"
	OrderStatusCancelled OrderStatus = "cancelled"
	OrderStatusRefunded  OrderStatus = "refunded"
)

type Sale struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	OrderNumber   string         `gorm:"uniqueIndex;not null" json:"order_number"`
	CustomerID    uint           `gorm:"not null" json:"customer_id"`
	Customer      Customer       `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	OrderDate     time.Time      `gorm:"not null" json:"order_date"`
	Status        OrderStatus    `gorm:"type:varchar(20);not null;default:'pending'" json:"status"`
	TotalAmount   float64        `gorm:"not null" json:"total_amount"`
	Notes         string         `json:"notes"`
	SaleItems     []SaleItem     `gorm:"foreignKey:SaleID;constraint:OnDelete:CASCADE" json:"items,omitempty"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type SaleItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SaleID    uint      `gorm:"not null" json:"sale_id"`
	ItemID    uint      `gorm:"not null" json:"item_id"`
	Item      Item      `gorm:"foreignKey:ItemID" json:"item,omitempty"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	UnitPrice float64   `gorm:"not null" json:"unit_price"`
	Subtotal  float64   `gorm:"not null" json:"subtotal"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSaleRequest struct {
	CustomerID uint              `json:"customer_id" validate:"required"`
	OrderDate  time.Time         `json:"order_date" validate:"required"`
	Notes      string            `json:"notes"`
	Items      []CreateSaleItem  `json:"items" validate:"required,min=1,dive"`
}

type CreateSaleItem struct {
	ItemID   uint `json:"item_id" validate:"required"`
	Quantity int  `json:"quantity" validate:"required,gt=0"`
}

type UpdateSaleRequest struct {
	Status OrderStatus `json:"status" validate:"omitempty,oneof=pending completed cancelled refunded"`
	Notes  string      `json:"notes"`
}
