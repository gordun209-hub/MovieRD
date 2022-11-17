package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	datastore "main/DataStore"
)

const UserLimit = 5

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func readAPIKeyFromFile() string {
	file, err := os.Open("./data/apikey")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func ReadJSONfile(filename string) datastore.Show {
	var movie datastore.Show
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &movie)
	return movie
}
