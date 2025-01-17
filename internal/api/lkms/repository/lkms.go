package lkmsRepository

import (
	"context"
	"errors"

	"github.com/braciate/braciate-be/internal/api/lkms"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type LkmsDB struct {
	ID         string
	Name       string
	CategoryID string
	LogoLink   string
	Type       int
}

func (r *LkmsRepository) CreateLkms(ctx context.Context, req entity.Lkms) (entity.Lkms, error) {
	var newLkms entity.Lkms
	argsKV := map[string]interface{}{
		"id":          req.ID,
		"name":        req.Name,
		"category_id": req.CategoryID,
		"logo_file":   req.LogoFile,
		"type":        req.Type,
	}

	query, args, err := sqlx.Named(queryCreateLkm, argsKV)
	if err != nil {
		r.log.Errorf("CreateLkm err:%v", err)
		return entity.Lkms{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&newLkms.ID, &newLkms.Name, &newLkms.CategoryID, &newLkms.LogoFile, &newLkms.Type); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" && pqErr.Constraint == "lkms_user_id_fkey" {
				r.log.Errorf("Foreign key violation: %v", err)
				return entity.Lkms{}, lkms.ErrForeignKeyViolation
			}

			if pqErr.Code.Name() == "unique_violation" {
				r.log.Errorf("Unique constarint violation: %v", err)
				return entity.Lkms{}, lkms.ErrUniqueViolation
			}
		}
		r.log.Errorf("CreateLkm err: %v", err)
		return entity.Lkms{}, err

	}

	return newLkms, nil

}

func (r *LkmsRepository) GetLkmsByCategoryIDAndType(ctx context.Context, id string, lkmType string) ([]entity.Lkms, error) {
	var allLkms []entity.Lkms

	argsKV := map[string]interface{}{
		"id":   id,
		"type": lkmType,
	}

	query, args, err := sqlx.Named(queryGetLkmsByCategory, argsKV)
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
		var lkm entity.Lkms
		if err := rows.StructScan(&lkm); err != nil {
			r.log.Errorf("Error scanning row: %v", err)
			return nil, err
		}
		allLkms = append(allLkms, lkm)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Row iteration error: %v", err)
		return nil, err
	}

	return allLkms, nil
}

func (r *LkmsRepository) UpdateLkm(ctx context.Context, UpdateLkms entity.Lkms) (entity.Lkms, error) {
	argsKV := map[string]interface{}{
		"id":          UpdateLkms.ID,
		"name":        UpdateLkms.Name,
		"category_id": UpdateLkms.CategoryID,
		"logo_file":   UpdateLkms.LogoFile,
		"type":        UpdateLkms.Type,
	}

	query, args, err := sqlx.Named(queryUpdateLkms, argsKV)
	if err != nil {
		r.log.Errorf("UpdateLkms err: %v", err)
		return entity.Lkms{}, err
	}
	query = r.DB.Rebind(query)

	if _, err := r.DB.ExecContext(ctx, query, args...); err != nil {
		r.log.Errorf("UpdateLkms err: %v", err)
		return entity.Lkms{}, err
	}

	return UpdateLkms, nil
}

func (r *LkmsRepository) GetLkmByID(ctx context.Context, id string) (entity.Lkms, error) {
	var getLkm entity.Lkms
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryGetLkmByID, argsKV)
	if err != nil {
		r.log.Errorf("GetLkm err: %v", err)
		return entity.Lkms{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&getLkm.ID, &getLkm.Name, &getLkm.LogoFile, &getLkm.Type, &getLkm.CategoryID); err != nil {
		r.log.Errorf("GetLkm err: %v", err)
		return entity.Lkms{}, err
	}

	return getLkm, nil
}

func (r *LkmsRepository) DeleteLkm(ctx context.Context, id string) (entity.Lkms, error) {
	var deleted entity.Lkms
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryDeleteLKMS, argsKV)
	if err != nil {
		r.log.Errorf("DeleteLkms err: %v", err)
		return entity.Lkms{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&deleted.ID, &deleted.Name, &deleted.CategoryID, &deleted.LogoFile, &deleted.Type); err != nil {
		r.log.Errorf("DeleteLkms err: %v", err)
		return entity.Lkms{}, err
	}

	return deleted, nil
}
