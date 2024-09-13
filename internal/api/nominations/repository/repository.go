package nominationsRepository

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
	NewClient(tx bool) (NominationsRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (NominationsRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &NominationsRepository{
		DB:  db,
		log: r.log,
	}, nil
}

type NominationsRepository struct {
	DB  sqlx.ExtContext
	log *logrus.Logger
}

type NominationsRepositoryItf interface {
	//Nominations
	CreateNomination(ctx context.Context, nomination entity.Nominations) (entity.Nominations, error)
	GetAllNominationsByCategoryID(ctx context.Context, id string) ([]entity.Nominations, error)
	//Categories
	CreateCategory(ctx context.Context, category entity.Categories) (entity.Categories, error)
	GetAllCategories(ctx context.Context) ([]entity.Categories, error)
}
