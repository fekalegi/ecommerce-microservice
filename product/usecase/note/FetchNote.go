package note

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (n noteImplementation) FetchNote(id int) (dto.JsonResponses, error) {
	note, err := n.repo.FetchNoteByID(id)
	if err == nil && note == nil {
		return command.NotFoundResponses("Note not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	tmpResponse := dto.ResponseBaseNote{
		ID:          note.ID,
		Title:       note.Title,
		Description: note.Description,
		Creator:     note.User.FullName,
	}

	response := dto.ResponseFetchNote{Note: tmpResponse}

	return command.SuccessResponses(response), nil
}
