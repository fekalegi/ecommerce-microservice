package role

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
)

func (n roleImplementation) FetchAllRole() (dto.JsonResponses, error) {
	roles, err := n.repo.FetchAllRole()
	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses(roles), nil
}
