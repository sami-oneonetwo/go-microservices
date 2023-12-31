package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/sami-oneonetwo/go-microservices/publisher/kafka"
	"github.com/sami-oneonetwo/go-microservices/publisher/messages"
)

func main() {

	// Init router
	r := chi.NewRouter()

	// Use the cors middleware to enable CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Update this with the allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)

	publisher, err := kafka.SetupKafkaPublisher()
	if err != nil {
		panic(err)
	}

	// Setup API routes
	r.Post("/message", messages.PublishMessage(publisher))

	http.ListenAndServe(":8888", r)
}
