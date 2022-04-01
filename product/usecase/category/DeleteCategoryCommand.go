package category

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) DeleteCategoryCommand(id uuid.UUID, userID int, levelUser int) (dto.JsonResponses, error) {
	current, err := i.repository.FetchCategoryByID(id)

	if err == nil && current == nil {
		return command.NotFoundResponses("Category Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	err = i.repository.DeleteCategory(id)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been deleted"), nil
}
