package wishlist

import (
	"ecommerce-microservice/product/repository"
	"ecommerce-microservice/product/usecase"
)

type impl struct {
	repository repository.WishlistRepository
}

func NewWishlistRepository(repo repository.WishlistRepository) usecase.WishlistService {
	return &impl{
		repository: repo,
	}
}
