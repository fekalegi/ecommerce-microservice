package role

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
)

func (n roleImplementation) DeleteRole(id int, userID int) (dto.JsonResponses, error) {
	role, err := n.repo.FetchRoleByID(id)
	if err == nil && role == nil {
		return command.NotFoundResponses("Role not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	err = n.repo.DeleteRole(id)
	if err != nil {
		return command.InternalServerResponses(""), err
	}
	return command.SuccessResponses("Success"), nil
}
