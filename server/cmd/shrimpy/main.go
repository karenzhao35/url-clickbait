package main

import (
	"log"
	"os"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"github.com/karenzhao35/shrimpy/internal/handlers"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	port, allowed := os.Getenv("PORT"), os.Getenv("ALLOWED_ORIGIN")
	if port == "" {
		port = "3000"
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Add CORS middleware
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowed)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	r.Mount("/", urlRoutes())
	http.ListenAndServe(":"+port, r)
}

func urlRoutes() chi.Router {
	r := chi.NewRouter()
	urlHandler := handlers.UrlHandler{}
	r.Post("/clickbait", urlHandler.CreateUrl)
	r.Get("/make-money-fast/{key}", urlHandler.Redirect)
	return r
}
