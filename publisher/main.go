package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"publisher/kafka"
	"publisher/messages"
)

func main() {

	fmt.Println("Setting up publisher service")

	// Init router
	r := chi.NewRouter()

	fmt.Println("pub 1")

	// Use the cors middleware to enable CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // Update this with the allowed origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fmt.Println("pub 2")

	r.Use(middleware.Logger)

	fmt.Println("pub 3")

	publisher, err := kafka.SetupKafkaPublisher()
	if err != nil {
		panic(err)
	}

	fmt.Println("pub 4")

	// Setup API routes
	r.Post("/", messages.PublishMessage(publisher))

	fmt.Println("pub 5")

	fmt.Println(os.Getenv("PORT"))

	http.ListenAndServe(":"+os.Getenv("PORT"), r)

	fmt.Println("pub final")
}
