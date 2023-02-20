package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/dani54196/mux-gorm-postgreSQL-api/db"
	"github.com/dani54196/mux-gorm-postgreSQL-api/routes"
)

func main() {
	// enviroment var
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// database connection
	db.DBConnection()

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", routes.HomeHandler)

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
