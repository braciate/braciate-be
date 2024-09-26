package lkmsHandler

import (
	lkmsService "github.com/braciate/braciate-be/internal/api/lkms/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LkmsHandler struct {
	lkmsService lkmsService.LkmsServiceItf
	log         *logrus.Logger
	validate    *validator.Validate
}

func New(log *logrus.Logger, lkmsService lkmsService.LkmsServiceItf, validate *validator.Validate) *LkmsHandler {
	return &LkmsHandler{
		lkmsService: lkmsService,
		log:         log,
		validate:    validate,
	}
}

func (h *LkmsHandler) Start(srv fiber.Router) {
	lkms := srv.Group("/lkms")
	lkms.Post("/create", h.CreateLkms)
	lkms.Get("/get/:id/:type", h.GetLkmsByCategoryIDAndType)
}
