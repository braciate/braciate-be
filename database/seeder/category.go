package seeder

import (
	"fmt"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
)

func CategorySeeder(db *sqlx.DB) error {
	categories := []entity.Categories{
		{ID: "bem", Name: "Badan Eksekutif Mahasiswa"},
		{ID: "hima", Name: "Himpunan Mahasiswa"},
		{ID: "ukm", Name: "Unit Kegiatan Mahasiswa"},
		{ID: "dpm", Name: "Dewan Perwakilan Mahasiswa"},
	}

	// Insert each category into the database
	for _, category := range categories {
		_, err := db.Exec("INSERT INTO categories (id, name) VALUES ($1, $2)", category.ID, category.Name)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", category.Name, err)
		}
	}

	return nil
}
