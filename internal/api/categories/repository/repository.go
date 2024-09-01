package categoriesRepository

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type Repository struct {
	DB  *sqlx.DB
	log *logrus.Logger
}

type RepositoryItf interface {
	NewClient(tx bool) (CategoriesRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (CategoriesRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &CategoriesRepository{
		DB:  db,
		log: r.log,
	}, nil
}

type CategoriesRepository struct {
	DB  sqlx.ExtContext
	log *logrus.Logger
}

type CategoriesRepositoryItf interface {
	CreateCategory(ctx context.Context, category entity.Categories) (string, error)
}
