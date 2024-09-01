package nominationsService

import (
	"github.com/braciate/braciate-be/internal/api/nominations"
	nominationsRepository "github.com/braciate/braciate-be/internal/api/nominations/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type NominationsService struct {
	nominationsRepository nominationsRepository.RepositoryItf
	log                   *logrus.Logger
}

type NominationsServiceItf interface {
	CreateNomination(ctx context.Context, request nominations.CreateNominationRequest) (nominations.CreateNominationResponse, error)
}

func New(log *logrus.Logger, repo nominationsRepository.RepositoryItf) NominationsServiceItf {
	return &NominationsService{
		nominationsRepository: repo,
		log:                   log,
	}
}
