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
	//Nominations
	CreateNomination(ctx context.Context, request entity.Nominations) (nominations.NominationResponse, error)
	GetAllNominationsByCategoryID(ctx context.Context, id string) ([]nominations.NominationResponse, error)
	UpdateNomination(ctx context.Context, req entity.Nominations, id string) (nominations.NominationResponse, error)

	//Categories
	CreateCategory(ctx context.Context, request entity.Categories) (nominations.CategoryResponse, error)
	GetAllCategories(ctx context.Context) ([]nominations.CategoryResponse, error)
	UpdateCategory(ctx context.Context, req entity.Categories, id string) (nominations.CategoryResponse, error)
}

func New(log *logrus.Logger, repo nominationsRepository.RepositoryItf) NominationsServiceItf {
	return &NominationsService{
		nominationsRepository: repo,
		log:                   log,
	}
}
