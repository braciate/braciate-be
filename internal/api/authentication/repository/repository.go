package authRepository

import (
	"github.com/braciate/braciate-be/internal/api/authentication"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type Repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	NewClient(tx bool) (AuthRepositoryItf, error)
}

func New(db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB: db,
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
		q: db,
	}, nil
}

type AuthRepository struct {
	q sqlx.ExtContext
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
