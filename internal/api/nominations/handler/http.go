package nominationsHandler

import (
	nominationsService "github.com/braciate/braciate-be/internal/api/nominations/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type NominationHandler struct {
	nominationsService nominationsService.NominationsServiceItf
	log                *logrus.Logger
	validate           *validator.Validate
}

func New(log *logrus.Logger, nominationsService nominationsService.NominationsServiceItf, validate *validator.Validate) *NominationHandler {
	return &NominationHandler{
		nominationsService: nominationsService,
		log:                log,
		validate:           validate,
	}
}

func (h *NominationHandler) Start(srv fiber.Router) {
	nominations := srv.Group("/nominations")
	nominations.Post("/create", h.CreateNominationHandler)

	categorries := srv.Group("/categories")
	categorries.Post("/create", h.CreateCategoryHandler)
}
