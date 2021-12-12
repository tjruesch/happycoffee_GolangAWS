package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/handlers"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/logs"
)

func main() {
	productsRouter := chi.NewRouter()
	setMiddlewares(productsRouter)
	addRoutes(productsRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/v1", productsRouter)

	logrus.Info("Starting HTTP server")

	http.ListenAndServe(":9090", rootRouter)
}

func addRoutes(router *chi.Mux) {
	router.Route("/products", func(r chi.Router) {
		r.Get("/", handlers.GetProducts)
		r.Post("/", handlers.AddNewProduct)
		r.Delete("/{productID}", handlers.DeleteProduct)
	})
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	router.Use(middleware.NoCache)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		return
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}
