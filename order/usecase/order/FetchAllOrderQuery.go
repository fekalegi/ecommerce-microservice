package order

import (
	"ecommerce-microservice/order/common"
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
	"math"
)

func (i impl) FetchAllOrderQuery(request dto.RequestFetchOrderByStatus, userID int, levelUser int) (dto.JsonResponsesPagination, error) {
	switch levelUser {
	case 1, 2, 3, 4:
	default:
		return command.UnauthorizedPaginationResponses("Unauthorized"), nil
	}

	if request.PageNumber == 0 {
		request.PageNumber = common.DefaultPageNumber
	}

	if request.PageSize == 0 {
		request.PageSize = common.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)
	order, totalData, err := i.repository.FetchAllOrderByUserID(userID, offset, request.PageSize, request.Status)
	if err == nil && order == nil {
		return command.NotFoundPaginationResponses("Order Not Found", []domain.Order{}), err
	} else if err != nil {
		return command.InternalServerPaginationResponses("Internal Server Error"), err
	}
	c := float64(totalData) / float64(request.PageSize)
	totalPages := int(math.Ceil(c))
	if totalPages == 0 {
		totalPages = 1
	}

	conditionNextPage := true
	if request.PageNumber >= totalPages {
		conditionNextPage = false
	}

	meta := dto.Meta{
		TotalPages:      totalPages,
		PageSize:        request.PageSize,
		TotalItem:       int(totalData),
		HasNextPage:     conditionNextPage,
		HasPreviousPage: !(request.PageNumber <= 1),
	}
	return command.SuccessPaginationResponses(order, meta), nil
}
