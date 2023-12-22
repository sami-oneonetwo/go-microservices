package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sami-oneonetwo/go-microservices/subscriber/connections"
	"github.com/sami-oneonetwo/go-microservices/subscriber/kafka"
	"github.com/sami-oneonetwo/go-microservices/subscriber/messages"
)

const (
	port              = ":8889"
	kafkaBrokers      = "kafka:9092"
	consumerGroupName = "test_consumer_group"
	topic             = "test-topic"
)

func main() {
	fmt.Println("Setting up server in 8889")
	http.HandleFunc("/", connections.Upgrade)

	// Create Kafka subscriber
	kafkaSubscriber, err := kafka.CreateKafkaSubscriber(kafkaBrokers, consumerGroupName)
	if err != nil {
		log.Fatalf("Error creating Kafka subscriber: %v", err)
	}

	ctx := context.Background()
	kafkaMessages, err := kafkaSubscriber.Subscribe(ctx, topic)

	go kafka.Listen(ctx, kafkaMessages, messages.SendMessage)
	http.ListenAndServe(":8889", nil)
}
