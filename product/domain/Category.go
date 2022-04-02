package domain

import "github.com/google/uuid"

type Category struct {
	ID   uuid.UUID `gorm:"primaryKey" json:"id"`
	Name string    `json:"name"`

	AuditTable AuditTable `gorm:"embedded" json:"-"`
	Products   []Product  `gorm:"foreignKey:CategoryID" json:"products"`
}
