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

func (r *NominationsRepository) GetNominatonByID(ctx context.Context, id string) (entity.Nominations, error) {
	var getNomination entity.Nominations
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryGetNominationByID, argsKV)
	if err != nil {
		r.log.Errorf("GetNomination err: %v", err)
		return entity.Nominations{}, err
	}

	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&getNomination.ID, &getNomination.Name, &getNomination.CategoryID); err != nil {
		r.log.Errorf("GetNomination err: %v", err)
		return entity.Nominations{}, err
	}

	return getNomination, nil

}
