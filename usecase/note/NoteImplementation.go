package note

import (
	"ecommerce-microservice/repository"
	"ecommerce-microservice/usecase"
)

type noteImplementation struct {
	repo repository.NoteRepository
}

func NewNoteImplementation(repo repository.NoteRepository) usecase.NoteService {
	return &noteImplementation{
		repo: repo,
	}
}
