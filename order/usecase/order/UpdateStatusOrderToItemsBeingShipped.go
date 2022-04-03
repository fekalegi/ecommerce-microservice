package order

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
	"time"
)

func (i impl) UpdateStatusOrderToItemsBeingShipped(orderID int64, userID int, levelUser int) (dto.JsonResponses, error) {
	order, err := i.repository.FetchOrderByID(orderID)
	if err == nil && order == nil {
		return command.NotFoundResponses("Order not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}
	switch order.Status {
	case 1, 2, 4, 5, 6:
		return command.BadRequestResponses("Status order is not valid, status should be 4"), nil
	case 3:
	default:
		return command.BadRequestResponses("Status order invalid"), nil
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
		StatusTo:   4,
		Date:       time.Now(),
		AuditTable: domain.AuditTable{CreatedBy: userID},
	}

	err = i.repository.UpdateOrderStatus(orderID, 4, history)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses("Success"), nil

}
