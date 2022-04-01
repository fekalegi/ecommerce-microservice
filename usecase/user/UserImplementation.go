package user

import (
	"ecommerce-microservice/common/interfaces"
	"ecommerce-microservice/repository"
	"ecommerce-microservice/usecase"
)

type userImplementation struct {
	repo   repository.UserRepository
	helper interfaces.HelperInterface
}

func NewUserImplementation(repo repository.UserRepository, helper interfaces.HelperInterface) usecase.UserService {
	return &userImplementation{
		repo: repo,
		helper: helper,
	}
}
