package main

import (
	"bufio"
	"log"
	"os"
)

const UserLimit = 5

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
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
