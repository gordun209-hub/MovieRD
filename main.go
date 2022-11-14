package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// simdilk dble urasmayak
// func main2() {
// 	Main()
// 	db := NewDB()
// 	defer db.Close()
// }

const (
	apiKey         = "ce50bf33b5c5d345eff1afb5ff756c63"
	baseURL        = "https://api.themoviedb.org/3"
	exampleRequest = "https://api.themoviedb.org/3/movie/550?api_key=ce50bf33b5c5d345eff1afb5ff756c63"
	reqq           = baseURL + "/movie/550?api_key=" + apiKey
)

var (
	ID   int
	Type string
)

func welcome(w http.ResponseWriter, _ *http.Request) {
	// apiKey := readAPIKeyFromFile()
	mov := getMovieData(3)
	io.WriteString(w, mov.String())
}

type Movie struct {
	ReleaseDate string `json:"release_date"`
	Title       string `json:"title"`
	Overview    string `json:"overview"`
}

func (m Movie) String() string {
	return fmt.Sprintf("Title: %s\nRelease Date: %s\nOverview: %s", m.Title, m.ReleaseDate, m.Overview)
}

func buildURL(id string) string {
	return fmt.Sprintf("%s/movie/%s?api_key=%s", baseURL, id, apiKey)
}

func getMovieData(_ int) Movie {
	var movie Movie
	resp, err := http.Get(reqq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&movie)
	return movie
}

func MovieInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["key"]
	w.WriteHeader(http.StatusOK)
	url := buildURL(id)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// display response

	var movie Movie
	json.NewDecoder(resp.Body).Decode(&movie)
	io.WriteString(w, movie.String())
	// response with body
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", welcome)
	r.HandleFunc("/movie/{key}", MovieInfoHandler)

	http.ListenAndServe(":8080", r)
}
