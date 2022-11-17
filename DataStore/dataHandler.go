package datastore

/*
This Module makes Handling data easy and contains structs for data
This is important for getting id's of movies for using in the API
*/
import (
	"encoding/json"
	"fmt"
	"os"
)

// Represents simple show data from the API
type Item struct {
	ID              string `json:"id"`
	Rank            string `json:"rank"`
	Title           string `json:"title"`
	FullTitle       string `json:"fullTitle"`
	Year            string `json:"year"`
	Image           string `json:"image"`
	Crew            string `json:"crew"`
	ImDBRating      string `json:"imDbRating"`
	ImDBRatingCount string `json:"imDbRatingCount"`
}

// Represents a show
type Show struct {
	Items []Item `json:"items"`
}

func (s Show) GetFirstAmount(amount int) []Item {
	items := make([]Item, amount)
	for i := 0; i < amount; i++ {
		items[i] = s.Items[i]
	}
	return items
}

func (s Show) FindShowByID(ID string) string {
	for _, item := range s.Items {
		if item.ID == ID {
			return item.Title
		}
	}
	return ""
}

func (s Show) FindShowByTitle(title string) string {
	for _, item := range s.Items {
		if item.Title == title {
			return item.ID
		}
	}
	return ""
}

// json format
func (s Show) StoreIDsAndTitle() {
	file, err := os.Create("movies.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(s)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (m Show) getMovieData(ID string) string {
	var movieData string
	for _, item := range m.Items {
		if item.ID == ID {
			movieData += fmt.Sprintf("ID: %s, Title: %s\n", item.ID, item.Title)
		}
	}
	return movieData
}

func (m Show) String() string {
	var movieData string
	for _, item := range m.Items {
		movieData += fmt.Sprintf("ID: %s, Title: %s\n", item.ID, item.Title)
	}
	return movieData
}
