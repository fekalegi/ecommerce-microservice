package command

import "ecommerce-microservice/order/common/dto"

func SuccessResponses(data interface{}) (result dto.JsonResponses) {
	return dto.JsonResponses{
		Status: "Success",
		Data:   data,
		Code:   200,
	}
}

func NotFoundResponses(status interface{}) (result dto.JsonResponses) {
	return dto.JsonResponses{
		Status: status,
		Data:   nil,
		Code:   404,
	}
}

func UnauthorizedResponses(status interface{}) (result dto.JsonResponses) {
	return dto.JsonResponses{
		Status: status,
		Data:   nil,
		Code:   401,
	}
}

func BadRequestResponses(data string) (result dto.JsonResponses) {
	return dto.JsonResponses{
		Status: data,
		Data:   nil,
		Code:   400,
	}
}

func InternalServerResponses(data string) (result dto.JsonResponses) {
	return dto.JsonResponses{
		Status: data,
		Data:   nil,
		Code:   500,
	}
}

func SuccessPaginationResponses(data interface{}, meta interface{}) (result dto.JsonResponsesPagination) {
	return dto.JsonResponsesPagination{
		Status: "Succeed",
		Code:   200,
		Data:   data,
		Meta:   meta,
	}
}

func InternalServerPaginationResponses(data string) (result dto.JsonResponsesPagination) {
	return dto.JsonResponsesPagination{
		Status: data,
		Code:   500,
		Data:   nil,
		Meta:   dto.Meta{},
	}
}

func NotFoundPaginationResponses(status string, data interface{}) (result dto.JsonResponsesPagination) {
	return dto.JsonResponsesPagination{
		Status: data,
		Code:   404,
		Data:   data,
		Meta:   dto.Meta{},
	}
}

func BadRequestPaginationResponses(data string) (result dto.JsonResponsesPagination) {
	return dto.JsonResponsesPagination{
		Status: data,
		Code:   400,
		Data:   nil,
		Meta:   dto.Meta{},
	}
}

func UnauthorizedPaginationResponses(data string) (result dto.JsonResponsesPagination) {
	return dto.JsonResponsesPagination{
		Status: data,
		Code:   500,
		Data:   nil,
		Meta:   dto.Meta{},
	}
}
