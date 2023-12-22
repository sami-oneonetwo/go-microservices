package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sami-oneonetwo/go-microservices/subscriber/connections"
	"github.com/sami-oneonetwo/go-microservices/subscriber/kafka"
	"github.com/sami-oneonetwo/go-microservices/subscriber/messages"
)

func main() {
	fmt.Println("Setting up server in 8889")
	http.HandleFunc("/", connections.Upgrade)
	ctx := context.Background()
	go kafka.Listen(ctx, messages.SendMessage)
	http.ListenAndServe(":8889", nil)
}
