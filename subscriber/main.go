package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"subscriber/connections"
	"subscriber/kafka"
)

func main() {
	fmt.Println("Setting up subscriber service")
	http.HandleFunc("/", connections.Upgrade)

	ctx := context.Background()
	go kafka.Listen(ctx, connections.SendMessage)

	http.ListenAndServe(":"+os.Getenv("SUBSCRIBER_PORT"), nil)
}
