package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	datastore "main/DataStore"

	_ "github.com/gorilla/mux"
)

func main() {
	Movies := ReadJSONfile("./data/top250TV.json")
	fmt.Println(Movies)
	FetchDataFromAPI()
}

func FetchDataFromAPI() datastore.Show {
	apiKey := readAPIKeyFromFile()
	url := "http://www.omdbapi.com/?apikey=" + apiKey + "&i=tt0944947"
	resp, err := http.Get(url)
	CheckError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	var movie datastore.Show
	json.Unmarshal(body, &movie)
	fmt.Println(movie)
	return movie
}
