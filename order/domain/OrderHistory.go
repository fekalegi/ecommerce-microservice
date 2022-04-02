package domain

import "time"

type OrderHistory struct {
	ID         int64     `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	OrderID    int64     `json:"order_id"`
	StatusFrom int       `json:"status_from"`
	StatusTo   int       `json:"status_to"`
	Date       time.Time `json:"date"`

	AuditTable AuditTable `gorm:"embedded" json:"-"`
}
