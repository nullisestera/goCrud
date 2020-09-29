package seed

import (
	"github.com/jinzhu/gorm"
)

// Seed model
type Seed struct {
	Name string
	Run  func(*gorm.DB) error
}
