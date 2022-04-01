package exception

import (
	"ecommerce-microservice/common"
	"ecommerce-microservice/common/models"
)

// NotFoundExceptions : Http Status Handle for 404
func NotFoundExceptions() models.JSONResponses {
	return models.JSONResponses{
		Code:   common.NotFoundCode,
		Status: "Data not found!",
	}
}

// BadRequestException : Http Status Handle for 400
func BadRequestException(message interface{}) models.JSONResponses {
	return models.JSONResponses{
		Code:   common.BadRequestCode,
		Status: message,
	}
}

func AuthException(message interface{}, code int) models.JSONResponses {
	return models.JSONResponses{
		Code:   code,
		Status: message,
	}
}

func ForbiddenException() models.JSONResponses {
	return models.JSONResponses{
		Code:   403,
		Status: "Access Forbidden",
	}
}
