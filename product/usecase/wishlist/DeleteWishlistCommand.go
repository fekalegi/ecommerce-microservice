package wishlist

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) DeleteWishlistCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error) {

	current, err := i.repository.FetchWishlistByID(id)
	if current == nil && err == nil {
		return command.NotFoundResponses("Wishlist not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	switch level {
	case 1, 2:
	case 3, 4:
		if userID != current.UserID {
			return command.UnauthorizedResponses("Unauthorized"), nil
		}
		break
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	err = i.repository.HardDeleteWishlist(id)

	if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	return command.SuccessResponses("Success"), nil

}
