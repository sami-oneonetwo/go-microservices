package messages

import (
	"net/http"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

// This function publishes a message to the Kafka topic
func PublishMessage(publisher message.Publisher) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		time := time.Now().Format(time.RFC850)
		messageBody := []byte(time)

		msg := message.NewMessage(watermill.NewUUID(), messageBody)

		if err := publisher.Publish("test-topic", msg); err != nil {
			http.Error(w, "Error publishing message to Kafka", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("Message published successfully"))
	}
}
