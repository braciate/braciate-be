package nominationsRepository

const (
	queryCreateNomination = `
	INSERT INTO Nominations (id, name, categories_id)
	VALUES (:id, :name, :categories_id)
	RETURNING id`
)
