package lkmsHandler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type lkmsHandler struct {
	log      *logrus.Logger
	validate *validator.Validate
}

func New(log *logrus.Logger, validate *validator.Validate) *lkmsHandler {
	return &lkmsHandler{
		log:      log,
		validate: validate,
	}
}

func (h *lkmsHandler) Start(srv fiber.Router) {

}
