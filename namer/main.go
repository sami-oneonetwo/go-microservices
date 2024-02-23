package main

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/goombaio/namegenerator"
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

	// Setup API routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		seed := time.Now().UTC().UnixNano()
		nameGenerator := namegenerator.NewNameGenerator(seed)
		name := nameGenerator.Generate()

		w.Write([]byte(name))
	})

	http.ListenAndServe(":"+os.Getenv("PORT"), r)

}
