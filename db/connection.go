package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
// alert!!! move to .env
var DSN = "host=localhost user=dani54196 password=example dbname=gorm port=5432"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Database connection successful")
	}
}
