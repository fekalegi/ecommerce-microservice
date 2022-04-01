package usecase

import (
	"ecommerce-microservice/product/common/dto"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProductCommand(request dto.RequestProduct, userID int, level int) (dto.JsonResponses, error)
	FetchAllProductQuery(request dto.RequestPagination) (dto.JsonResponsesPagination, error)
	FetchProductQuery(id uuid.UUID) (dto.JsonResponses, error)
	UpdateProductCommand(id uuid.UUID, request dto.RequestUpdateProduct, userID int, level int) (dto.JsonResponses, error)
	DeleteProductCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error)
}

type CategoryService interface {
	CreateCategoryCommand(request dto.RequestCategory, userID int, level int) (dto.JsonResponses, error)
	FetchAllCategoryQuery(request dto.RequestPagination) (dto.JsonResponsesPagination, error)
	FetchCategoryQuery(id uuid.UUID) (dto.JsonResponses, error)
	UpdateCategoryCommand(id uuid.UUID, request dto.RequestUpdateCategory, userID int, level int) (dto.JsonResponses, error)
	DeleteCategoryCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error)
}
