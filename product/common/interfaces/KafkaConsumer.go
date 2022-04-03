package interfaces

import "github.com/Shopify/sarama"

type ConsumerGroupHandler interface {
	sarama.ConsumerGroupHandler
	WaitReady()
	Reset()
}
