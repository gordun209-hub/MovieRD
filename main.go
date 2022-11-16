package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	datastore "main/DataStore"

	_ "github.com/gorilla/mux"
)

const (
	top250Movies   = "https://imdb-api.com/en/API/Top250Movies/k_hd2hitvi"
	top250TV       = "https://imdb-api.com/en/API/Top250TVs/k_hd2hitvi"
	apiKey         = "ce50bf33b5c5d345eff1afb5ff756c63"
	baseURL        = "https://api.themoviedb.org/3"
	exampleRequest = "https://api.themoviedb.org/3/movie/550?api_key=ce50bf33b5c5d345eff1afb5ff756c63"
	reqq           = baseURL + "/movie/550?api_key=" + apiKey
)

var (
	ID   int
	Type string
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
