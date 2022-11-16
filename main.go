package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	datastore "main/DataStore"

	_ "github.com/gorilla/mux"
)

func main() {
	Movies := ReadJSONfile("top250TV.json")
	fmt.Println(Movies.FindShowByID("tt7259746"))
	fmt.Println(Movies.FindShowByTitle("Queer Eye"))
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
