package order

import (
	"ecommerce-microservice/order/repository"
	"ecommerce-microservice/order/usecase"
)

type impl struct {
	repository repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) usecase.OrderService {
	return &impl{
		repository: repo,
	}
}
