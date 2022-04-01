package user

import (
	"ecommerce-microservice/user/common"
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
	"math"
)

func (u userImplementation) FetchAllUser(request dto.RequestPagination) (dto.JsonResponsesPagination, error) {
	if request.PageNumber == 0 {
		request.PageNumber = common.DefaultPageNumber
	}

	if request.PageSize == 0 {
		request.PageSize = common.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)

	users, totalData, err := u.repo.FindAllUser(offset, request.PageSize, request.Filter)
	if err == nil && users == nil {
		return command.NotFoundPaginationResponses("Users Not Found", []domain.User{}), err
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

	return command.SuccessPaginationResponses(users, meta), nil
}
