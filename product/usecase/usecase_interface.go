package usecase

import (
	"ecommerce-microservice/order/common/dto"
	"github.com/google/uuid"
)

type UserService interface {
	CheckLogin(username string, password string) (dto.JsonResponses, error)
	AddUser(request dto.RequestAddUser) (dto.JsonResponses, error)
	RefreshToken(userID int, authID uuid.UUID) (dto.JsonResponses, error)
	DeleteAuth(userID int, authID uuid.UUID) (dto.JsonResponses, error)
}

type NoteService interface {
	CreateNote(request dto.RequestNote, userID int) (dto.JsonResponses, error)
	FetchNote(id int) (dto.JsonResponses, error)
	FetchAllNote() (dto.JsonResponses, error)
	UpdateNote(id int, userID int, request dto.RequestNote) (dto.JsonResponses, error)
	DeleteNote(id int, userID int) (dto.JsonResponses, error)
}
