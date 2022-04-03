package domain

type SellerRating struct {
	SellerID int     `gorm:"primaryKey;AUTO_INCREMENT" json:"seller_id"`
	One      float64 `json:"one"`
	Two      float64 `json:"two"`
	Three    float64 `json:"three"`
	Four     float64 `json:"four"`
	Five     float64 `json:"five"`
}
