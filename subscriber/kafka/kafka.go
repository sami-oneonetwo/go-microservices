package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

var topic = "test-topic"
var brokerAddress = "kafka:9092"

func CreateKafkaSubscriber(brokers, consumerGroup string) (*kafka.Subscriber, error) {
	saramaSubscriberConfig := kafka.DefaultSaramaSubscriberConfig()
	saramaSubscriberConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
	saramaSubscriberConfig.Consumer.Group.Heartbeat.Interval = 3 * time.Second

	return kafka.NewSubscriber(
		kafka.SubscriberConfig{
			Brokers:               []string{brokers},
			Unmarshaler:           kafka.DefaultMarshaler{},
			OverwriteSaramaConfig: saramaSubscriberConfig,
			ConsumerGroup:         consumerGroup,
		},
		watermill.NewStdLogger(false, false),
	)
}

// Listens for messages from Kafka and sends them to a callback function
func Listen(ctx context.Context, messages <-chan *message.Message, cb func(msg string)) {

	for {
		select {
		case <-ctx.Done():
			return // Context canceled, exit the function
		case message, ok := <-messages:
			if !ok {
				return // Channel closed, exit the function
			}
			fmt.Println("Received from Kafka:", string(message.Payload))
			cb(string(message.Payload))
		}
	}
}
