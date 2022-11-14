package main

import (
	"net/http"
)

func Welcome(w http.ResponseWriter, _ *http.Request) {
	// apiKey := readAPIKeyFromFile()
}

// type Movie struct {
// 	ReleaseDate string `json:"release_date"`
// 	Title       string `json:"title"`
// 	Overview    string `json:"overview"`
// }

// func (m Movie) String() string {
// 	return fmt.Sprintf("Title: %s\nRelease Date: %s\nOverview: %s", m.Title, m.ReleaseDate, m.Overview)
// }
//
// func buildURL(id string) string {
// 	return fmt.Sprintf("%s/movie/%s?api_key=%s", baseURL, id, apiKey)
// }
//
// func GetMovieData(_ int) Movie {
// 	var movie Movie
// 	resp, err := http.Get(reqq)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer resp.Body.Close()
//
// 	json.NewDecoder(resp.Body).Decode(&movie)
// 	return movie
// }
//
// func MovieInfoHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id := vars["key"]
// 	w.WriteHeader(http.StatusOK)
// 	url := buildURL(id)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	// display response
//
// 	var movie Movie
// 	json.NewDecoder(resp.Body).Decode(&movie)
// 	io.WriteString(w, movie.String())
//
// 	// response with body
// }
