package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	Main()
	db := NewDB()
	defer db.Close()
	db.InsertUser("alaai", "aydin")

	users := db.getUsers()
	fmt.Println(users)
}

func readAPIKeyFromFile() string {
	file, err := os.Open("apikey")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
