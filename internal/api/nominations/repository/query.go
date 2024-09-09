package nominationsRepository

const (
	queryCreateNomination = `
		INSERT INTO Nominations (id, name, categories_id)
		VALUES (:id, :name, :categories_id)
		RETURNING id, name, categories_id`

	queryCreateCategory = `
		INSERT INTO Categories (id, name)
		VALUES (:id, :name)
		RETURNING id, name`
)
