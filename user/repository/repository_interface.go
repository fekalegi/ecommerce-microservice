package repository

import (
	"ecommerce-microservice/user/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	CheckLogin(username string, password string) (*domain.User, error)
	FindUserByID(id int) (*domain.User, error)
	FindUserByIDAndAuthID(id int, authID uuid.UUID) (*domain.User, error)
	AddUser(user domain.User) (*domain.User, error)
	UpdateAuthUUID(id int, authID uuid.UUID) error
	FindAllUser(offset int, limit int, filter string) ([]domain.User, int64, error)
	UpdateUser(request domain.User) error
	DeleteUser(id int) error
}

type RoleRepository interface {
	CreateRole(role domain.Role) error
	FetchRoleByID(id int) (*domain.Role, error)
	FetchAllRole() ([]domain.Role, error)
	UpdateRole(role domain.Role) error
	DeleteRole(id int) error
}
