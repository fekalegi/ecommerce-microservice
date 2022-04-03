package order

import (
	"ecommerce-microservice/order/common/interfaces"
	"ecommerce-microservice/order/repository"
	"ecommerce-microservice/order/usecase"
)

type impl struct {
	repository repository.OrderRepository
	producer   interfaces.ProducerService
}

func NewOrderService(repo repository.OrderRepository, producer interfaces.ProducerService) usecase.OrderService {
	return &impl{
		repository: repo,
		producer:   producer,
	}
}
