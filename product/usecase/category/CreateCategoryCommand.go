package category

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (i impl) CreateCategoryCommand(request dto.RequestCategory, userID int, levelUser int) (dto.JsonResponses, error) {
	unit := domain.Category{
		ID:   request.ID,
		Name: request.Name,
	}

	current, err := i.repository.FetchCategoryByID(request.ID)

	if current != nil && err == nil {
		err = i.repository.UpdateCategory(request.ID, unit)
	} else if err != nil {
		return command.InternalServerResponses(""), err
	} else {
		err = i.repository.CreateCategory(unit)
	}

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
