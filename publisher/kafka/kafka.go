package kafka

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v2/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

var topic = "test-topic"
var brokerAddress = "kafka:9092"

// This function sets up a Kafka publisher and returns it
func SetupKafkaPublisher() (message.Publisher, error) {

	// Init publisher and configure with kafka backend
	publisher, err := kafka.NewPublisher(
		kafka.PublisherConfig{
			Brokers:   []string{brokerAddress},
			Marshaler: kafka.DefaultMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
