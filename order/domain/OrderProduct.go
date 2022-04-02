package domain

import "github.com/google/uuid"

type OrderProduct struct {
	ID         int64     `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	OrderID    int64     `json:"order_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   float64   `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`

	AuditTable AuditTable `gorm:"embedded" json:"-"`
}
