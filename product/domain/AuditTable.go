package domain

import (
	"gorm.io/gorm"
	"time"
)

// AuditTable : embedded for entities with created, updated, delete at and by
type AuditTable struct {
	CreatedBy int            `json:"created_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedBy int            `json:"updated_by"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedBy int            `json:"deleted_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
