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
