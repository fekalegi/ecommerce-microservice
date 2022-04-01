package product

import (
	"ecommerce-microservice/product/common/command"
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

func (i impl) FetchProductQuery(id uuid.UUID) (dto.JsonResponses, error) {
	product, err := i.repository.FetchProductByID(id)
	if err == nil && product == nil {
		return command.NotFoundResponses("Product Not Found"), nil
	} else if err != nil {
		return command.InternalServerResponses("Internal Server Error"), err
	}

	return command.SuccessResponses(product), nil
}
