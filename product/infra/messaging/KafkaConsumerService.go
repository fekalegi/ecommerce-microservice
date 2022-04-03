package messaging

import (
	"ecommerce-microservice/product/common/interfaces"
	"ecommerce-microservice/product/common/models"
	"github.com/Shopify/sarama"
)

type multiAsyncConsumerGroupHandler struct {
	cfg *models.MultiAsyncConsumerConfig

	ready chan bool
}

func NewMultiAsyncConsumerGroupHandler(cfg *models.MultiAsyncConsumerConfig) interfaces.ConsumerGroupHandler {
	handler := multiAsyncConsumerGroupHandler{
		ready: make(chan bool, 0),
	}

	handler.cfg = cfg

	return &handler
}

func (h *multiAsyncConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	close(h.ready)
	return nil
}

func (h *multiAsyncConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *multiAsyncConsumerGroupHandler) WaitReady() {
	<-h.ready
	return
}

func (h *multiAsyncConsumerGroupHandler) Reset() {
	h.ready = make(chan bool, 0)
	return
}

func (h *multiAsyncConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	claimMsgChan := claim.Messages()

	for message := range claimMsgChan {
		h.cfg.BufChan <- &models.ConsumerSessionMessage{
			Session: session,
			Message: message,
		}
	}

	return nil
}
