package messaging

import (
	"context"
	"ecommerce-microservice/product/common/interfaces"
	"ecommerce-microservice/product/common/models"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

type ConsumerGroup struct {
	cg sarama.ConsumerGroup
}

func decodeNewOrderEventMessage(data []byte) (models.KafkaNewOrderEventDto, error) {
	var msg models.KafkaNewOrderEventDto
	err := json.Unmarshal(data, &msg)
	if err != nil {
		return models.KafkaNewOrderEventDto{}, err
	}
	return msg, nil
}

func StartMultiAsyncNewOrderEventConsumer(broker []string, topic string, channel chan models.KafkaNewOrderEventDto,
	channelNotif chan models.KafkaNewOrderEventDto) (
	*ConsumerGroup, error) {
	var bufChan = make(chan *models.ConsumerSessionMessage, 1000)
	go func() {
		for message := range bufChan {
			if value, err := decodeNewOrderEventMessage(message.Message.Value); err == nil {
				channel <- value
				channelNotif <- value
			}
		}
	}()

	handler := NewMultiAsyncConsumerGroupHandler(&models.MultiAsyncConsumerConfig{
		BufChan: bufChan,
	})
	consumer, err := NewConsumerGroup(broker, []string{topic}, "order-service-"+fmt.Sprintf("%d", time.Now().Unix()), handler)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func NewConsumerGroup(broker []string, topics []string, group string, handler interfaces.ConsumerGroupHandler) (*ConsumerGroup, error) {
	ctx := context.Background()
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V0_10_2_0
	cfg.Consumer.Offsets.Initial = sarama.OffsetNewest
	client, err := sarama.NewConsumerGroup(broker, group, cfg)
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			err := client.Consume(ctx, topics, handler)
			if err != nil {
				if err == sarama.ErrClosedConsumerGroup {
					break
				} else {
					panic(err)
				}
			}
			if ctx.Err() != nil {
				return
			}
			handler.Reset()
		}
	}()

	handler.WaitReady()

	return &ConsumerGroup{
		cg: client,
	}, nil
}

func (c *ConsumerGroup) Close() error {
	return c.cg.Close()
}
