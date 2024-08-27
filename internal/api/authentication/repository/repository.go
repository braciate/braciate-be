package authRepository

import (
	"github.com/braciate/braciate-be/internal/api/authentication"
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
	NewClient(tx bool) (AuthRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (AuthRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &AuthRepository{
		q:   db,
		log: r.log,
	}, nil
}

type AuthRepository struct {
	q   sqlx.ExtContext
	log *logrus.Logger
}

type AuthRepositoryItf interface {
	Commit() error
	Rollback() error

	GetUserByEmailOrNIM(ctx context.Context, identifier string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (string, error)
}

func (r *AuthRepository) Commit() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}

	return authentication.ErrCommitTransaction
}

func (r *AuthRepository) Rollback() error {
	if tx, ok := r.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}

	return authentication.ErrRollbackTransaction
}
