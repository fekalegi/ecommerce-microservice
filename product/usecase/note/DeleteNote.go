package note

import (
	"ecommerce-microservice/order/common/command"
	"ecommerce-microservice/order/common/dto"
)

func (n noteImplementation) DeleteNote(id int, userID int) (dto.JsonResponses, error) {
	note, err := n.repo.FetchNoteByID(id)
	if err == nil && note == nil {
		return command.NotFoundResponses("Note not found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	if note.Creator != userID {
		return command.BadRequestResponses("Only creator who can update the note"), nil
	}
	err = n.repo.DeleteNote(id)
	if err != nil {
		return command.InternalServerResponses(""), err
	}
	return command.SuccessResponses("Success"), nil
}
