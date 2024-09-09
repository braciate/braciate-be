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
	GetNominatonByID(ctx context.Context, id string) (nominations.NominationResponse, error)
	//Categories
	CreateCategory(ctx context.Context, request entity.Categories) (nominations.CategoryResponse, error)
	GetCategoriesByID(ctx context.Context, id string) (nominations.CategoryResponse, error)
}
