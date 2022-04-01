package note

import (
	"ecommerce-microservice/common/command"
	"ecommerce-microservice/common/dto"
	"ecommerce-microservice/domain"
)

func (n noteImplementation) UpdateNote(id int, userID int, request dto.RequestNote) (dto.JsonResponses, error) {
	note, err := n.repo.FetchNoteByID(id)
	if err == nil && note == nil {
		return command.NotFoundResponses("Note not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	if note.Creator != userID {
		return command.BadRequestResponses("Only creator who can update the note"), nil
	}

	newNote := domain.Note{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Creator:     userID,
	}

	err = n.repo.UpdateNote(newNote)
	if err != nil {
		return command.InternalServerResponses(""), err
	}

	response := dto.ResponseBaseNote{
		ID:          newNote.ID,
		Title:       newNote.Title,
		Description: newNote.Description,
		Creator:     note.User.FullName,
	}
	return command.SuccessResponses(response), nil
}
