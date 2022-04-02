package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) FetchRatingSellerBySellerID(sellerID int) (dto.JsonResponses, error) {
	sellerRating, err := i.repository.FindRatingSellerID(sellerID)
	if err == nil && sellerRating == nil {
		return command.NotFoundResponses("Rating not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	two := sellerRating.Two * 2
	three := sellerRating.Three * 3
	four := sellerRating.Four * 4
	five := sellerRating.Five * 5

	sum := sellerRating.One + two + three + four + five
	average := sum / 5

	response := dto.ResponseRatingSeller{Average: average}
	return command.SuccessResponses(response), nil
}
