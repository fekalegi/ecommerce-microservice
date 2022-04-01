package category

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) FetchCategoryQuery(id uuid.UUID) (dto.JsonResponses, error) {
	category, err := i.repository.FetchCategoryByID(id)
	if err == nil && category == nil {
		return command.NotFoundResponses("Category Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses(category), nil
}
