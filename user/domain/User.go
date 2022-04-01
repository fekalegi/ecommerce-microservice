package domain

import "github.com/google/uuid"

type User struct {
	ID       int       `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	Username string    `gorm:"uniqueIndex" json:"name"`
	Password string    `json:"password"`
	FullName string    `json:"full_name"`
	AuthUUID uuid.UUID `json:"auth_uuid"`
}
