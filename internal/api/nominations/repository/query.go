package nominationsRepository

const (
	queryCreateNomination = `
		INSERT INTO Nominations (id, name, category_id)
		VALUES (:id, :name, :category_id)
		RETURNING id, name, category_id`

	queryGetNominationByID = `
		SELECT id, name, category_id FROM Nominations WHERE id = :id`

	queryUpdateNomination = `
		UPDATE Nominations SET name = :name WHERE id = :id RETURNING id, name, category_id`

	queryCreateCategory = `
		INSERT INTO Categories (id, name)
		VALUES (:id, :name)
		RETURNING id, name`
)
