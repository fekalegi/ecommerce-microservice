package wishlist

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) FetchWishlistQuery(id uuid.UUID) (dto.JsonResponses, error) {
	wishlist, err := i.repository.FetchWishlistByID(id)
	if wishlist == nil && err == nil {
		return command.NotFoundResponses("Wishlist not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	return command.SuccessResponses(wishlist), nil

}
