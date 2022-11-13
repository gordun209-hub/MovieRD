package main

import (
	_ "bufio"
	_ "database/sql"
	_ "fmt"
	_ "log"
	_ "os"

	_ "github.com/lib/pq"

	_ "main/models"
)

func main() {
	Main()
	//		// open database
	//		db := model.NewDB()
	//		defer db.Close()
	//
	//		apiKey := readAPIKeyFromFile()
	//		fmt.Println(apiKey)
	//		// check db
	//		db.Ping()
	//
	//		// db.CreateUser("Alihan", "Aydin", []model.Movie{})
	//		// print all users
	//		rows := db.Query("SELECT * FROM users")
	//
	//		defer rows.Close()
	//	}
	//
	//	func readAPIKeyFromFile() string {
	//		file, err := os.Open("apikey")
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		defer file.Close()
	//
	//		scanner := bufio.NewScanner(file)
	//		scanner.Scan()
	//		return scanner.Text()
}
