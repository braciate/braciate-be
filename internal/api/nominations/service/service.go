package nominationsService

import (
	"github.com/braciate/braciate-be/internal/api/nominations"
	nominationsRepository "github.com/braciate/braciate-be/internal/api/nominations/repository"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type NominationsService struct {
	nominationsRepository nominationsRepository.RepositoryItf
	log                   *logrus.Logger
}

type NominationsServiceItf interface {
	CreateNomination(ctx context.Context, request entity.Nominations) (nominations.CreateNominationResponse, error)
	CreateCategory(ctx context.Context, request entity.Categories) (nominations.CreateCategoryResponse, error)
}

func New(log *logrus.Logger, repo nominationsRepository.RepositoryItf) NominationsServiceItf {
	return &NominationsService{
		nominationsRepository: repo,
		log:                   log,
	}
}
