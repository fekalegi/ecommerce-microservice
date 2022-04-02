package dto

import "github.com/google/uuid"

type RequestWishlist struct {
	ID        uuid.UUID `validate:"required" json:"id"`
	ProductID uuid.UUID `validate:"required" json:"product_id"`
}
