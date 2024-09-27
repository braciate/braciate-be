package votesRepository

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
	NewClient(tx bool) (UserVotesRepositoryItf, error)
}

func New(log *logrus.Logger, db *sqlx.DB) RepositoryItf {
	return &Repository{
		DB:  db,
		log: log,
	}
}

func (r *Repository) NewClient(tx bool) (UserVotesRepositoryItf, error) {
	var db sqlx.ExtContext

	db = r.DB

	if tx {
		var err error
		db, err = r.DB.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &UserVotesRepository{
		DB:  db,
		log: r.log,
	}, nil

}

type UserVotesRepository struct {
	DB  sqlx.ExtContext
	log *logrus.Logger
}

type UserVotesRepositoryItf interface {
	CreateUserVotes(ctx context.Context, votes entity.UserVotes) (entity.UserVotes, error)
	GetAllUserVotesByNomination(ctx context.Context, id string) ([]entity.UserVotes, error)
}
