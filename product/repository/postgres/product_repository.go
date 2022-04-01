package postgres

import (
	"ecommerce-microservice/product/domain"
	"ecommerce-microservice/product/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return productRepo{db}
}

func (u productRepo) CreateProduct(product domain.Product) error {
	err := u.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (u productRepo) FetchAllProduct(offset int, limit int, filter string) ([]domain.Product, int64, error) {
	var product []domain.Product
	var count int64

	err := u.db.
		Where("name ILIKE ? ", "%"+filter+"%").Offset(offset).Limit(limit).Find(&product).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []domain.Product{}, 0, err
	}

	return product, count, nil
}

func (u productRepo) FetchProductByID(id uuid.UUID) (*domain.Product, error) {
	var product domain.Product
	err := u.db.First(&product, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func (u productRepo) UpdateProduct(id uuid.UUID, request domain.Product) error {
	err := u.db.Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (u productRepo) DeleteProduct(id uuid.UUID) error {
	err := u.db.Delete(&domain.Product{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
