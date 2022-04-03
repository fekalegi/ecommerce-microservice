package dto

import "github.com/google/uuid"

type KafkaProduceOrderData struct {
	OrderID    int64     `json:"order_id"`
	SellerID   int       `json:"seller_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Quantity   float64   `json:"quantity"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
}
