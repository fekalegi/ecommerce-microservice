package postgres

import (
	"ecommerce-microservice/product/domain"
	"ecommerce-microservice/product/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return categoryRepo{db}
}

func (u categoryRepo) CreateCategory(category domain.Category) error {
	err := u.db.Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func (u categoryRepo) FetchAllCategory(offset int, limit int, filter string) ([]domain.Category, int64, error) {
	var category []domain.Category
	var count int64

	err := u.db.
		Where("name ILIKE ? ", "%"+filter+"%").Offset(offset).Limit(limit).Find(&category).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []domain.Category{}, 0, err
	}

	return category, count, nil
}

func (u categoryRepo) FetchCategoryByID(id uuid.UUID) (*domain.Category, error) {
	var category domain.Category
	err := u.db.First(&category, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &category, nil
}

func (u categoryRepo) UpdateCategory(id uuid.UUID, request domain.Category) error {
	err := u.db.Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (u categoryRepo) DeleteCategory(id uuid.UUID) error {
	err := u.db.Delete(&domain.Category{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
