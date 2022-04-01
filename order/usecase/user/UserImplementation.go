package user

import (
	"ecommerce-microservice/product/common/interfaces"
	"ecommerce-microservice/product/repository"
	"ecommerce-microservice/product/usecase"
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
