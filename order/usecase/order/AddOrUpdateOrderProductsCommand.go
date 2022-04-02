package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
)

func (i impl) AddOrUpdateOrderProductsCommand(request dto.RequestOrderProduct, orderID int64, userID int) (dto.JsonResponses, error) {
	TotalPrice := request.Price * request.Quantity

	newProduct := domain.OrderProduct{
		OrderID:    orderID,
		ProductID:  request.ProductID,
		Quantity:   request.Quantity,
		Price:      request.Price,
		TotalPrice: TotalPrice,
	}

	order, err := i.repository.FetchOrderByID(orderID)
	if order == nil && err == nil {
		return command.NotFoundResponses("Order not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	if order.UserID != userID {
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	orderProduct, err := i.repository.FindOrderProductByOrderIDAndProductID(orderID, request.ProductID)
	if err == nil && orderProduct == nil {
		err = i.repository.CreateOrderProduct(newProduct)
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	} else {
		err = i.repository.UpdateOrderProduct(orderProduct.ID, newProduct)
	}

	if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	return command.SuccessResponses("Success"), nil
}
