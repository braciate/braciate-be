package seeder

import (
	"fmt"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/jmoiron/sqlx"
)

func NominationSeeder(db *sqlx.DB) error {
	nominations := []entity.Nominations{
		{ID: "best-favorite-dpm", Name: "Best Favorite", CategoryID: "dpm"},
		{ID: "best-aspirative-dpm", Name: "Best Aspirative", CategoryID: "dpm"},
		{ID: "best-favorite-bem", Name: "Best Favorite", CategoryID: "bem"},
		{ID: "best-productive-bem", Name: "Best Productive", CategoryID: "bem"},
		{ID: "best-collaborative-bem", Name: "Best Collaborative", CategoryID: "bem"},
		{ID: "best-favorite-hima", Name: "Best Favorite", CategoryID: "hima"},
		{ID: "best-productive-hima", Name: "Best Productive", CategoryID: "hima"},
		{ID: "best-collaborative-hima", Name: "Best Collaborative", CategoryID: "hima"},
		{ID: "best-favorite-ukm", Name: "Best Favorite", CategoryID: "ukm"},
		{ID: "best-productive-ukm", Name: "Best Productive", CategoryID: "ukm"},
		{ID: "best-aoctc-ukm", Name: "Best Achivement Of Critical Thinking & Creativity", CategoryID: "ukm"},
		{ID: "best-art-sport-ukm", Name: "Best Of Art & Sport", CategoryID: "ukm"},
	}

	// Insert each category into the database
	for _, nomination := range nominations {
		_, err := db.Exec("INSERT INTO nominations (id, name, category_id) VALUES ($1, $2, $3)", nomination.ID, nomination.Name, nomination.CategoryID)
		if err != nil {
			return fmt.Errorf("failed to insert category %s: %w", nomination.Name, err)
		}
	}

	return nil
}
