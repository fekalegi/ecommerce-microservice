package repository

import (
	"ecommerce-microservice/product/domain"
	"github.com/google/uuid"
)

type UserRepository interface {
	CheckLogin(username string, password string) (*domain.User, error)
	FindUserByID(id int) (*domain.User, error)
	FindUserByIDAndAuthID(id int, authID uuid.UUID) (*domain.User, error)
	AddUser(user domain.User) (*domain.User, error)
	UpdateAuthUUID(id int, authID uuid.UUID) error
}

type NoteRepository interface {
	CreateNote(note domain.Note) (*domain.Note, error)
	FetchNoteByID(id int) (*domain.Note, error)
	FetchAllNote() ([]domain.Note, error)
	UpdateNote(note domain.Note) error
	DeleteNote(id int) error
}
