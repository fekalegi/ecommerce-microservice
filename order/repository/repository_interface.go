package repository

import (
	"ecommerce-microservice/order/domain"
	"github.com/google/uuid"
)

//go:generate mockgen -destination=../mocks/repository/mock_repository.go -package=mock_repository -source=repository_interface.go

type OrderRepository interface {
	CreateOrder(order domain.Order) (*domain.Order, error)
	CreateInitialOrder(order domain.Order, product domain.OrderProduct) error
	CreateOrderProduct(product domain.OrderProduct) error
	FindOrderProductByOrderIDAndProductID(orderID int64, productID uuid.UUID) (*domain.OrderProduct, error)
	FetchAllOrderByUserID(userID, offset, limit, status int) ([]domain.Order, int64, error)
	FetchOrderByID(id int64) (*domain.Order, error)
	UpdateOrder(id int64, request domain.Order) error
	UpdateOrderStatus(id int64, status int, history domain.OrderHistory) error
	UpdateOrderProduct(id int64, product domain.OrderProduct) error
	DeleteOrder(id int64) error
	DeleteOrderProduct(id int64) error
	FetchOrderProductID(orderProductID int64) (*domain.OrderProduct, error)
	CreateRatingSeller(seller domain.SellerRating) error
	UpdateRatingSeller(seller domain.SellerRating) error
	FindRatingSellerBySellerID(sellerID int) (domain.SellerRating, error)
	FindRatingSellerID(sellerID int) (*domain.SellerRating, error)
}
