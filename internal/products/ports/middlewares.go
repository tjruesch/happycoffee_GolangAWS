package ports

import (
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
	"github.com/truesch/happycoffee_GolangAWS/internal/products/logs"
)

func (s *server) SetMiddlewares() {
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	s.Router.Use(middleware.Recoverer)

	s.addCorsMiddleware()

	s.Router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
	s.Router.Use(middleware.NoCache)
}

func (s *server) addCorsMiddleware() {
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
	s.Router.Use(corsMiddleware.Handler)
}
