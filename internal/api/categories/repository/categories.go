package categoriesRepository

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type CategoriesDB struct {
	ID   string
	Name string
}

func (r *CategoriesRepository) CreateCategory(ctx context.Context, category entity.Categories) (string, error) {
	argsKV := map[string]interface{}{
		"id":   category.ID,
		"name": category.Name,
	}

	query, args, err := sqlx.Named(queryCreateCategory, argsKV)
	if err != nil {
		r.log.Errorf("CreateCategory err: %v", err)
		return "", err
	}
	query = r.DB.Rebind(query)

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		r.log.Errorf("CreateCategory err: %v", err)
		return "", err
	}

	return category.ID, nil
}
