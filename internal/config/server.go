package config

import (
	"fmt"
	"os"

	"github.com/braciate/braciate-be/database/postgres"
	authHandler "github.com/braciate/braciate-be/internal/api/authentication/handler"
	authRepository "github.com/braciate/braciate-be/internal/api/authentication/repository"
	authService "github.com/braciate/braciate-be/internal/api/authentication/service"
	lkmsHandler "github.com/braciate/braciate-be/internal/api/lkms/handler"
	lkmsRepository "github.com/braciate/braciate-be/internal/api/lkms/repository"
	lkmsService "github.com/braciate/braciate-be/internal/api/lkms/service"
	nominationsHandler "github.com/braciate/braciate-be/internal/api/nominations/handler"
	nominationsRepository "github.com/braciate/braciate-be/internal/api/nominations/repository"
	nominationsService "github.com/braciate/braciate-be/internal/api/nominations/service"
	votesHandler "github.com/braciate/braciate-be/internal/api/userVotes/handler"
	votesRepository "github.com/braciate/braciate-be/internal/api/userVotes/repository"
	votesService "github.com/braciate/braciate-be/internal/api/userVotes/service"
	broneAuth "github.com/braciate/braciate-be/internal/pkg/brone_auth"
	"github.com/braciate/braciate-be/internal/pkg/supabase"
	"github.com/braciate/braciate-be/internal/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"

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
	// Library
	broneAuths := broneAuth.New()
	validates := validator.NewValidator()

	// Authentication Domain
	authRepositorys := authRepository.New(s.log, s.db)
	authServices := authService.New(s.log, authRepositorys, broneAuths)
	authHandlers := authHandler.New(s.log, authServices, validates)

	// Nominations Domain
	nominationsRepositorys := nominationsRepository.New(s.log, s.db)
	nominationsServices := nominationsService.New(s.log, nominationsRepositorys)
	nominationsHandlers := nominationsHandler.New(s.log, nominationsServices, validates)

	// LKMS Domain
	supabase := supabase.NewSupabase()
	lkmsRepositorys := lkmsRepository.New(s.log, s.db)
	lkmsService := lkmsService.New(s.log, lkmsRepositorys, supabase)
	lkmsHandler := lkmsHandler.New(s.log, lkmsService, validates)

	//UserVotes Domain
	votesRepositorys := votesRepository.New(s.log, s.db)
	votesService := votesService.New(s.log, votesRepositorys)
	votesHandler := votesHandler.New(s.log, votesService, validates)
	s.checkHealth()
	s.handlers = append(s.handlers, authHandlers, nominationsHandlers, lkmsHandler, votesHandler)
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
	s.engine.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "🚗💨Beep Beep Your Server is Healthy!",
		})
	})
}
