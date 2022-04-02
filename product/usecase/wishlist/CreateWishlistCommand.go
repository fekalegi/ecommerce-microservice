package wishlist

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (i impl) CreateWishlistCommand(request dto.RequestWishlist, userID int, level int) (dto.JsonResponses, error) {
	newWishlist := domain.Wishlist{
		ID:        request.ID,
		ProductID: request.ProductID,
		UserID:    userID,
		AuditTable: domain.AuditTable{
			CreatedBy: userID,
		},
	}

	err := i.repository.CreateWishlist(newWishlist)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}
	return command.SuccessResponses("Success"), nil
}
