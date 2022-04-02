package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
)

func (i impl) CreateInitialOrderCommand(request dto.RequestInitialOrder, userID int, levelUser int) (dto.JsonResponses, error) {
	switch levelUser {
	case 1, 2, 3, 4:
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	order := domain.Order{
		UserID:     userID,
		SellerID:   request.SellerID,
		AuditTable: domain.AuditTable{CreatedBy: userID},
	}

	TotalPrice := request.Products.Price * request.Products.Quantity

	product := domain.OrderProduct{
		ProductID:  request.Products.ProductID,
		Quantity:   request.Products.Quantity,
		Price:      request.Products.Price,
		TotalPrice: TotalPrice,
		AuditTable: domain.AuditTable{CreatedBy: userID},
	}

	err := i.repository.CreateInitialOrder(order, product)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
