package product

import (
	"ecommerce-microservice/product/repository"
	"ecommerce-microservice/product/usecase"
)

type impl struct {
	repository repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) usecase.ProductService {
	return &impl{
		repository: repo,
	}
}
