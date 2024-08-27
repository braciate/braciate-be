package seeder

import (
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/bcrypt"
	"github.com/jmoiron/sqlx"
)

func UserSeeder(db *sqlx.DB) error {
	user := []struct {
		ID           string
		Username     string
		Password     string
		NIM          string
		Email        string
		Faculty      string
		StudyProgram string
		Role         entity.UserRole
	}{
		{ID: "admin1", Username: "Admin", NIM: "235150000000001", Email: "akbarfikri@student.ub.ac.id",
			StudyProgram: "Kedokteran", Role: entity.UserRoleAdmin, Faculty: "FILKOM"},
		{ID: "admin2", Username: "Admin", NIM: "235150000000001", Email: "nandanatyon@student.ub.ac.id",
			StudyProgram: "Kedokteran", Role: entity.UserRoleAdmin, Faculty: "FILKOM"},
		{ID: "delegation1", Username: "BEM FILKOM", NIM: "235150000000001", Email: "bemfilkom@braciate.ub.ac.id",
			StudyProgram: "Kedokteran", Role: entity.UserRoleDelegation, Faculty: "FILKOM", Password: "satuhatisatujiwafilkom"},
	}

	for _, v := range user {
		var hashPass string
		var err error
		switch v.Role {
		case entity.UserRoleAdmin:
			hashPass, err = bcrypt.HashPassword(v.Password)
			if err != nil {
				return err
			}
		default:
			hashPass, err = bcrypt.HashPassword(v.ID)
			if err != nil {
				return err
			}
		}

		_, err = db.Exec("INSERT INTO users (id, name, nim, email, faculty, study_program, role, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
			v.ID,
			v.Username,
			v.NIM,
			v.Email,
			v.Faculty,
			v.StudyProgram,
			v.Role,
			hashPass)
		if err != nil {
			return err
		}
	}

	return nil
}
