package user

import (
	"ecommerce-microservice/order/common/interfaces"
	"ecommerce-microservice/order/repository"
	"ecommerce-microservice/order/usecase"
)

type userImplementation struct {
	repo   repository.UserRepository
	helper interfaces.HelperInterface
}

func NewUserImplementation(repo repository.UserRepository, helper interfaces.HelperInterface) usecase.UserService {
	return &userImplementation{
		repo:   repo,
		helper: helper,
	}
}
