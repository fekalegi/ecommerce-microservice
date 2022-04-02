package product

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (i impl) CreateProductCommand(request dto.RequestProduct, userID int, levelUser int) (dto.JsonResponses, error) {
	switch levelUser {
	case 1, 2, 3:
	default:
		return command.UnauthorizedResponses("Unauthorized"), nil
	}

	product := domain.Product{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		CategoryID:  request.CategoryID,
		UserID:      userID,
	}

	current, err := i.repository.FetchProductByID(request.ID)

	if current != nil && err == nil {
		product.AuditTable = domain.AuditTable{UpdatedBy: userID}
		err = i.repository.UpdateProduct(request.ID, product)
	} else if err != nil {
		return command.InternalServerResponses(""), err
	} else {
		product.AuditTable = domain.AuditTable{CreatedBy: userID}
		err = i.repository.CreateProduct(product)
	}

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
