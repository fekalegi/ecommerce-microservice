package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
)

func (u userImplementation) FetchUser(id int) (dto.JsonResponses, error) {
	user, err := u.repo.FindUserByID(id)
	if err == nil && user == nil {
		return command.NotFoundResponses("user not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses(user), nil
}
