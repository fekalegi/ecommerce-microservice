package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) DeleteOrderProductCommand(orderProductID int64, userID int, levelUser int) (dto.JsonResponses, error) {
	orderProduct, err := i.repository.FetchOrderProductID(orderProductID)
	if err == nil && orderProduct == nil {
		return command.NotFoundResponses("Order Product Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	order, err := i.repository.FetchOrderByID(orderProduct.OrderID)
	if err == nil && order == nil {
		return command.NotFoundResponses("Order Product Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	switch levelUser {
	case 1, 2:
	case 3, 4:
		if order.UserID != userID {
			return command.UnauthorizedResponses("Unauthorized, User not the one order"), nil
		}
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	err = i.repository.DeleteOrderProduct(orderProductID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses("Success"), nil
}
