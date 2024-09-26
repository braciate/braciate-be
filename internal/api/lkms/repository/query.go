package lkmsRepository

const (
	queryCreateLkm = `
	Insert INTO Lkms(id, name, category_id, logo_file, type)
	VALUES(:id, :name, :category_id, :logo_file, :type)
	RETURNING id, name, category_id, logo_file, type`

	queryGetLkmsByCategory = `
	SELECT id, name, logo_file AS LogoFile, type, category_id AS CategoryID
	FROM Lkms
	WHERE category_id = :id AND type = :type`

	queryUpdateLkms = `
	UPDATE Lkms
	SET name = :name, category_id = :category_id, logo_file=:logo_file, type = :type
	WHERE id=:id
	RETURNING id, name, category_id, logo_file, type`

	queryGetLkmByID = `
	SELECT id, name, logo_file AS LogoFile, type, category_id AS CategoryID
	FROM Lkms
	WHERE id = :id`

	queryDeleteLKMS = `
	DELETE FROM Lkms WHERE id = :id RETURNING id, name, category_id, logo_file, type`
)
