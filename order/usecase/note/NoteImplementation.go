package note

import (
	"ecommerce-microservice/product/repository"
	"ecommerce-microservice/product/usecase"
)

type noteImplementation struct {
	repo repository.NoteRepository
}

func NewNoteImplementation(repo repository.NoteRepository) usecase.NoteService {
	return &noteImplementation{
		repo: repo,
	}
}
