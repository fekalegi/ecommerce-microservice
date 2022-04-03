package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) AddRatingSeller(request dto.RequestRatingSeller, userID int, orderID int64) (dto.JsonResponses, error) {

	order, err := i.repository.FetchOrderByID(orderID)
	if order == nil && err == nil {
		return command.NotFoundResponses("Order not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	sellerRating, err := i.repository.FindRatingSellerBySellerID(order.SellerID)

	switch request.Rating {
	case 1:
		sellerRating.One += 1
	case 2:
		sellerRating.Two += 1
	case 3:
		sellerRating.Three += 1
	case 4:
		sellerRating.Four += 1
	case 5:
		sellerRating.Five += 1
	default:
		return command.BadRequestResponses("Rating should be between 1 to 5"), nil
	}

	if err == nil && sellerRating.SellerID == 0 {
		sellerRating.SellerID = userID
		err = i.repository.CreateRatingSeller(sellerRating)
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	} else {
		err = i.repository.UpdateRatingSeller(sellerRating)
	}

	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses("success"), nil

}
