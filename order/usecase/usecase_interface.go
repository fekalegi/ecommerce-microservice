package usecase

import (
	"ecommerce-microservice/order/common/dto"
)

type OrderService interface {
	CreateInitialOrderCommand(request dto.RequestInitialOrder, userID int, level int) (dto.JsonResponses, error)
	AddOrUpdateOrderProductsCommand(request dto.RequestOrderProduct, orderID int64, userID int) (dto.JsonResponses, error)
	DeleteOrderCommand(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	DeleteOrderProductCommand(orderProductID int64, userID int, levelUser int) (dto.JsonResponses, error)
	FetchAllOrderQuery(request dto.RequestFetchOrderByStatus, userID int, levelUser int) (dto.JsonResponsesPagination, error)
	FetchOrderQuery(id int64, userID int, levelUser int) (dto.JsonResponses, error)
	UpdateStatusOrderToWaitForPayment(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	UpdateStatusOrderToPackingItems(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	UpdateStatusOrderToItemsBeingShipped(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	UpdateStatusOrderToItemsHaveArrived(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	UpdateStatusOrderToFinished(orderID int64, userID int, levelUser int) (dto.JsonResponses, error)
	AddRatingSeller(request dto.RequestRatingSeller, userID int, sellerID int, orderID int64) (dto.JsonResponses, error)
	FetchRatingSellerBySellerID(sellerID int) (dto.JsonResponses, error)
}
