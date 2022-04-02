package product

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	"github.com/google/uuid"
)

func (i impl) UpdateProductCommand(id uuid.UUID, request dto.RequestUpdateProduct, userID int, levelUser int) (dto.JsonResponses, error) {
	switch levelUser {
	case 1, 2, 3:
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}
	current, err := i.repository.FetchProductByID(id)

	if current == nil && err == nil {
		return command.NotFoundResponses("Product Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses(""), err
	}

	product := domain.Product{
		ID:          id,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryID:  request.CategoryID,
		UserID:      userID,
		AuditTable:  domain.AuditTable{UpdatedBy: userID},
	}

	err = i.repository.UpdateProduct(id, product)

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Data has been updated"), nil
}
