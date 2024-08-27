package authRepository

const (
	queryGetUserByEmailOrNIM = `
		SELECT id, name, email, nim, faculty, study_program, password, role
		FROM Users 
		WHERE email = :identifier OR nim = :identifier`

	queryCreateUser = `
		INSERT INTO Users (id, name, password, nim, email, faculty, study_program, role)
		VALUES (:id, :name, :password, :nim, :email, :faculty, :study_program, :role)
		RETURNING id`
)
