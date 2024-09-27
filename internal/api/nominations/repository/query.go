package nominationsRepository

const (
	queryCreateNomination = `
		INSERT INTO Nominations (id, name, category_id)
		VALUES (:id, :name, :category_id)
		RETURNING id, name, category_id`

	queryGetNominationByID = `
		SELECT id, name, category_id FROM Nominations WHERE id = :id`

	queryGetAllNominationByCategoryID = `
		SELECT id, name, category_id AS CategoryID
		FROM Nominations
		WHERE category_id = :id`

	queryUpdateNomination = `
		UPDATE Nominations
    	SET name = :name, category_id = :category_id
    	WHERE id = :id
    	RETURNING id, name, category_id`

	queryCreateCategory = `
		INSERT INTO Categories (id, name)
		VALUES (:id, :name)
		RETURNING id, name`

	queryGetCategoryByID = `
		SELECT id, name FROM Categories WHERE id = :id`

	queryUpdateCategory = `
		UPDATE Categories SET name = :name WHERE id = :id RETURNING id, name`

	queryGetAllCategories = `
		SELECT id, name
		FROM Categories`

	queryDeleteCategory = `
		DELETE FROM Categories WHERE id = :id RETURNING id, name;
		`

	queryDeleteNomination = `
		DELETE FROM Nominations WHERE id = :id RETURNING id, name, category_id;
		`
)
