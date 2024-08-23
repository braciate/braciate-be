package config

import (
	"fmt"
	"github.com/braciate/braciate-be/database/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/jmoiron/sqlx"
)

type Server struct {
	engine   *fiber.App
	db       *sqlx.DB
	log      *logrus.Logger
	handlers []handler
}

type handler interface {
	Start(srv fiber.Router)
}

func NewServer(fiberApp *fiber.App, log *logrus.Logger) (*Server, error) {
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		return nil, err
	}

	bootstrap := &Server{
		engine: fiberApp,
		db:     db,
		log:    log,
	}

	return bootstrap, nil
}

func (s *Server) RegisterHandler() {
	s.checkHealth()
	s.handlers = append(s.handlers)
}

func (s *Server) Run() error {
	s.engine.Use(cors.New())
	router := s.engine.Group("/api/v1")

	for _, h := range s.handlers {
		h.Start(router)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	if err := s.engine.Listen(fmt.Sprintf(":%s", port)); err != nil {
		return err
	}

	return nil
}

func (s *Server) checkHealth() {
	s.engine.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🚗💨Beep Beep Your Server is Healthy!",
		})
	})
}
