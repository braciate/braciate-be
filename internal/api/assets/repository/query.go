package assetsRepository

const (
	queryCreateAssets = `
	INSERT INTO Assets (id, user_id, lkm_id, nomination_id, url, created_at, updated_at)
VALUES (:id, :user_id, :lkm_id, :nomination_id, :url, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id, user_id, lkm_id, nomination_id, url;

	`

	queryGetAssetsFromNominationID = `
	SELECT id, user_id AS UserID, nomination_id AS NominationID, lkm_id AS LkmID, url
	FROM Assets
	WHERE nomination_id = :id`

	queryDeleteAssets = `
	DELETE FROM Assets
	WHERE id = :id
	RETURNING id, user_id, nomination_id, lkm_id, url`
)
