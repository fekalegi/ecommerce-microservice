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
}
