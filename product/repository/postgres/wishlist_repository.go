package postgres

import (
	"ecommerce-microservice/product/domain"
	"ecommerce-microservice/product/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type wishlistRepo struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) repository.WishlistRepository {
	return wishlistRepo{db}
}

func (w wishlistRepo) CreateWishlist(wishlist domain.Wishlist) error {
	err := w.db.Create(&wishlist).Error
	if err != nil {
		return nil
	}
	return err
}

func (w wishlistRepo) FetchAllWishlistByUserID(userID int, offset int, limit int, filter string) ([]domain.Wishlist, int64, error) {
	var wishlist []domain.Wishlist
	var count int64

	err := w.db.Preload("Product").
		Where("user_id = ?", userID).Offset(offset).Limit(limit).Find(&wishlist).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []domain.Wishlist{}, 0, err
	}

	return wishlist, count, nil
}

func (w wishlistRepo) FetchWishlistByID(id uuid.UUID) (*domain.Wishlist, error) {
	var wishlist domain.Wishlist
	err := w.db.Preload("Product").First(&wishlist, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &wishlist, nil
}

func (w wishlistRepo) HardDeleteWishlist(id uuid.UUID) error {
	err := w.db.Unscoped().Delete(&domain.Wishlist{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
