package postgres

import (
	"ecommerce-microservice/user/domain"
	"ecommerce-microservice/user/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return userRepository{
		db,
	}
}

func (u userRepository) DeleteUser(id int) error {
	err := u.db.Delete(&domain.User{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) UpdateUser(request domain.User) error {
	err := u.db.Where("id = ?", request.ID).Updates(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) FindAllUser(offset int, limit int, filter string) ([]domain.User, int64, error) {
	var user []domain.User
	var count int64

	err := u.db.
		Where("name ILIKE ? ", "%"+filter+"%").Offset(offset).Limit(limit).Find(&user).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []domain.User{}, 0, err
	}

	return user, count, nil
}

func (u userRepository) UpdateAuthUUID(id int, authID uuid.UUID) error {
	err := u.db.Model(&domain.User{}).Where("id = ?", id).Update("auth_uuid", authID).Error
	if err != nil {
		return err
	}
	return nil
}

func (u userRepository) FindUserByIDAndAuthID(id int, authID uuid.UUID) (*domain.User, error) {
	var user domain.User
	err := u.db.Where("auth_uuid = ?", authID).First(&user, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) AddUser(user domain.User) (*domain.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) FindUserByID(id int) (*domain.User, error) {
	var user domain.User
	err := u.db.Preload("Role").First(&user, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepository) CheckLogin(username string, password string) (*domain.User, error) {
	var user domain.User
	err := u.db.Preload("Role").Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
