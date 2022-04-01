package category

import (
	"ecommerce-microservice/product/repository"
	"ecommerce-microservice/product/usecase"
)

type impl struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) usecase.CategoryService {
	return &impl{
		repository: repo,
	}
}
