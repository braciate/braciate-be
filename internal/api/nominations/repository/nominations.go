package nominationsRepository

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type NominationsDB struct {
	ID         string
	Name       string
	CategoryID string
}

func (r *NominationsRepository) CreateNomination(ctx context.Context, nomination entity.Nominations) (string, error) {
	argsKV := map[string]interface{}{
		"id":            nomination.ID,
		"name":          nomination.Name,
		"categories_id": nomination.CategoryID,
	}

	query, args, err := sqlx.Named(queryCreateNomination, argsKV)
	if err != nil {
		r.log.Errorf("CreateNomination err: %v", err)
		return "", err
	}
	query = r.DB.Rebind(query)

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		r.log.Errorf("CreateNomination err: %v", err)
		return "", err
	}

	return nomination.ID, nil
}
