package user

import (
	"ecommerce-microservice/user/common/interfaces"
	"ecommerce-microservice/user/repository"
	"ecommerce-microservice/user/usecase"
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
