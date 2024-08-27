package authHandler

import (
	authService "github.com/braciate/braciate-be/internal/api/authentication/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	authService authService.AuthServiceItf
	validate    *validator.Validate
	log         *logrus.Logger
}

func New(log *logrus.Logger, authService authService.AuthServiceItf, validate *validator.Validate) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validate,
		log:         log,
	}
}

func (h *AuthHandler) Start(srv fiber.Router) {
	auth := srv.Group("/auth")

	auth.Post("/signin", h.handleSignin)
}
