package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) FetchOrderQuery(id int64, userID int, levelUser int) (dto.JsonResponses, error) {

	order, err := i.repository.FetchOrderByID(id)
	if err == nil && order == nil {
		return command.NotFoundResponses("Order Not Found"), nil
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

	return command.SuccessResponses(order), nil
}
