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

func (r *UserVotesRepository) GetAllUserVotesByNomination(ctx context.Context, id string) ([]entity.UserVotes, error) {
	var allUserVotes []entity.UserVotes

	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryGetUserVoteFromNominationID, argsKV)
	if err != nil {
		r.log.Errorf("Error generating query: %v", err)
		return nil, err
	}

	query = r.DB.Rebind(query)

	rows, err := r.DB.QueryxContext(ctx, query, args...)
	if err != nil {
		r.log.Errorf("Error executing query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var votes entity.UserVotes
		if err := rows.StructScan(&votes); err != nil {
			r.log.Errorf("Error scanning row: %v", err)
			return nil, err
		}
		allUserVotes = append(allUserVotes, votes)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Row iteration error: %v", err)
		return nil, err
	}

	return allUserVotes, nil
}

func (r *UserVotesRepository) DeleteUserVotes(ctx context.Context, id string) (entity.UserVotes, error) {
	var deleted entity.UserVotes
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryDeleteUserVotes, argsKV)
	if err != nil {
		r.log.Errorf("DeleteUserVotes err: %v", err)
		return entity.UserVotes{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&deleted.ID, &deleted.UserID, &deleted.NominationID, &deleted.LkmID); err != nil {
		r.log.Errorf("DeleteUserVotes err: %v", err)
		return entity.UserVotes{}, err
	}

	return deleted, nil
}
