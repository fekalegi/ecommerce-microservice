package wishlist

import (
	"ecommerce-microservice/product/common"
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"math"
)

func (i impl) FetchAllWishlistByUserID(request dto.RequestPagination, userID int) (dto.JsonResponsesPagination, error) {
	if request.PageNumber == 0 {
		request.PageNumber = common.DefaultPageNumber
	}

	if request.PageSize == 0 {
		request.PageSize = common.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)

	wishlists, totalData, err := i.repository.FetchAllWishlistByUserID(userID, offset, request.PageSize, request.Filter)
	if err != nil {
		return command.InternalServerPaginationResponses("Internal server error"), err
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

	return command.SuccessPaginationResponses(wishlists, meta), nil
}
