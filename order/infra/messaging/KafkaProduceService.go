package messaging

import (
	"ecommerce-microservice/order/common/interfaces"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

type KafkaProducer struct {
	p     sarama.AsyncProducer
	topic string
}

type KafkaCredentials struct {
	Topic  string
	Broker []string
}

func NewProducer(cfg KafkaCredentials) (interfaces.ProducerService, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	saramaConfig.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(cfg.Broker, saramaConfig)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		p:     producer,
		topic: cfg.Topic,
	}, nil
}

func (p *KafkaProducer) ProduceMessage(msg interface{}) error {
	marshalData, _ := json.Marshal(msg)
	log.Println(msg, p.topic)
	p.p.Input() <- &sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.ByteEncoder(marshalData),
	}

	select {
	case err := <-p.p.Errors():
		log.Println(err)
		return err
	case <-p.p.Successes():
		return nil
	}
}
