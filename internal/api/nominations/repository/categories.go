package nominationsRepository

import (
	"errors"
	"fmt"

	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/net/context"
)

type CategoriesDB struct {
	ID   string
	Name string
}

func (r *NominationsRepository) CreateCategory(ctx context.Context, category entity.Categories) (entity.Categories, error) {
	var newCategories entity.Categories
	argsKV := map[string]interface{}{
		"id":   category.ID,
		"name": category.Name,
	}

	query, args, err := sqlx.Named(queryCreateCategory, argsKV)
	if err != nil {
		r.log.Errorf("CreateCategory err: %v", err)
		return entity.Categories{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&newCategories.ID, &newCategories.Name); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" && pqErr.Constraint == "nominations_categories_id_fkey" {
				return entity.Categories{}, fmt.Errorf("category ID does not exist")
			}

			if pqErr.Code.Name() == "unique_violation" {
				return entity.Categories{}, fmt.Errorf("duplicate entry detected")
			}
		}

		r.log.Errorf("CreateCategory err: %v", err)
		return entity.Categories{}, err
	}

	return newCategories, nil
}

func (r *NominationsRepository) GetAllCategories(ctx context.Context) ([]entity.Categories, error) {
	var allCategories []entity.Categories

	rows, err := r.DB.QueryxContext(ctx, queryGetAllCategories)
	if err != nil {
		r.log.Errorf("GetCategories err: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entity.Categories
		if err := rows.StructScan(&category); err != nil {
			r.log.Errorf("StructScan err: %v", err)
			return nil, err
		}
		allCategories = append(allCategories, category)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("rows.Err: %v", err)
		return nil, err
	}

	return allCategories, nil
}

func (r *NominationsRepository) GetCategoryByID(ctx context.Context, id string) (entity.Categories, error) {
	var getCategories entity.Categories
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(quertGetCategoryByID, argsKV)
	if err != nil {
		r.log.Errorf("GetCategory err: %v", err)
		return entity.Categories{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&getCategories.ID, &getCategories.Name); err != nil {
		r.log.Errorf("GetCategory err: %v", err)
		return entity.Categories{}, err
	}

	return getCategories, nil
}

func (r *NominationsRepository) UpdateCategory(ctx context.Context, updateCategory entity.Categories) (entity.Categories, error) {
	argsKV := map[string]interface{}{
		"id":   updateCategory.ID,
		"name": updateCategory.Name,
	}

	query, args, err := sqlx.Named(queryUpdateCategory, argsKV)
	if err != nil {
		r.log.Errorf("UpdateCategory err: %v", err)
		return entity.Categories{}, err
	}
	query = r.DB.Rebind(query)

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		r.log.Errorf("UpdateCategory err: %v", err)
		return entity.Categories{}, err
	}

	return updateCategory, nil
}
