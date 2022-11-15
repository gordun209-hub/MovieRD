package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	Items []struct {
		ID              string `json:"id"`
		Rank            string `json:"rank"`
		Title           string `json:"title"`
		FullTitle       string `json:"fullTitle"`
		Year            string `json:"year"`
		Image           string `json:"image"`
		Crew            string `json:"crew"`
		ImDBRating      string `json:"imDbRating"`
		ImDBRatingCount string `json:"imDbRatingCount"`
	} `json:"items"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Welcome)
	Movies := ReadJSONfile()
	fmt.Println(FormatMovieData(Movies))

	// http.ListenAndServe(":8080", r)
}

func FormatMovieData(m Movie) string {
	var movieData string
	for _, item := range m.Items {
		movieData += fmt.Sprintf("Title: %s\n", item.Title)
	}
	return movieData
}

func ReadJSONfile() Movie {
	var movie Movie
	data, err := ioutil.ReadFile("top250TV.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(data, &movie)
	fmt.Println(movie.Items)
	return movie
}
