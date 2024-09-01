package categoriesRepository

const (
	queryCreateCategory = `
		INSERT INTO Categories (id, name)
		VALUES (:id, :name)
		RETURNING id`
)
