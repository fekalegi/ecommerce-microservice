package category

import (
	"ecommerce-microservice/product/common"
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	"math"
)

func (i impl) FetchAllCategoryQuery(request dto.RequestPagination) (dto.JsonResponsesPagination, error) {
	if request.PageNumber == 0 {
		request.PageNumber = common.DefaultPageNumber
	}

	if request.PageSize == 0 {
		request.PageSize = common.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)
	category, totalData, err := i.repository.FetchAllCategory(offset, request.PageSize, request.Filter)
	if err == nil && category == nil {
		return command.NotFoundPaginationResponses("Category Not Found", []domain.Category{}), err
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
	return command.SuccessPaginationResponses(category, meta), nil
}
