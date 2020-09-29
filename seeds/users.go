package seeds

import (
	"crud/models"

	"github.com/jinzhu/gorm"
)

// CreateContact function
func CreateUser(db *gorm.DB, name string, age uint) error {
	return db.Create(&models.User{
		Name: name,
		Age:  age,
	}).Error
}
