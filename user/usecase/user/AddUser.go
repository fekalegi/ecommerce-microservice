package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
)

func (u userImplementation) AddUser(request dto.RequestUser) (dto.JsonResponses, error) {
	user := domain.User{
		Username: request.Username,
		Password: request.Password,
		FullName: request.FullName,
		RoleID:   request.RoleID,
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
