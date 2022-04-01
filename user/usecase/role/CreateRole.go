package role

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
)

func (n roleImplementation) CreateRole(request dto.RequestRole, roleID int) (dto.JsonResponses, error) {
	role := domain.Role{
		Name:  request.Name,
		Level: request.Level,
	}

	err := n.repo.CreateRole(role)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), nil
	}

	return command.SuccessResponses("Success"), nil
}
