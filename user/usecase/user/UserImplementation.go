package user

import (
	"ecommerce-microservice/user/repository"
	"ecommerce-microservice/user/usecase"
	"github.com/fekalegi/custom-package/authentications/middlewares/common/interfaces"
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
