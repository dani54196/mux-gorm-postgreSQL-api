package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/dani54196/mux-gorm-postgreSQL-api/db"
	"github.com/dani54196/mux-gorm-postgreSQL-api/models"
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
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()

	// routes
	r.HandleFunc("/", routes.HomeHandler).Methods("GET")

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, r)
}
