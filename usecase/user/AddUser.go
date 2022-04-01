package user

import (
	"ecommerce-microservice/common/command"
	"ecommerce-microservice/common/dto"
	"ecommerce-microservice/domain"
)

func (u userImplementation) AddUser(request dto.RequestAddUser) (dto.JsonResponses, error) {
	user := domain.User{
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
	}
	newUser, err := u.repo.AddUser(user)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), nil
	}

	response := dto.ResponseAddUser{
		UserID: newUser.ID,
	}
	return dto.JsonResponses{
		Status: "success",
		Data:   response,
		Code:   201,
	}, nil
}
