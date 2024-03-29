package user

import (
	"ecommerce-microservice/user/common/command"
	"ecommerce-microservice/user/common/dto"
	"github.com/google/uuid"
	"strconv"
)

func (u userImplementation) RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error) {
	user, err := u.repo.FindUserByIDAndAuthID(userID, authID)
	if user == nil && err == nil {
		return command.NotFoundResponses("User not found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal server error"), err
	}

	newAuthID := uuid.New()
	err = u.repo.UpdateAuthUUID(user.ID, newAuthID)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), nil
	}
	conv := strconv.Itoa(user.ID)
	token, errCreateToken := u.helper.CreateJwtTokenLogin(conv, user.Username, newAuthID, user.Role.Level)
	if errCreateToken != nil {
		return command.InternalServerResponses(errCreateToken.Error()), err
	}

	response := dto.ResponseRefreshToken{
		AccessToken: token,
	}
	return command.SuccessResponses(response), nil
}
