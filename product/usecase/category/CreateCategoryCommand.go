package category

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (i impl) CreateCategoryCommand(request dto.RequestCategory, userID int, levelUser int) (dto.JsonResponses, error) {
	switch levelUser {
	case 1, 2, 3:
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	category := domain.Category{
		ID:   request.ID,
		Name: request.Name,
	}

	current, err := i.repository.FetchCategoryByID(request.ID)

	if current != nil && err == nil {
		category.AuditTable = domain.AuditTable{UpdatedBy: userID}
		err = i.repository.UpdateCategory(request.ID, category)
	} else if err != nil {
		return command.InternalServerResponses(""), err
	} else {
		category.AuditTable = domain.AuditTable{CreatedBy: userID}
		err = i.repository.CreateCategory(category)
	}

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
