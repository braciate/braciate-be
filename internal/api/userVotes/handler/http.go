package votesHandler

import (
	votesService "github.com/braciate/braciate-be/internal/api/userVotes/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserVotesHandler struct {
	userVotesService votesService.UserVotesServiceItf
	log              *logrus.Logger
	validate         *validator.Validate
}

func New(log *logrus.Logger, userVotesService votesService.UserVotesServiceItf, validate *validator.Validate) *UserVotesHandler {
	return &UserVotesHandler{
		userVotesService: userVotesService,
		log:              log,
		validate:         validate,
	}
}

func (h *UserVotesHandler) Start(srv fiber.Router) {
	userVotes := srv.Group("/userVotes")
	userVotes.Post("/create", h.CreateUserVotesHandler)

}
