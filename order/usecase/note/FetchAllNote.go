package note

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
)

func (n noteImplementation) FetchAllNote() (dto.JsonResponses, error) {
	notes, err := n.repo.FetchAllNote()
	if err != nil {
		return command.InternalServerResponses(""), err
	}
	var responseNotes []dto.ResponseBaseNote
	for _, v := range notes {
		tempNote := dto.ResponseBaseNote{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			Creator:     v.User.FullName,
		}
		responseNotes = append(responseNotes, tempNote)
	}
	response := dto.ResponseFetchAllNote{Notes: responseNotes}
	return command.SuccessResponses(response), nil
}
