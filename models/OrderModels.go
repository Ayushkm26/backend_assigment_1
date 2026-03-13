package models

import "time"

type Order struct {
	ID     uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID uint `gorm:"not null;index" json:"user_id"`

	Status string `gorm:"type:enum('pending','processing','completed');default:'pending'" json:"status"`

	TotalAmount float64 `gorm:"not null" json:"total_amount"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint    `gorm:"not null;index" json:"order_id"`
	ProductID uint    `gorm:"not null;index" json:"product_id"`
	Qty       int     `gorm:"not null" json:"qty"`
	Price     float64 `gorm:"not null" json:"price"`
}

type ReqOrder struct {
	UserID      uint    `json:"user_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}
