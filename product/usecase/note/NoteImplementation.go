package note

import (
	"ecommerce-microservice/order/repository"
	"ecommerce-microservice/order/usecase"
)

type noteImplementation struct {
	repo repository.NoteRepository
}

func NewNoteImplementation(repo repository.NoteRepository) usecase.NoteService {
	return &noteImplementation{
		repo: repo,
	}
}
