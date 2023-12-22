package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/sami-oneonetwo/go-microservices/frontend/fileserver"
)

func main() {
	// Init router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Setup / route and serve files from /public
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "public"))
	fileserver.Serve(r, "/", filesDir)

	http.ListenAndServe(":8080", r)
}
