package seeder

import (
	"flag"
	"github.com/braciate/braciate-be/database/postgres"
	"github.com/jmoiron/sqlx"
	"log"
	"sort"
)

type Seeder struct {
	Name     string
	Priority int
	Function func(db *sqlx.DB) error
}

var seeders = []Seeder{
	{Name: "user", Priority: 1, Function: UserSeeder},
	{Name: "category", Priority: 2, Function: CategorySeeder},
	{Name: "Lkms-Bem", Priority: 3, Function: LkmsBemSeeder},
	{Name: "Lkms-Hima", Priority: 3, Function: LkmsHimaSeeder},
	{Name: "Lkms-Dpm", Priority: 3, Function: LkmsDpmSeeder},
	{Name: "Lkms-Ukm-Kerohanian", Priority: 3, Function: LkmsUkmKerohanianSeeder},
	{Name: "Lkms-Ukm-Olahraga", Priority: 3, Function: LkmsUkmOlahragaSeeder},
	{Name: "Lkms-Ukm-Kesenian", Priority: 3, Function: LkmsUkmKesenianSeeder},
	{Name: "Lkms-Ukm-MinatKhusus", Priority: 3, Function: LkmsUkmMinatKhususSeeder},
	{Name: "Lkms-Ukm-Penalaran", Priority: 3, Function: LkmsUkmPenalaranSeeder},
	{Name: "nomination", Priority: 4, Function: NominationSeeder},
}

func Seed() {
	db, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}

	sort.Slice(seeders, func(i, j int) bool {
		return seeders[i].Priority < seeders[j].Priority
	})

	name := flag.String("name", "", "name for run explisit seeder")
	flag.Parse()

	if *name == "" {
		for _, seed := range seeders {
			if err := seed.Function(db); err != nil {
				log.Fatalf("failed seeding %s: %s\n", seed.Name, err)
			}
			log.Printf("seeding %s successfully\n", seed.Name)
		}
		return
	}

	for _, seed := range seeders {
		if seed.Name == *name {
			if err := seed.Function(db); err != nil {
				log.Fatalf("failed seeding %s: %s\n", seed.Name, err)
			}
			log.Printf("seeding %s successfully\n", seed.Name)
			return
		}
	}

	log.Fatalf("seed with name %s not found\n", *name)
}
