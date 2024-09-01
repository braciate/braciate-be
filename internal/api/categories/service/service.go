package categoriesService

import (
	"github.com/braciate/braciate-be/internal/api/categories"
	categoriesRepository "github.com/braciate/braciate-be/internal/api/categories/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type CategoriesService struct {
	categoriesRepository categoriesRepository.RepositoryItf
	log                  *logrus.Logger
}

type CategoriesServiceItf interface {
	CreateCategory(ctx context.Context, request categories.CreateCategoryRequest) (categories.CreateCategoryResponse, error)
}

func New(log *logrus.Logger, repo categoriesRepository.RepositoryItf) CategoriesServiceItf {
	return &CategoriesService{
		categoriesRepository: repo,
		log:                  log,
	}
}
