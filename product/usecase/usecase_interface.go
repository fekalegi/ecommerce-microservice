package usecase

import (
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	"github.com/google/uuid"
)

type ProductService interface {
	CreateProductCommand(request dto.RequestProduct, userID int, level int) (dto.JsonResponses, error)
	FetchAllProductQuery(request dto.RequestPagination) (dto.JsonResponsesPagination, error)
	FetchProductQuery(id uuid.UUID) (dto.JsonResponses, error)
	UpdateProductCommand(id uuid.UUID, request dto.RequestUpdateProduct, userID int, level int) (dto.JsonResponses, error)
	DeleteProductCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error)
	CreateProductTransactionHistory(request domain.ProductTransactionHistory) error
}

type CategoryService interface {
	CreateCategoryCommand(request dto.RequestCategory, userID int, level int) (dto.JsonResponses, error)
	FetchAllCategoryQuery(request dto.RequestPagination) (dto.JsonResponsesPagination, error)
	FetchCategoryQuery(id uuid.UUID) (dto.JsonResponses, error)
	UpdateCategoryCommand(id uuid.UUID, request dto.RequestUpdateCategory, userID int, level int) (dto.JsonResponses, error)
	DeleteCategoryCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error)
}

type WishlistService interface {
	CreateWishlistCommand(request dto.RequestWishlist, userID int, level int) (dto.JsonResponses, error)
	FetchAllWishlistByUserID(request dto.RequestPagination, userID int) (dto.JsonResponsesPagination, error)
	FetchWishlistQuery(id uuid.UUID) (dto.JsonResponses, error)
	DeleteWishlistCommand(id uuid.UUID, userID int, level int) (dto.JsonResponses, error)
}
