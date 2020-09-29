package models

import "github.com/jinzhu/gorm"

// Todo model será un struct que contenga gorm.Model (gorm, ORM para go)

// User model
type User struct {
	gorm.Model
	Name string `json:"nombre"`
	Age  uint   `json:"edad"`
}
