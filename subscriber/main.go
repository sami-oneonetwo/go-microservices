package main

import (
	"context"
	"fmt"
	"net/http"

	"subscriber/connections"
	"subscriber/kafka"
)

func main() {
	fmt.Println("Setting up server in 8889")
	http.HandleFunc("/", connections.Upgrade)

	ctx := context.Background()
	go kafka.Listen(ctx, connections.SendMessage)

	http.ListenAndServe(":8889", nil)
}
