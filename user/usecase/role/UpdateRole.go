package role

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
)

func (n roleImplementation) UpdateRole(id int, userID int, request dto.RequestRole) (dto.JsonResponses, error) {
	role, err := n.repo.FetchRoleByID(id)
	if err == nil && role == nil {
		return command.NotFoundResponses("Role not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	newRole := domain.Role{
		ID:    id,
		Name:  request.Name,
		Level: request.Level,
	}

	err = n.repo.UpdateRole(newRole)
	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
