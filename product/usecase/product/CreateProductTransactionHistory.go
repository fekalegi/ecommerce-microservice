package product

import "ecommerce-microservice/product/domain"

func (i impl) CreateProductTransactionHistory(request domain.ProductTransactionHistory) error {
	err := i.repository.CreateProductTransactionHistory(request)
	if err != nil {
		return err
	}

	err = i.repository.DecreaseStockProduct(request.ProductID, request.Quantity)
	if err != nil {
		return err
	}
	return nil

}
