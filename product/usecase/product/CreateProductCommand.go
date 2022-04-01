package product

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
)

func (i impl) CreateProductCommand(request dto.RequestProduct, userID int, levelUser int) (dto.JsonResponses, error) {
	unit := domain.Product{
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
		err = i.repository.UpdateProduct(request.ID, unit)
	} else if err != nil {
		return command.InternalServerResponses(""), err
	} else {
		err = i.repository.CreateProduct(unit)
	}

	if err != nil {
		return command.InternalServerResponses(""), err
	}

	return command.SuccessResponses("Success"), nil
}
