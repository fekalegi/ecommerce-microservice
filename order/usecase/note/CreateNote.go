package note

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (n noteImplementation) CreateNote(request dto.RequestNote, userID int) (dto.JsonResponses, error) {
	note := domain.Note{
		Title:       request.Title,
		Description: request.Description,
		Creator:     userID,
	}

	newNote, err := n.repo.CreateNote(note)
	if err != nil {
		return command.InternalServerResponses("Internal Server Error"), nil
	}

	response := dto.ResponseCreateNote{ID: newNote.ID}
	return dto.JsonResponses{
		Status: "success",
		Data:   response,
		Code:   201,
	}, nil
}
