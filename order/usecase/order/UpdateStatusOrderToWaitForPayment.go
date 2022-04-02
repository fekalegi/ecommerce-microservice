package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
	"time"
)

func (i impl) UpdateStatusOrderToWaitForPayment(orderID int64, userID int, levelUser int) (dto.JsonResponses, error) {
	order, err := i.repository.FetchOrderByID(orderID)
	if err == nil && order == nil {
		return command.NotFoundResponses("Order not found"), nil
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

	history := domain.OrderHistory{
		OrderID:    order.ID,
		StatusFrom: order.Status,
		StatusTo:   2,
		Date:       time.Now(),
		AuditTable: domain.AuditTable{CreatedBy: userID},
	}

	err = i.repository.UpdateOrderStatus(orderID, 2, history)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses("Success"), nil

}
