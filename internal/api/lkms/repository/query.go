package lkmsRepository

const (
	queryCreateLkm = `
	Insert INTO Lkms(id, name, category_id, logo_link, type)
	VALUES(:id, :name, :category_id, :logo_link, :type)
	RETURNING id, name, category_id, logo_link, type`
)
