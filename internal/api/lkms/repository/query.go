package lkmsRepository

const (
	queryCreateLkm = `
	Insert INTO Lkms(id, name, category_id, logo_file, type)
	VALUES(:id, :name, :category_id, :logo_file, :type)
	RETURNING id, name, category_id, logo_file, type`
)
