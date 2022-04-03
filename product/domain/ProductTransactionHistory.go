package domain

import "github.com/google/uuid"

type ProductTransactionHistory struct {
	OrderID    int64     `gorm:"primaryKey;autoIncrement:false" json:"order_id"`
	SellerID   int       `json:"seller_id"`
	ProductID  uuid.UUID `gorm:"primaryKey" json:"product_id"`
	Quantity   float64   `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
}
