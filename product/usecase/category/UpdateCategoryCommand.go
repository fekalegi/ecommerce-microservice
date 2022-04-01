package category

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	"github.com/google/uuid"
)

func (i impl) UpdateCategoryCommand(id uuid.UUID, request dto.RequestUpdateCategory, userID int, levelUser int) (dto.JsonResponses, error) {
	current, err := i.repository.FetchCategoryByID(id)

	if current == nil && err == nil {
		return command.NotFoundResponses("Category Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	category := domain.Category{
		ID:   id,
		Name: request.Name,
	}

	err = i.repository.UpdateCategory(id, category)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been updated"), nil
}
