package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"github.com/google/uuid"
	"log"
	"strconv"
)

func (u userImplementation) CheckLogin(username string, password string) (dto.JsonResponses, error) {
	user, err := u.repo.CheckLogin(username, password)
	if user == nil && err == nil {
		return command.NotFoundResponses("User not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	authID := uuid.New()
	log.Println(authID)
	err = u.repo.UpdateAuthUUID(user.ID, authID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}
	conv := strconv.Itoa(user.ID)

	token, errCreateToken := u.helper.CreateJwtTokenLogin(conv, user.Username, authID, user.Role.Level)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}

	refreshToken, errCreateToken := u.helper.CreateRefreshJwtTokenLogin(conv, authID)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}

	response := dto.ResponseLogin{
		RefreshToken: refreshToken,
		AccessToken:  token,
	}

	return dto.JsonResponses{
		Status: "success",
		Data:   response,
		Code:   201,
	}, nil
}
