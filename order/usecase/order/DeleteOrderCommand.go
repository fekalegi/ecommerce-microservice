package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) DeleteOrderCommand(orderID int64, userID int, levelUser int) (dto.JsonResponses, error) {

	current, err := i.repository.FetchOrderByID(orderID)

	if err == nil && current == nil {
		return command.NotFoundResponses("Order Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	switch levelUser {
	case 1, 2:
	case 3, 4:
		if userID != current.UserID {
			return command.UnauthorizedResponses("Unauthorized, user is not the one ordered"), nil
		}
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}
	err = i.repository.DeleteOrder(orderID)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been deleted"), nil
}
