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

		{ID: "adminDPU", Username: "adminDPU", NIM: "235150000000002", Email: "DPU@student.ub.ac.id",
			StudyProgram: "DPU", Role: entity.UserRoleDelegation, Faculty: "FILKOM", Password: os.Getenv("PASSWORDDPU")},
		{ID: "adminBEM", Username: "adminBEM", NIM: "235150000000003", Email: "BEM@student.ub.ac.id",
			StudyProgram: "BEM", Role: entity.UserRoleDelegation, Faculty: "FILKOM", Password: os.Getenv("PASSWORDBEM")},
		{ID: "adminHIMA", Username: "adminHIMA", NIM: "235150000000004", Email: "HIMA@student.ub.ac.id",
			StudyProgram: "HIMA", Role: entity.UserRoleDelegation, Faculty: "FILKOM", Password: os.Getenv("PASSWORDHIMA")},
		{ID: "adminUKM", Username: "adminUKM", NIM: "235150000000005", Email: "UKM@student.ub.ac.id",
			StudyProgram: "UKM", Role: entity.UserRoleDelegation, Faculty: "FILKOM", Password: os.Getenv("PASSWORDUKM")},
	}

	for _, v := range user {
		var hashPass string
		var err error
		switch v.Role {
		case entity.UserRoleDelegation:
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
