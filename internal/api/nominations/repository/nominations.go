package nominationsRepository

import (
	"errors"

	"github.com/braciate/braciate-be/internal/api/nominations"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/net/context"
)

type NominationsDB struct {
	ID         string
	Name       string
	CategoryID string
}

func (r *NominationsRepository) CreateNomination(ctx context.Context, nomination entity.Nominations) (entity.Nominations, error) {
	var newNomination entity.Nominations
	argsKV := map[string]interface{}{
		"id":            nomination.ID,
		"name":          nomination.Name,
		"categories_id": nomination.CategoryID,
	}

	query, args, err := sqlx.Named(queryCreateNomination, argsKV)
	if err != nil {
		r.log.Errorf("CreateNomination err: %v", err)
		return entity.Nominations{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&newNomination.ID, &newNomination.Name, &newNomination.CategoryID); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" && pqErr.Constraint == "nominations_categories_id_fkey" {
				r.log.Errorf("Foreign key violation: %v", err)
				return entity.Nominations{}, nominations.ErrForeignKeyViolation
			}

			if pqErr.Code.Name() == "unique_violation" {
				r.log.Errorf("Unique constraint violation: %v", err)
				return entity.Nominations{}, nominations.ErrUniqueViolation
			}
		}
		r.log.Errorf("CreateNomination err: %v", err)
		return entity.Nominations{}, err
	}

	return newNomination, nil
}

func (r *NominationsRepository) GetAllNominationsByCategoryID(ctx context.Context, id string) ([]entity.Nominations, error) {
	var allNominations []entity.Nominations

	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryGetAllNominationByCategoryID, argsKV)
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
		var nomination entity.Nominations
		if err := rows.StructScan(&nomination); err != nil {
			r.log.Errorf("Error scanning row: %v", err)
			return nil, err
		}
		allNominations = append(allNominations, nomination)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Row iteration error: %v", err)
		return nil, err
	}

	return allNominations, nil
}
