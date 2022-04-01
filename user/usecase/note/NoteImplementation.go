package note

import (
	"ecommerce-microservice/user/repository"
	"ecommerce-microservice/user/usecase"
)

type noteImplementation struct {
	repo repository.NoteRepository
}

func NewNoteImplementation(repo repository.NoteRepository) usecase.NoteService {
	return &noteImplementation{
		repo: repo,
	}
}
