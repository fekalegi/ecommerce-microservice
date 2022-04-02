package common

import "errors"

const (
	IsMigrate      = true
	NotFoundCode   = 404
	BadRequestCode = 400

	DefaultPageSize   = 10
	DefaultPageNumber = 1

	UserTypeInternal   = "Internal"
	UserTypeThirdParty = "Third-Party"

	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

func GetOrderStatus(code int) (status string, err error) {
	switch code {
	case 1:
		status = "Order in User Cart"
	case 2:
		status = "Order Waiting for Payment"
	case 3:
		status = "Seller Packing Items"
	case 4:
		status = "Items are being Shipped"
	case 5:
		status = "Items have arrived"
	case 6:
		status = "Order Finished"
	default:
		err = errors.New("failed to get status")
	}
	return
}

func GetCancelDetail(code int) (status string, err error) {
	switch code {
	case 1:
		status = "Order Canceled by Admin"
	case 2:
		status = "Order Canceled by User"
	case 3:
		status = "Order Canceled by Seller"
	default:
		err = errors.New("failed to get detail")
	}
	return
}
