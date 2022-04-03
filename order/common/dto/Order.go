package dto

import "github.com/google/uuid"

type RequestInitialOrder struct {
	SellerID int `json:"seller_id"`

	Products RequestOrderProduct `json:"products"`
}

type RequestOrderProduct struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  float64   `json:"quantity"`
	Price     float64   `json:"price"`
}

type RequestFetchOrderByStatus struct {
	Status     int `query:"status"`
	PageSize   int `query:"page_size"`
	PageNumber int `query:"page_number"`
}

type RequestRatingSeller struct {
	Rating int `validate:"required,number,oneof=1 2 3 4 5" enums:"1,2,3,4,5" json:"rating"`
}

type ResponseRatingSeller struct {
	Average float64 `json:"average"`
}
