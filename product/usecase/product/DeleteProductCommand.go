package product

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) DeleteProductCommand(id uuid.UUID, userID int, levelUser int) (dto.JsonResponses, error) {
	switch levelUser {
	case 1, 2, 3:
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}
	current, err := i.repository.FetchProductByID(id)

	if err == nil && current == nil {
		return command.NotFoundResponses("Product Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	err = i.repository.DeleteProduct(id)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been deleted"), nil
}
