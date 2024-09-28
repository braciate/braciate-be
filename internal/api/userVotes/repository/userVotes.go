package votesRepository

import (
	"context"
	"errors"

	uservotes "github.com/braciate/braciate-be/internal/api/userVotes"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type UserVotesDB struct {
	ID           string
	UserID       string
	NominationID string
	LkmID        string
}

func (r *UserVotesRepository) CreateUserVotes(ctx context.Context, votes entity.UserVotes) (entity.UserVotes, error) {
	var newVotes entity.UserVotes
	argsKV := map[string]interface{}{
		"id":            votes.ID,
		"user_id":       votes.UserID,
		"lkm_id":        votes.LkmID,
		"nomination_id": votes.NominationID,
	}

	query, args, err := sqlx.Named(queryCreateUserVotes, argsKV)
	if err != nil {
		r.log.Errorf("CreateUserVotes err: %v", err)
		return entity.UserVotes{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&newVotes.ID, &newVotes.UserID, &newVotes.LkmID, &newVotes.NominationID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" && pqErr.Constraint == "user_votes_user_id_fkey" || pqErr.Constraint == "user_votes_lkm_id_fkey" || pqErr.Constraint == "user_votes_nomination_id_fkey" {
				r.log.Errorf("Foreign key violation: %v", err)
				return entity.UserVotes{}, uservotes.ErrForeignKeyViolation
			}

			if pqErr.Code.Name() == "unique_violation" {
				r.log.Errorf("Unique constraint violation: %v", err)
				return entity.UserVotes{}, uservotes.ErrUniqueViolation
			}
		}
		r.log.Errorf("CreateUserVotes err: %v", err)
		return entity.UserVotes{}, err
	}

	return newVotes, nil
}
