package assetsRepository

import (
	"context"

	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	DB  *sqlx.DB
	log *logrus.Logger
}

type RepositoryItf interface {
	NewClient(tx bool) (AssetsRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (AssetsRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &AssetsRepository{
		DB:  db,
		log: r.log,
	}, nil

}

type AssetsRepository struct {
	DB  sqlx.ExtContext
	log *logrus.Logger
}

type AssetsRepositoryItf interface {
	CreateAssets(ctx context.Context, votes entity.Assets) (entity.Assets, error)
	GetAllAssetsByNomination(ctx context.Context, id string) ([]entity.Assets, error)
	DeleteAssets(ctx context.Context, id string) (entity.Assets, error)
}
