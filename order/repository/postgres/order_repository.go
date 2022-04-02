package postgres

import (
	"ecommerce-microservice/order/domain"
	"ecommerce-microservice/order/repository"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return orderRepo{db}
}

func (u orderRepo) FindRatingSellerID(sellerID int) (*domain.SellerRating, error) {
	var sellerRating domain.SellerRating
	err := u.db.First(&sellerRating, sellerID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &sellerRating, nil
}

func (u orderRepo) FindRatingSellerBySellerID(sellerID int) (domain.SellerRating, error) {
	var sellerRating domain.SellerRating
	err := u.db.First(&sellerRating, sellerID).Error
	if err != nil {
		return domain.SellerRating{}, err
	}
	return sellerRating, nil
}

func (u orderRepo) CreateRatingSeller(seller domain.SellerRating) error {
	err := u.db.Create(&seller).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) UpdateRatingSeller(seller domain.SellerRating) error {
	err := u.db.Updates(&seller).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) UpdateOrderStatus(id int64, status int, history domain.OrderHistory) error {
	tx := u.db.Begin()

	err := tx.Model(&domain.Order{}).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&history).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (u orderRepo) FetchOrderProductID(orderProductID int64) (*domain.OrderProduct, error) {
	var orderProduct domain.OrderProduct
	err := u.db.First(&orderProduct, orderProductID).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &orderProduct, nil
}

func (u orderRepo) DeleteOrderProduct(id int64) error {
	err := u.db.Delete(&domain.OrderProduct{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) UpdateOrderProduct(id int64, product domain.OrderProduct) error {
	err := u.db.Where("id = ? ", id).Updates(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) FindOrderProductByOrderIDAndProductID(orderID int64, productID uuid.UUID) (*domain.OrderProduct, error) {
	var product domain.OrderProduct
	err := u.db.Where("order_id = ? AND product_id = ? ", orderID, productID).First(&product).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &product, nil
}

func (u orderRepo) CreateOrderProduct(product domain.OrderProduct) error {
	err := u.db.Create(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) CreateInitialOrder(order domain.Order, product domain.OrderProduct) error {
	tx := u.db.Begin()

	err := tx.Create(&order).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	product.OrderID = order.ID
	err = tx.Create(&product).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (u orderRepo) CreateOrder(order domain.Order) (*domain.Order, error) {
	err := u.db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (u orderRepo) FetchAllOrderByUserID(userID, offset, limit, status int) ([]domain.Order, int64, error) {
	var order []domain.Order
	var count int64

	err := u.db.
		Where("user_id = ? AND status = ? ", userID, status).Offset(offset).Limit(limit).Find(&order).
		Offset(-1).Limit(-1).Count(&count).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, nil
	} else if err != nil {
		return []domain.Order{}, 0, err
	}

	return order, count, nil
}

func (u orderRepo) FetchOrderByID(id int64) (*domain.Order, error) {
	var order domain.Order
	err := u.db.First(&order, id).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &order, nil
}

func (u orderRepo) UpdateOrder(id int64, request domain.Order) error {
	err := u.db.Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return err
	}
	return nil
}

func (u orderRepo) DeleteOrder(id int64) error {
	tx := u.db.Begin()

	err := tx.Where("order_id = ? ", id).Delete(&domain.OrderProduct{}).Error
	if err != nil {
		tx.Rollback()
		return nil
	}

	err = tx.Delete(&domain.Order{}, id).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
