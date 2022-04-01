package role

import (
	"ecommerce-microservice/user/repository"
	"ecommerce-microservice/user/usecase"
)

type roleImplementation struct {
	repo repository.RoleRepository
}

func NewRoleImplementation(repo repository.RoleRepository) usecase.RoleService {
	return &roleImplementation{
		repo: repo,
	}
}
