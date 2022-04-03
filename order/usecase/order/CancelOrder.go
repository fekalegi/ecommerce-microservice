package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (i impl) CancelOrder(id int64, userID int, levelUser int) (dto.JsonResponses, error) {
	order, err := i.repository.FetchOrderByID(id)
	if err == nil && order == nil {
		return command.NotFoundResponses("order not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	switch levelUser {
	case 1, 2:
		err = i.repository.UpdateCancelOrder(id, 1)
	case 3:
		if order.SellerID == userID {
			err = i.repository.UpdateCancelOrder(id, 2)
		} else if order.UserID != userID {
			return command.UnauthorizedResponses("Unauthorized, User not the one order"), nil
		} else {
			err = i.repository.UpdateCancelOrder(id, 3)
		}
	case 4:
		if order.UserID != userID {
			return command.UnauthorizedResponses("Unauthorized, User not the one order"), nil
		}
		err = i.repository.UpdateCancelOrder(id, 2)
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses("Succeed"), nil
}
