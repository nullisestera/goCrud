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
		{
			Name: "CreateRick",
			Run: func(db *gorm.DB) error {
				return CreateContact(db, "John", 30, "123456789", "Caracas-Venezuela", "email@email.com", "desc1")
			},
		},
		{
			Name: "CreateMike",
			Run: func(db *gorm.DB) error {
				return CreateContact(db, "Mike", 29, "123456789", "Caracas-Venezuela", "email@email.com", "desc2")
			},
		},
	}
}
