package assetsRepository

import (
	"context"
	"errors"
	"github.com/braciate/braciate-be/internal/api/assets"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AssetsDB struct {
	ID           string
	UserID       string
	NominationID string
	LkmID        string
}

func (r *AssetsRepository) CreateAssets(ctx context.Context, asset entity.Assets) (entity.Assets, error) {
	var newAsset entity.Assets
	argsKV := map[string]interface{}{
		"id":            asset.ID,
		"user_id":       asset.UserID,
		"lkm_id":        asset.LkmID,
		"nomination_id": asset.NominationID,
		"url":           asset.Url,
	}

	query, args, err := sqlx.Named(queryCreateAssets, argsKV)
	if err != nil {
		r.log.Errorf("CreateAssets err: %v", err)
		return entity.Assets{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&newAsset.ID, &newAsset.UserID, &newAsset.LkmID, &newAsset.NominationID, &newAsset.Url); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code.Name() == "foreign_key_violation" && pqErr.Constraint == "assets_user_id_fkey" || pqErr.Constraint == "assets_lkm_id_fkey" || pqErr.Constraint == "assets_nomination_id_fkey" {
				r.log.Errorf("Foreign key violation: %v", err)
				return entity.Assets{}, assets.ErrForeignKeyViolation
			}

			if pqErr.Code.Name() == "unique_violation" {
				r.log.Errorf("Unique constraint violation: %v", err)
				return entity.Assets{}, assets.ErrUniqueViolation
			}
		}
		r.log.Errorf("CreateAssets err: %v", err)
		return entity.Assets{}, err
	}

	return newAsset, nil
}

func (r *AssetsRepository) GetAllAssetsByNomination(ctx context.Context, id string) ([]entity.Assets, error) {
	var allAssets []entity.Assets

	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryGetAssetsFromNominationID, argsKV)
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
		var votes entity.Assets
		if err := rows.StructScan(&votes); err != nil {
			r.log.Errorf("Error scanning row: %v", err)
			return nil, err
		}
		allAssets = append(allAssets, votes)
	}

	if err := rows.Err(); err != nil {
		r.log.Errorf("Row iteration error: %v", err)
		return nil, err
	}

	return allAssets, nil
}

func (r *AssetsRepository) DeleteAssets(ctx context.Context, id string) (entity.Assets, error) {
	var deleted entity.Assets
	argsKV := map[string]interface{}{
		"id": id,
	}

	query, args, err := sqlx.Named(queryDeleteAssets, argsKV)
	if err != nil {
		r.log.Errorf("DeleteAssets err: %v", err)
		return entity.Assets{}, err
	}
	query = r.DB.Rebind(query)

	if err := r.DB.QueryRowxContext(ctx, query, args...).Scan(&deleted.ID, &deleted.UserID, &deleted.NominationID, &deleted.LkmID, &deleted.Url); err != nil {
		r.log.Errorf("DeleteAssets err: %v", err)
		return entity.Assets{}, err
	}

	return deleted, nil
}
