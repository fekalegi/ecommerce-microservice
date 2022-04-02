package repository

import (
	"ecommerce-microservice/product/domain"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/repository/mock_repository.go -package=mock_repository -source=repository_interface.go

type ProductRepository interface {
	CreateProduct(product domain.Product) error
	FetchAllProduct(offset int, limit int, filter string) ([]domain.Product, int64, error)
	FetchProductByID(id uuid.UUID) (*domain.Product, error)
	UpdateProduct(id uuid.UUID, request domain.Product) error
	DeleteProduct(id uuid.UUID) error
}

type CategoryRepository interface {
	CreateCategory(category domain.Category) error
	FetchAllCategory(offset int, limit int, filter string) ([]domain.Category, int64, error)
	FetchCategoryByID(id uuid.UUID) (*domain.Category, error)
	UpdateCategory(id uuid.UUID, request domain.Category) error
	DeleteCategory(id uuid.UUID) error
}

type WishlistRepository interface {
	CreateWishlist(wishlist domain.Wishlist) error
	FetchAllWishlistByUserID(userID int, offset int, limit int, filter string) ([]domain.Wishlist, int64, error)
	FetchWishlistByID(id uuid.UUID) (*domain.Wishlist, error)
	HardDeleteWishlist(id uuid.UUID) error
}
