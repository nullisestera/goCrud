package seeds

import (
	"crud/seed"

	"github.com/jinzhu/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{
		{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "Jane", 30)

			},
		},
		{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				return CreateUser(db, "John", 30)
			},
		},
	}
}
