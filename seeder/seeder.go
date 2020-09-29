package main

import (
	"log"

	"crud/seeds"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbConn := openConnection()
	defer dbConn.Close()

	for _, seed := range seeds.All() {
		if err := seed.Run(dbConn); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(localhost:8889)/gocrud?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}
