package backgrounds

import (
	"ecommerce-microservice/product/common/models"
	"ecommerce-microservice/product/domain"
	"ecommerce-microservice/product/infra/messaging"
	"ecommerce-microservice/product/usecase"
	"log"
)

func StartConsumerOrder(broker []string, topic string, channel chan models.KafkaNewOrderEventDto, channelNotif chan models.KafkaNewOrderEventDto) {
	_, err := messaging.StartMultiAsyncNewOrderEventConsumer(broker, topic, channel, channelNotif)
	log.Println("Started consuming...")
	if err != nil {
		panic(err)
	}
}

func StartConsumerOrderHandler(channel chan models.KafkaNewOrderEventDto, productService usecase.ProductService) {
	go func() {
		for message := range channel {
			history := domain.ProductTransactionHistory{
				OrderID:    message.OrderID,
				SellerID:   message.SellerID,
				ProductID:  message.ProductID,
				Quantity:   message.Quantity,
				Price:      message.Price,
				TotalPrice: message.TotalPrice,
			}
			err := productService.CreateProductTransactionHistory(history)
			if err != nil {
				panic(err)
			} else {
				log.Println("Succeed")
			}

		}
	}()
}
