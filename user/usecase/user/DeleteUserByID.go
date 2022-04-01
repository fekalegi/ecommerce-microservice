package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
)

func (u userImplementation) DeleteUser(id int) (dto.JsonResponses, error) {
	current, err := u.repo.FindUserByID(id)

	if err == nil && current == nil {
		return command.NotFoundResponses("User Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	err = u.repo.DeleteUser(id)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been deleted"), nil
}
