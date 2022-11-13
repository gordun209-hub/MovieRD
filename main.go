package main

import (
	_ "database/sql"

	_ "github.com/lib/pq"

	model "main/models"
)

func main() {
	// open database
	db := model.NewDB()
	defer db.Close()

	// check db
	db.Ping()

	db.CreateUser("Alihan", "Aydin", []model.Movie{})
	// print all users
	rows := db.Query("SELECT * FROM users")

	defer rows.Close()
}
