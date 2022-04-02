package domain

type Order struct {
	ID                 int64   `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	UserID             int     `json:"user_id"`
	SellerID           int     `json:"seller_id"`
	TotalOrder         float64 `json:"total_order"`
	Status             int     `json:"status"`
	IsCanceled         bool    `json:"is_canceled"`
	CancellationStatus int     `json:"cancellation_status"`
	OrderRating        int     `json:"order_rating"`

	Products   []OrderProduct `gorm:"foreignKey:OrderID" json:"products"`
	AuditTable AuditTable     `gorm:"embedded" json:"-"`
}
