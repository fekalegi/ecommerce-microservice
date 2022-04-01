package usecase

import (
	"ecommerce-microservice/user/common/dto"
	"github.com/google/uuid"
)

type UserService interface {
	CheckLogin(username string, password string) (dto.JsonResponses, error)
	AddUser(request dto.RequestAddUser) (dto.JsonResponses, error)
	RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error)
	DeleteAuth(userID int, authID uuid.UUID) (dto.JsonResponses, error)
}
