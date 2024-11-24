package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/vitalii-minchuk/oklahoma/views/foo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	r := chi.NewRouter()
	r.Get("/", errorHandler(handler))
	// Start the server
	listenAddr := os.Getenv("PORT")
	log.Printf("HTTP server started at %s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) error {
	err := foo.Index().Render(r.Context(), w)
	if err != nil {
		return err
	}

	return nil // or an error
}

func errorHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
