package domain

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `json:"name"`

	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}
