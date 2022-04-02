package domain

import "github.com/google/uuid"

type Wishlist struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	UserID    int       `json:"user_id"`

	AuditTable AuditTable `gorm:"embedded" json:"-"`
	Product    Product    `gorm:"foreignKey:ProductID" json:"product"`
}
