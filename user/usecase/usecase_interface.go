package usecase

import (
	"ecommerce-microservice/user/common/dto"
	"github.com/google/uuid"
)

type UserService interface {
	CheckLogin(username string, password string) (dto.JsonResponses, error)
	AddUser(request dto.RequestUser) (dto.JsonResponses, error)
	RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error)
	DeleteAuth(userID int, authID uuid.UUID) (dto.JsonResponses, error)
	FetchAllUser(request dto.RequestPagination) (dto.JsonResponsesPagination, error)
	FetchUser(userID int) (dto.JsonResponses, error)
	UpdateUser(userID int, request dto.RequestUser) (dto.JsonResponses, error)
	DeleteUser(userID int) (dto.JsonResponses, error)
}

type RoleService interface {
	CreateRole(request dto.RequestRole, userID int) (dto.JsonResponses, error)
	FetchRole(id int) (dto.JsonResponses, error)
	FetchAllRole() (dto.JsonResponses, error)
	UpdateRole(id int, userID int, request dto.RequestRole) (dto.JsonResponses, error)
	DeleteRole(id int, userID int) (dto.JsonResponses, error)
}
