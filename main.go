package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

const (
	top250Movies   = "https://imdb-api.com/en/API/Top250Movies/k_hd2hitvi"
	top250TV       = "https://imdb-api.com/en/API/Top250TVs/k_hd2hitvi"
	apiKey         = "ce50bf33b5c5d345eff1afb5ff756c63"
	baseURL        = "https://api.themoviedb.org/3"
	exampleRequest = "https://api.themoviedb.org/3/movie/550?api_key=ce50bf33b5c5d345eff1afb5ff756c63"
	reqq           = baseURL + "/movie/550?api_key=" + apiKey
)

// func GenerateDatabaseFile() {
// 	resp, err := http.Get(top250TV)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
//
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	ioutil.WriteFile("top250TV.json", body, 0o644)
// }

var (
	ID   int
	Type string
)

type Movie struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func StoreMoviesAndIDs() []Movie {
	file, err := os.Open("top250TV.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// TODO read data from json
	var movies []Movie
	json.NewDecoder(file).Decode(&movies)
	fmt.Println(movies)

	return movies
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Welcome)
	moi := StoreMoviesAndIDs()
	fmt.Println(moi)

	// http.ListenAndServe(":8080", r)
}
