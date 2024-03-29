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

	TotalPrice := request.Products.Price * request.Products.Quantity

	order := domain.Order{
		UserID:     userID,
		SellerID:   request.SellerID,
		TotalOrder: TotalPrice,
		Status:     1,
		AuditTable: domain.AuditTable{CreatedBy: userID},
	}

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
