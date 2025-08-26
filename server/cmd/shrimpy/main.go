package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/karenzhao35/shrimpy/internal/handlers"
	"github.com/karenzhao35/shrimpy/internal/storage"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	allowed := os.Getenv("ALLOWED_ORIGIN")



	if err := storage.TestDB(); err != nil {
		fmt.Printf("Database test failed: %v\n", err)
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

	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			fmt.Printf("error writing response: %v", err)
		}
	})
	r.Mount("/", urlRoutes())
	http.ListenAndServe(":"+port, r)
}

func urlRoutes() chi.Router {
	r := chi.NewRouter()
	urlHandler := handlers.UrlHandler{}
	r.Post("/clickbait", urlHandler.CreateUrl)
	r.Get("/preview/{alias}", urlHandler.GetUrl)
	r.Get("/make-money-fast/{key}", urlHandler.Redirect)
	return r
}
