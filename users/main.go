package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/jackc/pgx/v5"
)

// User struct represents a user model
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"` // In a real-world scenario, hash and salt the password
}

func main() {
	fmt.Println("Setting up User Service")

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
	r.Options("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	r.Post("/register", RegisterHandler)

	port := os.Getenv("PORT")

	fmt.Printf("Server is listening on port %s...\n", port)
	http.ListenAndServe(":"+port, r)
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a User struct
	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	fmt.Println(newUser)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), "postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Perform user registration logic (you will implement this)
	// For now, just print the received user data
	fmt.Printf("Received registration request: %+v\n", newUser)

	// Send a response
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}
