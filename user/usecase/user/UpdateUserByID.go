package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
)

func (u userImplementation) UpdateUser(id int, request dto.RequestUser) (dto.JsonResponses, error) {
	current, err := u.repo.FindUserByID(id)

	if current == nil && err == nil {
		return command.NotFoundResponses("User Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	user := domain.User{
		ID:       id,
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
		RoleID:   request.RoleID,
	}

	err = u.repo.UpdateUser(user)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been updated"), nil
}
