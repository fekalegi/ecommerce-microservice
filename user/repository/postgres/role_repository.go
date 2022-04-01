package postgres

import (
	"ecommerce-microservice/user/domain"
	"ecommerce-microservice/user/repository"
	"errors"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return roleRepository{
		db,
	}
}

func (n roleRepository) CreateRole(role domain.Role) error {
	err := n.db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (n roleRepository) FetchRoleByID(id int) (*domain.Role, error) {
	var role domain.Role
	err := n.db.First(&role, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &role, nil
}

func (n roleRepository) FetchAllRole() ([]domain.Role, error) {
	var roles []domain.Role
	err := n.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (n roleRepository) UpdateRole(role domain.Role) error {
	err := n.db.Where("id = ?", role.ID).Updates(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (n roleRepository) DeleteRole(id int) error {
	err := n.db.Delete(&domain.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
