package seeds

import (
	"crud/models"

	"github.com/jinzhu/gorm"
)

// CreateContact function
func CreateContact(db *gorm.DB, name string, age uint, phone string, address string, email string, description string) error {
	return db.Create(&models.Contact{
		Nombre:      name,
		Edad:        age,
		Telefono:    phone,
		Direccion:   address,
		Email:       email,
		Descripcion: description,
	}).Error
}
