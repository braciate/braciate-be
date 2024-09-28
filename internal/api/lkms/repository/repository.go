package lkmsRepository

import (
	"context"

	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db  *sqlx.DB
	log *logrus.Logger
}

type RepositoryItf interface {
	NewClient(tx bool) (LkmsRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		db:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (LkmsRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.db

	if tx {
		var err error
		db, err = r.db.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &LkmsRepository{
		DB:  db,
		log: r.log,
	}, nil
}

type LkmsRepository struct {
	DB  sqlx.ExtContext
	log *logrus.Logger
}

type LkmsRepositoryItf interface {
	CreateLkms(ctx context.Context, req entity.Lkms) (entity.Lkms, error)
	GetLkmsByCategoryIDAndType(ctx context.Context, id string, lkmType string) ([]entity.Lkms, error)
	GetLkmByID(ctx context.Context, id string) (entity.Lkms, error)
	UpdateLkm(ctx context.Context, UpdateLkms entity.Lkms) (entity.Lkms, error)
	DeleteLkm(ctx context.Context, id string) (entity.Lkms, error)
}
